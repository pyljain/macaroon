package tools

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func setupDB(t *testing.T) *mongo.Client {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017/?replicaSet=rs0"))
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func TestGetCollections(t *testing.T) {
	db := setupDB(t)
	defer db.Disconnect(context.Background())

	res, _, err := GetCollections(db)(context.Background(), nil, nil)
	require.NoError(t, err)
	text := res.Content[0].(*mcp.TextContent).Text
	if text != "agentRuns" {
		t.Errorf("Expected 'agentRuns', got '%s'", text)
	}
}
