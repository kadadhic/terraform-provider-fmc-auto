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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type ICMPv4Object struct {
	Id          types.String `tfsdk:"id"`
	Domain      types.String `tfsdk:"domain"`
	IcmpType    types.Int64  `tfsdk:"icmp_type"`
	Code        types.Int64  `tfsdk:"code"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Overridable types.Bool   `tfsdk:"overridable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ICMPv4Object) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/icmpv4objects"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ICMPv4Object) toBody(ctx context.Context, state ICMPv4Object) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.IcmpType.IsNull() {
		body, _ = sjson.Set(body, "icmpType", data.IcmpType.ValueInt64())
	}
	if !data.Code.IsNull() {
		body, _ = sjson.Set(body, "code", data.Code.ValueInt64())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.Overridable.IsNull() {
		body, _ = sjson.Set(body, "overridable", data.Overridable.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ICMPv4Object) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("icmpType"); value.Exists() {
		data.IcmpType = types.Int64Value(value.Int())
	} else {
		data.IcmpType = types.Int64Null()
	}
	if value := res.Get("code"); value.Exists() {
		data.Code = types.Int64Value(value.Int())
	} else {
		data.Code = types.Int64Null()
	}
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("overridable"); value.Exists() {
		data.Overridable = types.BoolValue(value.Bool())
	} else {
		data.Overridable = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ICMPv4Object) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("icmpType"); value.Exists() && !data.IcmpType.IsNull() {
		data.IcmpType = types.Int64Value(value.Int())
	} else {
		data.IcmpType = types.Int64Null()
	}
	if value := res.Get("code"); value.Exists() && !data.Code.IsNull() {
		data.Code = types.Int64Value(value.Int())
	} else {
		data.Code = types.Int64Null()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("overridable"); value.Exists() && !data.Overridable.IsNull() {
		data.Overridable = types.BoolValue(value.Bool())
	} else {
		data.Overridable = types.BoolNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *ICMPv4Object) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.IcmpType.IsNull() {
		return false
	}
	if !data.Code.IsNull() {
		return false
	}
	if !data.Name.IsNull() {
		return false
	}
	if !data.Description.IsNull() {
		return false
	}
	if !data.Overridable.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull