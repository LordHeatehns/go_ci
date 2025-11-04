package handler_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/stretchr/testify/assert"

	"go_ci/handler"
	"go_ci/service"
)

func TestGetUsers_Success(t *testing.T) {
	app := fiber.New()

	mockSv := new(service.MockService)
	mockUsers := []map[string]interface{}{
		{"id": 1, "name": "John"},
	}

	mockSv.On("GetUsers").Return(mockUsers, nil)

	h := handler.NewHandlerMock(mockSv)

	app.Get("/users", h.GetUsers)

	req := httptest.NewRequest("GET", "/users", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)

	var body map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		log.Error(err)
	}

	assert.Equal(t, true, body["success"])
	assert.Equal(t, "success", body["message"])
	assert.Equal(t, float64(200), body["code"])
	assert.Len(t, body["data"].([]interface{}), 1)

	mockSv.AssertExpectations(t)
}
