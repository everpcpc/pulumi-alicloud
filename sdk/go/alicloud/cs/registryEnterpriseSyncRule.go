// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource will help you to manager Container Registry Enterprise Edition sync rules.
//
// For information about Container Registry Enterprise Edition sync rules and how to use it, see [Create a Sync Rule](https://www.alibabacloud.com/help/doc-detail/145280.htm)
//
// > **NOTE:** Available in v1.90.0+.
//
// > **NOTE:** You need to set your registry password in Container Registry Enterprise Edition console before use this resource.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cs.NewRegistryEnterpriseSyncRule(ctx, "_default", &cs.RegistryEnterpriseSyncRuleArgs{
// 			InstanceId:          pulumi.String("my-source-instance-id"),
// 			NamespaceName:       pulumi.String("my-source-namespace"),
// 			RepoName:            pulumi.String("my-source-repo"),
// 			TagFilter:           pulumi.String(".*"),
// 			TargetInstanceId:    pulumi.String("my-target-instance-id"),
// 			TargetNamespaceName: pulumi.String("my-target-namespace"),
// 			TargetRegionId:      pulumi.String("cn-hangzhou"),
// 			TargetRepoName:      pulumi.String("my-target-repo"),
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
// Container Registry Enterprise Edition sync rule can be imported using the id. Format to `{instance_id}:{namespace_name}:{rule_id}`, e.g.
//
// ```sh
//  $ pulumi import alicloud:cs/registryEnterpriseSyncRule:RegistryEnterpriseSyncRule default `cri-xxx:my-namespace:crsr-yyy`
// ```
type RegistryEnterpriseSyncRule struct {
	pulumi.CustomResourceState

	// ID of Container Registry Enterprise Edition source instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Name of Container Registry Enterprise Edition sync rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// Name of the source repository which should be set together with `targetRepoName`, if empty means that the synchronization scope is the entire namespace level.
	RepoName pulumi.StringPtrOutput `pulumi:"repoName"`
	// The uuid of Container Registry Enterprise Edition sync rule.
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
	// `FROM` or `TO`, the direction of synchronization. `FROM` means source instance, `TO` means target instance.
	SyncDirection pulumi.StringOutput `pulumi:"syncDirection"`
	// `REPO` or `NAMESPACE`,the scope that the synchronization rule applies.
	SyncScope pulumi.StringOutput `pulumi:"syncScope"`
	// The regular expression used to filter image tags for synchronization in the source repository.
	TagFilter pulumi.StringOutput `pulumi:"tagFilter"`
	// ID of Container Registry Enterprise Edition target instance to be synchronized.
	TargetInstanceId pulumi.StringOutput `pulumi:"targetInstanceId"`
	// Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
	TargetNamespaceName pulumi.StringOutput `pulumi:"targetNamespaceName"`
	// The target region to be synchronized.
	TargetRegionId pulumi.StringOutput `pulumi:"targetRegionId"`
	// Name of the target repository.
	TargetRepoName pulumi.StringPtrOutput `pulumi:"targetRepoName"`
}

