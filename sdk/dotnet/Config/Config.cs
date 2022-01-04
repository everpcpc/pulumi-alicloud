// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.AliCloud
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly Pulumi.Config __config = new Pulumi.Config("alicloud");

        private static readonly __Value<string?> _accessKey = new __Value<string?>(() => __config.Get("accessKey"));
        /// <summary>
        /// The access key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
        /// console.
        /// </summary>
        public static string? AccessKey
        {
            get => _accessKey.Get();
            set => _accessKey.Set(value);
        }

        private static readonly __Value<string?> _accountId = new __Value<string?>(() => __config.Get("accountId"));
        /// <summary>
        /// The account ID for some service API operations. You can retrieve this from the 'Security Settings' section of the
        /// Alibaba Cloud console.
        /// </summary>
        public static string? AccountId
        {
            get => _accountId.Get();
            set => _accountId.Set(value);
        }

        private static readonly __Value<Pulumi.AliCloud.Config.Types.AssumeRole?> _assumeRole = new __Value<Pulumi.AliCloud.Config.Types.AssumeRole?>(() => __config.GetObject<Pulumi.AliCloud.Config.Types.AssumeRole>("assumeRole"));
        public static Pulumi.AliCloud.Config.Types.AssumeRole? AssumeRole
        {
            get => _assumeRole.Get();
            set => _assumeRole.Set(value);
        }

        private static readonly __Value<int?> _clientConnectTimeout = new __Value<int?>(() => __config.GetInt32("clientConnectTimeout"));
        /// <summary>
        /// The maximum timeout of the client connection server.
        /// </summary>
        public static int? ClientConnectTimeout
        {
            get => _clientConnectTimeout.Get();
            set => _clientConnectTimeout.Set(value);
        }

        private static readonly __Value<int?> _clientReadTimeout = new __Value<int?>(() => __config.GetInt32("clientReadTimeout"));
        /// <summary>
        /// The maximum timeout of the client read request.
        /// </summary>
        public static int? ClientReadTimeout
        {
            get => _clientReadTimeout.Get();
            set => _clientReadTimeout.Set(value);
        }

        private static readonly __Value<string?> _configurationSource = new __Value<string?>(() => __config.Get("configurationSource"));
        /// <summary>
        /// Use this to mark a terraform configuration file source.
        /// </summary>
        public static string? ConfigurationSource
        {
            get => _configurationSource.Get();
            set => _configurationSource.Set(value);
        }

        private static readonly __Value<string?> _credentialsUri = new __Value<string?>(() => __config.Get("credentialsUri"));
        /// <summary>
        /// The URI of sidecar credentials service.
        /// </summary>
        public static string? CredentialsUri
        {
            get => _credentialsUri.Get();
            set => _credentialsUri.Set(value);
        }

        private static readonly __Value<string?> _ecsRoleName = new __Value<string?>(() => __config.Get("ecsRoleName") ?? Utilities.GetEnv("ALICLOUD_ECS_ROLE_NAME"));
        /// <summary>
        /// The RAM Role Name attached on a ECS instance for API operations. You can retrieve this from the 'Access Control' section
        /// of the Alibaba Cloud console.
        /// </summary>
        public static string? EcsRoleName
        {
            get => _ecsRoleName.Get();
            set => _ecsRoleName.Set(value);
        }

        private static readonly __Value<ImmutableArray<Pulumi.AliCloud.Config.Types.Endpoints>> _endpoints = new __Value<ImmutableArray<Pulumi.AliCloud.Config.Types.Endpoints>>(() => __config.GetObject<ImmutableArray<Pulumi.AliCloud.Config.Types.Endpoints>>("endpoints"));
        public static ImmutableArray<Pulumi.AliCloud.Config.Types.Endpoints> Endpoints
        {
            get => _endpoints.Get();
            set => _endpoints.Set(value);
        }

        private static readonly __Value<string?> _fc = new __Value<string?>(() => __config.Get("fc"));
        public static string? Fc
        {
            get => _fc.Get();
            set => _fc.Set(value);
        }

        private static readonly __Value<string?> _logEndpoint = new __Value<string?>(() => __config.Get("logEndpoint"));
        public static string? LogEndpoint
        {
            get => _logEndpoint.Get();
            set => _logEndpoint.Set(value);
        }

        private static readonly __Value<string?> _mnsEndpoint = new __Value<string?>(() => __config.Get("mnsEndpoint"));
        public static string? MnsEndpoint
        {
            get => _mnsEndpoint.Get();
            set => _mnsEndpoint.Set(value);
        }

        private static readonly __Value<string?> _otsInstanceName = new __Value<string?>(() => __config.Get("otsInstanceName"));
        public static string? OtsInstanceName
        {
            get => _otsInstanceName.Get();
            set => _otsInstanceName.Set(value);
        }

        private static readonly __Value<string?> _profile = new __Value<string?>(() => __config.Get("profile") ?? Utilities.GetEnv("ALICLOUD_PROFILE"));
        /// <summary>
        /// The profile for API operations. If not set, the default profile created with `aliyun configure` will be used.
        /// </summary>
        public static string? Profile
        {
            get => _profile.Get();
            set => _profile.Set(value);
        }

        private static readonly __Value<string?> _protocol = new __Value<string?>(() => __config.Get("protocol"));
        public static string? Protocol
        {
            get => _protocol.Get();
            set => _protocol.Set(value);
        }

        private static readonly __Value<string?> _region = new __Value<string?>(() => __config.Get("region") ?? Utilities.GetEnv("ALICLOUD_REGION"));
        /// <summary>
        /// The region where Alibaba Cloud operations will take place. Examples are cn-beijing, cn-hangzhou, eu-central-1, etc.
        /// </summary>
        public static string? Region
        {
            get => _region.Get();
            set => _region.Set(value);
        }

        private static readonly __Value<string?> _secretKey = new __Value<string?>(() => __config.Get("secretKey"));
        /// <summary>
        /// The secret key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
        /// console.
        /// </summary>
        public static string? SecretKey
        {
            get => _secretKey.Get();
            set => _secretKey.Set(value);
        }

        private static readonly __Value<string?> _secureTransport = new __Value<string?>(() => __config.Get("secureTransport"));
        /// <summary>
        /// The security transport for the assume role invoking.
        /// </summary>
        public static string? SecureTransport
        {
            get => _secureTransport.Get();
            set => _secureTransport.Set(value);
        }

        private static readonly __Value<string?> _securityToken = new __Value<string?>(() => __config.Get("securityToken"));
        /// <summary>
        /// security token. A security token is only required if you are using Security Token Service.
        /// </summary>
        public static string? SecurityToken
        {
            get => _securityToken.Get();
            set => _securityToken.Set(value);
        }

        private static readonly __Value<string?> _securityTransport = new __Value<string?>(() => __config.Get("securityTransport"));
        public static string? SecurityTransport
        {
            get => _securityTransport.Get();
            set => _securityTransport.Set(value);
        }

        private static readonly __Value<string?> _sharedCredentialsFile = new __Value<string?>(() => __config.Get("sharedCredentialsFile"));
        /// <summary>
        /// The path to the shared credentials file. If not set this defaults to ~/.aliyun/config.json
        /// </summary>
        public static string? SharedCredentialsFile
        {
            get => _sharedCredentialsFile.Get();
            set => _sharedCredentialsFile.Set(value);
        }

        private static readonly __Value<bool?> _skipRegionValidation = new __Value<bool?>(() => __config.GetBoolean("skipRegionValidation"));
        /// <summary>
        /// Skip static validation of region ID. Used by users of alternative AlibabaCloud-like APIs or users w/ access to regions
        /// that are not public (yet).
        /// </summary>
        public static bool? SkipRegionValidation
        {
            get => _skipRegionValidation.Get();
            set => _skipRegionValidation.Set(value);
        }

        private static readonly __Value<string?> _sourceIp = new __Value<string?>(() => __config.Get("sourceIp"));
        /// <summary>
        /// The source ip for the assume role invoking.
        /// </summary>
        public static string? SourceIp
        {
            get => _sourceIp.Get();
            set => _sourceIp.Set(value);
        }

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
                public string? Acr { get; set; } = null!;
                public string? Actiontrail { get; set; } = null!;
                public string? Adb { get; set; } = null!;
                public string? Alb { get; set; } = null!;
                public string? Alidfs { get; set; } = null!;
                public string? Alidns { get; set; } = null!;
                public string? Alikafka { get; set; } = null!;
                public string? Apigateway { get; set; } = null!;
                public string? Arms { get; set; } = null!;
                public string? Bastionhost { get; set; } = null!;
                public string? BrainIndustrial { get; set; } = null!;
                public string? Bssopenapi { get; set; } = null!;
                public string? Cas { get; set; } = null!;
                public string? Cassandra { get; set; } = null!;
                public string? Cbn { get; set; } = null!;
                public string? Cddc { get; set; } = null!;
                public string? Cdn { get; set; } = null!;
                public string? Cds { get; set; } = null!;
                public string? Clickhouse { get; set; } = null!;
                public string? Cloudauth { get; set; } = null!;
                public string? Cloudphone { get; set; } = null!;
                public string? Cloudsso { get; set; } = null!;
                public string? Cms { get; set; } = null!;
                public string? Config { get; set; } = null!;
                public string? Cr { get; set; } = null!;
                public string? Cs { get; set; } = null!;
                public string? Datahub { get; set; } = null!;
                public string? Dataworkspublic { get; set; } = null!;
                public string? Dbfs { get; set; } = null!;
                public string? Dcdn { get; set; } = null!;
                public string? Ddosbgp { get; set; } = null!;
                public string? Ddoscoo { get; set; } = null!;
                public string? Dds { get; set; } = null!;
                public string? Devopsrdc { get; set; } = null!;
                public string? Dg { get; set; } = null!;
                public string? Dm { get; set; } = null!;
                public string? DmsEnterprise { get; set; } = null!;
                public string? Dns { get; set; } = null!;
                public string? Drds { get; set; } = null!;
                public string? Dts { get; set; } = null!;
                public string? Eais { get; set; } = null!;
                public string? Eci { get; set; } = null!;
                public string? Ecs { get; set; } = null!;
                public string? Edsuser { get; set; } = null!;
                public string? Ehpc { get; set; } = null!;
                public string? Eipanycast { get; set; } = null!;
                public string? Elasticsearch { get; set; } = null!;
                public string? Emr { get; set; } = null!;
                public string? Ens { get; set; } = null!;
                public string? Ess { get; set; } = null!;
                public string? Eventbridge { get; set; } = null!;
                public string? Fc { get; set; } = null!;
                public string? Fnf { get; set; } = null!;
                public string? Ga { get; set; } = null!;
                public string? Gaplus { get; set; } = null!;
                public string? Gds { get; set; } = null!;
                public string? Gpdb { get; set; } = null!;
                public string? Gwsecd { get; set; } = null!;
                public string? Hbr { get; set; } = null!;
                public string? HcsSgw { get; set; } = null!;
                public string? Hitsdb { get; set; } = null!;
                public string? Imm { get; set; } = null!;
                public string? Imp { get; set; } = null!;
                public string? Ims { get; set; } = null!;
                public string? Iot { get; set; } = null!;
                public string? Kms { get; set; } = null!;
                public string? Kvstore { get; set; } = null!;
                public string? Location { get; set; } = null!;
                public string? Log { get; set; } = null!;
                public string? Market { get; set; } = null!;
                public string? Maxcompute { get; set; } = null!;
                public string? Mhub { get; set; } = null!;
                public string? Mns { get; set; } = null!;
                public string? Mscopensubscription { get; set; } = null!;
                public string? Mse { get; set; } = null!;
                public string? Nas { get; set; } = null!;
                public string? Ons { get; set; } = null!;
                public string? Onsproxy { get; set; } = null!;
                public string? Oos { get; set; } = null!;
                public string? Opensearch { get; set; } = null!;
                public string? Oss { get; set; } = null!;
                public string? Ots { get; set; } = null!;
                public string? Polardb { get; set; } = null!;
                public string? Privatelink { get; set; } = null!;
                public string? Pvtz { get; set; } = null!;
                public string? Quickbi { get; set; } = null!;
                public string? Quotas { get; set; } = null!;
                public string? RKvstore { get; set; } = null!;
                public string? Ram { get; set; } = null!;
                public string? Rds { get; set; } = null!;
                public string? Redisa { get; set; } = null!;
                public string? Resourcemanager { get; set; } = null!;
                public string? Resourcesharing { get; set; } = null!;
                public string? Ros { get; set; } = null!;
                public string? Sas { get; set; } = null!;
                public string? Scdn { get; set; } = null!;
                public string? Sddp { get; set; } = null!;
                public string? Serverless { get; set; } = null!;
                public string? Servicemesh { get; set; } = null!;
                public string? Sgw { get; set; } = null!;
                public string? Slb { get; set; } = null!;
                public string? Sts { get; set; } = null!;
                public string? Swas { get; set; } = null!;
                public string? Vod { get; set; } = null!;
                public string? Vpc { get; set; } = null!;
                public string? Vs { get; set; } = null!;
                public string? Waf { get; set; } = null!;
                public string? WafOpenapi { get; set; } = null!;
            }
        }
    }
}
