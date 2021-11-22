// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides a site monitor resource and it can be used to monitor public endpoints and websites.
// Details at https://www.alibabacloud.com/help/doc-detail/67907.htm
//
// Available in 1.72.0+
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cms"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cms.NewSiteMonitor(ctx, "basic", &cms.SiteMonitorArgs{
// 			Address:  pulumi.String("http://www.alibabacloud.com"),
// 			Interval: pulumi.Int(5),
// 			IspCities: cms.SiteMonitorIspCityArray{
// 				&cms.SiteMonitorIspCityArgs{
// 					City: pulumi.String("546"),
// 					Isp:  pulumi.String("465"),
// 				},
// 			},
// 			TaskName: pulumi.String("tf-testAccCmsSiteMonitor_basic"),
// 			TaskType: pulumi.String("HTTP"),
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
// Alarm rule can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:cms/siteMonitor:SiteMonitor alarm abc12345
// ```
type SiteMonitor struct {
	pulumi.CustomResourceState

	// The URL or IP address monitored by the site monitoring task.
	Address pulumi.StringOutput `pulumi:"address"`
	// The IDs of existing alert rules to be associated with the site monitoring task.
	AlertIds   pulumi.StringArrayOutput `pulumi:"alertIds"`
	CreateTime pulumi.StringOutput      `pulumi:"createTime"`
	// The monitoring interval of the site monitoring task. Unit: minutes. Valid values: 1, 5, and 15. Default value: 1.
	Interval pulumi.IntPtrOutput `pulumi:"interval"`
	// The detection points in a JSON array. For example, `[{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring.
	IspCities SiteMonitorIspCityArrayOutput `pulumi:"ispCities"`
	// The extended options of the protocol of the site monitoring task. The options vary according to the protocol.
	OptionsJson pulumi.StringPtrOutput `pulumi:"optionsJson"`
	// The name of the site monitoring task. The name must be 4 to 100 characters in length. The name can contain the following types of characters: letters, digits, and underscores.
	TaskName  pulumi.StringOutput `pulumi:"taskName"`
	TaskState pulumi.StringOutput `pulumi:"taskState"`
	// The protocol of the site monitoring task. Currently, site monitoring supports the following protocols: HTTP, Ping, TCP, UDP, DNS, SMTP, POP3, and FTP.
	TaskType   pulumi.StringOutput `pulumi:"taskType"`
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewSiteMonitor registers a new resource with the given unique name, arguments, and options.
func NewSiteMonitor(ctx *pulumi.Context,
	name string, args *SiteMonitorArgs, opts ...pulumi.ResourceOption) (*SiteMonitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.TaskName == nil {
		return nil, errors.New("invalid value for required argument 'TaskName'")
	}
	if args.TaskType == nil {
		return nil, errors.New("invalid value for required argument 'TaskType'")
	}
	var resource SiteMonitor
	err := ctx.RegisterResource("alicloud:cms/siteMonitor:SiteMonitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSiteMonitor gets an existing SiteMonitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSiteMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteMonitorState, opts ...pulumi.ResourceOption) (*SiteMonitor, error) {
	var resource SiteMonitor
	err := ctx.ReadResource("alicloud:cms/siteMonitor:SiteMonitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SiteMonitor resources.
type siteMonitorState struct {
	// The URL or IP address monitored by the site monitoring task.
	Address *string `pulumi:"address"`
	// The IDs of existing alert rules to be associated with the site monitoring task.
	AlertIds   []string `pulumi:"alertIds"`
	CreateTime *string  `pulumi:"createTime"`
	// The monitoring interval of the site monitoring task. Unit: minutes. Valid values: 1, 5, and 15. Default value: 1.
	Interval *int `pulumi:"interval"`
	// The detection points in a JSON array. For example, `[{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring.
	IspCities []SiteMonitorIspCity `pulumi:"ispCities"`
	// The extended options of the protocol of the site monitoring task. The options vary according to the protocol.
	OptionsJson *string `pulumi:"optionsJson"`
	// The name of the site monitoring task. The name must be 4 to 100 characters in length. The name can contain the following types of characters: letters, digits, and underscores.
	TaskName  *string `pulumi:"taskName"`
	TaskState *string `pulumi:"taskState"`
	// The protocol of the site monitoring task. Currently, site monitoring supports the following protocols: HTTP, Ping, TCP, UDP, DNS, SMTP, POP3, and FTP.
	TaskType   *string `pulumi:"taskType"`
	UpdateTime *string `pulumi:"updateTime"`
}

type SiteMonitorState struct {
	// The URL or IP address monitored by the site monitoring task.
	Address pulumi.StringPtrInput
	// The IDs of existing alert rules to be associated with the site monitoring task.
	AlertIds   pulumi.StringArrayInput
	CreateTime pulumi.StringPtrInput
	// The monitoring interval of the site monitoring task. Unit: minutes. Valid values: 1, 5, and 15. Default value: 1.
	Interval pulumi.IntPtrInput
	// The detection points in a JSON array. For example, `[{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring.
	IspCities SiteMonitorIspCityArrayInput
	// The extended options of the protocol of the site monitoring task. The options vary according to the protocol.
	OptionsJson pulumi.StringPtrInput
	// The name of the site monitoring task. The name must be 4 to 100 characters in length. The name can contain the following types of characters: letters, digits, and underscores.
	TaskName  pulumi.StringPtrInput
	TaskState pulumi.StringPtrInput
	// The protocol of the site monitoring task. Currently, site monitoring supports the following protocols: HTTP, Ping, TCP, UDP, DNS, SMTP, POP3, and FTP.
	TaskType   pulumi.StringPtrInput
	UpdateTime pulumi.StringPtrInput
}

func (SiteMonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteMonitorState)(nil)).Elem()
}

type siteMonitorArgs struct {
	// The URL or IP address monitored by the site monitoring task.
	Address string `pulumi:"address"`
	// The IDs of existing alert rules to be associated with the site monitoring task.
	AlertIds []string `pulumi:"alertIds"`
	// The monitoring interval of the site monitoring task. Unit: minutes. Valid values: 1, 5, and 15. Default value: 1.
	Interval *int `pulumi:"interval"`
	// The detection points in a JSON array. For example, `[{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring.
	IspCities []SiteMonitorIspCity `pulumi:"ispCities"`
	// The extended options of the protocol of the site monitoring task. The options vary according to the protocol.
	OptionsJson *string `pulumi:"optionsJson"`
	// The name of the site monitoring task. The name must be 4 to 100 characters in length. The name can contain the following types of characters: letters, digits, and underscores.
	TaskName string `pulumi:"taskName"`
	// The protocol of the site monitoring task. Currently, site monitoring supports the following protocols: HTTP, Ping, TCP, UDP, DNS, SMTP, POP3, and FTP.
	TaskType string `pulumi:"taskType"`
}

// The set of arguments for constructing a SiteMonitor resource.
type SiteMonitorArgs struct {
	// The URL or IP address monitored by the site monitoring task.
	Address pulumi.StringInput
	// The IDs of existing alert rules to be associated with the site monitoring task.
	AlertIds pulumi.StringArrayInput
	// The monitoring interval of the site monitoring task. Unit: minutes. Valid values: 1, 5, and 15. Default value: 1.
	Interval pulumi.IntPtrInput
	// The detection points in a JSON array. For example, `[{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring.
	IspCities SiteMonitorIspCityArrayInput
	// The extended options of the protocol of the site monitoring task. The options vary according to the protocol.
	OptionsJson pulumi.StringPtrInput
	// The name of the site monitoring task. The name must be 4 to 100 characters in length. The name can contain the following types of characters: letters, digits, and underscores.
	TaskName pulumi.StringInput
	// The protocol of the site monitoring task. Currently, site monitoring supports the following protocols: HTTP, Ping, TCP, UDP, DNS, SMTP, POP3, and FTP.
	TaskType pulumi.StringInput
}

func (SiteMonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteMonitorArgs)(nil)).Elem()
}

