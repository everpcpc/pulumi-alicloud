// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package kms

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/kms_plaintext.html.markdown.
func GetPlaintext(ctx *pulumi.Context, args *GetPlaintextArgs, opts ...pulumi.InvokeOption) (*GetPlaintextResult, error) {
	var rv GetPlaintextResult
	err := ctx.Invoke("alicloud:kms/getPlaintext:getPlaintext", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPlaintext.
type GetPlaintextArgs struct {
	// The ciphertext to be decrypted.
	CiphertextBlob string `pulumi:"ciphertextBlob"`
	EncryptionContext map[string]string `pulumi:"encryptionContext"`
}


// A collection of values returned by getPlaintext.
type GetPlaintextResult struct {
	CiphertextBlob string `pulumi:"ciphertextBlob"`
	EncryptionContext map[string]string `pulumi:"encryptionContext"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The globally unique ID of the CMK. It is the ID of the CMK used to decrypt ciphertext.
	KeyId string `pulumi:"keyId"`
	// The decrypted plaintext.
	Plaintext string `pulumi:"plaintext"`
}

