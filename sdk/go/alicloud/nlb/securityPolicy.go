// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nlb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a NLB Security Policy resource.
//
// For information about NLB Security Policy and how to use it, see [What is Security Policy](https://www.alibabacloud.com/help/en/server-load-balancer/latest/createsecuritypolicy-nlb).
//
// > **NOTE:** Available in v1.187.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/nlb"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultResourceGroups, err := resourcemanager.GetResourceGroups(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = nlb.NewSecurityPolicy(ctx, "defaultSecurityPolicy", &nlb.SecurityPolicyArgs{
//				ResourceGroupId:    *pulumi.String(defaultResourceGroups.Ids[0]),
//				SecurityPolicyName: pulumi.Any(_var.Name),
//				Ciphers: pulumi.StringArray{
//					pulumi.String("ECDHE-RSA-AES128-SHA"),
//					pulumi.String("ECDHE-ECDSA-AES128-SHA"),
//				},
//				TlsVersions: pulumi.StringArray{
//					pulumi.String("TLSv1.0"),
//					pulumi.String("TLSv1.1"),
//					pulumi.String("TLSv1.2"),
//				},
//				Tags: pulumi.AnyMap{
//					"Created": pulumi.Any("TF"),
//					"For":     pulumi.Any("Acceptance-test"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// NLB Security Policy can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:nlb/securityPolicy:SecurityPolicy example <id>
//
// ```
type SecurityPolicy struct {
	pulumi.CustomResourceState

	// The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
	// - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
	// - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
	// - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`
	Ciphers pulumi.StringArrayOutput `pulumi:"ciphers"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The name of the security policy. The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	SecurityPolicyName pulumi.StringOutput `pulumi:"securityPolicyName"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
	TlsVersions pulumi.StringArrayOutput `pulumi:"tlsVersions"`
}

// NewSecurityPolicy registers a new resource with the given unique name, arguments, and options.
func NewSecurityPolicy(ctx *pulumi.Context,
	name string, args *SecurityPolicyArgs, opts ...pulumi.ResourceOption) (*SecurityPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ciphers == nil {
		return nil, errors.New("invalid value for required argument 'Ciphers'")
	}
	if args.TlsVersions == nil {
		return nil, errors.New("invalid value for required argument 'TlsVersions'")
	}
	var resource SecurityPolicy
	err := ctx.RegisterResource("alicloud:nlb/securityPolicy:SecurityPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityPolicy gets an existing SecurityPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityPolicyState, opts ...pulumi.ResourceOption) (*SecurityPolicy, error) {
	var resource SecurityPolicy
	err := ctx.ReadResource("alicloud:nlb/securityPolicy:SecurityPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityPolicy resources.
type securityPolicyState struct {
	// The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
	// - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
	// - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
	// - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`
	Ciphers []string `pulumi:"ciphers"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The name of the security policy. The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	SecurityPolicyName *string `pulumi:"securityPolicyName"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
	TlsVersions []string `pulumi:"tlsVersions"`
}

type SecurityPolicyState struct {
	// The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
	// - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
	// - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
	// - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`
	Ciphers pulumi.StringArrayInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The name of the security policy. The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	SecurityPolicyName pulumi.StringPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
	TlsVersions pulumi.StringArrayInput
}

func (SecurityPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityPolicyState)(nil)).Elem()
}

type securityPolicyArgs struct {
	// The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
	// - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
	// - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
	// - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`
	Ciphers []string `pulumi:"ciphers"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The name of the security policy. The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	SecurityPolicyName *string `pulumi:"securityPolicyName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
	TlsVersions []string `pulumi:"tlsVersions"`
}

// The set of arguments for constructing a SecurityPolicy resource.
type SecurityPolicyArgs struct {
	// The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
	// - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
	// - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
	// - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`
	Ciphers pulumi.StringArrayInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The name of the security policy. The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	SecurityPolicyName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
	TlsVersions pulumi.StringArrayInput
}

func (SecurityPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityPolicyArgs)(nil)).Elem()
}

type SecurityPolicyInput interface {
	pulumi.Input

	ToSecurityPolicyOutput() SecurityPolicyOutput
	ToSecurityPolicyOutputWithContext(ctx context.Context) SecurityPolicyOutput
}

func (*SecurityPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityPolicy)(nil)).Elem()
}

func (i *SecurityPolicy) ToSecurityPolicyOutput() SecurityPolicyOutput {
	return i.ToSecurityPolicyOutputWithContext(context.Background())
}

func (i *SecurityPolicy) ToSecurityPolicyOutputWithContext(ctx context.Context) SecurityPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityPolicyOutput)
}

