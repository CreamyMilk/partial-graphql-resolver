package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/graph/generated"
	"example/graph/model"
	"fmt"
)

// GetUsers is the resolver for the getUsers field.
func (r *queryResolver) GetUsers(ctx context.Context, id string) ([]*model.User, error) {
	u := make([]*model.User, 5)
	for i := 0; i < 5; i++ {
		u[i] = AUser(fmt.Sprint(i))
	}

	return u, nil
}

// Todos is the resolver for the todos field.
func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	//log.Printf("🥹 The current User is %+v", obj)

	return GetTodos(), nil
}

// People is the resolver for the people field.
func (r *userResolver) People(ctx context.Context, obj *model.User) ([]*model.People, error) {
	//log.Printf("🙋 The current User is %+v", obj)

	return GetPeople(obj.ID), nil
}

// Inflow is the resolver for the inflow field.
func (r *userResolver) Inflow(ctx context.Context, obj *model.User) (int, error) {
	return 120, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
