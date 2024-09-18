---
page_title: "cloudflare_worker_secret Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_worker_secret (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filter` (Attributes) (see [below for nested schema](#nestedatt--filter))
- `name` (String) The name of this secret, this is what will be to access it inside the Worker.
- `type` (String) The type of secret to put.

<a id="nestedatt--filter"></a>
### Nested Schema for `filter`

Required:

- `account_id` (String) Identifier
- `dispatch_namespace` (String) Name of the Workers for Platforms dispatch namespace.
- `script_name` (String) Name of the script, used in URLs and route configuration.

