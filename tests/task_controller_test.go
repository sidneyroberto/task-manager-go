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
	"task-manager-go/outputs"
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

func openTasks(path string) []inputs.CreateTaskInput {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload []inputs.CreateTaskInput
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return payload
}

/* Inicialização dos testes */
func TestMain(m *testing.M) {
	// Antes dos testes
	dsn := configs.EnvMongoURI("TEST_DSN", "../.env")
	configs.ConnectToDatabase(dsn)
	configs.DB.Delete(&models.Task{}, "id > ?", 0)

	// Executa os testes
	code := m.Run()

	// Após os testes
	os.Exit(code)
}

func TestCreateValidTask(t *testing.T) {
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

	content := res.Body.Bytes()
	var taskOutput outputs.CreateTaskOutput
	err = json.Unmarshal(content, &taskOutput)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	task := taskOutput.Task
	t.Log("Task: ", task)
}

func TestCreateInvalidDescriptionTask(t *testing.T) {
	tasksInput := openTask("../data/invalid_description_task.json")
	jsonBody, marshalErr := json.Marshal(tasksInput)
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

	if res.Code != 400 {
		t.Errorf("Expected 400, received %d", res.Code)
	}
}

func TestCreateInvalidSeverityTask(t *testing.T) {
	createTaskInput := openTask("../data/invalid_severity_task.json")
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

	if res.Code != 400 {
		t.Errorf("Expected 400, received %d", res.Code)
	}
}

func TestGetAllTasks(t *testing.T) {
	configs.DB.Delete(&models.Task{}, "id > ?", 0)

	tasksInput := openTasks("../data/tasks.json")
	for _, createTaskInput := range tasksInput {
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
	}

	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		log.Fatal("Error during request: ", err)
	}

	res := httptest.NewRecorder()
	router := routes.SetupRouter()
	router.ServeHTTP(res, req)

	if res.Code != 200 {
		t.Errorf("Expected 200, received %d", res.Code)
	}

	content := res.Body.Bytes()
	var tasksOutput outputs.GetAllTasksOutput
	err = json.Unmarshal(content, &tasksOutput)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	tasks := tasksOutput.Tasks
	size := len(tasks)
	t.Log("Size: ", size)
	if size != 3 {
		t.Errorf("Expected 3, received %d", size)
	}
}
