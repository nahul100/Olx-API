# OLX API

A REST API for classified listings, built from scratch using **Go** and **PostgreSQL**. This project focuses on backend engineering fundamentals, API design, database integration, and practical software development practices.

## Project Diagrams

I documented the concepts and architecture I studied while building this project through visual diagrams.

[View Project Diagrams](https://viewer.diagrams.net/?tags=%7B%7D&lightbox=1&target=blank&highlight=0000ff&edit=_blank&layers=1&nav=1&dark=auto#G1zHPqDnMWet2svbXkZjBlS-JAnyDxkszX)

## Production Engineering Concepts & Practices

- **Configuration Management:** Manage application configuration through environment variables, including the server port, environment name, and database URL. Use `.env` for local development while keeping sensitive configuration out of Git.
- **Package Organization & Separation of Concerns:** Organize the application into separate Go packages for configuration, database connectivity, HTTP handlers, middleware, and shared HTTP utilities. Keep responsibilities isolated to improve maintainability, readability, scalability and testability.
- **Database Integration:** Connect Go applications to PostgreSQL using `database/sql` and the pgx driver.
- **Cloud Database:** Use **Neon** as the hosted PostgreSQL database and configure database connectivity through environment variables.
- **Database Connection Management:** Configure connection-pool limits, maximum connection lifetime, and startup connectivity checks with timeouts.
- **Database Migrations:** Manage schema changes through versioned up/down SQL migration files and execute migrations separately from application startup.
- **Seed Data Management:** Keep sample data in separate SQL seed files rather than mixing it with database schema definitions.
- **Dependency Injection:** Apply constructor-based dependency injection to provide database connections and loggers to handlers.
- **HTTP Routing:** Use Go's standard-library `http.ServeMux` for method-based routing, URL path parameters, and clean route registration without requiring an external routing dependency.
- **Handler Architecture:** Organize HTTP handlers by responsibility and explore closure-based handler factories.
- **DTOs & Request Validation:** Use Data Transfer Objects (DTOs) to structure incoming request data and validate request bodies before processing.
- **Custom Error Handling:** Implement a consistent JSON error envelope containing an error code, message, and optional field information.
- **Context-Aware Database Operations:** Propagate request-scoped contexts to database queries and writes using `QueryContext`, `ExecContext`, and `QueryRowContext` to support cancellation and deadline handling, preventing unnecessary long-running database operations.
- **Request ID Middleware:** Accept or generate unique `X-Request-ID` values, return them in response headers, and make them available throughout request handling.
- **Structured Logging:** Use Go's `log/slog` package to produce structured logs containing timestamps, severity levels, messages, request IDs, listing IDs, and error details.
- **SQL Injection Prevention:** Use parameterized SQL queries with PostgreSQL placeholders (`$1`, `$2`, etc.) through Go's `database/sql` package to separate user-supplied values from SQL statements and mitigate SQL injection risks.
- **HTTP Server Timeouts:** Configure `ReadTimeout`, `WriteTimeout`, and `IdleTimeout` to control connection handling and protect server resources.
- **API Development:** Implement listing creation, retrieval, and deletion endpoints with database-backed operations.
- **Cloud Deployment:** Deploy the Go API using **Render** and connect it to the hosted PostgreSQL database on Neon.


## Deployment

- **Application Hosting:** Render
- **PostgreSQL Database:** Neon
- **Backend Language:** Go

*This is a hands-on learning project focused on understanding backend development and applying practical engineering concepts. The diagrams also document additional concepts explored during development; they do not imply that every illustrated technology or component is implemented in the application.*