package tools

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type FindOneInput struct {
	CollectionName string `json:"collectionName" jsonschema:"Name of the collection to get a schema for"`
	Filter         string `json:"filter" jsonschema:"MongoDB filter in JSON format."`
}

func FindOne(db *mongo.Client) func(ctx context.Context, req *mcp.CallToolRequest, params *FindOneInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, req *mcp.CallToolRequest, params *FindOneInput) (*mcp.CallToolResult, any, error) {

		var result bson.M
		var mongoFilter bson.M
		err := json.Unmarshal([]byte(params.Filter), &mongoFilter)
		if err != nil {
			return nil, nil, err
		}

		record := db.Database("macaroon").Collection(params.CollectionName).FindOne(ctx, mongoFilter)
		if record.Err() != nil {
			return nil, nil, record.Err()
		}

		err = record.Decode(&result)
		if err != nil {
			return nil, nil, err
		}

		res, err := json.Marshal(result)
		if err != nil {
			return nil, nil, err
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: string(res)},
			},
		}, nil, nil
	}
}
