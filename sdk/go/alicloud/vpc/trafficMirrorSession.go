// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPC Traffic Mirror Session resource.
//
// For information about VPC Traffic Mirror Session and how to use it, see [What is Traffic Mirror Session](https://www.alibabacloud.com/help/en/doc-detail/261364.htm).
//
// > **NOTE:** Available in v1.142.0+.
//
// ## Import
//
// VPC Traffic Mirror Session can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:vpc/trafficMirrorSession:TrafficMirrorSession example <id>
//
// ```
type TrafficMirrorSession struct {
	pulumi.CustomResourceState

	// The dry run.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// Specifies whether to enable traffic mirror sessions. default to `false`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The priority of the traffic mirror session. Valid values: `1` to `32766`. A smaller value indicates a higher priority. You cannot specify the same priority for traffic mirror sessions that are created in the same region with the same Alibaba Cloud account.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// The state of the traffic mirror session. Valid values: `Creating`, `Created`, `Modifying` and `Deleting`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the filter.
	TrafficMirrorFilterId pulumi.StringOutput `pulumi:"trafficMirrorFilterId"`
	// The description of the traffic mirror session. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
	TrafficMirrorSessionDescription pulumi.StringPtrOutput `pulumi:"trafficMirrorSessionDescription"`
	// The name of the traffic mirror session. The name must be `2` to `128` characters in length and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
	TrafficMirrorSessionName pulumi.StringPtrOutput `pulumi:"trafficMirrorSessionName"`
	// The ID of the mirror source. You can specify only an elastic network interface (ENI) as the mirror source. **NOTE:** Only one mirror source can be added to a traffic mirror session.
	TrafficMirrorSourceIds pulumi.StringArrayOutput `pulumi:"trafficMirrorSourceIds"`
	// The ID of the mirror destination. You can specify only an ENI or a Server Load Balancer (SLB) instance as a mirror destination.
	TrafficMirrorTargetId pulumi.StringOutput `pulumi:"trafficMirrorTargetId"`
	// The type of the mirror destination. Valid values: `NetworkInterface` or `SLB`. `NetworkInterface`: an ENI. `SLB`: an internal-facing SLB instance
	TrafficMirrorTargetType pulumi.StringOutput `pulumi:"trafficMirrorTargetType"`
	// The VXLAN network identifier (VNI) that is used to distinguish different mirrored traffic. Valid values: `0` to `16777215`. You can specify VNIs for the traffic mirror destination to identify mirrored traffic from different sessions. If you do not specify a VNI, the system randomly allocates a VNI. If you want the system to randomly allocate a VNI, ignore this parameter.
	VirtualNetworkId pulumi.IntOutput `pulumi:"virtualNetworkId"`
}