// NewRegistryEnterpriseSyncRule registers a new resource with the given unique name, arguments, and options.
func NewRegistryEnterpriseSyncRule(ctx *pulumi.Context,
	name string, args *RegistryEnterpriseSyncRuleArgs, opts ...pulumi.ResourceOption) (*RegistryEnterpriseSyncRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.TagFilter == nil {
		return nil, errors.New("invalid value for required argument 'TagFilter'")
	}
	if args.TargetInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'TargetInstanceId'")
	}
	if args.TargetNamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'TargetNamespaceName'")
	}
	if args.TargetRegionId == nil {
		return nil, errors.New("invalid value for required argument 'TargetRegionId'")
	}
	var resource RegistryEnterpriseSyncRule
	err := ctx.RegisterResource("alicloud:cs/registryEnterpriseSyncRule:RegistryEnterpriseSyncRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryEnterpriseSyncRule gets an existing RegistryEnterpriseSyncRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryEnterpriseSyncRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryEnterpriseSyncRuleState, opts ...pulumi.ResourceOption) (*RegistryEnterpriseSyncRule, error) {
	var resource RegistryEnterpriseSyncRule
	err := ctx.ReadResource("alicloud:cs/registryEnterpriseSyncRule:RegistryEnterpriseSyncRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryEnterpriseSyncRule resources.
type registryEnterpriseSyncRuleState struct {
	// ID of Container Registry Enterprise Edition source instance.
	InstanceId *string `pulumi:"instanceId"`
	// Name of Container Registry Enterprise Edition sync rule.
	Name *string `pulumi:"name"`
	// Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
	NamespaceName *string `pulumi:"namespaceName"`
	// Name of the source repository which should be set together with `targetRepoName`, if empty means that the synchronization scope is the entire namespace level.
	RepoName *string `pulumi:"repoName"`
	// The uuid of Container Registry Enterprise Edition sync rule.
	RuleId *string `pulumi:"ruleId"`
	// `FROM` or `TO`, the direction of synchronization. `FROM` means source instance, `TO` means target instance.
	SyncDirection *string `pulumi:"syncDirection"`
	// `REPO` or `NAMESPACE`,the scope that the synchronization rule applies.
	SyncScope *string `pulumi:"syncScope"`
	// The regular expression used to filter image tags for synchronization in the source repository.
	TagFilter *string `pulumi:"tagFilter"`
	// ID of Container Registry Enterprise Edition target instance to be synchronized.
	TargetInstanceId *string `pulumi:"targetInstanceId"`
	// Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
	TargetNamespaceName *string `pulumi:"targetNamespaceName"`
	// The target region to be synchronized.
	TargetRegionId *string `pulumi:"targetRegionId"`
	// Name of the target repository.
	TargetRepoName *string `pulumi:"targetRepoName"`
}

type RegistryEnterpriseSyncRuleState struct {
	// ID of Container Registry Enterprise Edition source instance.
	InstanceId pulumi.StringPtrInput
	// Name of Container Registry Enterprise Edition sync rule.
	Name pulumi.StringPtrInput
	// Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
	NamespaceName pulumi.StringPtrInput
	// Name of the source repository which should be set together with `targetRepoName`, if empty means that the synchronization scope is the entire namespace level.
	RepoName pulumi.StringPtrInput
	// The uuid of Container Registry Enterprise Edition sync rule.
	RuleId pulumi.StringPtrInput
	// `FROM` or `TO`, the direction of synchronization. `FROM` means source instance, `TO` means target instance.
	SyncDirection pulumi.StringPtrInput
	// `REPO` or `NAMESPACE`,the scope that the synchronization rule applies.
	SyncScope pulumi.StringPtrInput
	// The regular expression used to filter image tags for synchronization in the source repository.
	TagFilter pulumi.StringPtrInput
	// ID of Container Registry Enterprise Edition target instance to be synchronized.
	TargetInstanceId pulumi.StringPtrInput
	// Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
	TargetNamespaceName pulumi.StringPtrInput
	// The target region to be synchronized.
	TargetRegionId pulumi.StringPtrInput
	// Name of the target repository.
	TargetRepoName pulumi.StringPtrInput
}

func (RegistryEnterpriseSyncRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryEnterpriseSyncRuleState)(nil)).Elem()
}

type registryEnterpriseSyncRuleArgs struct {
	// ID of Container Registry Enterprise Edition source instance.
	InstanceId string `pulumi:"instanceId"`
	// Name of Container Registry Enterprise Edition sync rule.
	Name *string `pulumi:"name"`
	// Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the source repository which should be set together with `targetRepoName`, if empty means that the synchronization scope is the entire namespace level.
	RepoName *string `pulumi:"repoName"`
	// The regular expression used to filter image tags for synchronization in the source repository.
	TagFilter string `pulumi:"tagFilter"`
	// ID of Container Registry Enterprise Edition target instance to be synchronized.
	TargetInstanceId string `pulumi:"targetInstanceId"`
	// Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
	TargetNamespaceName string `pulumi:"targetNamespaceName"`
	// The target region to be synchronized.
	TargetRegionId string `pulumi:"targetRegionId"`
	// Name of the target repository.
	TargetRepoName *string `pulumi:"targetRepoName"`
}

