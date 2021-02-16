// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// DNS domain can be imported using the id or domain name, e.g.
//
// ```sh
//  $ pulumi import alicloud:dns/dnsDomain:DnsDomain example aliyun.com
// ```
type DnsDomain struct {
	pulumi.CustomResourceState

	DnsServers pulumi.StringArrayOutput `pulumi:"dnsServers"`
	// The domain ID.
	DomainId pulumi.StringOutput `pulumi:"domainId"`
	// Name of the domain. This name without suffix can have a string of 1 to 63 characters(domain name subject, excluding suffix), must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// Id of the group in which the domain will add. If not supplied, then use default group.
	GroupId   pulumi.StringPtrOutput `pulumi:"groupId"`
	GroupName pulumi.StringOutput    `pulumi:"groupName"`
	// User language.
	Lang     pulumi.StringPtrOutput `pulumi:"lang"`
	PunyCode pulumi.StringOutput    `pulumi:"punyCode"`
	// Remarks information for your domain name.
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
	// The Id of resource group which the dns domain belongs.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewDnsDomain registers a new resource with the given unique name, arguments, and options.
func NewDnsDomain(ctx *pulumi.Context,
	name string, args *DnsDomainArgs, opts ...pulumi.ResourceOption) (*DnsDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	var resource DnsDomain
	err := ctx.RegisterResource("alicloud:dns/dnsDomain:DnsDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsDomain gets an existing DnsDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsDomainState, opts ...pulumi.ResourceOption) (*DnsDomain, error) {
	var resource DnsDomain
	err := ctx.ReadResource("alicloud:dns/dnsDomain:DnsDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsDomain resources.
type dnsDomainState struct {
	DnsServers []string `pulumi:"dnsServers"`
	// The domain ID.
	DomainId *string `pulumi:"domainId"`
	// Name of the domain. This name without suffix can have a string of 1 to 63 characters(domain name subject, excluding suffix), must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	DomainName *string `pulumi:"domainName"`
	// Id of the group in which the domain will add. If not supplied, then use default group.
	GroupId   *string `pulumi:"groupId"`
	GroupName *string `pulumi:"groupName"`
	// User language.
	Lang     *string `pulumi:"lang"`
	PunyCode *string `pulumi:"punyCode"`
	// Remarks information for your domain name.
	Remark *string `pulumi:"remark"`
	// The Id of resource group which the dns domain belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
}

type DnsDomainState struct {
	DnsServers pulumi.StringArrayInput
	// The domain ID.
	DomainId pulumi.StringPtrInput
	// Name of the domain. This name without suffix can have a string of 1 to 63 characters(domain name subject, excluding suffix), must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	DomainName pulumi.StringPtrInput
	// Id of the group in which the domain will add. If not supplied, then use default group.
	GroupId   pulumi.StringPtrInput
	GroupName pulumi.StringPtrInput
	// User language.
	Lang     pulumi.StringPtrInput
	PunyCode pulumi.StringPtrInput
	// Remarks information for your domain name.
	Remark pulumi.StringPtrInput
	// The Id of resource group which the dns domain belongs.
	ResourceGroupId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapInput
}

func (DnsDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsDomainState)(nil)).Elem()
}

type dnsDomainArgs struct {
	// Name of the domain. This name without suffix can have a string of 1 to 63 characters(domain name subject, excluding suffix), must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	DomainName string `pulumi:"domainName"`
	// Id of the group in which the domain will add. If not supplied, then use default group.
	GroupId *string `pulumi:"groupId"`
	// User language.
	Lang *string `pulumi:"lang"`
	// Remarks information for your domain name.
	Remark *string `pulumi:"remark"`
	// The Id of resource group which the dns domain belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a DnsDomain resource.
type DnsDomainArgs struct {
	// Name of the domain. This name without suffix can have a string of 1 to 63 characters(domain name subject, excluding suffix), must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	DomainName pulumi.StringInput
	// Id of the group in which the domain will add. If not supplied, then use default group.
	GroupId pulumi.StringPtrInput
	// User language.
	Lang pulumi.StringPtrInput
	// Remarks information for your domain name.
	Remark pulumi.StringPtrInput
	// The Id of resource group which the dns domain belongs.
	ResourceGroupId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapInput
}

func (DnsDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsDomainArgs)(nil)).Elem()
}

type DnsDomainInput interface {
	pulumi.Input

	ToDnsDomainOutput() DnsDomainOutput
	ToDnsDomainOutputWithContext(ctx context.Context) DnsDomainOutput
}

func (*DnsDomain) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsDomain)(nil))
}

func (i *DnsDomain) ToDnsDomainOutput() DnsDomainOutput {
	return i.ToDnsDomainOutputWithContext(context.Background())
}

func (i *DnsDomain) ToDnsDomainOutputWithContext(ctx context.Context) DnsDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsDomainOutput)
}

func (i *DnsDomain) ToDnsDomainPtrOutput() DnsDomainPtrOutput {
	return i.ToDnsDomainPtrOutputWithContext(context.Background())
}

func (i *DnsDomain) ToDnsDomainPtrOutputWithContext(ctx context.Context) DnsDomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsDomainPtrOutput)
}

