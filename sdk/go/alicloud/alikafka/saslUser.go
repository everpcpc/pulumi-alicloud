// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alikafka

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Alikafka sasl user resource.
//
// > **NOTE:** Available in 1.66.0+
//
// > **NOTE:**  Only the following regions support create alikafka sasl user.
// [`cn-hangzhou`,`cn-beijing`,`cn-shenzhen`,`cn-shanghai`,`cn-qingdao`,`cn-hongkong`,`cn-huhehaote`,`cn-zhangjiakou`,`cn-chengdu`,`cn-heyuan`,`ap-southeast-1`,`ap-southeast-3`,`ap-southeast-5`,`ap-south-1`,`ap-northeast-1`,`eu-central-1`,`eu-west-1`,`us-west-1`,`us-east-1`]
//
// For information about Alikafka sasl user and how to use it, see [What is Alikafka sasl user a](https://www.alibabacloud.com/help/en/doc-detail/162221.html)
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/alikafka"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			username := "testusername"
//			if param := cfg.Get("username"); param != "" {
//				username = param
//			}
//			password := "testpassword"
//			if param := cfg.Get("password"); param != "" {
//				password = param
//			}
//			defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
//				CidrBlock: pulumi.String("172.16.0.0/12"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
//				VpcId:     defaultNetwork.ID(),
//				CidrBlock: pulumi.String("172.16.0.0/24"),
//				ZoneId:    *pulumi.String(defaultZones.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := alikafka.NewInstance(ctx, "defaultInstance", &alikafka.InstanceArgs{
//				PartitionNum: pulumi.Int(50),
//				DiskType:     pulumi.Int(1),
//				DiskSize:     pulumi.Int(500),
//				DeployType:   pulumi.Int(5),
//				IoMax:        pulumi.Int(20),
//				VswitchId:    defaultSwitch.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = alikafka.NewSaslUser(ctx, "defaultSaslUser", &alikafka.SaslUserArgs{
//				InstanceId: defaultInstance.ID(),
//				Username:   pulumi.String(username),
//				Password:   pulumi.String(password),
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
// Alikafka Sasl User can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:alikafka/saslUser:SaslUser example <instance_id>:<username>
//
// ```
type SaslUser struct {
	pulumi.CustomResourceState

	// ID of the ALIKAFKA Instance that owns the groups.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// An KMS encrypts password used to a db account. You have to specify one of `password` and `kmsEncryptedPassword` fields.
	KmsEncryptedPassword pulumi.StringPtrOutput `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a user with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapOutput `pulumi:"kmsEncryptionContext"`
	// Operation password. It may consist of letters, digits, or underlines, with a length of 1 to 64 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The authentication mechanism. Valid values: `plain`, `scram`. Default value: `plain`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// Username for the sasl user. The length should between 1 to 64 characters. The characters can only contain 'a'-'z', 'A'-'Z', '0'-'9', '_' and '-'.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewSaslUser registers a new resource with the given unique name, arguments, and options.
func NewSaslUser(ctx *pulumi.Context,
	name string, args *SaslUserArgs, opts ...pulumi.ResourceOption) (*SaslUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	var resource SaslUser
	err := ctx.RegisterResource("alicloud:alikafka/saslUser:SaslUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSaslUser gets an existing SaslUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSaslUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SaslUserState, opts ...pulumi.ResourceOption) (*SaslUser, error) {
	var resource SaslUser
	err := ctx.ReadResource("alicloud:alikafka/saslUser:SaslUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SaslUser resources.
type saslUserState struct {
	// ID of the ALIKAFKA Instance that owns the groups.
	InstanceId *string `pulumi:"instanceId"`
	// An KMS encrypts password used to a db account. You have to specify one of `password` and `kmsEncryptedPassword` fields.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a user with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// Operation password. It may consist of letters, digits, or underlines, with a length of 1 to 64 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
	Password *string `pulumi:"password"`
	// The authentication mechanism. Valid values: `plain`, `scram`. Default value: `plain`.
	Type *string `pulumi:"type"`
	// Username for the sasl user. The length should between 1 to 64 characters. The characters can only contain 'a'-'z', 'A'-'Z', '0'-'9', '_' and '-'.
	Username *string `pulumi:"username"`
}

type SaslUserState struct {
	// ID of the ALIKAFKA Instance that owns the groups.
	InstanceId pulumi.StringPtrInput
	// An KMS encrypts password used to a db account. You have to specify one of `password` and `kmsEncryptedPassword` fields.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a user with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// Operation password. It may consist of letters, digits, or underlines, with a length of 1 to 64 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
	Password pulumi.StringPtrInput
	// The authentication mechanism. Valid values: `plain`, `scram`. Default value: `plain`.
	Type pulumi.StringPtrInput
	// Username for the sasl user. The length should between 1 to 64 characters. The characters can only contain 'a'-'z', 'A'-'Z', '0'-'9', '_' and '-'.
	Username pulumi.StringPtrInput
}

func (SaslUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*saslUserState)(nil)).Elem()
}

type saslUserArgs struct {
	// ID of the ALIKAFKA Instance that owns the groups.
	InstanceId string `pulumi:"instanceId"`
	// An KMS encrypts password used to a db account. You have to specify one of `password` and `kmsEncryptedPassword` fields.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a user with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// Operation password. It may consist of letters, digits, or underlines, with a length of 1 to 64 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
	Password *string `pulumi:"password"`
	// The authentication mechanism. Valid values: `plain`, `scram`. Default value: `plain`.
	Type *string `pulumi:"type"`
	// Username for the sasl user. The length should between 1 to 64 characters. The characters can only contain 'a'-'z', 'A'-'Z', '0'-'9', '_' and '-'.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a SaslUser resource.
type SaslUserArgs struct {
	// ID of the ALIKAFKA Instance that owns the groups.
	InstanceId pulumi.StringInput
	// An KMS encrypts password used to a db account. You have to specify one of `password` and `kmsEncryptedPassword` fields.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a user with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// Operation password. It may consist of letters, digits, or underlines, with a length of 1 to 64 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
	Password pulumi.StringPtrInput
	// The authentication mechanism. Valid values: `plain`, `scram`. Default value: `plain`.
	Type pulumi.StringPtrInput
	// Username for the sasl user. The length should between 1 to 64 characters. The characters can only contain 'a'-'z', 'A'-'Z', '0'-'9', '_' and '-'.
	Username pulumi.StringInput
}

func (SaslUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*saslUserArgs)(nil)).Elem()
}

type SaslUserInput interface {
	pulumi.Input

	ToSaslUserOutput() SaslUserOutput
	ToSaslUserOutputWithContext(ctx context.Context) SaslUserOutput
}

func (*SaslUser) ElementType() reflect.Type {
	return reflect.TypeOf((**SaslUser)(nil)).Elem()
}

func (i *SaslUser) ToSaslUserOutput() SaslUserOutput {
	return i.ToSaslUserOutputWithContext(context.Background())
}

func (i *SaslUser) ToSaslUserOutputWithContext(ctx context.Context) SaslUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SaslUserOutput)
}

// SaslUserArrayInput is an input type that accepts SaslUserArray and SaslUserArrayOutput values.
// You can construct a concrete instance of `SaslUserArrayInput` via:
//
//	SaslUserArray{ SaslUserArgs{...} }
type SaslUserArrayInput interface {
	pulumi.Input

	ToSaslUserArrayOutput() SaslUserArrayOutput
	ToSaslUserArrayOutputWithContext(context.Context) SaslUserArrayOutput
}

type SaslUserArray []SaslUserInput

func (SaslUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SaslUser)(nil)).Elem()
}

func (i SaslUserArray) ToSaslUserArrayOutput() SaslUserArrayOutput {
	return i.ToSaslUserArrayOutputWithContext(context.Background())
}

func (i SaslUserArray) ToSaslUserArrayOutputWithContext(ctx context.Context) SaslUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SaslUserArrayOutput)
}

// SaslUserMapInput is an input type that accepts SaslUserMap and SaslUserMapOutput values.
// You can construct a concrete instance of `SaslUserMapInput` via:
//
//	SaslUserMap{ "key": SaslUserArgs{...} }
type SaslUserMapInput interface {
	pulumi.Input

	ToSaslUserMapOutput() SaslUserMapOutput
	ToSaslUserMapOutputWithContext(context.Context) SaslUserMapOutput
}

type SaslUserMap map[string]SaslUserInput

func (SaslUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SaslUser)(nil)).Elem()
}

func (i SaslUserMap) ToSaslUserMapOutput() SaslUserMapOutput {
	return i.ToSaslUserMapOutputWithContext(context.Background())
}

func (i SaslUserMap) ToSaslUserMapOutputWithContext(ctx context.Context) SaslUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SaslUserMapOutput)
}

type SaslUserOutput struct{ *pulumi.OutputState }

func (SaslUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SaslUser)(nil)).Elem()
}

func (o SaslUserOutput) ToSaslUserOutput() SaslUserOutput {
	return o
}

func (o SaslUserOutput) ToSaslUserOutputWithContext(ctx context.Context) SaslUserOutput {
	return o
}

// ID of the ALIKAFKA Instance that owns the groups.
func (o SaslUserOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SaslUser) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// An KMS encrypts password used to a db account. You have to specify one of `password` and `kmsEncryptedPassword` fields.
func (o SaslUserOutput) KmsEncryptedPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SaslUser) pulumi.StringPtrOutput { return v.KmsEncryptedPassword }).(pulumi.StringPtrOutput)
}

// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a user with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
func (o SaslUserOutput) KmsEncryptionContext() pulumi.MapOutput {
	return o.ApplyT(func(v *SaslUser) pulumi.MapOutput { return v.KmsEncryptionContext }).(pulumi.MapOutput)
}

// Operation password. It may consist of letters, digits, or underlines, with a length of 1 to 64 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
func (o SaslUserOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SaslUser) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The authentication mechanism. Valid values: `plain`, `scram`. Default value: `plain`.
func (o SaslUserOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SaslUser) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// Username for the sasl user. The length should between 1 to 64 characters. The characters can only contain 'a'-'z', 'A'-'Z', '0'-'9', '_' and '-'.
func (o SaslUserOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *SaslUser) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type SaslUserArrayOutput struct{ *pulumi.OutputState }

func (SaslUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SaslUser)(nil)).Elem()
}

func (o SaslUserArrayOutput) ToSaslUserArrayOutput() SaslUserArrayOutput {
	return o
}

func (o SaslUserArrayOutput) ToSaslUserArrayOutputWithContext(ctx context.Context) SaslUserArrayOutput {
	return o
}

func (o SaslUserArrayOutput) Index(i pulumi.IntInput) SaslUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SaslUser {
		return vs[0].([]*SaslUser)[vs[1].(int)]
	}).(SaslUserOutput)
}

type SaslUserMapOutput struct{ *pulumi.OutputState }

func (SaslUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SaslUser)(nil)).Elem()
}

func (o SaslUserMapOutput) ToSaslUserMapOutput() SaslUserMapOutput {
	return o
}

func (o SaslUserMapOutput) ToSaslUserMapOutputWithContext(ctx context.Context) SaslUserMapOutput {
	return o
}

func (o SaslUserMapOutput) MapIndex(k pulumi.StringInput) SaslUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SaslUser {
		return vs[0].(map[string]*SaslUser)[vs[1].(string)]
	}).(SaslUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SaslUserInput)(nil)).Elem(), &SaslUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*SaslUserArrayInput)(nil)).Elem(), SaslUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SaslUserMapInput)(nil)).Elem(), SaslUserMap{})
	pulumi.RegisterOutputType(SaslUserOutput{})
	pulumi.RegisterOutputType(SaslUserArrayOutput{})
	pulumi.RegisterOutputType(SaslUserMapOutput{})
}
