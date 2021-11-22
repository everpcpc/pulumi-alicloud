// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Load Balancer CA Certificate is used by the listener of the protocol https.
//
// For information about slb and how to use it, see [What is Server Load Balancer](https://www.alibabacloud.com/help/doc-detail/27539.htm).
//
// For information about CA Certificate and how to use it, see [Configure CA Certificate](https://www.alibabacloud.com/help/doc-detail/85968.htm).
//
// ## Example Usage
//
// * using CA certificate content
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/slb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := slb.NewCaCertificate(ctx, "foo", &slb.CaCertificateArgs{
// 			CaCertificate:     pulumi.String(fmt.Sprintf("%v%v%v", "-----BEGIN CERTIFICATE-----\n", "MIIDRjCCAq+gAwIBAgIJAJnI******90EAxEG/bJJyOm5LqoiA=\n", "-----END CERTIFICATE-----\n")),
// 			CaCertificateName: pulumi.String("tf-testAccSlbCACertificate"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// * using CA certificate file
//
// ```go
// package main
//
// import (
// 	"fmt"
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/slb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := slb.NewCaCertificate(ctx, "foo_file", &slb.CaCertificateArgs{
// 			CaCertificateName: pulumi.String("tf-testAccSlbCACertificate"),
// 			CaCertificate:     readFileOrPanic(fmt.Sprintf("%v%v", path.Module, "/ca_certificate.pem")),
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
// Server Load balancer CA Certificate can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:slb/caCertificate:CaCertificate example abc123456
// ```
type CaCertificate struct {
	pulumi.CustomResourceState

	// the content of the CA certificate.
	CaCertificate     pulumi.StringOutput `pulumi:"caCertificate"`
	CaCertificateName pulumi.StringOutput `pulumi:"caCertificateName"`
	// Field `name` has been deprecated from provider version 1.123.1. New field `caCertificateName` instead
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'ca_certificate_name' instead
	Name pulumi.StringOutput `pulumi:"name"`
	// The Id of resource group which the slbCa certificate belongs.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewCaCertificate registers a new resource with the given unique name, arguments, and options.
func NewCaCertificate(ctx *pulumi.Context,
	name string, args *CaCertificateArgs, opts ...pulumi.ResourceOption) (*CaCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CaCertificate == nil {
		return nil, errors.New("invalid value for required argument 'CaCertificate'")
	}
	var resource CaCertificate
	err := ctx.RegisterResource("alicloud:slb/caCertificate:CaCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCaCertificate gets an existing CaCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCaCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CaCertificateState, opts ...pulumi.ResourceOption) (*CaCertificate, error) {
	var resource CaCertificate
	err := ctx.ReadResource("alicloud:slb/caCertificate:CaCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CaCertificate resources.
type caCertificateState struct {
	// the content of the CA certificate.
	CaCertificate     *string `pulumi:"caCertificate"`
	CaCertificateName *string `pulumi:"caCertificateName"`
	// Field `name` has been deprecated from provider version 1.123.1. New field `caCertificateName` instead
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'ca_certificate_name' instead
	Name *string `pulumi:"name"`
	// The Id of resource group which the slbCa certificate belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type CaCertificateState struct {
	// the content of the CA certificate.
	CaCertificate     pulumi.StringPtrInput
	CaCertificateName pulumi.StringPtrInput
	// Field `name` has been deprecated from provider version 1.123.1. New field `caCertificateName` instead
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'ca_certificate_name' instead
	Name pulumi.StringPtrInput
	// The Id of resource group which the slbCa certificate belongs.
	ResourceGroupId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (CaCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*caCertificateState)(nil)).Elem()
}

type caCertificateArgs struct {
	// the content of the CA certificate.
	CaCertificate     string  `pulumi:"caCertificate"`
	CaCertificateName *string `pulumi:"caCertificateName"`
	// Field `name` has been deprecated from provider version 1.123.1. New field `caCertificateName` instead
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'ca_certificate_name' instead
	Name *string `pulumi:"name"`
	// The Id of resource group which the slbCa certificate belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a CaCertificate resource.
type CaCertificateArgs struct {
	// the content of the CA certificate.
	CaCertificate     pulumi.StringInput
	CaCertificateName pulumi.StringPtrInput
	// Field `name` has been deprecated from provider version 1.123.1. New field `caCertificateName` instead
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'ca_certificate_name' instead
	Name pulumi.StringPtrInput
	// The Id of resource group which the slbCa certificate belongs.
	ResourceGroupId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (CaCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*caCertificateArgs)(nil)).Elem()
}

type CaCertificateInput interface {
	pulumi.Input

	ToCaCertificateOutput() CaCertificateOutput
	ToCaCertificateOutputWithContext(ctx context.Context) CaCertificateOutput
}

func (*CaCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((*CaCertificate)(nil))
}

func (i *CaCertificate) ToCaCertificateOutput() CaCertificateOutput {
	return i.ToCaCertificateOutputWithContext(context.Background())
}

func (i *CaCertificate) ToCaCertificateOutputWithContext(ctx context.Context) CaCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaCertificateOutput)
}

func (i *CaCertificate) ToCaCertificatePtrOutput() CaCertificatePtrOutput {
	return i.ToCaCertificatePtrOutputWithContext(context.Background())
}

func (i *CaCertificate) ToCaCertificatePtrOutputWithContext(ctx context.Context) CaCertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaCertificatePtrOutput)
}

type CaCertificatePtrInput interface {
	pulumi.Input

	ToCaCertificatePtrOutput() CaCertificatePtrOutput
	ToCaCertificatePtrOutputWithContext(ctx context.Context) CaCertificatePtrOutput
}

type caCertificatePtrType CaCertificateArgs

func (*caCertificatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CaCertificate)(nil))
}

func (i *caCertificatePtrType) ToCaCertificatePtrOutput() CaCertificatePtrOutput {
	return i.ToCaCertificatePtrOutputWithContext(context.Background())
}

func (i *caCertificatePtrType) ToCaCertificatePtrOutputWithContext(ctx context.Context) CaCertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaCertificatePtrOutput)
}

