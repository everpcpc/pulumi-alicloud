// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vod

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VOD Domain resource.
//
// For information about VOD Domain and how to use it, see [What is Domain](https://www.alibabacloud.com/help/product/29932.html).
//
// > **NOTE:** Available in v1.136.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vod"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vod.NewDomain(ctx, "_default", &vod.DomainArgs{
// 			DomainName: pulumi.String("your_domain_name"),
// 			Scope:      pulumi.String("domestic"),
// 			Sources: vod.DomainSourceArray{
// 				&vod.DomainSourceArgs{
// 					SourceContent: pulumi.String("your_source_content"),
// 					SourcePort:    pulumi.String("80"),
// 					SourceType:    pulumi.String("domain"),
// 				},
// 			},
// 			Tags: pulumi.AnyMap{
// 				"key1": pulumi.Any("value1"),
// 				"key2": pulumi.Any("value2"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// VOD Domain can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:vod/domain:Domain example <domain_name>
// ```
type Domain struct {
	pulumi.CustomResourceState

	// The name of the certificate. The value of this parameter is returned if HTTPS is enabled.
	CertName pulumi.StringOutput `pulumi:"certName"`
	// The URL that is used for health checks.
	CheckUrl pulumi.StringPtrOutput `pulumi:"checkUrl"`
	// The CNAME that is assigned to the domain name for CDN. You must add a CNAME record in the system of your Domain Name System (DNS) service provider to map the domain name for CDN to the CNAME.
	Cname pulumi.StringOutput `pulumi:"cname"`
	// The description of the domain name for CDN.
	Description pulumi.StringOutput `pulumi:"description"`
	// The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The time when the domain name for CDN was added. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	GmtCreated pulumi.StringOutput `pulumi:"gmtCreated"`
	// The last time when the domain name for CDN was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	GmtModified pulumi.StringOutput `pulumi:"gmtModified"`
	// This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
	Sources DomainSourceArrayOutput `pulumi:"sources"`
	// Indicates whether the Secure Sockets Layer (SSL) certificate is enabled. Valid values: `on`,`off`.
	SslProtocol pulumi.StringOutput `pulumi:"sslProtocol"`
	// The public key of the certificate. The value of this parameter is returned if HTTPS is enabled.
	SslPub pulumi.StringOutput `pulumi:"sslPub"`
	// The status of the domain name for CDN. Value values:
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	// * `Key`: It can be up to 64 characters in length. It cannot be a null string.
	// * `Value`: It can be up to 128 characters in length. It can be a null string.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The top-level domain name.
	TopLevelDomain pulumi.StringPtrOutput `pulumi:"topLevelDomain"`
	// The weight of the origin server.
	Weight pulumi.StringOutput `pulumi:"weight"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.Sources == nil {
		return nil, errors.New("invalid value for required argument 'Sources'")
	}
	var resource Domain
	err := ctx.RegisterResource("alicloud:vod/domain:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("alicloud:vod/domain:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
	// The name of the certificate. The value of this parameter is returned if HTTPS is enabled.
	CertName *string `pulumi:"certName"`
	// The URL that is used for health checks.
	CheckUrl *string `pulumi:"checkUrl"`
	// The CNAME that is assigned to the domain name for CDN. You must add a CNAME record in the system of your Domain Name System (DNS) service provider to map the domain name for CDN to the CNAME.
	Cname *string `pulumi:"cname"`
	// The description of the domain name for CDN.
	Description *string `pulumi:"description"`
	// The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
	DomainName *string `pulumi:"domainName"`
	// The time when the domain name for CDN was added. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	GmtCreated *string `pulumi:"gmtCreated"`
	// The last time when the domain name for CDN was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	GmtModified *string `pulumi:"gmtModified"`
	// This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
	Scope *string `pulumi:"scope"`
	// The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
	Sources []DomainSource `pulumi:"sources"`
	// Indicates whether the Secure Sockets Layer (SSL) certificate is enabled. Valid values: `on`,`off`.
	SslProtocol *string `pulumi:"sslProtocol"`
	// The public key of the certificate. The value of this parameter is returned if HTTPS is enabled.
	SslPub *string `pulumi:"sslPub"`
	// The status of the domain name for CDN. Value values:
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	// * `Key`: It can be up to 64 characters in length. It cannot be a null string.
	// * `Value`: It can be up to 128 characters in length. It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
	// The top-level domain name.
	TopLevelDomain *string `pulumi:"topLevelDomain"`
	// The weight of the origin server.
	Weight *string `pulumi:"weight"`
}

type DomainState struct {
	// The name of the certificate. The value of this parameter is returned if HTTPS is enabled.
	CertName pulumi.StringPtrInput
	// The URL that is used for health checks.
	CheckUrl pulumi.StringPtrInput
	// The CNAME that is assigned to the domain name for CDN. You must add a CNAME record in the system of your Domain Name System (DNS) service provider to map the domain name for CDN to the CNAME.
	Cname pulumi.StringPtrInput
	// The description of the domain name for CDN.
	Description pulumi.StringPtrInput
	// The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
	DomainName pulumi.StringPtrInput
	// The time when the domain name for CDN was added. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	GmtCreated pulumi.StringPtrInput
	// The last time when the domain name for CDN was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	GmtModified pulumi.StringPtrInput
	// This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
	Scope pulumi.StringPtrInput
	// The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
	Sources DomainSourceArrayInput
	// Indicates whether the Secure Sockets Layer (SSL) certificate is enabled. Valid values: `on`,`off`.
	SslProtocol pulumi.StringPtrInput
	// The public key of the certificate. The value of this parameter is returned if HTTPS is enabled.
	SslPub pulumi.StringPtrInput
	// The status of the domain name for CDN. Value values:
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	// * `Key`: It can be up to 64 characters in length. It cannot be a null string.
	// * `Value`: It can be up to 128 characters in length. It can be a null string.
	Tags pulumi.MapInput
	// The top-level domain name.
	TopLevelDomain pulumi.StringPtrInput
	// The weight of the origin server.
	Weight pulumi.StringPtrInput
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// The URL that is used for health checks.
	CheckUrl *string `pulumi:"checkUrl"`
	// The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
	DomainName string `pulumi:"domainName"`
	// This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
	Scope *string `pulumi:"scope"`
	// The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
	Sources []DomainSource `pulumi:"sources"`
	// A mapping of tags to assign to the resource.
	// * `Key`: It can be up to 64 characters in length. It cannot be a null string.
	// * `Value`: It can be up to 128 characters in length. It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
	// The top-level domain name.
	TopLevelDomain *string `pulumi:"topLevelDomain"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// The URL that is used for health checks.
	CheckUrl pulumi.StringPtrInput
	// The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
	DomainName pulumi.StringInput
	// This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
	Scope pulumi.StringPtrInput
	// The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
	Sources DomainSourceArrayInput
	// A mapping of tags to assign to the resource.
	// * `Key`: It can be up to 64 characters in length. It cannot be a null string.
	// * `Value`: It can be up to 128 characters in length. It can be a null string.
	Tags pulumi.MapInput
	// The top-level domain name.
	TopLevelDomain pulumi.StringPtrInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (*Domain) ElementType() reflect.Type {
	return reflect.TypeOf((*Domain)(nil))
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

func (i *Domain) ToDomainPtrOutput() DomainPtrOutput {
	return i.ToDomainPtrOutputWithContext(context.Background())
}

func (i *Domain) ToDomainPtrOutputWithContext(ctx context.Context) DomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPtrOutput)
}

