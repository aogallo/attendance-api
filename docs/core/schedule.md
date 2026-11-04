# Schedule

## Purpose

Schedule represents the planned or expected attendance periods associated with a person within an organization. Every person must have an associated schedule.

## Business Context

Schedules may support different business needs. A teacher may have expected teaching periods, an employee may have work shifts, and a member may have expected attendance periods. These are contextual examples of a person's schedule, not separate kinds of person in the shared core.

The schedule belongs to the person, not to an establishment, group, or activity. Establishment opening hours describe when a location operates; they are separate from, and must not replace, the person's required schedule.

Person-specific schedule exceptions may be configured for an individual. The event types relevant to a person, where supported, must be interpreted together with the applicable schedule when producing reports.

## Business Purpose

Schedule describes when the person is expected to attend. All lateness and schedule-compliance reports must compare confirmed attendance events with that person's applicable schedule and configured event meanings. The first intended attendance report includes lateness, early departure, and absence.

## Scope

Schedule represents an expectation, not proof that a person was present. It is distinct from an attendance record, an attendance event type, an establishment, and an establishment's opening hours. Schedule mismatches do not by themselves block otherwise successful attendance captures.

The exact reporting calculations are unresolved. Lateness and early-departure reporting depend on the relevant schedule and configured event meanings; the system must not infer departure from the count or order of captures alone. Absence reporting requires a defined schedule period and cutoff. The handling of holidays, time zones, overlapping schedules, and personal exceptions also requires definition.

## Core Principle

Every person must have a schedule. It describes planned or expected attendance and must remain distinct from actual attendance and establishment opening hours. Other attendance contexts do not replace the person's schedule as the reference for lateness and schedule-compliance reporting. A schedule mismatch does not, by itself, prevent a successful attendance capture.
