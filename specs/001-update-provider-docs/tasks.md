---
description: "Task list for feature 001-update-provider-docs"
---

# Tasks: Update Provider Documentation

**Input**: Design documents from `/specs/001-update-provider-docs/`
**Prerequisites**: plan.md ✅, spec.md ✅, research.md ✅, data-model.md ✅, contracts/ ✅

**Tests**: No acceptance tests required (adding Description strings does not change runtime behaviour per constitution Principle IV).

**Organization**: Phases 1–2 are foundational setup. Phase 3 (US1) is the bulk of the work — descriptions per service. Phase 4 (US2) adds CI enforcement. Phase 5 (US3) validates examples.

**⚠️ Implementation note**: US2 tooling (Phases 1–2) MUST be complete before US1 description work can be fully verified. However, description additions can begin immediately in parallel — just run `go generate ./...` locally to check each service's output before committing.

## Format: `[ID] [P?] [Story?] Description`

- **[P]**: Can run in parallel (different files, no shared dependencies)
- **[US]**: Which user story this task belongs to
- All file paths are relative to the repo root

---

## Phase 1: Setup

**Purpose**: Wire tfplugindocs tooling into the project.

- [x] T001 Create `tools.go` at the repo root with a `//go:build tools` tag and a blank import of `github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs` to pin it in go.mod
- [x] T002 Add `//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name twilio` to the top of `main.go`

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Vendor the new dependency and establish a clean generation baseline.

**⚠️ CRITICAL**: US1 verification depends on generation working. US2 CI enforcement depends on this phase being complete. US3 example validation is independent.

- [x] T003 Run `go get github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs && go mod tidy && go mod vendor` and commit the updated `go.mod`, `go.sum`, and `vendor/` directory
- [x] T004 Run `go generate ./...` to produce the initial regenerated docs baseline; commit all resulting changes to `docs/resources/` and `docs/data-sources/`
- [x] T005 Run `make terrafmt-docs` immediately after T004 and confirm it produces no diff; if it does, commit those formatting changes too

**Checkpoint**: `go generate ./...` runs cleanly, `docs/` is committed, `make terrafmt-docs` is a no-op. US1 description work can now be verified locally.

---

## Phase 3: User Story 1 — Attribute Descriptions (Priority: P1) 🎯 MVP

**Goal**: Every schema attribute across all resources and data sources has a non-empty Description, visible in Terraform Registry docs.

**Independent Test**: Run `go generate ./...` after completing all tasks in this phase and verify the output of `docs/resources/*.md` shows populated Argument Reference and Attribute Reference tables. Run `git diff --exit-code docs/` to confirm docs match schema.

**Convention reference**: Follow `specs/001-update-provider-docs/contracts/description-conventions.md` for all descriptions.

**Pattern for each service task below**:
1. Open each `resource_*.go` and `data_source_*.go` file in the service directory
2. Add `Description: "..."` to every `schema.Schema{}` entry without one
3. Run `go generate ./...` and inspect the generated markdown for that service
4. Run `make terrafmt-docs` to confirm formatting is clean
5. Commit: `docs(service): add schema descriptions and regenerate docs`

