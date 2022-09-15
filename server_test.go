package main

import (
	"testing"
	"net/http/httptest"
	"github.com/gofiber/fiber/v2"
)

func Test_CreateRoomRoute(t *testing.T) {
	app := fiber.New()

	Router(app)

	//Test 1 - Basic Routing 200
	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req, 1)
}