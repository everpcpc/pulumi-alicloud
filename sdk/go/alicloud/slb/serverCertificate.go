// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Load Balancer Server Certificate is an ssl Certificate used by the listener of the protocol https.
//
// For information about slb and how to use it, see [What is Server Load Balancer](https://www.alibabacloud.com/help/doc-detail/27539.htm).
//
// For information about Server Certificate and how to use it, see [Configure Server Certificate](https://www.alibabacloud.com/help/doc-detail/85968.htm).
//
// ## Import
//
// Server Load balancer Server Certificate can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:slb/serverCertificate:ServerCertificate example abc123456
// ```
type ServerCertificate struct {
	pulumi.CustomResourceState

	// Deprecated: Field 'alicloud_certifacte_id' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_id' replaces it.
	AlicloudCertifacteId pulumi.StringPtrOutput `pulumi:"alicloudCertifacteId"`
	// Deprecated: Field 'alicloud_certifacte_name' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_name' replaces it.
	AlicloudCertifacteName pulumi.StringPtrOutput `pulumi:"alicloudCertifacteName"`
	// an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateId pulumi.StringPtrOutput `pulumi:"alicloudCertificateId"`
	// the name of the certificate specified by `alicloudCertificateId`.but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateName pulumi.StringPtrOutput `pulumi:"alicloudCertificateName"`
	// the region of the certificate specified by `alicloudCertificateId`. but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateRegionId pulumi.StringPtrOutput `pulumi:"alicloudCertificateRegionId"`
	// Name of the Server Certificate.
	Name pulumi.StringOutput `pulumi:"name"`
	// the content of privat key of the ssl certificate specified by `serverCertificate`. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
	PrivateKey pulumi.StringPtrOutput `pulumi:"privateKey"`
	// The Id of resource group which the slb server certificate belongs.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// the content of the ssl certificate. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
	ServerCertificate pulumi.StringPtrOutput `pulumi:"serverCertificate"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewServerCertificate registers a new resource with the given unique name, arguments, and options.
func NewServerCertificate(ctx *pulumi.Context,
	name string, args *ServerCertificateArgs, opts ...pulumi.ResourceOption) (*ServerCertificate, error) {
	if args == nil {
		args = &ServerCertificateArgs{}
	}

	var resource ServerCertificate
	err := ctx.RegisterResource("alicloud:slb/serverCertificate:ServerCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerCertificate gets an existing ServerCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerCertificateState, opts ...pulumi.ResourceOption) (*ServerCertificate, error) {
	var resource ServerCertificate
	err := ctx.ReadResource("alicloud:slb/serverCertificate:ServerCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerCertificate resources.
type serverCertificateState struct {
	// Deprecated: Field 'alicloud_certifacte_id' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_id' replaces it.
	AlicloudCertifacteId *string `pulumi:"alicloudCertifacteId"`
	// Deprecated: Field 'alicloud_certifacte_name' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_name' replaces it.
	AlicloudCertifacteName *string `pulumi:"alicloudCertifacteName"`
	// an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateId *string `pulumi:"alicloudCertificateId"`
	// the name of the certificate specified by `alicloudCertificateId`.but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateName *string `pulumi:"alicloudCertificateName"`
	// the region of the certificate specified by `alicloudCertificateId`. but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateRegionId *string `pulumi:"alicloudCertificateRegionId"`
	// Name of the Server Certificate.
	Name *string `pulumi:"name"`
	// the content of privat key of the ssl certificate specified by `serverCertificate`. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
	PrivateKey *string `pulumi:"privateKey"`
	// The Id of resource group which the slb server certificate belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// the content of the ssl certificate. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
	ServerCertificate *string `pulumi:"serverCertificate"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type ServerCertificateState struct {
	// Deprecated: Field 'alicloud_certifacte_id' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_id' replaces it.
	AlicloudCertifacteId pulumi.StringPtrInput
	// Deprecated: Field 'alicloud_certifacte_name' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_name' replaces it.
	AlicloudCertifacteName pulumi.StringPtrInput
	// an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateId pulumi.StringPtrInput
	// the name of the certificate specified by `alicloudCertificateId`.but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateName pulumi.StringPtrInput
	// the region of the certificate specified by `alicloudCertificateId`. but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateRegionId pulumi.StringPtrInput
	// Name of the Server Certificate.
	Name pulumi.StringPtrInput
	// the content of privat key of the ssl certificate specified by `serverCertificate`. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
	PrivateKey pulumi.StringPtrInput
	// The Id of resource group which the slb server certificate belongs.
	ResourceGroupId pulumi.StringPtrInput
	// the content of the ssl certificate. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
	ServerCertificate pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (ServerCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverCertificateState)(nil)).Elem()
}

