// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rocketmq

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a list of ONS Groups in an Alibaba Cloud account according to the specified filters.
//
// > **NOTE:** Available in 1.53.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/rocketmq"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		name := "onsInstanceName"
// 		if param := cfg.Get("name"); param != "" {
// 			name = param
// 		}
// 		groupName := "GID-onsGroupDatasourceName"
// 		if param := cfg.Get("groupName"); param != "" {
// 			groupName = param
// 		}
// 		defaultInstance, err := rocketmq.NewInstance(ctx, "defaultInstance", &rocketmq.InstanceArgs{
// 			InstanceName: pulumi.String(name),
// 			Remark:       pulumi.String("default_ons_instance_remark"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultGroup, err := rocketmq.NewGroup(ctx, "defaultGroup", &rocketmq.GroupArgs{
// 			GroupName:  pulumi.String(groupName),
// 			InstanceId: defaultInstance.ID(),
// 			Remark:     pulumi.String("dafault_ons_group_remark"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstGroupName", groupsDs.ApplyT(func(groupsDs rocketmq.GetGroupsResult) (string, error) {
// 			return groupsDs.Groups[0].GroupName, nil
// 		}).(pulumi.StringOutput))
// 		return nil
// 	})
// }
// ```
func GetGroups(ctx *pulumi.Context, args *GetGroupsArgs, opts ...pulumi.InvokeOption) (*GetGroupsResult, error) {
	var rv GetGroupsResult
	err := ctx.Invoke("alicloud:rocketmq/getGroups:getGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroups.
type GetGroupsArgs struct {
	// A regex string to filter results by the group name.
	GroupIdRegex *string `pulumi:"groupIdRegex"`
	// Specify the protocol applicable to the created Group ID. Valid values: `tcp`, `http`. Default to `tcp`.
	GroupType *string `pulumi:"groupType"`
	// A list of group names.
	Ids []string `pulumi:"ids"`
	// ID of the ONS Instance that owns the groups.
	InstanceId string  `pulumi:"instanceId"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// A map of tags assigned to the Ons instance.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getGroups.
type GetGroupsResult struct {
	GroupIdRegex *string `pulumi:"groupIdRegex"`
	// Specify the protocol applicable to the created Group ID.
	GroupType *string `pulumi:"groupType"`
	// A list of groups. Each element contains the following attributes:
	Groups []GetGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of group names.
	Ids        []string `pulumi:"ids"`
	InstanceId string   `pulumi:"instanceId"`
	NameRegex  *string  `pulumi:"nameRegex"`
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A map of tags assigned to the Ons group.
	Tags map[string]interface{} `pulumi:"tags"`
}

func GetGroupsOutput(ctx *pulumi.Context, args GetGroupsOutputArgs, opts ...pulumi.InvokeOption) GetGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGroupsResult, error) {
			args := v.(GetGroupsArgs)
			r, err := GetGroups(ctx, &args, opts...)
			return *r, err
		}).(GetGroupsResultOutput)
}

// A collection of arguments for invoking getGroups.
type GetGroupsOutputArgs struct {
	// A regex string to filter results by the group name.
	GroupIdRegex pulumi.StringPtrInput `pulumi:"groupIdRegex"`
	// Specify the protocol applicable to the created Group ID. Valid values: `tcp`, `http`. Default to `tcp`.
	GroupType pulumi.StringPtrInput `pulumi:"groupType"`
	// A list of group names.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// ID of the ONS Instance that owns the groups.
	InstanceId pulumi.StringInput    `pulumi:"instanceId"`
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// A map of tags assigned to the Ons instance.
	Tags pulumi.MapInput `pulumi:"tags"`
}

func (GetGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getGroups.
type GetGroupsResultOutput struct{ *pulumi.OutputState }

func (GetGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsResult)(nil)).Elem()
}

func (o GetGroupsResultOutput) ToGetGroupsResultOutput() GetGroupsResultOutput {
	return o
}

func (o GetGroupsResultOutput) ToGetGroupsResultOutputWithContext(ctx context.Context) GetGroupsResultOutput {
	return o
}

func (o GetGroupsResultOutput) GroupIdRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.GroupIdRegex }).(pulumi.StringPtrOutput)
}

// Specify the protocol applicable to the created Group ID.
func (o GetGroupsResultOutput) GroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.GroupType }).(pulumi.StringPtrOutput)
}

// A list of groups. Each element contains the following attributes:
func (o GetGroupsResultOutput) Groups() GetGroupsGroupArrayOutput {
	return o.ApplyT(func(v GetGroupsResult) []GetGroupsGroup { return v.Groups }).(GetGroupsGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of group names.
func (o GetGroupsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetGroupsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetGroupsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetGroupsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetGroupsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetGroupsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// A map of tags assigned to the Ons group.
func (o GetGroupsResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetGroupsResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGroupsResultOutput{})
}
