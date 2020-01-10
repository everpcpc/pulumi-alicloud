// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fc

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides the Function Compute services of the current Alibaba Cloud user.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/fc_services.html.markdown.
func LookupServices(ctx *pulumi.Context, args *GetServicesArgs) (*GetServicesResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["ids"] = args.Ids
		inputs["nameRegex"] = args.NameRegex
		inputs["outputFile"] = args.OutputFile
	}
	outputs, err := ctx.Invoke("alicloud:fc/getServices:getServices", inputs)
	if err != nil {
		return nil, err
	}
	return &GetServicesResult{
		Ids: outputs["ids"],
		NameRegex: outputs["nameRegex"],
		Names: outputs["names"],
		OutputFile: outputs["outputFile"],
		Services: outputs["services"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getServices.
type GetServicesArgs struct {
	Ids interface{}
	// A regex string to filter results by FC service name.
	// * `ids` (Optional, Available in 1.53.0+) - A list of FC services ids.
	NameRegex interface{}
	OutputFile interface{}
}

// A collection of values returned by getServices.
type GetServicesResult struct {
	// A list of FC services ids.
	Ids interface{}
	NameRegex interface{}
	// A list of FC services names.
	Names interface{}
	OutputFile interface{}
	// A list of FC services. Each element contains the following attributes:
	Services interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
