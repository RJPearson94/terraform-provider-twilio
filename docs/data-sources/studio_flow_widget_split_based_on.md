---
page_title: "twilio_studio_flow_widget_split_based_on Data Source - twilio"
subcategory: "Studio"
description: |-
  
---

# twilio_studio_flow_widget_split_based_on Data Source

Use this data source to generate the JSON for the Studio Flow split based on widget. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition. See the [docs](https://www.twilio.com/docs/studio/widget-library/split-based-on) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

```hcl
data "twilio_studio_flow_widget_split_based_on" "split_based_on" {
  name = "SplitBasedOn"

  transitions {
    matches {
      next = "test"
      conditions {
        arguments     = ["{{contact.channel.address}}"]
        friendly_name = "If value equal_to test"
        type          = "equal_to"
        value         = "test"
      }
    }
  }

  input = "{{contact.channel.address}}"
}
```

## Schema

### Required

- `input` (String) The value or Liquid template expression to evaluate against the match conditions (e.g. `{{widgets.my_widget.parsed.status}}`)
- `name` (String) The unique name of this widget within the flow, used to reference it in transitions

### Optional

- `offset` (Block List, Max: 1) The position of this widget in the Studio visual editor (see [below for nested schema](#nestedblock--offset))
- `transitions` (Block List, Max: 1) The conditional routing rules and fallback transition for this widget (see [below for nested schema](#nestedblock--transitions))

### Read-Only

- `id` (String) The ID of this resource.
- `json` (String) A JSON string representation of the widget state, for use as an entry in the `states` list of a `twilio_studio_flow_definition` data source

<a id="nestedblock--offset"></a>
### Nested Schema for `offset`

Optional:

- `x` (Number) The x-axis position. Defaults to 0
- `y` (Number) The y-axis position. Defaults to 0


<a id="nestedblock--transitions"></a>
### Nested Schema for `transitions`

Optional:

- `matches` (Block List) A list of match rules, each routing to a different widget based on conditions (see [below for nested schema](#nestedblock--transitions--matches))
- `no_match` (String) The name of the next widget when none of the match conditions are satisfied

<a id="nestedblock--transitions--matches"></a>
### Nested Schema for `transitions.matches`

Required:

- `conditions` (Block List, Min: 1) The list of conditions that must all be true for this match to fire (see [below for nested schema](#nestedblock--transitions--matches--conditions))
- `next` (String) The name of the next widget when all conditions in this match are satisfied

<a id="nestedblock--transitions--matches--conditions"></a>
### Nested Schema for `transitions.matches.conditions`

Required:

- `arguments` (List of String) The arguments to pass to the condition operator
- `friendly_name` (String) A human-readable label for this condition
- `type` (String) The comparison operator. Valid values: `equal_to`, `not_equal_to`, `matches_any_of`, `does_not_match_any_of`, `is_blank`, `is_not_blank`, `regex`, `contains`, `does_not_contain`, `starts_with`, `does_not_start_with`, `less_than`, `greater_than`, `is_before_time`, `is_after_time`, `is_before_date`, `is_after_date`
- `value` (String) The value to compare the input against