- [x] T006 [P] [US1] Add `Description` to all `schema.Schema` entries in `twilio/internal/services/studio/` (resource_studio_flow.go and all data sources)
- [x] T007 [P] [US1] Add `Description` to all `schema.Schema` entries in `twilio/internal/services/messaging/` (messaging service, phone number, short code, alpha sender)
- [x] T008 [P] [US1] Add `Description` to all `schema.Schema` entries in `twilio/internal/services/taskrouter/` (workspace, task queue, worker, workflow, activity)
- [x] T009 [P] [US1] Add `Description` to all `schema.Schema` entries in `twilio/internal/services/conversations/` (conversation, participant, message, service, webhook)
- [x] T010 [P] [US1] Add `Description` to all `schema.Schema` entries in `twilio/internal/services/verify/` (service, rate limit, rate limit bucket, messaging configuration)
- [x] T011 [P] [US1] Add `Description` to all `schema.Schema` entries in `twilio/internal/services/serverless/` (service, environment, function, asset, build, deployment, variable)
- [x] T012 [P] [US1] Add `Description` to all `schema.Schema` entries in `twilio/internal/services/proxy/` (service, phone number, short code, session, participant)
- [x] T013 [P] [US1] Add `Description` to all `schema.Schema` entries in `twilio/internal/services/sync/` (service, document, list, list item, map, map item, stream)
- [x] T014 [P] [US1] Add `Description` to all `schema.Schema` entries in `twilio/internal/services/account/` (details, sub-account)
- [x] T015 [P] [US1] Add `Description` to all `schema.Schema` entries in `twilio/internal/services/chat/` (service, channel, member, message, role, user, webhook)
- [x] T016 [P] [US1] Add `Description` to all `schema.Schema` entries in `twilio/internal/services/credentials/` (aws, public_key)
- [x] T017 [P] [US1] Add `Description` to all `schema.Schema` entries in `twilio/internal/services/flex/` (flow, plugin, plugin configuration, plugin release)
- [x] T018 [P] [US1] Add `Description` to all `schema.Schema` entries in `twilio/internal/services/iam/` (api_key)
- [x] T019 [P] [US1] Add `Description` to all `schema.Schema` entries in `twilio/internal/services/phone_number/` (available numbers, incoming phone numbers)
- [x] T020 [P] [US1] Add `Description` to all `schema.Schema` entries in `twilio/internal/services/sip/` (domain, credential list, ip access control list, ip address)
- [x] T021 [P] [US1] Add `Description` to all `schema.Schema` entries in `twilio/internal/services/sip_trunking/` (trunk, credential list, ip access control list, origination url, phone number)
- [x] T022 [P] [US1] Add `Description` to all `schema.Schema` entries in `twilio/internal/services/twiml/` (app, apps data source)
- [x] T023 [P] [US1] Add `Description` to all `schema.Schema` entries in `twilio/internal/services/video/` (composition hook, composition settings, room)
- [x] T024 [US1] Run `go generate ./...` across the full provider after all service description tasks are complete; run `git diff --exit-code docs/` to confirm generated docs match the updated schema; run `make terrafmt-docs` to confirm no formatting diff

**Checkpoint**: All 18 services have descriptions. `go generate ./...` produces a clean diff. `make terrafmt-docs` is a no-op. Manual review of `docs/resources/studio_flow.md`, `docs/resources/messaging_service.md`, and `docs/resources/taskrouter_workspace.md` confirms populated Argument/Attribute Reference tables.

---

## Phase 4: User Story 2 — Single-Command Documentation Generation (Priority: P2)

**Goal**: CI enforces that docs are always regenerated after schema changes; a PR with stale docs is blocked automatically.

**Independent Test**: Push a branch with a schema description change but without running `go generate ./...`. Verify the CI pipeline fails with a message directing the contributor to run `go generate ./...`.

- [x] T025 [US2] Add a doc-currency CI step to `.github/workflows/terraform_provider.yml` after the `Build` step: run `go generate ./...` then `git diff --exit-code docs/ || (echo "Run 'go generate ./...' and commit the docs/ changes" && exit 1)`
- [x] T026 [P] [US2] Extend the `paths` trigger in `.github/workflows/terraform_provider.yml` to include `docs/**` and `examples/**` so doc-only and example-only PRs also trigger the pipeline
- [ ] T027 [US2] Verify the updated CI workflow passes on the current branch by pushing and confirming all steps (including the new doc-currency step) succeed

**Checkpoint**: A PR with an undocumented schema change is blocked by CI. The new step message clearly instructs the contributor on how to fix it.

---

## Phase 5: User Story 3 — Valid, Runnable Examples (Priority: P3)

**Goal**: All files in `examples/` are valid HCL that matches the current schema.

