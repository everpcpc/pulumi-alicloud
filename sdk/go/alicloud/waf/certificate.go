// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a WAF Certificate resource.
//
// For information about WAF Certificate and how to use it, see [What is Certificate](https://www.alibabacloud.com/help/doc-detail/28517.htm).
//
// > **NOTE:** Available in v1.135.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/waf"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := waf.NewCertificate(ctx, "_default", &waf.CertificateArgs{
// 			Certificate:     pulumi.String("your_certificate"),
// 			CertificateName: pulumi.String("your_certificate_name"),
// 			Domain:          pulumi.String("your_domain_name"),
// 			InstanceId:      pulumi.String("your_instance_id"),
// 			PrivateKey:      pulumi.String("your_private_key"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = waf.NewCertificate(ctx, "default2", &waf.CertificateArgs{
// 			CertificateId: pulumi.String("your_certificate_id"),
// 			Domain:        pulumi.String("your_domain_name"),
// 			InstanceId:    pulumi.String("your_instance_id"),
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
// WAF Certificate can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:waf/certificate:Certificate example <instance_id>:<domain>:<certificate_id>
// ```
type Certificate struct {
	pulumi.CustomResourceState

	// Certificate file content.
	Certificate pulumi.StringPtrOutput `pulumi:"certificate"`
	// The certificate id is automatically generated when you upload your certificate content.**NOTE:** you can also use Certificate ID saved in SSL.
	CertificateId pulumi.StringOutput `pulumi:"certificateId"`
	// Certificate file name.
	CertificateName pulumi.StringOutput `pulumi:"certificateName"`
	// The domain that you want to add to WAF.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The ID of the WAF instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The private key.
	PrivateKey pulumi.StringPtrOutput `pulumi:"privateKey"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource Certificate
	err := ctx.RegisterResource("alicloud:waf/certificate:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("alicloud:waf/certificate:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// Certificate file content.
	Certificate *string `pulumi:"certificate"`
	// The certificate id is automatically generated when you upload your certificate content.**NOTE:** you can also use Certificate ID saved in SSL.
	CertificateId *string `pulumi:"certificateId"`
	// Certificate file name.
	CertificateName *string `pulumi:"certificateName"`
	// The domain that you want to add to WAF.
	Domain *string `pulumi:"domain"`
	// The ID of the WAF instance.
	InstanceId *string `pulumi:"instanceId"`
	// The private key.
	PrivateKey *string `pulumi:"privateKey"`
}

type CertificateState struct {
	// Certificate file content.
	Certificate pulumi.StringPtrInput
	// The certificate id is automatically generated when you upload your certificate content.**NOTE:** you can also use Certificate ID saved in SSL.
	CertificateId pulumi.StringPtrInput
	// Certificate file name.
	CertificateName pulumi.StringPtrInput
	// The domain that you want to add to WAF.
	Domain pulumi.StringPtrInput
	// The ID of the WAF instance.
	InstanceId pulumi.StringPtrInput
	// The private key.
	PrivateKey pulumi.StringPtrInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// Certificate file content.
	Certificate *string `pulumi:"certificate"`
	// The certificate id is automatically generated when you upload your certificate content.**NOTE:** you can also use Certificate ID saved in SSL.
	CertificateId *string `pulumi:"certificateId"`
	// Certificate file name.
	CertificateName *string `pulumi:"certificateName"`
	// The domain that you want to add to WAF.
	Domain string `pulumi:"domain"`
	// The ID of the WAF instance.
	InstanceId string `pulumi:"instanceId"`
	// The private key.
	PrivateKey *string `pulumi:"privateKey"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// Certificate file content.
	Certificate pulumi.StringPtrInput
	// The certificate id is automatically generated when you upload your certificate content.**NOTE:** you can also use Certificate ID saved in SSL.
	CertificateId pulumi.StringPtrInput
	// Certificate file name.
	CertificateName pulumi.StringPtrInput
	// The domain that you want to add to WAF.
	Domain pulumi.StringInput
	// The ID of the WAF instance.
	InstanceId pulumi.StringInput
	// The private key.
	PrivateKey pulumi.StringPtrInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil))
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

