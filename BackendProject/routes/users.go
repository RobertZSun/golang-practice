package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/backend/models"
	"github.com/backend/utils"
)

func signupUser(context *gin.Context) {
	var user models.User

	error := context.ShouldBindJSON(&user)

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse signup data.",
			"error":   error,
		})
		return
	}

	result, SaveErr := user.Save()

	if SaveErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save user.",
			"error":   SaveErr,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created.",
		"user":    result,
	})

}

func loginUser(context *gin.Context) {
	var user models.User

	error := context.ShouldBindJSON(&user)

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse login data.",
			"error":   error,
		})
		return
	}

	validateError := user.ValidateCredentials()

	if validateError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not login.",
			"error":   validateError,
		})
		return
	}

	token, tokenError := utils.GenerateToken(user.Email, user.ID)

	if tokenError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not generate token.",
			"error":   tokenError,
		})
		return

	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login success.",
		"token":   token,
	})

}
