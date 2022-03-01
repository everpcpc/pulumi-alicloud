// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package polardb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a PolarDB account privilege resource and used to grant several database some access privilege. A database can be granted by multiple account.
//
// > **NOTE:** Available in v1.67.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/polardb"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		creation := "PolarDB"
// 		if param := cfg.Get("creation"); param != "" {
// 			creation = param
// 		}
// 		name := "dbaccountprivilegebasic"
// 		if param := cfg.Get("name"); param != "" {
// 			name = param
// 		}
// 		defaultZones, err := alicloud.GetZones(ctx, &GetZonesArgs{
// 			AvailableResourceCreation: pulumi.StringRef(creation),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
// 			CidrBlock: pulumi.String("172.16.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
// 			VpcId:     defaultNetwork.ID(),
// 			CidrBlock: pulumi.String("172.16.0.0/24"),
// 			ZoneId:    pulumi.String(defaultZones.Zones[0].Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		cluster, err := polardb.NewCluster(ctx, "cluster", &polardb.ClusterArgs{
// 			DbType:      pulumi.String("MySQL"),
// 			DbVersion:   pulumi.String("8.0"),
// 			PayType:     pulumi.String("PostPaid"),
// 			DbNodeClass: pulumi.String("polar.mysql.x4.large"),
// 			VswitchId:   defaultSwitch.ID(),
// 			Description: pulumi.String(name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		db, err := polardb.NewDatabase(ctx, "db", &polardb.DatabaseArgs{
// 			DbClusterId: cluster.ID(),
// 			DbName:      pulumi.String("tftestdatabase"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		account, err := polardb.NewAccount(ctx, "account", &polardb.AccountArgs{
// 			DbClusterId:        cluster.ID(),
// 			AccountName:        pulumi.String("tftestnormal"),
// 			AccountPassword:    pulumi.String("Test12345"),
// 			AccountDescription: pulumi.String(name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = polardb.NewAccountPrivilege(ctx, "privilege", &polardb.AccountPrivilegeArgs{
// 			DbClusterId:      cluster.ID(),
// 			AccountName:      account.AccountName,
// 			AccountPrivilege: pulumi.String("ReadOnly"),
// 			DbNames: pulumi.StringArray{
// 				db.DbName,
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
// PolarDB account privilege can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:polardb/accountPrivilege:AccountPrivilege example "pc-12345:tf_account:ReadOnly"
// ```
type AccountPrivilege struct {
	pulumi.CustomResourceState

	// A specified account name.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"], ["DMLOnly", "DDLOnly"] added since version v1.101.0. Default to "ReadOnly".
	AccountPrivilege pulumi.StringPtrOutput `pulumi:"accountPrivilege"`
	// The Id of cluster in which account belongs.
	DbClusterId pulumi.StringOutput `pulumi:"dbClusterId"`
	// List of specified database name.
	DbNames pulumi.StringArrayOutput `pulumi:"dbNames"`
}

// NewAccountPrivilege registers a new resource with the given unique name, arguments, and options.
func NewAccountPrivilege(ctx *pulumi.Context,
	name string, args *AccountPrivilegeArgs, opts ...pulumi.ResourceOption) (*AccountPrivilege, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DbClusterId == nil {
		return nil, errors.New("invalid value for required argument 'DbClusterId'")
	}
	if args.DbNames == nil {
		return nil, errors.New("invalid value for required argument 'DbNames'")
	}
	var resource AccountPrivilege
	err := ctx.RegisterResource("alicloud:polardb/accountPrivilege:AccountPrivilege", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountPrivilege gets an existing AccountPrivilege resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountPrivilege(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountPrivilegeState, opts ...pulumi.ResourceOption) (*AccountPrivilege, error) {
	var resource AccountPrivilege
	err := ctx.ReadResource("alicloud:polardb/accountPrivilege:AccountPrivilege", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountPrivilege resources.
type accountPrivilegeState struct {
	// A specified account name.
	AccountName *string `pulumi:"accountName"`
	// The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"], ["DMLOnly", "DDLOnly"] added since version v1.101.0. Default to "ReadOnly".
	AccountPrivilege *string `pulumi:"accountPrivilege"`
	// The Id of cluster in which account belongs.
	DbClusterId *string `pulumi:"dbClusterId"`
	// List of specified database name.
	DbNames []string `pulumi:"dbNames"`
}

type AccountPrivilegeState struct {
	// A specified account name.
	AccountName pulumi.StringPtrInput
	// The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"], ["DMLOnly", "DDLOnly"] added since version v1.101.0. Default to "ReadOnly".
	AccountPrivilege pulumi.StringPtrInput
	// The Id of cluster in which account belongs.
	DbClusterId pulumi.StringPtrInput
	// List of specified database name.
	DbNames pulumi.StringArrayInput
}

func (AccountPrivilegeState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountPrivilegeState)(nil)).Elem()
}

type accountPrivilegeArgs struct {
	// A specified account name.
	AccountName string `pulumi:"accountName"`
	// The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"], ["DMLOnly", "DDLOnly"] added since version v1.101.0. Default to "ReadOnly".
	AccountPrivilege *string `pulumi:"accountPrivilege"`
	// The Id of cluster in which account belongs.
	DbClusterId string `pulumi:"dbClusterId"`
	// List of specified database name.
	DbNames []string `pulumi:"dbNames"`
}

// The set of arguments for constructing a AccountPrivilege resource.
type AccountPrivilegeArgs struct {
	// A specified account name.
	AccountName pulumi.StringInput
	// The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"], ["DMLOnly", "DDLOnly"] added since version v1.101.0. Default to "ReadOnly".
	AccountPrivilege pulumi.StringPtrInput
	// The Id of cluster in which account belongs.
	DbClusterId pulumi.StringInput
	// List of specified database name.
	DbNames pulumi.StringArrayInput
}

func (AccountPrivilegeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountPrivilegeArgs)(nil)).Elem()
}

