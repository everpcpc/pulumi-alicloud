// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package adb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an ADB connection resource to allocate an Internet connection string for ADB cluster.
//
// > **NOTE:** Each ADB instance will allocate a intranet connnection string automatically and its prifix is ADB instance ID.
//  To avoid unnecessary conflict, please specified a internet connection prefix before applying the resource.
//
// > **NOTE:** Available in v1.81.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/adb"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		creation := "ADB"
// 		if param := cfg.Get("creation"); param != "" {
// 			creation = param
// 		}
// 		name := "adbaccountmysql"
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
// 		cluster, err := adb.NewCluster(ctx, "cluster", &adb.ClusterArgs{
// 			DbClusterVersion:  pulumi.String("3.0"),
// 			DbClusterCategory: pulumi.String("Cluster"),
// 			DbNodeClass:       pulumi.String("C8"),
// 			DbNodeCount:       pulumi.Int(2),
// 			DbNodeStorage:     pulumi.Int(200),
// 			PayType:           pulumi.String("PostPaid"),
// 			VswitchId:         defaultSwitch.ID(),
// 			Description:       pulumi.String(name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = adb.NewConnection(ctx, "connection", &adb.ConnectionArgs{
// 			DbClusterId:      cluster.ID(),
// 			ConnectionPrefix: pulumi.String("testabc"),
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
// ADB connection can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:adb/connection:Connection example am-12345678
// ```
type Connection struct {
	pulumi.CustomResourceState

	// Prefix of the cluster public endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter. Default to `<db_cluster_id> + tf`.
	ConnectionPrefix pulumi.StringOutput `pulumi:"connectionPrefix"`
	// Connection cluster string.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// The Id of cluster that can run database.
	DbClusterId pulumi.StringOutput `pulumi:"dbClusterId"`
	// The ip address of connection string.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// Connection cluster port.
	Port pulumi.StringOutput `pulumi:"port"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbClusterId == nil {
		return nil, errors.New("invalid value for required argument 'DbClusterId'")
	}
	var resource Connection
	err := ctx.RegisterResource("alicloud:adb/connection:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("alicloud:adb/connection:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
	// Prefix of the cluster public endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter. Default to `<db_cluster_id> + tf`.
	ConnectionPrefix *string `pulumi:"connectionPrefix"`
	// Connection cluster string.
	ConnectionString *string `pulumi:"connectionString"`
	// The Id of cluster that can run database.
	DbClusterId *string `pulumi:"dbClusterId"`
	// The ip address of connection string.
	IpAddress *string `pulumi:"ipAddress"`
	// Connection cluster port.
	Port *string `pulumi:"port"`
}

type ConnectionState struct {
	// Prefix of the cluster public endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter. Default to `<db_cluster_id> + tf`.
	ConnectionPrefix pulumi.StringPtrInput
	// Connection cluster string.
	ConnectionString pulumi.StringPtrInput
	// The Id of cluster that can run database.
	DbClusterId pulumi.StringPtrInput
	// The ip address of connection string.
	IpAddress pulumi.StringPtrInput
	// Connection cluster port.
	Port pulumi.StringPtrInput
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// Prefix of the cluster public endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter. Default to `<db_cluster_id> + tf`.
	ConnectionPrefix *string `pulumi:"connectionPrefix"`
	// The Id of cluster that can run database.
	DbClusterId string `pulumi:"dbClusterId"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// Prefix of the cluster public endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter. Default to `<db_cluster_id> + tf`.
	ConnectionPrefix pulumi.StringPtrInput
	// The Id of cluster that can run database.
	DbClusterId pulumi.StringInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}

type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput
}

func (Connection) ElementType() reflect.Type {
	return reflect.TypeOf((*Connection)(nil)).Elem()
}

func (i Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

type ConnectionOutput struct {
	*pulumi.OutputState
}

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionOutput)(nil)).Elem()
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConnectionOutput{})
}
