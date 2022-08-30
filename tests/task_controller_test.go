package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"task-manager-go/configs"
	"task-manager-go/inputs"
	"task-manager-go/models"
	"task-manager-go/routes"
	"testing"
)

func openTask(path string) inputs.CreateTaskInput {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload inputs.CreateTaskInput
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return payload
}

/* Inicialização dos testes */
func TestMain(m *testing.M) {
	dsn := configs.EnvMongoURI("TEST_DSN", "../.env")
	configs.ConnectToDatabase(dsn)
	configs.DB.Delete(&models.Task{}, "id > ?", 0)

	code := m.Run()

	os.Exit(code)
}

func TestCreateTask(t *testing.T) {
	createTaskInput := openTask("../data/task.json")
	jsonBody, marshalErr := json.Marshal(createTaskInput)
	if marshalErr != nil {
		log.Fatal("Error during Marshal(): ", marshalErr)
	}

	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("POST", "/tasks", bodyReader)
	if err != nil {
		log.Fatal("Error during request: ", err)
	}

	res := httptest.NewRecorder()
	router := routes.SetupRouter()
	router.ServeHTTP(res, req)

	if res.Code != 201 {
		t.Errorf("Expected 201, received %d", res.Code)
	}
}
