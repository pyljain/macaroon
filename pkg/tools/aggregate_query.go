package tools

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AggregateQueryInput struct {
	CollectionName string `json:"collectionName" jsonschema:"Name of the collection to run the aggregate query on"`
	Pipeline       string `json:"pipeline" jsonschema:"MongoDB aggregate pipeline to execute"`
}

func AggregateQuery(db *mongo.Client) func(ctx context.Context, req *mcp.CallToolRequest, params *AggregateQueryInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, req *mcp.CallToolRequest, params *AggregateQueryInput) (*mcp.CallToolResult, any, error) {

		var result []bson.M
		var mongoPipeline bson.A
		err := json.Unmarshal([]byte(params.Pipeline), &mongoPipeline)
		if err != nil {
			return nil, nil, err
		}

		cursor, err := db.Database("macaroon").Collection(params.CollectionName).Aggregate(ctx, mongoPipeline)
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
