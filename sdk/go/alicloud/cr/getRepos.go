// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a list Container Registry repositories on Alibaba Cloud.
//
// > **NOTE:** Available in v1.35.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		myRepos, err := cr.GetRepos(ctx, &cr.GetReposArgs{
// 			NameRegex:  pulumi.StringRef("my-repos"),
// 			OutputFile: pulumi.StringRef("my-repo-json"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("output", myRepos.Repos)
// 		return nil
// 	})
// }
// ```
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
	Namespace  *string `pulumi:"namespace"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getRepos.
type GetReposResult struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of matched Container Registry Repositories. Its element is set to `names`.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of repository names.
	Names []string `pulumi:"names"`
	// Name of container registry namespace where repo is located.
	Namespace  *string `pulumi:"namespace"`
	OutputFile *string `pulumi:"outputFile"`
	// A list of matched Container Registry Namespaces. Each element contains the following attributes:
	Repos []GetReposRepo `pulumi:"repos"`
}

func GetReposOutput(ctx *pulumi.Context, args GetReposOutputArgs, opts ...pulumi.InvokeOption) GetReposResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetReposResult, error) {
			args := v.(GetReposArgs)
			r, err := GetRepos(ctx, &args, opts...)
			return *r, err
		}).(GetReposResultOutput)
}

// A collection of arguments for invoking getRepos.
type GetReposOutputArgs struct {
	// Boolean, false by default, only repository attributes are exported. Set to true if domain list and tags belong to this repository are needed. See `tags` in attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A regex string to filter results by repository name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// Name of container registry namespace where the repositories are located in.
	Namespace  pulumi.StringPtrInput `pulumi:"namespace"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetReposOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReposArgs)(nil)).Elem()
}

// A collection of values returned by getRepos.
type GetReposResultOutput struct{ *pulumi.OutputState }

func (GetReposResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReposResult)(nil)).Elem()
}

func (o GetReposResultOutput) ToGetReposResultOutput() GetReposResultOutput {
	return o
}

func (o GetReposResultOutput) ToGetReposResultOutputWithContext(ctx context.Context) GetReposResultOutput {
	return o
}

func (o GetReposResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetReposResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetReposResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetReposResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of matched Container Registry Repositories. Its element is set to `names`.
func (o GetReposResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetReposResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetReposResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetReposResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of repository names.
func (o GetReposResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetReposResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

// Name of container registry namespace where repo is located.
func (o GetReposResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetReposResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o GetReposResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetReposResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// A list of matched Container Registry Namespaces. Each element contains the following attributes:
func (o GetReposResultOutput) Repos() GetReposRepoArrayOutput {
	return o.ApplyT(func(v GetReposResult) []GetReposRepo { return v.Repos }).(GetReposRepoArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetReposResultOutput{})
}
