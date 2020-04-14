// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticsearch

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Elasticsearch instance resource. It contains data nodes, dedicated master node(optional) and etc. It can be associated with private IP whitelists and kibana IP whitelist.
//
// > **NOTE:** Only one operation is supported in a request. So if `dataNodeSpec` and `dataNodeDiskSize` are both changed, system will respond error.
//
// > **NOTE:** At present, `version` can not be modified once instance has been created.
type Instance struct {
	pulumi.CustomResourceState

	// The Elasticsearch cluster's data node quantity, between 2 and 50.
	DataNodeAmount pulumi.IntOutput `pulumi:"dataNodeAmount"`
	// The single data node storage space.
	// - `cloudSsd`: An SSD disk, supports a maximum of 2048 GiB (2 TB).
	// - `cloudEfficiency` An ultra disk, supports a maximum of 5120 GiB (5 TB). If the data to be stored is larger than 2048 GiB, an ultra disk can only support the following data sizes (GiB): [`2560`, `3072`, `3584`, `4096`, `4608`, `5120`].
	DataNodeDiskSize pulumi.IntOutput `pulumi:"dataNodeDiskSize"`
	// The data node disk type. Supported values: cloud_ssd, cloud_efficiency.
	DataNodeDiskType pulumi.StringOutput `pulumi:"dataNodeDiskType"`
	// The data node specifications of the Elasticsearch instance.
	DataNodeSpec pulumi.StringOutput `pulumi:"dataNodeSpec"`
	// The description of instance. It a string of 0 to 30 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Instance connection domain (only VPC network access supported).
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. From version 1.69.0, the Elasticsearch cluster allows you to update your instanceChargeYpe from `PostPaid` to `PrePaid`, the following attributes are required: `period`. But, updating from `PostPaid` to `PrePaid` is not supported.
	InstanceChargeType pulumi.StringPtrOutput `pulumi:"instanceChargeType"`
	// Kibana console domain (Internet access supported).
	KibanaDomain pulumi.StringOutput `pulumi:"kibanaDomain"`
	// Kibana console port.
	KibanaPort pulumi.IntOutput `pulumi:"kibanaPort"`
	// Set the Kibana's IP whitelist in internet network.
	KibanaWhitelists pulumi.StringArrayOutput `pulumi:"kibanaWhitelists"`
	// An KMS encrypts password used to a instance. If the `password` is filled in, this field will be ignored, but you have to specify one of `password` and `kmsEncryptedPassword` fields.
	KmsEncryptedPassword pulumi.StringPtrOutput `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapOutput `pulumi:"kmsEncryptionContext"`
	// The dedicated master node spec. If specified, dedicated master node will be created.
	MasterNodeSpec pulumi.StringPtrOutput `pulumi:"masterNodeSpec"`
	// The password of the instance. The password can be 8 to 30 characters in length and must contain three of the following conditions: uppercase letters, lowercase letters, numbers, and special characters (`!@#$%^&*()_+-=`).
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The duration that you will buy Elasticsearch instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1. From version 1.69.2, when to modify this value, the resource can renewal a `PrePaid` instance.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Instance connection port.
	Port pulumi.IntOutput `pulumi:"port"`
	// Set the instance's IP whitelist in VPC network.
	PrivateWhitelists pulumi.StringArrayOutput `pulumi:"privateWhitelists"`
	PublicWhitelists  pulumi.StringArrayOutput `pulumi:"publicWhitelists"`
	// The Elasticsearch instance status. Includes `active`, `activating`, `inactive`. Some operations are denied when status is not `active`.
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	// - key: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:". It cannot contain "http://" and "https://". It cannot be a null string.
	// - value: It can be up to 128 characters in length. It cannot contain "http://" and "https://". It can be a null string.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Elasticsearch version. Supported values: `5.5.3_with_X-Pack`, `6.3_with_X-Pack` and `6.7_with_X-Pack`.
	Version pulumi.StringOutput `pulumi:"version"`
	// The ID of VSwitch.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
	// The Multi-AZ supported for Elasticsearch, between 1 and 3. The `dataNodeAmount` value must be an integral multiple of the `zoneCount` value.
	ZoneCount pulumi.IntPtrOutput `pulumi:"zoneCount"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil || args.DataNodeAmount == nil {
		return nil, errors.New("missing required argument 'DataNodeAmount'")
	}
	if args == nil || args.DataNodeDiskSize == nil {
		return nil, errors.New("missing required argument 'DataNodeDiskSize'")
	}
	if args == nil || args.DataNodeDiskType == nil {
		return nil, errors.New("missing required argument 'DataNodeDiskType'")
	}
	if args == nil || args.DataNodeSpec == nil {
		return nil, errors.New("missing required argument 'DataNodeSpec'")
	}
	if args == nil || args.Version == nil {
		return nil, errors.New("missing required argument 'Version'")
	}
	if args == nil || args.VswitchId == nil {
		return nil, errors.New("missing required argument 'VswitchId'")
	}
	if args == nil {
		args = &InstanceArgs{}
	}
	var resource Instance
	err := ctx.RegisterResource("alicloud:elasticsearch/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("alicloud:elasticsearch/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The Elasticsearch cluster's data node quantity, between 2 and 50.
	DataNodeAmount *int `pulumi:"dataNodeAmount"`
	// The single data node storage space.
	// - `cloudSsd`: An SSD disk, supports a maximum of 2048 GiB (2 TB).
	// - `cloudEfficiency` An ultra disk, supports a maximum of 5120 GiB (5 TB). If the data to be stored is larger than 2048 GiB, an ultra disk can only support the following data sizes (GiB): [`2560`, `3072`, `3584`, `4096`, `4608`, `5120`].
	DataNodeDiskSize *int `pulumi:"dataNodeDiskSize"`
	// The data node disk type. Supported values: cloud_ssd, cloud_efficiency.
	DataNodeDiskType *string `pulumi:"dataNodeDiskType"`
	// The data node specifications of the Elasticsearch instance.
	DataNodeSpec *string `pulumi:"dataNodeSpec"`
	// The description of instance. It a string of 0 to 30 characters.
	Description *string `pulumi:"description"`
	// Instance connection domain (only VPC network access supported).
	Domain *string `pulumi:"domain"`
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. From version 1.69.0, the Elasticsearch cluster allows you to update your instanceChargeYpe from `PostPaid` to `PrePaid`, the following attributes are required: `period`. But, updating from `PostPaid` to `PrePaid` is not supported.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// Kibana console domain (Internet access supported).
	KibanaDomain *string `pulumi:"kibanaDomain"`
	// Kibana console port.
	KibanaPort *int `pulumi:"kibanaPort"`
	// Set the Kibana's IP whitelist in internet network.
	KibanaWhitelists []string `pulumi:"kibanaWhitelists"`
	// An KMS encrypts password used to a instance. If the `password` is filled in, this field will be ignored, but you have to specify one of `password` and `kmsEncryptedPassword` fields.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// The dedicated master node spec. If specified, dedicated master node will be created.
	MasterNodeSpec *string `pulumi:"masterNodeSpec"`
	// The password of the instance. The password can be 8 to 30 characters in length and must contain three of the following conditions: uppercase letters, lowercase letters, numbers, and special characters (`!@#$%^&*()_+-=`).
	Password *string `pulumi:"password"`
	// The duration that you will buy Elasticsearch instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1. From version 1.69.2, when to modify this value, the resource can renewal a `PrePaid` instance.
	Period *int `pulumi:"period"`
	// Instance connection port.
	Port *int `pulumi:"port"`
	// Set the instance's IP whitelist in VPC network.
	PrivateWhitelists []string `pulumi:"privateWhitelists"`
	PublicWhitelists  []string `pulumi:"publicWhitelists"`
	// The Elasticsearch instance status. Includes `active`, `activating`, `inactive`. Some operations are denied when status is not `active`.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	// - key: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:". It cannot contain "http://" and "https://". It cannot be a null string.
	// - value: It can be up to 128 characters in length. It cannot contain "http://" and "https://". It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
	// Elasticsearch version. Supported values: `5.5.3_with_X-Pack`, `6.3_with_X-Pack` and `6.7_with_X-Pack`.
	Version *string `pulumi:"version"`
	// The ID of VSwitch.
	VswitchId *string `pulumi:"vswitchId"`
	// The Multi-AZ supported for Elasticsearch, between 1 and 3. The `dataNodeAmount` value must be an integral multiple of the `zoneCount` value.
	ZoneCount *int `pulumi:"zoneCount"`
}

