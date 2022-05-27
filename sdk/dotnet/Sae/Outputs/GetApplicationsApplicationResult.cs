// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae.Outputs
{

    [OutputType]
    public sealed class GetApplicationsApplicationResult
    {
        /// <summary>
        /// The ARN of the RAM role required when pulling images across accounts.
        /// </summary>
        public readonly string AcrAssumeRoleArn;
        /// <summary>
        /// Application description information. No more than 1024 characters.
        /// </summary>
        public readonly string AppDescription;
        /// <summary>
        /// Application Name. Combinations of numbers, letters, and dashes (-) are allowed. It must start with a letter and the maximum length is 36 characters.
        /// </summary>
        public readonly string AppName;
        /// <summary>
        /// The first ID of the resource.
        /// </summary>
        public readonly string ApplicationId;
        /// <summary>
        /// Mirror start command. The command must be an executable object in the container. For example: sleep. Setting this command will cause the original startup command of the mirror to become invalid.
        /// </summary>
        public readonly string Command;
        /// <summary>
        /// Mirror startup command parameters. The parameters required for the above start command. For example: 1d.
        /// </summary>
        public readonly string CommandArgs;
        /// <summary>
        /// ConfigMap mount description.
        /// </summary>
        public readonly string ConfigMapMountDesc;
        /// <summary>
        /// The CPU required for each instance, in millicores, cannot be 0.
        /// </summary>
        public readonly int Cpu;
        /// <summary>
        /// Indicates That the Application of the Creation Time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Custom host mapping in the container. For example: [{"hostName":"samplehost","ip":"127.0.0.1"}].
        /// </summary>
        public readonly string CustomHostAlias;
        /// <summary>
        /// The operating environment used by the Pandora application.
        /// </summary>
        public readonly string EdasContainerVersion;
        /// <summary>
        /// The virtual switch where the elastic network card of the application instance is located. The switch must be located in the aforementioned VPC. The switch also has a binding relationship with the SAE namespace. If it is left blank, the default is the vSwitch ID bound to the namespace.
        /// </summary>
        public readonly string Envs;
        /// <summary>
        /// The ID of the Application.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Mirror address. Only Image type applications can configure the mirror address.
        /// </summary>
        public readonly string ImageUrl;
        /// <summary>
        /// The JAR package starts application parameters. Application default startup command: $JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs.
        /// </summary>
        public readonly string JarStartArgs;
        /// <summary>
        /// The JAR package starts the application option. Application default startup command: $JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs.
        /// </summary>
        public readonly string JarStartOptions;
        /// <summary>
        /// The JDK version that the deployment package depends on. Image type applications are not supported.
        /// </summary>
        public readonly string Jdk;
        /// <summary>
        /// Container health check. Containers that fail the health check will be shut down and restored. Currently, only the method of issuing commands in the container is supported.
        /// </summary>
        public readonly string Liveness;
        /// <summary>
        /// The memory required for each instance, in MB, cannot be 0. One-to-one correspondence with CPU.
        /// </summary>
        public readonly int Memory;
        /// <summary>
        /// The Minimum Available Instance. On the Change Had Promised during the Available Number of Instances to Be.
        /// </summary>
        public readonly int MinReadyInstances;
        /// <summary>
        /// Mount description information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetApplicationsApplicationMountDescResult> MountDescs;
        /// <summary>
        /// Mount point of NAS in application VPC.
        /// </summary>
        public readonly string MountHost;
        /// <summary>
        /// SAE namespace ID. Only namespaces whose names are lowercase letters and dashes (-) are supported, and must start with a letter. The namespace can be obtained by calling the DescribeNamespaceList interface.
        /// </summary>
        public readonly string NamespaceId;
        /// <summary>
        /// ID of the mounted NAS, Must be in the same region as the cluster. It must have an available mount point creation quota, or its mount point must be on a switch in the VPC. If it is not filled in and the mountDescs field is present, a NAS will be automatically purchased and mounted on the switch in the VPC by default.
        /// </summary>
        public readonly string NasId;
        /// <summary>
        /// OSS AccessKey ID.
        /// </summary>
        public readonly string OssAkId;
        /// <summary>
        /// OSS  AccessKey Secret.
        /// </summary>
        public readonly string OssAkSecret;
        /// <summary>
        /// OSS mount description information.
        /// </summary>
        public readonly string OssMountDescs;
        /// <summary>
        /// The OSS mount detail.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetApplicationsApplicationOssMountDetailResult> OssMountDetails;
        /// <summary>
        /// Application package type. Support FatJar, War and Image.
        /// </summary>
        public readonly string PackageType;
        /// <summary>
        /// Deployment package address. Only FatJar or War type applications can configure the deployment package address.
        /// </summary>
        public readonly string PackageUrl;
        /// <summary>
        /// The version number of the deployment package. Required when the Package Type is War and FatJar.
        /// </summary>
        public readonly string PackageVersion;
        /// <summary>
        /// The PHP application monitors the mount path, and you need to ensure that the PHP server will load the configuration file of this path. You don't need to pay attention to the configuration content, SAE will automatically render the correct configuration file.
        /// </summary>
        public readonly string PhpArmsConfigLocation;
        /// <summary>
        /// PHP configuration file content.
        /// </summary>
        public readonly string PhpConfig;
        /// <summary>
        /// PHP application startup configuration mount path, you need to ensure that the PHP server will start using this configuration file.
        /// </summary>
        public readonly string PhpConfigLocation;
        /// <summary>
        /// Execute the script after startup, the format is like: {"exec":{"command":["cat","/etc/group"]}}.
        /// </summary>
        public readonly string PostStart;
        /// <summary>
        /// Execute the script before stopping, the format is like: {"exec":{"command":["cat","/etc/group"]}}.
        /// </summary>
        public readonly string PreStop;
        /// <summary>
        /// Application startup status checks, containers that fail multiple health checks will be shut down and restarted. Containers that do not pass the health check will not receive SLB traffic. For example: {"exec":{"command":["sh","-c","cat /home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds ":2}.
        /// </summary>
        public readonly string Readiness;
        public readonly string RegionId;
        /// <summary>
        /// Initial number of instances.
        /// </summary>
        public readonly int Replicas;
        public readonly string RepoName;
        public readonly string RepoNamespace;
        public readonly string RepoOriginType;
        /// <summary>
        /// Security group ID.
        /// </summary>
        public readonly string SecurityGroupId;
        /// <summary>
        /// SLS  configuration.
        /// </summary>
        public readonly string SlsConfigs;
        /// <summary>
        /// The status of the resource.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// Graceful offline timeout, the default is 30, the unit is seconds. The value range is 1~60.
        /// </summary>
        public readonly int TerminationGracePeriodSeconds;
        /// <summary>
        /// Time zone, the default value is Asia/Shanghai.
        /// </summary>
        public readonly string Timezone;
        /// <summary>
        /// Tomcat file configuration, set to "" or "{}" means to delete the configuration:  useDefaultConfig: Whether to use a custom configuration, if it is true, it means that the custom configuration is not used; if it is false, it means that the custom configuration is used. If you do not use custom configuration, the following parameter configuration will not take effect.  contextInputType: Select the access path of the application.  war: No need to fill in the custom path, the access path of the application is the WAR package name. root: No need to fill in the custom path, the access path of the application is /. custom: You need to fill in the custom path in the custom path below. contextPath: custom path, this parameter only needs to be configured when the contextInputType type is custom.  httpPort: The port range is 1024~65535. Ports less than 1024 need Root permission to operate. Because the container is configured with Admin permissions, please fill in a port greater than 1024. If not configured, the default is 8080. maxThreads: Configure the number of connections in the connection pool, the default size is 400. uriEncoding: Tomcat encoding format, including UTF-8, ISO-8859-1, GBK and GB2312. If not set, the default is ISO-8859-1. useBodyEncoding: Whether to use BodyEncoding for URL.
        /// </summary>
        public readonly string TomcatConfig;
        /// <summary>
        /// The VPC corresponding to the SAE namespace. In SAE, a namespace can only correspond to one VPC and cannot be modified. Creating a SAE application in the namespace for the first time will form a binding relationship. Multiple namespaces can correspond to a VPC. If you leave it blank, it will default to the VPC ID bound to the namespace.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// The vswitch id.
        /// </summary>
        public readonly string VswitchId;
        /// <summary>
        /// WAR package launch application option. Application default startup command: java $JAVA_OPTS $CATALINA_OPTS [-Options] org.apache.catalina.startup.Bootstrap "$@" start.
        /// </summary>
        public readonly string WarStartOptions;
        /// <summary>
        /// The version of tomcat that the deployment package depends on. Image type applications are not supported.
        /// </summary>
        public readonly string WebContainer;

        [OutputConstructor]
        private GetApplicationsApplicationResult(
            string acrAssumeRoleArn,

            string appDescription,

            string appName,

            string applicationId,

            string command,

            string commandArgs,

            string configMapMountDesc,

            int cpu,

            string createTime,

            string customHostAlias,

            string edasContainerVersion,

            string envs,

            string id,

            string imageUrl,

            string jarStartArgs,

            string jarStartOptions,

            string jdk,

            string liveness,

            int memory,

            int minReadyInstances,

            ImmutableArray<Outputs.GetApplicationsApplicationMountDescResult> mountDescs,

            string mountHost,

            string namespaceId,

            string nasId,

            string ossAkId,

            string ossAkSecret,

            string ossMountDescs,

            ImmutableArray<Outputs.GetApplicationsApplicationOssMountDetailResult> ossMountDetails,

            string packageType,

            string packageUrl,

            string packageVersion,

            string phpArmsConfigLocation,

            string phpConfig,

            string phpConfigLocation,

            string postStart,

            string preStop,

            string readiness,

            string regionId,

            int replicas,

            string repoName,

            string repoNamespace,

            string repoOriginType,

            string securityGroupId,

            string slsConfigs,

            string status,

            ImmutableDictionary<string, object> tags,

            int terminationGracePeriodSeconds,

            string timezone,

            string tomcatConfig,

            string vpcId,

            string vswitchId,

            string warStartOptions,

            string webContainer)
        {
            AcrAssumeRoleArn = acrAssumeRoleArn;
            AppDescription = appDescription;
            AppName = appName;
            ApplicationId = applicationId;
            Command = command;
            CommandArgs = commandArgs;
            ConfigMapMountDesc = configMapMountDesc;
            Cpu = cpu;
            CreateTime = createTime;
            CustomHostAlias = customHostAlias;
            EdasContainerVersion = edasContainerVersion;
            Envs = envs;
            Id = id;
            ImageUrl = imageUrl;
            JarStartArgs = jarStartArgs;
            JarStartOptions = jarStartOptions;
            Jdk = jdk;
            Liveness = liveness;
            Memory = memory;
            MinReadyInstances = minReadyInstances;
            MountDescs = mountDescs;
            MountHost = mountHost;
            NamespaceId = namespaceId;
            NasId = nasId;
            OssAkId = ossAkId;
            OssAkSecret = ossAkSecret;
            OssMountDescs = ossMountDescs;
            OssMountDetails = ossMountDetails;
            PackageType = packageType;
            PackageUrl = packageUrl;
            PackageVersion = packageVersion;
            PhpArmsConfigLocation = phpArmsConfigLocation;
            PhpConfig = phpConfig;
            PhpConfigLocation = phpConfigLocation;
            PostStart = postStart;
            PreStop = preStop;
            Readiness = readiness;
            RegionId = regionId;
            Replicas = replicas;
            RepoName = repoName;
            RepoNamespace = repoNamespace;
            RepoOriginType = repoOriginType;
            SecurityGroupId = securityGroupId;
            SlsConfigs = slsConfigs;
            Status = status;
            Tags = tags;
            TerminationGracePeriodSeconds = terminationGracePeriodSeconds;
            Timezone = timezone;
            TomcatConfig = tomcatConfig;
            VpcId = vpcId;
            VswitchId = vswitchId;
            WarStartOptions = warStartOptions;
            WebContainer = webContainer;
        }
    }
}
