package mongodb

import (
	"github.com/bobomurod/go-auth-bolilerplate/internal/core/ports"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	client *mongo.Client
	db     *mongo.Database
	logger ports.Logger
}

func NewUserRepository(client *mongo.Client) *MongoUserRepository {
	return &MongoUserRepository{client: client}
}
