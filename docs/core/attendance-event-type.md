# Attendance Event Type

## Purpose

Attendance Event Type describes the purpose assigned to a successful attendance capture when an organization's configured workflow distinguishes between different kinds of events.

## Business Context

Depending on an organization's attendance process and the configuration applicable to a person, event types may include work entry, work departure, break start, and break end. Other organizations may need a simpler model in which successful captures are recorded as separate events without assigning an entry or departure meaning.

The type of event is distinct from the Attendance Method. For example, facial recognition or a fingerprint may identify the person, while a button or other explicit selection indicates the purpose of the capture.

## Business Purpose

Attendance Event Type gives an organization a clear way to interpret a successful capture when its workflow requires distinctions such as arrival, departure, or a break. This can support schedule comparisons and reports without making those distinctions mandatory for every organization.

## Scope

An event type is determined by the configured attendance workflow and the settings applicable to the person, using an explicit selection or another rule that is deliberately defined for the product. The platform must not infer entry, departure, or break type merely from the order or number of previous captures.

Each successful capture creates a separate Attendance event. If no specific event type applies, the event is still recorded, but the platform must not invent an entry or departure classification. The exact event types supported in the first release and the way a person or device selects them remain unresolved.

Attendance Event Type does not identify the person, validate their identity, define their schedule, or decide whether they may enter a facility.

## Core Principle

Event meaning must follow the configured business process rather than assumptions about capture sequence. Event types are configurable attendance semantics, not identification methods and not access-control decisions.
