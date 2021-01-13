// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./acl";
export * from "./aclRule";
export * from "./clientUser";
export * from "./dnatEntry";
export * from "./getGroups";
export * from "./getInstances";
export * from "./getService";
export * from "./getTopics";
export * from "./group";
export * from "./instance";
export * from "./qos";
export * from "./qosCar";
export * from "./qosPolicy";
export * from "./snatEntry";
export * from "./topic";

// Import resources to register:
import { Acl } from "./acl";
import { AclRule } from "./aclRule";
import { ClientUser } from "./clientUser";
import { DnatEntry } from "./dnatEntry";
import { Group } from "./group";
import { Instance } from "./instance";
import { Qos } from "./qos";
import { QosCar } from "./qosCar";
import { QosPolicy } from "./qosPolicy";
import { SnatEntry } from "./snatEntry";
import { Topic } from "./topic";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:rocketmq/acl:Acl":
                return new Acl(name, <any>undefined, { urn })
            case "alicloud:rocketmq/aclRule:AclRule":
                return new AclRule(name, <any>undefined, { urn })
            case "alicloud:rocketmq/clientUser:ClientUser":
                return new ClientUser(name, <any>undefined, { urn })
            case "alicloud:rocketmq/dnatEntry:DnatEntry":
                return new DnatEntry(name, <any>undefined, { urn })
            case "alicloud:rocketmq/group:Group":
                return new Group(name, <any>undefined, { urn })
            case "alicloud:rocketmq/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "alicloud:rocketmq/qos:Qos":
                return new Qos(name, <any>undefined, { urn })
            case "alicloud:rocketmq/qosCar:QosCar":
                return new QosCar(name, <any>undefined, { urn })
            case "alicloud:rocketmq/qosPolicy:QosPolicy":
                return new QosPolicy(name, <any>undefined, { urn })
            case "alicloud:rocketmq/snatEntry:SnatEntry":
                return new SnatEntry(name, <any>undefined, { urn })
            case "alicloud:rocketmq/topic:Topic":
                return new Topic(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "rocketmq/acl", _module)
pulumi.runtime.registerResourceModule("alicloud", "rocketmq/aclRule", _module)
pulumi.runtime.registerResourceModule("alicloud", "rocketmq/clientUser", _module)
pulumi.runtime.registerResourceModule("alicloud", "rocketmq/dnatEntry", _module)
pulumi.runtime.registerResourceModule("alicloud", "rocketmq/group", _module)
pulumi.runtime.registerResourceModule("alicloud", "rocketmq/instance", _module)
pulumi.runtime.registerResourceModule("alicloud", "rocketmq/qos", _module)
pulumi.runtime.registerResourceModule("alicloud", "rocketmq/qosCar", _module)
pulumi.runtime.registerResourceModule("alicloud", "rocketmq/qosPolicy", _module)
pulumi.runtime.registerResourceModule("alicloud", "rocketmq/snatEntry", _module)
pulumi.runtime.registerResourceModule("alicloud", "rocketmq/topic", _module)
