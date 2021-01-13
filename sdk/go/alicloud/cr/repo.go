// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource will help you to manager Container Registry repositories.
//
// > **NOTE:** Available in v1.35.0+.
//
// > **NOTE:** You need to set your registry password in Container Registry console before use this resource.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/cr"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cr.NewNamespace(ctx, "my_namespace", &cr.NamespaceArgs{
// 			AutoCreate:        pulumi.Bool(false),
// 			DefaultVisibility: pulumi.String("PUBLIC"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cr.NewRepo(ctx, "my_repo", &cr.RepoArgs{
// 			Namespace: my_namespace.Name,
// 			Summary:   pulumi.String("this is summary of my new repo"),
// 			RepoType:  pulumi.String("PUBLIC"),
// 			Detail:    pulumi.String("this is a public repo"),
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
// Container Registry repository can be imported using the `namespace/repository`, e.g.
//
// ```sh
//  $ pulumi import alicloud:cr/repo:Repo default `my-namespace/my-repo`
// ```
type Repo struct {
	pulumi.CustomResourceState

	// The repository specific information. MarkDown format is supported, and the length limit is 2000.
	Detail pulumi.StringPtrOutput `pulumi:"detail"`
	// The repository domain list.
	DomainList RepoDomainListOutput `pulumi:"domainList"`
	// Name of container registry repository.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of container registry namespace where repository is located.
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// `PUBLIC` or `PRIVATE`, repo's visibility.
	RepoType pulumi.StringOutput `pulumi:"repoType"`
	// The repository general information. It can contain 1 to 80 characters.
	Summary pulumi.StringOutput `pulumi:"summary"`
}

// NewRepo registers a new resource with the given unique name, arguments, and options.
func NewRepo(ctx *pulumi.Context,
	name string, args *RepoArgs, opts ...pulumi.ResourceOption) (*Repo, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Namespace == nil {
		return nil, errors.New("invalid value for required argument 'Namespace'")
	}
	if args.RepoType == nil {
		return nil, errors.New("invalid value for required argument 'RepoType'")
	}
	if args.Summary == nil {
		return nil, errors.New("invalid value for required argument 'Summary'")
	}
	var resource Repo
	err := ctx.RegisterResource("alicloud:cr/repo:Repo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepo gets an existing Repo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepoState, opts ...pulumi.ResourceOption) (*Repo, error) {
	var resource Repo
	err := ctx.ReadResource("alicloud:cr/repo:Repo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Repo resources.
type repoState struct {
	// The repository specific information. MarkDown format is supported, and the length limit is 2000.
	Detail *string `pulumi:"detail"`
	// The repository domain list.
	DomainList *RepoDomainList `pulumi:"domainList"`
	// Name of container registry repository.
	Name *string `pulumi:"name"`
	// Name of container registry namespace where repository is located.
	Namespace *string `pulumi:"namespace"`
	// `PUBLIC` or `PRIVATE`, repo's visibility.
	RepoType *string `pulumi:"repoType"`
	// The repository general information. It can contain 1 to 80 characters.
	Summary *string `pulumi:"summary"`
}

type RepoState struct {
	// The repository specific information. MarkDown format is supported, and the length limit is 2000.
	Detail pulumi.StringPtrInput
	// The repository domain list.
	DomainList RepoDomainListPtrInput
	// Name of container registry repository.
	Name pulumi.StringPtrInput
	// Name of container registry namespace where repository is located.
	Namespace pulumi.StringPtrInput
	// `PUBLIC` or `PRIVATE`, repo's visibility.
	RepoType pulumi.StringPtrInput
	// The repository general information. It can contain 1 to 80 characters.
	Summary pulumi.StringPtrInput
}

func (RepoState) ElementType() reflect.Type {
	return reflect.TypeOf((*repoState)(nil)).Elem()
}

type repoArgs struct {
	// The repository specific information. MarkDown format is supported, and the length limit is 2000.
	Detail *string `pulumi:"detail"`
	// Name of container registry repository.
	Name *string `pulumi:"name"`
	// Name of container registry namespace where repository is located.
	Namespace string `pulumi:"namespace"`
	// `PUBLIC` or `PRIVATE`, repo's visibility.
	RepoType string `pulumi:"repoType"`
	// The repository general information. It can contain 1 to 80 characters.
	Summary string `pulumi:"summary"`
}

// The set of arguments for constructing a Repo resource.
type RepoArgs struct {
	// The repository specific information. MarkDown format is supported, and the length limit is 2000.
	Detail pulumi.StringPtrInput
	// Name of container registry repository.
	Name pulumi.StringPtrInput
	// Name of container registry namespace where repository is located.
	Namespace pulumi.StringInput
	// `PUBLIC` or `PRIVATE`, repo's visibility.
	RepoType pulumi.StringInput
	// The repository general information. It can contain 1 to 80 characters.
	Summary pulumi.StringInput
}

func (RepoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repoArgs)(nil)).Elem()
}

type RepoInput interface {
	pulumi.Input

	ToRepoOutput() RepoOutput
	ToRepoOutputWithContext(ctx context.Context) RepoOutput
}

func (Repo) ElementType() reflect.Type {
	return reflect.TypeOf((*Repo)(nil)).Elem()
}

func (i Repo) ToRepoOutput() RepoOutput {
	return i.ToRepoOutputWithContext(context.Background())
}

func (i Repo) ToRepoOutputWithContext(ctx context.Context) RepoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepoOutput)
}

type RepoOutput struct {
	*pulumi.OutputState
}

func (RepoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepoOutput)(nil)).Elem()
}

func (o RepoOutput) ToRepoOutput() RepoOutput {
	return o
}

func (o RepoOutput) ToRepoOutputWithContext(ctx context.Context) RepoOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RepoOutput{})
}