// NewTrafficMirrorSession registers a new resource with the given unique name, arguments, and options.
func NewTrafficMirrorSession(ctx *pulumi.Context,
	name string, args *TrafficMirrorSessionArgs, opts ...pulumi.ResourceOption) (*TrafficMirrorSession, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	if args.TrafficMirrorFilterId == nil {
		return nil, errors.New("invalid value for required argument 'TrafficMirrorFilterId'")
	}
	if args.TrafficMirrorSourceIds == nil {
		return nil, errors.New("invalid value for required argument 'TrafficMirrorSourceIds'")
	}
	if args.TrafficMirrorTargetId == nil {
		return nil, errors.New("invalid value for required argument 'TrafficMirrorTargetId'")
	}
	if args.TrafficMirrorTargetType == nil {
		return nil, errors.New("invalid value for required argument 'TrafficMirrorTargetType'")
	}
	var resource TrafficMirrorSession
	err := ctx.RegisterResource("alicloud:vpc/trafficMirrorSession:TrafficMirrorSession", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficMirrorSession gets an existing TrafficMirrorSession resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficMirrorSession(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficMirrorSessionState, opts ...pulumi.ResourceOption) (*TrafficMirrorSession, error) {
	var resource TrafficMirrorSession
	err := ctx.ReadResource("alicloud:vpc/trafficMirrorSession:TrafficMirrorSession", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficMirrorSession resources.
type trafficMirrorSessionState struct {
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// Specifies whether to enable traffic mirror sessions. default to `false`.
	Enabled *bool `pulumi:"enabled"`
	// The priority of the traffic mirror session. Valid values: `1` to `32766`. A smaller value indicates a higher priority. You cannot specify the same priority for traffic mirror sessions that are created in the same region with the same Alibaba Cloud account.
	Priority *int `pulumi:"priority"`
	// The state of the traffic mirror session. Valid values: `Creating`, `Created`, `Modifying` and `Deleting`.
	Status *string `pulumi:"status"`
	// The ID of the filter.
	TrafficMirrorFilterId *string `pulumi:"trafficMirrorFilterId"`
	// The description of the traffic mirror session. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
	TrafficMirrorSessionDescription *string `pulumi:"trafficMirrorSessionDescription"`
	// The name of the traffic mirror session. The name must be `2` to `128` characters in length and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
	TrafficMirrorSessionName *string `pulumi:"trafficMirrorSessionName"`
	// The ID of the mirror source. You can specify only an elastic network interface (ENI) as the mirror source. **NOTE:** Only one mirror source can be added to a traffic mirror session.
	TrafficMirrorSourceIds []string `pulumi:"trafficMirrorSourceIds"`
	// The ID of the mirror destination. You can specify only an ENI or a Server Load Balancer (SLB) instance as a mirror destination.
	TrafficMirrorTargetId *string `pulumi:"trafficMirrorTargetId"`
	// The type of the mirror destination. Valid values: `NetworkInterface` or `SLB`. `NetworkInterface`: an ENI. `SLB`: an internal-facing SLB instance
	TrafficMirrorTargetType *string `pulumi:"trafficMirrorTargetType"`
	// The VXLAN network identifier (VNI) that is used to distinguish different mirrored traffic. Valid values: `0` to `16777215`. You can specify VNIs for the traffic mirror destination to identify mirrored traffic from different sessions. If you do not specify a VNI, the system randomly allocates a VNI. If you want the system to randomly allocate a VNI, ignore this parameter.
	VirtualNetworkId *int `pulumi:"virtualNetworkId"`
}

type TrafficMirrorSessionState struct {
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// Specifies whether to enable traffic mirror sessions. default to `false`.
	Enabled pulumi.BoolPtrInput
	// The priority of the traffic mirror session. Valid values: `1` to `32766`. A smaller value indicates a higher priority. You cannot specify the same priority for traffic mirror sessions that are created in the same region with the same Alibaba Cloud account.
	Priority pulumi.IntPtrInput
	// The state of the traffic mirror session. Valid values: `Creating`, `Created`, `Modifying` and `Deleting`.
	Status pulumi.StringPtrInput
	// The ID of the filter.
	TrafficMirrorFilterId pulumi.StringPtrInput
	// The description of the traffic mirror session. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
	TrafficMirrorSessionDescription pulumi.StringPtrInput
	// The name of the traffic mirror session. The name must be `2` to `128` characters in length and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
	TrafficMirrorSessionName pulumi.StringPtrInput
	// The ID of the mirror source. You can specify only an elastic network interface (ENI) as the mirror source. **NOTE:** Only one mirror source can be added to a traffic mirror session.
	TrafficMirrorSourceIds pulumi.StringArrayInput
	// The ID of the mirror destination. You can specify only an ENI or a Server Load Balancer (SLB) instance as a mirror destination.
	TrafficMirrorTargetId pulumi.StringPtrInput
	// The type of the mirror destination. Valid values: `NetworkInterface` or `SLB`. `NetworkInterface`: an ENI. `SLB`: an internal-facing SLB instance
	TrafficMirrorTargetType pulumi.StringPtrInput
	// The VXLAN network identifier (VNI) that is used to distinguish different mirrored traffic. Valid values: `0` to `16777215`. You can specify VNIs for the traffic mirror destination to identify mirrored traffic from different sessions. If you do not specify a VNI, the system randomly allocates a VNI. If you want the system to randomly allocate a VNI, ignore this parameter.
	VirtualNetworkId pulumi.IntPtrInput
}

func (TrafficMirrorSessionState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorSessionState)(nil)).Elem()
}

type trafficMirrorSessionArgs struct {
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// Specifies whether to enable traffic mirror sessions. default to `false`.
	Enabled *bool `pulumi:"enabled"`
	// The priority of the traffic mirror session. Valid values: `1` to `32766`. A smaller value indicates a higher priority. You cannot specify the same priority for traffic mirror sessions that are created in the same region with the same Alibaba Cloud account.
	Priority int `pulumi:"priority"`
	// The ID of the filter.
	TrafficMirrorFilterId string `pulumi:"trafficMirrorFilterId"`
	// The description of the traffic mirror session. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
	TrafficMirrorSessionDescription *string `pulumi:"trafficMirrorSessionDescription"`
	// The name of the traffic mirror session. The name must be `2` to `128` characters in length and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
	TrafficMirrorSessionName *string `pulumi:"trafficMirrorSessionName"`
	// The ID of the mirror source. You can specify only an elastic network interface (ENI) as the mirror source. **NOTE:** Only one mirror source can be added to a traffic mirror session.
	TrafficMirrorSourceIds []string `pulumi:"trafficMirrorSourceIds"`
	// The ID of the mirror destination. You can specify only an ENI or a Server Load Balancer (SLB) instance as a mirror destination.
	TrafficMirrorTargetId string `pulumi:"trafficMirrorTargetId"`
	// The type of the mirror destination. Valid values: `NetworkInterface` or `SLB`. `NetworkInterface`: an ENI. `SLB`: an internal-facing SLB instance
	TrafficMirrorTargetType string `pulumi:"trafficMirrorTargetType"`
	// The VXLAN network identifier (VNI) that is used to distinguish different mirrored traffic. Valid values: `0` to `16777215`. You can specify VNIs for the traffic mirror destination to identify mirrored traffic from different sessions. If you do not specify a VNI, the system randomly allocates a VNI. If you want the system to randomly allocate a VNI, ignore this parameter.
	VirtualNetworkId *int `pulumi:"virtualNetworkId"`
}

// The set of arguments for constructing a TrafficMirrorSession resource.
type TrafficMirrorSessionArgs struct {
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// Specifies whether to enable traffic mirror sessions. default to `false`.
	Enabled pulumi.BoolPtrInput
	// The priority of the traffic mirror session. Valid values: `1` to `32766`. A smaller value indicates a higher priority. You cannot specify the same priority for traffic mirror sessions that are created in the same region with the same Alibaba Cloud account.
	Priority pulumi.IntInput
	// The ID of the filter.
	TrafficMirrorFilterId pulumi.StringInput
	// The description of the traffic mirror session. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
	TrafficMirrorSessionDescription pulumi.StringPtrInput
	// The name of the traffic mirror session. The name must be `2` to `128` characters in length and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
	TrafficMirrorSessionName pulumi.StringPtrInput
	// The ID of the mirror source. You can specify only an elastic network interface (ENI) as the mirror source. **NOTE:** Only one mirror source can be added to a traffic mirror session.
	TrafficMirrorSourceIds pulumi.StringArrayInput
	// The ID of the mirror destination. You can specify only an ENI or a Server Load Balancer (SLB) instance as a mirror destination.
	TrafficMirrorTargetId pulumi.StringInput
	// The type of the mirror destination. Valid values: `NetworkInterface` or `SLB`. `NetworkInterface`: an ENI. `SLB`: an internal-facing SLB instance
	TrafficMirrorTargetType pulumi.StringInput
	// The VXLAN network identifier (VNI) that is used to distinguish different mirrored traffic. Valid values: `0` to `16777215`. You can specify VNIs for the traffic mirror destination to identify mirrored traffic from different sessions. If you do not specify a VNI, the system randomly allocates a VNI. If you want the system to randomly allocate a VNI, ignore this parameter.
	VirtualNetworkId pulumi.IntPtrInput
}

func (TrafficMirrorSessionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorSessionArgs)(nil)).Elem()
}

