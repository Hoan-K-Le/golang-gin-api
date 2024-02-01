package controllers

import (
	"context"
    "fmt"
    "log"

    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "github.com/Hoan-K-Le/golang-gin-api-ecom/configs"

    "github.com/Hoan-K-Le/golang-gin-api-ecom/models"
    helper "github.com/Hoan-K-Le/golang-gin-api-ecom/helpers"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
)
var userCollection *mongo.Collection = configs.OpenCollection(configs.Client, "user")
var validate = validator.New()

// HashPassword used to encrypt the password before it is stored in the DB
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

// Verify the password
func VerifyPassword(userPassword string, providedPassword string) (bool,string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("Login or password is incorrect")
		check = false
	}
	return check, msg
}

func SignUp() gin.HandlerFunc {
    return func(c *gin.Context) {
        var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
        var user models.User

        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        validationErr := validate.Struct(user)
        if validationErr != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
            return
        }

        count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
        defer cancel()
        if err != nil {
            log.Panic(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
            return
        }

        password := HashPassword(*user.Password)
        user.Password = &password

        if count > 0 {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "this email"})
            return
        }

        user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
        user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
        user.ID = primitive.NewObjectID()
        token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.Username, user.ID.Hex())
        user.Token = &token
        user.Refresh_token = &refreshToken

        resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
        if insertErr != nil {
            msg := fmt.Sprintf("User item was not created")
            c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
            return
        }
        defer cancel()

        c.JSON(200, resultInsertionNumber)

    }
}

//Login is the api used to get a single user
func Login() gin.HandlerFunc {
    return func(c *gin.Context) {
        var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
        var user models.User
        var foundUser models.User
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
        defer cancel()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "login or passowrd is incorrect"})
            return
        }

        passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
        defer cancel()
        if passwordIsValid != true {
            c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
            return
        }
       

        token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.Username, foundUser.ID.Hex())
        helper.UpdateAllTokens(token, refreshToken, foundUser.ID.Hex())
        c.SetCookie("auth_token", token, 3600, "/", "localhost", false, true)


        c.JSON(http.StatusOK, gin.H{"token":token, "foundUser":foundUser})

    }
}