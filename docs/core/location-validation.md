# Location Validation

## Purpose

Location Validation describes an optional check used to determine whether an attendance capture is associated with an acceptable location in relation to an establishment.

## Business Context

Some organizations need location context for attendance. Others may not need it, including those using a different capture process or supporting remote or virtual activities.

## Business Purpose

When enabled and location information is available, Location Validation provides additional context for the organization's attendance process without making location a mandatory requirement for every customer.

## Scope

Location Validation is applied only when enabled by the organization. If location information is unavailable in the mobile or web application, the person can still record attendance because attendance capture is the core product capability. In that case the record must not imply that the location was successfully validated.

The acceptable radius, its configuration scope, and the outcome when an available location is outside the permitted area or too imprecise remain unresolved. No universal radius should be assumed.

Location Validation is separate from identity verification. Being near an establishment does not itself identify a person, and identifying a person does not by itself prove physical presence there.

## Core Principle

Location validation is optional. An unavailable location does not block attendance, but the platform must distinguish an unavailable location from a successful location validation.