type SiteMonitorInput interface {
	pulumi.Input

	ToSiteMonitorOutput() SiteMonitorOutput
	ToSiteMonitorOutputWithContext(ctx context.Context) SiteMonitorOutput
}

func (*SiteMonitor) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteMonitor)(nil))
}

func (i *SiteMonitor) ToSiteMonitorOutput() SiteMonitorOutput {
	return i.ToSiteMonitorOutputWithContext(context.Background())
}

func (i *SiteMonitor) ToSiteMonitorOutputWithContext(ctx context.Context) SiteMonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteMonitorOutput)
}

func (i *SiteMonitor) ToSiteMonitorPtrOutput() SiteMonitorPtrOutput {
	return i.ToSiteMonitorPtrOutputWithContext(context.Background())
}

func (i *SiteMonitor) ToSiteMonitorPtrOutputWithContext(ctx context.Context) SiteMonitorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteMonitorPtrOutput)
}

type SiteMonitorPtrInput interface {
	pulumi.Input

	ToSiteMonitorPtrOutput() SiteMonitorPtrOutput
	ToSiteMonitorPtrOutputWithContext(ctx context.Context) SiteMonitorPtrOutput
}

type siteMonitorPtrType SiteMonitorArgs

func (*siteMonitorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteMonitor)(nil))
}

func (i *siteMonitorPtrType) ToSiteMonitorPtrOutput() SiteMonitorPtrOutput {
	return i.ToSiteMonitorPtrOutputWithContext(context.Background())
}

func (i *siteMonitorPtrType) ToSiteMonitorPtrOutputWithContext(ctx context.Context) SiteMonitorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteMonitorPtrOutput)
}

// SiteMonitorArrayInput is an input type that accepts SiteMonitorArray and SiteMonitorArrayOutput values.
// You can construct a concrete instance of `SiteMonitorArrayInput` via:
//
//          SiteMonitorArray{ SiteMonitorArgs{...} }
type SiteMonitorArrayInput interface {
	pulumi.Input

	ToSiteMonitorArrayOutput() SiteMonitorArrayOutput
	ToSiteMonitorArrayOutputWithContext(context.Context) SiteMonitorArrayOutput
}

