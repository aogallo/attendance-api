# Attendance Capture Attempt

## Purpose

Attendance Capture Attempt represents an effort to capture or register attendance that did not result in a confirmed Attendance record.

## Business Context

An attempt may be unsuccessful because the person could not be identified, the provided credential was not accepted, a capture could not be completed, or a configured attendance validation was not satisfied. The actual reasons depend on the selected method and workflow.

## Business Purpose

Recording unsuccessful attempts allows an organization to review capture problems, understand recurring issues, and distinguish an unsuccessful attempt from a person's confirmed attendance.

## Scope

This concept covers unsuccessful attendance captures and their business outcome. It is separate from Attendance: an unsuccessful attempt must not be counted as confirmed attendance.

The attempt log is not a general activity log for the entire platform, and it is not an access-control decision. Membership checks, permission to enter, and other industry-specific eligibility decisions remain outside the shared attendance core.

## Core Principle

An unsuccessful capture must remain distinguishable from a successful attendance record. The platform must not create a confirmed Attendance record merely because an attempt was made.
