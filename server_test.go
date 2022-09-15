package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/vagrant-technology/squad-leader/store"
	// "github.com/vagrant-technology/squad-leader/game"
)

func Test_CreateRoomRoute(t *testing.T) {
	app := fiber.New()

	Router(app)

	// ----- Test 1 - Basic Routing 200
	getReq := httptest.NewRequest("GET", "/", nil)
	getResp, _ := app.Test(getReq, 1)

	if getResp.StatusCode != 200 {
		t.Fatal("Basic Root GET request failing - probably app/router not initializing properly")
	}

	defer getResp.Body.Close()

	respString, err := io.ReadAll(getResp.Body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(respString))

	// ----- Test 2 - Create Room Route 200
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
}