type TrafficMirrorSessionInput interface {
	pulumi.Input

	ToTrafficMirrorSessionOutput() TrafficMirrorSessionOutput
	ToTrafficMirrorSessionOutputWithContext(ctx context.Context) TrafficMirrorSessionOutput
}

func (*TrafficMirrorSession) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficMirrorSession)(nil)).Elem()
}

func (i *TrafficMirrorSession) ToTrafficMirrorSessionOutput() TrafficMirrorSessionOutput {
	return i.ToTrafficMirrorSessionOutputWithContext(context.Background())
}

func (i *TrafficMirrorSession) ToTrafficMirrorSessionOutputWithContext(ctx context.Context) TrafficMirrorSessionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMirrorSessionOutput)
}

// TrafficMirrorSessionArrayInput is an input type that accepts TrafficMirrorSessionArray and TrafficMirrorSessionArrayOutput values.
// You can construct a concrete instance of `TrafficMirrorSessionArrayInput` via:
//
//	TrafficMirrorSessionArray{ TrafficMirrorSessionArgs{...} }
type TrafficMirrorSessionArrayInput interface {
	pulumi.Input

	ToTrafficMirrorSessionArrayOutput() TrafficMirrorSessionArrayOutput
	ToTrafficMirrorSessionArrayOutputWithContext(context.Context) TrafficMirrorSessionArrayOutput
}

