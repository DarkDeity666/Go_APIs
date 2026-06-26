package repositories

import (
	"context"
	"github.com/DarkDeity666/Go_APIs/internal/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

//this line is always needed as this has so be passed in every database funcion
//eg
// func FindByEmail(Client *mongo.Client,email string){}
// func CreateUser(Client *mongo.Client,user User){}
// func Delete(client * mongo.Client, id string){}

// note how many times we are passing client here so we can use struct

type UserRepository struct {
	client *mongo.Client
}


func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{
		client: client,
	}
}

func (r *UserRepository) collection() *mongo.Collection {
	return r.client.
		Database("social").
		Collection("users")
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	return nil, nil
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	return nil
}