// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud
{
    /// <summary>
    /// The provider type for the alicloud package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/index.html.markdown.
    /// </summary>
    public partial class Provider : Pulumi.ProviderResource
    {
        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, ResourceOptions? options = null)
            : base("alicloud", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private static ResourceOptions MakeResourceOptions(ResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new ResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = ResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access key for API operations. You can retrieve this from the 'Security Management' section of the
        /// Alibaba Cloud console.
        /// </summary>
        [Input("accessKey")]
        public Input<string>? AccessKey { get; set; }

        /// <summary>
        /// The account ID for some service API operations. You can retrieve this from the 'Security Settings' section
        /// of the Alibaba Cloud console.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        [Input("assumeRole", json: true)]
        public Input<Inputs.ProviderAssumeRoleArgs>? AssumeRole { get; set; }

        /// <summary>
        /// Use this to mark a terraform configuration file source.
        /// </summary>
        [Input("configurationSource")]
        public Input<string>? ConfigurationSource { get; set; }

        /// <summary>
        /// The RAM Role Name attached on a ECS instance for API operations. You can retrieve this from the 'Access
        /// Control' section of the Alibaba Cloud console.
        /// </summary>
        [Input("ecsRoleName")]
        public Input<string>? EcsRoleName { get; set; }

        [Input("endpoints", json: true)]
        private InputList<Inputs.ProviderEndpointsArgs>? _endpoints;
        public InputList<Inputs.ProviderEndpointsArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Inputs.ProviderEndpointsArgs>());
            set => _endpoints = value;
        }

        [Input("fc")]
        public Input<string>? Fc { get; set; }

        [Input("logEndpoint")]
        public Input<string>? LogEndpoint { get; set; }

        [Input("mnsEndpoint")]
        public Input<string>? MnsEndpoint { get; set; }

        [Input("otsInstanceName")]
        public Input<string>? OtsInstanceName { get; set; }

        /// <summary>
        /// The profile for API operations. If not set, the default profile created with `aliyun configure` will be
        /// used.
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        /// <summary>
        /// The region where Alibaba Cloud operations will take place. Examples are cn-beijing, cn-hangzhou,
        /// eu-central-1, etc.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The secret key for API operations. You can retrieve this from the 'Security Management' section of the
        /// Alibaba Cloud console.
        /// </summary>
        [Input("secretKey")]
        public Input<string>? SecretKey { get; set; }

        /// <summary>
        /// security token. A security token is only required if you are using Security Token Service.
        /// </summary>
        [Input("securityToken")]
        public Input<string>? SecurityToken { get; set; }

        /// <summary>
        /// The path to the shared credentials file. If not set this defaults to ~/.aliyun/config.json
        /// </summary>
        [Input("sharedCredentialsFile")]
        public Input<string>? SharedCredentialsFile { get; set; }

        /// <summary>
        /// Skip static validation of region ID. Used by users of alternative AlibabaCloud-like APIs or users w/ access
        /// to regions that are not public (yet).
        /// </summary>
        [Input("skipRegionValidation", json: true)]
        public Input<bool>? SkipRegionValidation { get; set; }

        public ProviderArgs()
        {
            AccessKey = Utilities.GetEnv("ALICLOUD_ACCESS_KEY");
            AccountId = Utilities.GetEnv("ALICLOUD_ACCOUNT_ID");
            EcsRoleName = Utilities.GetEnv("ALICLOUD_ECS_ROLE_NAME");
            Profile = Utilities.GetEnv("ALICLOUD_PROFILE");
            Region = Utilities.GetEnv("ALICLOUD_REGION");
            SecretKey = Utilities.GetEnv("ALICLOUD_SECRET_KEY");
            SecurityToken = Utilities.GetEnv("ALICLOUD_SECURITY_TOKEN");
            SharedCredentialsFile = Utilities.GetEnv("ALICLOUD_SHARED_CREDENTIALS_FILE");
        }
    }

    namespace Inputs
    {

    public sealed class ProviderAssumeRoleArgs : Pulumi.ResourceArgs
    {
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("sessionExpiration")]
        public Input<int>? SessionExpiration { get; set; }

        [Input("sessionName")]
        public Input<string>? SessionName { get; set; }

        public ProviderAssumeRoleArgs()
        {
        }
    }

    public sealed class ProviderEndpointsArgs : Pulumi.ResourceArgs
    {
        [Input("actiontrail")]
        public Input<string>? Actiontrail { get; set; }

        [Input("adb")]
        public Input<string>? Adb { get; set; }

        [Input("alikafka")]
        public Input<string>? Alikafka { get; set; }

        [Input("apigateway")]
        public Input<string>? Apigateway { get; set; }

        [Input("bssopenapi")]
        public Input<string>? Bssopenapi { get; set; }

        [Input("cas")]
        public Input<string>? Cas { get; set; }

        [Input("cdn")]
        public Input<string>? Cdn { get; set; }

        [Input("cen")]
        public Input<string>? Cen { get; set; }

        [Input("cms")]
        public Input<string>? Cms { get; set; }

        [Input("cr")]
        public Input<string>? Cr { get; set; }

        [Input("cs")]
        public Input<string>? Cs { get; set; }

        [Input("datahub")]
        public Input<string>? Datahub { get; set; }

        [Input("ddosbgp")]
        public Input<string>? Ddosbgp { get; set; }

        [Input("ddoscoo")]
        public Input<string>? Ddoscoo { get; set; }

        [Input("dds")]
        public Input<string>? Dds { get; set; }

        [Input("dns")]
        public Input<string>? Dns { get; set; }

        [Input("drds")]
        public Input<string>? Drds { get; set; }

        [Input("ecs")]
        public Input<string>? Ecs { get; set; }

        [Input("elasticsearch")]
        public Input<string>? Elasticsearch { get; set; }

        [Input("emr")]
        public Input<string>? Emr { get; set; }

        [Input("ess")]
        public Input<string>? Ess { get; set; }

        [Input("fc")]
        public Input<string>? Fc { get; set; }

        [Input("gpdb")]
        public Input<string>? Gpdb { get; set; }

        [Input("kms")]
        public Input<string>? Kms { get; set; }

        [Input("kvstore")]
        public Input<string>? Kvstore { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("log")]
        public Input<string>? Log { get; set; }

        [Input("market")]
        public Input<string>? Market { get; set; }

        [Input("mns")]
        public Input<string>? Mns { get; set; }

        [Input("nas")]
        public Input<string>? Nas { get; set; }

        [Input("ons")]
        public Input<string>? Ons { get; set; }

        [Input("oss")]
        public Input<string>? Oss { get; set; }

        [Input("ots")]
        public Input<string>? Ots { get; set; }

        [Input("polardb")]
        public Input<string>? Polardb { get; set; }

        [Input("pvtz")]
        public Input<string>? Pvtz { get; set; }

        [Input("ram")]
        public Input<string>? Ram { get; set; }

        [Input("rds")]
        public Input<string>? Rds { get; set; }

        [Input("slb")]
        public Input<string>? Slb { get; set; }

        [Input("sts")]
        public Input<string>? Sts { get; set; }

        [Input("vpc")]
        public Input<string>? Vpc { get; set; }

        public ProviderEndpointsArgs()
        {
        }
    }
    }
}
