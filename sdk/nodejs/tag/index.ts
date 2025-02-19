// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetMetaTagsArgs, GetMetaTagsResult, GetMetaTagsOutputArgs } from "./getMetaTags";
export const getMetaTags: typeof import("./getMetaTags").getMetaTags = null as any;
export const getMetaTagsOutput: typeof import("./getMetaTags").getMetaTagsOutput = null as any;
utilities.lazyLoad(exports, ["getMetaTags","getMetaTagsOutput"], () => require("./getMetaTags"));

export { PolicyArgs, PolicyState } from "./policy";
export type Policy = import("./policy").Policy;
export const Policy: typeof import("./policy").Policy = null as any;
utilities.lazyLoad(exports, ["Policy"], () => require("./policy"));

export { PolicyAttachmentArgs, PolicyAttachmentState } from "./policyAttachment";
export type PolicyAttachment = import("./policyAttachment").PolicyAttachment;
export const PolicyAttachment: typeof import("./policyAttachment").PolicyAttachment = null as any;
utilities.lazyLoad(exports, ["PolicyAttachment"], () => require("./policyAttachment"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:tag/policy:Policy":
                return new Policy(name, <any>undefined, { urn })
            case "alicloud:tag/policyAttachment:PolicyAttachment":
                return new PolicyAttachment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "tag/policy", _module)
pulumi.runtime.registerResourceModule("alicloud", "tag/policyAttachment", _module)
