# Working with Project Documentation and Coding Agents

Spec-Driven Development (SDD) is an optional aid, not a requirement for contributing to this project. Contributors do not need prior SDD knowledge, coding agents, or specialized tools.

## Shared expectations, optional tooling

Everyone works from the same agreed behavior, acceptance criteria, and approved decisions, and verifies their changes with the available checks. These expectations do not require a particular development methodology.

Contributors may work manually from a task and submit a pull request, or use SDD and coding agents to assist them. Existing documents and specifications remain the shared source of truth; do not create separate requirements for manual and tool-assisted work. Plain Markdown is sufficient to read and update them.

The following lightweight SDD sequence and agent prompts are an optional guide, not mandatory steps for every contributor or small change:

`Project Foundation → Clarify → Specify → Plan → Tasks → Implementation → Verify`

## Foundation history and new decisions

The initial technical foundation is approved in `docs/technical-decisions/`. The foundation/clarification prompts describe how those decisions were reached; do not rerun them as if accepted choices were absent. Use read-only analysis for genuinely new decisions, distinguishing proposals from approved outcomes.

Approve only the technical decisions that block the first feature. Record accepted durable decisions in `docs/technical-decisions/` or ADRs, and leave the rest explicitly deferred.

## Clarify a feature or behavior change

Read the current feature specification and its relevant domain, policy, and decision links. For onboarding, `intent.md` preserves original motivation and clarification context, while `spec.md` is the proposed consolidation. Resolve material gaps without reopening unrelated confirmed decisions. A proposed document does not become accepted because it was moved or indexed.

## Specify

Maintain `specs/<feature>/spec.md` as living user-visible behavior and testable acceptance criteria, not a disposable implementation checklist. Its status distinguishes proposed from accepted behavior. Reference cross-cutting policies rather than independently redefining them. Do not put database schemas, classes, endpoints, or framework details into behavioral requirements unless needed to define an approved external API contract.

## Plan and tasks

Ask the agent to propose the implementation design and break it into small tasks. Review security, organization isolation, invitation expiry, email delivery, tests, and future client migration implications. Keep this plan consistent with approved technical decisions.

## Implementation and verification

Implement only the approved scope. Run the actual tests and checks available. Report what ran, what passed, and what was not verified. If implementation reveals a missing product decision, return to Clarify/Specify rather than guessing.

## One owner per kind of information

| Information | Current owner | Maintenance rule |
| --- | --- | --- |
| Domain meanings and role boundaries | `docs/core/`, `docs/actors/` | Maintain definitions; reference them from features. |
| Confirmed general product rules and open product questions | `docs/product-decisions.md` | Keep product decisions here; link specialized policies instead of duplicating their numeric values and full rules. |
| Administrative access policy | `docs/security/administrative-access.md` | Own passwords, MFA, sessions, recovery, and authentication-link rules. |
| Email delivery policy | `docs/operations/email-delivery.md` | Own current sender/language, delivery responsibilities, dated quotas, and quota failure handling. |
| Significant technical choice and its rationale | `docs/technical-decisions/` | Existing TD records are ADRs. Preserve IDs/links, alternatives, approval status, and consequences. |
| Feature behavior and acceptance checks | `specs/<feature>/spec.md` | Remains useful after implementation; references policies and domain definitions. Until accepted, it does not override confirmed sources. |
| Delivery sequence, tasks, and implementation evidence | Feature plan or adopted task/PR workflow | Temporary execution context, not a second rulebook. A separate plan is optional. |
| Agent entry points and commands | `AGENTS.md` | Concise routing and safeguards; contributors without agents use the same underlying docs. |

An ADR answers why a significant solution was selected, not every detail of what a feature must do. A new implementation finding first goes to the appropriate owner: product uncertainty for clarification, current behavior to its spec/policy, or an architectural proposal to an ADR. Do not create one ADR per field, timeout, helper, or routine edit. A proposal needs approval before it is recorded as accepted.

Keep summaries clearly labeled and link the owner. When an approved rule changes, update that source, affected feature criteria, contract/tests and relevant references together. Do not resolve contradictions using file recency alone; preserve approval evidence and raise unresolved conflicts. Accepted ADRs retain the original decision rationale; a materially changed decision is linked as superseding/superseded rather than silently rewriting its history. Existing compatibility headings may point to the current policy.

## Plan closeout and archive

1. Keep an active plan proposed or in progress until the stated outcomes and checks have actually been observed. Never archive unfinished work merely to reduce file count.
2. At completion, update the living spec, contract, tests and operational docs to describe the delivered behavior. Extract significant accepted technical choices into ADRs only when needed; do not leave durable rules exclusively in the plan.
3. Record the outcome, verification evidence, linked commits/PRs, and remaining/deferred work in the plan or adopted delivery record. Mark the plan completed; unresolved outcomes remain explicit.
4. A completed plan may remain clearly marked or move to `specs/<feature>/archive/`, with an updated reference from its former entry point. Archive is historical context, not a source of current requirements. Git history also preserves removed transient detail when no useful archive is needed.

The current onboarding specification and implementation plan have separate owner approval for local implementation. The plan remains active and unimplemented; its remaining design/security gates and operational authorization still apply. This documentation-lifecycle decision did not itself approve the delivery boundary or mark the feature implemented. No new tool, generator, framework, or mandatory agent workflow is introduced.
