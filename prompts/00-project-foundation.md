# Agent Prompt — Greenfield Project Foundation (Read Only)

Use this prompt first because API v2 is a new implementation and this repository does not yet contain application code.

---

Act as a senior software architect and pragmatic AI-native engineering partner. Help me establish the minimum technical foundation for Attendance API v2.

## Context

- This is a greenfield API implementation. The current repository is empty of application code; do not assume an existing architecture, framework, database, CI pipeline, or code convention.
- Product and domain documents exist in `docs/` and are the source of truth for confirmed business decisions.
- API v2 will eventually replace API v1 for an existing client, but the existing v1 implementation/client may not be available in this workspace. Do not claim to have inspected them unless their source or verifiable documentation is actually provided.
- The first feature intended for implementation is Organization Onboarding.
- The shared product is an attendance platform for organizations in different business contexts. Keep the core industry-neutral.
- Technical decisions have not yet been approved.

## Read first

1. `AGENTS.md`
2. `docs/README.md`
3. `docs/product-decisions.md`
4. `docs/use-cases/principal-flow.md`
5. Relevant documents under `docs/core/` and `docs/actors/`
6. `specs/organization-onboarding/intent.md`

## Restrictions

- This is analysis and recommendation only. Do not create application code, install packages, initialize a framework, alter dependency manifests, or modify external systems.
- Do not silently choose a technology or turn a recommendation into an approved decision.
- Do not assume that API v2 must be a separate service, a separate repository, a `/v2` route prefix, or a replacement deployment. Explain reasonable options and consequences.
- Do not assume specific requirements about cloud provider, hosting, traffic volume, team size, data residency, budget, or existing v1 contract that are not documented. List material unknowns.
- Do not print or request secrets.
- Ground product facts in the docs and label technical advice as a proposal.

## Task

Recommend the smallest sound technical foundation needed to begin Organization Onboarding, with room to evolve the attendance product. Avoid designing the entire future platform prematurely.

Analyze these decision areas:

1. Language, runtime, and API framework.
2. API style and contract/documentation approach.
3. Persistence and database migration strategy.
4. Authentication and authorization for Platform Administrator and Organization Administrator.
5. Organization isolation / multi-tenant data protection.
6. Invitation tokens, 15-minute expiry, email delivery, and safe acceptance flow.
7. API versioning and future migration of the existing client from v1 to v2.
8. Validation, error response conventions, logging, auditability, and monitoring.
9. Automated testing, formatting/linting/type checks, local development, CI, and deployment assumptions.
10. Integration boundaries for attendance methods and location validation, without implementing biometric workflows now.

For each area, state whether it blocks the first Organization Onboarding implementation or can be deferred. Focus on decisions with near-term impact.

## Expected report

Produce a concise report with:

1. **Product constraints from docs** — confirmed decisions that affect technical design.
2. **Assumptions and missing context** — what is unknown because no application code or v1 source is available here.
3. **Decision table** — decision, plausible options, key trade-offs, recommended option, consequences, and whether approval is needed now.
4. **Minimum recommended baseline** — a coherent proposal, not a maximal architecture.
5. **Organization Onboarding implications** — which technical choices must be settled before implementing this feature.
6. **Deferred decisions** — decisions that can safely wait.
7. **Questions for the product owner** — ask only questions whose answers materially change the recommendation.
8. **Proposed next step** — the minimal set of decisions to approve and the ADRs or decision records to create.

Do not edit files or create specifications, plans, tasks, code, or ADRs during this first step. Stop after the report and wait for approval.
