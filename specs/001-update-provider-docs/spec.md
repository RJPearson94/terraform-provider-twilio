# Feature Specification: Update Provider Documentation to Align with Terraform Best Practices

**Feature Branch**: `001-update-provider-docs`
**Created**: 2026-04-03
**Status**: Draft
**Input**: User description: "update provider documentation and generation to align to terraform provider best practices"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Attribute Descriptions in Registry (Priority: P1)

A Terraform user navigates to the Twilio provider page on the Terraform Registry to
understand what arguments a resource accepts. Today, the Argument Reference table lists
field names with no descriptions — users must consult the Twilio API docs separately to
understand each field's purpose, valid values, and whether it is computed or configurable.

After this feature, every attribute in every resource and data source has a description
explaining its purpose, constraints, valid values, and computed behaviour — visible
directly in the Registry without leaving the page.

**Why this priority**: This is the primary user-facing gap. It directly prevents
self-service use of the provider without additional external documentation.

**Independent Test**: Generate docs locally and verify that every entry in the Argument
Reference and Attribute Reference tables for any resource has a non-empty description.

**Acceptance Scenarios**:

1. **Given** a user views `twilio_studio_flow` in the Registry,
   **When** they look at the `status` argument,
   **Then** they can read that it accepts `draft` or `published` without leaving the page.

2. **Given** any computed attribute (e.g. `sid`, `date_created`, `url`),
   **When** a user reads its description,
   **Then** the description states that the value is assigned by Twilio and cannot be set.

3. **Given** an enum attribute (e.g. `inbound_method` on a messaging resource),
   **When** a user reads its description,
   **Then** all valid values are listed in the description.

---

### User Story 2 - Single-Command Documentation Generation (Priority: P2)

A contributor updates a resource schema — adds an attribute, changes an enum's valid
values, or corrects a default. Today the `docs/` markdown must be updated manually,
causing drift between code and documentation over time.

After this feature, running one command regenerates all documentation from schema
descriptions and examples, keeping docs and code permanently in sync.

**Why this priority**: Without automated generation, descriptions added in P1 will drift
from the `docs/` markdown as the schema evolves, defeating the purpose of adding them.

**Independent Test**: Change one schema attribute description, run the generation command,
and verify the change appears in the corresponding `docs/` file without manual editing.

**Acceptance Scenarios**:

1. **Given** a contributor has updated a schema description,
   **When** they run `go generate ./...`,
   **Then** the corresponding `docs/resources/` or `docs/data-sources/` file is updated
   to reflect the new description without any further manual steps.

2. **Given** the generation command has been run,
   **When** a contributor runs `make terrafmt-docs`,
   **Then** no formatting changes are required (the generated output is already clean).

3. **Given** CI runs on a pull request that changes a schema attribute,
   **When** the contributor did not run `go generate ./...`,
   **Then** CI detects a diff in `docs/` and fails, blocking the merge.

---

### User Story 3 - Valid, Runnable Examples (Priority: P3)

A Terraform user copies an example from the Registry to bootstrap their configuration.
Examples may reference argument names that do not match the current schema, or omit
required arguments, causing immediate plan errors.

After this feature, all examples in `examples/` are valid HCL that reflects the current
schema and can be used directly without modification.

**Why this priority**: Broken examples undermine trust, but the current examples are
mostly correct — this is a validation and correction pass, not a rewrite.

**Independent Test**: Run `make validate-all-examples` against all examples and verify
the command exits cleanly.

**Acceptance Scenarios**:

1. **Given** all examples in `examples/` are checked,
   **When** `make validate-all-examples` is run,
   **Then** it exits with code 0 and reports no errors.

2. **Given** a new resource is added as part of a future PR,
   **When** the contributor opens that PR,
   **Then** at least one example for that resource exists in `examples/` and passes
   validation.

---

### Edge Cases

- Attributes shared across many resources (e.g. `account_sid`, `date_created`, `url`)
  should have consistent description wording — establish a convention for these common
  fields so descriptions are uniform without being duplicate prose.
- The `twilio_serverless_deployment` resource intentionally creates a new deployment on
  delete rather than removing one — its description must reflect this API limitation so
  users are not surprised by accumulating deployments.
- Any attribute that accepts a free-form JSON value (e.g. `definition` on studio flow)
  must describe the expected JSON structure or link to where the schema is documented.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: Every `schema.Schema` entry for a user-facing attribute in every resource
  and data source MUST include a non-empty `Description` string.
- **FR-002**: Descriptions for `Computed: true` attributes MUST state that the value is
  assigned by Twilio and is read-only.
- **FR-003**: Descriptions for enum-constrained attributes MUST list all valid values.
- **FR-004**: A `//go:generate` directive MUST exist in the provider so that
  `go generate ./...` triggers regeneration of all documentation in `docs/`.
- **FR-005**: The generation tool MUST produce output that matches the existing `docs/`
  structure — resources in `docs/resources/` and data sources in `docs/data-sources/`.
- **FR-006**: CI MUST detect when `docs/` is out of date with the schema and fail the
  build on any PR that changes schema without regenerating docs.
- **FR-007**: All files in `examples/` MUST pass HCL validation as reported by
  `make validate-all-examples`.
- **FR-008**: `make terrafmt-docs` MUST produce no diff after generation runs — generated
  output must be clean without a separate formatting pass.

### Key Entities

- **Schema attribute**: A field defined in a `schema.Schema` struct with Type,
  Required/Optional/Computed flags, and a Description string.
- **Documentation file**: A markdown file in `docs/resources/` or `docs/data-sources/`
  that documents one resource or data source for the Terraform Registry.
- **Example**: A `.tf` file in `examples/` that demonstrates a working resource
  configuration using current argument names and required fields.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: 100% of schema attributes across all resources and data sources have a
  non-empty Description — zero undescribed attributes remain.
- **SC-002**: `go generate ./...` completes without error and produces a clean `docs/`
  directory (confirmed by `make terrafmt-docs` reporting no diff).
- **SC-003**: CI enforces documentation currency — a PR with a schema change and no doc
  regeneration is blocked automatically.
- **SC-004**: `make validate-all-examples` exits with code 0 against all examples.
- **SC-005**: A Terraform user can understand any resource argument's purpose, valid
  values, and behaviour from the Registry page alone — verified by manual review of three
  representative resources (studio flow, messaging service, taskrouter workspace).

## Assumptions

- The existing `docs/` markdown files will serve as the foundation; the generation tool
  will overwrite them from schema descriptions and examples.
- `tfplugindocs` is compatible with `terraform-plugin-sdk/v2` providers and will be used
  as the generation tool via `go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs`.
- Adding `Description` fields to schema attributes does not change runtime provider
  behaviour and does not require new acceptance tests.
- The existing `make validate-all-examples` target will be used as-is for FR-007.
- CI additions are extensions of the existing GitHub Actions workflow, not a new pipeline.
- Work will proceed service-by-service (studio → messaging → taskrouter → conversations
  → verify → serverless → proxy → sync → remaining services) to allow incremental PR
  review without one enormous change.
