// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package polardb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `polardb.getAccounts` data source provides a collection of PolarDB cluster database account available in Alibaba Cloud account.
// Filters support regular expression for the account name, searches by clusterId.
//
// > **NOTE:** Available in v1.70.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/polardb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			polardbClustersDs, err := polardb.GetClusters(ctx, &polardb.GetClustersArgs{
//				DescriptionRegex: pulumi.StringRef("pc-\\w+"),
//				Status:           pulumi.StringRef("Running"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_default, err := polardb.GetAccounts(ctx, &polardb.GetAccountsArgs{
//				DbClusterId: polardbClustersDs.Clusters[0].Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("account", _default.Accounts[0].AccountName)
//			return nil
//		})
//	}
//
// ```
func GetAccounts(ctx *pulumi.Context, args *GetAccountsArgs, opts ...pulumi.InvokeOption) (*GetAccountsResult, error) {
	var rv GetAccountsResult
	err := ctx.Invoke("alicloud:polardb/getAccounts:getAccounts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccounts.
type GetAccountsArgs struct {
	// The polarDB cluster ID.
	DbClusterId string `pulumi:"dbClusterId"`
	// A regex string to filter results by account name.
	NameRegex *string `pulumi:"nameRegex"`
}

// A collection of values returned by getAccounts.
type GetAccountsResult struct {
	// A list of PolarDB cluster accounts. Each element contains the following attributes:
	Accounts    []GetAccountsAccount `pulumi:"accounts"`
	DbClusterId string               `pulumi:"dbClusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	NameRegex *string `pulumi:"nameRegex"`
	// Account name of the cluster.
	Names []string `pulumi:"names"`
}

func GetAccountsOutput(ctx *pulumi.Context, args GetAccountsOutputArgs, opts ...pulumi.InvokeOption) GetAccountsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAccountsResult, error) {
			args := v.(GetAccountsArgs)
			r, err := GetAccounts(ctx, &args, opts...)
			var s GetAccountsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAccountsResultOutput)
}

// A collection of arguments for invoking getAccounts.
type GetAccountsOutputArgs struct {
	// The polarDB cluster ID.
	DbClusterId pulumi.StringInput `pulumi:"dbClusterId"`
	// A regex string to filter results by account name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
}

func (GetAccountsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountsArgs)(nil)).Elem()
}

// A collection of values returned by getAccounts.
type GetAccountsResultOutput struct{ *pulumi.OutputState }

func (GetAccountsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountsResult)(nil)).Elem()
}

func (o GetAccountsResultOutput) ToGetAccountsResultOutput() GetAccountsResultOutput {
	return o
}

func (o GetAccountsResultOutput) ToGetAccountsResultOutputWithContext(ctx context.Context) GetAccountsResultOutput {
	return o
}

// A list of PolarDB cluster accounts. Each element contains the following attributes:
func (o GetAccountsResultOutput) Accounts() GetAccountsAccountArrayOutput {
	return o.ApplyT(func(v GetAccountsResult) []GetAccountsAccount { return v.Accounts }).(GetAccountsAccountArrayOutput)
}

func (o GetAccountsResultOutput) DbClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountsResult) string { return v.DbClusterId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAccountsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAccountsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccountsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// Account name of the cluster.
func (o GetAccountsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAccountsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccountsResultOutput{})
}
