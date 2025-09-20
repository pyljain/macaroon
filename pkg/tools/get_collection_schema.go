package tools

import (
	"context"
	_ "embed"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

//go:embed schemas/runs.json
var runsData string

type GetCollectionSchemaInput struct {
	CollectionName string `json:"collectionName" jsonschema:"Name of the collection to get a schema for"`
}

func GetCollectionSchema(db *mongo.Client) func(ctx context.Context, req *mcp.CallToolRequest, params *GetCollectionSchemaInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, req *mcp.CallToolRequest, params *GetCollectionSchemaInput) (*mcp.CallToolResult, any, error) {
		schema := "Not Found"
		switch params.CollectionName {
		case "runs":
			schema = runsData
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: schema},
			},
		}, nil, nil
	}
}
