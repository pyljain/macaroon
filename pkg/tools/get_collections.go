package tools

import (
	"context"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type GetCollectionsInput struct {
}

func GetCollections(db *mongo.Client) func(ctx context.Context, req *mcp.CallToolRequest, params *GetCollectionsInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, req *mcp.CallToolRequest, params *GetCollectionsInput) (*mcp.CallToolResult, any, error) {

		var result []string
		cursor, err := db.Database("macaroon").ListCollections(ctx, bson.M{})
		if err != nil {
			return nil, nil, err
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var data bson.M
			err := cursor.Decode(&data)
			if err != nil {
				return nil, nil, err
			}
			if name, ok := data["name"].(string); ok {
				result = append(result, name)
			}
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: strings.Join(result, "\n")},
			},
		}, nil, nil
	}
}
