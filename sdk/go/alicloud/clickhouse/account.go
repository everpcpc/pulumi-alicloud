// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clickhouse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Click House Account resource.
//
// For information about Click House Account and how to use it, see [What is Account](https://www.alibabacloud.com/product/clickhouse).
//
// > **NOTE:** Available in v1.134.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/clickhouse"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		name := "testaccountname"
// 		if param := cfg.Get("name"); param != "" {
// 			name = param
// 		}
// 		pwd := "Tf-testpwd"
// 		if param := cfg.Get("pwd"); param != "" {
// 			pwd = param
// 		}
// 		defaultRegions, err := clickhouse.GetRegions(ctx, &clickhouse.GetRegionsArgs{
// 			Current: pulumi.BoolRef(true),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		defaultNetworks, err := vpc.GetNetworks(ctx, &vpc.GetNetworksArgs{
// 			NameRegex: pulumi.StringRef("default-NODELETING"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		defaultSwitches, err := vpc.GetSwitches(ctx, &vpc.GetSwitchesArgs{
// 			VpcId:  pulumi.StringRef(defaultNetworks.Ids[0]),
// 			ZoneId: pulumi.StringRef(defaultRegions.Regions[0].ZoneIds[0].ZoneId),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		defaultDbCluster, err := clickhouse.NewDbCluster(ctx, "defaultDbCluster", &clickhouse.DbClusterArgs{
// 			DbClusterVersion:     pulumi.String("20.3.10.75"),
// 			Category:             pulumi.String("Basic"),
// 			DbClusterClass:       pulumi.String("S8"),
// 			DbClusterNetworkType: pulumi.String("vpc"),
// 			DbClusterDescription: pulumi.String(name),
// 			DbNodeGroupCount:     pulumi.Int(1),
// 			PaymentType:          pulumi.String("PayAsYouGo"),
// 			DbNodeStorage:        pulumi.String("500"),
// 			StorageType:          pulumi.String("cloud_essd"),
// 			VswitchId:            pulumi.String(defaultSwitches.Vswitches[0].Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = clickhouse.NewAccount(ctx, "defaultAccount", &clickhouse.AccountArgs{
// 			DbClusterId:        defaultDbCluster.ID(),
// 			AccountDescription: pulumi.String("your_description"),
// 			AccountName:        pulumi.String(name),
// 			AccountPassword:    pulumi.String(pwd),
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
// Click House Account can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:clickhouse/account:Account example <db_cluster_id>:<account_name>
// ```
type Account struct {
	pulumi.CustomResourceState

	// In Chinese, English letter. May contain Chinese and English characters, lowercase letters, numbers, and underscores (_), the dash (-). Cannot start with http:// and https:// at the beginning. Length is from 2 to 256 characters.
	AccountDescription pulumi.StringPtrOutput `pulumi:"accountDescription"`
	// Account name: lowercase letters, numbers, underscores, lowercase letter; length no more than 16 characters.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// The account password: uppercase letters, lowercase letters, lowercase letters, numbers, and special characters (special character! #$%^& author (s):_+-=) in a length of 8-32 bit.
	AccountPassword pulumi.StringOutput `pulumi:"accountPassword"`
	// The list of databases to which you want to grant permissions. Separate databases with commas (,).
	AllowDatabases pulumi.StringOutput `pulumi:"allowDatabases"`
	// The list of dictionaries to which you want to grant permissions. Separate dictionaries with commas (,).
	AllowDictionaries pulumi.StringOutput `pulumi:"allowDictionaries"`
	// The db cluster id.
	DbClusterId pulumi.StringOutput `pulumi:"dbClusterId"`
	// Specifies whether to grant DDL permissions to the database account. Valid values: `true` and `false`.
	// -`true`: grants DDL permissions to the database account.
	// -`false`: does not grant DDL permissions to the database account.
	DdlAuthority pulumi.BoolOutput `pulumi:"ddlAuthority"`
	// Specifies whether to grant DML permissions to the database account. Valid values: `all` and `readonly,modify`.
	DmlAuthority pulumi.StringOutput `pulumi:"dmlAuthority"`
	// The status of the resource. Valid Status: `Creating`,`Available`,`Deleting`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The list of all databases. Separate databases with commas (,).
	TotalDatabases pulumi.StringOutput `pulumi:"totalDatabases"`
	// The list of all dictionaries. Separate dictionaries with commas (,).
	TotalDictionaries pulumi.StringOutput `pulumi:"totalDictionaries"`
	// The type of the database account. Valid values: `Normal` or `Super`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.AccountPassword == nil {
		return nil, errors.New("invalid value for required argument 'AccountPassword'")
	}
	if args.DbClusterId == nil {
		return nil, errors.New("invalid value for required argument 'DbClusterId'")
	}
	var resource Account
	err := ctx.RegisterResource("alicloud:clickhouse/account:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("alicloud:clickhouse/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// In Chinese, English letter. May contain Chinese and English characters, lowercase letters, numbers, and underscores (_), the dash (-). Cannot start with http:// and https:// at the beginning. Length is from 2 to 256 characters.
	AccountDescription *string `pulumi:"accountDescription"`
	// Account name: lowercase letters, numbers, underscores, lowercase letter; length no more than 16 characters.
	AccountName *string `pulumi:"accountName"`
	// The account password: uppercase letters, lowercase letters, lowercase letters, numbers, and special characters (special character! #$%^& author (s):_+-=) in a length of 8-32 bit.
	AccountPassword *string `pulumi:"accountPassword"`
	// The list of databases to which you want to grant permissions. Separate databases with commas (,).
	AllowDatabases *string `pulumi:"allowDatabases"`
	// The list of dictionaries to which you want to grant permissions. Separate dictionaries with commas (,).
	AllowDictionaries *string `pulumi:"allowDictionaries"`
	// The db cluster id.
	DbClusterId *string `pulumi:"dbClusterId"`
	// Specifies whether to grant DDL permissions to the database account. Valid values: `true` and `false`.
	// -`true`: grants DDL permissions to the database account.
	// -`false`: does not grant DDL permissions to the database account.
	DdlAuthority *bool `pulumi:"ddlAuthority"`
	// Specifies whether to grant DML permissions to the database account. Valid values: `all` and `readonly,modify`.
	DmlAuthority *string `pulumi:"dmlAuthority"`
	// The status of the resource. Valid Status: `Creating`,`Available`,`Deleting`.
	Status *string `pulumi:"status"`
	// The list of all databases. Separate databases with commas (,).
	TotalDatabases *string `pulumi:"totalDatabases"`
	// The list of all dictionaries. Separate dictionaries with commas (,).
	TotalDictionaries *string `pulumi:"totalDictionaries"`
	// The type of the database account. Valid values: `Normal` or `Super`.
	Type *string `pulumi:"type"`
}

