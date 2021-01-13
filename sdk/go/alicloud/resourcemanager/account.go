// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package resourcemanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Resource Manager Account resource. Member accounts are containers for resources in a resource directory. These accounts isolate resources and serve as organizational units in the resource directory. You can create member accounts in a folder and then manage them in a unified manner.
// For information about Resource Manager Account and how to use it, see [What is Resource Manager Account](https://www.alibabacloud.com/help/en/doc-detail/111231.htm).
//
// > **NOTE:** Available in v1.83.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/resourcemanager"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		f1, err := resourcemanager.NewFolder(ctx, "f1", &resourcemanager.FolderArgs{
// 			FolderName: pulumi.String("test1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = resourcemanager.NewAccount(ctx, "example", &resourcemanager.AccountArgs{
// 			DisplayName: pulumi.String("RDAccount"),
// 			FolderId:    f1.ID(),
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
// Resource Manager Account can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:resourcemanager/account:Account example 13148890145*****
// ```
type Account struct {
	pulumi.CustomResourceState

	// Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The ID of the parent folder.
	FolderId pulumi.StringPtrOutput `pulumi:"folderId"`
	// Ways for members to join the resource directory. Valid values: `invited`, `created`.
	JoinMethod pulumi.StringOutput `pulumi:"joinMethod"`
	// The time when the member joined the resource directory.
	JoinTime pulumi.StringOutput `pulumi:"joinTime"`
	// The modification time of the invitation.
	ModifyTime pulumi.StringOutput `pulumi:"modifyTime"`
	// Settlement account ID. If the value is empty, the current account will be used for settlement.
	PayerAccountId pulumi.StringPtrOutput `pulumi:"payerAccountId"`
	// Resource directory ID.
	ResourceDirectoryId pulumi.StringOutput `pulumi:"resourceDirectoryId"`
	// Member joining status. Valid values: `CreateSuccess`,`CreateVerifying`,`CreateFailed`,`CreateExpired`,`CreateCancelled`,`PromoteVerifying`,`PromoteFailed`,`PromoteExpired`,`PromoteCancelled`,`PromoteSuccess`,`InviteSuccess`,`Removed`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Member type. The value of `ResourceAccount` indicates the resource account.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource Account
	err := ctx.RegisterResource("alicloud:resourcemanager/account:Account", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:resourcemanager/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
	DisplayName *string `pulumi:"displayName"`
	// The ID of the parent folder.
	FolderId *string `pulumi:"folderId"`
	// Ways for members to join the resource directory. Valid values: `invited`, `created`.
	JoinMethod *string `pulumi:"joinMethod"`
	// The time when the member joined the resource directory.
	JoinTime *string `pulumi:"joinTime"`
	// The modification time of the invitation.
	ModifyTime *string `pulumi:"modifyTime"`
	// Settlement account ID. If the value is empty, the current account will be used for settlement.
	PayerAccountId *string `pulumi:"payerAccountId"`
	// Resource directory ID.
	ResourceDirectoryId *string `pulumi:"resourceDirectoryId"`
	// Member joining status. Valid values: `CreateSuccess`,`CreateVerifying`,`CreateFailed`,`CreateExpired`,`CreateCancelled`,`PromoteVerifying`,`PromoteFailed`,`PromoteExpired`,`PromoteCancelled`,`PromoteSuccess`,`InviteSuccess`,`Removed`.
	Status *string `pulumi:"status"`
	// Member type. The value of `ResourceAccount` indicates the resource account.
	Type *string `pulumi:"type"`
}

type AccountState struct {
	// Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
	DisplayName pulumi.StringPtrInput
	// The ID of the parent folder.
	FolderId pulumi.StringPtrInput
	// Ways for members to join the resource directory. Valid values: `invited`, `created`.
	JoinMethod pulumi.StringPtrInput
	// The time when the member joined the resource directory.
	JoinTime pulumi.StringPtrInput
	// The modification time of the invitation.
	ModifyTime pulumi.StringPtrInput
	// Settlement account ID. If the value is empty, the current account will be used for settlement.
	PayerAccountId pulumi.StringPtrInput
	// Resource directory ID.
	ResourceDirectoryId pulumi.StringPtrInput
	// Member joining status. Valid values: `CreateSuccess`,`CreateVerifying`,`CreateFailed`,`CreateExpired`,`CreateCancelled`,`PromoteVerifying`,`PromoteFailed`,`PromoteExpired`,`PromoteCancelled`,`PromoteSuccess`,`InviteSuccess`,`Removed`.
	Status pulumi.StringPtrInput
	// Member type. The value of `ResourceAccount` indicates the resource account.
	Type pulumi.StringPtrInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
	DisplayName string `pulumi:"displayName"`
	// The ID of the parent folder.
	FolderId *string `pulumi:"folderId"`
	// Settlement account ID. If the value is empty, the current account will be used for settlement.
	PayerAccountId *string `pulumi:"payerAccountId"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
	DisplayName pulumi.StringInput
	// The ID of the parent folder.
	FolderId pulumi.StringPtrInput
	// Settlement account ID. If the value is empty, the current account will be used for settlement.
	PayerAccountId pulumi.StringPtrInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (Account) ElementType() reflect.Type {
	return reflect.TypeOf((*Account)(nil)).Elem()
}

func (i Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

type AccountOutput struct {
	*pulumi.OutputState
}

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountOutput)(nil)).Elem()
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
}