// The set of arguments for constructing a RegistryEnterpriseSyncRule resource.
type RegistryEnterpriseSyncRuleArgs struct {
	// ID of Container Registry Enterprise Edition source instance.
	InstanceId pulumi.StringInput
	// Name of Container Registry Enterprise Edition sync rule.
	Name pulumi.StringPtrInput
	// Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
	NamespaceName pulumi.StringInput
	// Name of the source repository which should be set together with `targetRepoName`, if empty means that the synchronization scope is the entire namespace level.
	RepoName pulumi.StringPtrInput
	// The regular expression used to filter image tags for synchronization in the source repository.
	TagFilter pulumi.StringInput
	// ID of Container Registry Enterprise Edition target instance to be synchronized.
	TargetInstanceId pulumi.StringInput
	// Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
	TargetNamespaceName pulumi.StringInput
	// The target region to be synchronized.
	TargetRegionId pulumi.StringInput
	// Name of the target repository.
	TargetRepoName pulumi.StringPtrInput
}

func (RegistryEnterpriseSyncRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryEnterpriseSyncRuleArgs)(nil)).Elem()
}

type RegistryEnterpriseSyncRuleInput interface {
	pulumi.Input

	ToRegistryEnterpriseSyncRuleOutput() RegistryEnterpriseSyncRuleOutput
	ToRegistryEnterpriseSyncRuleOutputWithContext(ctx context.Context) RegistryEnterpriseSyncRuleOutput
}

func (*RegistryEnterpriseSyncRule) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryEnterpriseSyncRule)(nil))
}

func (i *RegistryEnterpriseSyncRule) ToRegistryEnterpriseSyncRuleOutput() RegistryEnterpriseSyncRuleOutput {
	return i.ToRegistryEnterpriseSyncRuleOutputWithContext(context.Background())
}

func (i *RegistryEnterpriseSyncRule) ToRegistryEnterpriseSyncRuleOutputWithContext(ctx context.Context) RegistryEnterpriseSyncRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryEnterpriseSyncRuleOutput)
}

func (i *RegistryEnterpriseSyncRule) ToRegistryEnterpriseSyncRulePtrOutput() RegistryEnterpriseSyncRulePtrOutput {
	return i.ToRegistryEnterpriseSyncRulePtrOutputWithContext(context.Background())
}

func (i *RegistryEnterpriseSyncRule) ToRegistryEnterpriseSyncRulePtrOutputWithContext(ctx context.Context) RegistryEnterpriseSyncRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryEnterpriseSyncRulePtrOutput)
}

type RegistryEnterpriseSyncRulePtrInput interface {
	pulumi.Input

	ToRegistryEnterpriseSyncRulePtrOutput() RegistryEnterpriseSyncRulePtrOutput
	ToRegistryEnterpriseSyncRulePtrOutputWithContext(ctx context.Context) RegistryEnterpriseSyncRulePtrOutput
}

type registryEnterpriseSyncRulePtrType RegistryEnterpriseSyncRuleArgs

func (*registryEnterpriseSyncRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryEnterpriseSyncRule)(nil))
}

func (i *registryEnterpriseSyncRulePtrType) ToRegistryEnterpriseSyncRulePtrOutput() RegistryEnterpriseSyncRulePtrOutput {
	return i.ToRegistryEnterpriseSyncRulePtrOutputWithContext(context.Background())
}

func (i *registryEnterpriseSyncRulePtrType) ToRegistryEnterpriseSyncRulePtrOutputWithContext(ctx context.Context) RegistryEnterpriseSyncRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryEnterpriseSyncRulePtrOutput)
}

