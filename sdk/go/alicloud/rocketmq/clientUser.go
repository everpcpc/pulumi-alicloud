// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rocketmq

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Sag ClientUser resource. This topic describes how to manage accounts as an administrator. After you configure the network, you can create multiple accounts and distribute them to end users so that clients can access Alibaba Cloud.
// 
// For information about Sag ClientUser and how to use it, see [What is Sag ClientUser](https://www.alibabacloud.com/help/doc-detail/108326.htm).
// 
// > **NOTE:** Available in 1.65.0+
// 
// > **NOTE:** Only the following regions support. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/sag_client_user.html.markdown.
type ClientUser struct {
	pulumi.CustomResourceState

	// The SAG APP bandwidth that the user can use. Unit: Kbit/s. Maximum value: 2000 Kbit/s.
	Bandwidth pulumi.IntOutput `pulumi:"bandwidth"`
	// The IP address of the SAG APP. If you specify this parameter, the current account always uses the specified IP address.Note The IP address must be in the private CIDR block of the SAG client.If you do not specify this parameter, the system automatically allocates an IP address from the private CIDR block of the SAG client. In this case, each re-connection uses a different IP address.
	ClientIp pulumi.StringPtrOutput `pulumi:"clientIp"`
	KmsEncryptedPassword pulumi.StringPtrOutput `pulumi:"kmsEncryptedPassword"`
	KmsEncryptionContext pulumi.MapOutput `pulumi:"kmsEncryptionContext"`
	// The password used to log on to the SAG APP.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.
	Password pulumi.StringOutput `pulumi:"password"`
	// The ID of the SAG instance created for the SAG APP.
	SagId pulumi.StringOutput `pulumi:"sagId"`
	// The email address of the user. The administrator uses this address to send the account information for logging on to the APP to the user.
	UserMail pulumi.StringOutput `pulumi:"userMail"`
	// The user name. User names in the same SAG APP must be unique.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewClientUser registers a new resource with the given unique name, arguments, and options.
func NewClientUser(ctx *pulumi.Context,
	name string, args *ClientUserArgs, opts ...pulumi.ResourceOption) (*ClientUser, error) {
	if args == nil || args.Bandwidth == nil {
		return nil, errors.New("missing required argument 'Bandwidth'")
	}
	if args == nil || args.SagId == nil {
		return nil, errors.New("missing required argument 'SagId'")
	}
	if args == nil || args.UserMail == nil {
		return nil, errors.New("missing required argument 'UserMail'")
	}
	if args == nil {
		args = &ClientUserArgs{}
	}
	var resource ClientUser
	err := ctx.RegisterResource("alicloud:rocketmq/clientUser:ClientUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientUser gets an existing ClientUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientUserState, opts ...pulumi.ResourceOption) (*ClientUser, error) {
	var resource ClientUser
	err := ctx.ReadResource("alicloud:rocketmq/clientUser:ClientUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientUser resources.
type clientUserState struct {
	// The SAG APP bandwidth that the user can use. Unit: Kbit/s. Maximum value: 2000 Kbit/s.
	Bandwidth *int `pulumi:"bandwidth"`
	// The IP address of the SAG APP. If you specify this parameter, the current account always uses the specified IP address.Note The IP address must be in the private CIDR block of the SAG client.If you do not specify this parameter, the system automatically allocates an IP address from the private CIDR block of the SAG client. In this case, each re-connection uses a different IP address.
	ClientIp *string `pulumi:"clientIp"`
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// The password used to log on to the SAG APP.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.
	Password *string `pulumi:"password"`
	// The ID of the SAG instance created for the SAG APP.
	SagId *string `pulumi:"sagId"`
	// The email address of the user. The administrator uses this address to send the account information for logging on to the APP to the user.
	UserMail *string `pulumi:"userMail"`
	// The user name. User names in the same SAG APP must be unique.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.
	UserName *string `pulumi:"userName"`
}

type ClientUserState struct {
	// The SAG APP bandwidth that the user can use. Unit: Kbit/s. Maximum value: 2000 Kbit/s.
	Bandwidth pulumi.IntPtrInput
	// The IP address of the SAG APP. If you specify this parameter, the current account always uses the specified IP address.Note The IP address must be in the private CIDR block of the SAG client.If you do not specify this parameter, the system automatically allocates an IP address from the private CIDR block of the SAG client. In this case, each re-connection uses a different IP address.
	ClientIp pulumi.StringPtrInput
	KmsEncryptedPassword pulumi.StringPtrInput
	KmsEncryptionContext pulumi.MapInput
	// The password used to log on to the SAG APP.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.
	Password pulumi.StringPtrInput
	// The ID of the SAG instance created for the SAG APP.
	SagId pulumi.StringPtrInput
	// The email address of the user. The administrator uses this address to send the account information for logging on to the APP to the user.
	UserMail pulumi.StringPtrInput
	// The user name. User names in the same SAG APP must be unique.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.
	UserName pulumi.StringPtrInput
}

func (ClientUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientUserState)(nil)).Elem()
}

type clientUserArgs struct {
	// The SAG APP bandwidth that the user can use. Unit: Kbit/s. Maximum value: 2000 Kbit/s.
	Bandwidth int `pulumi:"bandwidth"`
	// The IP address of the SAG APP. If you specify this parameter, the current account always uses the specified IP address.Note The IP address must be in the private CIDR block of the SAG client.If you do not specify this parameter, the system automatically allocates an IP address from the private CIDR block of the SAG client. In this case, each re-connection uses a different IP address.
	ClientIp *string `pulumi:"clientIp"`
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// The password used to log on to the SAG APP.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.
	Password *string `pulumi:"password"`
	// The ID of the SAG instance created for the SAG APP.
	SagId string `pulumi:"sagId"`
	// The email address of the user. The administrator uses this address to send the account information for logging on to the APP to the user.
	UserMail string `pulumi:"userMail"`
	// The user name. User names in the same SAG APP must be unique.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.
	UserName *string `pulumi:"userName"`
}

// The set of arguments for constructing a ClientUser resource.
type ClientUserArgs struct {
	// The SAG APP bandwidth that the user can use. Unit: Kbit/s. Maximum value: 2000 Kbit/s.
	Bandwidth pulumi.IntInput
	// The IP address of the SAG APP. If you specify this parameter, the current account always uses the specified IP address.Note The IP address must be in the private CIDR block of the SAG client.If you do not specify this parameter, the system automatically allocates an IP address from the private CIDR block of the SAG client. In this case, each re-connection uses a different IP address.
	ClientIp pulumi.StringPtrInput
	KmsEncryptedPassword pulumi.StringPtrInput
	KmsEncryptionContext pulumi.MapInput
	// The password used to log on to the SAG APP.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.
	Password pulumi.StringPtrInput
	// The ID of the SAG instance created for the SAG APP.
	SagId pulumi.StringInput
	// The email address of the user. The administrator uses this address to send the account information for logging on to the APP to the user.
	UserMail pulumi.StringInput
	// The user name. User names in the same SAG APP must be unique.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.
	UserName pulumi.StringPtrInput
}

func (ClientUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientUserArgs)(nil)).Elem()
}

