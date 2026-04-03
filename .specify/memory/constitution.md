<!--
SYNC IMPACT REPORT
==================
Version change: 1.0.0 → 1.1.0

Modified principles:
  - II. Consistent Resource Pattern — added: immutable fields MUST use ForceNew, sensitive
    attributes MUST be marked Sensitive: true
  - IV. Acceptance-Test Coverage — added: disappears test mandatory, randomised names,
    multi-step update testing, regression two-commit workflow

Modified sections:
  - Documentation Standards — added: schema descriptions written first, go generate ./...
    for doc generation

Added sections: none

Removed sections: none

Templates reviewed:
  - .specify/templates/plan-template.md       ✅ Constitution Check gate present; no update needed
  - .specify/templates/spec-template.md       ✅ Acceptance-scenario format compatible; no update needed
  - .specify/templates/tasks-template.md      ✅ Phase structure compatible; no update needed
  - .specify/templates/agent-file-template.md ✅ Generic; no update needed
  - No commands/*.md files found              ✅ (directory absent)

Follow-up TODOs: none
-->

# Terraform Provider for Twilio Constitution

## Core Principles

### I. Service Module Encapsulation

Each Twilio product MUST be implemented as a self-contained module under
`twilio/internal/services/<service>/`, and MUST expose its resources and data
sources exclusively through the `ServiceRegistration` interface
(`twilio/common/service_registration.go`). Modules MUST be registered centrally
in `twilio/services.go`.

Cross-service imports are forbidden. Shared logic belongs in `twilio/utils/` or
`twilio/common/` only.

**Rationale**: Encapsulation keeps each product independently navigable and
prevents tangled dependencies that make future Twilio API surface additions
difficult to scope.

### II. Consistent Resource Pattern (NON-NEGOTIABLE)

Every resource MUST conform to the established pattern without deviation:

- A factory function (e.g. `resourceXxx() *schema.Resource`) returns the schema.
- All CRUD callbacks accept `context.Context` and return `diag.Diagnostics`.
- The client MUST be accessed as `meta.(*common.TwilioClient).<Service>`.
- All Twilio SDK calls MUST use `WithContext()` variants so timeouts propagate.
- Default timeouts: Create/Update/Delete 10 min, Read 5 min.
- 404 responses MUST be detected with `utils.IsNotFoundError(err)` and handled
  by calling `d.SetId("")` and returning without error (soft delete).
- Attributes that cannot be updated in-place MUST use `ForceNew: true` so
  Terraform destroys and recreates the resource rather than producing an
  unactionable diff.
- Attributes that hold secrets, tokens, or credentials MUST be marked
  `Sensitive: true` so Terraform redacts their values from plan and state output.
- Optional string fields MUST use `utils.OptionalString(d, "field")`.
- JSON schema fields MUST use `structure.SuppressJsonDiff` for diff suppression.
- Timestamps MUST be stored in `time.RFC3339` format.

**Rationale**: Consistent patterns reduce cognitive load, make code review
predictable, and ensure all resources behave identically under Terraform's
lifecycle model. `ForceNew` and `Sensitive` are non-negotiable safety properties
that protect users from data loss and credential exposure.

### III. SDK-Mediated API Access

All communication with the Twilio API MUST go through
`github.com/RJPearson94/twilio-sdk-go`. Direct HTTP clients, `net/http` calls
to Twilio endpoints, or alternative SDKs are forbidden.

SID validation MUST use the validators in `twilio/utils/validation.go`. New
Twilio resource types MUST add a dedicated validator there following the existing
pattern (`validation.StringMatch(regexp.MustCompile(...), "")`).

**Rationale**: SDK mediation centralises retry logic, credential handling, edge/
region routing, and backoff — duplicating this in individual resources creates
divergence and maintenance burden.

### IV. Acceptance-Test Coverage (NON-NEGOTIABLE)

New resources and data sources MUST ship with acceptance tests that cover all of
the following scenarios:

**Required scenarios for every resource:**

- **Basic lifecycle** — Create, Read, and Delete with minimum required config.
- **Update** — Combine into one test function as a multi-step sequence: step 1
  applies the basic config, step 2 applies an expanded config; verifies in-place
  update behaviour without reprovisioning.
- **Import** — `ImportState: true` with `ImportStateVerify: true` verifies the
  import path produces state identical to a fresh apply.
- **Disappears** — Deletes the resource via the API inside a `CheckState`
  callback and sets `ExpectNonEmptyPlan: true`; verifies the provider detects
  external deletion and handles it gracefully rather than erroring.

**Test authoring rules:**

- Tests MUST use `acceptance.PreCheck(t)` and `acceptance.TestAccProviderFactories`.
- Destroy checks MUST verify the resource no longer exists via the API.
- Resource names in test configs MUST be randomised (e.g. using
  `sdkacctest.RandomWithPrefix`) to prevent conflicts between parallel runs.
- Tests provision real Twilio resources and incur cost — they MUST clean up
  after themselves unconditionally; rely on `CheckDestroy` and deferred cleanup.

**Regression tests:**

Bugs fixed by code changes MUST be accompanied by a regression test. Use the
two-commit workflow: first commit adds the failing test (proving the bug exists),
second commit adds the fix (test now passes). Name the test
`TestAccTwilioXxx_regressionGH<number>` and link to the issue in a comment.

Unit tests alone are insufficient; the acceptance test suite is the authoritative
correctness gate.

**Rationale**: The disappears pattern and multi-step update tests are
industry-standard provider quality signals. Randomised names prevent false
failures from parallel execution. The regression two-commit workflow gives
reviewers an independent verification of the bug.

### V. Simplicity and Minimal Surface

Implement exactly what the Twilio API surface requires — no more. Specifically:

- Do not add provider-side defaults, computed logic, or derived fields unless the
  API itself provides them.
- Do not create helpers or abstractions for one-off operations.
- Do not add optional schema attributes speculatively for future API changes.
- Complexity introduced beyond the raw API contract MUST be justified in the
  plan's Complexity Tracking table.

**Rationale**: Over-abstraction in providers creates drift from the API contract
and makes future API changes harder to accommodate cleanly.

## Documentation Standards

Schema field descriptions MUST be written before documentation templates. Every
`schema.Schema` entry for a user-facing attribute MUST include a `Description`
string that states: purpose, constraints, defaults (if any), and whether the
field is computed. These descriptions are the source of truth for generated docs
and MUST remain accurate as the API evolves.

Every resource and data source MUST have a documentation template in `docs/`
following Terraform Registry conventions. Working HCL examples MUST be provided
in `examples/` and kept in sync with the schema. Examples MUST be valid,
runnable HCL reflecting current argument and attribute names.

After any schema or template edit, regenerate documentation:

```bash
go generate ./...
```

Format generated docs before merging:

```bash
make terrafmt-docs
```

## Quality Gates

The following MUST pass before any code is merged:

- `make fmt` — Go code formatting (goimports + gofmt -s)
- `make terraform-fmt` — Terraform example formatting
- `make terrafmt-docs` — generated documentation formatting
- `make test` — All unit tests (30 s timeout, 4 parallel workers)
- Acceptance tests for the affected service MUST be run locally and pass before
  opening a PR for a new resource or data source.

**Pre-submission checklist for new resources:**

- [ ] All CRUD operations implemented
- [ ] `ForceNew: true` on all immutable attributes
- [ ] `Sensitive: true` on all credential/secret attributes
- [ ] Import implemented and tested
- [ ] Disappears test included
- [ ] Multi-step update test included
- [ ] Schema field `Description` set on every user-facing attribute
- [ ] Documentation template and examples present
- [ ] `go generate ./...` run and output committed

CI (`make build` + `make test`) runs on every push and PR via GitHub Actions.

## Governance

This constitution supersedes all informal conventions in this repository.
Any practice that conflicts with a stated principle MUST be updated to comply.

**Amendment procedure**: Open a PR that updates this file with an incremented
`CONSTITUTION_VERSION`, states the rationale for the change, and describes any
migration needed for existing resources. The PR description serves as the
amendment record.

**Versioning policy**:
- MAJOR — principle removed, redefined, or backward-incompatible governance change.
- MINOR — new principle or section added; material expansion of existing guidance.
- PATCH — clarification, wording, or non-semantic refinement.

**Compliance review**: Constitution Check gates in feature plans MUST be
evaluated against the current version of this file before Phase 0 research and
re-checked after Phase 1 design.

**Version**: 1.1.0 | **Ratified**: 2026-04-03 | **Last Amended**: 2026-04-03
