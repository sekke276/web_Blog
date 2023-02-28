package mongo_repository

import (
	"context"
	"errors"
	"web_Blogs/pkg/entities"
	"web_Blogs/pkg/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryMongo struct {
	collection *mongo.Collection
}

func NewUserMongoRepository(db *mongo.Database) repository.UserRepository {
	return &UserRepositoryMongo{
		collection: db.Collection("users"),
	}
}

func (r *UserRepositoryMongo) FindUserByID(id string) (*entities.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid Id")
	}
	ctx := context.Background()
	var user entities.User
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryMongo) Create(user *entities.UserRequest) error {
	ctx := context.Background()
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepositoryMongo) GetUserByUsername(username string) (*entities.User, error) {
	ctx := context.Background()
	var user entities.User
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
