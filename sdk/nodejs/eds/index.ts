// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./ecdPolicyGroup";
export * from "./getNasFileSystems";
export * from "./getPolicyGroups";
export * from "./getSimpleOfficeSites";
export * from "./nasFileSystem";
export * from "./simpleOfficeSite";

// Import resources to register:
import { EcdPolicyGroup } from "./ecdPolicyGroup";
import { NasFileSystem } from "./nasFileSystem";
import { SimpleOfficeSite } from "./simpleOfficeSite";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:eds/ecdPolicyGroup:EcdPolicyGroup":
                return new EcdPolicyGroup(name, <any>undefined, { urn })
            case "alicloud:eds/nasFileSystem:NasFileSystem":
                return new NasFileSystem(name, <any>undefined, { urn })
            case "alicloud:eds/simpleOfficeSite:SimpleOfficeSite":
                return new SimpleOfficeSite(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "eds/ecdPolicyGroup", _module)
pulumi.runtime.registerResourceModule("alicloud", "eds/nasFileSystem", _module)
pulumi.runtime.registerResourceModule("alicloud", "eds/simpleOfficeSite", _module)
