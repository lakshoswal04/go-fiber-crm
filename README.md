 Go Fiber CRM 

A modern Customer Relationship Management (CRM) system built with Go Fiber framework and SQLite database.

 Features

- REST API: Full CRUD operations for lead management
- Go Fiber: Fast and lightweight web framework
- SQLite Database: Pure Go implementation (CGO-free)
- GORM: Elegant ORM for database operations


## 🛠️ Tech Stack

- Backend: Go (Golang)
- Database: SQLite with GORM
- SQLite Driver: modernc.org/sqlite (pure Go, no CGO dependencies)

  Project Structure

```
go-fiber-crm/
├── .gitignore
├── README.md
├── go.mod
├── go.sum
├── main.go              # Application entry point
├── database/
│   └── database.go      # Database connection setup
└── lead/
    └── lead.go          # Lead model and handlers
```

 Getting Started

 Prerequisites

- Go 1.19 or higher
- Git

 Installation

1. Clone the repository
   ```bash
   git clone https://github.com/lakshoswal04/go-fiber-crm.git
   cd go-fiber-crm
   ```

2. Install dependencies
   ```bash
   go mod tidy
   ```

3. Run the application
   ```bash
   go run main.go
   ```

4. Server will start on
   ```
   http://localhost:3000
   ```

 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/lead` | Get all leads |
| GET | `/api/v1/lead/:id` | Get a specific lead |
| POST | `/api/v1/lead` | Create a new lead |
| DELETE | `/api/v1/lead/:id` | Delete a lead |

Lead Model

```json
{
  "id": 1,
  "name": "John Doe",
  "company": "Tech Corp",
  "email": "john@techcorp.com",
  "phone": "+1234567890",
  "created_at": "2025-01-01T00:00:00Z",
  "updated_at": "2025-01-01T00:00:00Z"
}
```

 API Usage Examples

Create a Lead
```bash
curl -X POST http://localhost:3000/api/v1/lead \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "company": "Tech Corp",
    "email": "john@techcorp.com",
    "phone": "+1234567890"
  }'
```

Get All Leads
```bash
curl http://localhost:3000/api/v1/lead
```

 Get a Specific Lead
```bash
curl http://localhost:3000/api/v1/lead/1
```

Delete a Lead
```bash
curl -X DELETE http://localhost:3000/api/v1/lead/1
```

Development

 Building the Application
```bash
go build -o crm main.go
```

 Running Tests
```bash
go test ./...
```
 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

Project Link: [https://github.com/lakshoswal04/go-fiber-crm](https://github.com/lakshoswal04/go-fiber-crm)

---

⭐ **Star this repository if you found it helpful!**
