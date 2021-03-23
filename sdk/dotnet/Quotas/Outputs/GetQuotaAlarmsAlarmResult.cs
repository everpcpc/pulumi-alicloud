// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Quotas.Outputs
{

    [OutputType]
    public sealed class GetQuotaAlarmsAlarmResult
    {
        /// <summary>
        /// The first ID of the resource.
        /// </summary>
        public readonly string AlarmId;
        /// <summary>
        /// The ID of the Quota Alarm.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Product Code.
        /// </summary>
        public readonly string ProductCode;
        /// <summary>
        /// The Quota Action Code.
        /// </summary>
        public readonly string QuotaActionCode;
        /// <summary>
        /// The name of Quota Alarm.
        /// </summary>
        public readonly string QuotaAlarmName;
        /// <summary>
        /// The Quota Dimensions.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetQuotaAlarmsAlarmQuotaDimensionResult> QuotaDimensions;
        /// <summary>
        /// The threshold of Quota Alarm.
        /// </summary>
        public readonly double Threshold;
        /// <summary>
        /// The threshold percent of Quota Alarm.
        /// </summary>
        public readonly double ThresholdPercent;
        /// <summary>
        /// The WebHook of Quota Alarm.
        /// </summary>
        public readonly string WebHook;

        [OutputConstructor]
        private GetQuotaAlarmsAlarmResult(
            string alarmId,

            string id,

            string productCode,

            string quotaActionCode,

            string quotaAlarmName,

            ImmutableArray<Outputs.GetQuotaAlarmsAlarmQuotaDimensionResult> quotaDimensions,

            double threshold,

            double thresholdPercent,

            string webHook)
        {
            AlarmId = alarmId;
            Id = id;
            ProductCode = productCode;
            QuotaActionCode = quotaActionCode;
            QuotaAlarmName = quotaAlarmName;
            QuotaDimensions = quotaDimensions;
            Threshold = threshold;
            ThresholdPercent = thresholdPercent;
            WebHook = webHook;
        }
    }
}