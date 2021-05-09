---
page_title: "Twilio Studio Flow Widget - Split Based On"
subcategory: "Studio"
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

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the split based on widget
- `input` - (Mandatory) The value or expression that is being evaluated to split on
- `transitions` - (Optional) A `transitions` block as documented below
- `offset` - (Optional) A `offset` block as documented below

---

A `transitions` block supports the following:

- `no_match` - (Optional) The widget to transition to when no match is found
- `matches` - (Optional) A list of `match` blocks as documented below

---

A `match` block supports the following:

- `next` - (Mandatory) The widget to transition to when the value or expression matches the conditions
- `conditions` - (Mandatory) A list of `condition` blocks as documented below

~> At least 1 condition block must be supplied

---

A `condition` block supports the following:

- `arguments` - (Mandatory) A list of arguments to evaluate
- `friendly_name` - (Mandatory) The name of the condition
- `type` - (Mandatory) The type/ operator to use when comparing the arguments and value. Valid values include: `equal_to`,`not_equal_to`,`matches_any_of`,`does_not_match_any_of`,`is_blank`,`is_not_blank`,`regex`,`contains`,`does_not_contain`,`starts_with`,`does_not_start_with`,`less_than`,`greater_than`,`is_before_time`,`is_after_time`,`is_before_date` or `is_after_date`.
- `value` - (Mandatory) The value or values to compare against

~> Due to data type and validation restrictions liquid templates are not supported for the `transitions.matches.conditions.type` argument. Please see the widget documentation to determine whether other arguments support liquid templates

---

An `offset` block supports the following:

- `x` - (Optional) The x coordinate to display the split based on widget in the Studio console. The default value is `0`
- `y` - (Optional) The y coordinate to display the split based on widget in the Studio console. The default value is `0`

## Attributes Reference

The following attributes are exported:

- `id` - The name of the split based on widget
- `json` - The JSON state definition for the split based on widget
