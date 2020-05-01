// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

//
// Deprecated: This resource has been deprecated in favour of DnsDomain
type Domain struct {
	pulumi.CustomResourceState

	// A list of the dns server name.
	DnsServers pulumi.StringArrayOutput `pulumi:"dnsServers"`
	// The domain ID.
	DomainId pulumi.StringOutput `pulumi:"domainId"`
	// Id of the group in which the domain will add. If not supplied, then use default group.
	GroupId pulumi.StringPtrOutput `pulumi:"groupId"`
	// Name of the domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Id of resource group which the dns belongs.
	ResourceGroupId pulumi.StringPtrOutput `pulumi:"resourceGroupId"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		args = &DomainArgs{}
	}
	var resource Domain
	err := ctx.RegisterResource("alicloud:dns/domain:Domain", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:dns/domain:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
	// A list of the dns server name.
	DnsServers []string `pulumi:"dnsServers"`
	// The domain ID.
	DomainId *string `pulumi:"domainId"`
	// Id of the group in which the domain will add. If not supplied, then use default group.
	GroupId *string `pulumi:"groupId"`
	// Name of the domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	Name *string `pulumi:"name"`
	// The Id of resource group which the dns belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
}

type DomainState struct {
	// A list of the dns server name.
	DnsServers pulumi.StringArrayInput
	// The domain ID.
	DomainId pulumi.StringPtrInput
	// Id of the group in which the domain will add. If not supplied, then use default group.
	GroupId pulumi.StringPtrInput
	// Name of the domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	Name pulumi.StringPtrInput
	// The Id of resource group which the dns belongs.
	ResourceGroupId pulumi.StringPtrInput
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// Id of the group in which the domain will add. If not supplied, then use default group.
	GroupId *string `pulumi:"groupId"`
	// Name of the domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	Name *string `pulumi:"name"`
	// The Id of resource group which the dns belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// Id of the group in which the domain will add. If not supplied, then use default group.
	GroupId pulumi.StringPtrInput
	// Name of the domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	Name pulumi.StringPtrInput
	// The Id of resource group which the dns belongs.
	ResourceGroupId pulumi.StringPtrInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}
