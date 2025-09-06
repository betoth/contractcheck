# ==== OS-aware shell & helpers ====
# Use cmd.exe on Windows and /bin/sh elsewhere. Implement quiet RM/MKDIR.
ifeq ($(OS),Windows_NT)
  SHELL := cmd.exe
  .SHELLFLAGS := /Q /C
  # Use PowerShell for quiet recursive removal
  RM_DIR  = powershell -NoProfile -Command "if (Test-Path '$(1)') { Remove-Item '$(1)' -Recurse -Force -ErrorAction SilentlyContinue }"
  RM_FILE = powershell -NoProfile -Command "if (Test-Path '$(1)') { Remove-Item '$(1)' -Force -ErrorAction SilentlyContinue }"
  MKDIR_P = if not exist "$(1)" mkdir "$(1)"
else
  SHELL := /bin/sh
  .SHELLFLAGS := -c
  RM_DIR  = rm -rf "$(1)"
  RM_FILE = rm -f "$(1)"
  MKDIR_P = mkdir -p "$(1)"
endif

# ==== Basic metadata (kept for future use) ====
VERSION  ?= 0.1.0
COMMIT   := $(shell git rev-parse --short HEAD 2>/dev/null || echo none)

# ==== Tools ====
GO ?= go

# ==== Targets (desktop/Wails only for now) ====
.PHONY: help tidy fmt vet test cover desktop-dev desktop-build desktop-clean print-version

help:
	@echo "Targets:"
	@echo "  tidy           - go mod tidy"
	@echo "  fmt            - go fmt ./..."
	@echo "  vet            - go vet ./..."
	@echo "  test           - go test ./... (unit tests)"
	@echo "  cover          - go test w/ coverage summary"
	@echo "  desktop-dev    - run Wails dev (desktop app)"
	@echo "  desktop-build  - build Wails app (generates bindings + bundles frontend)"
	@echo "  desktop-clean  - remove only Wails build/bin artifacts"
	@echo "  print-version  - show VERSION and COMMIT"

tidy:
	$(GO) mod tidy

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

# --- tests (quiet on Windows) ---
ifeq ($(OS),Windows_NT)
test:
	powershell -NoProfile -Command "$(GO) test ./..."

cover:
	powershell -NoProfile -Command "$(GO) test -coverprofile=coverage.out ./...; $(GO) tool cover -func=coverage.out; if (Test-Path 'coverage.out') { Remove-Item 'coverage.out' -Force }"
else
test:
	$(GO) test ./...

cover:
	$(GO) test -coverprofile=coverage.out ./...
	$(GO) tool cover -func=coverage.out
	@$(call RM_FILE,coverage.out)
endif

# --- desktop (Wails) ---
desktop-dev:
	wails dev

desktop-build:
	wails build

# Clean only the bin folder inside build (preserve icons and manifests)
desktop-clean:
	@$(call RM_DIR,build/bin)

print-version:
	@echo VERSION=$(VERSION)
	@echo COMMIT=$(COMMIT)
