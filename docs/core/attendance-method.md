# Attendance Method

## Purpose

Attendance Method describes how a person is identified or how the information needed to capture attendance is provided.

## Business Context

Organizations may enable multiple methods, such as PIN, fingerprint recognition, and facial recognition. Enabled methods are configured at the organization level and are available across the organization's establishments by default.

Facial recognition is the first method to be delivered and the product's primary attendance method. It is the only attendance method in the first attendance release. PIN and fingerprint recognition are deferred to later releases and are not available methods or fallbacks in the first release. This scope does not automatically enable facial recognition for an organization, define fallback behavior, or select a biometric implementation. The core remains independent of the method used so additional methods can be supported later.

A person uses one enabled method for each capture. If only one method is enabled, it is the only available method. If several methods are enabled, the policy for assigning a method to a person or allowing the person to choose one has not yet been decided.

If an organization disables a method, future captures using that method cannot produce confirmed attendance. Configuration changes do not rewrite historical attendance records. The treatment of existing method enrollments if the method is later re-enabled remains unresolved.

## Business Purpose

Attendance Method allows the platform to support different identification and capture options while preserving a consistent meaning for a confirmed Attendance record.

## Scope

Attendance Method describes how a person is identified or how an attendance capture begins. It does not define the purpose of the attendance event. For example, facial recognition may identify a person, while a separate workflow selection indicates that the event represents an entry, departure, or break.

It does not define the entire attendance workflow, the result of a capture attempt, the attendance record itself, or whether a person is allowed to enter a facility.

The enrollment, confirmation, recovery, replacement, and re-enrollment process may differ by method and still needs to be defined. Email confirmation by code is a candidate proposal; it is not yet an approved requirement and does not itself enroll a PIN, fingerprint, or facial reference.

## Core Principle

The shared attendance core must not depend on a single identification or capture method. Organizations can enable supported methods that fit their needs without changing the fundamental meaning of Attendance. A disabled method must not be accepted for future captures.
