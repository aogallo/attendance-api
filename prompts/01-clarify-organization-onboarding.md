# Agent Prompt — Clarify Organization Onboarding

Use this after the owner approves the relevant technical foundation decisions.

---

Act as a product-minded software engineer. Help clarify Organization Onboarding before writing its feature specification.

## Read first

- `AGENTS.md`
- `docs/product-decisions.md`
- `docs/use-cases/principal-flow.md`
- `specs/organization-onboarding/intent.md`
- Accepted technical decisions in `docs/technical-decisions/`

## Goal

Find remaining product ambiguities that could change Organization Onboarding behavior. Ask focused questions and explain why each answer matters. Do not write implementation code or reopen approved technical decisions unless a contradiction or new material risk is found.

## Rules

- Treat decisions marked confirmed in product documents as authoritative.
- Do not turn proposals/open questions into requirements silently.
- Ask the highest-impact question first and ask small batches, preferably one question at a time.
- Separate product decisions from technical decisions.
- Defer details that do not block the first release instead of expanding scope.
- Do not ask questions already answered in the documents.
- Do not modify files until asked to update the intent or create the specification.

## Expected outcome

Summarize confirmed behavior, remaining questions and whether they block specification/implementation, explicit out-of-scope behavior, and the smallest useful first release. Do not implement anything.
