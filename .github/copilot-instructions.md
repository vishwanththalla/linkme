# Copilot Instructions for linkme

This repository implements a **simple link management platform** written in Go. Even though most of the source files are currently stubs, the directory layout and naming conventions reveal the intended architecture.  AI agents should treat this as a layered, idiomatic Go web service and follow the patterns already established.

## 📁 High‑level architecture

1. **`cmd/server/main.go`** – program entry point. In a finished service it wires up the HTTP server, middleware, and routes; it will call packages under `internal/`.
2. **`internal/` packages** – used for encapsulation (Go's `internal` directory prevents external import). Key subpackages:
   - `config` – environment loading (from `.env`) and application settings.
   - `database` – establishes a DB connection (Postgres, using env vars from `.env`).
   - `models` – data structures for `User`, `Link`, etc.
   - `services` – business logic that operates on models and talks to the database.
   - `handlers` – HTTP handlers that receive requests, call services, and render templates or JSON.
   - `middleware` – HTTP middleware such as authentication; often uses `utils/jwt.go`.
   - `utils` – helpers like password hashing (`hash.go`) and JWT creation/verification (`jwt.go`).
3. **`templates/`** – HTML templates (`dashboard.html`, `login.html`, `register.html`) rendered by handlers.
4. **`.env`** – configuration parameters (port, DB credentials, `JWT_SECRET`). Look at the file for all keys.

> **Dataflow example**: an authenticated request to `/links` is intercepted by `auth_middleware.go`, which validates a token using `utils/jwt.go`. The corresponding handler in `link_handler.go` calls `link_service.go` to fetch links via `database/db.go` and the `models.Link` struct, then renders `dashboard.html`.

## ⚙️ Developer workflows

* **Run locally**: set up environment (copy `.env` and adjust secrets) and execute
  ```sh
  go run ./cmd/server
  ```
  or build with `go build ./cmd/server`.
* **Tests**: there are no tests yet, but add them under the same packages. Run `go test ./...` to execute all tests once they exist.
* **Dependency management**: use Go modules. If `go.mod` is missing, initialize it with `go mod init github.com/yourusername/linkme` and add imports as you develop.
* **Database**: assumes PostgreSQL; connection parameters come from `.env`. Migrate schemas manually or with a tool of your choice.
* **Templates**: use `html/template` or similar; handlers should load them at startup.

> **Note**: there is no build script or Makefile yet. Use standard `go` commands.

## 🔁 Conventions & patterns

* Package names match directory names and are imported by other packages (e.g. `internal/services`).
* Each logical component (handler, service, model) has its own file with a comment at the top. When implementing new features, follow that pattern and keep files small.
* `internal` ensures code is only usable within the repository.
* Handlers should be thin: validate input, call services, return responses. Business logic belongs in `services`.
* Use the `utils` package for cross‑cutting helpers (currently hashing and JWT only).
* Repository‑specific constant names (e.g. `Link`, `User`) are defined in the model files; extend them there.

## 📌 Integration points

* **Environment** – loaded from `.env` using whatever library you choose (e.g., `github.com/joho/godotenv`).
* **Database** – Postgres via `internal/database`. Expect a `DB` object exported for use by services.
* **HTTP** – uses Go’s standard `net/http`; middleware typically wraps `http.Handler`.
* **Templates** – located in `templates/` and should be loaded once and passed to handlers or stored globally.
* **JWT** – secrets from `.env`; `utils/jwt.go` contains helpers for signing/verifying.

## ✅ Practical examples

* When adding a new API route, create a handler in `internal/handlers` and export a function such as:
  ```go
  func GetLinks(w http.ResponseWriter, r *http.Request) { … }
  ```
  Wire it in `main.go` with `http.HandleFunc("/links", handlers.GetLinks)`.
* Add business rules to services (`internal/services/link_service.go`) and keep database queries in `internal/database`.
* For authentication, mimic existing stubs: `auth_middleware.go` should read the `Authorization` header and call `utils.ValidateToken`.

---

✏️ **Feedback request**: please read the above and let me know if any section is unclear or missing details that would help an AI agent “hit the ground running.” Are there other conventions or workflows used by the team that I haven’t captured? Once I have your input I can iterate on this file.