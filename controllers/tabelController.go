package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/novitaekaari/restraurant-management/database"
	"github.com/novitaekaari/restraurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var tabelCollection *mongo.Collection = database.OpenCollection(database.Client, "tabel")

func GetTabels() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing tabel items"})
		}
		var allTabels []bson.M
		if err = result.All(ctx, &allTabels); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allTabels)
	}
}
func GetTabel() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		tabelId := c.Param("tabel_id")
		var tabel models.Tabel

		err := tabelCollection.FindOne(ctx, bson.M{"tabel_id": tabelId}).Decode(&tabel)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the tabels"})
		}
		c.JSON(http.StatusOK, tabel)
	}
}
func CreateTabel() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var tabel models.Tabel

		if err := c.BindJSON(&tabel); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(tabel)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return

		}

		tabel.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		tabel.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		tabel.ID = primitive.NewObjectID()
		tabel.Tabel_id = tabel.ID.Hex()

		result, insertErr := tabelCollection.InsertOne(ctx, tabel)

		if insertErr != nil {
			msg := fmt.Sprintf("Tabel item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()

		c.JSON(http.StatusOK, result)

	}
}

func UpdateTabel() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var tabel models.Tabel

		tabelId := c.Param("tabel_id")

		if err := c.BindJSON(&tabel); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var updateObj primitive.D

		if tabel.Number_of_guests != nil {
			updateObj = append(updateObj, bson.E{"number_of_guests", tabel.Number_of_guests})
			return
		}

		if tabel.Tabel_number != nil {
			updateObj = append(updateObj, bson.E{"tabel_number", tabel.Tabel_number})
		}

		tabel.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		upsert := true

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		filter := bson.M{"tabel_id": tabelId}

		result, err := tabelCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			&opt,
		)

		if err != nil {
			msg := fmt.Sprintf("table item update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})

		}

		defer cancel()
		c.JSON(http.StatusOK, result)

	}

}
