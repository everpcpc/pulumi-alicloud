// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package log

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// SLS log audit exists in the form of log service app.
//
// In addition to inheriting all SLS functions, it also enhances the real-time automatic centralized collection of audit related logs across multi cloud products under multi accounts, and provides support for storage, query and information summary required by audit. It covers actiontrail, OSS, NAS, SLB, API gateway, RDS, WAF, cloud firewall, cloud security center and other products.
//
// > **NOTE:** Available in 1.81.0
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/log"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := log.NewAudit(ctx, "example", &log.AuditArgs{
// 			Aliuid:      pulumi.String("12345678"),
// 			DisplayName: pulumi.String("tf-audit-test"),
// 			VariableMap: pulumi.StringMap{
// 				"actiontrail_enabled": pulumi.String("true"),
// 				"actiontrail_ttl":     pulumi.String("180"),
// 				"oss_access_enabled":  pulumi.String("true"),
// 				"oss_access_ttl":      pulumi.String("180"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// Multiple accounts Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/log"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := log.NewAudit(ctx, "example", &log.AuditArgs{
// 			Aliuid:      pulumi.String("12345678"),
// 			DisplayName: pulumi.String("tf-audit-test"),
// 			MultiAccounts: pulumi.StringArray{
// 				pulumi.String("123456789123"),
// 				pulumi.String("12345678912300123"),
// 			},
// 			VariableMap: pulumi.StringMap{
// 				"actiontrail_enabled": pulumi.String("true"),
// 				"actiontrail_ttl":     pulumi.String("180"),
// 				"oss_access_enabled":  pulumi.String("true"),
// 				"oss_access_ttl":      pulumi.String("180"),
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
// Log alert can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:log/audit:Audit example tf-audit-test
// ```
type Audit struct {
	pulumi.CustomResourceState

	// Aliuid value of your account.
	Aliuid pulumi.StringOutput `pulumi:"aliuid"`
	// Name of SLS log audit.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Multi-account configuration, please fill in multiple aliuid.
	MultiAccounts pulumi.StringArrayOutput `pulumi:"multiAccounts"`
	// Log audit detailed configuration.
	VariableMap pulumi.MapOutput `pulumi:"variableMap"`
}

// NewAudit registers a new resource with the given unique name, arguments, and options.
func NewAudit(ctx *pulumi.Context,
	name string, args *AuditArgs, opts ...pulumi.ResourceOption) (*Audit, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Aliuid == nil {
		return nil, errors.New("invalid value for required argument 'Aliuid'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource Audit
	err := ctx.RegisterResource("alicloud:log/audit:Audit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAudit gets an existing Audit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAudit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuditState, opts ...pulumi.ResourceOption) (*Audit, error) {
	var resource Audit
	err := ctx.ReadResource("alicloud:log/audit:Audit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Audit resources.
type auditState struct {
	// Aliuid value of your account.
	Aliuid *string `pulumi:"aliuid"`
	// Name of SLS log audit.
	DisplayName *string `pulumi:"displayName"`
	// Multi-account configuration, please fill in multiple aliuid.
	MultiAccounts []string `pulumi:"multiAccounts"`
	// Log audit detailed configuration.
	VariableMap map[string]interface{} `pulumi:"variableMap"`
}

type AuditState struct {
	// Aliuid value of your account.
	Aliuid pulumi.StringPtrInput
	// Name of SLS log audit.
	DisplayName pulumi.StringPtrInput
	// Multi-account configuration, please fill in multiple aliuid.
	MultiAccounts pulumi.StringArrayInput
	// Log audit detailed configuration.
	VariableMap pulumi.MapInput
}

func (AuditState) ElementType() reflect.Type {
	return reflect.TypeOf((*auditState)(nil)).Elem()
}

type auditArgs struct {
	// Aliuid value of your account.
	Aliuid string `pulumi:"aliuid"`
	// Name of SLS log audit.
	DisplayName string `pulumi:"displayName"`
	// Multi-account configuration, please fill in multiple aliuid.
	MultiAccounts []string `pulumi:"multiAccounts"`
	// Log audit detailed configuration.
	VariableMap map[string]interface{} `pulumi:"variableMap"`
}

// The set of arguments for constructing a Audit resource.
type AuditArgs struct {
	// Aliuid value of your account.
	Aliuid pulumi.StringInput
	// Name of SLS log audit.
	DisplayName pulumi.StringInput
	// Multi-account configuration, please fill in multiple aliuid.
	MultiAccounts pulumi.StringArrayInput
	// Log audit detailed configuration.
	VariableMap pulumi.MapInput
}

func (AuditArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*auditArgs)(nil)).Elem()
}

type AuditInput interface {
	pulumi.Input

	ToAuditOutput() AuditOutput
	ToAuditOutputWithContext(ctx context.Context) AuditOutput
}

func (*Audit) ElementType() reflect.Type {
	return reflect.TypeOf((*Audit)(nil))
}

func (i *Audit) ToAuditOutput() AuditOutput {
	return i.ToAuditOutputWithContext(context.Background())
}

func (i *Audit) ToAuditOutputWithContext(ctx context.Context) AuditOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditOutput)
}

func (i *Audit) ToAuditPtrOutput() AuditPtrOutput {
	return i.ToAuditPtrOutputWithContext(context.Background())
}

func (i *Audit) ToAuditPtrOutputWithContext(ctx context.Context) AuditPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditPtrOutput)
}

