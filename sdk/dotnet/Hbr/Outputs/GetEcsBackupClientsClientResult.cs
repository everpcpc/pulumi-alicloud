// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Hbr.Outputs
{

    [OutputType]
    public sealed class GetEcsBackupClientsClientResult
    {
        /// <summary>
        /// The Client System Architecture (Only the ECS File Backup Client Is Available. Valid Values: `AMD64` , `386`.
        /// </summary>
        public readonly string ArchType;
        /// <summary>
        /// Client protected status.
        /// </summary>
        public readonly string BackupStatus;
        /// <summary>
        /// The Client Type. Valid Values: `ECS_CLIENT` (ECS File Backup Client).
        /// </summary>
        public readonly string ClientType;
        /// <summary>
        /// Client Version.
        /// </summary>
        public readonly string ClientVersion;
        /// <summary>
        /// The Client Creates a Time. Unix Time Seconds.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The Data Plane Data Access Point Type. Valid Values: `PUBLIC`, `VPC`, `CLASSIC`.
        /// </summary>
        public readonly string DataNetworkType;
        /// <summary>
        /// The Data Plane Proxy Settings. Valid Values: `DISABLE`, `USE_CONTROL_PROXY`, `CUSTOM`. **Note**: `USE_CONTROL_PROXY` (Default, the same with Control Plane), `CUSTOM` (Custom Configuration Items for the HTTP Protocol).
        /// </summary>
        public readonly string DataProxySetting;
        /// <summary>
        /// The first ID of the resource.
        /// </summary>
        public readonly string EcsBackupClientId;
        /// <summary>
        /// The ECS Host Name.
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// The ID of the Ecs Backup Client.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of ECS Instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// ECS Instance Names.
        /// </summary>
        public readonly string InstanceName;
        /// <summary>
        /// Client Last Heartbeat Time. Unix Time Seconds.
        /// </summary>
        public readonly string LastHeartBeatTime;
        /// <summary>
        /// The Latest Client Version.
        /// </summary>
        public readonly string MaxClientVersion;
        /// <summary>
        /// A Single Backup Task Uses for Example, Instances Can Be Grouped According to CPU Core Count, 0 Means No Restrictions.
        /// </summary>
        public readonly string MaxCpuCore;
        /// <summary>
        /// A Single Backup Task Parallel Work, the Number of 0 Means No Restrictions.
        /// </summary>
        public readonly string MaxWorker;
        /// <summary>
        /// The Client System Type (Only the ECS File Backup Client Is Available. Possible Values: * windows * linux.
        /// </summary>
        public readonly string OsType;
        /// <summary>
        /// Instance Must Not Use the Intranet IP Address.
        /// </summary>
        public readonly string PrivateIpv4;
        /// <summary>
        /// Custom Data Plane Proxy Server Host Address.
        /// </summary>
        public readonly string ProxyHost;
        /// <summary>
        /// Custom Data Plane Proxy Password.
        /// </summary>
        public readonly string ProxyPassword;
        /// <summary>
        /// Custom Data Plane Proxy Server Host Port.
        /// </summary>
        public readonly string ProxyPort;
        /// <summary>
        /// Custom Data Plane Proxy Server User Name.
        /// </summary>
        public readonly string ProxyUser;
        /// <summary>
        /// The status of the resource.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Client Update Time. Unix Time Seconds.
        /// </summary>
        public readonly string UpdatedTime;
        /// <summary>
        /// Indicates Whether to Use the Https Transport Data Plane Data.
        /// </summary>
        public readonly bool UseHttps;
        /// <summary>
        /// The Zone ID.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private GetEcsBackupClientsClientResult(
            string archType,

            string backupStatus,

            string clientType,

            string clientVersion,

            string createTime,

            string dataNetworkType,

            string dataProxySetting,

            string ecsBackupClientId,

            string hostname,

            string id,

            string instanceId,

            string instanceName,

            string lastHeartBeatTime,

            string maxClientVersion,

            string maxCpuCore,

            string maxWorker,

            string osType,

            string privateIpv4,

            string proxyHost,

            string proxyPassword,

            string proxyPort,

            string proxyUser,

            string status,

            string updatedTime,

            bool useHttps,

            string zoneId)
        {
            ArchType = archType;
            BackupStatus = backupStatus;
            ClientType = clientType;
            ClientVersion = clientVersion;
            CreateTime = createTime;
            DataNetworkType = dataNetworkType;
            DataProxySetting = dataProxySetting;
            EcsBackupClientId = ecsBackupClientId;
            Hostname = hostname;
            Id = id;
            InstanceId = instanceId;
            InstanceName = instanceName;
            LastHeartBeatTime = lastHeartBeatTime;
            MaxClientVersion = maxClientVersion;
            MaxCpuCore = maxCpuCore;
            MaxWorker = maxWorker;
            OsType = osType;
            PrivateIpv4 = privateIpv4;
            ProxyHost = proxyHost;
            ProxyPassword = proxyPassword;
            ProxyPort = proxyPort;
            ProxyUser = proxyUser;
            Status = status;
            UpdatedTime = updatedTime;
            UseHttps = useHttps;
            ZoneId = zoneId;
        }
    }
}