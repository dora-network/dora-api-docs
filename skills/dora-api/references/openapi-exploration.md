# Tips for Exploring OpenAPI Specifications

When working with an OpenAPI/Swagger specification to understand an API, these commands and approaches can be helpful.

## Quick Overview
```bash
# Get basic info about the spec
jq '.info' api-spec.json

# List all tags (API categories)
jq '.tags[] | .name' api-spec.json

# Count endpoints
jq '.paths | length' api-spec.json
```

## Finding Endpoints and Operations
```bash
# List all operation IDs (unique identifier for each endpoint)
jq -r '.paths[] | .[] | .operationId' api-spec.json | sort

# List all paths and their methods
jq -r 'to_entries[] | "\(.key): \(.value | keys | join(", "))"' api-spec.json

# Find endpoints by keyword in summary or description
jq -r '.paths[] | .[] | select(.summary | test("user"; "i")) | "\(.operationId): \(.summary)"' api-spec.json
```

## Examining Components (Schemas)
```bash
# List all schema definitions
jq -r '.components.schemas | keys[]' api-spec.json

# Get a specific schema (pretty-printed)
jq '.components.schemas.UserCreatedResponse' api-spec.json

# List all properties of a schema
jq '.components.schemas.UserCreatedResponse.properties | keys[]' api-spec.json

# Find schemas that reference a specific component
jq -r '.components.schemas[] | select(.properties | has("userId")) | keys[0]' api-spec.json
```

## Security and Authentication
```bash
# Check security schemes
jq '.components.securitySchemes' api-spec.json

# See which endpoints require authentication
jq -r '.paths[] | .[] | select(.security != []) | "\(.operationId): \(.summary // "no summary")"' api-spec.json
```

## Response Codes
```bash
# See what status codes an endpoint returns
jq -r '.paths["/v1/user"].get.responses | keys[]' api-spec.json

# Find all endpoints that can return 404
jq -r '.paths[] | .[] | select(.responses["404"]) | .operationId' api-spec.json
```

## Examples
```bash
# Check if an operation has examples
jq -r '.paths[] | .[] | select(.requestBody | has("examples")) | .operationId' api-spec.json

# Look for examples in responses
jq -r '.paths[] | .[] | select(.responses["200"].content | has("application/json")) | .operationId' api-spec.json
```

## Common Patterns to Look For
- Pagination: Look for `page`, `limit`, `offset` parameters
- Filtering: Look for query parameters with descriptive names
- Sorting: Look for `sort`, `order` parameters
- Dates: Look for `format: "date-time"` in string parameters
- UUIDs: Look for `format: "uuid"` in path or query parameters