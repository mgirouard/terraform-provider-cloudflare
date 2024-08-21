---
page_title: "cloudflare_zone_cache_variants Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zone_cache_variants (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `zone_id` (String) Identifier

### Optional

- `id` (String) ID of the zone setting.
- `modified_on` (String) last time this setting was modified.
- `value` (Attributes) Value of the zone setting. (see [below for nested schema](#nestedatt--value))

<a id="nestedatt--value"></a>
### Nested Schema for `value`

Optional:

- `avif` (List of String) List of strings with the MIME types of all the variants that should be served for avif.
- `bmp` (List of String) List of strings with the MIME types of all the variants that should be served for bmp.
- `gif` (List of String) List of strings with the MIME types of all the variants that should be served for gif.
- `jp2` (List of String) List of strings with the MIME types of all the variants that should be served for jp2.
- `jpeg` (List of String) List of strings with the MIME types of all the variants that should be served for jpeg.
- `jpg` (List of String) List of strings with the MIME types of all the variants that should be served for jpg.
- `jpg2` (List of String) List of strings with the MIME types of all the variants that should be served for jpg2.
- `png` (List of String) List of strings with the MIME types of all the variants that should be served for png.
- `tif` (List of String) List of strings with the MIME types of all the variants that should be served for tif.
- `tiff` (List of String) List of strings with the MIME types of all the variants that should be served for tiff.
- `webp` (List of String) List of strings with the MIME types of all the variants that should be served for webp.

