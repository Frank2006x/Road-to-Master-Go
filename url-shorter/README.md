# URL Shortener

A simple URL shortener service built with Go, Fiber v3, and SQLite. It allows you to shorten long URLs into compact identifiers and redirect to the original URLs.

## Features

- Shorten URLs with unique IDs generated using nanoid
- Redirect short URLs to original links
- Track click counts for each short URL
- List all shortened URLs
- Local SQLite database for persistence

## Prerequisites

- Go 1.26.1 or later
- Internet connection for downloading dependencies

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Frank2006x/Road-to-Master-Go.git
   cd Road-to-Master-Go
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Run the server:
   ```bash
   go run server.go
   ```

The server will start on `http://localhost:3000`.

## Usage

### Shorten a URL

Send a POST request to `/shorten/` with a JSON body containing the URL:

```bash
curl -X POST http://localhost:3000/shorten/ \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com"}'
```

Response:

```json
{
  "id": "abc123def4"
}
```

### Redirect to Original URL

Visit `http://localhost:3000/{id}` (e.g., `http://localhost:3000/abc123def4`) to redirect to the original URL. The click count will be incremented.

### List All URLs

Send a GET request to `/shorten/`:

```bash
curl http://localhost:3000/shorten/
```

Response:

```json
{
  "urls": [
    {
      "id": "abc123def4",
      "url": "https://example.com",
      "short_url": "abc123def4",
      "count": 5
    }
  ]
}
```

## Database

The application uses SQLite (`test.db`) with the following schema:

```sql
CREATE TABLE urls (
  id TEXT PRIMARY KEY,
  url TEXT NOT NULL,
  short_url TEXT NOT NULL,
  count INTEGER DEFAULT 0
);
```

The table is created automatically on first run.

## API Endpoints

- `POST /shorten/` - Shorten a URL
- `GET /shorten/` - List all shortened URLs
- `GET /:id` - Redirect to the original URL

## Development

- The main entry point is `server.go`
- Controllers are in `src/controller/`
- Routes are in `src/router/`
- Database logic is in `src/db/`

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is open source. Feel free to use and modify.
