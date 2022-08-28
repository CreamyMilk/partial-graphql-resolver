package graph

import (
	"example/graph/model"
	"fmt"
	"log"
)

func AUser(randomId string) *model.User {
	x := model.User{
		ID:   fmt.Sprint(randomId),
		Name: "People Who win",
	}

	return &x
}

func GetTodos() []*model.Todo {
	//	log.Println("✅ Somone called this func")
	t := []*model.Todo{
		{
			ID:   "1",
			Text: "This is atodo",
			Done: true,
		},
		{
			ID:   "1",
			Text: "This is atodo",
			Done: true,
		},
	}

	return t
}
func GetPeople(userId string) []*model.People {
	log.Println("🌈 User is", userId)

	p := []*model.People{
		{
			ID:   userId,
			Name: "This person",
		},
		{
			ID:   "defualt",
			Name: "This is another person",
		},
	}

	return p
}
