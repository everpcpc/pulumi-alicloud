// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Arms.Outputs
{

    [OutputType]
    public sealed class GetIntegrationExportersIntegrationExporterResult
    {
        /// <summary>
        /// The ID of the Prometheus instance.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// Integration Exporter Type.
        /// </summary>
        public readonly string ExporterType;
        /// <summary>
        /// The ID of the Integration Exporter. It formats as `&lt;cluster_id&gt;:&lt;integration_type&gt;:&lt;instance_id&gt;`.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of the Integration Exporter instance.
        /// </summary>
        public readonly int InstanceId;
        /// <summary>
        /// The name of the instance.
        /// </summary>
        public readonly string InstanceName;
        /// <summary>
        /// The type of prometheus integration.
        /// </summary>
        public readonly string IntegrationType;
        /// <summary>
        /// Exporter configuration parameter json string.
        /// </summary>
        public readonly string Param;
        /// <summary>
        /// Monitor the target address.
        /// </summary>
        public readonly string Target;
        /// <summary>
        /// The version information.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetIntegrationExportersIntegrationExporterResult(
            string clusterId,

            string exporterType,

            string id,

            int instanceId,

            string instanceName,

            string integrationType,

            string param,

            string target,

            string version)
        {
            ClusterId = clusterId;
            ExporterType = exporterType;
            Id = id;
            InstanceId = instanceId;
            InstanceName = instanceName;
            IntegrationType = integrationType;
            Param = param;
            Target = target;
            Version = version;
        }
    }
}
