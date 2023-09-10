package services

import (
	"context"
	"go-sample/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	collection *mongo.Collection
}

func NewUserService(database *mongo.Database, collectionName string) *UserService {
	collection := database.Collection(collectionName)

	return &UserService{
		collection: collection,
	}
}
func (us *UserService) CreateUser(user models.User) error {
	newUser := models.NewUser(user.Username, user.Email, user.Password)
	_, err := us.collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Println("Error creating user:", err)
		return err
	}
	return nil
}

func (us *UserService) FindUserByEmail(email string) (models.User, error) {
	var user models.User
	filter := bson.M{"email": email}

	err := us.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.User{}, nil
		}
		log.Println("Error finding user:", err)
		return models.User{}, err
	}

	return user, nil
}

func (us *UserService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