// RegistryEnterpriseSyncRuleArrayInput is an input type that accepts RegistryEnterpriseSyncRuleArray and RegistryEnterpriseSyncRuleArrayOutput values.
// You can construct a concrete instance of `RegistryEnterpriseSyncRuleArrayInput` via:
//
//          RegistryEnterpriseSyncRuleArray{ RegistryEnterpriseSyncRuleArgs{...} }
type RegistryEnterpriseSyncRuleArrayInput interface {
	pulumi.Input

	ToRegistryEnterpriseSyncRuleArrayOutput() RegistryEnterpriseSyncRuleArrayOutput
	ToRegistryEnterpriseSyncRuleArrayOutputWithContext(context.Context) RegistryEnterpriseSyncRuleArrayOutput
}

type RegistryEnterpriseSyncRuleArray []RegistryEnterpriseSyncRuleInput

func (RegistryEnterpriseSyncRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegistryEnterpriseSyncRule)(nil)).Elem()
}

func (i RegistryEnterpriseSyncRuleArray) ToRegistryEnterpriseSyncRuleArrayOutput() RegistryEnterpriseSyncRuleArrayOutput {
	return i.ToRegistryEnterpriseSyncRuleArrayOutputWithContext(context.Background())
}

func (i RegistryEnterpriseSyncRuleArray) ToRegistryEnterpriseSyncRuleArrayOutputWithContext(ctx context.Context) RegistryEnterpriseSyncRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryEnterpriseSyncRuleArrayOutput)
}

// RegistryEnterpriseSyncRuleMapInput is an input type that accepts RegistryEnterpriseSyncRuleMap and RegistryEnterpriseSyncRuleMapOutput values.
// You can construct a concrete instance of `RegistryEnterpriseSyncRuleMapInput` via:
//
//          RegistryEnterpriseSyncRuleMap{ "key": RegistryEnterpriseSyncRuleArgs{...} }
type RegistryEnterpriseSyncRuleMapInput interface {
	pulumi.Input

	ToRegistryEnterpriseSyncRuleMapOutput() RegistryEnterpriseSyncRuleMapOutput
	ToRegistryEnterpriseSyncRuleMapOutputWithContext(context.Context) RegistryEnterpriseSyncRuleMapOutput
}

type RegistryEnterpriseSyncRuleMap map[string]RegistryEnterpriseSyncRuleInput

func (RegistryEnterpriseSyncRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegistryEnterpriseSyncRule)(nil)).Elem()
}

func (i RegistryEnterpriseSyncRuleMap) ToRegistryEnterpriseSyncRuleMapOutput() RegistryEnterpriseSyncRuleMapOutput {
	return i.ToRegistryEnterpriseSyncRuleMapOutputWithContext(context.Background())
}

func (i RegistryEnterpriseSyncRuleMap) ToRegistryEnterpriseSyncRuleMapOutputWithContext(ctx context.Context) RegistryEnterpriseSyncRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryEnterpriseSyncRuleMapOutput)
}

type RegistryEnterpriseSyncRuleOutput struct{ *pulumi.OutputState }

func (RegistryEnterpriseSyncRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryEnterpriseSyncRule)(nil))
}

func (o RegistryEnterpriseSyncRuleOutput) ToRegistryEnterpriseSyncRuleOutput() RegistryEnterpriseSyncRuleOutput {
	return o
}

func (o RegistryEnterpriseSyncRuleOutput) ToRegistryEnterpriseSyncRuleOutputWithContext(ctx context.Context) RegistryEnterpriseSyncRuleOutput {
	return o
}

func (o RegistryEnterpriseSyncRuleOutput) ToRegistryEnterpriseSyncRulePtrOutput() RegistryEnterpriseSyncRulePtrOutput {
	return o.ToRegistryEnterpriseSyncRulePtrOutputWithContext(context.Background())
}

func (o RegistryEnterpriseSyncRuleOutput) ToRegistryEnterpriseSyncRulePtrOutputWithContext(ctx context.Context) RegistryEnterpriseSyncRulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RegistryEnterpriseSyncRule) *RegistryEnterpriseSyncRule {
		return &v
	}).(RegistryEnterpriseSyncRulePtrOutput)
}

