# Working with a Coding Agent

Spec-Driven Development (SDD) is an optional aid, not a requirement for contributing to this project. Contributors do not need prior SDD knowledge, coding agents, or specialized tools.

## Shared expectations, optional tooling

Everyone works from the same agreed behavior, acceptance criteria, and approved decisions, and verifies their changes with the available checks. These expectations do not require a particular development methodology.

Contributors may work manually from a task and submit a pull request, or use SDD and coding agents to assist them. Existing documents and specifications remain the shared source of truth; do not create separate requirements for manual and tool-assisted work. Plain Markdown is sufficient to read and update them.

The following lightweight SDD sequence and agent prompts are an optional guide, not mandatory steps for every contributor or small change:

`Project Foundation → Clarify → Specify → Plan → Tasks → Implementation → Verify`

## First session: project foundation (greenfield)

The repository has no application implementation yet. Use `prompts/00-project-foundation.md` in analysis/read-only mode. Ask the agent to recommend a minimal technical foundation using the product documents, state unknowns, and separate recommendations from approved decisions. Do not let it initialize the project or write application code in this step.

Approve only the technical decisions that block the first feature. Record accepted durable decisions in `docs/technical-decisions/` or ADRs, and leave the rest explicitly deferred.

## Second session: clarify Organization Onboarding

Use `prompts/01-clarify-organization-onboarding.md`. The agent should read the approved technical decisions as context but must not reopen confirmed product decisions unless a conflict is found. Resolve material product questions before writing the specification.

## Specify

Create `specs/organization-onboarding/spec.md` with user-visible behavior and testable acceptance criteria. Do not put database schemas, classes, endpoints, or framework details into the behavioral specification unless needed to define an external API contract.

## Plan and tasks

Ask the agent to propose the implementation design and break it into small tasks. Review security, organization isolation, invitation expiry, email delivery, tests, and future client migration implications. Keep this plan consistent with approved technical decisions.

## Implementation and verification

Implement only the approved scope. Run the actual tests and checks available. Report what ran, what passed, and what was not verified. If implementation reveals a missing product decision, return to Clarify/Specify rather than guessing.

## Decision rule

- Product truth lives in `docs/`.
- Approved technical decisions live in `docs/technical-decisions/` (or an accepted ADR convention).
- Feature behavior lives in `specs/<feature>/spec.md`.
- When an implementation plan is useful, it may live in `plan.md`; a separate plan file is not required for every contribution.
- Tasks and verification evidence live alongside the feature specification or in the project's adopted workflow.
- `AGENTS.md` should remain concise and contain reusable agent instructions, not full product/architecture essays.
