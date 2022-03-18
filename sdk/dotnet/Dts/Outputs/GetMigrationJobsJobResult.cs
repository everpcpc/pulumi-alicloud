// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dts.Outputs
{

    [OutputType]
    public sealed class GetMigrationJobsJobResult
    {
        /// <summary>
        /// Whether or not to execute DTS supports schema migration, full data migration, or full-data initialization.
        /// </summary>
        public readonly bool DataInitialization;
        /// <summary>
        /// Whether to perform incremental data migration for migration types or synchronization values include:
        /// </summary>
        public readonly bool DataSynchronization;
        /// <summary>
        /// The Migration object, in the format of JSON strings.
        /// </summary>
        public readonly string DbList;
        /// <summary>
        /// The name of migrate the database.
        /// </summary>
        public readonly string DestinationEndpointDataBaseName;
        /// <summary>
        /// The type of destination database.
        /// </summary>
        public readonly string DestinationEndpointEngineName;
        /// <summary>
        /// The ID of destination instance.
        /// </summary>
        public readonly string DestinationEndpointInstanceId;
        /// <summary>
        /// The type of destination instance.
        /// </summary>
        public readonly string DestinationEndpointInstanceType;
        /// <summary>
        /// The ip of source endpoint.
        /// </summary>
        public readonly string DestinationEndpointIp;
        /// <summary>
        /// The SID of Oracle database.
        /// </summary>
        public readonly string DestinationEndpointOracleSid;
        /// <summary>
        /// The port of source endpoint.
        /// </summary>
        public readonly string DestinationEndpointPort;
        /// <summary>
        /// The region of destination instance.
        /// </summary>
        public readonly string DestinationEndpointRegion;
        /// <summary>
        /// The username of database account.
        /// </summary>
        public readonly string DestinationEndpointUserName;
        /// <summary>
        /// The Migration instance ID. The ID of `alicloud.dts.MigrationInstance`.
        /// </summary>
        public readonly string DtsInstanceId;
        /// <summary>
        /// The ID of the Migration Job.
        /// </summary>
        public readonly string DtsJobId;
        /// <summary>
        /// The name of synchronization job.
        /// </summary>
        public readonly string DtsJobName;
        /// <summary>
        /// The ID of the Migration Job. Its value is same as `dts_job_id`.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The payment type of the Migration Instance.
        /// </summary>
        public readonly string PaymentType;
        /// <summary>
        /// The name of migrate the database.
        /// </summary>
        public readonly string SourceEndpointDatabaseName;
        /// <summary>
        /// The type of source database.
        /// </summary>
        public readonly string SourceEndpointEngineName;
        /// <summary>
        /// The ID of source instance.
        /// </summary>
        public readonly string SourceEndpointInstanceId;
        /// <summary>
        /// The type of source instance.
        /// </summary>
        public readonly string SourceEndpointInstanceType;
        /// <summary>
        /// The ip of source endpoint.
        /// </summary>
        public readonly string SourceEndpointIp;
        /// <summary>
        /// The SID of Oracle database.
        /// </summary>
        public readonly string SourceEndpointOracleSid;
        /// <summary>
        /// The Alibaba Cloud account ID to which the source instance belongs.
        /// </summary>
        public readonly string SourceEndpointOwnerId;
        /// <summary>
        /// The port of source endpoint.
        /// </summary>
        public readonly string SourceEndpointPort;
        /// <summary>
        /// The region of source instance.
        /// </summary>
        public readonly string SourceEndpointRegion;
        /// <summary>
        /// The name of the role configured for the cloud account to which the source instance belongs.
        /// </summary>
        public readonly string SourceEndpointRole;
        /// <summary>
        /// The username of database account.
        /// </summary>
        public readonly string SourceEndpointUserName;
        /// <summary>
        /// The status of the resource.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Whether to perform a database table structure to migrate or initialization.
        /// </summary>
        public readonly bool StructureInitialization;

        [OutputConstructor]
        private GetMigrationJobsJobResult(
            bool dataInitialization,

            bool dataSynchronization,

            string dbList,

            string destinationEndpointDataBaseName,

            string destinationEndpointEngineName,

            string destinationEndpointInstanceId,

            string destinationEndpointInstanceType,

            string destinationEndpointIp,

            string destinationEndpointOracleSid,

            string destinationEndpointPort,

            string destinationEndpointRegion,

            string destinationEndpointUserName,

            string dtsInstanceId,

            string dtsJobId,

            string dtsJobName,

            string id,

            string paymentType,

            string sourceEndpointDatabaseName,

            string sourceEndpointEngineName,

            string sourceEndpointInstanceId,

            string sourceEndpointInstanceType,

            string sourceEndpointIp,

            string sourceEndpointOracleSid,

            string sourceEndpointOwnerId,

            string sourceEndpointPort,

            string sourceEndpointRegion,

            string sourceEndpointRole,

            string sourceEndpointUserName,

            string status,

            bool structureInitialization)
        {
            DataInitialization = dataInitialization;
            DataSynchronization = dataSynchronization;
            DbList = dbList;
            DestinationEndpointDataBaseName = destinationEndpointDataBaseName;
            DestinationEndpointEngineName = destinationEndpointEngineName;
            DestinationEndpointInstanceId = destinationEndpointInstanceId;
            DestinationEndpointInstanceType = destinationEndpointInstanceType;
            DestinationEndpointIp = destinationEndpointIp;
            DestinationEndpointOracleSid = destinationEndpointOracleSid;
            DestinationEndpointPort = destinationEndpointPort;
            DestinationEndpointRegion = destinationEndpointRegion;
            DestinationEndpointUserName = destinationEndpointUserName;
            DtsInstanceId = dtsInstanceId;
            DtsJobId = dtsJobId;
            DtsJobName = dtsJobName;
            Id = id;
            PaymentType = paymentType;
            SourceEndpointDatabaseName = sourceEndpointDatabaseName;
            SourceEndpointEngineName = sourceEndpointEngineName;
            SourceEndpointInstanceId = sourceEndpointInstanceId;
            SourceEndpointInstanceType = sourceEndpointInstanceType;
            SourceEndpointIp = sourceEndpointIp;
            SourceEndpointOracleSid = sourceEndpointOracleSid;
            SourceEndpointOwnerId = sourceEndpointOwnerId;
            SourceEndpointPort = sourceEndpointPort;
            SourceEndpointRegion = sourceEndpointRegion;
            SourceEndpointRole = sourceEndpointRole;
            SourceEndpointUserName = sourceEndpointUserName;
            Status = status;
            StructureInitialization = structureInitialization;
        }
    }
}
