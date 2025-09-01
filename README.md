# contractcheck

Desktop application built with **Go** and **Wails (React)**, structured to evolve with Clean/Hexagonal Architecture.

## Requirements
- Go 1.24.x
- Node.js (LTS) + npm
- Wails CLI
  ```bash
  go install github.com/wailsapp/wails/v2/cmd/wails@latest
  ```

## Quick start (development)
In the repository root:
```bash
make tidy            # go mod tidy
cd frontend && npm install && cd ..
make desktop-dev     # runs Wails dev (starts Vite + opens the app)
```

## Build (desktop)
```bash
make desktop-build   # produces ./build/bin/<outputfilename>
```

## Tests
```bash
make test            # go test ./...
```
