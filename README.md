# Hearts

A Go web server experiment building the foundation for a **Spades**-style multiplayer card game (despite the repository's name), using a custom **flat-file, "No-SQL" JSON storage scheme** instead of a traditional database.

> **Status: Archived / Phase 1 dropped.** Per the repository description, this project was *"dropped at phase 1 due to inefficiency."* It contains user sign-up/sign-in scaffolding, session cookies, and a WebSocket endpoint, but **no card-game logic (dealing, turns, scoring, etc.) has been implemented.** This README documents the project as it stands for archival/reference purposes.

---

## Table of Contents

- [Overview](#overview)
- [Why It Was Dropped](#why-it-was-dropped)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [The "No-SQL" Storage Approach](#the-no-sql-storage-approach)
- [Routes](#routes)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Security Notes](#security-notes)
- [Known Issues](#known-issues)
- [Acknowledgments](#acknowledgments)
- [License](#license)

## Overview

Hearts is a Go (`net/http` + `gorilla/websocket`) backend that renders server-side HTML templates (`.gohtml`) for account creation and sign-in, then serves a chat/dashboard-style "main" page over a session cookie. Instead of a conventional database (SQL or a hosted document store), each user's data is written to its own flat file on disk using a custom `.gojson` extension — effectively a hand-rolled, per-user "No-SQL" data store.

The title of the main page template and the repository description ("A No-SQL version of spades") indicate the intended end product was a **Spades card game** with real-time messaging between players via WebSockets, though the game itself was never built.

## Why It Was Dropped

Based on the repository description and the current code, the flat-file storage approach ran into scalability/design limitations before any game logic was added:

- Every sign-in/sign-up request lists the entire `user/` directory (`ioutil.ReadDir`) and reads user files individually rather than through an indexed lookup.
- There is no locking, indexing, or querying layer — just direct file reads/writes per username.
- This works for a handful of test accounts but doesn't scale, hence "inefficiency" as the stated reason for stopping at phase 1.

## Tech Stack

| Concern | Technology |
|---|---|
| Language | Go `1.15` |
| HTTP routing | Standard library `net/http` (a `gorilla/mux` dependency is present in `go.mod` but not used in `routes.go`, which relies on `http.HandleFunc`) |
| Real-time messaging | [gorilla/websocket](https://github.com/gorilla/websocket) `v1.4.2` |
| Templating | Go standard library `text/template`, with `.gohtml` template files |
| Data storage | Custom flat-file JSON ("`.gojson`") storage — one file per user, no external database |
| Frontend | Bootstrap-based HTML/CSS/JS templates (see [Acknowledgments](#acknowledgments)) |

## Project Structure

```
hearts/
├── main.go                    # Entry point; struct definitions; page-rendering handlers; starts server on :2020
├── routes.go                  # Route table (http.HandleFunc registrations)
├── handlers.go                # Sign-up/sign-in/logout logic, user-page rendering, WebSocket endpoint
├── go.mod / go.sum             # Go module dependencies
├── helpers/
│   └── helper.go                # File read/parse helpers, base64 encode/decode, an (unused) HMAC signer
├── user/                         # "Database" directory — one .gojson file per registered user
│   ├── Noah24.gojson
│   ├── alvinv3.gojson
│   ├── jam_vaz24.gojson
│   └── spades_25.gojson
└── app/
    ├── sign-up.gohtml            # Registration form template
    ├── sign-in.gohtml            # Login form template
    ├── index.gohtml               # Main/dashboard page template (title: "SPADES")
    ├── 404-page.gohtml             # Not-found page template
    └── assets/                      # Bundled third-party frontend assets (CSS, JS, fonts, images, sounds)
```

## The "No-SQL" Storage Approach

Instead of a database, each user is represented by a JSON file named `<username>.gojson` inside the `user/` directory, shaped like:

```json
{
    "person": {
        "firstname": "...",
        "lastname": "...",
        "username": "...",
        "password": "...",
        "blocked": "true"
    }
}
```

- **Sign-up** (`signupAction` in `handlers.go`) marshals a new `person` struct to JSON and writes it to `user/<username>.gojson`, after scanning the directory to check the username isn't already taken.
- **Sign-in** (`signinAction`) reads the matching `.gojson` file, unmarshals it into a `map[string]interface{}`, and compares the submitted credentials against the stored `username`/`password`/`blocked` fields.
- **Session state** is a single cookie (`user_id`) set to the username — there is no server-side session store or token expiry.

## Routes

Registered in `routes.go`, served on **port 2020**:

| Path | Method(s) | Handler | Purpose |
|---|---|---|---|
| `/` | GET | `signup` | Renders the sign-up page |
| `/signin` | GET | `sigin` | Renders the sign-in page |
| `/signupAction` | POST | `signupAction` | Creates a new `.gojson` user file |
| `/signinAction` | POST | `signinAction` | Validates credentials and sets the session cookie |
| `/main` | GET | `renderUser` | Renders the main page for the logged-in user (reads cookie → `.gojson` file) |
| `/logout` | GET | `logOut` | Clears the session cookie |
| `/ws` | GET (upgrades to WebSocket) | `wsEndpoint` | Echoes messages back over a WebSocket connection (chat/game-message scaffold) |
| `/error` | GET | `errorPage` | Renders a 404-style error page |
| `/assets/*` | GET | — | Serves static frontend assets from `app/assets` |
| `*/..` (catch-all) | GET | `lostPage` | Fallback 404 handler for unmatched routes |

## Prerequisites

- **Go** `1.15` or later

## Getting Started

```bash
git clone https://github.com/joshuaomonemu/hearts.git
cd hearts
go mod tidy
go run .
```

The server starts on **`http://localhost:2020`**. Visit `/` to sign up, then `/signin` to log in and reach `/main`.

> Note: several `.gojson` files already exist in `user/` from earlier testing. These are sample/test data and should not be treated as real accounts — see [Security Notes](#security-notes).

## Security Notes

This project was an early-stage prototype and should **not** be used as-is beyond local experimentation:

- **Passwords are stored in plain text** inside the `.gojson` files (no hashing, e.g. bcrypt/argon2). Anyone with file access to `user/` can read every password directly.
- **Sample user files with plaintext test credentials are committed to the repository** (`user/*.gojson`). If any of these credentials were reused elsewhere, they should be treated as compromised and rotated. Going forward, test/user data directories like this should be excluded via `.gitignore`.
- `helpers/helper.go` includes an HMAC-signing helper (`getCode`) with a **hardcoded key (`"ourkey"`)**. It doesn't currently appear to be called anywhere in the request-handling code, but if it's reused later, the key should be moved to an environment variable / secrets manager.
- The WebSocket upgrader in `handlers.go` sets `CheckOrigin` to always return `true`, which disables the browser's origin check and would allow cross-origin WebSocket connections in a production deployment.
- Session identification is a single unsigned cookie holding the raw username, with no expiry, CSRF protection, or server-side session invalidation.

## Known Issues

- `go.mod`'s module path is `main.go`, which is unconventional (a proper module path, e.g. `github.com/joshuaomonemu/hearts`, is recommended).
- `gorilla/mux` is listed as a dependency in `go.mod` but the routing code uses the standard library's `http.HandleFunc` directly — the dependency appears unused.
- The catch-all "lost route" pattern in `routes.go` (`"*/.."`) is registered via `http.HandleFunc`, which does not support wildcard patterns the way the comment implies; unmatched paths are actually handled by Go's default `http.ServeMux` behavior rather than by `lostPage` as intended.
- `signinAction`'s existence check for a user file has inverted logic (it errors out when the file *is* found in the directory listing rather than when it's missing), which likely does not behave as intended.
- No automated tests are present.
- No `.gitignore` is present, and there's no license file.

## Acknowledgments

The bundled frontend templates under `app/` are third-party UI kits, not original work of this project:

- `app/index.gohtml` and its assets are mirrored from a Bootstrap **"Chatvia"** chat application template by [Themesbrand](https://themesbrand.com/), captured via HTTrack.
- `app/404-page.gohtml` is based on the **"Concept"** Bootstrap 4 admin dashboard template by [Colorlib](https://colorlib.com/).

If distributing this repository publicly, confirm the licensing terms of these bundled templates before reuse.

## License

No license file is currently included in this repository. If you intend for others to use, modify, or distribute this project, consider adding a `LICENSE` file (e.g. [MIT](https://choosealicense.com/licenses/mit/)) to clarify usage terms.
