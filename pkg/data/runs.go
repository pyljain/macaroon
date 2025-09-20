package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Run struct {
	RunID         int       `bson:"RunID"`
	Timestamp     time.Time `bson:"timestamp"`
	UserID        string    `bson:"userID"`
	ModelUsed     string    `bson:"modelUsed"`
	InputTokens   int       `bson:"inputTokens"`
	OutputTokens  int       `bson:"outputTokens"`
	TotalTime     float64   `bson:"totalTime"`
	TotalSpend    float64   `bson:"totalSpend"`
	FilesAttached bool      `bson:"filesAttached"`
	HttpCode      int       `bson:"httpCode"`
}

func randomUserID() string {
	return fmt.Sprintf("user_%d", rand.Intn(50)+1)
}

func randomModel() string {
	models := []string{"gpt-4", "gpt-4o-mini", "gpt-3.5-turbo", "llama-3-70b", "mistral-large"}
	return models[rand.Intn(len(models))]
}

func randomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func main() {
	rand.Seed(time.Now().UnixNano())

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/?replicaSet=rs0")
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("macaroon").Collection("agentRuns")

	var runs []interface{}

	for i := 1; i <= 100; i++ {
		inputTokens := randomInt(100, 5000)
		outputTokens := randomInt(50, 3000)
		totalTime := rand.Float64() * 10
		totalSpend := float64(inputTokens+outputTokens) * 0.000002

		run := Run{
			RunID:         i,
			Timestamp:     time.Now().AddDate(0, 0, -randomInt(0, 30)),
			UserID:        randomUserID(),
			ModelUsed:     randomModel(),
			InputTokens:   inputTokens,
			OutputTokens:  outputTokens,
			TotalTime:     totalTime,
			TotalSpend:    totalSpend,
			FilesAttached: rand.Float64() < 0.3,
			HttpCode:      []int{200, 400, 401, 403, 404, 429, 500}[rand.Intn(7)],
		}
		runs = append(runs, run)
	}

	_, err = collection.InsertMany(context.TODO(), runs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 100 agent run documents successfully")
}