type AuditPtrInput interface {
	pulumi.Input

	ToAuditPtrOutput() AuditPtrOutput
	ToAuditPtrOutputWithContext(ctx context.Context) AuditPtrOutput
}

type auditPtrType AuditArgs

func (*auditPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Audit)(nil))
}

func (i *auditPtrType) ToAuditPtrOutput() AuditPtrOutput {
	return i.ToAuditPtrOutputWithContext(context.Background())
}

func (i *auditPtrType) ToAuditPtrOutputWithContext(ctx context.Context) AuditPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditPtrOutput)
}

// AuditArrayInput is an input type that accepts AuditArray and AuditArrayOutput values.
// You can construct a concrete instance of `AuditArrayInput` via:
//
//          AuditArray{ AuditArgs{...} }
type AuditArrayInput interface {
	pulumi.Input

	ToAuditArrayOutput() AuditArrayOutput
	ToAuditArrayOutputWithContext(context.Context) AuditArrayOutput
}

type AuditArray []AuditInput

func (AuditArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Audit)(nil))
}

func (i AuditArray) ToAuditArrayOutput() AuditArrayOutput {
	return i.ToAuditArrayOutputWithContext(context.Background())
}

func (i AuditArray) ToAuditArrayOutputWithContext(ctx context.Context) AuditArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditArrayOutput)
}

// AuditMapInput is an input type that accepts AuditMap and AuditMapOutput values.
// You can construct a concrete instance of `AuditMapInput` via:
//
//          AuditMap{ "key": AuditArgs{...} }
type AuditMapInput interface {
	pulumi.Input

	ToAuditMapOutput() AuditMapOutput
	ToAuditMapOutputWithContext(context.Context) AuditMapOutput
}

type AuditMap map[string]AuditInput

func (AuditMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Audit)(nil))
}

func (i AuditMap) ToAuditMapOutput() AuditMapOutput {
	return i.ToAuditMapOutputWithContext(context.Background())
}

func (i AuditMap) ToAuditMapOutputWithContext(ctx context.Context) AuditMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditMapOutput)
}

type AuditOutput struct {
	*pulumi.OutputState
}

func (AuditOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Audit)(nil))
}

func (o AuditOutput) ToAuditOutput() AuditOutput {
	return o
}

func (o AuditOutput) ToAuditOutputWithContext(ctx context.Context) AuditOutput {
	return o
}

func (o AuditOutput) ToAuditPtrOutput() AuditPtrOutput {
	return o.ToAuditPtrOutputWithContext(context.Background())
}

func (o AuditOutput) ToAuditPtrOutputWithContext(ctx context.Context) AuditPtrOutput {
	return o.ApplyT(func(v Audit) *Audit {
		return &v
	}).(AuditPtrOutput)
}

type AuditPtrOutput struct {
	*pulumi.OutputState
}

func (AuditPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Audit)(nil))
}

func (o AuditPtrOutput) ToAuditPtrOutput() AuditPtrOutput {
	return o
}

func (o AuditPtrOutput) ToAuditPtrOutputWithContext(ctx context.Context) AuditPtrOutput {
	return o
}

type AuditArrayOutput struct{ *pulumi.OutputState }

func (AuditArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Audit)(nil))
}

func (o AuditArrayOutput) ToAuditArrayOutput() AuditArrayOutput {
	return o
}

func (o AuditArrayOutput) ToAuditArrayOutputWithContext(ctx context.Context) AuditArrayOutput {
	return o
}

func (o AuditArrayOutput) Index(i pulumi.IntInput) AuditOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Audit {
		return vs[0].([]Audit)[vs[1].(int)]
	}).(AuditOutput)
}

type AuditMapOutput struct{ *pulumi.OutputState }

func (AuditMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Audit)(nil))
}

func (o AuditMapOutput) ToAuditMapOutput() AuditMapOutput {
	return o
}

func (o AuditMapOutput) ToAuditMapOutputWithContext(ctx context.Context) AuditMapOutput {
	return o
}

func (o AuditMapOutput) MapIndex(k pulumi.StringInput) AuditOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Audit {
		return vs[0].(map[string]Audit)[vs[1].(string)]
	}).(AuditOutput)
}

func init() {
	pulumi.RegisterOutputType(AuditOutput{})
	pulumi.RegisterOutputType(AuditPtrOutput{})
	pulumi.RegisterOutputType(AuditArrayOutput{})
	pulumi.RegisterOutputType(AuditMapOutput{})
}
