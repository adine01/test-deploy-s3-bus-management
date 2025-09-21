# Bus Management Service

A microservice for managing buses and staff in a transportation system built with Go and Gin framework.

## Features

- **Bus Management**: Create, read, update, and delete buses
- **Staff Management**: Manage staff members including drivers, conductors, mechanics, and supervisors
- **Database Integration**: PostgreSQL with pgx/v5 driver
- **API Documentation**: OpenAPI 3.0 specification
- **Health Monitoring**: Health check endpoint
- **CORS Support**: Cross-origin resource sharing enabled

## API Endpoints

### Health

- `GET /health` - Service health check

### Buses

- `GET /api/buses` - Get all buses
- `POST /api/buses` - Create a new bus
- `GET /api/buses/{id}` - Get bus by ID
- `PUT /api/buses/{id}` - Update bus
- `DELETE /api/buses/{id}` - Delete bus

### Staff

- `GET /api/staff` - Get all staff
- `POST /api/staff` - Create a new staff member
- `GET /api/staff/{id}` - Get staff member by ID
- `PUT /api/staff/{id}` - Update staff member
- `DELETE /api/staff/{id}` - Delete staff member

## Data Models

### Bus

- ID, Plate Number, Model, Capacity, Status (active/maintenance/retired)
- Timestamps: created_at, updated_at

### Staff

- ID, Name, Email, Phone, Position (driver/conductor/mechanic/supervisor)
- License Number (optional), Status (active/inactive)
- Timestamps: created_at, updated_at

## Environment Variables

- `DATABASE_URL` - PostgreSQL connection string
- `PORT` - Service port (default: 8081)
- `GIN_MODE` - Gin mode (release for production)

## API Documentation

The complete API documentation is available in `openapi.yaml` following OpenAPI 3.0 specification.

## Choreo Deployment

This service is configured for deployment on WSO2 Choreo platform:

- Endpoint configuration: `.choreo/endpoints.yaml`
- OpenAPI specification: `openapi.yaml`
- Docker support: `Dockerfile`

## Local Development

1. Set up environment variables in `.env` file
2. Run with: `go run .`
3. Service will be available at `http://localhost:8081`

## Dependencies

- Go 1.21+
- Gin web framework
- PostgreSQL with pgx/v5 driver
- Supabase (optional, for managed PostgreSQL)
