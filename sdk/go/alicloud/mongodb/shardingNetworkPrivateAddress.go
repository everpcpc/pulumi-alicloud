// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a MongoDB Sharding Network Private Address resource.
//
// For information about MongoDB Sharding Network Private Address and how to use it, see [What is Sharding Network Private Address](https://www.alibabacloud.com/help/en/doc-detail/141403.html).
//
// > **NOTE:** Available in v1.157.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/mongodb"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultZones, err := mongodb.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetworks, err := vpc.GetNetworks(ctx, &vpc.GetNetworksArgs{
//				NameRegex: pulumi.StringRef("default-NODELETING"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultSwitches, err := vpc.GetSwitches(ctx, &vpc.GetSwitchesArgs{
//				VpcId:  pulumi.StringRef(defaultNetworks.Ids[0]),
//				ZoneId: pulumi.StringRef(defaultZones.Zones[0].Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultShardingInstance, err := mongodb.NewShardingInstance(ctx, "defaultShardingInstance", &mongodb.ShardingInstanceArgs{
//				ZoneId:        *pulumi.String(defaultZones.Zones[0].Id),
//				VswitchId:     *pulumi.String(defaultSwitches.Ids[0]),
//				EngineVersion: pulumi.String("4.2"),
//				MongoLists: mongodb.ShardingInstanceMongoListArray{
//					&mongodb.ShardingInstanceMongoListArgs{
//						NodeClass: pulumi.String("dds.mongos.mid"),
//					},
//					&mongodb.ShardingInstanceMongoListArgs{
//						NodeClass: pulumi.String("dds.mongos.mid"),
//					},
//				},
//				ShardLists: mongodb.ShardingInstanceShardListArray{
//					&mongodb.ShardingInstanceShardListArgs{
//						NodeClass:   pulumi.String("dds.shard.mid"),
//						NodeStorage: pulumi.Int(10),
//					},
//					&mongodb.ShardingInstanceShardListArgs{
//						NodeClass:   pulumi.String("dds.shard.mid"),
//						NodeStorage: pulumi.Int(10),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = mongodb.NewShardingNetworkPrivateAddress(ctx, "example", &mongodb.ShardingNetworkPrivateAddressArgs{
//				DbInstanceId: defaultShardingInstance.ID(),
//				NodeId: defaultShardingInstance.ShardLists.ApplyT(func(shardLists []mongodb.ShardingInstanceShardList) (*string, error) {
//					return &shardLists[0].NodeId, nil
//				}).(pulumi.StringPtrOutput),
//				ZoneId:          defaultShardingInstance.ZoneId,
//				AccountName:     pulumi.String("example_value"),
//				AccountPassword: pulumi.String("YourPassword+12345"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// MongoDB Sharding Network Private Address can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:mongodb/shardingNetworkPrivateAddress:ShardingNetworkPrivateAddress example <db_instance_id>:<node_id>
//
// ```
type ShardingNetworkPrivateAddress struct {
	pulumi.CustomResourceState

	// The name of the account.
	// - The name must be 4 to 16 characters in length and can contain lowercase letters, digits, and underscores (_). It must start with a lowercase letter.
	// - You need to set the account name and password only when you apply for an endpoint for a shard or Configserver node for the first time. In this case, the account name and password are used for all shard and Configserver nodes.
	// - The permissions of this account are fixed to read-only.
	AccountName pulumi.StringPtrOutput `pulumi:"accountName"`
	// Account password.
	// - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%^&*()_+-=`.
	// - The password must be 8 to 32 characters in length.
	AccountPassword pulumi.StringPtrOutput `pulumi:"accountPassword"`
	// The db instance id.
	DbInstanceId pulumi.StringOutput `pulumi:"dbInstanceId"`
	// The endpoint of the instance.
	NetworkAddresses ShardingNetworkPrivateAddressNetworkAddressArrayOutput `pulumi:"networkAddresses"`
	// The ID of the Shard node or the ConfigServer node.
	NodeId pulumi.StringOutput `pulumi:"nodeId"`
	// The zone ID of the instance.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewShardingNetworkPrivateAddress registers a new resource with the given unique name, arguments, and options.
func NewShardingNetworkPrivateAddress(ctx *pulumi.Context,
	name string, args *ShardingNetworkPrivateAddressArgs, opts ...pulumi.ResourceOption) (*ShardingNetworkPrivateAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'DbInstanceId'")
	}
	if args.NodeId == nil {
		return nil, errors.New("invalid value for required argument 'NodeId'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	if args.AccountPassword != nil {
		args.AccountPassword = pulumi.ToSecret(args.AccountPassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accountPassword",
	})
	opts = append(opts, secrets)
	var resource ShardingNetworkPrivateAddress
	err := ctx.RegisterResource("alicloud:mongodb/shardingNetworkPrivateAddress:ShardingNetworkPrivateAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetShardingNetworkPrivateAddress gets an existing ShardingNetworkPrivateAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShardingNetworkPrivateAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShardingNetworkPrivateAddressState, opts ...pulumi.ResourceOption) (*ShardingNetworkPrivateAddress, error) {
	var resource ShardingNetworkPrivateAddress
	err := ctx.ReadResource("alicloud:mongodb/shardingNetworkPrivateAddress:ShardingNetworkPrivateAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ShardingNetworkPrivateAddress resources.
type shardingNetworkPrivateAddressState struct {
	// The name of the account.
	// - The name must be 4 to 16 characters in length and can contain lowercase letters, digits, and underscores (_). It must start with a lowercase letter.
	// - You need to set the account name and password only when you apply for an endpoint for a shard or Configserver node for the first time. In this case, the account name and password are used for all shard and Configserver nodes.
	// - The permissions of this account are fixed to read-only.
	AccountName *string `pulumi:"accountName"`
	// Account password.
	// - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%^&*()_+-=`.
	// - The password must be 8 to 32 characters in length.
	AccountPassword *string `pulumi:"accountPassword"`
	// The db instance id.
	DbInstanceId *string `pulumi:"dbInstanceId"`
	// The endpoint of the instance.
	NetworkAddresses []ShardingNetworkPrivateAddressNetworkAddress `pulumi:"networkAddresses"`
	// The ID of the Shard node or the ConfigServer node.
	NodeId *string `pulumi:"nodeId"`
	// The zone ID of the instance.
	ZoneId *string `pulumi:"zoneId"`
}

type ShardingNetworkPrivateAddressState struct {
	// The name of the account.
	// - The name must be 4 to 16 characters in length and can contain lowercase letters, digits, and underscores (_). It must start with a lowercase letter.
	// - You need to set the account name and password only when you apply for an endpoint for a shard or Configserver node for the first time. In this case, the account name and password are used for all shard and Configserver nodes.
	// - The permissions of this account are fixed to read-only.
	AccountName pulumi.StringPtrInput
	// Account password.
	// - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%^&*()_+-=`.
	// - The password must be 8 to 32 characters in length.
	AccountPassword pulumi.StringPtrInput
	// The db instance id.
	DbInstanceId pulumi.StringPtrInput
	// The endpoint of the instance.
	NetworkAddresses ShardingNetworkPrivateAddressNetworkAddressArrayInput
	// The ID of the Shard node or the ConfigServer node.
	NodeId pulumi.StringPtrInput
	// The zone ID of the instance.
	ZoneId pulumi.StringPtrInput
}

func (ShardingNetworkPrivateAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*shardingNetworkPrivateAddressState)(nil)).Elem()
}

type shardingNetworkPrivateAddressArgs struct {
	// The name of the account.
	// - The name must be 4 to 16 characters in length and can contain lowercase letters, digits, and underscores (_). It must start with a lowercase letter.
	// - You need to set the account name and password only when you apply for an endpoint for a shard or Configserver node for the first time. In this case, the account name and password are used for all shard and Configserver nodes.
	// - The permissions of this account are fixed to read-only.
	AccountName *string `pulumi:"accountName"`
	// Account password.
	// - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%^&*()_+-=`.
	// - The password must be 8 to 32 characters in length.
	AccountPassword *string `pulumi:"accountPassword"`
	// The db instance id.
	DbInstanceId string `pulumi:"dbInstanceId"`
	// The ID of the Shard node or the ConfigServer node.
	NodeId string `pulumi:"nodeId"`
	// The zone ID of the instance.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a ShardingNetworkPrivateAddress resource.
type ShardingNetworkPrivateAddressArgs struct {
	// The name of the account.
	// - The name must be 4 to 16 characters in length and can contain lowercase letters, digits, and underscores (_). It must start with a lowercase letter.
	// - You need to set the account name and password only when you apply for an endpoint for a shard or Configserver node for the first time. In this case, the account name and password are used for all shard and Configserver nodes.
	// - The permissions of this account are fixed to read-only.
	AccountName pulumi.StringPtrInput
	// Account password.
	// - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%^&*()_+-=`.
	// - The password must be 8 to 32 characters in length.
	AccountPassword pulumi.StringPtrInput
	// The db instance id.
	DbInstanceId pulumi.StringInput
	// The ID of the Shard node or the ConfigServer node.
	NodeId pulumi.StringInput
	// The zone ID of the instance.
	ZoneId pulumi.StringInput
}

func (ShardingNetworkPrivateAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shardingNetworkPrivateAddressArgs)(nil)).Elem()
}

type ShardingNetworkPrivateAddressInput interface {
	pulumi.Input

	ToShardingNetworkPrivateAddressOutput() ShardingNetworkPrivateAddressOutput
	ToShardingNetworkPrivateAddressOutputWithContext(ctx context.Context) ShardingNetworkPrivateAddressOutput
}

func (*ShardingNetworkPrivateAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**ShardingNetworkPrivateAddress)(nil)).Elem()
}

