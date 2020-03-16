// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package polardb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a PolarDB endpoint address resource to allocate an Internet endpoint address string for PolarDB instance.
//
// > **NOTE:** Available in v1.68.0+. Each PolarDB instance will allocate a intranet connection string automatically and its prefix is Cluster ID.
//  To avoid unnecessary conflict, please specified a internet connection prefix before applying the resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/polardb_endpoint_address.html.markdown.
type EndpointAddress struct {
	pulumi.CustomResourceState

	// Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <db_endpoint_id> + 'tf'.
	ConnectionPrefix pulumi.StringOutput `pulumi:"connectionPrefix"`
	// Connection cluster or endpoint string.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// The Id of cluster that can run database.
	DbClusterId pulumi.StringOutput `pulumi:"dbClusterId"`
	// The Id of endpoint that can run database.
	DbEndpointId pulumi.StringOutput `pulumi:"dbEndpointId"`
	// The ip address of connection string.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// Internet connection net type. Valid value: `Public`. Default to `Public`. Currently supported only `Public`.
	NetType pulumi.StringPtrOutput `pulumi:"netType"`
	// Connection cluster or endpoint port.
	Port pulumi.StringOutput `pulumi:"port"`
}

// NewEndpointAddress registers a new resource with the given unique name, arguments, and options.
func NewEndpointAddress(ctx *pulumi.Context,
	name string, args *EndpointAddressArgs, opts ...pulumi.ResourceOption) (*EndpointAddress, error) {
	if args == nil || args.DbClusterId == nil {
		return nil, errors.New("missing required argument 'DbClusterId'")
	}
	if args == nil || args.DbEndpointId == nil {
		return nil, errors.New("missing required argument 'DbEndpointId'")
	}
	if args == nil {
		args = &EndpointAddressArgs{}
	}
	var resource EndpointAddress
	err := ctx.RegisterResource("alicloud:polardb/endpointAddress:EndpointAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointAddress gets an existing EndpointAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointAddressState, opts ...pulumi.ResourceOption) (*EndpointAddress, error) {
	var resource EndpointAddress
	err := ctx.ReadResource("alicloud:polardb/endpointAddress:EndpointAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointAddress resources.
type endpointAddressState struct {
	// Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <db_endpoint_id> + 'tf'.
	ConnectionPrefix *string `pulumi:"connectionPrefix"`
	// Connection cluster or endpoint string.
	ConnectionString *string `pulumi:"connectionString"`
	// The Id of cluster that can run database.
	DbClusterId *string `pulumi:"dbClusterId"`
	// The Id of endpoint that can run database.
	DbEndpointId *string `pulumi:"dbEndpointId"`
	// The ip address of connection string.
	IpAddress *string `pulumi:"ipAddress"`
	// Internet connection net type. Valid value: `Public`. Default to `Public`. Currently supported only `Public`.
	NetType *string `pulumi:"netType"`
	// Connection cluster or endpoint port.
	Port *string `pulumi:"port"`
}

type EndpointAddressState struct {
	// Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <db_endpoint_id> + 'tf'.
	ConnectionPrefix pulumi.StringPtrInput
	// Connection cluster or endpoint string.
	ConnectionString pulumi.StringPtrInput
	// The Id of cluster that can run database.
	DbClusterId pulumi.StringPtrInput
	// The Id of endpoint that can run database.
	DbEndpointId pulumi.StringPtrInput
	// The ip address of connection string.
	IpAddress pulumi.StringPtrInput
	// Internet connection net type. Valid value: `Public`. Default to `Public`. Currently supported only `Public`.
	NetType pulumi.StringPtrInput
	// Connection cluster or endpoint port.
	Port pulumi.StringPtrInput
}

func (EndpointAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointAddressState)(nil)).Elem()
}

type endpointAddressArgs struct {
	// Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <db_endpoint_id> + 'tf'.
	ConnectionPrefix *string `pulumi:"connectionPrefix"`
	// The Id of cluster that can run database.
	DbClusterId string `pulumi:"dbClusterId"`
	// The Id of endpoint that can run database.
	DbEndpointId string `pulumi:"dbEndpointId"`
	// Internet connection net type. Valid value: `Public`. Default to `Public`. Currently supported only `Public`.
	NetType *string `pulumi:"netType"`
}

// The set of arguments for constructing a EndpointAddress resource.
type EndpointAddressArgs struct {
	// Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <db_endpoint_id> + 'tf'.
	ConnectionPrefix pulumi.StringPtrInput
	// The Id of cluster that can run database.
	DbClusterId pulumi.StringInput
	// The Id of endpoint that can run database.
	DbEndpointId pulumi.StringInput
	// Internet connection net type. Valid value: `Public`. Default to `Public`. Currently supported only `Public`.
	NetType pulumi.StringPtrInput
}

func (EndpointAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointAddressArgs)(nil)).Elem()
}

