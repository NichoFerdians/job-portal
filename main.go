package main

import (
	"github.com/NichoFerdians/job-portal/db"
	routes "github.com/NichoFerdians/job-portal/routes"
)

func main() {
	db.Connect()
	r := routes.SetupRoutes()
	r.Run()
}
