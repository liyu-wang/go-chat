# Go Chat

A real-time chat application built with Go and WebSockets.

## Features

- Real-time messaging using WebSockets
- User-friendly web interface
- Google OAuth authentication (configurable)
- Clean package structure with internal modules

## Project Structure

```
go-chat/
├── auth.go                  # Authentication logic
├── client.go                # WebSocket client implementation
├── room.go                  # Chat room management
├── main.go                  # Application entry point
├── go-chat                  # Compiled executable
├── templates/               # HTML templates
│   ├── chat.html
│   └── login.html
├── internal/
│   └── trace/               # Tracing package
│       ├── trace.go
│       └── tracer_test.go
├── .env                     # Environment variables (local only, not in git)
├── .env.example             # Environment variables template
├── setup.sh                 # Setup script to load environment
├── go.mod                   # Go module definition
├── go.sum                   # Go dependencies lock file
└── README.md                # This file
```

## Prerequisites

- Go 1.22 or higher
- Git

## Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/liyu-wang/go-chat.git
cd go-chat
```

### 2. Create .env file

Copy the example environment file and add your credentials:

```bash
cp .env.example .env
```

Then edit `.env` and add your Google OAuth credentials:

```bash
# .env
GOOGLE_CLIENT_ID=your_client_id_here
GOOGLE_CLIENT_SECRET=your_client_secret_here
```

**How to get Google OAuth credentials:**
1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Create a new project
3. Enable the Google+ API
4. Create OAuth 2.0 credentials (Web application)
5. Add authorized redirect URIs
6. Copy the Client ID and Client Secret to `.env`

### 3. Load environment variables

```bash
source setup.sh
```

You should see the output:
```
✓ Environment variables loaded successfully
```

### 4. Build the application

```bash
go build -o go-chat
```

This creates an executable named `go-chat` in the current directory.

### 5. Run the application

```bash
./go-chat -addr :8080
```

Then open your browser and navigate to:

```
http://localhost:8080
```

## Command-line Flags

- `-addr` - The server address (default: `:8080`)

Example:
```bash
./go-chat -addr :3000
```

## Development

### Running from the project root

```bash
cd /path/to/go-chat
source setup.sh
go run . -addr :8080
```

### Building for different platforms

```bash
# macOS ARM64 (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o go-chat-darwin-arm64

# macOS Intel
GOOS=darwin GOARCH=amd64 go build -o go-chat-darwin-amd64

# Linux
GOOS=linux GOARCH=amd64 go build -o go-chat-linux

# Windows
GOOS=windows GOARCH=amd64 go build -o go-chat.exe
```

## Environment Variables

The application uses the following environment variables:

| Variable | Description | Required |
|----------|-------------|----------|
| `GOOGLE_CLIENT_ID` | Google OAuth Client ID | No (for basic chat) |
| `GOOGLE_CLIENT_SECRET` | Google OAuth Client Secret | No (for basic chat) |
| `SESSION_SECRET` | Secret key for Gothic session encryption | No (for basic chat) |

These variables are loaded from the `.env` file via `setup.sh`.

**⚠️ Important:** Never commit the `.env` file to version control. The `.gitignore` file already contains `.env` to prevent accidental commits.

## Project Packages

### Root Level Files

- **auth.go** - Authentication and authorization logic
- **client.go** - WebSocket client implementation and message handling
- **room.go** - Chat room management and broadcast functionality
- **main.go** - HTTP server setup and route handlers

### `internal/trace`

Tracing and debugging utilities:
- **trace.go** - Tracer interface and context helpers
- **tracer_test.go** - Unit tests for tracing functionality

## Troubleshooting

### Templates not found

If you get a "templates not found" error:
- Make sure you're running from the project root directory
- Or run using the full path: `/path/to/go-chat/go-chat -addr :8080`

### Environment variables not loaded

Make sure you:
1. Have created the `.env` file
2. Run `source setup.sh` before starting the application
3. Verify variables are set: `echo $GOOGLE_CLIENT_ID`

### Port already in use

If port 8080 is already in use:

```bash
./go-chat -addr :3000
```

Or find and kill the process using port 8080:

```bash
lsof -ti:8080 | xargs kill -9
```

## License

MIT License - See LICENSE file for details

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
