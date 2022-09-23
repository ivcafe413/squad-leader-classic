package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"

	// "io"
	"net/http/httptest"
	"testing"

	//"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
	//"github.com/valyala/fasthttp/fasthttputil"

	// "github.com/gofiber/websocket/v2"
	"github.com/gorilla/websocket"
	"github.com/vagrant-technology/squad-leader/auth"
	// "github.com/vagrant-technology/squad-leader/game"
)

func Test_CreateRoomRoute(t *testing.T) {
	fmt.Println("--------------------")
	fmt.Println("Test 1: Create Room")
	fmt.Println("--------------------")

	app := fiber.New()

	ConfigureApi(app)

	// ----- Create Room -----
	testUser := struct {
		Username string `json:"username"`
	}{
		Username: "TestUser",
	}
	//fmt.Println(testUser)
	jsonUser, _ := json.Marshal(testUser)
	// fmt.Println(jsonUser)

	postReq := httptest.NewRequest("POST", "/api/CreateRoom", bytes.NewReader(jsonUser))
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
		Room string     `json:"room"`
		User auth.User `json:"user"`
	}{}
	jsonErr := json.NewDecoder(postResp.Body).Decode(&respJson)
	if jsonErr != nil {
		t.Fatal(jsonErr)
	}

	fmt.Println("Room ID: " + respJson.Room)
	fmt.Println("Username: " + respJson.User.Username)
}

func Test_CreateAndJoinRoom(t *testing.T) {
	fmt.Println("--------------------")
	fmt.Println("Test 2: Create and Join Room")
	fmt.Println("--------------------")

	//Testing Config
	testConfig := fiber.Config{
		DisableStartupMessage: true,
	}

	// Basic API/WS Config
	app := fiber.New(testConfig)

	ConfigureWS(app)
	ConfigureApi(app)

	//ln := fasthttputil.NewInmemoryListener()
	ln, _ := net.Listen("tcp", "localhost:3001")
	go func() {
		_ = app.Listener(ln)
	}()

	defer func() {
		_ = app.Shutdown()
		_ = ln.Close()
	}()

	// ----- Step 1: Create and verify Game Room -----
	u := struct {
		Username string `json:"username"`
	}{Username: "Alpha Dog"}
	payload, _ := json.Marshal(u)

	cReq := httptest.NewRequest("POST", "/api/CreateRoom", bytes.NewReader(payload))
	cReq.Header.Set("Content-Type", "application/json")

	cRes, err := app.Test(cReq)

	if err != nil {
		t.Fatal(err.Error())
	}

	defer cRes.Body.Close()

	roomPayload := struct {
		Room string     `json:"room"`
		User auth.User `json:"user"`
	}{}

	if err := json.NewDecoder(cRes.Body).Decode(&roomPayload); err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println("Room ID: " + roomPayload.Room)
	fmt.Println("Username: " + roomPayload.User.Username)

	// ----- Step 2: Join the Room Lobby -----

	//testUrl := "ws://localhost:3000/ws"
	wsUrl := "ws://" + ln.Addr().String() + "/ws/" + roomPayload.Room + "/" + roomPayload.User.Username
	fmt.Println("WS Join URL: " + wsUrl)

	ws, _, joinErr := websocket.DefaultDialer.Dial(wsUrl, nil)
	//joinReq := httptest.NewRequest("GET", "/ws/" + roomPayload.Room + "/" + fmt.Sprint(roomPayload.User.ID), nil)
	//joinRes, joinErr := app.Test(joinReq)

	if joinErr != nil {
		t.Fatal(joinErr.Error())
	}

	defer ws.Close()

	// ----- Step 3: Send a test message -----
	if err := ws.WriteMessage(websocket.TextMessage, []byte("ready")); err != nil {
		t.Fatalf("%v", err)
	}

	// Receive the response
	_, reply, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("%v", err)
	}
	fmt.Println(string(reply))

	lobbyState := make(map[string]bool)
	err = json.Unmarshal(reply, &lobbyState)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(lobbyState)

	// Deferred closes
}
