package controllers

import (
	"context"
	"github.com/fixwa/go-prices-tracker/database"
	"github.com/fixwa/go-prices-tracker/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"
)

func init() {
	database.ConnectDatabase()
}

func ListProducts(c *gin.Context) {
	products := getProducts()

	c.JSON(http.StatusOK, products)
}

func getProducts() []models.Product {
	productsCollection := database.Db.Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	cursor, err := productsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var result []models.Product
	if err = cursor.All(ctx, &result); err != nil {
		log.Fatal(err)
	}

	return result
}