type SiteMonitorArray []SiteMonitorInput

func (SiteMonitorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SiteMonitor)(nil)).Elem()
}

func (i SiteMonitorArray) ToSiteMonitorArrayOutput() SiteMonitorArrayOutput {
	return i.ToSiteMonitorArrayOutputWithContext(context.Background())
}

func (i SiteMonitorArray) ToSiteMonitorArrayOutputWithContext(ctx context.Context) SiteMonitorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteMonitorArrayOutput)
}

// SiteMonitorMapInput is an input type that accepts SiteMonitorMap and SiteMonitorMapOutput values.
// You can construct a concrete instance of `SiteMonitorMapInput` via:
//
//          SiteMonitorMap{ "key": SiteMonitorArgs{...} }
type SiteMonitorMapInput interface {
	pulumi.Input

	ToSiteMonitorMapOutput() SiteMonitorMapOutput
	ToSiteMonitorMapOutputWithContext(context.Context) SiteMonitorMapOutput
}

type SiteMonitorMap map[string]SiteMonitorInput

func (SiteMonitorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SiteMonitor)(nil)).Elem()
}

func (i SiteMonitorMap) ToSiteMonitorMapOutput() SiteMonitorMapOutput {
	return i.ToSiteMonitorMapOutputWithContext(context.Background())
}

func (i SiteMonitorMap) ToSiteMonitorMapOutputWithContext(ctx context.Context) SiteMonitorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteMonitorMapOutput)
}

type SiteMonitorOutput struct{ *pulumi.OutputState }

func (SiteMonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteMonitor)(nil))
}

func (o SiteMonitorOutput) ToSiteMonitorOutput() SiteMonitorOutput {
	return o
}

func (o SiteMonitorOutput) ToSiteMonitorOutputWithContext(ctx context.Context) SiteMonitorOutput {
	return o
}

func (o SiteMonitorOutput) ToSiteMonitorPtrOutput() SiteMonitorPtrOutput {
	return o.ToSiteMonitorPtrOutputWithContext(context.Background())
}

func (o SiteMonitorOutput) ToSiteMonitorPtrOutputWithContext(ctx context.Context) SiteMonitorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SiteMonitor) *SiteMonitor {
		return &v
	}).(SiteMonitorPtrOutput)
}

type SiteMonitorPtrOutput struct{ *pulumi.OutputState }

func (SiteMonitorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteMonitor)(nil))
}

func (o SiteMonitorPtrOutput) ToSiteMonitorPtrOutput() SiteMonitorPtrOutput {
	return o
}

func (o SiteMonitorPtrOutput) ToSiteMonitorPtrOutputWithContext(ctx context.Context) SiteMonitorPtrOutput {
	return o
}

func (o SiteMonitorPtrOutput) Elem() SiteMonitorOutput {
	return o.ApplyT(func(v *SiteMonitor) SiteMonitor {
		if v != nil {
			return *v
		}
		var ret SiteMonitor
		return ret
	}).(SiteMonitorOutput)
}

type SiteMonitorArrayOutput struct{ *pulumi.OutputState }

func (SiteMonitorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SiteMonitor)(nil))
}

func (o SiteMonitorArrayOutput) ToSiteMonitorArrayOutput() SiteMonitorArrayOutput {
	return o
}

func (o SiteMonitorArrayOutput) ToSiteMonitorArrayOutputWithContext(ctx context.Context) SiteMonitorArrayOutput {
	return o
}

func (o SiteMonitorArrayOutput) Index(i pulumi.IntInput) SiteMonitorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SiteMonitor {
		return vs[0].([]SiteMonitor)[vs[1].(int)]
	}).(SiteMonitorOutput)
}

type SiteMonitorMapOutput struct{ *pulumi.OutputState }

func (SiteMonitorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SiteMonitor)(nil))
}

func (o SiteMonitorMapOutput) ToSiteMonitorMapOutput() SiteMonitorMapOutput {
	return o
}

func (o SiteMonitorMapOutput) ToSiteMonitorMapOutputWithContext(ctx context.Context) SiteMonitorMapOutput {
	return o
}

func (o SiteMonitorMapOutput) MapIndex(k pulumi.StringInput) SiteMonitorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SiteMonitor {
		return vs[0].(map[string]SiteMonitor)[vs[1].(string)]
	}).(SiteMonitorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SiteMonitorInput)(nil)).Elem(), &SiteMonitor{})
	pulumi.RegisterInputType(reflect.TypeOf((*SiteMonitorPtrInput)(nil)).Elem(), &SiteMonitor{})
	pulumi.RegisterInputType(reflect.TypeOf((*SiteMonitorArrayInput)(nil)).Elem(), SiteMonitorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SiteMonitorMapInput)(nil)).Elem(), SiteMonitorMap{})
	pulumi.RegisterOutputType(SiteMonitorOutput{})
	pulumi.RegisterOutputType(SiteMonitorPtrOutput{})
	pulumi.RegisterOutputType(SiteMonitorArrayOutput{})
	pulumi.RegisterOutputType(SiteMonitorMapOutput{})
}
