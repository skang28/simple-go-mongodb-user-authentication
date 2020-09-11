package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/skang28/golang_gin_mongo/form"
	"github.com/skang28/golang_gin_mongo/models"
	"github.com/skang28/golang_gin_mongo/auth"
	"github.com/skang28/golang_gin_mongo/hasher"
)

var userModel = new(models.UserModel)

//AccountController houses controller methods
type AccountController struct{}

//Register function handles registering accounts
func (u *AccountController) Register(c *gin.Context) {
	var data form.NewUser

	//check to see if all fields are provided in request body
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message":"ensure all fields are filled"})
		c.Abort()
		return
	}

	result, _ := userModel.GetUserByEmail(data.Email)

	if result.Email != "" {
		c.JSON(403, gin.H{"message": "Email already in use"})
		c.Abort()
		return
	}

	//check for error when creating account, if there is an error abort request
	err := userModel.Register(data)
	if err != nil {
		c.JSON(400, gin.H{"message":"failed to create account"})
		c.Abort()
		return
	}

	c.JSON(201, gin.H{"message":"account created"})
}

//Login function allows user to login and get jwt token
func (u *AccountController) Login(c *gin.Context) {
	var data form.LoginUser

	if c.BindJSON(&data) !=nil {
		c.JSON(406, gin.H{"message":"Missing details"})
		c.Abort()
		return
	}

	result, err := userModel.GetUserByEmail(data.Email)

	if result.Email == "" {
		c.JSON(404, gin.H{"message":"account not found"})
		c.Abort()
		return
	}

	if err != nil {
		c.JSON(400, gin.H{"message":"error"})
		c.Abort()
		return
	}

	hashedPassword := []byte(result.Password)
	password := []byte(data.Password)
	err = hasher.ComparePasswords(password, hashedPassword)

	if err != nil {
		c.JSON(403, gin.H{"message":"invalid password"})
		c.Abort()
		return
	}

	jwtToken, err2 := auth.GenerateToken(data.Email)

	if err2 != nil {
		c.JSON(403, gin.H{"message":"error"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message":"Successfully logged in","token":jwtToken})
}
