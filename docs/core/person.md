# Person

## Purpose

Person represents an individual whose participation or presence may be recorded through an organization's attendance process.

## Business Context

The concept is shared across organizations in different industries, including schools, companies, gyms, and other settings that need to record attendance. Depending on its context, an individual may be referred to as a student, teacher, employee, member, visitor, or participant.

These designations describe the person's role or relationship within a particular organization. They do not change the meaning of Person in the shared attendance core.

## Business Purpose

Person provides a consistent way for an organization to identify the individual associated with an attendance record, regardless of the organization's industry or terminology.

## Scope

Person represents the individual within the context of one organization. The organization manages its own person record and related information. Person information is not shared or automatically linked across organizations, even when two organizations may have records for the same real-world individual.

Person does not, by itself, represent a student enrollment, employment relationship, gym membership, schedule, or attendance record. Those concepts belong to their corresponding business contexts or processes.

Every person must have an associated schedule within their organization. The schedule belongs to the person, not the establishment. All lateness and schedule-compliance reports use the person's applicable schedule as their reference; establishment opening hours do not replace it. The assignment process and timing remain to be specified.

Person represents attendance participation, not administrative authority or a login account. An Organization Administrator who only administers does not require a Person record or schedule. If that administrator also records attendance, the same individual has an associated Person record within the organization, with a required schedule and the same attendance and reporting rules as other people; a second login account is not required merely to combine these responsibilities. The Platform Administrator has no attendance participation or schedule for its platform role.

In the first attendance release, an ordinary Person records attendance through facial recognition on an organization shared device in administrator-enabled attendance mode, without an individual login account or sign-in. Registration of the Person and facial enrollment are distinct from creating an administrative account. If that person later becomes an Organization Administrator, administrative access is granted through the invitation/account flow while their Person, schedule, and attendance history remain distinct from the administrative role. Personal-device capture is deferred.

## Core Principle

Person must remain independent of any specific industry. Each organization controls its own person information, and the core must not assume that a person record or its information is shared with another organization.
