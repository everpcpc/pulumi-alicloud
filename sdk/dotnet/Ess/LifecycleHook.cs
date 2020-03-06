// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ess
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/ess_lifecycle_hook.html.markdown.
    /// </summary>
    public partial class LifecycleHook : Pulumi.CustomResource
    {
        /// <summary>
        /// Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses. Applicable value: CONTINUE, ABANDON, default value: CONTINUE.
        /// </summary>
        [Output("defaultResult")]
        public Output<string?> DefaultResult { get; private set; } = null!;

        /// <summary>
        /// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the default_result parameter. Default value: 600.
        /// </summary>
        [Output("heartbeatTimeout")]
        public Output<int?> HeartbeatTimeout { get; private set; } = null!;

        /// <summary>
        /// Type of Scaling activity attached to lifecycle hook. Supported value: SCALE_OUT, SCALE_IN.
        /// </summary>
        [Output("lifecycleTransition")]
        public Output<string> LifecycleTransition { get; private set; } = null!;

        /// <summary>
        /// The name of the lifecycle hook, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is lifecycle hook id.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Arn of notification target.
        /// </summary>
        [Output("notificationArn")]
        public Output<string> NotificationArn { get; private set; } = null!;

        /// <summary>
        /// Additional information that you want to include when Auto Scaling sends a message to the notification target.
        /// </summary>
        [Output("notificationMetadata")]
        public Output<string> NotificationMetadata { get; private set; } = null!;

        /// <summary>
        /// The ID of the Auto Scaling group to which you want to assign the lifecycle hook.
        /// </summary>
        [Output("scalingGroupId")]
        public Output<string> ScalingGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a LifecycleHook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LifecycleHook(string name, LifecycleHookArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ess/lifecycleHook:LifecycleHook", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private LifecycleHook(string name, Input<string> id, LifecycleHookState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ess/lifecycleHook:LifecycleHook", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LifecycleHook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LifecycleHook Get(string name, Input<string> id, LifecycleHookState? state = null, CustomResourceOptions? options = null)
        {
            return new LifecycleHook(name, id, state, options);
        }
    }

    public sealed class LifecycleHookArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses. Applicable value: CONTINUE, ABANDON, default value: CONTINUE.
        /// </summary>
        [Input("defaultResult")]
        public Input<string>? DefaultResult { get; set; }

        /// <summary>
        /// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the default_result parameter. Default value: 600.
        /// </summary>
        [Input("heartbeatTimeout")]
        public Input<int>? HeartbeatTimeout { get; set; }

        /// <summary>
        /// Type of Scaling activity attached to lifecycle hook. Supported value: SCALE_OUT, SCALE_IN.
        /// </summary>
        [Input("lifecycleTransition", required: true)]
        public Input<string> LifecycleTransition { get; set; } = null!;

        /// <summary>
        /// The name of the lifecycle hook, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is lifecycle hook id.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Arn of notification target.
        /// </summary>
        [Input("notificationArn")]
        public Input<string>? NotificationArn { get; set; }

        /// <summary>
        /// Additional information that you want to include when Auto Scaling sends a message to the notification target.
        /// </summary>
        [Input("notificationMetadata")]
        public Input<string>? NotificationMetadata { get; set; }

        /// <summary>
        /// The ID of the Auto Scaling group to which you want to assign the lifecycle hook.
        /// </summary>
        [Input("scalingGroupId", required: true)]
        public Input<string> ScalingGroupId { get; set; } = null!;

        public LifecycleHookArgs()
        {
        }
    }

    public sealed class LifecycleHookState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses. Applicable value: CONTINUE, ABANDON, default value: CONTINUE.
        /// </summary>
        [Input("defaultResult")]
        public Input<string>? DefaultResult { get; set; }

        /// <summary>
        /// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the default_result parameter. Default value: 600.
        /// </summary>
        [Input("heartbeatTimeout")]
        public Input<int>? HeartbeatTimeout { get; set; }

        /// <summary>
        /// Type of Scaling activity attached to lifecycle hook. Supported value: SCALE_OUT, SCALE_IN.
        /// </summary>
        [Input("lifecycleTransition")]
        public Input<string>? LifecycleTransition { get; set; }

        /// <summary>
        /// The name of the lifecycle hook, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is lifecycle hook id.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Arn of notification target.
        /// </summary>
        [Input("notificationArn")]
        public Input<string>? NotificationArn { get; set; }

        /// <summary>
        /// Additional information that you want to include when Auto Scaling sends a message to the notification target.
        /// </summary>
        [Input("notificationMetadata")]
        public Input<string>? NotificationMetadata { get; set; }

        /// <summary>
        /// The ID of the Auto Scaling group to which you want to assign the lifecycle hook.
        /// </summary>
        [Input("scalingGroupId")]
        public Input<string>? ScalingGroupId { get; set; }

        public LifecycleHookState()
        {
        }
    }
}
