package tools

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/require"
)

func TestAggregateQuery(t *testing.T) {
	db := setupDB(t)
	ctx := context.Background()
	defer db.Disconnect(ctx)

	tt := []struct {
		description    string
		pipeline       string
		collectionName string
		recordCount    int
		expectedError  bool
	}{
		{
			description:    "When a valid collectionName and pipeline is valid, records are returned",
			pipeline:       `[{ "$count": "totalRuns" }]`,
			collectionName: "agentRuns",
			recordCount:    1,
			expectedError:  false,
		},
		{
			description:    "When an invalid collectionName and pipeline is valid, 0 records are returned",
			pipeline:       `[{ "$count": "totalRuns" }]`,
			collectionName: "agentRunsFake",
			recordCount:    0,
			expectedError:  false,
		},
		{
			description:    "When pipeline syntax is incorrect an error is returned",
			pipeline:       "{\"user }",
			collectionName: "agentRuns",
			recordCount:    0,
			expectedError:  true,
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			aggregateFn := AggregateQuery(db)
			res, _, err := aggregateFn(ctx, nil, &AggregateQueryInput{
				CollectionName: test.collectionName,
				Pipeline:       test.pipeline,
			})

			if test.expectedError {
				require.Error(t, err)
				return
			}

			text := res.Content[0].(*mcp.TextContent).Text
			data := []interface{}{}
			err = json.Unmarshal([]byte(text), &data)
			require.NoError(t, err)

			require.Equal(t, test.recordCount, len(data))
		})
	}
}
