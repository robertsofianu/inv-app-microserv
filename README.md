# inv-app-microserv

A microservices-based inventory management application.

## Project Structure

```
invetory-app/
  api_server/
    createNirSvc/      # Service for creating NIR (goods receipt notes)
    CRUD/              # CRUD and authentication services
    dbutils/           # Database utilities and migrations
    home_svc/          # Home/dashboard service
    productsSvc/       # Product management service
    Makefile           # Build and run scripts
    version.txt        # Application version
  ui/
    login.html         # Login page
    createNIR/         # UI for creating NIRs
    home/              # Home page UI
    products/          # Products UI
    images/            # Static images
README.md
```

## Features

- User authentication (login)
- NIR (goods receipt note) creation and management
- Product management
- SQLite database for persistence
- Web-based UI

## Getting Started

### Prerequisites

- Go 1.18+
- Node.js (for UI development, if needed)
- SQLite3

### Build and Run

1. **Backend**

   Navigate to the API server directory and build:

   ```sh
   cd invetory-app/api_server
   make build
   make run
   ```

   The backend will be available at `http://localhost:8080`.

2. **Frontend**

   Open the HTML files in `invetory-app/ui/` in your browser, or serve them with a static file server.

### API Endpoints

- `POST /login` — User login
- `GET /appInfo` — Application info (name, version, build time)
- `POST /nir/create` — Create a new NIR
- Additional endpoints for products and home details

## Configuration

- Update database paths in [`dbutils/db.go`](invetory-app/api_server/dbutils/db.go) and [`CRUD/login_server.go`](invetory-app/api_server/CRUD/login_server.go) as needed.
- Application version is managed in [`version.txt`](invetory-app/api_server/version.txt) and injected at build time.

## License

MIT License

---

For more details, see the code in each service directory.