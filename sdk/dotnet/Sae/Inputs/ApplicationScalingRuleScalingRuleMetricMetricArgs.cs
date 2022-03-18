// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae.Inputs
{

    public sealed class ApplicationScalingRuleScalingRuleMetricMetricArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// According to different `metric_type`, set the target value of the corresponding monitoring index.
        /// </summary>
        [Input("metricTargetAverageUtilization")]
        public Input<int>? MetricTargetAverageUtilization { get; set; }

        /// <summary>
        /// Monitoring indicator trigger condition. Valid values: `CPU`, `MEMORY`, `tcpActiveConn`, `SLB_QPS` and `SLB_RT`. The values are described as follows:
        /// - CPU: CPU usage.
        /// - MEMORY: MEMORY usage.
        /// - tcpActiveConn: the average number of TCP active connections for a single instance in 30 seconds.
        /// - SLB_QPS: the average public network SLB QPS of a single instance within 15 seconds.
        /// - SLB_RT: the average response time of public network SLB within 15 seconds.
        /// </summary>
        [Input("metricType")]
        public Input<string>? MetricType { get; set; }

        public ApplicationScalingRuleScalingRuleMetricMetricArgs()
        {
        }
    }
}
