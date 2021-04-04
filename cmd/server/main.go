package main

import (
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gofiber/fiber/v2"
	"github.com/yogparra/go-fiber-rest-api.git/config"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance
const dbName = "products"
const mongoURI = "mongodb://user:password@localhost:27017/" + dbName + "?authSource=admin"

type Products struct {
	ID     			string  `json:"id,omitempty" bson:"_id,omitempty"`
	Brand   		string 	`json:"brand"`
	Description		string 	`json:"description"`
	Image			string 	`json:"image"`
	Price 			float64 `json:"price"`
}

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

func main() {

	Println("USERNAME" + config.Config("USERNAME"))
	Println("PASSWORD" + config.Config("PASSWORD"))

	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/api/v1/products", func(c *fiber.Ctx) error {
		
		query := bson.D{{}}
		cursor, err := mg.Db.Collection("products").Find(c.Context(), query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		var products []Products = make([]Products, 0)

		if err := cursor.All(c.Context(), &products); err != nil {
			return c.Status(500).SendString(err.Error())

		}
	
		return c.JSON(products)
	})

	log.Fatal(app.Listen(":3000"))
}
