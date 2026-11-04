# Attendance API v2 — Start Here

This is a greenfield API project. Spec-Driven Development (SDD) is an optional aid, not a requirement for contributing. Contributors do not need SDD knowledge, coding agents, or specialized tools. See [Working with a Coding Agent](docs/agent-workflow.md) for the shared expectations and optional workflow.

For contributors choosing to work with a coding agent, the suggested order is:

1. Copy/merge these docs into the new repository.
2. Start with `prompts/00-project-foundation.md` in read-only/analysis mode. Ask the agent to compare reasonable technical options using the domain requirements, identify the minimum decisions needed for Organization Onboarding, and recommend a minimal foundation. No code changes yet.
3. Review the report. Approve or defer decisions explicitly and record accepted ones in `docs/technical-decisions/`.
4. Use `prompts/01-clarify-organization-onboarding.md` to resolve remaining business questions for the first feature.
5. Create the behavior-first spec with `prompts/02-create-organization-onboarding-spec.md`.
6. Review the plan and tasks before enabling implementation.

API v2 identifies the second generation of the implementation, not an HTTP path. It does not prescribe a `/v2` route prefix or any other API versioning strategy; those technical decisions require explicit owner approval.

The existing API v1 and client are outside this repository unless separately supplied. The agent cannot inspect or verify their internals from this greenfield repository alone.
