package AttProgram

import (
	"math/rand"
	"time"
)

func getUserById(userId int) User {
	fakeUsers := [2]User{}
	var foundUser User
	for _, user := range  fakeUsers {
		if user.Id == userId {
			foundUser = user
			break
		}
	}
	return  foundUser
}

func  PopulateUsersList(count int) []User  {
	var usersCollection []User
	for i:=0; i<count; i++   {
		user := User{Id: i, TrainingResults: CreateTrainingRecords(i),
			Username: "User" + string(i)}
		usersCollection = append(usersCollection, user)
	}
	return usersCollection
}

func CreateTrainingRecords(userId int) []TrainingRecord {
	var resultsCollection []TrainingRecord
	recCount := rand.Intn(10)
	for i:=0; i<=recCount; i++ {
		rec := TrainingRecord{ID:string(userId*11 + i),
			Comment: "This is a training" + string(i),
			TrainingResult: CreateResult(i)}
		resultsCollection = append(resultsCollection, rec)
	}
	return resultsCollection
}

func CreateResult(recNumber int) Results {
	var exr []string
	if recNumber%2 == 0 {
		exr = []string{"Бегит","Пресс качат","Анжумания"}
	} else {
		exr = []string{"Бегит", "Приседат", "Пресс качат", "Анжумания"}
	}
	res := Results{BurnedCalories: rand.Intn(600),
		TrainingDuration: time.Duration(rand.Intn(60))*time.Minute,
		Exercises: exr}
	return res
}
