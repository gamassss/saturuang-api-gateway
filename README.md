# Overview

**SatuRuang** is a project aimed at developing a comprehensive real-time communication platform. This **_API Gateway_** repo serves as the central hub for routing requests to various microservices within the **SatuRuang** ecosystem.

Currently, both the API Gateway and other related services are still under development. The goal of SatuRuang is to provide a robust and scalable solution for real-time chat and collaboration, enabling users to join random group chat rooms and interact seamlessly.

This repository contains the API Gateway service, which is responsible for managing and forwarding incoming requests to the appropriate microservices, handling authentication, and ensuring smooth communication across the platform.


## System Design

![saturuang-microservices-architecture-design](https://github.com/user-attachments/assets/858dc2f9-578a-4f26-ad09-2945c4764214)

## Folder Structure

```bash
project-root/
│
├── cmd/
│   └── api-gateway/
│       └── main.go          # Application entry point
│
├── internal/                # Private logic and business code
│   └── gateway/
│       ├── handler/         # Contain all handlers for services
│       ├── middleware/      # Custom middleware
│       └── service/         # Service forwarding logic
│
├── pkg/
│   ├── jwtutils/            # JWT utilities dan klaim
│   └── config/              # Configuration for environment, secrets, etc.
│       └── config.go
├── .env
├── Dockerfile
```

## Technology Stack

- **Programming Language**: Go
- **Framework**: Echo
- **Database**: MongoDB (for other services, not directly in API Gateway)
- **CI/CD**: Docker (for containerization)

## How to Run

1. **Clone the repository**:
   ```bash
   git clone https://github.com/gamassss/saturuang-api-gateway.git
   ```
2. **Navigate and install dependencies**:
   ```bash
   cd saturuang-api-gateway
   go mod tidy
   ``` 
3. **Run the application**:
   ```bash
   go run cmd/api-gateway/main.go
   ``` 
