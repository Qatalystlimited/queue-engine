# Queue Engine

A gRPC-based queue service built with Go and PostgreSQL (Supabase).

## Architecture

```
Node.js Client
     ↓ gRPC
Queue Engine (Go)
     ↓
PostgreSQL (Supabase)
```

## Deliverables

### Phase 1 — gRPC Skeleton (Current)
- [x] `queue.proto` — service contract defined
- [x] gRPC code generated (`proto/queue.pb.go`, `proto/queue_grpc.pb.go`)
- [x] Go server running on port `50051`
- [x] `JoinQueue` endpoint responding (stub)
- [x] `GetPosition` endpoint defined

### Phase 2 — Real Queue Logic
- [x] Connect to Supabase (PostgreSQL)
- [ ] Implement transaction-safe `JoinQueue`
- [ ] Implement `GetPosition` with live DB query

### Phase 3 — Production Hardening
- [ ] Error handling & input validation
- [ ] Dockerize the service
- [ ] Environment config via `.env`

## gRPC Endpoints

| Method         | Input                    | Output                          |
|----------------|--------------------------|---------------------------------|
| `JoinQueue`    | `user_id`, `queue_id`    | `ticket_id`, `position`, `status` |
| `GetPosition`  | `user_id`, `queue_id`    | `position`                      |

## Quick Start

```bash
# .env
cp .env .env.example

# Generate gRPC code
protoc --go_out=. --go-grpc_out=. queue.proto

# Run server
go run cmd/main.go
```

Server starts on `localhost:50051`.

## Tech Stack

- **Language:** Go
- **Transport:** gRPC / Protocol Buffers
- **Database:** PostgreSQL via Supabase
