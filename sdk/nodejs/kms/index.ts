// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AliasArgs, AliasState } from "./alias";
export type Alias = import("./alias").Alias;
export const Alias: typeof import("./alias").Alias = null as any;
utilities.lazyLoad(exports, ["Alias"], () => require("./alias"));

export { CiphertextArgs, CiphertextState } from "./ciphertext";
export type Ciphertext = import("./ciphertext").Ciphertext;
export const Ciphertext: typeof import("./ciphertext").Ciphertext = null as any;
utilities.lazyLoad(exports, ["Ciphertext"], () => require("./ciphertext"));

export { GetAliasesArgs, GetAliasesResult, GetAliasesOutputArgs } from "./getAliases";
export const getAliases: typeof import("./getAliases").getAliases = null as any;
export const getAliasesOutput: typeof import("./getAliases").getAliasesOutput = null as any;
utilities.lazyLoad(exports, ["getAliases","getAliasesOutput"], () => require("./getAliases"));

export { GetCiphertextArgs, GetCiphertextResult, GetCiphertextOutputArgs } from "./getCiphertext";
export const getCiphertext: typeof import("./getCiphertext").getCiphertext = null as any;
export const getCiphertextOutput: typeof import("./getCiphertext").getCiphertextOutput = null as any;
utilities.lazyLoad(exports, ["getCiphertext","getCiphertextOutput"], () => require("./getCiphertext"));

export { GetKeyVersionsArgs, GetKeyVersionsResult, GetKeyVersionsOutputArgs } from "./getKeyVersions";
export const getKeyVersions: typeof import("./getKeyVersions").getKeyVersions = null as any;
export const getKeyVersionsOutput: typeof import("./getKeyVersions").getKeyVersionsOutput = null as any;
utilities.lazyLoad(exports, ["getKeyVersions","getKeyVersionsOutput"], () => require("./getKeyVersions"));

export { GetKeysArgs, GetKeysResult, GetKeysOutputArgs } from "./getKeys";
export const getKeys: typeof import("./getKeys").getKeys = null as any;
export const getKeysOutput: typeof import("./getKeys").getKeysOutput = null as any;
utilities.lazyLoad(exports, ["getKeys","getKeysOutput"], () => require("./getKeys"));

export { GetPlaintextArgs, GetPlaintextResult, GetPlaintextOutputArgs } from "./getPlaintext";
export const getPlaintext: typeof import("./getPlaintext").getPlaintext = null as any;
export const getPlaintextOutput: typeof import("./getPlaintext").getPlaintextOutput = null as any;
utilities.lazyLoad(exports, ["getPlaintext","getPlaintextOutput"], () => require("./getPlaintext"));

export { GetSecretVersionsArgs, GetSecretVersionsResult, GetSecretVersionsOutputArgs } from "./getSecretVersions";
export const getSecretVersions: typeof import("./getSecretVersions").getSecretVersions = null as any;
export const getSecretVersionsOutput: typeof import("./getSecretVersions").getSecretVersionsOutput = null as any;
utilities.lazyLoad(exports, ["getSecretVersions","getSecretVersionsOutput"], () => require("./getSecretVersions"));

export { GetSecretsArgs, GetSecretsResult, GetSecretsOutputArgs } from "./getSecrets";
export const getSecrets: typeof import("./getSecrets").getSecrets = null as any;
export const getSecretsOutput: typeof import("./getSecrets").getSecretsOutput = null as any;
utilities.lazyLoad(exports, ["getSecrets","getSecretsOutput"], () => require("./getSecrets"));

export { GetServiceArgs, GetServiceResult, GetServiceOutputArgs } from "./getService";
export const getService: typeof import("./getService").getService = null as any;
export const getServiceOutput: typeof import("./getService").getServiceOutput = null as any;
utilities.lazyLoad(exports, ["getService","getServiceOutput"], () => require("./getService"));

export { KeyArgs, KeyState } from "./key";
export type Key = import("./key").Key;
export const Key: typeof import("./key").Key = null as any;
utilities.lazyLoad(exports, ["Key"], () => require("./key"));

export { KeyVersionArgs, KeyVersionState } from "./keyVersion";
export type KeyVersion = import("./keyVersion").KeyVersion;
export const KeyVersion: typeof import("./keyVersion").KeyVersion = null as any;
utilities.lazyLoad(exports, ["KeyVersion"], () => require("./keyVersion"));

export { SecretArgs, SecretState } from "./secret";
export type Secret = import("./secret").Secret;
export const Secret: typeof import("./secret").Secret = null as any;
utilities.lazyLoad(exports, ["Secret"], () => require("./secret"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:kms/alias:Alias":
                return new Alias(name, <any>undefined, { urn })
            case "alicloud:kms/ciphertext:Ciphertext":
                return new Ciphertext(name, <any>undefined, { urn })
            case "alicloud:kms/key:Key":
                return new Key(name, <any>undefined, { urn })
            case "alicloud:kms/keyVersion:KeyVersion":
                return new KeyVersion(name, <any>undefined, { urn })
            case "alicloud:kms/secret:Secret":
                return new Secret(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "kms/alias", _module)
pulumi.runtime.registerResourceModule("alicloud", "kms/ciphertext", _module)
pulumi.runtime.registerResourceModule("alicloud", "kms/key", _module)
pulumi.runtime.registerResourceModule("alicloud", "kms/keyVersion", _module)
pulumi.runtime.registerResourceModule("alicloud", "kms/secret", _module)
