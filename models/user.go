package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/skang28/golang_gin_mongo/form"
	"github.com/skang28/golang_gin_mongo/hasher"
)

//User struct defines the user object
type User struct {
	ID int `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
	Email string `json:"email" bson:"email"`
	IsVerified bool `json:"is_verified" bson:"is_verified"`
}

//UserModel defines the model structure, houses methods
type UserModel struct{}

//Register function handles registration
func (u *UserModel) Register(data form.NewUser) error {
    // Connect to the user collection
    collection := dbConnect.Use(databaseName, "user")
    // Assign result to error object while saving user
    err := collection.Insert(bson.M{
        "name": data.Name,
		"password": hasher.GenerateHashedPassword([]byte(data.Password)),
		"email": data.Email,
        "is_verified": false,
    })

    return err
}

//GetUserByEmail fetches user by email
func (u *UserModel) GetUserByEmail(email string) (user User, err error) {
	collection := dbConnect.Use(databaseName, "user")
	err = collection.Find(bson.M{"email":email}).One(&user)
	return user, err
}