type InstanceState struct {
	// The Elasticsearch cluster's data node quantity, between 2 and 50.
	DataNodeAmount pulumi.IntPtrInput
	// The single data node storage space.
	// - `cloudSsd`: An SSD disk, supports a maximum of 2048 GiB (2 TB).
	// - `cloudEfficiency` An ultra disk, supports a maximum of 5120 GiB (5 TB). If the data to be stored is larger than 2048 GiB, an ultra disk can only support the following data sizes (GiB): [`2560`, `3072`, `3584`, `4096`, `4608`, `5120`].
	DataNodeDiskSize pulumi.IntPtrInput
	// The data node disk type. Supported values: cloud_ssd, cloud_efficiency.
	DataNodeDiskType pulumi.StringPtrInput
	// The data node specifications of the Elasticsearch instance.
	DataNodeSpec pulumi.StringPtrInput
	// The description of instance. It a string of 0 to 30 characters.
	Description pulumi.StringPtrInput
	// Instance connection domain (only VPC network access supported).
	Domain pulumi.StringPtrInput
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. From version 1.69.0, the Elasticsearch cluster allows you to update your instanceChargeYpe from `PostPaid` to `PrePaid`, the following attributes are required: `period`. But, updating from `PostPaid` to `PrePaid` is not supported.
	InstanceChargeType pulumi.StringPtrInput
	// Kibana console domain (Internet access supported).
	KibanaDomain pulumi.StringPtrInput
	// Kibana console port.
	KibanaPort pulumi.IntPtrInput
	// Set the Kibana's IP whitelist in internet network.
	KibanaWhitelists pulumi.StringArrayInput
	// An KMS encrypts password used to a instance. If the `password` is filled in, this field will be ignored, but you have to specify one of `password` and `kmsEncryptedPassword` fields.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// The dedicated master node spec. If specified, dedicated master node will be created.
	MasterNodeSpec pulumi.StringPtrInput
	// The password of the instance. The password can be 8 to 30 characters in length and must contain three of the following conditions: uppercase letters, lowercase letters, numbers, and special characters (`!@#$%^&*()_+-=`).
	Password pulumi.StringPtrInput
	// The duration that you will buy Elasticsearch instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1. From version 1.69.2, when to modify this value, the resource can renewal a `PrePaid` instance.
	Period pulumi.IntPtrInput
	// Instance connection port.
	Port pulumi.IntPtrInput
	// Set the instance's IP whitelist in VPC network.
	PrivateWhitelists pulumi.StringArrayInput
	PublicWhitelists  pulumi.StringArrayInput
	// The Elasticsearch instance status. Includes `active`, `activating`, `inactive`. Some operations are denied when status is not `active`.
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	// - key: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:". It cannot contain "http://" and "https://". It cannot be a null string.
	// - value: It can be up to 128 characters in length. It cannot contain "http://" and "https://". It can be a null string.
	Tags pulumi.MapInput
	// Elasticsearch version. Supported values: `5.5.3_with_X-Pack`, `6.3_with_X-Pack` and `6.7_with_X-Pack`.
	Version pulumi.StringPtrInput
	// The ID of VSwitch.
	VswitchId pulumi.StringPtrInput
	// The Multi-AZ supported for Elasticsearch, between 1 and 3. The `dataNodeAmount` value must be an integral multiple of the `zoneCount` value.
	ZoneCount pulumi.IntPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The Elasticsearch cluster's data node quantity, between 2 and 50.
	DataNodeAmount int `pulumi:"dataNodeAmount"`
	// The single data node storage space.
	// - `cloudSsd`: An SSD disk, supports a maximum of 2048 GiB (2 TB).
	// - `cloudEfficiency` An ultra disk, supports a maximum of 5120 GiB (5 TB). If the data to be stored is larger than 2048 GiB, an ultra disk can only support the following data sizes (GiB): [`2560`, `3072`, `3584`, `4096`, `4608`, `5120`].
	DataNodeDiskSize int `pulumi:"dataNodeDiskSize"`
	// The data node disk type. Supported values: cloud_ssd, cloud_efficiency.
	DataNodeDiskType string `pulumi:"dataNodeDiskType"`
	// The data node specifications of the Elasticsearch instance.
	DataNodeSpec string `pulumi:"dataNodeSpec"`
	// The description of instance. It a string of 0 to 30 characters.
	Description *string `pulumi:"description"`
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. From version 1.69.0, the Elasticsearch cluster allows you to update your instanceChargeYpe from `PostPaid` to `PrePaid`, the following attributes are required: `period`. But, updating from `PostPaid` to `PrePaid` is not supported.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// Set the Kibana's IP whitelist in internet network.
	KibanaWhitelists []string `pulumi:"kibanaWhitelists"`
	// An KMS encrypts password used to a instance. If the `password` is filled in, this field will be ignored, but you have to specify one of `password` and `kmsEncryptedPassword` fields.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// The dedicated master node spec. If specified, dedicated master node will be created.
	MasterNodeSpec *string `pulumi:"masterNodeSpec"`
	// The password of the instance. The password can be 8 to 30 characters in length and must contain three of the following conditions: uppercase letters, lowercase letters, numbers, and special characters (`!@#$%^&*()_+-=`).
	Password *string `pulumi:"password"`
	// The duration that you will buy Elasticsearch instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1. From version 1.69.2, when to modify this value, the resource can renewal a `PrePaid` instance.
	Period *int `pulumi:"period"`
	// Set the instance's IP whitelist in VPC network.
	PrivateWhitelists []string `pulumi:"privateWhitelists"`
	PublicWhitelists  []string `pulumi:"publicWhitelists"`
	// A mapping of tags to assign to the resource.
	// - key: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:". It cannot contain "http://" and "https://". It cannot be a null string.
	// - value: It can be up to 128 characters in length. It cannot contain "http://" and "https://". It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
	// Elasticsearch version. Supported values: `5.5.3_with_X-Pack`, `6.3_with_X-Pack` and `6.7_with_X-Pack`.
	Version string `pulumi:"version"`
	// The ID of VSwitch.
	VswitchId string `pulumi:"vswitchId"`
	// The Multi-AZ supported for Elasticsearch, between 1 and 3. The `dataNodeAmount` value must be an integral multiple of the `zoneCount` value.
	ZoneCount *int `pulumi:"zoneCount"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The Elasticsearch cluster's data node quantity, between 2 and 50.
	DataNodeAmount pulumi.IntInput
	// The single data node storage space.
	// - `cloudSsd`: An SSD disk, supports a maximum of 2048 GiB (2 TB).
	// - `cloudEfficiency` An ultra disk, supports a maximum of 5120 GiB (5 TB). If the data to be stored is larger than 2048 GiB, an ultra disk can only support the following data sizes (GiB): [`2560`, `3072`, `3584`, `4096`, `4608`, `5120`].
	DataNodeDiskSize pulumi.IntInput
	// The data node disk type. Supported values: cloud_ssd, cloud_efficiency.
	DataNodeDiskType pulumi.StringInput
	// The data node specifications of the Elasticsearch instance.
	DataNodeSpec pulumi.StringInput
	// The description of instance. It a string of 0 to 30 characters.
	Description pulumi.StringPtrInput
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. From version 1.69.0, the Elasticsearch cluster allows you to update your instanceChargeYpe from `PostPaid` to `PrePaid`, the following attributes are required: `period`. But, updating from `PostPaid` to `PrePaid` is not supported.
	InstanceChargeType pulumi.StringPtrInput
	// Set the Kibana's IP whitelist in internet network.
	KibanaWhitelists pulumi.StringArrayInput
	// An KMS encrypts password used to a instance. If the `password` is filled in, this field will be ignored, but you have to specify one of `password` and `kmsEncryptedPassword` fields.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// The dedicated master node spec. If specified, dedicated master node will be created.
	MasterNodeSpec pulumi.StringPtrInput
	// The password of the instance. The password can be 8 to 30 characters in length and must contain three of the following conditions: uppercase letters, lowercase letters, numbers, and special characters (`!@#$%^&*()_+-=`).
	Password pulumi.StringPtrInput
	// The duration that you will buy Elasticsearch instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1. From version 1.69.2, when to modify this value, the resource can renewal a `PrePaid` instance.
	Period pulumi.IntPtrInput
	// Set the instance's IP whitelist in VPC network.
	PrivateWhitelists pulumi.StringArrayInput
	PublicWhitelists  pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	// - key: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:". It cannot contain "http://" and "https://". It cannot be a null string.
	// - value: It can be up to 128 characters in length. It cannot contain "http://" and "https://". It can be a null string.
	Tags pulumi.MapInput
	// Elasticsearch version. Supported values: `5.5.3_with_X-Pack`, `6.3_with_X-Pack` and `6.7_with_X-Pack`.
	Version pulumi.StringInput
	// The ID of VSwitch.
	VswitchId pulumi.StringInput
	// The Multi-AZ supported for Elasticsearch, between 1 and 3. The `dataNodeAmount` value must be an integral multiple of the `zoneCount` value.
	ZoneCount pulumi.IntPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}
