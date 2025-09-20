package tools

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/require"
)

func TestFind(t *testing.T) {
	db := setupDB(t)
	ctx := context.Background()
	defer db.Disconnect(ctx)

	tt := []struct {
		description    string
		filter         string
		collectionName string
		recordCount    int
		expectedError  bool
	}{
		{
			description:    "When a valid collectionName and filter is input, records are returned",
			filter:         "{\"userID\": \"user_36\"}",
			collectionName: "agentRuns",
			recordCount:    2,
			expectedError:  false,
		},
		{
			description:    "When a valid collectionName and invalid filter is input records are returned",
			filter:         "{\"userID\": \"user_007\"}",
			collectionName: "agentRuns",
			recordCount:    0,
			expectedError:  false,
		},
		{
			description:    "When an invalid collectionName and filter is input, 0 records are returned",
			filter:         "{\"userID\": \"user_007\"}",
			collectionName: "agentRunsFake",
			recordCount:    0,
			expectedError:  false,
		},
		{
			description:    "When filter syntax is incorrect an error is returned",
			filter:         "{\"user }",
			collectionName: "agentRuns",
			recordCount:    0,
			expectedError:  true,
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			findFn := Find(db)
			res, _, err := findFn(ctx, nil, &FindInput{
				CollectionName: test.collectionName,
				Filter:         test.filter,
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