type TrafficMirrorSessionArray []TrafficMirrorSessionInput

func (TrafficMirrorSessionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficMirrorSession)(nil)).Elem()
}

func (i TrafficMirrorSessionArray) ToTrafficMirrorSessionArrayOutput() TrafficMirrorSessionArrayOutput {
	return i.ToTrafficMirrorSessionArrayOutputWithContext(context.Background())
}

func (i TrafficMirrorSessionArray) ToTrafficMirrorSessionArrayOutputWithContext(ctx context.Context) TrafficMirrorSessionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMirrorSessionArrayOutput)
}

// TrafficMirrorSessionMapInput is an input type that accepts TrafficMirrorSessionMap and TrafficMirrorSessionMapOutput values.
// You can construct a concrete instance of `TrafficMirrorSessionMapInput` via:
//
//	TrafficMirrorSessionMap{ "key": TrafficMirrorSessionArgs{...} }
type TrafficMirrorSessionMapInput interface {
	pulumi.Input

	ToTrafficMirrorSessionMapOutput() TrafficMirrorSessionMapOutput
	ToTrafficMirrorSessionMapOutputWithContext(context.Context) TrafficMirrorSessionMapOutput
}

type TrafficMirrorSessionMap map[string]TrafficMirrorSessionInput

func (TrafficMirrorSessionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficMirrorSession)(nil)).Elem()
}

func (i TrafficMirrorSessionMap) ToTrafficMirrorSessionMapOutput() TrafficMirrorSessionMapOutput {
	return i.ToTrafficMirrorSessionMapOutputWithContext(context.Background())
}

func (i TrafficMirrorSessionMap) ToTrafficMirrorSessionMapOutputWithContext(ctx context.Context) TrafficMirrorSessionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMirrorSessionMapOutput)
}

type TrafficMirrorSessionOutput struct{ *pulumi.OutputState }

func (TrafficMirrorSessionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficMirrorSession)(nil)).Elem()
}

func (o TrafficMirrorSessionOutput) ToTrafficMirrorSessionOutput() TrafficMirrorSessionOutput {
	return o
}

func (o TrafficMirrorSessionOutput) ToTrafficMirrorSessionOutputWithContext(ctx context.Context) TrafficMirrorSessionOutput {
	return o
}

// The dry run.
func (o TrafficMirrorSessionOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TrafficMirrorSession) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// Specifies whether to enable traffic mirror sessions. default to `false`.
func (o TrafficMirrorSessionOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TrafficMirrorSession) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The priority of the traffic mirror session. Valid values: `1` to `32766`. A smaller value indicates a higher priority. You cannot specify the same priority for traffic mirror sessions that are created in the same region with the same Alibaba Cloud account.
func (o TrafficMirrorSessionOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *TrafficMirrorSession) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

// The state of the traffic mirror session. Valid values: `Creating`, `Created`, `Modifying` and `Deleting`.
func (o TrafficMirrorSessionOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorSession) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The ID of the filter.
func (o TrafficMirrorSessionOutput) TrafficMirrorFilterId() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorSession) pulumi.StringOutput { return v.TrafficMirrorFilterId }).(pulumi.StringOutput)
}

