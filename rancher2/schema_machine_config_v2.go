package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

//Schemas

func machineConfigV2Fields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"generate_name": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Cluster V2 generate name. The pattern to generate machine config name. e.g  generate_name=\"prod-pool1\" will generate \"nc-prod-pool1-?????\" names",
		},
		"fleet_namespace": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
			Default:  "fleet-default",
		},
		"amazonec2_config": {
			Type:          schema.TypeList,
			MaxItems:      1,
			Optional:      true,
			ConflictsWith: []string{"azure_config", "digitalocean_config", "linode_config", "openstack_config", "vsphere_config"},
			Elem: &schema.Resource{
				Schema: machineConfigV2Amazonec2Fields(),
			},
		},
		"azure_config": {
			Type:          schema.TypeList,
			MaxItems:      1,
			Optional:      true,
			ConflictsWith: []string{"amazonec2_config", "digitalocean_config", "linode_config", "openstack_config", "vsphere_config"},
			Elem: &schema.Resource{
				Schema: machineConfigV2AzureFields(),
			},
		},
		"digitalocean_config": {
			Type:          schema.TypeList,
			MaxItems:      1,
			Optional:      true,
			ConflictsWith: []string{"amazonec2_config", "azure_config", "linode_config", "openstack_config", "vsphere_config"},
			Elem: &schema.Resource{
				Schema: machineConfigV2DigitaloceanFields(),
			},
		},
		"kind": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"linode_config": {
			Type:          schema.TypeList,
			MaxItems:      1,
			Optional:      true,
			ConflictsWith: []string{"amazonec2_config", "azure_config", "digitalocean_config", "openstack_config", "vsphere_config"},
			Elem: &schema.Resource{
				Schema: machineConfigV2LinodeFields(),
			},
		},
		"openstack_config": {
			Type:          schema.TypeList,
			MaxItems:      1,
			Optional:      true,
			ConflictsWith: []string{"amazonec2_config", "azure_config", "digitalocean_config", "linode_config", "vsphere_config"},
			Elem: &schema.Resource{
				Schema: machineConfigV2OpenstackFields(),
			},
		},
		"vsphere_config": {
			Type:          schema.TypeList,
			MaxItems:      1,
			Optional:      true,
			ConflictsWith: []string{"amazonec2_config", "azure_config", "digitalocean_config", "linode_config", "openstack_config"},
			Elem: &schema.Resource{
				Schema: machineConfigV2VmwarevsphereFields(),
			},
		},
		"resource_version": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}

	for k, v := range commonAnnotationLabelFields() {
		s[k] = v
	}

	return s
}