type AccountState struct {
	// In Chinese, English letter. May contain Chinese and English characters, lowercase letters, numbers, and underscores (_), the dash (-). Cannot start with http:// and https:// at the beginning. Length is from 2 to 256 characters.
	AccountDescription pulumi.StringPtrInput
	// Account name: lowercase letters, numbers, underscores, lowercase letter; length no more than 16 characters.
	AccountName pulumi.StringPtrInput
	// The account password: uppercase letters, lowercase letters, lowercase letters, numbers, and special characters (special character! #$%^& author (s):_+-=) in a length of 8-32 bit.
	AccountPassword pulumi.StringPtrInput
	// The list of databases to which you want to grant permissions. Separate databases with commas (,).
	AllowDatabases pulumi.StringPtrInput
	// The list of dictionaries to which you want to grant permissions. Separate dictionaries with commas (,).
	AllowDictionaries pulumi.StringPtrInput
	// The db cluster id.
	DbClusterId pulumi.StringPtrInput
	// Specifies whether to grant DDL permissions to the database account. Valid values: `true` and `false`.
	// -`true`: grants DDL permissions to the database account.
	// -`false`: does not grant DDL permissions to the database account.
	DdlAuthority pulumi.BoolPtrInput
	// Specifies whether to grant DML permissions to the database account. Valid values: `all` and `readonly,modify`.
	DmlAuthority pulumi.StringPtrInput
	// The status of the resource. Valid Status: `Creating`,`Available`,`Deleting`.
	Status pulumi.StringPtrInput
	// The list of all databases. Separate databases with commas (,).
	TotalDatabases pulumi.StringPtrInput
	// The list of all dictionaries. Separate dictionaries with commas (,).
	TotalDictionaries pulumi.StringPtrInput
	// The type of the database account. Valid values: `Normal` or `Super`.
	Type pulumi.StringPtrInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// In Chinese, English letter. May contain Chinese and English characters, lowercase letters, numbers, and underscores (_), the dash (-). Cannot start with http:// and https:// at the beginning. Length is from 2 to 256 characters.
	AccountDescription *string `pulumi:"accountDescription"`
	// Account name: lowercase letters, numbers, underscores, lowercase letter; length no more than 16 characters.
	AccountName string `pulumi:"accountName"`
	// The account password: uppercase letters, lowercase letters, lowercase letters, numbers, and special characters (special character! #$%^& author (s):_+-=) in a length of 8-32 bit.
	AccountPassword string `pulumi:"accountPassword"`
	// The list of databases to which you want to grant permissions. Separate databases with commas (,).
	AllowDatabases *string `pulumi:"allowDatabases"`
	// The list of dictionaries to which you want to grant permissions. Separate dictionaries with commas (,).
	AllowDictionaries *string `pulumi:"allowDictionaries"`
	// The db cluster id.
	DbClusterId string `pulumi:"dbClusterId"`
	// Specifies whether to grant DDL permissions to the database account. Valid values: `true` and `false`.
	// -`true`: grants DDL permissions to the database account.
	// -`false`: does not grant DDL permissions to the database account.
	DdlAuthority *bool `pulumi:"ddlAuthority"`
	// Specifies whether to grant DML permissions to the database account. Valid values: `all` and `readonly,modify`.
	DmlAuthority *string `pulumi:"dmlAuthority"`
	// The list of all databases. Separate databases with commas (,).
	TotalDatabases *string `pulumi:"totalDatabases"`
	// The list of all dictionaries. Separate dictionaries with commas (,).
	TotalDictionaries *string `pulumi:"totalDictionaries"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// In Chinese, English letter. May contain Chinese and English characters, lowercase letters, numbers, and underscores (_), the dash (-). Cannot start with http:// and https:// at the beginning. Length is from 2 to 256 characters.
	AccountDescription pulumi.StringPtrInput
	// Account name: lowercase letters, numbers, underscores, lowercase letter; length no more than 16 characters.
	AccountName pulumi.StringInput
	// The account password: uppercase letters, lowercase letters, lowercase letters, numbers, and special characters (special character! #$%^& author (s):_+-=) in a length of 8-32 bit.
	AccountPassword pulumi.StringInput
	// The list of databases to which you want to grant permissions. Separate databases with commas (,).
	AllowDatabases pulumi.StringPtrInput
	// The list of dictionaries to which you want to grant permissions. Separate dictionaries with commas (,).
	AllowDictionaries pulumi.StringPtrInput
	// The db cluster id.
	DbClusterId pulumi.StringInput
	// Specifies whether to grant DDL permissions to the database account. Valid values: `true` and `false`.
	// -`true`: grants DDL permissions to the database account.
	// -`false`: does not grant DDL permissions to the database account.
	DdlAuthority pulumi.BoolPtrInput
	// Specifies whether to grant DML permissions to the database account. Valid values: `all` and `readonly,modify`.
	DmlAuthority pulumi.StringPtrInput
	// The list of all databases. Separate databases with commas (,).
	TotalDatabases pulumi.StringPtrInput
	// The list of all dictionaries. Separate dictionaries with commas (,).
	TotalDictionaries pulumi.StringPtrInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (*Account) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

// AccountArrayInput is an input type that accepts AccountArray and AccountArrayOutput values.
// You can construct a concrete instance of `AccountArrayInput` via:
//
//          AccountArray{ AccountArgs{...} }
type AccountArrayInput interface {
	pulumi.Input

	ToAccountArrayOutput() AccountArrayOutput
	ToAccountArrayOutputWithContext(context.Context) AccountArrayOutput
}

type AccountArray []AccountInput

func (AccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Account)(nil)).Elem()
}

func (i AccountArray) ToAccountArrayOutput() AccountArrayOutput {
	return i.ToAccountArrayOutputWithContext(context.Background())
}

func (i AccountArray) ToAccountArrayOutputWithContext(ctx context.Context) AccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountArrayOutput)
}

// AccountMapInput is an input type that accepts AccountMap and AccountMapOutput values.
// You can construct a concrete instance of `AccountMapInput` via:
//
//          AccountMap{ "key": AccountArgs{...} }
type AccountMapInput interface {
	pulumi.Input

	ToAccountMapOutput() AccountMapOutput
	ToAccountMapOutputWithContext(context.Context) AccountMapOutput
}

type AccountMap map[string]AccountInput

func (AccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Account)(nil)).Elem()
}

func (i AccountMap) ToAccountMapOutput() AccountMapOutput {
	return i.ToAccountMapOutputWithContext(context.Background())
}

func (i AccountMap) ToAccountMapOutputWithContext(ctx context.Context) AccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountMapOutput)
}

type AccountOutput struct{ *pulumi.OutputState }

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

type AccountArrayOutput struct{ *pulumi.OutputState }

func (AccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Account)(nil)).Elem()
}

func (o AccountArrayOutput) ToAccountArrayOutput() AccountArrayOutput {
	return o
}

func (o AccountArrayOutput) ToAccountArrayOutputWithContext(ctx context.Context) AccountArrayOutput {
	return o
}

func (o AccountArrayOutput) Index(i pulumi.IntInput) AccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Account {
		return vs[0].([]*Account)[vs[1].(int)]
	}).(AccountOutput)
}

type AccountMapOutput struct{ *pulumi.OutputState }

func (AccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Account)(nil)).Elem()
}

func (o AccountMapOutput) ToAccountMapOutput() AccountMapOutput {
	return o
}

func (o AccountMapOutput) ToAccountMapOutputWithContext(ctx context.Context) AccountMapOutput {
	return o
}

func (o AccountMapOutput) MapIndex(k pulumi.StringInput) AccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Account {
		return vs[0].(map[string]*Account)[vs[1].(string)]
	}).(AccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountInput)(nil)).Elem(), &Account{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountArrayInput)(nil)).Elem(), AccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountMapInput)(nil)).Elem(), AccountMap{})
	pulumi.RegisterOutputType(AccountOutput{})
	pulumi.RegisterOutputType(AccountArrayOutput{})
	pulumi.RegisterOutputType(AccountMapOutput{})
}
