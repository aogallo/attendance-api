# TD-003: Owner-hosted initial delivery and deferred customer installations

- **Status:** Accepted
- **Date:** 2026-11-03
- **Owner:** Product/Engineering owner

## Context

The owner intends to host and operate the initial attendance platform, including its authentication integration. A possible later commercial model would deliver the software for a customer to operate using that customer's enterprise identity provider. The owner requested that this scenario be preserved for future analysis without implementing it in the first delivery.

## Options considered

### Option A — Owner-hosted initial platform; defer customer-operated installations

- Benefits: one initial operating model and a smaller implementation scope, while preserving the rationale for future identity-provider integration.
- Costs and risks: customer-operated delivery will require a later design and compatibility assessment; it is not guaranteed to work through configuration alone.
- Onboarding implications: build for the owner's hosted platform and its selected authentication integration, not simultaneous customer deployment variants.

### Option B — Support owner-hosted and customer-operated delivery immediately

- Benefits: immediate flexibility for buyers who require their own infrastructure and identity provider.
- Costs and risks: adds installation, upgrades, identity-provider compatibility, security configuration, and support obligations before they are needed.
- Onboarding implications: expands initial implementation and verification scope; not selected for the first delivery.

## Decision

### Initial scope

- The first delivery is hosted and operated by the project owner.
- Authentication is part of that owner-operated platform. [TD-004](004-initial-administrative-authentication.md) subsequently selects Supabase Auth for the initial administrative accounts; this does not approve custom-built authentication, Keycloak, or customer enterprise-provider integration.
- Cloud vendor, infrastructure topology, deployment automation, and production sizing remain separate decisions.
- Do not implement customer-operated installations, per-customer identity-provider connections, or a generic multi-provider framework as part of the initial scope solely to anticipate this scenario.

### Deferred scenario

A future buyer may receive an independently operated installation of the software and authenticate its administrators using its existing enterprise identity provider. Its hosting could be on the customer's premises or in its cloud; no specific deployment environment is approved now.

The current Controller/Service/Repository boundaries should keep provider-specific authentication details out of business rules. Organization isolation, application permissions, invitations, and attendance participation must not be replaced by assumptions about the initial provider. This is a design consideration, not a claim that every provider is interchangeable or already supported.

### Revisit when a concrete customer requires it

Before approving or estimating that delivery, assess:

1. The customer's identity provider, supported protocols such as OIDC or SAML, and whether direct integration or a broker is appropriate.
2. Identity verification, stable external account identifiers, account linking, administrator provisioning, and mapping to the application's organization-scoped permissions. Successful corporate login must not automatically grant administrative authority.
3. Configuration, secret management, logout, revocation, session behavior, and environment-specific trust boundaries.
4. Hosting requirements, installation packaging, database provisioning, Flyway migrations, upgrade compatibility, backups, observability, and recovery.
5. Who operates and supports each component, identity-provider licensing and costs, software distribution/licensing terms, and data-handling requirements.
6. Acceptance and integration tests against the customer's actual provider and deployment environment.

Do not assume one identity provider per school, support for multiple customer providers in the initial hosted platform, delivery of source code, or a particular license or support agreement. Those are separate future decisions.

## Consequences

- Positive consequences: the initial hosted delivery remains focused and the future commercial scenario has a documented starting point.
- Trade-offs: customer-specific deployment and authentication may require new adapters or changes later; no zero-effort portability is promised.
- Follow-up work: complete the Supabase Auth integration design under TD-004 and select the operating environment. Reopen the deferred scenario only when a concrete customer requirement provides the necessary context and the owner approves the scope.

## Related decisions

- [TD-001: Go HTTP API and SQL-owned persistence](001-api-foundation.md).
- [TD-002: Feature-oriented layers and an OpenAPI contract](002-application-structure-and-api-contract.md).
