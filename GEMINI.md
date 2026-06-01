# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```sh
# Run all tests
./scripts/test.sh
# or directly:
go test github.com/hatch-ed-com/ri-sdk-go/pkg/rapididentity

# Run a single test
go test github.com/hatch-ed-com/ri-sdk-go/pkg/rapididentity -run TestGetConnectFiles

# Format code
./scripts/fmt.sh
# or directly:
go fmt ./...
```

## Architecture

This is a Go SDK for the RapidIdentity REST API. All SDK code lives in `pkg/rapididentity/` as a single flat package with no external dependencies (stdlib only, `go 1.23.2`).

**Client initialization** (`RapidIdentity.go`): The `Client` struct wraps an `*http.Client` and holds either a `serviceIdentityKey` (for service identity auth) or a `*Session` (for user session auth). `New(Options)` creates the client and, if `RapidIdentityUser` credentials are provided, immediately POSTs to `/api/rest/sessions` to establish a session. `client.Close()` should always be deferred — it DELETEs the session if one exists.

**Per-endpoint files**: Each API endpoint is implemented in its own file named after the operation (e.g., `GetConnectFiles.go`). Each file defines:
- An `Input` struct for request parameters
- An `Output` struct for the response
- A method on `*Client` that calls `GenerateRequest`, does `httpClient.Do`, then `ReceiveResponse`, then JSON-unmarshals into the output struct

**Shared helpers** (`RapidIdentity.go`):
- `GenerateRequest` — builds an `*http.Request` with `Authorization: Bearer <token>`, `UserAgent`, and `Accept: application/json` headers
- `ReceiveResponse` — reads the response body and returns an error (`RapidIdentityError`) for non-2xx status codes
- `DoCustomRequest` — for API calls not yet wrapped by the SDK; the path is relative to `/api/rest/` (e.g., `"admin/workflow/resources"`)

**Error handling**: Errors are returned as `RapidIdentityError` (implements `error`) containing `Method`, `ReqUrl`, `Message`, `Reason`, and `Code`. Callers should use `errors.As(err, &riError)` to extract typed error details.

**Tests** (`*_test.go`): All tests use `httptest.NewServer` with a `http.ServeMux`. The `setup()` helper in `RapidIdentity_test.go` creates a test client and mux. Tests verify HTTP method, headers, query params, and response unmarshaling. Tests run in parallel (`t.Parallel()`).

**`MainProject` constant**: Use `rapididentity.MainProject` (value `"<Main>"`) when referring to the default Connect project — some endpoints treat an empty string differently from `<Main>`.

## Code Conventions

- Every exported method requires a Go doc comment with a `//meta:operation <METHOD> <PATH>` line that maps to the RapidIdentity OpenAPI spec.
- When releasing, update the `Version` constant in `RapidIdentity.go` — it is sent in the `User-Agent` header.
- The `examples/` directory contains runnable `main.go` programs demonstrating each SDK method.
