# Organization

## Purpose

Organization represents a customer or institution that uses the platform to manage its attendance process.

## Business Context

An organization may be a school, company, gym, pharmacy, retailer, or another type of institution. Its context influences the language it uses, the way attendance is taken, and the optional capabilities it needs.

An organization may operate one or more establishments and maintain its own people, attendance records, schedules, and attendance configuration.

Creating an organization requires its name and the email address of the intended initial Organization Administrator, even if that administrator's account does not exist yet. Saving the email identifies the intended recipient for a later invitation; it does not automatically create an account, grant permissions, create a Person or schedule, or issue an invitation. The invitation's validity period starts when its own record is created, not when the organization is created.

The Platform Administrator creates the organization before inviting its Organization Administrators. Creation does not require an immediate invitation; the organization may exist before any administrator has accepted one. Later invitations refer to that existing organization, and administrative authority is granted only through successful acceptance. Organization existence alone does not establish readiness for attendance operations.

## Business Purpose

Organization establishes the business boundary for attendance management. It provides the context in which people are registered, attendance is recorded, and relevant settings are configured.

## Scope

Organization represents the customer and its operating context; it does not contain every business concept used by that customer. For example, a gym's organization is not itself a membership or a branch, and a company's organization is not itself an employment contract or payroll process.

Each organization manages its own people and attendance information. Configuration for one organization must not be treated as the configuration for another organization.

## Core Principle

The shared core must support organizations from different industries without requiring the meaning of Organization to be redesigned for each one. Industry-specific processes may extend the platform, but Organization remains the boundary for that customer's attendance context and information.
