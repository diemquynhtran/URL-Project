package controllers

import (
	"fmt"
	"learning-go/database"
	"learning-go/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize()  {
	server.DB = database.InitDb()
	server.DB.AutoMigrate(&models.Url{})
	server.Router = mux.NewRouter()
	server.initializeRoutes()
	server.Run(":8080")
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port " + addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}