type DomainPtrInput interface {
	pulumi.Input

	ToDomainPtrOutput() DomainPtrOutput
	ToDomainPtrOutputWithContext(ctx context.Context) DomainPtrOutput
}

type domainPtrType DomainArgs

func (*domainPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil))
}

func (i *domainPtrType) ToDomainPtrOutput() DomainPtrOutput {
	return i.ToDomainPtrOutputWithContext(context.Background())
}

func (i *domainPtrType) ToDomainPtrOutputWithContext(ctx context.Context) DomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPtrOutput)
}

// DomainArrayInput is an input type that accepts DomainArray and DomainArrayOutput values.
// You can construct a concrete instance of `DomainArrayInput` via:
//
//          DomainArray{ DomainArgs{...} }
type DomainArrayInput interface {
	pulumi.Input

	ToDomainArrayOutput() DomainArrayOutput
	ToDomainArrayOutputWithContext(context.Context) DomainArrayOutput
}

type DomainArray []DomainInput

func (DomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Domain)(nil)).Elem()
}

func (i DomainArray) ToDomainArrayOutput() DomainArrayOutput {
	return i.ToDomainArrayOutputWithContext(context.Background())
}

func (i DomainArray) ToDomainArrayOutputWithContext(ctx context.Context) DomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainArrayOutput)
}

