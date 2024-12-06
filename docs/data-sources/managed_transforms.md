---
page_title: "cloudflare_managed_transforms Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_managed_transforms (Data Source)



## Example Usage

```terraform
data "cloudflare_managed_transforms" "example_managed_transforms" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `zone_id` (String) Identifier

### Read-Only

- `managed_request_headers` (Attributes List) (see [below for nested schema](#nestedatt--managed_request_headers))
- `managed_response_headers` (Attributes List) (see [below for nested schema](#nestedatt--managed_response_headers))

<a id="nestedatt--managed_request_headers"></a>
### Nested Schema for `managed_request_headers`

Read-Only:

- `enabled` (Boolean) When true, the Managed Transform is enabled.
- `id` (String) Human-readable identifier of the Managed Transform.


<a id="nestedatt--managed_response_headers"></a>
### Nested Schema for `managed_response_headers`

Read-Only:

- `enabled` (Boolean) When true, the Managed Transform is enabled.
- `id` (String) Human-readable identifier of the Managed Transform.

