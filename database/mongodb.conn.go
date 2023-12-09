package database

import (
	"context"
	"cv_api/initializers"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	usr  string
	pwd  string
	host string
	// port     string
	database string
	client   *mongo.Client
)

func GetCollection(collection string) *mongo.Collection {

	initializers.LoadEnv()

	usr = os.Getenv("DB_USER")
	pwd = os.Getenv("DB_PASS")
	host = os.Getenv("DB_HOST")
	database = os.Getenv("DB_NAME")

	// // uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", usr, pwd, host, port)
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", usr, pwd, host)

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	var err error
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database(database).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client.Database(database).Collection(collection)
}
