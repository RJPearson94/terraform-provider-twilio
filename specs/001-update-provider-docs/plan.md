# Implementation Plan: Update Provider Documentation

**Branch**: `001-update-provider-docs` | **Date**: 2026-04-03 | **Spec**: [spec.md](spec.md)
**Input**: Feature specification from `/specs/001-update-provider-docs/spec.md`

## Summary

Add `Description` strings to every schema attribute across all ~80 resources and ~138
data sources, wire up `tfplugindocs` for automated doc generation via `go generate ./...`,
and add a CI step that enforces docs are always regenerated after schema changes. A
description-writing convention document ensures consistency across the provider.

## Technical Context

**Language/Version**: Go 1.24
**Primary Dependencies**: terraform-plugin-sdk/v2 (existing), tfplugindocs v0.x (new — added via tools.go)
**Storage**: Markdown files in `docs/` (generated), Go source files in `twilio/internal/services/`
**Testing**: `go generate ./...` + `git diff --exit-code docs/` (currency); `make validate-all-examples` (examples); `make terrafmt-docs` (formatting)
**Target Platform**: Developer toolchain (generation); Terraform Registry (output)
**Project Type**: Terraform provider plugin
**Performance Goals**: N/A — generation runs once per PR, not at runtime
**Constraints**: Must not change provider runtime behaviour; no acceptance test changes required
**Scale/Scope**: ~153 Go files (80 resources + 73 data sources across 19 services); ~82 existing resource docs + ~138 data source docs to regenerate

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

| Principle | Status | Notes |
|---|---|---|
| I. Service Module Encapsulation | ✅ Pass | Changes stay within existing service directories; no cross-service imports introduced |
| II. Consistent Resource Pattern | ✅ Pass | Only `Description` fields added to existing `schema.Schema` entries; no CRUD changes |
| III. SDK-Mediated API Access | ✅ Pass | No API calls; tooling change only |
| IV. Acceptance Test Coverage | ✅ Pass | Constitution explicitly excludes Description additions from acceptance test requirement |
| V. Simplicity and Minimal Surface | ✅ Pass | No new abstractions; only data strings added to existing schema structs |
| Documentation Standards | ✅ Pass | This feature IS the documentation standards compliance work; adds go:generate and enforces descriptions |
| Quality Gates | ✅ Pass | Adds CI enforcement on top of existing gates |

No violations. No Complexity Tracking required.

**Post-design re-check**: All principles still pass. tools.go is the only new file added to the repo root; it uses a standard Go build tag and introduces no runtime complexity.

## Project Structure

### Documentation (this feature)

```text
specs/001-update-provider-docs/
├── plan.md                           # This file
├── research.md                       # Tool selection, CI approach, convention decisions
├── data-model.md                     # Schema description entity, doc file entity, CI entity
├── quickstart.md                     # How to run generation locally
├── contracts/
│   └── description-conventions.md   # Canonical writing conventions for descriptions
└── tasks.md                          # Created by /speckit-tasks (next step)
```

### Source Code (repository root)

```text
tools.go                              # NEW — blank import of tfplugindocs (build tag: tools)
main.go                               # MODIFIED — add //go:generate directive
go.mod                                # MODIFIED — add tfplugindocs dependency
go.sum                                # MODIFIED — updated checksums
vendor/                               # MODIFIED — vendored tfplugindocs + transitive deps
.github/workflows/terraform_provider.yml  # MODIFIED — add doc-currency CI step + paths trigger

twilio/internal/services/studio/
├── resource_studio_flow.go           # MODIFIED — add Description to all schema fields
├── data_source_studio_flow.go        # MODIFIED — add Description to all schema fields
├── resource_studio_flow_widget_*.go  # MODIFIED — add Description where present
└── ...

twilio/internal/services/messaging/  # MODIFIED — same pattern
twilio/internal/services/taskrouter/ # MODIFIED — same pattern
twilio/internal/services/conversations/ # MODIFIED — same pattern
twilio/internal/services/verify/     # MODIFIED — same pattern
twilio/internal/services/serverless/ # MODIFIED — same pattern
twilio/internal/services/proxy/      # MODIFIED — same pattern
twilio/internal/services/sync/       # MODIFIED — same pattern
twilio/internal/services/*/          # MODIFIED — all remaining services

docs/resources/*.md                  # REGENERATED — by go generate ./...
docs/data-sources/*.md               # REGENERATED — by go generate ./...
```

**Structure Decision**: Single project, modifying existing files in-place. No new
directories in `twilio/`. New file `tools.go` at repo root for tool dependency pinning.

## Complexity Tracking

> No violations — not required.
