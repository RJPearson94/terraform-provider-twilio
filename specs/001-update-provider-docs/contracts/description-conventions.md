# Contract: Schema Description Conventions

**Feature**: 001-update-provider-docs
**Date**: 2026-04-03

This document defines the writing conventions for `Description` strings on `schema.Schema`
entries. All contributors MUST follow these conventions so descriptions are consistent
across the provider.

---

## General Rules

- Descriptions are plain English sentences. Capitalise the first word. No trailing period required.
- Maximum 200 characters per description (soft limit — prefer conciseness).
- Do not include implementation details: no API endpoint paths, no SDK method names, no Go types.
- Do not restate the field name: `"friendly_name"` → not "The friendly_name of the resource" → use "A human-readable label for the resource".

---

## Field Category Conventions

### Required configurable fields

State what the value controls. Example:
```go
"friendly_name": {
    Description: "A human-readable label for the flow, unique within the account",
},
```

### Optional configurable fields

State purpose and any default. Example:
```go
"commit_message": {
    Description: "A description of the changes made in this flow revision",
},
"validity_period": {
    Description: "How long, in seconds, the message is valid. Defaults to 36000 (10 hours)",
},
```

### Computed-only fields (Computed: true, not Optional)

MUST state the value is set by Twilio. Pattern: "The \<description\>. Set by Twilio."

```go
"sid": {
    Description: "The unique SID assigned to this flow by Twilio",
},
"account_sid": {
    Description: "The SID of the account that owns this resource",
},
"date_created": {
    Description: "The date and time the resource was created, in RFC 3339 format",
},
"date_updated": {
    Description: "The date and time the resource was last updated, in RFC 3339 format",
},
"url": {
    Description: "The absolute URL of the resource",
},
```

### Enum-constrained fields

MUST list all valid values in backticks. Example:
```go
"status": {
    Description: "The status of the flow. Valid values: `draft`, `published`",
},
"inbound_method": {
    Description: "The HTTP method for inbound requests. Valid values: `GET`, `POST`. Defaults to `POST`",
},
"prioritize_queue_order": {
    Description: "How TaskRouter should prioritise reservations. Valid values: `FIFO`, `LIFO`",
},
```

### Parent SID fields (ForceNew: true)

MUST mention that changing the value forces resource recreation. Example:
```go
"service_sid": {
    Description: "The SID of the Serverless service this environment belongs to. Changing this forces a new resource to be created",
},
"workspace_sid": {
    Description: "The SID of the TaskRouter workspace this resource belongs to. Changing this forces a new resource to be created",
},
```

### Sensitive fields (Sensitive: true)

Describe the value normally. Do not add "secret" or "sensitive" to the description —
Terraform already redacts these values in output.

```go
"aws_secret_access_key": {
    Description: "The AWS secret access key credential for the binding",
},
```

### JSON fields (using SuppressJsonDiff)

State that the value is a JSON string. Example:
```go
"definition": {
    Description: "A JSON string defining the flow. See the Twilio Studio Flow documentation for the schema",
},
```

---

## Common Shared Fields

Use exactly these descriptions for the common fields that appear across many resources.
Copy-paste to ensure consistency:

| Field name | Description string |
|---|---|
| `sid` | `"The unique SID assigned to this [resource] by Twilio"` |
| `account_sid` | `"The SID of the account that owns this resource"` |
| `date_created` | `"The date and time the resource was created, in RFC 3339 format"` |
| `date_updated` | `"The date and time the resource was last updated, in RFC 3339 format"` |
| `url` | `"The absolute URL of the resource"` |
| `friendly_name` | Contextual — describe what the name identifies, do not use "A friendly name" |

---

## Review checklist for description PRs

- [ ] No field has an empty Description
- [ ] Computed fields say the value is set by Twilio
- [ ] Enum fields list all valid values
- [ ] ForceNew fields mention resource recreation
- [ ] No field restates its own name as the description
- [ ] JSON fields say the value is a JSON string
