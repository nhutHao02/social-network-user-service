# social-network-user-service
## Project Summary
This is project about social network that allows users to share content, images, and emotions, and have real-time communication capabilities, while ensuring high performance, security, and scalability using the microservices architecture.

#### Technologies:
- Back-end:
  - Language: Go.
  - Frameworks/Platforms: Gin-Gonic, gRPC, Swagger, JWT, Google-Wire, SQLX, Redis, Zap, WebSocket.
  - Database: MariaDB, MongoDB.
- Front-end:
  - Language: JavaScript.
  - Frameworks/Platforms: React, Tailwind CSS, FireBase.

## The project includes repositories
- [common-service](https://github.com/nhutHao02/social-network-common-service)
- [user-service](https://github.com/nhutHao02/social-network-user-service)
- [tweet-service](https://github.com/nhutHao02/social-network-tweet-service)
- [chat-service](https://github.com/nhutHao02/social-network-chat-service)
- [notification-service](https://github.com/nhutHao02/social-network-notification-service)
- [Front-end-service (in progress)](https://github.com/nhutHao02/)

## This service
This is the service that provides the APIs related to the User Information and Authentication.

## ER Diagram
![ER Diagram](https://github.com/user-attachments/assets/f67593e1-3b3a-4fc0-a3cd-e2480a2ee616)

## Project structure
```
.
├── config
│   ├── config.go
│   └── local
│       └── config.yaml
├── database
│   └── database.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── install.sh
├── internal
│   ├── api
│   │   ├── grpc
│   │   │   ├── grpc_server.go
│   │   │   └── user_grpc.go
│   │   ├── http
│   │   │   ├── http_server.go
│   │   │   └── v1
│   │   │       ├── route.go
│   │   │       └── user_handler.go
│   │   └── server.go
│   ├── application
│   │   ├── imp
│   │   │   └── user_service_imp.go
│   │   └── user_service.go
│   ├── domain
│   │   ├── entity
│   │   │   ├── location.go
│   │   │   └── user.go
│   │   ├── interface
│   │   │   └── user
│   │   │       └── user_repository.go
│   │   └── model
│   │       ├── auth.go
│   │       ├── follow.go
│   │       ├── location.go
│   │       └── user.go
│   ├── infrastructure
│   │   └── user
│   │       ├── command_repository.go
│   │       └── query_repository.go
│   ├── wire_gen.go
│   └── wire.go
├── main.go
├── Makefile
├── migrations
│   ├── 000001_location.down.sql
│   ├── 000001_location.up.sql
│   ├── 000002_user.down.sql
│   ├── 000002_user.up.sql
│   ├── 000003_follow.down.sql
│   └── 000003_follow.up.sql
├── pkg
│   ├── common
│   │   └── response.go
│   ├── constants
│   │   └── constants.go
│   ├── grpc
│   │   ├── proto
│   │   │   ├── user_message.proto
│   │   │   └── user.proto
│   │   ├── user_grpc.pb.go
│   │   ├── user_message.pb.go
│   │   └── user.pb.go
│   └── redis
│       └── redis.go
├── README.md
└── startup
    └── startup.go
```