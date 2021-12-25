package AttProgram

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["userId"])
	user := getUserById(id)
	json.NewEncoder(writer).Encode(user)
}

func addTrainingResult(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-type", "application/json")
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["userId"])
	user := getUserById(id)
	var result TrainingRecord
	json.NewDecoder(request.Body).Decode(&result)
	user.TrainingResults = append(user.TrainingResults, result)
}

func getTrainingResult(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["userId"])
	user := getUserById(id)
	for _, result := range user.TrainingResults {
		if result.ID == params["trainingId"] {
			json.NewEncoder(writer).Encode(result)
			return
		}
	}
}

func getTrainingsHistory(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["userId"])
	user := getUserById(id)
	json.NewEncoder(writer).Encode(user.TrainingResults)
}
