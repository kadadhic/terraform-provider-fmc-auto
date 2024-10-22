---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_port_group Resource - terraform-provider-fmc"
subcategory: "Objects"
description: |-
  This resource can manage a Port Group.
---

# fmc_port_group (Resource)

This resource can manage a Port Group.

## Example Usage

```terraform
resource "fmc_port_group" "example" {
  name        = "portgroup_obj1"
  type        = "PortObjectGroup"
  description = "My port group description"
  ports = [
    {
      id   = "0050568A-232D-0ed3-0000-004294971602"
      type = "ProtocolPortObject"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) User-created name of the object.
- `ports` (Attributes Set) (see [below for nested schema](#nestedatt--ports))
- `type` (String)

### Optional

- `description` (String) Optional user-created description.
- `domain` (String) The name of the FMC domain
- `overridable` (Boolean) Indicates whether object values can be overridden.

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--ports"></a>
### Nested Schema for `ports`

Required:

- `type` (String) - Choices: `ProtocolPortObject`, `ICMPV6Object`, `ICMPV4Object`

Optional:

- `id` (String) UUID of the port (such as fmc_port.test.id, etc.).

## Import

Import is supported using the following syntax:

```shell
terraform import fmc_port_group.example "<id>"
```