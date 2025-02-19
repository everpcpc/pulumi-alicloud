// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package polardb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a PolarDB endpoint resource to manage endpoint of PolarDB cluster.
//
// > **NOTE:** After v1.80.0 and before v1.121.0, you can only use this resource to manage the custom endpoint. Since v1.121.0, you also can import the primary endpoint and the cluster endpoint, to modify their ssl status and so on.
//
// > **NOTE:** The primary endpoint and the default cluster endpoint can not be created or deleted manually.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/polardb"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			creation := "PolarDB"
//			if param := cfg.Get("creation"); param != "" {
//				creation = param
//			}
//			name := "polardbconnectionbasic"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef(creation),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
//				VpcId:       defaultNetwork.ID(),
//				CidrBlock:   pulumi.String("172.16.0.0/24"),
//				ZoneId:      *pulumi.String(defaultZones.Zones[0].Id),
//				VswitchName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultCluster, err := polardb.NewCluster(ctx, "defaultCluster", &polardb.ClusterArgs{
//				DbType:      pulumi.String("MySQL"),
//				DbVersion:   pulumi.String("8.0"),
//				PayType:     pulumi.String("PostPaid"),
//				DbNodeClass: pulumi.String("polar.mysql.x4.large"),
//				VswitchId:   defaultSwitch.ID(),
//				Description: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = polardb.NewEndpoint(ctx, "endpoint", &polardb.EndpointArgs{
//				DbClusterId:  defaultCluster.ID(),
//				EndpointType: pulumi.String("Custom"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Argument Reference
//
// The following arguments are supported:
//
//   - `dbClusterId` - (Required, ForceNew) The Id of cluster that can run database.
//   - `endpointType` - (Optional, ForceNew) Type of the endpoint. Before v1.121.0, it only can be `Custom`. since v1.121.0, `Custom`, `Cluster`, `Primary` are valid, default to `Custom`. However when creating a new endpoint, it also only can be `Custom`.
//   - `readWriteMode` - (Optional) Read or write mode. Valid values are `ReadWrite`, `ReadOnly`. When creating a new custom endpoint, default to `ReadOnly`.
//   - `nodes` - (Optional) Node id list for endpoint configuration. At least 2 nodes if specified, or if the cluster has more than 3 nodes, read-only endpoint is allowed to mount only one node. Default is all nodes.
//   - `autoAddNewNodes` - (Optional) Whether the new node automatically joins the default cluster address. Valid values are `Enable`, `Disable`. When creating a new custom endpoint, default to `Disable`.
//   - `endpointConfig` - (Optional) The advanced settings of the endpoint of Apsara PolarDB clusters are in JSON format. Including the settings of consistency level, transaction splitting, connection pool, and offload reads from primary node. For more details, see the [description of EndpointConfig in the Request parameters table for details](https://www.alibabacloud.com/help/doc-detail/116593.htm).
//   - `sslEnabled` - (Optional, Available in v1.121.0+) Specifies how to modify the SSL encryption status. Valid values: `Disable`, `Enable`, `Update`.
//   - `netType` - (Optional, Available in v1.121.0+) The network type of the endpoint address.
//   - `sslAutoRotate` - (Available in v1.132.0+) Specifies whether automatic rotation of SSL certificates is enabled. Valid values: `Enable`,`Disable`.
//   - `sslCertificateUrl` - (Available in v1.132.0+) Specifies SSL certificate download link.\
//     **NOTE:** For a PolarDB for MySQL cluster, this parameter is required, and only one connection string in each endpoint can enable the ssl, for other notes, see [Configure SSL encryption](https://www.alibabacloud.com/help/doc-detail/153182.htm).\
//     For a PolarDB for PostgreSQL cluster or a PolarDB-O cluster, this parameter is not required, by default, SSL encryption is enabled for all endpoints.
//   - `dbEndpointDescription` - (Optional, Available in v1.201.0+) The name of the endpoint.
//
// ## Import
//
// PolarDB endpoint can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:polardb/endpoint:Endpoint example pc-abc123456:pe-abc123456
//
// ```
type Endpoint struct {
	pulumi.CustomResourceState

	AutoAddNewNodes       pulumi.StringOutput    `pulumi:"autoAddNewNodes"`
	DbClusterId           pulumi.StringOutput    `pulumi:"dbClusterId"`
	DbEndpointDescription pulumi.StringPtrOutput `pulumi:"dbEndpointDescription"`
	// (Available in v1.161.0+) The ID of the cluster endpoint.
	DbEndpointId   pulumi.StringOutput `pulumi:"dbEndpointId"`
	EndpointConfig pulumi.MapOutput    `pulumi:"endpointConfig"`
	// Type of endpoint.
	EndpointType      pulumi.StringPtrOutput   `pulumi:"endpointType"`
	NetType           pulumi.StringPtrOutput   `pulumi:"netType"`
	Nodes             pulumi.StringArrayOutput `pulumi:"nodes"`
	ReadWriteMode     pulumi.StringOutput      `pulumi:"readWriteMode"`
	SslAutoRotate     pulumi.StringPtrOutput   `pulumi:"sslAutoRotate"`
	SslCertificateUrl pulumi.StringOutput      `pulumi:"sslCertificateUrl"`
	// (Available in v1.121.0+) The SSL connection string.
	SslConnectionString pulumi.StringOutput    `pulumi:"sslConnectionString"`
	SslEnabled          pulumi.StringPtrOutput `pulumi:"sslEnabled"`
	// (Available in v1.121.0+) The time when the SSL certificate expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	SslExpireTime pulumi.StringOutput `pulumi:"sslExpireTime"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbClusterId == nil {
		return nil, errors.New("invalid value for required argument 'DbClusterId'")
	}
	var resource Endpoint
	err := ctx.RegisterResource("alicloud:polardb/endpoint:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("alicloud:polardb/endpoint:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
	AutoAddNewNodes       *string `pulumi:"autoAddNewNodes"`
	DbClusterId           *string `pulumi:"dbClusterId"`
	DbEndpointDescription *string `pulumi:"dbEndpointDescription"`
	// (Available in v1.161.0+) The ID of the cluster endpoint.
	DbEndpointId   *string                `pulumi:"dbEndpointId"`
	EndpointConfig map[string]interface{} `pulumi:"endpointConfig"`
	// Type of endpoint.
	EndpointType      *string  `pulumi:"endpointType"`
	NetType           *string  `pulumi:"netType"`
	Nodes             []string `pulumi:"nodes"`
	ReadWriteMode     *string  `pulumi:"readWriteMode"`
	SslAutoRotate     *string  `pulumi:"sslAutoRotate"`
	SslCertificateUrl *string  `pulumi:"sslCertificateUrl"`
	// (Available in v1.121.0+) The SSL connection string.
	SslConnectionString *string `pulumi:"sslConnectionString"`
	SslEnabled          *string `pulumi:"sslEnabled"`
	// (Available in v1.121.0+) The time when the SSL certificate expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	SslExpireTime *string `pulumi:"sslExpireTime"`
}

