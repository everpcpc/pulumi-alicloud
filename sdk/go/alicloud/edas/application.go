// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package edas

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates an EDAS ecs application on EDAS. The application will be deployed when `groupId` and `warUrl` are given.
//
// > **NOTE:** Available in 1.82.0+
//
// ## Import
//
// EDAS application can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:edas/application:Application app app_Id
// ```
type Application struct {
	pulumi.CustomResourceState

	// Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
	ApplicationName pulumi.StringOutput `pulumi:"applicationName"`
	// The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
	BuildPackId pulumi.IntPtrOutput `pulumi:"buildPackId"`
	// The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The description of the application that you want to create.
	Descriotion pulumi.StringPtrOutput `pulumi:"descriotion"`
	// The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
	EcuInfos pulumi.StringArrayOutput `pulumi:"ecuInfos"`
	// The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
	GroupId pulumi.StringPtrOutput `pulumi:"groupId"`
	// The URL for health checking of the application.
	HealthCheckUrl pulumi.StringPtrOutput `pulumi:"healthCheckUrl"`
	// The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
	LogicalRegionId pulumi.StringPtrOutput `pulumi:"logicalRegionId"`
	// The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
	PackageType pulumi.StringOutput `pulumi:"packageType"`
	// The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
	PackageVersion pulumi.StringPtrOutput `pulumi:"packageVersion"`
	// The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
	WarUrl pulumi.StringPtrOutput `pulumi:"warUrl"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationName'")
	}
	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.PackageType == nil {
		return nil, errors.New("invalid value for required argument 'PackageType'")
	}
	var resource Application
	err := ctx.RegisterResource("alicloud:edas/application:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("alicloud:edas/application:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
	ApplicationName *string `pulumi:"applicationName"`
	// The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
	BuildPackId *int `pulumi:"buildPackId"`
	// The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
	ClusterId *string `pulumi:"clusterId"`
	// The description of the application that you want to create.
	Descriotion *string `pulumi:"descriotion"`
	// The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
	EcuInfos []string `pulumi:"ecuInfos"`
	// The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
	GroupId *string `pulumi:"groupId"`
	// The URL for health checking of the application.
	HealthCheckUrl *string `pulumi:"healthCheckUrl"`
	// The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
	LogicalRegionId *string `pulumi:"logicalRegionId"`
	// The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
	PackageType *string `pulumi:"packageType"`
	// The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
	PackageVersion *string `pulumi:"packageVersion"`
	// The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
	WarUrl *string `pulumi:"warUrl"`
}

type ApplicationState struct {
	// Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
	ApplicationName pulumi.StringPtrInput
	// The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
	BuildPackId pulumi.IntPtrInput
	// The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
	ClusterId pulumi.StringPtrInput
	// The description of the application that you want to create.
	Descriotion pulumi.StringPtrInput
	// The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
	EcuInfos pulumi.StringArrayInput
	// The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
	GroupId pulumi.StringPtrInput
	// The URL for health checking of the application.
	HealthCheckUrl pulumi.StringPtrInput
	// The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
	LogicalRegionId pulumi.StringPtrInput
	// The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
	PackageType pulumi.StringPtrInput
	// The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
	PackageVersion pulumi.StringPtrInput
	// The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
	WarUrl pulumi.StringPtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
	ApplicationName string `pulumi:"applicationName"`
	// The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
	BuildPackId *int `pulumi:"buildPackId"`
	// The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
	ClusterId string `pulumi:"clusterId"`
	// The description of the application that you want to create.
	Descriotion *string `pulumi:"descriotion"`
	// The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
	EcuInfos []string `pulumi:"ecuInfos"`
	// The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
	GroupId *string `pulumi:"groupId"`
	// The URL for health checking of the application.
	HealthCheckUrl *string `pulumi:"healthCheckUrl"`
	// The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
	LogicalRegionId *string `pulumi:"logicalRegionId"`
	// The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
	PackageType string `pulumi:"packageType"`
	// The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
	PackageVersion *string `pulumi:"packageVersion"`
	// The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
	WarUrl *string `pulumi:"warUrl"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
	ApplicationName pulumi.StringInput
	// The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
	BuildPackId pulumi.IntPtrInput
	// The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
	ClusterId pulumi.StringInput
	// The description of the application that you want to create.
	Descriotion pulumi.StringPtrInput
	// The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
	EcuInfos pulumi.StringArrayInput
	// The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
	GroupId pulumi.StringPtrInput
	// The URL for health checking of the application.
	HealthCheckUrl pulumi.StringPtrInput
	// The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
	LogicalRegionId pulumi.StringPtrInput
	// The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
	PackageType pulumi.StringInput
	// The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
	PackageVersion pulumi.StringPtrInput
	// The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
	WarUrl pulumi.StringPtrInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (Application) ElementType() reflect.Type {
	return reflect.TypeOf((*Application)(nil)).Elem()
}

func (i Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

type ApplicationOutput struct {
	*pulumi.OutputState
}

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationOutput)(nil)).Elem()
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApplicationOutput{})
}