type RegistryEnterpriseSyncRulePtrOutput struct{ *pulumi.OutputState }

func (RegistryEnterpriseSyncRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryEnterpriseSyncRule)(nil))
}

func (o RegistryEnterpriseSyncRulePtrOutput) ToRegistryEnterpriseSyncRulePtrOutput() RegistryEnterpriseSyncRulePtrOutput {
	return o
}

func (o RegistryEnterpriseSyncRulePtrOutput) ToRegistryEnterpriseSyncRulePtrOutputWithContext(ctx context.Context) RegistryEnterpriseSyncRulePtrOutput {
	return o
}

func (o RegistryEnterpriseSyncRulePtrOutput) Elem() RegistryEnterpriseSyncRuleOutput {
	return o.ApplyT(func(v *RegistryEnterpriseSyncRule) RegistryEnterpriseSyncRule {
		if v != nil {
			return *v
		}
		var ret RegistryEnterpriseSyncRule
		return ret
	}).(RegistryEnterpriseSyncRuleOutput)
}

type RegistryEnterpriseSyncRuleArrayOutput struct{ *pulumi.OutputState }

func (RegistryEnterpriseSyncRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegistryEnterpriseSyncRule)(nil))
}

func (o RegistryEnterpriseSyncRuleArrayOutput) ToRegistryEnterpriseSyncRuleArrayOutput() RegistryEnterpriseSyncRuleArrayOutput {
	return o
}

func (o RegistryEnterpriseSyncRuleArrayOutput) ToRegistryEnterpriseSyncRuleArrayOutputWithContext(ctx context.Context) RegistryEnterpriseSyncRuleArrayOutput {
	return o
}

func (o RegistryEnterpriseSyncRuleArrayOutput) Index(i pulumi.IntInput) RegistryEnterpriseSyncRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RegistryEnterpriseSyncRule {
		return vs[0].([]RegistryEnterpriseSyncRule)[vs[1].(int)]
	}).(RegistryEnterpriseSyncRuleOutput)
}

type RegistryEnterpriseSyncRuleMapOutput struct{ *pulumi.OutputState }

func (RegistryEnterpriseSyncRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RegistryEnterpriseSyncRule)(nil))
}

func (o RegistryEnterpriseSyncRuleMapOutput) ToRegistryEnterpriseSyncRuleMapOutput() RegistryEnterpriseSyncRuleMapOutput {
	return o
}

func (o RegistryEnterpriseSyncRuleMapOutput) ToRegistryEnterpriseSyncRuleMapOutputWithContext(ctx context.Context) RegistryEnterpriseSyncRuleMapOutput {
	return o
}

func (o RegistryEnterpriseSyncRuleMapOutput) MapIndex(k pulumi.StringInput) RegistryEnterpriseSyncRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RegistryEnterpriseSyncRule {
		return vs[0].(map[string]RegistryEnterpriseSyncRule)[vs[1].(string)]
	}).(RegistryEnterpriseSyncRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryEnterpriseSyncRuleInput)(nil)).Elem(), &RegistryEnterpriseSyncRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryEnterpriseSyncRulePtrInput)(nil)).Elem(), &RegistryEnterpriseSyncRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryEnterpriseSyncRuleArrayInput)(nil)).Elem(), RegistryEnterpriseSyncRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryEnterpriseSyncRuleMapInput)(nil)).Elem(), RegistryEnterpriseSyncRuleMap{})
	pulumi.RegisterOutputType(RegistryEnterpriseSyncRuleOutput{})
	pulumi.RegisterOutputType(RegistryEnterpriseSyncRulePtrOutput{})
	pulumi.RegisterOutputType(RegistryEnterpriseSyncRuleArrayOutput{})
	pulumi.RegisterOutputType(RegistryEnterpriseSyncRuleMapOutput{})
}
