// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oos.Outputs
{

    [OutputType]
    public sealed class GetTemplatesTemplateResult
    {
        /// <summary>
        /// The category of template.
        /// </summary>
        public readonly string Category;
        /// <summary>
        /// The creator of the template.
        /// </summary>
        public readonly string CreatedBy;
        /// <summary>
        /// The template whose creation time is less than or equal to the specified time. The format is: YYYY-MM-DDThh:mm::ssZ.
        /// </summary>
        public readonly string CreatedDate;
        /// <summary>
        /// Description of the OOS Template.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Is it triggered successfully.
        /// </summary>
        public readonly bool HasTrigger;
        /// <summary>
        /// ID of the OOS Template. The value is same as template_name.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The sharing type of the template. Valid values: `Private`, `Public`.
        /// </summary>
        public readonly string ShareType;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// The format of the template. Valid values: `JSON`, `YAML`.
        /// </summary>
        public readonly string TemplateFormat;
        /// <summary>
        /// ID of the OOS Template resource.
        /// </summary>
        public readonly string TemplateId;
        /// <summary>
        /// Name of the OOS Template.
        /// </summary>
        public readonly string TemplateName;
        /// <summary>
        /// The type of OOS Template.
        /// </summary>
        public readonly string TemplateType;
        /// <summary>
        /// Version of the OOS Template.
        /// </summary>
        public readonly string TemplateVersion;
        /// <summary>
        /// The user who updated the template.
        /// </summary>
        public readonly string UpdatedBy;
        /// <summary>
        /// The time when the template was updated.
        /// </summary>
        public readonly string UpdatedDate;

        [OutputConstructor]
        private GetTemplatesTemplateResult(
            string category,

            string createdBy,

            string createdDate,

            string description,

            bool hasTrigger,

            string id,

            string shareType,

            ImmutableDictionary<string, object> tags,

            string templateFormat,

            string templateId,

            string templateName,

            string templateType,

            string templateVersion,

            string updatedBy,

            string updatedDate)
        {
            Category = category;
            CreatedBy = createdBy;
            CreatedDate = createdDate;
            Description = description;
            HasTrigger = hasTrigger;
            Id = id;
            ShareType = shareType;
            Tags = tags;
            TemplateFormat = templateFormat;
            TemplateId = templateId;
            TemplateName = templateName;
            TemplateType = templateType;
            TemplateVersion = templateVersion;
            UpdatedBy = updatedBy;
            UpdatedDate = updatedDate;
        }
    }
}