type EndpointState struct {
	AutoAddNewNodes       pulumi.StringPtrInput
	DbClusterId           pulumi.StringPtrInput
	DbEndpointDescription pulumi.StringPtrInput
	// (Available in v1.161.0+) The ID of the cluster endpoint.
	DbEndpointId   pulumi.StringPtrInput
	EndpointConfig pulumi.MapInput
	// Type of endpoint.
	EndpointType      pulumi.StringPtrInput
	NetType           pulumi.StringPtrInput
	Nodes             pulumi.StringArrayInput
	ReadWriteMode     pulumi.StringPtrInput
	SslAutoRotate     pulumi.StringPtrInput
	SslCertificateUrl pulumi.StringPtrInput
	// (Available in v1.121.0+) The SSL connection string.
	SslConnectionString pulumi.StringPtrInput
	SslEnabled          pulumi.StringPtrInput
	// (Available in v1.121.0+) The time when the SSL certificate expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	SslExpireTime pulumi.StringPtrInput
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	AutoAddNewNodes       *string                `pulumi:"autoAddNewNodes"`
	DbClusterId           string                 `pulumi:"dbClusterId"`
	DbEndpointDescription *string                `pulumi:"dbEndpointDescription"`
	EndpointConfig        map[string]interface{} `pulumi:"endpointConfig"`
	// Type of endpoint.
	EndpointType  *string  `pulumi:"endpointType"`
	NetType       *string  `pulumi:"netType"`
	Nodes         []string `pulumi:"nodes"`
	ReadWriteMode *string  `pulumi:"readWriteMode"`
	SslAutoRotate *string  `pulumi:"sslAutoRotate"`
	SslEnabled    *string  `pulumi:"sslEnabled"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	AutoAddNewNodes       pulumi.StringPtrInput
	DbClusterId           pulumi.StringInput
	DbEndpointDescription pulumi.StringPtrInput
	EndpointConfig        pulumi.MapInput
	// Type of endpoint.
	EndpointType  pulumi.StringPtrInput
	NetType       pulumi.StringPtrInput
	Nodes         pulumi.StringArrayInput
	ReadWriteMode pulumi.StringPtrInput
	SslAutoRotate pulumi.StringPtrInput
	SslEnabled    pulumi.StringPtrInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}

type EndpointInput interface {
	pulumi.Input

	ToEndpointOutput() EndpointOutput
	ToEndpointOutputWithContext(ctx context.Context) EndpointOutput
}

func (*Endpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (i *Endpoint) ToEndpointOutput() EndpointOutput {
	return i.ToEndpointOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointOutput)
}

// EndpointArrayInput is an input type that accepts EndpointArray and EndpointArrayOutput values.
// You can construct a concrete instance of `EndpointArrayInput` via:
//
//	EndpointArray{ EndpointArgs{...} }
type EndpointArrayInput interface {
	pulumi.Input

	ToEndpointArrayOutput() EndpointArrayOutput
	ToEndpointArrayOutputWithContext(context.Context) EndpointArrayOutput
}

type EndpointArray []EndpointInput

func (EndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Endpoint)(nil)).Elem()
}

func (i EndpointArray) ToEndpointArrayOutput() EndpointArrayOutput {
	return i.ToEndpointArrayOutputWithContext(context.Background())
}

func (i EndpointArray) ToEndpointArrayOutputWithContext(ctx context.Context) EndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointArrayOutput)
}

// EndpointMapInput is an input type that accepts EndpointMap and EndpointMapOutput values.
// You can construct a concrete instance of `EndpointMapInput` via:
//
//	EndpointMap{ "key": EndpointArgs{...} }
type EndpointMapInput interface {
	pulumi.Input

	ToEndpointMapOutput() EndpointMapOutput
	ToEndpointMapOutputWithContext(context.Context) EndpointMapOutput
}

type EndpointMap map[string]EndpointInput

func (EndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Endpoint)(nil)).Elem()
}

func (i EndpointMap) ToEndpointMapOutput() EndpointMapOutput {
	return i.ToEndpointMapOutputWithContext(context.Background())
}

func (i EndpointMap) ToEndpointMapOutputWithContext(ctx context.Context) EndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointMapOutput)
}

type EndpointOutput struct{ *pulumi.OutputState }

func (EndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

func (o EndpointOutput) AutoAddNewNodes() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.AutoAddNewNodes }).(pulumi.StringOutput)
}

func (o EndpointOutput) DbClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.DbClusterId }).(pulumi.StringOutput)
}

func (o EndpointOutput) DbEndpointDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.DbEndpointDescription }).(pulumi.StringPtrOutput)
}

// (Available in v1.161.0+) The ID of the cluster endpoint.
func (o EndpointOutput) DbEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.DbEndpointId }).(pulumi.StringOutput)
}

func (o EndpointOutput) EndpointConfig() pulumi.MapOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.MapOutput { return v.EndpointConfig }).(pulumi.MapOutput)
}

// Type of endpoint.
func (o EndpointOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.EndpointType }).(pulumi.StringPtrOutput)
}

func (o EndpointOutput) NetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.NetType }).(pulumi.StringPtrOutput)
}

func (o EndpointOutput) Nodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringArrayOutput { return v.Nodes }).(pulumi.StringArrayOutput)
}

func (o EndpointOutput) ReadWriteMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.ReadWriteMode }).(pulumi.StringOutput)
}

func (o EndpointOutput) SslAutoRotate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.SslAutoRotate }).(pulumi.StringPtrOutput)
}

func (o EndpointOutput) SslCertificateUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.SslCertificateUrl }).(pulumi.StringOutput)
}

// (Available in v1.121.0+) The SSL connection string.
func (o EndpointOutput) SslConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.SslConnectionString }).(pulumi.StringOutput)
}

func (o EndpointOutput) SslEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.SslEnabled }).(pulumi.StringPtrOutput)
}

// (Available in v1.121.0+) The time when the SSL certificate expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
func (o EndpointOutput) SslExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.SslExpireTime }).(pulumi.StringOutput)
}

type EndpointArrayOutput struct{ *pulumi.OutputState }

func (EndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Endpoint)(nil)).Elem()
}

func (o EndpointArrayOutput) ToEndpointArrayOutput() EndpointArrayOutput {
	return o
}

func (o EndpointArrayOutput) ToEndpointArrayOutputWithContext(ctx context.Context) EndpointArrayOutput {
	return o
}

func (o EndpointArrayOutput) Index(i pulumi.IntInput) EndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Endpoint {
		return vs[0].([]*Endpoint)[vs[1].(int)]
	}).(EndpointOutput)
}

type EndpointMapOutput struct{ *pulumi.OutputState }

func (EndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Endpoint)(nil)).Elem()
}

func (o EndpointMapOutput) ToEndpointMapOutput() EndpointMapOutput {
	return o
}

func (o EndpointMapOutput) ToEndpointMapOutputWithContext(ctx context.Context) EndpointMapOutput {
	return o
}

func (o EndpointMapOutput) MapIndex(k pulumi.StringInput) EndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Endpoint {
		return vs[0].(map[string]*Endpoint)[vs[1].(string)]
	}).(EndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointInput)(nil)).Elem(), &Endpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointArrayInput)(nil)).Elem(), EndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointMapInput)(nil)).Elem(), EndpointMap{})
	pulumi.RegisterOutputType(EndpointOutput{})
	pulumi.RegisterOutputType(EndpointArrayOutput{})
	pulumi.RegisterOutputType(EndpointMapOutput{})
}
