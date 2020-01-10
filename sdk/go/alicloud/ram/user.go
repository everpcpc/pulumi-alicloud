// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/ram_user.html.markdown.
type User struct {
	s *pulumi.ResourceState
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOpt) (*User, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["comments"] = nil
		inputs["displayName"] = nil
		inputs["email"] = nil
		inputs["force"] = nil
		inputs["mobile"] = nil
		inputs["name"] = nil
	} else {
		inputs["comments"] = args.Comments
		inputs["displayName"] = args.DisplayName
		inputs["email"] = args.Email
		inputs["force"] = args.Force
		inputs["mobile"] = args.Mobile
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("alicloud:ram/user:User", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &User{s: s}, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.ID, state *UserState, opts ...pulumi.ResourceOpt) (*User, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["comments"] = state.Comments
		inputs["displayName"] = state.DisplayName
		inputs["email"] = state.Email
		inputs["force"] = state.Force
		inputs["mobile"] = state.Mobile
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("alicloud:ram/user:User", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &User{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *User) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *User) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Comment of the RAM user. This parameter can have a string of 1 to 128 characters.
func (r *User) Comments() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["comments"])
}

// Name of the RAM user which for display. This name can have a string of 1 to 128 characters or Chinese characters, must contain only alphanumeric characters or Chinese characters or hyphens, such as "-",".", and must not end with a hyphen.
func (r *User) DisplayName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["displayName"])
}

// Email of the RAM user.
func (r *User) Email() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["email"])
}

// This parameter is used for resource destroy. Default value is `false`.
func (r *User) Force() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["force"])
}

// Phone number of the RAM user. This number must contain an international area code prefix, just look like this: 86-18600008888.
func (r *User) Mobile() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["mobile"])
}

// Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
func (r *User) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering User resources.
type UserState struct {
	// Comment of the RAM user. This parameter can have a string of 1 to 128 characters.
	Comments interface{}
	// Name of the RAM user which for display. This name can have a string of 1 to 128 characters or Chinese characters, must contain only alphanumeric characters or Chinese characters or hyphens, such as "-",".", and must not end with a hyphen.
	DisplayName interface{}
	// Email of the RAM user.
	Email interface{}
	// This parameter is used for resource destroy. Default value is `false`.
	Force interface{}
	// Phone number of the RAM user. This number must contain an international area code prefix, just look like this: 86-18600008888.
	Mobile interface{}
	// Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
	Name interface{}
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// Comment of the RAM user. This parameter can have a string of 1 to 128 characters.
	Comments interface{}
	// Name of the RAM user which for display. This name can have a string of 1 to 128 characters or Chinese characters, must contain only alphanumeric characters or Chinese characters or hyphens, such as "-",".", and must not end with a hyphen.
	DisplayName interface{}
	// Email of the RAM user.
	Email interface{}
	// This parameter is used for resource destroy. Default value is `false`.
	Force interface{}
	// Phone number of the RAM user. This number must contain an international area code prefix, just look like this: 86-18600008888.
	Mobile interface{}
	// Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
	Name interface{}
}
