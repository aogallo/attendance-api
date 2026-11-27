# P1 preserved implementation delivery slices
Deliver existing P1 in independently buildable PRs; preserve chore/p1-local-preservation intact.
Base: reviewed authentication layout 25951f0 (PR #4); source snapshot a19dd6b. No merges or provider operations.

## Specs

S1. Owner: “crea los PR para esas dos ramas”. Later selected “Dividila en PR más chicos — recomendado” after whole-candidate review exceeded native evidence budget. Publish smaller draft PRs, keep tests/docs with each behavior, preserve original branch without rewriting it.
S2. Preserve existing P1 behavior: protected organization creation, industry-neutral domain, issuer/subject Identity, input normalization, NUL rejection before persistence, safe HTTP errors and no invitation/person/email side effects. This is packaging, not new behavior.
S3. Keep unavailable original T4/T10 independent reports and accepted T11 historical limitation explicit. Native review is additional proof, not a replacement. No verifier retries, no fabricated merge-readiness or live-provider claims.

## Tasks

D1 | S1-S3 | inline packaging | in_progress | commit: pending | Organization domain/controller, HTTP Problem helpers, complete existing tests and scoped library documentation. Stack on #4.
D2 | S1-S3 | inline packaging | pending | commit: pending | Protected application wiring and embedded contract/docs with existing seam tests. Full vendored evidence may still exceed native review budget; do not omit it.
D3 | S1-S3 | inline packaging | pending | commit: pending | Transactional persistence, migration, Compose and real HTTP integration harness.
D4 | S1-S3 | inline packaging | pending | commit: pending | Executable and logging with tests and local run instructions.
D5 | S1-S3 | inline packaging | pending | commit: pending | Developer tooling/hook checks and final documentation reconciliation.

## Log

L1. Owner requested both branch PRs, then selected “Dividila en PR más chicos — recomendado”. Auth layout PR #4 published as draft; native review review-153e37b51fd99ed7 approved and acknowledged. Whole preservation review failed preflight lens_context_budget_exceeded without authority or reviewer execution; no retry of that candidate. Smaller packaging authorized; preserve a19dd6b intact. Historical functional checks passed including Docker lint/OpenAPI/Flyway/PostgreSQL; execute applicable checks for each isolated slice. Engram mirror unavailable due incompatible binary.
L2. D1 exact source set: internal/organization/*.go plus internal/httpapi/problem.go and types.go. No pgx/module change, application, runtime or Swagger. Expected roughly 550 changed lines including complete existing tests/docs; coherent library boundary retained rather than dropping coverage to meet the advisory 400-line target. Remaining delivery follows D2-D5 dependencies, without any merge authorization.
