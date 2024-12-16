// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &DeviceHAPairDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceHAPairDataSource{}
)

func NewDeviceHAPairDataSource() datasource.DataSource {
	return &DeviceHAPairDataSource{}
}

type DeviceHAPairDataSource struct {
	client *fmc.Client
}

func (d *DeviceHAPairDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_ha_pair"
}

func (d *DeviceHAPairDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Device HA Pair.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Optional:            true,
				Computed:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "The name of the FMC domain",
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the High Availability Pair.",
				Optional:            true,
				Computed:            true,
			},
			"primary_device_id": schema.StringAttribute{
				MarkdownDescription: "ID of primary FTD in the HA Pair.",
				Computed:            true,
			},
			"secondary_device_id": schema.StringAttribute{
				MarkdownDescription: "ID of secondary FTD in the HA Pair.",
				Computed:            true,
			},
			"is_encryption_enabled": schema.BoolAttribute{
				MarkdownDescription: "Boolean field to enable encryption",
				Computed:            true,
			},
			"use_same_link_for_failovers": schema.BoolAttribute{
				MarkdownDescription: "Boolean field to enable same link for failovers",
				Computed:            true,
			},
			"shared_key": schema.StringAttribute{
				MarkdownDescription: "Pass the unique shared key if needed.",
				Computed:            true,
			},
			"enc_key_generation_scheme": schema.StringAttribute{
				MarkdownDescription: "Select the encyption key generation scheme.",
				Computed:            true,
			},
			"lan_failover_standby_ip": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"lan_failover_active_ip": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"lan_failover_name": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"lan_failover_netmask": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"lan_failover_ipv6_addr": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"lan_failover_interface_name": schema.StringAttribute{
				MarkdownDescription: "Name of physical interface",
				Optional:            true,
				Computed:            true,
			},
			"lan_failover_interface_id": schema.StringAttribute{
				MarkdownDescription: "ID of physical interface.",
				Computed:            true,
			},
			"lan_failover_interface_type": schema.StringAttribute{
				MarkdownDescription: "Type of physical interface.",
				Computed:            true,
			},
			"stateful_failover_standby_ip": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"stateful_failover_active_ip": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"stateful_failover_name": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"stateful_failover_subnet_mask": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"stateful_failover_ipv6_addr": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"stateful_failover_interface_name": schema.StringAttribute{
				MarkdownDescription: "Name of physical interface",
				Optional:            true,
				Computed:            true,
			},
			"stateful_failover_interface_id": schema.StringAttribute{
				MarkdownDescription: "ID of physical interface.",
				Computed:            true,
			},
			"stateful_failover_interface_type": schema.StringAttribute{
				MarkdownDescription: "Type of physical interface.",
				Computed:            true,
			},
			"action": schema.StringAttribute{
				MarkdownDescription: "FTD HA PUT operation action. Specifically used for manual switch. HA Break will be triggered when you run terraform destroy",
				Computed:            true,
			},
		},
	}
}
func (d *DeviceHAPairDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *DeviceHAPairDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*FmcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceHAPairDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DeviceHAPair

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set request domain if provided
	reqMods := [](func(*fmc.Req)){}
	if !config.Domain.IsNull() && config.Domain.ValueString() != "" {
		reqMods = append(reqMods, fmc.DomainName(config.Domain.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))
	if config.Id.IsNull() && !config.Name.IsNull() {
		offset := 0
		limit := 1000
		for page := 1; ; page++ {
			queryString := fmt.Sprintf("?limit=%d&offset=%d", limit, offset)
			res, err := d.client.Get(config.getPath()+queryString, reqMods...)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
				return
			}
			if value := res.Get("items"); len(value.Array()) > 0 {
				value.ForEach(func(k, v gjson.Result) bool {
					if config.Name.ValueString() == v.Get("name").String() {
						config.Id = types.StringValue(v.Get("id").String())
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
						return false
					}
					return true
				})
			}
			if !config.Id.IsNull() || !res.Get("paging.next.0").Exists() {
				break
			}
			offset += limit
		}

		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %s", config.Name.ValueString()))
			return
		}
	}
	urlPath := config.getPath() + "/" + url.QueryEscape(config.Id.ValueString())
	res, err := d.client.Get(urlPath, reqMods...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read