func (i *ShardingNetworkPrivateAddress) ToShardingNetworkPrivateAddressOutput() ShardingNetworkPrivateAddressOutput {
	return i.ToShardingNetworkPrivateAddressOutputWithContext(context.Background())
}

func (i *ShardingNetworkPrivateAddress) ToShardingNetworkPrivateAddressOutputWithContext(ctx context.Context) ShardingNetworkPrivateAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShardingNetworkPrivateAddressOutput)
}

// ShardingNetworkPrivateAddressArrayInput is an input type that accepts ShardingNetworkPrivateAddressArray and ShardingNetworkPrivateAddressArrayOutput values.
// You can construct a concrete instance of `ShardingNetworkPrivateAddressArrayInput` via:
//
//	ShardingNetworkPrivateAddressArray{ ShardingNetworkPrivateAddressArgs{...} }
type ShardingNetworkPrivateAddressArrayInput interface {
	pulumi.Input

	ToShardingNetworkPrivateAddressArrayOutput() ShardingNetworkPrivateAddressArrayOutput
	ToShardingNetworkPrivateAddressArrayOutputWithContext(context.Context) ShardingNetworkPrivateAddressArrayOutput
}

type ShardingNetworkPrivateAddressArray []ShardingNetworkPrivateAddressInput

func (ShardingNetworkPrivateAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ShardingNetworkPrivateAddress)(nil)).Elem()
}

func (i ShardingNetworkPrivateAddressArray) ToShardingNetworkPrivateAddressArrayOutput() ShardingNetworkPrivateAddressArrayOutput {
	return i.ToShardingNetworkPrivateAddressArrayOutputWithContext(context.Background())
}

func (i ShardingNetworkPrivateAddressArray) ToShardingNetworkPrivateAddressArrayOutputWithContext(ctx context.Context) ShardingNetworkPrivateAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShardingNetworkPrivateAddressArrayOutput)
}

// ShardingNetworkPrivateAddressMapInput is an input type that accepts ShardingNetworkPrivateAddressMap and ShardingNetworkPrivateAddressMapOutput values.
// You can construct a concrete instance of `ShardingNetworkPrivateAddressMapInput` via:
//
//	ShardingNetworkPrivateAddressMap{ "key": ShardingNetworkPrivateAddressArgs{...} }
type ShardingNetworkPrivateAddressMapInput interface {
	pulumi.Input

	ToShardingNetworkPrivateAddressMapOutput() ShardingNetworkPrivateAddressMapOutput
	ToShardingNetworkPrivateAddressMapOutputWithContext(context.Context) ShardingNetworkPrivateAddressMapOutput
}

type ShardingNetworkPrivateAddressMap map[string]ShardingNetworkPrivateAddressInput

func (ShardingNetworkPrivateAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ShardingNetworkPrivateAddress)(nil)).Elem()
}

func (i ShardingNetworkPrivateAddressMap) ToShardingNetworkPrivateAddressMapOutput() ShardingNetworkPrivateAddressMapOutput {
	return i.ToShardingNetworkPrivateAddressMapOutputWithContext(context.Background())
}

func (i ShardingNetworkPrivateAddressMap) ToShardingNetworkPrivateAddressMapOutputWithContext(ctx context.Context) ShardingNetworkPrivateAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShardingNetworkPrivateAddressMapOutput)
}

type ShardingNetworkPrivateAddressOutput struct{ *pulumi.OutputState }

func (ShardingNetworkPrivateAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShardingNetworkPrivateAddress)(nil)).Elem()
}

func (o ShardingNetworkPrivateAddressOutput) ToShardingNetworkPrivateAddressOutput() ShardingNetworkPrivateAddressOutput {
	return o
}

func (o ShardingNetworkPrivateAddressOutput) ToShardingNetworkPrivateAddressOutputWithContext(ctx context.Context) ShardingNetworkPrivateAddressOutput {
	return o
}

