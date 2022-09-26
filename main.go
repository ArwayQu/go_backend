package main

import (
	"ginDemo/routers"
	"net/http"

	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database("tasker").Collection("tasks")

	r := gin.Default()

	var router *gin.Engine
	router = routers.InitRouter()

	s := &http.Server{
		Addr:    "127.0.0.1:8899",
		Handler: router,
	}
	s.ListenAndServe()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