type serverCertificateArgs struct {
	// Deprecated: Field 'alicloud_certifacte_id' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_id' replaces it.
	AlicloudCertifacteId *string `pulumi:"alicloudCertifacteId"`
	// Deprecated: Field 'alicloud_certifacte_name' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_name' replaces it.
	AlicloudCertifacteName *string `pulumi:"alicloudCertifacteName"`
	// an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateId *string `pulumi:"alicloudCertificateId"`
	// the name of the certificate specified by `alicloudCertificateId`.but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateName *string `pulumi:"alicloudCertificateName"`
	// the region of the certificate specified by `alicloudCertificateId`. but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateRegionId *string `pulumi:"alicloudCertificateRegionId"`
	// Name of the Server Certificate.
	Name *string `pulumi:"name"`
	// the content of privat key of the ssl certificate specified by `serverCertificate`. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
	PrivateKey *string `pulumi:"privateKey"`
	// The Id of resource group which the slb server certificate belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// the content of the ssl certificate. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
	ServerCertificate *string `pulumi:"serverCertificate"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a ServerCertificate resource.
type ServerCertificateArgs struct {
	// Deprecated: Field 'alicloud_certifacte_id' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_id' replaces it.
	AlicloudCertifacteId pulumi.StringPtrInput
	// Deprecated: Field 'alicloud_certifacte_name' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_name' replaces it.
	AlicloudCertifacteName pulumi.StringPtrInput
	// an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateId pulumi.StringPtrInput
	// the name of the certificate specified by `alicloudCertificateId`.but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateName pulumi.StringPtrInput
	// the region of the certificate specified by `alicloudCertificateId`. but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateRegionId pulumi.StringPtrInput
	// Name of the Server Certificate.
	Name pulumi.StringPtrInput
	// the content of privat key of the ssl certificate specified by `serverCertificate`. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
	PrivateKey pulumi.StringPtrInput
	// The Id of resource group which the slb server certificate belongs.
	ResourceGroupId pulumi.StringPtrInput
	// the content of the ssl certificate. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
	ServerCertificate pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (ServerCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverCertificateArgs)(nil)).Elem()
}

type ServerCertificateInput interface {
	pulumi.Input

	ToServerCertificateOutput() ServerCertificateOutput
	ToServerCertificateOutputWithContext(ctx context.Context) ServerCertificateOutput
}

func (ServerCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerCertificate)(nil)).Elem()
}

func (i ServerCertificate) ToServerCertificateOutput() ServerCertificateOutput {
	return i.ToServerCertificateOutputWithContext(context.Background())
}

func (i ServerCertificate) ToServerCertificateOutputWithContext(ctx context.Context) ServerCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCertificateOutput)
}

type ServerCertificateOutput struct {
	*pulumi.OutputState
}

func (ServerCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerCertificateOutput)(nil)).Elem()
}

func (o ServerCertificateOutput) ToServerCertificateOutput() ServerCertificateOutput {
	return o
}

func (o ServerCertificateOutput) ToServerCertificateOutputWithContext(ctx context.Context) ServerCertificateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServerCertificateOutput{})
}
