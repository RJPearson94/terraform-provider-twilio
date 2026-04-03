# Quickstart: Documentation Generation

**Feature**: 001-update-provider-docs
**Date**: 2026-04-03

---

## Prerequisites

- Go 1.24 installed
- `make tools` already run (installs linting tools)
- Twilio provider builds locally (`make build` succeeds)

---

## One-time setup: add tfplugindocs dependency

```bash
# Add the tools.go file (already done as part of this feature)
# Then update go.mod and vendor:
go get github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
go mod tidy
go mod vendor
```

---

## Regenerate all documentation

```bash
go generate ./...
```

This runs the `//go:generate` directive in `main.go`, which invokes `tfplugindocs` to
regenerate all files in `docs/resources/` and `docs/data-sources/` from the current
schema descriptions and examples.

---

## Format generated docs

```bash
make terrafmt-docs
```

Run this after `go generate ./...`. If it produces no changes, the generated output is
already clean. If it does produce changes, commit them alongside the generation output.

---

## Validate examples

```bash
make validate-all-examples
```

Checks that all `.tf` files in `examples/` are valid HCL. Fix any failures before opening
a PR.

---

## Verify docs are up to date (the CI check, run locally)

```bash
go generate ./...
git diff --exit-code docs/
```

If the `git diff` command exits non-zero, docs are stale — commit the updated `docs/`
files alongside the schema changes.

---

## Workflow for adding descriptions to a service

1. Check out branch `001-update-provider-docs` (or your own branch off it).
2. Open each resource and data source file for the target service (e.g. `twilio/internal/services/studio/`).
3. Add `Description:` to every `schema.Schema` entry following `contracts/description-conventions.md`.
4. Run `go generate ./...` to regenerate docs.
5. Run `make terrafmt-docs` to format.
6. Run `make validate-all-examples` to check examples.
7. Run `git diff --exit-code docs/` to confirm docs match schema.
8. Commit with message: `docs(studio): add schema descriptions and regenerate docs`.
9. Open PR. CI will verify all checks pass.

---

## Troubleshooting

**`go generate ./...` fails with "package not found"**
→ Run `go mod tidy && go mod vendor` first.

**`make terrafmt-docs` produces changes after generation**
→ Commit the formatting changes too — both generation and formatting outputs must be committed.

**`git diff docs/` shows changes after generation**
→ Commit the updated `docs/` files. CI enforces that docs are always up to date.

**`make validate-all-examples` fails**
→ Open the failing `.tf` file and check for attribute names that no longer exist in the
schema. Fix the example to use current attribute names.