type DnsDomainPtrInput interface {
	pulumi.Input

	ToDnsDomainPtrOutput() DnsDomainPtrOutput
	ToDnsDomainPtrOutputWithContext(ctx context.Context) DnsDomainPtrOutput
}

type dnsDomainPtrType DnsDomainArgs

func (*dnsDomainPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsDomain)(nil))
}

func (i *dnsDomainPtrType) ToDnsDomainPtrOutput() DnsDomainPtrOutput {
	return i.ToDnsDomainPtrOutputWithContext(context.Background())
}

func (i *dnsDomainPtrType) ToDnsDomainPtrOutputWithContext(ctx context.Context) DnsDomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsDomainPtrOutput)
}

// DnsDomainArrayInput is an input type that accepts DnsDomainArray and DnsDomainArrayOutput values.
// You can construct a concrete instance of `DnsDomainArrayInput` via:
//
//          DnsDomainArray{ DnsDomainArgs{...} }
type DnsDomainArrayInput interface {
	pulumi.Input

	ToDnsDomainArrayOutput() DnsDomainArrayOutput
	ToDnsDomainArrayOutputWithContext(context.Context) DnsDomainArrayOutput
}

type DnsDomainArray []DnsDomainInput

func (DnsDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DnsDomain)(nil))
}

func (i DnsDomainArray) ToDnsDomainArrayOutput() DnsDomainArrayOutput {
	return i.ToDnsDomainArrayOutputWithContext(context.Background())
}

func (i DnsDomainArray) ToDnsDomainArrayOutputWithContext(ctx context.Context) DnsDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsDomainArrayOutput)
}

// DnsDomainMapInput is an input type that accepts DnsDomainMap and DnsDomainMapOutput values.
// You can construct a concrete instance of `DnsDomainMapInput` via:
//
//          DnsDomainMap{ "key": DnsDomainArgs{...} }
type DnsDomainMapInput interface {
	pulumi.Input

	ToDnsDomainMapOutput() DnsDomainMapOutput
	ToDnsDomainMapOutputWithContext(context.Context) DnsDomainMapOutput
}

type DnsDomainMap map[string]DnsDomainInput

func (DnsDomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DnsDomain)(nil))
}

func (i DnsDomainMap) ToDnsDomainMapOutput() DnsDomainMapOutput {
	return i.ToDnsDomainMapOutputWithContext(context.Background())
}

func (i DnsDomainMap) ToDnsDomainMapOutputWithContext(ctx context.Context) DnsDomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsDomainMapOutput)
}

type DnsDomainOutput struct {
	*pulumi.OutputState
}

func (DnsDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsDomain)(nil))
}

func (o DnsDomainOutput) ToDnsDomainOutput() DnsDomainOutput {
	return o
}

func (o DnsDomainOutput) ToDnsDomainOutputWithContext(ctx context.Context) DnsDomainOutput {
	return o
}

func (o DnsDomainOutput) ToDnsDomainPtrOutput() DnsDomainPtrOutput {
	return o.ToDnsDomainPtrOutputWithContext(context.Background())
}

func (o DnsDomainOutput) ToDnsDomainPtrOutputWithContext(ctx context.Context) DnsDomainPtrOutput {
	return o.ApplyT(func(v DnsDomain) *DnsDomain {
		return &v
	}).(DnsDomainPtrOutput)
}

type DnsDomainPtrOutput struct {
	*pulumi.OutputState
}

func (DnsDomainPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsDomain)(nil))
}

func (o DnsDomainPtrOutput) ToDnsDomainPtrOutput() DnsDomainPtrOutput {
	return o
}

func (o DnsDomainPtrOutput) ToDnsDomainPtrOutputWithContext(ctx context.Context) DnsDomainPtrOutput {
	return o
}

type DnsDomainArrayOutput struct{ *pulumi.OutputState }

func (DnsDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DnsDomain)(nil))
}

func (o DnsDomainArrayOutput) ToDnsDomainArrayOutput() DnsDomainArrayOutput {
	return o
}

func (o DnsDomainArrayOutput) ToDnsDomainArrayOutputWithContext(ctx context.Context) DnsDomainArrayOutput {
	return o
}

func (o DnsDomainArrayOutput) Index(i pulumi.IntInput) DnsDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DnsDomain {
		return vs[0].([]DnsDomain)[vs[1].(int)]
	}).(DnsDomainOutput)
}

type DnsDomainMapOutput struct{ *pulumi.OutputState }

func (DnsDomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DnsDomain)(nil))
}

func (o DnsDomainMapOutput) ToDnsDomainMapOutput() DnsDomainMapOutput {
	return o
}

func (o DnsDomainMapOutput) ToDnsDomainMapOutputWithContext(ctx context.Context) DnsDomainMapOutput {
	return o
}

func (o DnsDomainMapOutput) MapIndex(k pulumi.StringInput) DnsDomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DnsDomain {
		return vs[0].(map[string]DnsDomain)[vs[1].(string)]
	}).(DnsDomainOutput)
}

func init() {
	pulumi.RegisterOutputType(DnsDomainOutput{})
	pulumi.RegisterOutputType(DnsDomainPtrOutput{})
	pulumi.RegisterOutputType(DnsDomainArrayOutput{})
	pulumi.RegisterOutputType(DnsDomainMapOutput{})
}
