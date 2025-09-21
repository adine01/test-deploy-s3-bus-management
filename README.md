# Bus Management Service

Bus and staff management microservice for the Choreo platform testing project.

## Features

- Bus registration and management
- Staff registration and management
- CRUD operations for both buses and staff
- In-memory storage for quick testing

## API Endpoints

### Health Check

- `GET /health` - Service health check

### Bus Management

- `POST /api/buses` - Register new bus
- `GET /api/buses` - List all buses
- `GET /api/buses/:id` - Get specific bus
- `PUT /api/buses/:id` - Update bus information
- `DELETE /api/buses/:id` - Delete bus

### Staff Management

- `POST /api/staff` - Register new staff member
- `GET /api/staff` - List all staff
- `GET /api/staff/:id` - Get specific staff member
- `PUT /api/staff/:id` - Update staff information
- `DELETE /api/staff/:id` - Delete staff member

## Request/Response Examples

### Register Bus

```bash
POST /api/buses
Content-Type: application/json

{
  "plate_number": "ABC-1234",
  "model": "Toyota Coaster",
  "capacity": 30
}
```

Response:

```json
{
  "id": 1,
  "plate_number": "ABC-1234",
  "model": "Toyota Coaster",
  "capacity": 30,
  "status": "active",
  "created_at": "2025-09-21T13:30:00Z",
  "updated_at": "2025-09-21T13:30:00Z"
}
```

### Register Staff

```bash
POST /api/staff
Content-Type: application/json

{
  "name": "John Driver",
  "email": "john.driver@example.com",
  "phone": "+1234567890",
  "position": "driver",
  "license_no": "DL123456"
}
```

Response:

```json
{
  "id": 1,
  "name": "John Driver",
  "email": "john.driver@example.com",
  "phone": "+1234567890",
  "position": "driver",
  "license_no": "DL123456",
  "status": "active",
  "created_at": "2025-09-21T13:30:00Z",
  "updated_at": "2025-09-21T13:30:00Z"
}
```

## Running the Service

```bash
# Install dependencies
go mod tidy

# Run the service
go run .

# Or build and run
go build -o bus-management
./bus-management
```

## Environment Variables

- `PORT` - Server port (default: 8081)
- `GIN_MODE` - Gin framework mode (debug/release)
- `DB_HOST` - Database host
- `DB_PORT` - Database port
- `DB_USER` - Database user
- `DB_PASSWORD` - Database password
- `DB_NAME` - Database name

## Docker

```bash
# Build image
docker build -t bus-management .

# Run container
docker run -p 8081:8081 bus-management
```

## Data Models

### Bus

- `id` - Unique identifier
- `plate_number` - Bus registration plate number
- `model` - Bus model/type
- `capacity` - Passenger capacity
- `status` - Bus status (active, maintenance, retired)
- `created_at` - Creation timestamp
- `updated_at` - Last update timestamp

### Staff

- `id` - Unique identifier
- `name` - Staff member name
- `email` - Email address
- `phone` - Phone number
- `position` - Job position (driver, conductor, mechanic)
- `license_no` - Driver's license number (optional)
- `status` - Staff status (active, inactive)
- `created_at` - Creation timestamp
- `updated_at` - Last update timestamp
