package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"heading": "Pokedex",
	})
}

func main() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/data/pokemon")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
		return
	}
	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("Failed to ping MongoDB:", err)
		return
	}
	fmt.Println("Connected to MongoDB!")
	pokemon := client.Database("data").Collection("pokemon")
	router := gin.Default()
	router.LoadHTMLGlob("../views/*")
	router.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/index"
		router.HandleContext(c)
	})
	router.GET("/index", newPage)
	router.GET("/data", func(c *gin.Context) {
		var results []bson.M
		jsonData, err := json.Marshal(results)
		cursor, err := pokemon.Find(context.TODO(), bson.D{})
		if err != nil {
			panic(err)
		}
		if err = cursor.All(context.TODO(), &results); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, jsonData)
	})
	router.Run(":42069")
}