// The name of the account.
// - The name must be 4 to 16 characters in length and can contain lowercase letters, digits, and underscores (_). It must start with a lowercase letter.
// - You need to set the account name and password only when you apply for an endpoint for a shard or Configserver node for the first time. In this case, the account name and password are used for all shard and Configserver nodes.
// - The permissions of this account are fixed to read-only.
func (o ShardingNetworkPrivateAddressOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShardingNetworkPrivateAddress) pulumi.StringPtrOutput { return v.AccountName }).(pulumi.StringPtrOutput)
}

// Account password.
// - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%^&*()_+-=`.
// - The password must be 8 to 32 characters in length.
func (o ShardingNetworkPrivateAddressOutput) AccountPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShardingNetworkPrivateAddress) pulumi.StringPtrOutput { return v.AccountPassword }).(pulumi.StringPtrOutput)
}

// The db instance id.
func (o ShardingNetworkPrivateAddressOutput) DbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ShardingNetworkPrivateAddress) pulumi.StringOutput { return v.DbInstanceId }).(pulumi.StringOutput)
}

// The endpoint of the instance.
func (o ShardingNetworkPrivateAddressOutput) NetworkAddresses() ShardingNetworkPrivateAddressNetworkAddressArrayOutput {
	return o.ApplyT(func(v *ShardingNetworkPrivateAddress) ShardingNetworkPrivateAddressNetworkAddressArrayOutput {
		return v.NetworkAddresses
	}).(ShardingNetworkPrivateAddressNetworkAddressArrayOutput)
}

// The ID of the Shard node or the ConfigServer node.
func (o ShardingNetworkPrivateAddressOutput) NodeId() pulumi.StringOutput {
	return o.ApplyT(func(v *ShardingNetworkPrivateAddress) pulumi.StringOutput { return v.NodeId }).(pulumi.StringOutput)
}

// The zone ID of the instance.
func (o ShardingNetworkPrivateAddressOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *ShardingNetworkPrivateAddress) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type ShardingNetworkPrivateAddressArrayOutput struct{ *pulumi.OutputState }

func (ShardingNetworkPrivateAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ShardingNetworkPrivateAddress)(nil)).Elem()
}

func (o ShardingNetworkPrivateAddressArrayOutput) ToShardingNetworkPrivateAddressArrayOutput() ShardingNetworkPrivateAddressArrayOutput {
	return o
}

func (o ShardingNetworkPrivateAddressArrayOutput) ToShardingNetworkPrivateAddressArrayOutputWithContext(ctx context.Context) ShardingNetworkPrivateAddressArrayOutput {
	return o
}

func (o ShardingNetworkPrivateAddressArrayOutput) Index(i pulumi.IntInput) ShardingNetworkPrivateAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ShardingNetworkPrivateAddress {
		return vs[0].([]*ShardingNetworkPrivateAddress)[vs[1].(int)]
	}).(ShardingNetworkPrivateAddressOutput)
}

type ShardingNetworkPrivateAddressMapOutput struct{ *pulumi.OutputState }

func (ShardingNetworkPrivateAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ShardingNetworkPrivateAddress)(nil)).Elem()
}

func (o ShardingNetworkPrivateAddressMapOutput) ToShardingNetworkPrivateAddressMapOutput() ShardingNetworkPrivateAddressMapOutput {
	return o
}

func (o ShardingNetworkPrivateAddressMapOutput) ToShardingNetworkPrivateAddressMapOutputWithContext(ctx context.Context) ShardingNetworkPrivateAddressMapOutput {
	return o
}

func (o ShardingNetworkPrivateAddressMapOutput) MapIndex(k pulumi.StringInput) ShardingNetworkPrivateAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ShardingNetworkPrivateAddress {
		return vs[0].(map[string]*ShardingNetworkPrivateAddress)[vs[1].(string)]
	}).(ShardingNetworkPrivateAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ShardingNetworkPrivateAddressInput)(nil)).Elem(), &ShardingNetworkPrivateAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShardingNetworkPrivateAddressArrayInput)(nil)).Elem(), ShardingNetworkPrivateAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShardingNetworkPrivateAddressMapInput)(nil)).Elem(), ShardingNetworkPrivateAddressMap{})
	pulumi.RegisterOutputType(ShardingNetworkPrivateAddressOutput{})
	pulumi.RegisterOutputType(ShardingNetworkPrivateAddressArrayOutput{})
	pulumi.RegisterOutputType(ShardingNetworkPrivateAddressMapOutput{})
}
