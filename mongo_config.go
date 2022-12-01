package gobe

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnector struct {
	*mongo.Database
}

type MongoBaseConfig struct {
	DBName     string `mapstructure:"db_name" json:"db_name"`
	DBHost     string `mapstructure:"db_host" json:"db_host"`
	DBPort     string `mapstructure:"db_port" json:"db_port"`
	DBUsername string `mapstructure:"db_username" json:"db_username"`
	DBPassword string `mapstructure:"db_password" json:"db_password"`
	URI        string `mapstructure:"uri" json:"uri"`
}

// Initialize new connection to MongoDB
func NewMongoConfig(config *MongoBaseConfig) MongoConnector {
	uri := config.URI
	if uri == "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort)
	}
	client := options.Client()
	client.ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, client)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Connected!")
	}

	db := mongoClient.Database(config.DBName)
	return MongoConnector{db}

}
