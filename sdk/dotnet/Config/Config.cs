// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.AliCloud
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("alicloud");
        /// <summary>
        /// The access key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
        /// console.
        /// </summary>
        public static string? AccessKey { get; set; } = __config.Get("accessKey") ?? Utilities.GetEnv("ALICLOUD_ACCESS_KEY");

        /// <summary>
        /// The account ID for some service API operations. You can retrieve this from the 'Security Settings' section of the
        /// Alibaba Cloud console.
        /// </summary>
        public static string? AccountId { get; set; } = __config.Get("accountId") ?? Utilities.GetEnv("ALICLOUD_ACCOUNT_ID");

        public static Pulumi.AliCloud.Config.Types.AssumeRole? AssumeRole { get; set; } = __config.GetObject<Pulumi.AliCloud.Config.Types.AssumeRole>("assumeRole");

        /// <summary>
        /// Use this to mark a terraform configuration file source.
        /// </summary>
        public static string? ConfigurationSource { get; set; } = __config.Get("configurationSource");

        /// <summary>
        /// The RAM Role Name attached on a ECS instance for API operations. You can retrieve this from the 'Access Control' section
        /// of the Alibaba Cloud console.
        /// </summary>
        public static string? EcsRoleName { get; set; } = __config.Get("ecsRoleName") ?? Utilities.GetEnv("ALICLOUD_ECS_ROLE_NAME");

        public static ImmutableArray<Pulumi.AliCloud.Config.Types.Endpoints> Endpoints { get; set; } = __config.GetObject<ImmutableArray<Pulumi.AliCloud.Config.Types.Endpoints>>("endpoints");

        public static string? Fc { get; set; } = __config.Get("fc");

        public static string? LogEndpoint { get; set; } = __config.Get("logEndpoint");

        public static string? MnsEndpoint { get; set; } = __config.Get("mnsEndpoint");

        public static string? OtsInstanceName { get; set; } = __config.Get("otsInstanceName");

        /// <summary>
        /// The profile for API operations. If not set, the default profile created with `aliyun configure` will be used.
        /// </summary>
        public static string? Profile { get; set; } = __config.Get("profile") ?? Utilities.GetEnv("ALICLOUD_PROFILE");

        public static string? Protocol { get; set; } = __config.Get("protocol");

        /// <summary>
        /// The region where Alibaba Cloud operations will take place. Examples are cn-beijing, cn-hangzhou, eu-central-1, etc.
        /// </summary>
        public static string? Region { get; set; } = __config.Get("region") ?? Utilities.GetEnv("ALICLOUD_REGION");

        /// <summary>
        /// The secret key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
        /// console.
        /// </summary>
        public static string? SecretKey { get; set; } = __config.Get("secretKey") ?? Utilities.GetEnv("ALICLOUD_SECRET_KEY");

        /// <summary>
        /// security token. A security token is only required if you are using Security Token Service.
        /// </summary>
        public static string? SecurityToken { get; set; } = __config.Get("securityToken") ?? Utilities.GetEnv("ALICLOUD_SECURITY_TOKEN");

        /// <summary>
        /// The path to the shared credentials file. If not set this defaults to ~/.aliyun/config.json
        /// </summary>
        public static string? SharedCredentialsFile { get; set; } = __config.Get("sharedCredentialsFile") ?? Utilities.GetEnv("ALICLOUD_SHARED_CREDENTIALS_FILE");

        /// <summary>
        /// Skip static validation of region ID. Used by users of alternative AlibabaCloud-like APIs or users w/ access to regions
        /// that are not public (yet).
        /// </summary>
        public static bool? SkipRegionValidation { get; set; } = __config.GetBoolean("skipRegionValidation");

        public static class Types
        {

             public class AssumeRole
             {
                public string? Policy { get; set; } = null!;
                public string RoleArn { get; set; }
                public int? SessionExpiration { get; set; }
                public string? SessionName { get; set; } = null!;
            }

             public class Endpoints
             {
                public string? Actiontrail { get; set; } = null!;
                public string? Adb { get; set; } = null!;
                public string? Alidns { get; set; } = null!;
                public string? Alikafka { get; set; } = null!;
                public string? Apigateway { get; set; } = null!;
                public string? Bssopenapi { get; set; } = null!;
                public string? Cas { get; set; } = null!;
                public string? Cassandra { get; set; } = null!;
                public string? Cbn { get; set; } = null!;
                public string? Cdn { get; set; } = null!;
                public string? Cen { get; set; } = null!;
                public string? Cms { get; set; } = null!;
                public string? Cr { get; set; } = null!;
                public string? Cs { get; set; } = null!;
                public string? Datahub { get; set; } = null!;
                public string? Ddosbgp { get; set; } = null!;
                public string? Ddoscoo { get; set; } = null!;
                public string? Dds { get; set; } = null!;
                public string? DmsEnterprise { get; set; } = null!;
                public string? Dns { get; set; } = null!;
                public string? Drds { get; set; } = null!;
                public string? Eci { get; set; } = null!;
                public string? Ecs { get; set; } = null!;
                public string? Elasticsearch { get; set; } = null!;
                public string? Emr { get; set; } = null!;
                public string? Ess { get; set; } = null!;
                public string? Fc { get; set; } = null!;
                public string? Gpdb { get; set; } = null!;
                public string? Kms { get; set; } = null!;
                public string? Kvstore { get; set; } = null!;
                public string? Location { get; set; } = null!;
                public string? Log { get; set; } = null!;
                public string? Market { get; set; } = null!;
                public string? Maxcompute { get; set; } = null!;
                public string? Mns { get; set; } = null!;
                public string? Nas { get; set; } = null!;
                public string? Ons { get; set; } = null!;
                public string? Oss { get; set; } = null!;
                public string? Ots { get; set; } = null!;
                public string? Polardb { get; set; } = null!;
                public string? Pvtz { get; set; } = null!;
                public string? Ram { get; set; } = null!;
                public string? Rds { get; set; } = null!;
                public string? Resourcemanager { get; set; } = null!;
                public string? Slb { get; set; } = null!;
                public string? Sts { get; set; } = null!;
                public string? Vpc { get; set; } = null!;
                public string? WafOpenapi { get; set; } = null!;
            }
        }
    }
}