func (i *Certificate) ToCertificatePtrOutput() CertificatePtrOutput {
	return i.ToCertificatePtrOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePtrOutput)
}

type CertificatePtrInput interface {
	pulumi.Input

	ToCertificatePtrOutput() CertificatePtrOutput
	ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput
}

type certificatePtrType CertificateArgs

func (*certificatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil))
}

func (i *certificatePtrType) ToCertificatePtrOutput() CertificatePtrOutput {
	return i.ToCertificatePtrOutputWithContext(context.Background())
}

func (i *certificatePtrType) ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePtrOutput)
}

// CertificateArrayInput is an input type that accepts CertificateArray and CertificateArrayOutput values.
// You can construct a concrete instance of `CertificateArrayInput` via:
//
//          CertificateArray{ CertificateArgs{...} }
type CertificateArrayInput interface {
	pulumi.Input

	ToCertificateArrayOutput() CertificateArrayOutput
	ToCertificateArrayOutputWithContext(context.Context) CertificateArrayOutput
}

type CertificateArray []CertificateInput

func (CertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Certificate)(nil)).Elem()
}

func (i CertificateArray) ToCertificateArrayOutput() CertificateArrayOutput {
	return i.ToCertificateArrayOutputWithContext(context.Background())
}

func (i CertificateArray) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateArrayOutput)
}

// CertificateMapInput is an input type that accepts CertificateMap and CertificateMapOutput values.
// You can construct a concrete instance of `CertificateMapInput` via:
//
//          CertificateMap{ "key": CertificateArgs{...} }
type CertificateMapInput interface {
	pulumi.Input

	ToCertificateMapOutput() CertificateMapOutput
	ToCertificateMapOutputWithContext(context.Context) CertificateMapOutput
}

type CertificateMap map[string]CertificateInput

func (CertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Certificate)(nil)).Elem()
}

func (i CertificateMap) ToCertificateMapOutput() CertificateMapOutput {
	return i.ToCertificateMapOutputWithContext(context.Background())
}

func (i CertificateMap) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateMapOutput)
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil))
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificatePtrOutput() CertificatePtrOutput {
	return o.ToCertificatePtrOutputWithContext(context.Background())
}

func (o CertificateOutput) ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Certificate) *Certificate {
		return &v
	}).(CertificatePtrOutput)
}

type CertificatePtrOutput struct{ *pulumi.OutputState }

func (CertificatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil))
}

func (o CertificatePtrOutput) ToCertificatePtrOutput() CertificatePtrOutput {
	return o
}

func (o CertificatePtrOutput) ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput {
	return o
}

func (o CertificatePtrOutput) Elem() CertificateOutput {
	return o.ApplyT(func(v *Certificate) Certificate {
		if v != nil {
			return *v
		}
		var ret Certificate
		return ret
	}).(CertificateOutput)
}

type CertificateArrayOutput struct{ *pulumi.OutputState }

func (CertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Certificate)(nil))
}

func (o CertificateArrayOutput) ToCertificateArrayOutput() CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) Index(i pulumi.IntInput) CertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Certificate {
		return vs[0].([]Certificate)[vs[1].(int)]
	}).(CertificateOutput)
}

type CertificateMapOutput struct{ *pulumi.OutputState }

func (CertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Certificate)(nil))
}

func (o CertificateMapOutput) ToCertificateMapOutput() CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) MapIndex(k pulumi.StringInput) CertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Certificate {
		return vs[0].(map[string]Certificate)[vs[1].(string)]
	}).(CertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateInput)(nil)).Elem(), &Certificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificatePtrInput)(nil)).Elem(), &Certificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateArrayInput)(nil)).Elem(), CertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateMapInput)(nil)).Elem(), CertificateMap{})
	pulumi.RegisterOutputType(CertificateOutput{})
	pulumi.RegisterOutputType(CertificatePtrOutput{})
	pulumi.RegisterOutputType(CertificateArrayOutput{})
	pulumi.RegisterOutputType(CertificateMapOutput{})
}
