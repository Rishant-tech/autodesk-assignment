# autodesk-assignment

Repository scaffold with two standalone Go assignments.

## Structure

```text
autodesk-assignment/
├── assignment1/
│   ├── reverse.go
│   ├── reverse_test.go
│   └── main.go
├── assignment2/
│   ├── updater.go
│   ├── updater_test.go
│   └── main.go
├── README.md
├── go.mod
└── .gitignore
```

## Notes

- Each assignment folder is self-contained.
- `main.go` files provide a simple executable entry point for each assignment.
- The repository uses only the Go standard library.
    locked_until    TIMESTAMPTZ,                   -- NULL when not locked
    last_login_at   TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Sessions table
CREATE TABLE sessions (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    expires_at  TIMESTAMPTZ NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

Migrations run automatically on startup via embedded SQL files — no manual setup required.

---

## Security Design

| Concern              | Implementation                                                     |
|----------------------|--------------------------------------------------------------------|
| Password storage     | bcrypt with default cost (min cost 10)                             |
| Session tokens       | Random UUIDs generated server-side, stored in PostgreSQL           |
| Session expiry       | Checked on every command, configurable timeout                     |
| Account lockout      | Configurable threshold and lockout duration                        |
| 2FA                  | TOTP (RFC 6238), compatible with Google Authenticator              |
| 2FA verification     | Code verified before secret is saved; code required to disable     |

---

## Tech Stack

| Layer       | Technology                                                                 |
|-------------|----------------------------------------------------------------------------|
| Language    | Go 1.24                                                                    |
| Database    | PostgreSQL 16                                                              |
| DB Driver   | [pgx v5](https://github.com/jackc/pgx)                                    |
| Migrations  | [golang-migrate](https://github.com/golang-migrate/migrate) (embedded SQL) |
| CLI         | [chzyer/readline](https://github.com/chzyer/readline) (history + tab-completion) |
| Terminal UI | [pterm](https://github.com/pterm/pterm)                                    |
| 2FA / TOTP  | [pquerna/otp](https://github.com/pquerna/otp)                              |
| QR Code     | [mdp/qrterminal](https://github.com/mdp/qrterminal)                        |
| Passwords   | [golang.org/x/crypto/bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) |
| Containers  | Docker + Docker Compose                                                    |

---

## Stopping the App

To stop all containers:

```bash
docker compose down
```

To stop and also delete the database volume (all data):

```bash
docker compose down -v
```