// The description of the traffic mirror session. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
func (o TrafficMirrorSessionOutput) TrafficMirrorSessionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficMirrorSession) pulumi.StringPtrOutput { return v.TrafficMirrorSessionDescription }).(pulumi.StringPtrOutput)
}

// The name of the traffic mirror session. The name must be `2` to `128` characters in length and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
func (o TrafficMirrorSessionOutput) TrafficMirrorSessionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficMirrorSession) pulumi.StringPtrOutput { return v.TrafficMirrorSessionName }).(pulumi.StringPtrOutput)
}

// The ID of the mirror source. You can specify only an elastic network interface (ENI) as the mirror source. **NOTE:** Only one mirror source can be added to a traffic mirror session.
func (o TrafficMirrorSessionOutput) TrafficMirrorSourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TrafficMirrorSession) pulumi.StringArrayOutput { return v.TrafficMirrorSourceIds }).(pulumi.StringArrayOutput)
}

// The ID of the mirror destination. You can specify only an ENI or a Server Load Balancer (SLB) instance as a mirror destination.
func (o TrafficMirrorSessionOutput) TrafficMirrorTargetId() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorSession) pulumi.StringOutput { return v.TrafficMirrorTargetId }).(pulumi.StringOutput)
}

// The type of the mirror destination. Valid values: `NetworkInterface` or `SLB`. `NetworkInterface`: an ENI. `SLB`: an internal-facing SLB instance
func (o TrafficMirrorSessionOutput) TrafficMirrorTargetType() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorSession) pulumi.StringOutput { return v.TrafficMirrorTargetType }).(pulumi.StringOutput)
}

// The VXLAN network identifier (VNI) that is used to distinguish different mirrored traffic. Valid values: `0` to `16777215`. You can specify VNIs for the traffic mirror destination to identify mirrored traffic from different sessions. If you do not specify a VNI, the system randomly allocates a VNI. If you want the system to randomly allocate a VNI, ignore this parameter.
func (o TrafficMirrorSessionOutput) VirtualNetworkId() pulumi.IntOutput {
	return o.ApplyT(func(v *TrafficMirrorSession) pulumi.IntOutput { return v.VirtualNetworkId }).(pulumi.IntOutput)
}

type TrafficMirrorSessionArrayOutput struct{ *pulumi.OutputState }

func (TrafficMirrorSessionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficMirrorSession)(nil)).Elem()
}

func (o TrafficMirrorSessionArrayOutput) ToTrafficMirrorSessionArrayOutput() TrafficMirrorSessionArrayOutput {
	return o
}

func (o TrafficMirrorSessionArrayOutput) ToTrafficMirrorSessionArrayOutputWithContext(ctx context.Context) TrafficMirrorSessionArrayOutput {
	return o
}

func (o TrafficMirrorSessionArrayOutput) Index(i pulumi.IntInput) TrafficMirrorSessionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TrafficMirrorSession {
		return vs[0].([]*TrafficMirrorSession)[vs[1].(int)]
	}).(TrafficMirrorSessionOutput)
}

type TrafficMirrorSessionMapOutput struct{ *pulumi.OutputState }

func (TrafficMirrorSessionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficMirrorSession)(nil)).Elem()
}

func (o TrafficMirrorSessionMapOutput) ToTrafficMirrorSessionMapOutput() TrafficMirrorSessionMapOutput {
	return o
}

func (o TrafficMirrorSessionMapOutput) ToTrafficMirrorSessionMapOutputWithContext(ctx context.Context) TrafficMirrorSessionMapOutput {
	return o
}

func (o TrafficMirrorSessionMapOutput) MapIndex(k pulumi.StringInput) TrafficMirrorSessionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TrafficMirrorSession {
		return vs[0].(map[string]*TrafficMirrorSession)[vs[1].(string)]
	}).(TrafficMirrorSessionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMirrorSessionInput)(nil)).Elem(), &TrafficMirrorSession{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMirrorSessionArrayInput)(nil)).Elem(), TrafficMirrorSessionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMirrorSessionMapInput)(nil)).Elem(), TrafficMirrorSessionMap{})
	pulumi.RegisterOutputType(TrafficMirrorSessionOutput{})
	pulumi.RegisterOutputType(TrafficMirrorSessionArrayOutput{})
	pulumi.RegisterOutputType(TrafficMirrorSessionMapOutput{})
}
