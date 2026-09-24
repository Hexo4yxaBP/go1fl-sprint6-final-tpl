# Morse Converter API

Small Go HTTP service that converts text ↔ Morse code via a file upload form.

Course project (Yandex Practicum, Sprint 6 final). Useful as a portfolio sample of a tiny Go API: handlers, service layer, and a reusable `pkg/morse` package.

## What it does

- Serves a simple HTML upload form at `/`
- Accepts a text file via `POST /upload` (`multipart/form-data`, field `myFile`)
- Auto-detects direction:
  - mostly Morse (`.`, `-`, spaces) → converts to Cyrillic text
  - otherwise → converts text to Morse
- Alphabet: Russian letters, digits, and common punctuation (`pkg/morse`)

## Layout

```
cmd/main.go           # entrypoint
internal/server/      # HTTP server setup
internal/handlers/    # / and /upload
internal/service/     # MorseConvert direction logic
internal/config/      # optional JSON config loader
pkg/morse/            # encode / decode library
index.html            # upload form
cmd/server.config     # sample timeouts / addr (not loaded by default)
```

## Run locally

Requires Go 1.24+.

```bash
# from repo root
go run ./cmd
```

Open http://localhost:8080/

Or build a binary:

```bash
go build -o ./bin/morse-api ./cmd
./bin/morse-api
```

## API

| Method | Path      | Body                         | Result                          |
|--------|-----------|------------------------------|---------------------------------|
| GET    | `/`       | —                            | HTML upload form                |
| POST   | `/upload` | `myFile` (multipart file)    | Converted text/Morse in body    |

Example with curl:

```bash
echo 'ПРИВЕТ' > sample.txt
curl -F "myFile=@sample.txt" http://localhost:8080/upload
```

## Notes

- Listen address is hardcoded to `:8080` in `internal/server` (JSON config exists but is commented out).
- On upload the handler also writes a result file in the working directory (UTC timestamp + original extension).
- Module path in `go.mod` still points at the course template module name.

## Status

Educational / portfolio sample — not production hardening (no auth, no size limits beyond multipart memory, result files written next to the process).
