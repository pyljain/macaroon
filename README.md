# Macaroon

A MongoDB Model Context Protocol (MCP) server that provides natural language access to MongoDB databases through Claude and other AI assistants.

## Overview

Macaroon is an MCP server implementation that bridges MongoDB databases with AI assistants, enabling natural language queries and database operations. Built with Go and the MCP SDK, it provides a secure and efficient way to interact with MongoDB collections through conversational interfaces.

## Features

- **Collection Discovery**: List and describe available MongoDB collections
- **Schema Analysis**: Get natural language descriptions of collection schemas and field structures
- **Query Operations**: Perform find, findOne, and aggregate queries on MongoDB collections
- **Natural Language Interface**: Interact with your MongoDB data using plain English through AI assistants
- **Connection Management**: Efficient connection pooling and management for MongoDB instances

## Architecture

The server exposes several MCP tools that AI assistants can use:

- `get-collections`: Retrieve a list of available collections with descriptions
- `get-schema-for-collection`: Get schema information for specific collections
- `find`: Execute MongoDB find queries with filters and options
- `find-one`: Execute findOne queries to retrieve single documents
- `run-aggregate-query`: Run complex MongoDB aggregation pipelines

## Prerequisites

- Go 1.24.0 or later
- MongoDB instance (local or remote)
- Environment variables configured (see Configuration section)

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd mongomcp
```

2. Install dependencies:
```bash
go mod tidy
```

3. Build the application:
```bash
go build -o mongomcp
```

## Configuration

Set the following environment variables:

- `MONGO_CONNECTION_STRING`: MongoDB connection string (e.g., `mongodb://localhost:27017`)
- `HOST_PORT`: Port for the MCP server to listen on (e.g., `3000`)

Example:
```bash
export MONGO_CONNECTION_STRING="mongodb://localhost:27017"
export HOST_PORT="3000"
```

## Usage

1. Start the MCP server:
```bash
./mongomcp
```

2. The server will listen on the configured port and be ready to accept MCP connections from AI assistants.

3. Configure your AI assistant (like Claude) to connect to the MCP server at `http://localhost:3000` (or your configured port).

## Database Structure

The server connects to a database named "macaroon" by default. Ensure your MongoDB instance has this database with the collections you want to query.

## Examples

Here are example natural language questions you can ask through AI assistants for different types of data:

### For a "runs" Collection (AI API Usage Tracking)

**Schema**: RunID, timestamp, userID, modelUsed, inputTokens, outputTokens, totalTime, totalSpend, filesAttached, httpCode

**Example Questions:**
- "Show me all the runs from the last week"
- "Which AI model is used most frequently?"
- "What's the total spend for each user this month?"
- "Find all runs that took longer than 10 seconds"
- "How many successful runs (HTTP 200) were there today?"
- "What's the average cost per run for GPT-4 vs Claude?"
- "Show me runs where files were attached and the cost was over $1"
- "Which users spent the most on API calls this week?"
- "What's the daily spending trend for the past month?"
- "Find all failed runs (non-200 HTTP codes) and their error patterns"

### For a "metrics" Collection (System Metrics)

**Schema**: model, totalTokens

**Example Questions:**
- "What's the total token usage across all models?"
- "Which model consumes the most tokens on average?"
- "Show me token usage trends over time"
- "Compare token efficiency between different models"

### For an E-commerce "products" Collection

**Schema**: productId, name, category, price, inStock, rating, reviews

**Example Questions:**
- "Show me all products under $50 that are in stock"
- "What are the highest-rated products in the electronics category?"
- "Find products with more than 100 reviews and a rating above 4.5"
- "Which categories have the most expensive products?"
- "Show me out-of-stock items that need restocking"

### For a "users" Collection

**Schema**: userId, email, createdAt, lastLogin, subscription, preferences

**Example Questions:**
- "How many users signed up this month?"
- "Show me users who haven't logged in for over 30 days"
- "What percentage of users have premium subscriptions?"
- "Find users created in the last week with specific preferences"
- "Which users are most active based on login frequency?"

### Complex Aggregation Queries

The system also supports complex aggregation queries through natural language:
- "Group users by subscription type and show average usage"
- "Calculate monthly revenue trends from the orders collection"
- "Find the correlation between product price and rating"
- "Show daily active users for the past quarter"
- "Analyze seasonal trends in product sales"

### How It Works

1. Ask your question in natural language through an AI assistant
2. The AI assistant translates your question into appropriate MongoDB queries
3. Macaroon executes the query against your database
4. Results are returned in a readable format
5. Follow-up questions can dive deeper into the data

The system handles various query types including simple finds, complex filters, aggregation pipelines, and statistical analysis.

## Development

### Running Tests

```bash
go test ./...
```

### Project Structure

```
mongomcp/
├── main.go                 # Main server entry point
├── logging.go             # HTTP request logging middleware  
├── go.mod                 # Go module dependencies
├── pkg/
│   ├── tools/            # MCP tool implementations
│   │   ├── get_collections.go
│   │   ├── get_collection_schema.go
│   │   ├── find.go
│   │   ├── find_one.go
│   │   ├── aggregate_query.go
│   │   └── *_test.go     # Test files
│   └── data/            # Data models and schemas
└── README.md
```

### Key Dependencies

- **MCP SDK**: `github.com/modelcontextprotocol/go-sdk` - Model Context Protocol implementation
- **MongoDB Driver**: `go.mongodb.org/mongo-driver/v2` - Official MongoDB Go driver
- **Testing**: `github.com/stretchr/testify` - Testing utilities

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## License

[Add your license information here]

## Support

[Add support information here]