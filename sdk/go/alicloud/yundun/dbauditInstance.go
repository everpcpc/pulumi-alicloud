// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package yundun

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Cloud DBaudit instance resource ("Yundun_dbaudit" is the short term of this product).
// 
// > **NOTE:** The endpoint of bssopenapi used only support "business.aliyuncs.com" at present.
// 
// > **NOTE:** Available in 1.62.0+ .
// 
// > **NOTE:** In order to destroy Cloud DBaudit instance , users are required to apply for white list first
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/yundun_dbaudit_instance.html.markdown.
type DBAuditInstance struct {
	pulumi.CustomResourceState

	// Description of the instance. This name can have a string of 1 to 63 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// Duration for initially producing the instance. Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Plan code of the Cloud DBAudit to produce. (alpha.professional, alpha.basic, alpha.premium) 
	PlanCode pulumi.StringOutput `pulumi:"planCode"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// vSwtich ID configured to audit
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
}

// NewDBAuditInstance registers a new resource with the given unique name, arguments, and options.
func NewDBAuditInstance(ctx *pulumi.Context,
	name string, args *DBAuditInstanceArgs, opts ...pulumi.ResourceOption) (*DBAuditInstance, error) {
	if args == nil || args.Description == nil {
		return nil, errors.New("missing required argument 'Description'")
	}
	if args == nil || args.PlanCode == nil {
		return nil, errors.New("missing required argument 'PlanCode'")
	}
	if args == nil || args.VswitchId == nil {
		return nil, errors.New("missing required argument 'VswitchId'")
	}
	if args == nil {
		args = &DBAuditInstanceArgs{}
	}
	var resource DBAuditInstance
	err := ctx.RegisterResource("alicloud:yundun/dBAuditInstance:DBAuditInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDBAuditInstance gets an existing DBAuditInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDBAuditInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DBAuditInstanceState, opts ...pulumi.ResourceOption) (*DBAuditInstance, error) {
	var resource DBAuditInstance
	err := ctx.ReadResource("alicloud:yundun/dBAuditInstance:DBAuditInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DBAuditInstance resources.
type dbauditInstanceState struct {
	// Description of the instance. This name can have a string of 1 to 63 characters.
	Description *string `pulumi:"description"`
	// Duration for initially producing the instance. Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
	Period *int `pulumi:"period"`
	// Plan code of the Cloud DBAudit to produce. (alpha.professional, alpha.basic, alpha.premium) 
	PlanCode *string `pulumi:"planCode"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// vSwtich ID configured to audit
	VswitchId *string `pulumi:"vswitchId"`
}

type DBAuditInstanceState struct {
	// Description of the instance. This name can have a string of 1 to 63 characters.
	Description pulumi.StringPtrInput
	// Duration for initially producing the instance. Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
	Period pulumi.IntPtrInput
	// Plan code of the Cloud DBAudit to produce. (alpha.professional, alpha.basic, alpha.premium) 
	PlanCode pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// vSwtich ID configured to audit
	VswitchId pulumi.StringPtrInput
}

func (DBAuditInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbauditInstanceState)(nil)).Elem()
}

type dbauditInstanceArgs struct {
	// Description of the instance. This name can have a string of 1 to 63 characters.
	Description string `pulumi:"description"`
	// Duration for initially producing the instance. Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
	Period *int `pulumi:"period"`
	// Plan code of the Cloud DBAudit to produce. (alpha.professional, alpha.basic, alpha.premium) 
	PlanCode string `pulumi:"planCode"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// vSwtich ID configured to audit
	VswitchId string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a DBAuditInstance resource.
type DBAuditInstanceArgs struct {
	// Description of the instance. This name can have a string of 1 to 63 characters.
	Description pulumi.StringInput
	// Duration for initially producing the instance. Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
	Period pulumi.IntPtrInput
	// Plan code of the Cloud DBAudit to produce. (alpha.professional, alpha.basic, alpha.premium) 
	PlanCode pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// vSwtich ID configured to audit
	VswitchId pulumi.StringInput
}

func (DBAuditInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbauditInstanceArgs)(nil)).Elem()
}

