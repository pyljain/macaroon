package main

import (
	"fmt"
	"log"
	"mongomcp/pkg/tools"
	"net/http"
	"os"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	port := os.Getenv("HOST_PORT")
	mongoConnectionString := os.Getenv("MONGO_CONNECTION_STRING")

	opts := options.Client().ApplyURI(mongoConnectionString).
		SetMaxPoolSize(100).
		SetMaxConnIdleTime(30 * time.Second)

	mongoClient, err := mongo.Connect(opts)
	if err != nil {
		log.Printf("Unable to connect to MongoDB: %v", err)
		os.Exit(-1)
	}

	server := mcp.NewServer(&mcp.Implementation{Name: "macaroon", Version: "v1.0.0"}, nil)

	mcp.AddTool(server, &mcp.Tool{Name: "get-collections", Description: "This can be used to get a natural language description of the collecitons that exist"}, tools.GetCollections(mongoClient))
	mcp.AddTool(server, &mcp.Tool{Name: "get-schema-for-collection", Description: "This can be used to get a natural language description of keys / schema in a collection"}, tools.GetCollectionSchema(mongoClient))
	// mcp.AddTool(server, &mcp.Tool{Name: "get-exemplar-sample-queries", Description: "Exemplar examples of user questions in natural language & corresponding queries and results for reference"}, tools.GetExemplars)
	mcp.AddTool(server, &mcp.Tool{Name: "find", Description: "Run a find Mongo query for a collection"}, tools.Find(mongoClient))
	mcp.AddTool(server, &mcp.Tool{Name: "find-one", Description: "Run a findOne Mongo query for a collection"}, tools.FindOne(mongoClient))
	mcp.AddTool(server, &mcp.Tool{Name: "run-aggregate-query", Description: "Run a Mongo aggregate query"}, tools.AggregateQuery(mongoClient))

	handler := mcp.NewStreamableHTTPHandler(func(req *http.Request) *mcp.Server {
		log.Printf("New request from %s", req.RemoteAddr)
		return server
	}, nil)

	handlerWithLogging := loggingHandler(handler)

	serverURL := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("Listening on %s", serverURL)

	if err := http.ListenAndServe(serverURL, handlerWithLogging); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
