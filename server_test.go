package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	// "io"
	"net/http/httptest"
	"testing"

	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/websocket/v2"
	"github.com/gorilla/websocket"
	"github.com/vagrant-technology/squad-leader/game"
	"github.com/vagrant-technology/squad-leader/store"
	// "github.com/vagrant-technology/squad-leader/game"
)

func Test_CreateRoomRoute(t *testing.T) {
	app := fiber.New()

	Router(app)

	// ----- Test 1 - Basic Routing 200
	// getReq := httptest.NewRequest("GET", "/", nil)
	// getResp, _ := app.Test(getReq, 1)

	// if getResp.StatusCode != 200 {
	// 	t.Fatal("Basic Root GET request failing - probably app/router not initializing properly")
	// }

	// defer getResp.Body.Close()

	// respString, err := io.ReadAll(getResp.Body)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// fmt.Println(string(respString))

	// ----- Create Room -----
	testUser := struct {
		Username string `json:"username"`
	}{
		Username: "TestUser",
	}
	//fmt.Println(testUser)
	jsonUser, _ := json.Marshal(testUser)
	// fmt.Println(jsonUser)

	postReq := httptest.NewRequest("POST", "/CreateRoom", bytes.NewReader(jsonUser))
	postReq.Header.Set("Content-Type", "application/json") //Necessary

	postResp, postErr := app.Test(postReq)
	fmt.Println("Post Response: " + fmt.Sprint(postResp.StatusCode))

	if postErr != nil {
		fmt.Println(postErr)
		t.Fatal(postErr)
	}

	if postResp.StatusCode != fiber.StatusOK {
		//fmt.Println(err.Error())
		t.Fatal("Not a 200 Post")
	}

	defer postResp.Body.Close()

	//respString, err = io.ReadAll(postResp.Body)
	//respRoom := new(game.Room)
	respJson := struct {
		Room string `json:"room"`
		User store.User `json:"user"`
	}{}
	jsonErr := json.NewDecoder(postResp.Body).Decode(&respJson)
	if jsonErr != nil {
		t.Fatal(jsonErr)
	}

	fmt.Println("Room ID: " + respJson.Room)
	fmt.Println("Username: " + respJson.User.Username)
	fmt.Println("--------------------")
}

func Test_CreateAndJoinRoom(t * testing.T) {
	// Basic API/WS Config
	app := fiber.New()
	ConfigureWS(app)
	go game.ClientHub()
	Router(app)

	// ----- Step 1: Create and verify Game Room -----
	u := struct {
		Username string `json:"username"`
	} { Username: "Alpha Dog" }
	payload, _ := json.Marshal(u)

	cReq := httptest.NewRequest("POST", "/CreateRoom", bytes.NewReader(payload))
	cReq.Header.Set("Content-Type", "application/json")
	
	cRes, err := app.Test(cReq)

	if err != nil {
		t.Fatal(err.Error())
	}

	defer cRes.Body.Close()

	roomPayload := struct {
		Room string `json:"room"`
		User store.User `json:"user"`
	}{}
	
	if err := json.NewDecoder(cRes.Body).Decode(&roomPayload); err!= nil {
		t.Fatal(err.Error())
	}

	fmt.Println("Room ID: " + roomPayload.Room)
	fmt.Println("Username: " + roomPayload.User.Username)

	// ----- Step 2: Join the Room -----
	
	//joinReq := httptest.NewRequest("GET", "/ws/" + roomPayload.Room + "/" + fmt.Sprint(roomPayload.User.ID), nil)
	joinRes, joinErr := app.Test(joinReq)

	if joinErr != nil {
		t.Fatal(joinErr.Error())
	}

	if joinRes.StatusCode != fiber.StatusOK {
		t.Fatalf("Response is %v, not 200", joinRes.StatusCode)
	}

	fmt.Println("Connection: " + joinRes.Header.Get("Connection"))
	fmt.Println("Upgrade: " + joinRes.Header.Get("Upgrade"))

	defer joinRes.Body.Close()
}