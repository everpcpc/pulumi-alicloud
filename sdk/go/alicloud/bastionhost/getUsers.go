// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bastionhost

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Bastionhost Users of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.133.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/bastionhost"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ids, err := bastionhost.GetUsers(ctx, &bastionhost.GetUsersArgs{
// 			InstanceId: "example_value",
// 			Ids: []string{
// 				"1",
// 				"10",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("bastionhostUserId1", ids.Users[0].Id)
// 		opt0 := "^my-User"
// 		nameRegex, err := bastionhost.GetUsers(ctx, &bastionhost.GetUsersArgs{
// 			InstanceId: "example_value",
// 			NameRegex:  &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("bastionhostUserId2", nameRegex.Users[0].Id)
// 		return nil
// 	})
// }
// ```
func GetUsers(ctx *pulumi.Context, args *GetUsersArgs, opts ...pulumi.InvokeOption) (*GetUsersResult, error) {
	var rv GetUsersResult
	err := ctx.Invoke("alicloud:bastionhost/getUsers:getUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUsers.
type GetUsersArgs struct {
	// Specify the New Created the User's Display Name. Supports up to 128 Characters.
	DisplayName *string `pulumi:"displayName"`
	// A list of User IDs.
	Ids []string `pulumi:"ids"`
	// You Want to Query the User the Bastion Host ID of.
	InstanceId string `pulumi:"instanceId"`
	// Specify the New of the User That Created a Different Mobile Phone Number from Your.
	Mobile *string `pulumi:"mobile"`
	// A regex string to filter results by User name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// Specify the New of the User That Created the Source. Valid Values: Local: Local User RAM: Ram User.
	Source *string `pulumi:"source"`
	// Specify the Newly Created User Is Uniquely Identified. Indicates That the Parameter Is a Bastion Host Corresponding to the User with the Ram User's Unique Identifier. The Newly Created User Source Grant Permission to a RAM User (That Is, Source Used as a Ram), this Parameter Is Required. You Can Call Access Control of Listusers Interface from the Return Data Userid to Obtain the Parameters.
	SourceUserId *string `pulumi:"sourceUserId"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// Specify the New User Name. This Parameter Is Only by Letters, Lowercase Letters, Numbers, and Underscores (_), Supports up to 128 Characters.
	UserName *string `pulumi:"userName"`
}

// A collection of values returned by getUsers.
type GetUsersResult struct {
	DisplayName *string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id           string         `pulumi:"id"`
	Ids          []string       `pulumi:"ids"`
	InstanceId   string         `pulumi:"instanceId"`
	Mobile       *string        `pulumi:"mobile"`
	NameRegex    *string        `pulumi:"nameRegex"`
	Names        []string       `pulumi:"names"`
	OutputFile   *string        `pulumi:"outputFile"`
	Source       *string        `pulumi:"source"`
	SourceUserId *string        `pulumi:"sourceUserId"`
	Status       *string        `pulumi:"status"`
	UserName     *string        `pulumi:"userName"`
	Users        []GetUsersUser `pulumi:"users"`
}