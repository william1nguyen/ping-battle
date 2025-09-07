# Ping Battle

Ping Battle is a small demo project written in Go that demonstrates how to build a clean and maintainable REST API with Redis as the backend.  

The project includes session management, rate limiting, distributed locking, counters, leaderboards, and unique user counting.

## Features

- User session management with TTL  
- Rate limiting (maximum 2 pings per minute per user)  
- Distributed lock (prevent concurrent pings)  
- Counter per user  
- Leaderboard with Redis Sorted Set  
- Unique user count with Redis HyperLogLog  
- Clean project structure aligned with Go standards  
- Environment-based configuration  
- Dockerized for easy deployment  


## Requirements

- Go 1.21 or later  
- Redis 7.x  
- Docker & Docker Compose (optional)  

## Configuration

All configuration is loaded from environment variables.  
A `.env` file can be used for local development.

## Running Locally

Start Redis with Docker Compose: 
```bash
make start-dev-env
```

Run server with `make`:
```bash
make run
```

## API Endpoints

### Create session
```bash
POST /game/session
```

Example:
```bash
curl -X POST http://localhost:8080/game/session
 -d "username=alice"
```

Response:

```bash
{ 
    "username": "alice", 
    "sessionID": "uuid-here" 
}
```

### Ping
```bash
GET /game/ping?sessionID={id}
```

Example:

```bash
curl -X GET "http://localhost:8080/game/ping?sessionID=uuid-here"
```

Response:
```bash
{ 
    "message": "pong" 
}
```

### Leaderboard
```bash
GET /game/top
```

Example:
```bash
curl -X GET http://localhost:8080/game/top
```

Response:
```bash
[ 
    { 
        "username": "alice", 
        "score": 5 
    }, 
    { 
        "username": "bob", 
        "score": 3 
    } 
] 
```

### Unique user count
```bash
GET /game/count
```

Example:
```bash
curl -X GET http://localhost:8080/game/count
```

Response:
```bash
{ 
    "unique_users": 2 
}
```