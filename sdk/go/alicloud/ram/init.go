// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "alicloud:ram/accessKey:AccessKey":
		r = &AccessKey{}
	case "alicloud:ram/accountAlias:AccountAlias":
		r = &AccountAlias{}
	case "alicloud:ram/accountPasswordPolicy:AccountPasswordPolicy":
		r = &AccountPasswordPolicy{}
	case "alicloud:ram/alias:Alias":
		r = &Alias{}
	case "alicloud:ram/group:Group":
		r = &Group{}
	case "alicloud:ram/groupMembership:GroupMembership":
		r = &GroupMembership{}
	case "alicloud:ram/groupPolicyAttachment:GroupPolicyAttachment":
		r = &GroupPolicyAttachment{}
	case "alicloud:ram/loginProfile:LoginProfile":
		r = &LoginProfile{}
	case "alicloud:ram/policy:Policy":
		r = &Policy{}
	case "alicloud:ram/role:Role":
		r = &Role{}
	case "alicloud:ram/roleAttachment:RoleAttachment":
		r = &RoleAttachment{}
	case "alicloud:ram/rolePolicyAttachment:RolePolicyAttachment":
		r = &RolePolicyAttachment{}
	case "alicloud:ram/samlProvider:SamlProvider":
		r = &SamlProvider{}
	case "alicloud:ram/user:User":
		r = &User{}
	case "alicloud:ram/userPolicyAttachment:UserPolicyAttachment":
		r = &UserPolicyAttachment{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := alicloud.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"alicloud",
		"ram/accessKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ram/accountAlias",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ram/accountPasswordPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ram/alias",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ram/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ram/groupMembership",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ram/groupPolicyAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ram/loginProfile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ram/policy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ram/role",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ram/roleAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ram/rolePolicyAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ram/samlProvider",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ram/user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ram/userPolicyAttachment",
		&module{version},
	)
}
