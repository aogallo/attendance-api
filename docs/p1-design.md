# P1 organization creation library

This slice provides the organization service/controller and route-registration library, tested with an in-memory repository double. It does not yet provide a running authenticated endpoint, PostgreSQL persistence or Swagger UI. Production application wiring and persistence follow separately.

## HTTP and input behavior

The controller implements the approved `POST /api/v1/organizations` contract. Successful creation returns `201` with `id`, normalized `name` and `initialAdministratorEmail`, `attendanceConfiguration: null`, and `createdAt`. No `Location` header points to an unimplemented GET operation.

| Input | Rule |
| --- | --- |
| Name | Trim outer whitespace; 1–200 Unicode characters. Reject NUL (U+0000) with `422`, `validation_failed` and field `name: invalid` before invoking persistence. |
| Email | Trim outer whitespace; simple address at most 254 UTF-8 bytes, preserve local part and lowercase domain; no display names or lists. |
| Request | JSON object, known fields only, one JSON value, maximum 16 KiB, JSON media type. |

Errors use `application/problem+json`, safe stable codes and `Cache-Control: no-store`. Malformed/unknown/trailing JSON returns `400`; oversized requests `413`; unsupported media `415`; field validation `422`. Repository errors map to `403` for forbidden, `503` for unavailable, or generic `500`, without exposing dependency errors. The library creates no identity, role assignment, invitation, Person, schedule or email.

## Package and authority boundaries

Shared domain types/interfaces live in `types.go`, errors in `errors.go`, and validation limits in `constants.go`, within the owning package. `organization.Identity` carries issuer and subject through `WithIdentity`/`IdentityFromContext`; incomplete or absent identities are rejected.

`RegisterRoutes` owns endpoint declarations and applies the guard passed by the application. These context helpers do not authenticate callers or grant platform privileges. Production wiring must supply trusted token/session verification, and the concrete repository must enforce platform authority. Tests inject identities and repositories; that is not a runtime authentication bypass or proof of hosted-provider enforcement.

## Verification

```sh
go test ./internal/organization ./internal/httpapi -count=1
go test -race ./... -count=1
go vet ./...
go build ./...
```

Existing tests cover input boundaries/normalization, exact success and NUL rejection responses, zero repository calls/state mutation for NUL rejection, safe error mapping, complete Identity propagation and route-guard application. Database integration and executable checks belong to later slices. The original whole-P1 independent report remains unavailable; this slice is not declared merge-ready merely because local tests pass.
