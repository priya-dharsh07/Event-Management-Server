package controllers

import (
	"context"
	"net/http"
	"time"

	"event-management/config"
	"event-management/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var eventCollection *mongo.Collection

func InitEventController() {
	eventCollection = config.GetCollection("events")
}

func CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.BindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.CreatedAt = time.Now()
	_, err := eventCollection.InsertOne(context.TODO(), event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating event"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully"})
}

func GetEvents(c *gin.Context) {
	cursor, err := eventCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching events"})
		return
	}
	var events []models.Event
	if err = cursor.All(context.TODO(), &events); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding events"})
		return
	}
	c.JSON(http.StatusOK, events)
}
