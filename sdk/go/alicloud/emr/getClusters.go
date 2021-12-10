// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Emr Clusters of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.146.0+.
func GetClusters(ctx *pulumi.Context, args *GetClustersArgs, opts ...pulumi.InvokeOption) (*GetClustersResult, error) {
	var rv GetClustersResult
	err := ctx.Invoke("alicloud:emr/getClusters:getClusters", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusters.
type GetClustersArgs struct {
	// The name of the associated cluster.
	ClusterName *string `pulumi:"clusterName"`
	// The cluster type list.
	ClusterTypeLists []string `pulumi:"clusterTypeLists"`
	// How to create a cluster.
	CreateType *string `pulumi:"createType"`
	// The default status.
	DefaultStatus *bool `pulumi:"defaultStatus"`
	// The hosting type of the cluster.
	DepositType *string `pulumi:"depositType"`
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Cluster IDs.
	Ids []string `pulumi:"ids"`
	// The is desc.
	IsDesc *bool `pulumi:"isDesc"`
	// The host type of the cluster. The default is ECS.
	MachineType *string `pulumi:"machineType"`
	// A regex string to filter results by Cluster name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The Resource Group ID.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status list. Valid values: `ABNORMAL`, `CREATE_FAILED`, `CREATING`, `IDLE`, `RELEASED`, `RELEASE_FAILED`, `RELEASING`, `RUNNING`, `WAIT_FOR_PAY`.
	StatusLists []string `pulumi:"statusLists"`
	// The VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getClusters.
type GetClustersResult struct {
	ClusterName      *string              `pulumi:"clusterName"`
	ClusterTypeLists []string             `pulumi:"clusterTypeLists"`
	Clusters         []GetClustersCluster `pulumi:"clusters"`
	CreateType       *string              `pulumi:"createType"`
	DefaultStatus    *bool                `pulumi:"defaultStatus"`
	DepositType      *string              `pulumi:"depositType"`
	EnableDetails    *bool                `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id              string   `pulumi:"id"`
	Ids             []string `pulumi:"ids"`
	IsDesc          *bool    `pulumi:"isDesc"`
	MachineType     *string  `pulumi:"machineType"`
	NameRegex       *string  `pulumi:"nameRegex"`
	Names           []string `pulumi:"names"`
	OutputFile      *string  `pulumi:"outputFile"`
	ResourceGroupId *string  `pulumi:"resourceGroupId"`
	StatusLists     []string `pulumi:"statusLists"`
	VpcId           *string  `pulumi:"vpcId"`
}

func GetClustersOutput(ctx *pulumi.Context, args GetClustersOutputArgs, opts ...pulumi.InvokeOption) GetClustersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetClustersResult, error) {
			args := v.(GetClustersArgs)
			r, err := GetClusters(ctx, &args, opts...)
			return *r, err
		}).(GetClustersResultOutput)
}

// A collection of arguments for invoking getClusters.
type GetClustersOutputArgs struct {
	// The name of the associated cluster.
	ClusterName pulumi.StringPtrInput `pulumi:"clusterName"`
	// The cluster type list.
	ClusterTypeLists pulumi.StringArrayInput `pulumi:"clusterTypeLists"`
	// How to create a cluster.
	CreateType pulumi.StringPtrInput `pulumi:"createType"`
	// The default status.
	DefaultStatus pulumi.BoolPtrInput `pulumi:"defaultStatus"`
	// The hosting type of the cluster.
	DepositType pulumi.StringPtrInput `pulumi:"depositType"`
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Cluster IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The is desc.
	IsDesc pulumi.BoolPtrInput `pulumi:"isDesc"`
	// The host type of the cluster. The default is ECS.
	MachineType pulumi.StringPtrInput `pulumi:"machineType"`
	// A regex string to filter results by Cluster name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The Resource Group ID.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// The status list. Valid values: `ABNORMAL`, `CREATE_FAILED`, `CREATING`, `IDLE`, `RELEASED`, `RELEASE_FAILED`, `RELEASING`, `RUNNING`, `WAIT_FOR_PAY`.
	StatusLists pulumi.StringArrayInput `pulumi:"statusLists"`
	// The VPC ID.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (GetClustersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClustersArgs)(nil)).Elem()
}

// A collection of values returned by getClusters.
type GetClustersResultOutput struct{ *pulumi.OutputState }

func (GetClustersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClustersResult)(nil)).Elem()
}

func (o GetClustersResultOutput) ToGetClustersResultOutput() GetClustersResultOutput {
	return o
}

func (o GetClustersResultOutput) ToGetClustersResultOutputWithContext(ctx context.Context) GetClustersResultOutput {
	return o
}

func (o GetClustersResultOutput) ClusterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.ClusterName }).(pulumi.StringPtrOutput)
}

func (o GetClustersResultOutput) ClusterTypeLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClustersResult) []string { return v.ClusterTypeLists }).(pulumi.StringArrayOutput)
}

func (o GetClustersResultOutput) Clusters() GetClustersClusterArrayOutput {
	return o.ApplyT(func(v GetClustersResult) []GetClustersCluster { return v.Clusters }).(GetClustersClusterArrayOutput)
}

func (o GetClustersResultOutput) CreateType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.CreateType }).(pulumi.StringPtrOutput)
}

func (o GetClustersResultOutput) DefaultStatus() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *bool { return v.DefaultStatus }).(pulumi.BoolPtrOutput)
}

func (o GetClustersResultOutput) DepositType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.DepositType }).(pulumi.StringPtrOutput)
}

func (o GetClustersResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetClustersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetClustersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClustersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetClustersResultOutput) IsDesc() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *bool { return v.IsDesc }).(pulumi.BoolPtrOutput)
}

func (o GetClustersResultOutput) MachineType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.MachineType }).(pulumi.StringPtrOutput)
}

func (o GetClustersResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetClustersResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClustersResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetClustersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetClustersResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o GetClustersResultOutput) StatusLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClustersResult) []string { return v.StatusLists }).(pulumi.StringArrayOutput)
}

func (o GetClustersResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClustersResultOutput{})
}
