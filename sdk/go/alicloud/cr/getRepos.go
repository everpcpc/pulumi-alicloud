// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cr

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides a list Container Registry repositories on Alibaba Cloud.
// 
// > **NOTE:** Available in v1.35.0+
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/cr_repos.html.markdown.
func GetRepos(ctx *pulumi.Context, args *GetReposArgs, opts ...pulumi.InvokeOption) (*GetReposResult, error) {
	var rv GetReposResult
	err := ctx.Invoke("alicloud:cr/getRepos:getRepos", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRepos.
type GetReposArgs struct {
	// Boolean, false by default, only repository attributes are exported. Set to true if domain list and tags belong to this repository are needed. See `tags` in attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A regex string to filter results by repository name.
	NameRegex *string `pulumi:"nameRegex"`
	// Name of container registry namespace where the repositories are located in.
	Namespace *string `pulumi:"namespace"`
	OutputFile *string `pulumi:"outputFile"`
}


// A collection of values returned by getRepos.
type GetReposResult struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of matched Container Registry Repositories. Its element is set to `names`.
	Ids []string `pulumi:"ids"`
	NameRegex *string `pulumi:"nameRegex"`
	// A list of repository names.
	Names []string `pulumi:"names"`
	// Name of container registry namespace where repo is located.
	Namespace *string `pulumi:"namespace"`
	OutputFile *string `pulumi:"outputFile"`
	// A list of matched Container Registry Namespaces. Each element contains the following attributes:
	Repos []GetReposRepo `pulumi:"repos"`
}

