---
name: provider-docs
description: Generate, update, and validate Terraform provider documentation from Go schema code. Use when adding or changing resources, data sources, or schema fields to keep docs/ aligned with the provider implementation.
---

# Terraform Provider Documentation

Generate and maintain Terraform Registry documentation directly from provider schema code. No external tooling (tfplugindocs) is needed — Claude reads the schema and writes the markdown.

## When to Use

- After adding a new resource or data source
- After changing schema fields (adding, removing, renaming, changing Required/Optional/Computed)
- After updating Description strings on schema fields
- To audit docs against current schema for drift
- To create docs for a resource or data source that doesn't have one yet

## Document Format

Every resource doc lives at `docs/resources/<name>.md` and every data source doc at `docs/data-sources/<name>.md`. The format follows the Terraform Registry spec:

```markdown
---
page_title: "twilio_<name> <Resource|Data Source> - twilio"
subcategory: "<Service>"
description: |-
  
---

# twilio_<name> <Resource|Data Source>

<One-line description with link to Twilio API docs>

<Optional second line linking to product page>

## Example Usage

\```hcl
<Minimal, runnable example>
\```

## Schema

### Required

- `field` (Type) Description from schema

### Optional

- `field` (Type) Description from schema

### Read-Only

- `field` (Type) Description from schema

<Nested schema blocks as needed>

## Import  (resources only)

<Import format and example>
```

## Workflow

### 1. Identify affected docs

Map the code change to doc targets:
- New resource -> new `docs/resources/<name>.md`
- New data source -> new `docs/data-sources/<name>.md`
- Schema field change -> update the Schema section of the existing doc

### 2. Read the schema source

Read the Go source file(s) to extract the full schema:
- `twilio/internal/services/<service>/resource_<name>.go`
- `twilio/internal/services/<service>/data_source_<name>.go`

For each `schema.Schema` entry, extract:
- Field name (map key)
- Type (`schema.TypeString`, `schema.TypeInt`, `schema.TypeBool`, `schema.TypeList`, `schema.TypeSet`, `schema.TypeMap`)
- Required / Optional / Computed flags
- Description string
- Default value
- ForceNew flag
- Sensitive flag
- Nested `Elem` schema (for blocks)
- ValidateFunc constraints (enum values, ranges)
- ConflictsWith / ExactlyOneOf / AtLeastOneOf / RequiredWith

### 3. Read the existing doc (if any)

If a doc already exists, preserve:
- The **frontmatter** (page_title, subcategory)
- The **description paragraphs** (between the title and Example Usage)
- The **Example Usage** section (all HCL examples)
- The **Import** section (for resources)

Only regenerate the **Schema** section from the current source code.

### 4. Write the Schema section

Follow this format exactly — it matches what Terraform Registry renders:

**Type mapping:**

| Go Type | Doc Type |
|---------|----------|
| `schema.TypeString` | `String` |
| `schema.TypeInt` | `Number` |
| `schema.TypeFloat` | `Number` |
| `schema.TypeBool` | `Boolean` |
| `schema.TypeList` with `Elem: &schema.Schema{}` | `List of <ElementType>` |
| `schema.TypeSet` with `Elem: &schema.Schema{}` | `Set of <ElementType>` |
| `schema.TypeMap` | `Map of String` |
| `schema.TypeList` with `Elem: &schema.Resource{}` | `Block List` |
| `schema.TypeSet` with `Elem: &schema.Resource{}` | `Block Set` |

**Block attributes** get an extra suffix: `(see [below for nested schema](#nestedblock--<name>))` and MaxItems noted if 1: `Block List, Max: 1`.

**Grouping order:**
1. `### Required` — fields with `Required: true`
2. `### Optional` — fields with `Optional: true` (include `Timeouts` block here if present)
3. `### Read-Only` — fields with `Computed: true` and NOT Required/Optional (always include `id`)

**Each field line:**
```
- `<name>` (<Type>) <Description>
```

**Nested blocks** appear after the main schema:
```markdown
<a id="nestedblock--<name>"></a>
### Nested Schema for `<name>`

Required/Optional/Read-Only sections as above
```

### 5. Validate

After writing docs:
- Confirm every schema field appears in the doc
- Confirm Required/Optional/Computed grouping matches the schema
- Confirm nested blocks have anchors and are documented
- Confirm examples use current field names

## Subcategory Mapping

| Service directory | Subcategory |
|-------------------|-------------|
| account | Account |
| chat | Programmable Chat |
| conversations | Conversations |
| credentials | Credentials |
| flex | Flex |
| iam | IAM |
| messaging | Messaging |
| phone_number | Phone Numbers |
| proxy | Proxy |
| serverless | Serverless |
| sip | SIP |
| sip_trunking | SIP Trunking |
| studio | Studio |
| sync | Sync |
| taskrouter | TaskRouter |
| twiml | TwiML |
| verify | Verify |
| video | Video |
| voice | Voice |

## Quality Rules

- Every schema field MUST have a Description in the Go source — add one before generating docs
- Never describe fields that don't exist in the schema
- Never duplicate the Schema section content in manual prose above it
- Keep examples minimal, realistic, and using current field names
- Include import instructions for every resource that supports import
- Follow `specs/001-update-provider-docs/contracts/description-conventions.md` for Description wording

## Registry Publication Rules

- Use semantic version tags prefixed with `v` (for example `v1.2.3`)
- Create release tags from the default branch
- Keep `terraform-registry-manifest.json` in the repository root
- Docs are versioned in Registry and switchable with the version selector

## Load References On Demand

- Read `references/hashicorp-provider-docs.md` for Registry publication rules and official links
- Read `specs/001-update-provider-docs/contracts/description-conventions.md` for Description conventions
