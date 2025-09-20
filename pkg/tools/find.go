package tools

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type FindInput struct {
	CollectionName string `json:"collectionName" jsonschema:"Name of the collection to get a schema for"`
	Filter         string `json:"filter" jsonschema:"MongoDB filter in JSON format."`
}

func Find(db *mongo.Client) func(ctx context.Context, req *mcp.CallToolRequest, params *FindInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, req *mcp.CallToolRequest, params *FindInput) (*mcp.CallToolResult, any, error) {

		var result []bson.M
		var mongoFilter bson.M
		err := json.Unmarshal([]byte(params.Filter), &mongoFilter)
		if err != nil {
			return nil, nil, err
		}

		cursor, err := db.Database("macaroon").Collection(params.CollectionName).Find(ctx, mongoFilter)
		if err != nil {
			return nil, nil, err
		}

		data := bson.M{}
		for cursor.Next(ctx) {
			err := cursor.Decode(&data)
			if err != nil {
				return nil, nil, err
			}
			result = append(result, data)
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
