package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/NichoFerdians/job-portal/db"
	"github.com/NichoFerdians/job-portal/helpers"
	"github.com/NichoFerdians/job-portal/models"
	"github.com/NichoFerdians/job-portal/token"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	var input models.LoginUserInput
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Con.Where("user_name = ?", input.UserName).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	err := helpers.CheckPassword(input.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errorc": err.Error()})

		return
	}

	SECRET := os.Getenv("JWT_SECRET")

	tokenMaker, err := token.NewJWTMaker(SECRET)
	if err != nil {
		log.Fatalf("cannot create token maker: %s", err)
	}

	duration := time.Hour
	accessToken, err := tokenMaker.CreateToken(user.UserName, duration)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rsp := models.LoginUserResponse{
		AccessToken: accessToken,
	}

	c.JSON(http.StatusOK, gin.H{"data": rsp})
}
