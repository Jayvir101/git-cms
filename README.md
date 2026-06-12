# Git CMS

A Git-backed Headless CMS built in Go.

The goal of this project is to explore backend architecture, content management systems, Git-based storage, authentication, and REST API design while following production-grade engineering practices.

Instead of storing content in a traditional database, content is stored as Markdown files with frontmatter metadata and versioned using Git.

---

## Project Goals

This project aims to implement:

* REST API
* JWT Authentication
* Markdown-based Content Storage
* Frontmatter Metadata Parsing
* Slug Generation
* Filesystem Storage Abstraction
* Git Integration
* Role-ready Architecture
* Production-grade Project Structure

---

## High-Level Architecture

```text
Client
   │
   ▼
HTTP Router (Chi)
   │
   ▼
Middleware
   ├── Logging
   ├── Recovery
   └── Authentication
   │
   ▼
Handlers
   │
   ▼
Services
   │
   ├── Content Service
   ├── Auth Service
   └── Git Service
   │
   ▼
Storage Interfaces
   │
   ▼
Filesystem
   │
   ▼
Markdown Files
   │
   ▼
Git Repository
```

---

## Request Flow

```text
POST /api/posts

      │
      ▼

Content Handler

      │
      ▼

Content Service

      │
      ▼

Generate Slug

      │
      ▼

Generate Frontmatter

      │
      ▼

Filesystem Store

      │
      ▼

Write Markdown File

      │
      ▼

Git Service

      │
      ▼

Commit Changes
```

---

## Technology Stack

| Purpose             | Library                 |
| ------------------- | ----------------------- |
| Language            | Go                      |
| Router              | Chi                     |
| Logging             | Zap                     |
| Config              | godotenv                |
| Validation          | go-playground/validator |
| Authentication      | golang-jwt/jwt          |
| Password Hashing    | bcrypt                  |
| Markdown Rendering  | Goldmark                |
| Frontmatter Parsing | yaml.v3                 |
| Git Integration     | go-git                  |
| Testing             | testing, httptest       |

---

## Project Structure

```text
git-cms/
├── cmd/
│   └── server/
│       └── main.go

├── internal/

│   ├── api/
│   │   ├── handler/
│   │   ├── middleware/
│   │   └── router.go

│   ├── auth/
│   │   ├── service.go
│   │   └── store.go

│   ├── content/
│   │   ├── service.go
│   │   ├── store.go
│   │   ├── markdown.go
│   │   └── slug.go

│   ├── git/
│   │   └── service.go

│   └── config/
│       └── config.go

├── pkg/
│   └── apierr/
│       └── errors.go

├── content/

├── .env.example
├── go.mod
└── go.sum
```

---

## Package Responsibilities

### cmd/server

Application entry point.

Responsible for:

* Loading configuration
* Creating dependencies
* Wiring services
* Starting the HTTP server

Contains no business logic.

---

### internal/api

HTTP layer.

Responsible for:

* Route registration
* Request parsing
* Response writing
* Middleware execution

Contains no business logic.

---

### internal/auth

Authentication domain.

Responsible for:

* User registration
* Password hashing
* JWT generation
* JWT validation

---

### internal/content

Content management domain.

Responsible for:

* CRUD operations
* Markdown generation
* Frontmatter handling
* Slug generation

---

### internal/git

Git integration layer.

Responsible for:

* Repository initialization
* Staging changes
* Creating commits
* Pushing to remote repositories

---

### internal/config

Configuration management.

Responsible for:

* Loading environment variables
* Providing typed configuration objects

---

### pkg/apierr

Shared API error definitions.

Responsible for:

* Error codes
* HTTP status mapping
* Consistent API error responses

---

## Content Format

Posts are stored as Markdown files.

Example:

```markdown
---
title: Embedded Systems in Cars
slug: embedded-systems-in-cars
tags:
  - embedded
  - automotive
published: true
---

Modern vehicles contain dozens of ECUs that...
```

Stored as:

```text
content/embedded-systems-in-cars.md
```

---

## Planned API Endpoints

### Health

```http
GET /health
```

---

### Authentication

```http
POST /api/auth/register
POST /api/auth/login
```

---

### Content

```http
GET    /api/posts
GET    /api/posts/{slug}

POST   /api/posts

PUT    /api/posts/{slug}

DELETE /api/posts/{slug}
```

---

## Development Roadmap

### Phase 1

* [ ] Project Foundation
* [ ] Configuration
* [ ] Logging
* [ ] Health Endpoint

### Phase 2

* [ ] Storage Abstraction
* [ ] Filesystem Store

### Phase 3

* [ ] Markdown Storage
* [ ] Frontmatter Parsing
* [ ] Slug Generation

### Phase 4

* [ ] Content CRUD API

### Phase 5

* [ ] Authentication
* [ ] JWT
* [ ] Password Hashing

### Phase 6

* [ ] Auth Middleware

### Phase 7

* [ ] Git Integration

### Phase 8

* [ ] Testing

### Phase 9

* [ ] Dockerization
* [ ] Production Hardening

---

## Engineering Principles

This project follows several principles:

### Separation of Concerns

Handlers handle HTTP.

Services handle business logic.

Stores handle persistence.

---

### Dependency Inversion

Services depend on interfaces rather than concrete implementations.

Example:

```text
ContentService
        │
        ▼
ContentStore Interface
        │
        ▼
FilesystemStore
```

This allows storage implementations to be swapped without changing business logic.

---

### Content as Code

Content is treated as source code.

Benefits:

* Full version history
* Easy rollback
* Human-readable storage
* Git-based collaboration
* Audit trail

---

## Future Enhancements

Potential future improvements:

* PostgreSQL User Store
* Role-Based Access Control (RBAC)
* Refresh Tokens
* Search Integration
* Meilisearch
* OpenTelemetry
* Prometheus Metrics
* Docker Compose
* GitHub Actions CI/CD
* Webhook-based Deployments
* Multi-user Content Workflows

---

## Learning Objectives

By building this project, the goal is to gain practical experience with:

* Go Backend Development
* REST API Design
* Authentication & Authorization
* Filesystem-Based Persistence
* Git Internals
* Dependency Injection
* Clean Architecture
* Testing Strategies
* Production Service Design
* Observability Fundamentals

---

## Status

🚧 Work In Progress
