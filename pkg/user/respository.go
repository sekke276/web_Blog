package user

import (
	"context"
	"errors"
	"web_Blogs/pkg/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Database
}

func NewUserMongoRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindUserByID(id string) (*entities.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid Id")
	}
	collection := r.db.Collection("Users")
	ctx := context.Background()
	var user entities.User
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *entities.User) error {
	collection := r.db.Collection("Users")
	ctx := context.Background()
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByUserName(username string) (*entities.User, error) {
	collection := r.db.Collection("Users")
	ctx := context.Background()
	var user entities.User
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
