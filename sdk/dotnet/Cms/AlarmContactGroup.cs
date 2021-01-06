// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cms
{
    /// <summary>
    /// Provides a CMS Alarm Contact Group resource.
    /// 
    /// For information about CMS Alarm Contact Group and how to use it, see [What is Alarm Contact Group](https://www.alibabacloud.com/help/en/doc-detail/114929.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.101.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new AliCloud.Cms.AlarmContactGroup("example", new AliCloud.Cms.AlarmContactGroupArgs
    ///         {
    ///             AlarmContactGroupName = "tf-test",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CMS Alarm Contact Group can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cms/alarmContactGroup:AlarmContactGroup example tf-testacc123
    /// ```
    /// </summary>
    public partial class AlarmContactGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the alarm group.
        /// </summary>
        [Output("alarmContactGroupName")]
        public Output<string> AlarmContactGroupName { get; private set; } = null!;

        /// <summary>
        /// The name of the alert contact.
        /// </summary>
        [Output("contacts")]
        public Output<ImmutableArray<string>> Contacts { get; private set; } = null!;

        /// <summary>
        /// The description of the alert group.
        /// </summary>
        [Output("describe")]
        public Output<string?> Describe { get; private set; } = null!;

        /// <summary>
        /// Whether to open weekly subscription.
        /// </summary>
        [Output("enableSubscribed")]
        public Output<bool> EnableSubscribed { get; private set; } = null!;


        /// <summary>
        /// Create a AlarmContactGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AlarmContactGroup(string name, AlarmContactGroupArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cms/alarmContactGroup:AlarmContactGroup", name, args ?? new AlarmContactGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AlarmContactGroup(string name, Input<string> id, AlarmContactGroupState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cms/alarmContactGroup:AlarmContactGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AlarmContactGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AlarmContactGroup Get(string name, Input<string> id, AlarmContactGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new AlarmContactGroup(name, id, state, options);
        }
    }

    public sealed class AlarmContactGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the alarm group.
        /// </summary>
        [Input("alarmContactGroupName", required: true)]
        public Input<string> AlarmContactGroupName { get; set; } = null!;

        [Input("contacts")]
        private InputList<string>? _contacts;

        /// <summary>
        /// The name of the alert contact.
        /// </summary>
        public InputList<string> Contacts
        {
            get => _contacts ?? (_contacts = new InputList<string>());
            set => _contacts = value;
        }

        /// <summary>
        /// The description of the alert group.
        /// </summary>
        [Input("describe")]
        public Input<string>? Describe { get; set; }

        /// <summary>
        /// Whether to open weekly subscription.
        /// </summary>
        [Input("enableSubscribed")]
        public Input<bool>? EnableSubscribed { get; set; }

        public AlarmContactGroupArgs()
        {
        }
    }

    public sealed class AlarmContactGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the alarm group.
        /// </summary>
        [Input("alarmContactGroupName")]
        public Input<string>? AlarmContactGroupName { get; set; }

        [Input("contacts")]
        private InputList<string>? _contacts;

        /// <summary>
        /// The name of the alert contact.
        /// </summary>
        public InputList<string> Contacts
        {
            get => _contacts ?? (_contacts = new InputList<string>());
            set => _contacts = value;
        }

        /// <summary>
        /// The description of the alert group.
        /// </summary>
        [Input("describe")]
        public Input<string>? Describe { get; set; }

        /// <summary>
        /// Whether to open weekly subscription.
        /// </summary>
        [Input("enableSubscribed")]
        public Input<bool>? EnableSubscribed { get; set; }

        public AlarmContactGroupState()
        {
        }
    }
}