// SecurityPolicyArrayInput is an input type that accepts SecurityPolicyArray and SecurityPolicyArrayOutput values.
// You can construct a concrete instance of `SecurityPolicyArrayInput` via:
//
//	SecurityPolicyArray{ SecurityPolicyArgs{...} }
type SecurityPolicyArrayInput interface {
	pulumi.Input

	ToSecurityPolicyArrayOutput() SecurityPolicyArrayOutput
	ToSecurityPolicyArrayOutputWithContext(context.Context) SecurityPolicyArrayOutput
}

type SecurityPolicyArray []SecurityPolicyInput

func (SecurityPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityPolicy)(nil)).Elem()
}

func (i SecurityPolicyArray) ToSecurityPolicyArrayOutput() SecurityPolicyArrayOutput {
	return i.ToSecurityPolicyArrayOutputWithContext(context.Background())
}

func (i SecurityPolicyArray) ToSecurityPolicyArrayOutputWithContext(ctx context.Context) SecurityPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityPolicyArrayOutput)
}

// SecurityPolicyMapInput is an input type that accepts SecurityPolicyMap and SecurityPolicyMapOutput values.
// You can construct a concrete instance of `SecurityPolicyMapInput` via:
//
//	SecurityPolicyMap{ "key": SecurityPolicyArgs{...} }
type SecurityPolicyMapInput interface {
	pulumi.Input

	ToSecurityPolicyMapOutput() SecurityPolicyMapOutput
	ToSecurityPolicyMapOutputWithContext(context.Context) SecurityPolicyMapOutput
}

type SecurityPolicyMap map[string]SecurityPolicyInput

func (SecurityPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityPolicy)(nil)).Elem()
}

func (i SecurityPolicyMap) ToSecurityPolicyMapOutput() SecurityPolicyMapOutput {
	return i.ToSecurityPolicyMapOutputWithContext(context.Background())
}

func (i SecurityPolicyMap) ToSecurityPolicyMapOutputWithContext(ctx context.Context) SecurityPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityPolicyMapOutput)
}

type SecurityPolicyOutput struct{ *pulumi.OutputState }

func (SecurityPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityPolicy)(nil)).Elem()
}

func (o SecurityPolicyOutput) ToSecurityPolicyOutput() SecurityPolicyOutput {
	return o
}

func (o SecurityPolicyOutput) ToSecurityPolicyOutputWithContext(ctx context.Context) SecurityPolicyOutput {
	return o
}

// The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
// - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
// - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
// - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`
func (o SecurityPolicyOutput) Ciphers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityPolicy) pulumi.StringArrayOutput { return v.Ciphers }).(pulumi.StringArrayOutput)
}

// The ID of the resource group.
func (o SecurityPolicyOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityPolicy) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The name of the security policy. The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
func (o SecurityPolicyOutput) SecurityPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityPolicy) pulumi.StringOutput { return v.SecurityPolicyName }).(pulumi.StringOutput)
}

// The status of the resource.
func (o SecurityPolicyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityPolicy) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o SecurityPolicyOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *SecurityPolicy) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
func (o SecurityPolicyOutput) TlsVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityPolicy) pulumi.StringArrayOutput { return v.TlsVersions }).(pulumi.StringArrayOutput)
}

type SecurityPolicyArrayOutput struct{ *pulumi.OutputState }

func (SecurityPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityPolicy)(nil)).Elem()
}

func (o SecurityPolicyArrayOutput) ToSecurityPolicyArrayOutput() SecurityPolicyArrayOutput {
	return o
}

func (o SecurityPolicyArrayOutput) ToSecurityPolicyArrayOutputWithContext(ctx context.Context) SecurityPolicyArrayOutput {
	return o
}

func (o SecurityPolicyArrayOutput) Index(i pulumi.IntInput) SecurityPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecurityPolicy {
		return vs[0].([]*SecurityPolicy)[vs[1].(int)]
	}).(SecurityPolicyOutput)
}

type SecurityPolicyMapOutput struct{ *pulumi.OutputState }

func (SecurityPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityPolicy)(nil)).Elem()
}

func (o SecurityPolicyMapOutput) ToSecurityPolicyMapOutput() SecurityPolicyMapOutput {
	return o
}

func (o SecurityPolicyMapOutput) ToSecurityPolicyMapOutputWithContext(ctx context.Context) SecurityPolicyMapOutput {
	return o
}

func (o SecurityPolicyMapOutput) MapIndex(k pulumi.StringInput) SecurityPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecurityPolicy {
		return vs[0].(map[string]*SecurityPolicy)[vs[1].(string)]
	}).(SecurityPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityPolicyInput)(nil)).Elem(), &SecurityPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityPolicyArrayInput)(nil)).Elem(), SecurityPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityPolicyMapInput)(nil)).Elem(), SecurityPolicyMap{})
	pulumi.RegisterOutputType(SecurityPolicyOutput{})
	pulumi.RegisterOutputType(SecurityPolicyArrayOutput{})
	pulumi.RegisterOutputType(SecurityPolicyMapOutput{})
}