**Independent Test**: Run `make validate-all-examples` and confirm exit code 0.

- [x] T028 [US3] Run `make validate-all-examples` against the current `examples/` directory and record any failures
- [x] T029 [US3] For each failing example identified in T028, update the `.tf` file to use current attribute names and include all required arguments; re-run `make validate-all-examples` until exit code is 0

**Checkpoint**: `make validate-all-examples` exits with code 0 across all examples.

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: Final validation and CI completeness.

- [x] T030 [P] Add `make validate-all-examples` as a CI step in `.github/workflows/terraform_provider.yml` (after the doc-currency step added in T025) to enforce example validity on every PR
- [x] T031 Manually review the generated Registry docs for three representative resources — `docs/resources/studio_flow.md`, `docs/resources/messaging_service.md`, `docs/resources/taskrouter_workspace.md` — and confirm every Argument Reference and Attribute Reference entry has a meaningful, non-empty description

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies — start immediately
- **Foundational (Phase 2)**: Depends on Phase 1 (tools.go + go:generate must exist before vendoring)
- **US1 — Descriptions (Phase 3)**: Can begin after Phase 1; fully verifiable after Phase 2
- **US2 — CI Enforcement (Phase 4)**: Depends on Phase 2 (generation must work before CI can enforce it); run after most US1 work is merged
- **US3 — Examples (Phase 5)**: Independent of Phases 3 and 4; can run at any time after Phase 2
- **Polish (Phase 6)**: Depends on Phases 3, 4, and 5 being complete

### User Story Dependencies

- **US1 (P1)**: Can start after Phase 1 (locally verifiable); fully CI-enforced after US2
- **US2 (P2)**: Depends on Phase 2 completion; independent of US1 description content
- **US3 (P3)**: Independent of US1 and US2 (different files — `examples/` only)

### Within Phase 3 (US1)

- T006–T023 are all [P] — each touches a different service directory, no shared files
- T024 (final generation + diff check) depends on T006–T023 all being complete

### Parallel Opportunities

```bash
# Phase 3: all 18 service description tasks can run in parallel across team members
# Each task is a self-contained directory:
Task: "Add descriptions to twilio/internal/services/studio/"      # T006
Task: "Add descriptions to twilio/internal/services/messaging/"   # T007
Task: "Add descriptions to twilio/internal/services/taskrouter/"  # T008
# ... etc
```

---

## Implementation Strategy

### MVP First (US1 — Studio service only)

1. Complete Phase 1: Setup (T001–T002)
2. Complete Phase 2: Foundational (T003–T005)
3. Complete T006 (studio service descriptions only)
4. Run `go generate ./...` and verify studio docs are populated
5. **STOP and VALIDATE**: Confirm `docs/resources/studio_flow.md` shows descriptions in Registry format
6. Merge — demonstrates the full end-to-end flow with one service

### Incremental Delivery (recommended)

1. Complete Phases 1–2 (tooling)
2. Add descriptions one service per PR (T006 → T007 → … → T023)
3. Each PR triggers `go generate ./...` + doc diff check
4. Add CI enforcement (Phase 4) after the first few service PRs are merged
5. Run example validation (Phase 5) and Polish (Phase 6) last

### Parallel Team Strategy

With multiple contributors:
- One contributor: Phases 1–2 (tooling — serialize to avoid go.mod conflicts)
- Multiple contributors in parallel: T006–T023 (one service each, different directories)
- One contributor: T025–T027 (CI enforcement, once tooling is merged)

---

## Notes

- [P] tasks use different service directories — safe to run in parallel
- Commit message convention: `docs(<service>): add schema descriptions and regenerate docs`
- Run `go generate ./...` + `make terrafmt-docs` + `git diff docs/` after every service batch before committing
- The contracts/description-conventions.md file is the source of truth for description wording
- Avoid: writing descriptions that restate the field name, leaving enum fields without value lists, forgetting the "Changing this forces a new resource" note on ForceNew fields