// CaCertificateArrayInput is an input type that accepts CaCertificateArray and CaCertificateArrayOutput values.
// You can construct a concrete instance of `CaCertificateArrayInput` via:
//
//          CaCertificateArray{ CaCertificateArgs{...} }
type CaCertificateArrayInput interface {
	pulumi.Input

	ToCaCertificateArrayOutput() CaCertificateArrayOutput
	ToCaCertificateArrayOutputWithContext(context.Context) CaCertificateArrayOutput
}

type CaCertificateArray []CaCertificateInput

func (CaCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CaCertificate)(nil)).Elem()
}

func (i CaCertificateArray) ToCaCertificateArrayOutput() CaCertificateArrayOutput {
	return i.ToCaCertificateArrayOutputWithContext(context.Background())
}

func (i CaCertificateArray) ToCaCertificateArrayOutputWithContext(ctx context.Context) CaCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaCertificateArrayOutput)
}

// CaCertificateMapInput is an input type that accepts CaCertificateMap and CaCertificateMapOutput values.
// You can construct a concrete instance of `CaCertificateMapInput` via:
//
//          CaCertificateMap{ "key": CaCertificateArgs{...} }
type CaCertificateMapInput interface {
	pulumi.Input

	ToCaCertificateMapOutput() CaCertificateMapOutput
	ToCaCertificateMapOutputWithContext(context.Context) CaCertificateMapOutput
}

type CaCertificateMap map[string]CaCertificateInput

func (CaCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CaCertificate)(nil)).Elem()
}

func (i CaCertificateMap) ToCaCertificateMapOutput() CaCertificateMapOutput {
	return i.ToCaCertificateMapOutputWithContext(context.Background())
}

func (i CaCertificateMap) ToCaCertificateMapOutputWithContext(ctx context.Context) CaCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaCertificateMapOutput)
}

type CaCertificateOutput struct{ *pulumi.OutputState }

func (CaCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CaCertificate)(nil))
}

func (o CaCertificateOutput) ToCaCertificateOutput() CaCertificateOutput {
	return o
}

func (o CaCertificateOutput) ToCaCertificateOutputWithContext(ctx context.Context) CaCertificateOutput {
	return o
}

func (o CaCertificateOutput) ToCaCertificatePtrOutput() CaCertificatePtrOutput {
	return o.ToCaCertificatePtrOutputWithContext(context.Background())
}

func (o CaCertificateOutput) ToCaCertificatePtrOutputWithContext(ctx context.Context) CaCertificatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CaCertificate) *CaCertificate {
		return &v
	}).(CaCertificatePtrOutput)
}

type CaCertificatePtrOutput struct{ *pulumi.OutputState }

func (CaCertificatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CaCertificate)(nil))
}

func (o CaCertificatePtrOutput) ToCaCertificatePtrOutput() CaCertificatePtrOutput {
	return o
}

func (o CaCertificatePtrOutput) ToCaCertificatePtrOutputWithContext(ctx context.Context) CaCertificatePtrOutput {
	return o
}

func (o CaCertificatePtrOutput) Elem() CaCertificateOutput {
	return o.ApplyT(func(v *CaCertificate) CaCertificate {
		if v != nil {
			return *v
		}
		var ret CaCertificate
		return ret
	}).(CaCertificateOutput)
}

type CaCertificateArrayOutput struct{ *pulumi.OutputState }

func (CaCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CaCertificate)(nil))
}

func (o CaCertificateArrayOutput) ToCaCertificateArrayOutput() CaCertificateArrayOutput {
	return o
}

func (o CaCertificateArrayOutput) ToCaCertificateArrayOutputWithContext(ctx context.Context) CaCertificateArrayOutput {
	return o
}

func (o CaCertificateArrayOutput) Index(i pulumi.IntInput) CaCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CaCertificate {
		return vs[0].([]CaCertificate)[vs[1].(int)]
	}).(CaCertificateOutput)
}

type CaCertificateMapOutput struct{ *pulumi.OutputState }

func (CaCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CaCertificate)(nil))
}

func (o CaCertificateMapOutput) ToCaCertificateMapOutput() CaCertificateMapOutput {
	return o
}

func (o CaCertificateMapOutput) ToCaCertificateMapOutputWithContext(ctx context.Context) CaCertificateMapOutput {
	return o
}

func (o CaCertificateMapOutput) MapIndex(k pulumi.StringInput) CaCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CaCertificate {
		return vs[0].(map[string]CaCertificate)[vs[1].(string)]
	}).(CaCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CaCertificateInput)(nil)).Elem(), &CaCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*CaCertificatePtrInput)(nil)).Elem(), &CaCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*CaCertificateArrayInput)(nil)).Elem(), CaCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CaCertificateMapInput)(nil)).Elem(), CaCertificateMap{})
	pulumi.RegisterOutputType(CaCertificateOutput{})
	pulumi.RegisterOutputType(CaCertificatePtrOutput{})
	pulumi.RegisterOutputType(CaCertificateArrayOutput{})
	pulumi.RegisterOutputType(CaCertificateMapOutput{})
}
