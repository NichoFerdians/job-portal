package routes

import (
	"log"
	"os"

	ctrls "github.com/NichoFerdians/job-portal/controllers"
	"github.com/NichoFerdians/job-portal/token"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

// var tokenMaker token.Maker

func SetupRoutes() *gin.Engine {

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	SECRET := os.Getenv("JWT_SECRET")

	tokenMaker, err := token.NewJWTMaker(SECRET)
	if err != nil {
		log.Fatalf("cannot create token maker: %s", err)
	}

	v1 := router.Group("/api")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", ctrls.LoginUser)
		}

		orders := v1.Group("/v1").Use(ctrls.AuthMiddleware(tokenMaker))
		{
			orders.GET("/job/list", ctrls.JobList)
			orders.GET("/job/detail", ctrls.JobDetail)
		}
	}

	return router
}
