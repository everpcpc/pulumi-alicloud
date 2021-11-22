// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfirewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Cloud Firewall Control Policies of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.129.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cloudfirewall"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudfirewall.GetControlPolicies(ctx, &cloudfirewall.GetControlPoliciesArgs{
// 			Direction: "in",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetControlPolicies(ctx *pulumi.Context, args *GetControlPoliciesArgs, opts ...pulumi.InvokeOption) (*GetControlPoliciesResult, error) {
	var rv GetControlPoliciesResult
	err := ctx.Invoke("alicloud:cloudfirewall/getControlPolicies:getControlPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getControlPolicies.
type GetControlPoliciesArgs struct {
	// The action that Cloud Firewall performs on the traffic. Valid values: `accept`, `drop`, `log`.
	AclAction *string `pulumi:"aclAction"`
	// The unique ID of the access control policy.
	AclUuid *string `pulumi:"aclUuid"`
	// The description of the access control policy.
	Description *string `pulumi:"description"`
	// The destination address defined in the access control policy.
	Destination *string `pulumi:"destination"`
	// The direction of traffic to which the access control policy applies. Valid values: `in`, `out`.
	Direction string `pulumi:"direction"`
	// The ip version.
	IpVersion *string `pulumi:"ipVersion"`
	// DestPortGroupPorts. Valid values: `en`, `zh`.
	Lang       *string `pulumi:"lang"`
	OutputFile *string `pulumi:"outputFile"`
	// The protocol type of traffic to which the access control policy applies. Valid values: If `direction` is `in`, the valid value is `ANY`. If `direction` is `out`, the valid values are `ANY`, `TCP`, `UDP`, `ICMP`.
	Proto *string `pulumi:"proto"`
	// The source address defined in the access control policy.
	Source *string `pulumi:"source"`
	// The source IP address of the request.
	SourceIp *string `pulumi:"sourceIp"`
}

// A collection of values returned by getControlPolicies.
type GetControlPoliciesResult struct {
	AclAction   *string `pulumi:"aclAction"`
	AclUuid     *string `pulumi:"aclUuid"`
	Description *string `pulumi:"description"`
	Destination *string `pulumi:"destination"`
	Direction   string  `pulumi:"direction"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                     `pulumi:"id"`
	Ids        []string                   `pulumi:"ids"`
	IpVersion  *string                    `pulumi:"ipVersion"`
	Lang       *string                    `pulumi:"lang"`
	OutputFile *string                    `pulumi:"outputFile"`
	Policies   []GetControlPoliciesPolicy `pulumi:"policies"`
	Proto      *string                    `pulumi:"proto"`
	Source     *string                    `pulumi:"source"`
	SourceIp   *string                    `pulumi:"sourceIp"`
}

func GetControlPoliciesOutput(ctx *pulumi.Context, args GetControlPoliciesOutputArgs, opts ...pulumi.InvokeOption) GetControlPoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetControlPoliciesResult, error) {
			args := v.(GetControlPoliciesArgs)
			r, err := GetControlPolicies(ctx, &args, opts...)
			return *r, err
		}).(GetControlPoliciesResultOutput)
}

// A collection of arguments for invoking getControlPolicies.
type GetControlPoliciesOutputArgs struct {
	// The action that Cloud Firewall performs on the traffic. Valid values: `accept`, `drop`, `log`.
	AclAction pulumi.StringPtrInput `pulumi:"aclAction"`
	// The unique ID of the access control policy.
	AclUuid pulumi.StringPtrInput `pulumi:"aclUuid"`
	// The description of the access control policy.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The destination address defined in the access control policy.
	Destination pulumi.StringPtrInput `pulumi:"destination"`
	// The direction of traffic to which the access control policy applies. Valid values: `in`, `out`.
	Direction pulumi.StringInput `pulumi:"direction"`
	// The ip version.
	IpVersion pulumi.StringPtrInput `pulumi:"ipVersion"`
	// DestPortGroupPorts. Valid values: `en`, `zh`.
	Lang       pulumi.StringPtrInput `pulumi:"lang"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The protocol type of traffic to which the access control policy applies. Valid values: If `direction` is `in`, the valid value is `ANY`. If `direction` is `out`, the valid values are `ANY`, `TCP`, `UDP`, `ICMP`.
	Proto pulumi.StringPtrInput `pulumi:"proto"`
	// The source address defined in the access control policy.
	Source pulumi.StringPtrInput `pulumi:"source"`
	// The source IP address of the request.
	SourceIp pulumi.StringPtrInput `pulumi:"sourceIp"`
}

func (GetControlPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetControlPoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getControlPolicies.
type GetControlPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetControlPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetControlPoliciesResult)(nil)).Elem()
}

func (o GetControlPoliciesResultOutput) ToGetControlPoliciesResultOutput() GetControlPoliciesResultOutput {
	return o
}

func (o GetControlPoliciesResultOutput) ToGetControlPoliciesResultOutputWithContext(ctx context.Context) GetControlPoliciesResultOutput {
	return o
}

func (o GetControlPoliciesResultOutput) AclAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.AclAction }).(pulumi.StringPtrOutput)
}

func (o GetControlPoliciesResultOutput) AclUuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.AclUuid }).(pulumi.StringPtrOutput)
}

func (o GetControlPoliciesResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetControlPoliciesResultOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.Destination }).(pulumi.StringPtrOutput)
}

func (o GetControlPoliciesResultOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) string { return v.Direction }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetControlPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetControlPoliciesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetControlPoliciesResultOutput) IpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.IpVersion }).(pulumi.StringPtrOutput)
}

func (o GetControlPoliciesResultOutput) Lang() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.Lang }).(pulumi.StringPtrOutput)
}

func (o GetControlPoliciesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetControlPoliciesResultOutput) Policies() GetControlPoliciesPolicyArrayOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) []GetControlPoliciesPolicy { return v.Policies }).(GetControlPoliciesPolicyArrayOutput)
}

func (o GetControlPoliciesResultOutput) Proto() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.Proto }).(pulumi.StringPtrOutput)
}

func (o GetControlPoliciesResultOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.Source }).(pulumi.StringPtrOutput)
}

func (o GetControlPoliciesResultOutput) SourceIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.SourceIp }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetControlPoliciesResultOutput{})
}
