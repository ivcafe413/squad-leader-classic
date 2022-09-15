package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/vagrant-technology/squad-leader/game"
)

func Test_CreateRoomRoute(t *testing.T) {
	app := fiber.New()

	Router(app)

	// ----- Test 1 - Basic Routing 200
	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req, 1)

	if resp.StatusCode != 200 {
		t.Fatal("Basic Root GET request failing - probably app/router not initializing properly")
	}

	defer resp.Body.Close()

	respString, err := io.ReadAll(resp.Body)
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
	jsonUser, _ := json.Marshal(testUser)
	req = httptest.NewRequest("POST", "/CreateRoom", bytes.NewBuffer(jsonUser))
	resp, err = app.Test(req, 1)

	if resp.StatusCode != 200 {
		fmt.Println(err)
		t.Fatal(err)
	}

	defer resp.Body.Close()

	//respString, err = io.ReadAll(resp.Body)
	respRoom := new(game.Room)
	json.NewDecoder(resp.Body).Decode(respRoom)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Room ID: " + respRoom.ID.String())
}
