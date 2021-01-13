// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./accessGroup";
export * from "./accessRule";
export * from "./fileSystem";
export * from "./getAccessGroups";
export * from "./getAccessRules";
export * from "./getFileSystems";
export * from "./getMountTargets";
export * from "./getProtocols";
export * from "./getService";
export * from "./mountTarget";

// Import resources to register:
import { AccessGroup } from "./accessGroup";
import { AccessRule } from "./accessRule";
import { FileSystem } from "./fileSystem";
import { MountTarget } from "./mountTarget";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:nas/accessGroup:AccessGroup":
                return new AccessGroup(name, <any>undefined, { urn })
            case "alicloud:nas/accessRule:AccessRule":
                return new AccessRule(name, <any>undefined, { urn })
            case "alicloud:nas/fileSystem:FileSystem":
                return new FileSystem(name, <any>undefined, { urn })
            case "alicloud:nas/mountTarget:MountTarget":
                return new MountTarget(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "nas/accessGroup", _module)
pulumi.runtime.registerResourceModule("alicloud", "nas/accessRule", _module)
pulumi.runtime.registerResourceModule("alicloud", "nas/fileSystem", _module)
pulumi.runtime.registerResourceModule("alicloud", "nas/mountTarget", _module)
