package tools

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/require"
)

func TestFindOne(t *testing.T) {
	db := setupDB(t)
	ctx := context.Background()
	defer db.Disconnect(ctx)

	tt := []struct {
		description    string
		filter         string
		collectionName string
		expectedError  bool
	}{
		{
			description:    "When a valid collectionName and filter are provided, a record is returned",
			filter:         "{\"userID\": \"user_36\"}",
			collectionName: "agentRuns",
			expectedError:  false,
		},
		{
			description:    "When a valid collectionName and invalid filter, no records are returned",
			filter:         "{\"userID\": \"user_007\"}",
			collectionName: "agentRuns",
			expectedError:  true,
		},
		{
			description:    "When an invalid collectionName and valid filter is input, 0 records are returned",
			filter:         "{\"userID\": \"user_007\"}",
			collectionName: "agentRunsFake",
			expectedError:  true,
		},
		{
			description:    "When filter syntax is incorrect an error is returned",
			filter:         "{\"user }",
			collectionName: "agentRuns",
			expectedError:  true,
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			findOneFn := FindOne(db)
			res, _, err := findOneFn(ctx, nil, &FindOneInput{
				CollectionName: test.collectionName,
				Filter:         test.filter,
			})

			if test.expectedError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			text := res.Content[0].(*mcp.TextContent).Text
			var data interface{}
			err = json.Unmarshal([]byte(text), &data)
			require.NoError(t, err)
		})
	}
}
