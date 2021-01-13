// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an RDS account privilege resource and used to grant several database some access privilege. A database can be granted by multiple account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/rds"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		creation := "Rds"
// 		if param := cfg.Get("creation"); param != "" {
// 			creation = param
// 		}
// 		name := "dbaccountprivilegebasic"
// 		if param := cfg.Get("name"); param != "" {
// 			name = param
// 		}
// 		opt0 := creation
// 		defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
// 			AvailableResourceCreation: &opt0,
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
// 			VpcId:            defaultNetwork.ID(),
// 			CidrBlock:        pulumi.String("172.16.0.0/24"),
// 			AvailabilityZone: pulumi.String(defaultZones.Zones[0].Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		instance, err := rds.NewInstance(ctx, "instance", &rds.InstanceArgs{
// 			Engine:          pulumi.String("MySQL"),
// 			EngineVersion:   pulumi.String("5.6"),
// 			InstanceType:    pulumi.String("rds.mysql.s1.small"),
// 			InstanceStorage: pulumi.Int(10),
// 			VswitchId:       defaultSwitch.ID(),
// 			InstanceName:    pulumi.String(name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		var db []*rds.Database
// 		for key0, _ := range 2 {
// 			__res, err := rds.NewDatabase(ctx, fmt.Sprintf("db-%v", key0), &rds.DatabaseArgs{
// 				InstanceId:  instance.ID(),
// 				Description: pulumi.String("from terraform"),
// 			})
// 			if err != nil {
// 				return err
// 			}
// 			db = append(db, __res)
// 		}
// 		account, err := rds.NewAccount(ctx, "account", &rds.AccountArgs{
// 			InstanceId:  instance.ID(),
// 			Password:    pulumi.String("Test12345"),
// 			Description: pulumi.String("from terraform"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		var splat0 pulumi.StringArray
// 		for _, val0 := range db {
// 			splat0 = append(splat0, val0.Name)
// 		}
// 		_, err = rds.NewAccountPrivilege(ctx, "privilege", &rds.AccountPrivilegeArgs{
// 			InstanceId:  instance.ID(),
// 			AccountName: account.Name,
// 			Privilege:   pulumi.String("ReadOnly"),
// 			DbNames:     toPulumiStringArray(splat0),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// func toPulumiStringArray(arr []string) pulumi.StringArray {
// 	var pulumiArr pulumi.StringArray
// 	for _, v := range arr {
// 		pulumiArr = append(pulumiArr, pulumi.String(v))
// 	}
// 	return pulumiArr
// }
// ```
//
// ## Import
//
// RDS account privilege can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:rds/accountPrivilege:AccountPrivilege example "rm-12345:tf_account:ReadOnly"
// ```
type AccountPrivilege struct {
	pulumi.CustomResourceState

	// A specified account name.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// List of specified database name.
	DbNames pulumi.StringArrayOutput `pulumi:"dbNames"`
	// The Id of instance in which account belongs.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The privilege of one account access database. Valid values:
	// - ReadOnly: This value is only for MySQL, MariaDB and SQL Server
	// - ReadWrite: This value is only for MySQL, MariaDB and SQL Server
	// - DDLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
	// - DMLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
	// - DBOwner: (Available in 1.64.0+) This value is only for SQL Server and PostgreSQL.
	Privilege pulumi.StringPtrOutput `pulumi:"privilege"`
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
	if args.DbNames == nil {
		return nil, errors.New("invalid value for required argument 'DbNames'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource AccountPrivilege
	err := ctx.RegisterResource("alicloud:rds/accountPrivilege:AccountPrivilege", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:rds/accountPrivilege:AccountPrivilege", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountPrivilege resources.
type accountPrivilegeState struct {
	// A specified account name.
	AccountName *string `pulumi:"accountName"`
	// List of specified database name.
	DbNames []string `pulumi:"dbNames"`
	// The Id of instance in which account belongs.
	InstanceId *string `pulumi:"instanceId"`
	// The privilege of one account access database. Valid values:
	// - ReadOnly: This value is only for MySQL, MariaDB and SQL Server
	// - ReadWrite: This value is only for MySQL, MariaDB and SQL Server
	// - DDLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
	// - DMLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
	// - DBOwner: (Available in 1.64.0+) This value is only for SQL Server and PostgreSQL.
	Privilege *string `pulumi:"privilege"`
}

type AccountPrivilegeState struct {
	// A specified account name.
	AccountName pulumi.StringPtrInput
	// List of specified database name.
	DbNames pulumi.StringArrayInput
	// The Id of instance in which account belongs.
	InstanceId pulumi.StringPtrInput
	// The privilege of one account access database. Valid values:
	// - ReadOnly: This value is only for MySQL, MariaDB and SQL Server
	// - ReadWrite: This value is only for MySQL, MariaDB and SQL Server
	// - DDLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
	// - DMLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
	// - DBOwner: (Available in 1.64.0+) This value is only for SQL Server and PostgreSQL.
	Privilege pulumi.StringPtrInput
}

func (AccountPrivilegeState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountPrivilegeState)(nil)).Elem()
}

type accountPrivilegeArgs struct {
	// A specified account name.
	AccountName string `pulumi:"accountName"`
	// List of specified database name.
	DbNames []string `pulumi:"dbNames"`
	// The Id of instance in which account belongs.
	InstanceId string `pulumi:"instanceId"`
	// The privilege of one account access database. Valid values:
	// - ReadOnly: This value is only for MySQL, MariaDB and SQL Server
	// - ReadWrite: This value is only for MySQL, MariaDB and SQL Server
	// - DDLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
	// - DMLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
	// - DBOwner: (Available in 1.64.0+) This value is only for SQL Server and PostgreSQL.
	Privilege *string `pulumi:"privilege"`
}

// The set of arguments for constructing a AccountPrivilege resource.
type AccountPrivilegeArgs struct {
	// A specified account name.
	AccountName pulumi.StringInput
	// List of specified database name.
	DbNames pulumi.StringArrayInput
	// The Id of instance in which account belongs.
	InstanceId pulumi.StringInput
	// The privilege of one account access database. Valid values:
	// - ReadOnly: This value is only for MySQL, MariaDB and SQL Server
	// - ReadWrite: This value is only for MySQL, MariaDB and SQL Server
	// - DDLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
	// - DMLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
	// - DBOwner: (Available in 1.64.0+) This value is only for SQL Server and PostgreSQL.
	Privilege pulumi.StringPtrInput
}

func (AccountPrivilegeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountPrivilegeArgs)(nil)).Elem()
}

type AccountPrivilegeInput interface {
	pulumi.Input

	ToAccountPrivilegeOutput() AccountPrivilegeOutput
	ToAccountPrivilegeOutputWithContext(ctx context.Context) AccountPrivilegeOutput
}

func (AccountPrivilege) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountPrivilege)(nil)).Elem()
}

func (i AccountPrivilege) ToAccountPrivilegeOutput() AccountPrivilegeOutput {
	return i.ToAccountPrivilegeOutputWithContext(context.Background())
}

func (i AccountPrivilege) ToAccountPrivilegeOutputWithContext(ctx context.Context) AccountPrivilegeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPrivilegeOutput)
}

type AccountPrivilegeOutput struct {
	*pulumi.OutputState
}

func (AccountPrivilegeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountPrivilegeOutput)(nil)).Elem()
}

func (o AccountPrivilegeOutput) ToAccountPrivilegeOutput() AccountPrivilegeOutput {
	return o
}

func (o AccountPrivilegeOutput) ToAccountPrivilegeOutputWithContext(ctx context.Context) AccountPrivilegeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AccountPrivilegeOutput{})
}
