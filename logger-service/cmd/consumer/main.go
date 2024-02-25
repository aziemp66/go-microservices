package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log-service/internal/data"
	"log-service/internal/rabbitmq"
	"log-service/internal/validation"
	"os"
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

var client *mongo.Client

const (
	mongoURL    = "mongodb://mongo:27017"
	rabbitmqURL = "amqp://user:password@rabbitmq:5672/"
)

func main() {
	en := en.New()
	id := id.New()
	uni := ut.New(en, id)

	trans, isFound := uni.GetTranslator("id")
	if !isFound {
		panic(errors.New("translator not found"))
	}

	v := validator.New()
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("id"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})
	// en_translations.RegisterDefaultTranslations(v, trans)
	id_translations.RegisterDefaultTranslations(v, trans)

	validate := &validation.Validate{
		Validator: v,
		Trans:     trans,
	}

	// connect to mongo
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}
	client = mongoClient

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	models := data.New(client)

	consumer, err := rabbitmq.NewConsumer(rabbitmqURL, &models, validate)
	if err != nil {
		log.Panic(err)
	}

	defer consumer.Channel.Close()
	defer consumer.Connection.Close()

	consumer.Consume(context.Background())

	fmt.Println("Running")
}

func connectToMongo() (*mongo.Client, error) {
	// create connection options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
	})

	// connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting:", err)
		return nil, err
	}

	return c, nil
}
