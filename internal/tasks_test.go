package internal

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	group := r.Group("/api/v1")
	TaskRoutes(group)

	return r
}

func TestCreateTask_HappyPath(t *testing.T) {

	router := setupRouter()

	body := CreateTaskRequest{
		Title:  "Learn Go",
		Status: "todo",
	}

	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/api/v1/tasks",
		bytes.NewBuffer(jsonBody),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusCreated {
		t.Errorf("expected status 200/201 got %d", w.Code)
	}
}

func TestCreateTask_TitleEmpty(t *testing.T) {

	router := setupRouter()

	body := CreateTaskRequest{
		Title:  "",
		Status: "todo",
	}

	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/api/v1/tasks",
		bytes.NewBuffer(jsonBody),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 got %d", w.Code)
	}
}

func TestCreateTask_TitleTooLong(t *testing.T) {

	router := setupRouter()

	longTitle := strings.Repeat("a", 201)

	body := CreateTaskRequest{
		Title:  longTitle,
		Status: "todo",
	}

	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/api/v1/tasks",
		bytes.NewBuffer(jsonBody),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 got %d", w.Code)
	}
}

func TestCreateTask_InvalidStatus(t *testing.T) {

	router := setupRouter()

	body := CreateTaskRequest{
		Title:  "Learn Go",
		Status: "invalid",
	}

	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/api/v1/tasks",
		bytes.NewBuffer(jsonBody),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 got %d", w.Code)
	}
}