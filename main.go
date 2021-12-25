package AttProgram

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main(args string)  {
	router := mux.NewRouter()
	router.HandleFunc("/users/{userId}/profile", getUser).Methods("GET")
	router.HandleFunc("/users/{userId}/trainings", addTrainingResult).Methods("POST")
	router.HandleFunc("/users/{userId}/trainings", getTrainingsHistory).Methods("GET")
	router.HandleFunc("/users/{userId}/trainings/{trainingId}/result", getTrainingResult).Methods("GET")
	http.ListenAndServe(":5000", router)
}