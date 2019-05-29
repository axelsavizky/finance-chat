package users

import (
	"finance-chat/src/domain/errors"
	"finance-chat/src/domain/users/entities"
	"finance-chat/src/domain/users/repository/database"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)



func HandleLogin(c *gin.Context) {
	loginRequest := entities.UserRequest{}

	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, entities.LoginErrorResponse{Error:"The request is not valid"})
		return
	}

	hashedPassword, err := database.Login(loginRequest.Username)
	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			c.JSON(http.StatusNotFound, entities.LoginErrorResponse{Error: err.Error()})
		case errors.ErrBadRequest:
			c.JSON(http.StatusBadRequest, entities.LoginErrorResponse{Error: err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, entities.LoginErrorResponse{Error: "An error has occurred"})
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(loginRequest.Password))
	if err != nil {
		c.JSON(http.StatusNotFound, entities.LoginErrorResponse{Error:"username or password is wrong"})
		return
	}

	c.Status(http.StatusAccepted)
}

func HandleSignup(c *gin.Context) {
	signupRequest := entities.UserRequest{}

	err := c.ShouldBindJSON(&signupRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, entities.LoginErrorResponse{Error:"the request is not valid"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.LoginErrorResponse{Error:"there was a problem handling the request"})
		return
	}

	err = database.Signup(signupRequest.Username, string(hashedPassword))
	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			c.JSON(http.StatusNotFound, entities.LoginErrorResponse{Error: err.Error()})
		case errors.ErrBadRequest:
			c.JSON(http.StatusBadRequest, entities.LoginErrorResponse{Error: err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, entities.LoginErrorResponse{Error: "An error has occurred"})
		}
		return
	}

	c.Status(http.StatusCreated)
}
