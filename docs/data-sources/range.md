---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_range Data Source - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This data source can read the Range.
---

# fmc_range (Data Source)

This data source can read the Range.

## Example Usage

```terraform
data "fmc_range" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domain` (String) The name of the FMC domain
- `id` (String) The id of the object
- `name` (String) User-created name of the resource.

### Read-Only

- `description` (String) Optional user-created description.
- `ip_range` (String) Range of addresses, IPv4 or IPv6.
- `overridable` (Boolean) Indicates whether object values can be overridden.
