// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const key = new alicloud.kms.Key("key", {
 *     description: "example key",
 *     isEnabled: true,
 * });
 * // Encrypt plaintext 'example'
 * const encrypted = new alicloud.kms.Ciphertext("encrypted", {
 *     keyId: key.id,
 *     plaintext: "example",
 * });
 * const plaintext = alicloud.kms.getPlaintextOutput({
 *     ciphertextBlob: encrypted.ciphertextBlob,
 * });
 * export const decrypted = plaintext.apply(plaintext => plaintext.plaintext);
 * ```
 */
export function getPlaintext(args: GetPlaintextArgs, opts?: pulumi.InvokeOptions): Promise<GetPlaintextResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("alicloud:kms/getPlaintext:getPlaintext", {
        "ciphertextBlob": args.ciphertextBlob,
        "encryptionContext": args.encryptionContext,
    }, opts);
}

/**
 * A collection of arguments for invoking getPlaintext.
 */
export interface GetPlaintextArgs {
    /**
     * The ciphertext to be decrypted.
     */
    ciphertextBlob: string;
    /**
     * -
     * (Optional) The Encryption context. If you specify this parameter in the Encrypt or GenerateDataKey API operation, it is also required when you call the Decrypt API operation. For more information, see [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm).
     */
    encryptionContext?: {[key: string]: string};
}

/**
 * A collection of values returned by getPlaintext.
 */
export interface GetPlaintextResult {
    readonly ciphertextBlob: string;
    readonly encryptionContext?: {[key: string]: string};
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The globally unique ID of the CMK. It is the ID of the CMK used to decrypt ciphertext.
     */
    readonly keyId: string;
    /**
     * The decrypted plaintext.
     */
    readonly plaintext: string;
}

export function getPlaintextOutput(args: GetPlaintextOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPlaintextResult> {
    return pulumi.output(args).apply(a => getPlaintext(a, opts))
}

/**
 * A collection of arguments for invoking getPlaintext.
 */
export interface GetPlaintextOutputArgs {
    /**
     * The ciphertext to be decrypted.
     */
    ciphertextBlob: pulumi.Input<string>;
    /**
     * -
     * (Optional) The Encryption context. If you specify this parameter in the Encrypt or GenerateDataKey API operation, it is also required when you call the Decrypt API operation. For more information, see [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm).
     */
    encryptionContext?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
