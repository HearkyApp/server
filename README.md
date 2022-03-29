# UpMeet Server

This repository contains the REST API server for UpMeet.

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/24916d5be28b4c378ef207d1a0a48019)](https://www.codacy.com/gh/UpMeetApp/server/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=UpMeetApp/server&amp;utm_campaign=Badge_Grade)

# Development Setup

1. Clone the repository
2. Install dependencies (`go mod download`)
3. Run `docker-compose -f dev.docker-compose.yml up -d` to start up the development stack. (Currently only PostgreSQL)
4. Setup [environment variables](#environment-variables) (.env file is supported)
5. Run `go run cmd/server/main.go` to start the server

# Environment Variables

- `UPMEET_DEBUG`: Whether to enable debug mode.
- `UPMEET_BIND_ADDRESS`: The address to bind the web server to.
- `UPMEET_FIREBASE_CREDENTIALS`: The Firebase service account key JSON file encoded into Base64.
- `UPMEET_POSTGRES_HOST`: The hostname of the PostgreSQL server.
- `UPMEET_POSTGRES_PORT`: The port of the PostgreSQL server.
- `UPMEET_POSTGRES_USER`: The username of the PostgreSQL server.
- `UPMEET_POSTGRES_PASSWORD`: The password of the PostgreSQL server.
- `UPMEET_POSTGRES_DATABASE`: The PostgreSQL database name.
- `UPMEET_POSTGRES_SSL`: The PostgreSQL SSL mode.