type AccountPrivilegeInput interface {
	pulumi.Input

	ToAccountPrivilegeOutput() AccountPrivilegeOutput
	ToAccountPrivilegeOutputWithContext(ctx context.Context) AccountPrivilegeOutput
}

func (*AccountPrivilege) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountPrivilege)(nil)).Elem()
}

func (i *AccountPrivilege) ToAccountPrivilegeOutput() AccountPrivilegeOutput {
	return i.ToAccountPrivilegeOutputWithContext(context.Background())
}

func (i *AccountPrivilege) ToAccountPrivilegeOutputWithContext(ctx context.Context) AccountPrivilegeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPrivilegeOutput)
}

// AccountPrivilegeArrayInput is an input type that accepts AccountPrivilegeArray and AccountPrivilegeArrayOutput values.
// You can construct a concrete instance of `AccountPrivilegeArrayInput` via:
//
//          AccountPrivilegeArray{ AccountPrivilegeArgs{...} }
type AccountPrivilegeArrayInput interface {
	pulumi.Input

	ToAccountPrivilegeArrayOutput() AccountPrivilegeArrayOutput
	ToAccountPrivilegeArrayOutputWithContext(context.Context) AccountPrivilegeArrayOutput
}

type AccountPrivilegeArray []AccountPrivilegeInput

func (AccountPrivilegeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountPrivilege)(nil)).Elem()
}

func (i AccountPrivilegeArray) ToAccountPrivilegeArrayOutput() AccountPrivilegeArrayOutput {
	return i.ToAccountPrivilegeArrayOutputWithContext(context.Background())
}

func (i AccountPrivilegeArray) ToAccountPrivilegeArrayOutputWithContext(ctx context.Context) AccountPrivilegeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPrivilegeArrayOutput)
}

// AccountPrivilegeMapInput is an input type that accepts AccountPrivilegeMap and AccountPrivilegeMapOutput values.
// You can construct a concrete instance of `AccountPrivilegeMapInput` via:
//
//          AccountPrivilegeMap{ "key": AccountPrivilegeArgs{...} }
type AccountPrivilegeMapInput interface {
	pulumi.Input

	ToAccountPrivilegeMapOutput() AccountPrivilegeMapOutput
	ToAccountPrivilegeMapOutputWithContext(context.Context) AccountPrivilegeMapOutput
}

type AccountPrivilegeMap map[string]AccountPrivilegeInput

func (AccountPrivilegeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountPrivilege)(nil)).Elem()
}

func (i AccountPrivilegeMap) ToAccountPrivilegeMapOutput() AccountPrivilegeMapOutput {
	return i.ToAccountPrivilegeMapOutputWithContext(context.Background())
}

func (i AccountPrivilegeMap) ToAccountPrivilegeMapOutputWithContext(ctx context.Context) AccountPrivilegeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPrivilegeMapOutput)
}

type AccountPrivilegeOutput struct{ *pulumi.OutputState }

func (AccountPrivilegeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountPrivilege)(nil)).Elem()
}

func (o AccountPrivilegeOutput) ToAccountPrivilegeOutput() AccountPrivilegeOutput {
	return o
}

func (o AccountPrivilegeOutput) ToAccountPrivilegeOutputWithContext(ctx context.Context) AccountPrivilegeOutput {
	return o
}

type AccountPrivilegeArrayOutput struct{ *pulumi.OutputState }

func (AccountPrivilegeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountPrivilege)(nil)).Elem()
}

func (o AccountPrivilegeArrayOutput) ToAccountPrivilegeArrayOutput() AccountPrivilegeArrayOutput {
	return o
}

func (o AccountPrivilegeArrayOutput) ToAccountPrivilegeArrayOutputWithContext(ctx context.Context) AccountPrivilegeArrayOutput {
	return o
}

func (o AccountPrivilegeArrayOutput) Index(i pulumi.IntInput) AccountPrivilegeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccountPrivilege {
		return vs[0].([]*AccountPrivilege)[vs[1].(int)]
	}).(AccountPrivilegeOutput)
}

type AccountPrivilegeMapOutput struct{ *pulumi.OutputState }

func (AccountPrivilegeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountPrivilege)(nil)).Elem()
}

func (o AccountPrivilegeMapOutput) ToAccountPrivilegeMapOutput() AccountPrivilegeMapOutput {
	return o
}

func (o AccountPrivilegeMapOutput) ToAccountPrivilegeMapOutputWithContext(ctx context.Context) AccountPrivilegeMapOutput {
	return o
}

func (o AccountPrivilegeMapOutput) MapIndex(k pulumi.StringInput) AccountPrivilegeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccountPrivilege {
		return vs[0].(map[string]*AccountPrivilege)[vs[1].(string)]
	}).(AccountPrivilegeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPrivilegeInput)(nil)).Elem(), &AccountPrivilege{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPrivilegeArrayInput)(nil)).Elem(), AccountPrivilegeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPrivilegeMapInput)(nil)).Elem(), AccountPrivilegeMap{})
	pulumi.RegisterOutputType(AccountPrivilegeOutput{})
	pulumi.RegisterOutputType(AccountPrivilegeArrayOutput{})
	pulumi.RegisterOutputType(AccountPrivilegeMapOutput{})
}