// DomainMapInput is an input type that accepts DomainMap and DomainMapOutput values.
// You can construct a concrete instance of `DomainMapInput` via:
//
//          DomainMap{ "key": DomainArgs{...} }
type DomainMapInput interface {
	pulumi.Input

	ToDomainMapOutput() DomainMapOutput
	ToDomainMapOutputWithContext(context.Context) DomainMapOutput
}

type DomainMap map[string]DomainInput

func (DomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Domain)(nil)).Elem()
}

func (i DomainMap) ToDomainMapOutput() DomainMapOutput {
	return i.ToDomainMapOutputWithContext(context.Background())
}

func (i DomainMap) ToDomainMapOutputWithContext(ctx context.Context) DomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainMapOutput)
}

type DomainOutput struct{ *pulumi.OutputState }

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Domain)(nil))
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

func (o DomainOutput) ToDomainPtrOutput() DomainPtrOutput {
	return o.ToDomainPtrOutputWithContext(context.Background())
}

func (o DomainOutput) ToDomainPtrOutputWithContext(ctx context.Context) DomainPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Domain) *Domain {
		return &v
	}).(DomainPtrOutput)
}

type DomainPtrOutput struct{ *pulumi.OutputState }

func (DomainPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil))
}

func (o DomainPtrOutput) ToDomainPtrOutput() DomainPtrOutput {
	return o
}

func (o DomainPtrOutput) ToDomainPtrOutputWithContext(ctx context.Context) DomainPtrOutput {
	return o
}

func (o DomainPtrOutput) Elem() DomainOutput {
	return o.ApplyT(func(v *Domain) Domain {
		if v != nil {
			return *v
		}
		var ret Domain
		return ret
	}).(DomainOutput)
}

type DomainArrayOutput struct{ *pulumi.OutputState }

func (DomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Domain)(nil))
}

func (o DomainArrayOutput) ToDomainArrayOutput() DomainArrayOutput {
	return o
}

func (o DomainArrayOutput) ToDomainArrayOutputWithContext(ctx context.Context) DomainArrayOutput {
	return o
}

func (o DomainArrayOutput) Index(i pulumi.IntInput) DomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Domain {
		return vs[0].([]Domain)[vs[1].(int)]
	}).(DomainOutput)
}

type DomainMapOutput struct{ *pulumi.OutputState }

func (DomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Domain)(nil))
}

func (o DomainMapOutput) ToDomainMapOutput() DomainMapOutput {
	return o
}

func (o DomainMapOutput) ToDomainMapOutputWithContext(ctx context.Context) DomainMapOutput {
	return o
}

func (o DomainMapOutput) MapIndex(k pulumi.StringInput) DomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Domain {
		return vs[0].(map[string]Domain)[vs[1].(string)]
	}).(DomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainInput)(nil)).Elem(), &Domain{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainPtrInput)(nil)).Elem(), &Domain{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainArrayInput)(nil)).Elem(), DomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainMapInput)(nil)).Elem(), DomainMap{})
	pulumi.RegisterOutputType(DomainOutput{})
	pulumi.RegisterOutputType(DomainPtrOutput{})
	pulumi.RegisterOutputType(DomainArrayOutput{})
	pulumi.RegisterOutputType(DomainMapOutput{})
}
