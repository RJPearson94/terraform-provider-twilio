# Research: Update Provider Documentation

**Feature**: 001-update-provider-docs
**Date**: 2026-04-03

---

## Decision 1: Documentation generation tool

**Decision**: Use `tfplugindocs` (`github.com/hashicorp/terraform-plugin-docs`).

**Rationale**: tfplugindocs is the HashiCorp-maintained standard generator for both
terraform-plugin-sdk/v2 and terraform-plugin-framework providers. It reads `Description`
fields from `schema.Schema` structs, combines them with HCL examples from `examples/`,
and produces Terraform Registry-compatible markdown in `docs/`. The existing `docs/`
directory already uses the layout tfplugindocs expects (`docs/resources/`,
`docs/data-sources/`), so the tool can overwrite files in-place.

**Alternatives considered**:
- Manual docs maintenance — already the current approach; known to drift from schema.
- Custom doc generator — unnecessary complexity given tfplugindocs covers the use case.

---

## Decision 2: go:generate directive placement

**Decision**: Add a `//go:generate` comment to `main.go` (the standard location for
Terraform providers) and a `tools.go` file to pin `tfplugindocs` in go.mod.

**Rationale**: A `tools.go` file with a blank import under a `//go:build tools` build tag
is the idiomatic Go pattern for pinning tool dependencies in go.mod without including
them in the binary. This allows `go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs`
to resolve via the module graph. The existing `make generate` target already runs
`go generate ./...`, so no Makefile changes are needed.

**Directive (in main.go)**:
```go
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name twilio
```

**tools.go (new file at repo root)**:
```go
//go:build tools
// +build tools

package main

import (
    _ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
)
```

**Alternatives considered**:
- Adding tfplugindocs to `make tools` as a `go install` — fragile because version is not
  pinned in go.mod, and different developers may get different versions.
- Separate script — adds friction compared to the standard `go generate ./...` flow.

---

## Decision 3: Vendor compatibility

**Decision**: Run `go mod tidy && go mod vendor` after adding tfplugindocs to tools.go,
then commit the updated `vendor/` directory.

**Rationale**: The project uses `go mod vendor` (evidenced by the `vendor/` directory
and `make download` running `go mod vendor`). All dependencies must be vendored for CI
and offline builds to work. tfplugindocs and its transitive dependencies must be added
to `vendor/` via `go mod vendor`.

**Alternatives considered**:
- `-mod=mod` flag to bypass vendor for generation — inconsistent with project convention.

---

## Decision 4: CI doc-currency enforcement

**Decision**: Add a new CI step after `make generate` that runs `git diff --exit-code docs/`
to fail the build if generated docs differ from committed docs.

**Rationale**: Without enforcement, contributors will forget to regenerate. The check
is cheap (no network, no compilation) and gives a clear error message pointing to the
fix (`go generate ./...`). The existing GitHub Actions workflow at
`.github/workflows/terraform_provider.yml` is the right place to add it.

**New CI step**:
```yaml
- name: Check documentation is up to date
  run: |
    go generate ./...
    git diff --exit-code docs/ || (echo "Run 'go generate ./...' and commit the docs/ changes" && exit 1)
```

The `paths` trigger in the workflow must be extended to include `docs/**` so that
doc-only PRs also run the check.

**Alternatives considered**:
- Terraform Cloud Workspace automation — overkill for a file diff check.
- `make` target that wraps the check — acceptable but the raw git diff is self-evident.

---

## Decision 5: Description writing order and convention

**Decision**: Write descriptions service-by-service in priority order:
studio → messaging → taskrouter → conversations → verify → serverless → proxy → sync
→ remaining services. Each PR covers one service.

**Conventions established** (see contracts/description-conventions.md for full detail):
- Computed fields: "Set by Twilio. Cannot be configured."
- Enum fields: list values in backticks, e.g. `` `draft` or `published` ``.
- Parent SID fields: "The SID of the parent [Resource]. Changing this forces a new resource."
- Timestamps: "The date and time the [resource] was [created/updated] in RFC 3339 format."
- URL fields: "The absolute URL of the [resource]."

**Rationale**: Service-by-service PRs keep review size manageable. Conventions ensure
consistency without reviewing each description independently.

---

## Decision 6: Example validation

**Decision**: Use the existing `make validate-all-examples` target as-is. Fix any
examples that currently fail before closing the feature.

**Rationale**: The Makefile already has this target. Running it as a pre-merge check
(added to CI alongside the doc-currency step) catches broken examples without additional
tooling.

**Finding from current state**: `make validate-all-examples` likely passes today — the
existing examples are mostly correct. Any failures found during this feature will be
fixed as part of the same PR as the affected resource's description work.
