// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oss

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to put a object(content or file) to a oss bucket.
//
// ## Example Usage
// ### Uploading a file to a bucket
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/oss"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := oss.NewBucketObject(ctx, "object_source", &oss.BucketObjectArgs{
// 			Bucket: pulumi.String("your_bucket_name"),
// 			Key:    pulumi.String("new_object_key"),
// 			Source: pulumi.String("path/to/file"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Uploading a content to a bucket
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/oss"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := oss.NewBucket(ctx, "example", &oss.BucketArgs{
// 			Bucket: pulumi.String("your_bucket_name"),
// 			Acl:    pulumi.String("public-read"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = oss.NewBucketObject(ctx, "object_content", &oss.BucketObjectArgs{
// 			Bucket:  example.Bucket,
// 			Key:     pulumi.String("new_object_key"),
// 			Content: pulumi.String("the content that you want to upload."),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type BucketObject struct {
	pulumi.CustomResourceState

	// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/52284.htm) to apply. Defaults to "private".
	Acl pulumi.StringPtrOutput `pulumi:"acl"`
	// The name of the bucket to put the file in.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Specifies caching behavior along the request/reply chain. Read [RFC2616 Cache-Control](https://www.ietf.org/rfc/rfc2616.txt) for further details.
	CacheControl pulumi.StringPtrOutput `pulumi:"cacheControl"`
	// The literal content being uploaded to the bucket.
	Content pulumi.StringPtrOutput `pulumi:"content"`
	// Specifies presentational information for the object. Read [RFC2616 Content-Disposition](https://www.ietf.org/rfc/rfc2616.txt) for further details.
	ContentDisposition pulumi.StringPtrOutput `pulumi:"contentDisposition"`
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [RFC2616 Content-Encoding](https://www.ietf.org/rfc/rfc2616.txt) for further details.
	ContentEncoding pulumi.StringPtrOutput `pulumi:"contentEncoding"`
	// the content length of request.
	ContentLength pulumi.StringOutput `pulumi:"contentLength"`
	// The MD5 value of the content. Read [MD5](https://www.alibabacloud.com/help/doc-detail/31978.htm) for computing method.
	ContentMd5 pulumi.StringPtrOutput `pulumi:"contentMd5"`
	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType pulumi.StringOutput `pulumi:"contentType"`
	// the ETag generated for the object (an MD5 sum of the object content).
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Specifies expire date for the the request/response. Read [RFC2616 Expires](https://www.ietf.org/rfc/rfc2616.txt) for further details.
	Expires pulumi.StringPtrOutput `pulumi:"expires"`
	// The name of the object once it is in the bucket.
	Key pulumi.StringOutput `pulumi:"key"`
	// Specifies the primary key managed by KMS. This parameter is valid when the value of `serverSideEncryption` is set to KMS.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// Specifies server-side encryption of the object in OSS. Valid values are `AES256`, `KMS`. Default value is `AES256`.
	ServerSideEncryption pulumi.StringPtrOutput `pulumi:"serverSideEncryption"`
	// The path to the source file being uploaded to the bucket.
	Source pulumi.StringPtrOutput `pulumi:"source"`
	// A unique version ID value for the object, if bucket versioning is enabled.
	VersionId pulumi.StringOutput `pulumi:"versionId"`
}

// NewBucketObject registers a new resource with the given unique name, arguments, and options.
func NewBucketObject(ctx *pulumi.Context,
	name string, args *BucketObjectArgs, opts ...pulumi.ResourceOption) (*BucketObject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource BucketObject
	err := ctx.RegisterResource("alicloud:oss/bucketObject:BucketObject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketObject gets an existing BucketObject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketObject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketObjectState, opts ...pulumi.ResourceOption) (*BucketObject, error) {
	var resource BucketObject
	err := ctx.ReadResource("alicloud:oss/bucketObject:BucketObject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketObject resources.
type bucketObjectState struct {
	// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/52284.htm) to apply. Defaults to "private".
	Acl *string `pulumi:"acl"`
	// The name of the bucket to put the file in.
	Bucket *string `pulumi:"bucket"`
	// Specifies caching behavior along the request/reply chain. Read [RFC2616 Cache-Control](https://www.ietf.org/rfc/rfc2616.txt) for further details.
	CacheControl *string `pulumi:"cacheControl"`
	// The literal content being uploaded to the bucket.
	Content *string `pulumi:"content"`
	// Specifies presentational information for the object. Read [RFC2616 Content-Disposition](https://www.ietf.org/rfc/rfc2616.txt) for further details.
	ContentDisposition *string `pulumi:"contentDisposition"`
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [RFC2616 Content-Encoding](https://www.ietf.org/rfc/rfc2616.txt) for further details.
	ContentEncoding *string `pulumi:"contentEncoding"`
	// the content length of request.
	ContentLength *string `pulumi:"contentLength"`
	// The MD5 value of the content. Read [MD5](https://www.alibabacloud.com/help/doc-detail/31978.htm) for computing method.
	ContentMd5 *string `pulumi:"contentMd5"`
	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType *string `pulumi:"contentType"`
	// the ETag generated for the object (an MD5 sum of the object content).
	Etag *string `pulumi:"etag"`
	// Specifies expire date for the the request/response. Read [RFC2616 Expires](https://www.ietf.org/rfc/rfc2616.txt) for further details.
	Expires *string `pulumi:"expires"`
	// The name of the object once it is in the bucket.
	Key *string `pulumi:"key"`
	// Specifies the primary key managed by KMS. This parameter is valid when the value of `serverSideEncryption` is set to KMS.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Specifies server-side encryption of the object in OSS. Valid values are `AES256`, `KMS`. Default value is `AES256`.
	ServerSideEncryption *string `pulumi:"serverSideEncryption"`
	// The path to the source file being uploaded to the bucket.
	Source *string `pulumi:"source"`
	// A unique version ID value for the object, if bucket versioning is enabled.
	VersionId *string `pulumi:"versionId"`
}

type BucketObjectState struct {
	// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/52284.htm) to apply. Defaults to "private".
	Acl pulumi.StringPtrInput
	// The name of the bucket to put the file in.
	Bucket pulumi.StringPtrInput
	// Specifies caching behavior along the request/reply chain. Read [RFC2616 Cache-Control](https://www.ietf.org/rfc/rfc2616.txt) for further details.
	CacheControl pulumi.StringPtrInput
	// The literal content being uploaded to the bucket.
	Content pulumi.StringPtrInput
	// Specifies presentational information for the object. Read [RFC2616 Content-Disposition](https://www.ietf.org/rfc/rfc2616.txt) for further details.
	ContentDisposition pulumi.StringPtrInput
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [RFC2616 Content-Encoding](https://www.ietf.org/rfc/rfc2616.txt) for further details.
	ContentEncoding pulumi.StringPtrInput
	// the content length of request.
	ContentLength pulumi.StringPtrInput
	// The MD5 value of the content. Read [MD5](https://www.alibabacloud.com/help/doc-detail/31978.htm) for computing method.
	ContentMd5 pulumi.StringPtrInput
	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType pulumi.StringPtrInput
	// the ETag generated for the object (an MD5 sum of the object content).
	Etag pulumi.StringPtrInput
	// Specifies expire date for the the request/response. Read [RFC2616 Expires](https://www.ietf.org/rfc/rfc2616.txt) for further details.
	Expires pulumi.StringPtrInput
	// The name of the object once it is in the bucket.
	Key pulumi.StringPtrInput
	// Specifies the primary key managed by KMS. This parameter is valid when the value of `serverSideEncryption` is set to KMS.
	KmsKeyId pulumi.StringPtrInput
	// Specifies server-side encryption of the object in OSS. Valid values are `AES256`, `KMS`. Default value is `AES256`.
	ServerSideEncryption pulumi.StringPtrInput
	// The path to the source file being uploaded to the bucket.
	Source pulumi.StringPtrInput
	// A unique version ID value for the object, if bucket versioning is enabled.
	VersionId pulumi.StringPtrInput
}

func (BucketObjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketObjectState)(nil)).Elem()
}

type bucketObjectArgs struct {
	// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/52284.htm) to apply. Defaults to "private".
	Acl *string `pulumi:"acl"`
	// The name of the bucket to put the file in.
	Bucket string `pulumi:"bucket"`
	// Specifies caching behavior along the request/reply chain. Read [RFC2616 Cache-Control](https://www.ietf.org/rfc/rfc2616.txt) for further details.
	CacheControl *string `pulumi:"cacheControl"`
	// The literal content being uploaded to the bucket.
	Content *string `pulumi:"content"`
	// Specifies presentational information for the object. Read [RFC2616 Content-Disposition](https://www.ietf.org/rfc/rfc2616.txt) for further details.
	ContentDisposition *string `pulumi:"contentDisposition"`
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [RFC2616 Content-Encoding](https://www.ietf.org/rfc/rfc2616.txt) for further details.
	ContentEncoding *string `pulumi:"contentEncoding"`
	// The MD5 value of the content. Read [MD5](https://www.alibabacloud.com/help/doc-detail/31978.htm) for computing method.
	ContentMd5 *string `pulumi:"contentMd5"`
	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType *string `pulumi:"contentType"`
	// Specifies expire date for the the request/response. Read [RFC2616 Expires](https://www.ietf.org/rfc/rfc2616.txt) for further details.
	Expires *string `pulumi:"expires"`
	// The name of the object once it is in the bucket.
	Key string `pulumi:"key"`
	// Specifies the primary key managed by KMS. This parameter is valid when the value of `serverSideEncryption` is set to KMS.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Specifies server-side encryption of the object in OSS. Valid values are `AES256`, `KMS`. Default value is `AES256`.
	ServerSideEncryption *string `pulumi:"serverSideEncryption"`
	// The path to the source file being uploaded to the bucket.
	Source *string `pulumi:"source"`
}

// The set of arguments for constructing a BucketObject resource.
type BucketObjectArgs struct {
	// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/52284.htm) to apply. Defaults to "private".
	Acl pulumi.StringPtrInput
	// The name of the bucket to put the file in.
	Bucket pulumi.StringInput
	// Specifies caching behavior along the request/reply chain. Read [RFC2616 Cache-Control](https://www.ietf.org/rfc/rfc2616.txt) for further details.
	CacheControl pulumi.StringPtrInput
	// The literal content being uploaded to the bucket.
	Content pulumi.StringPtrInput
	// Specifies presentational information for the object. Read [RFC2616 Content-Disposition](https://www.ietf.org/rfc/rfc2616.txt) for further details.
	ContentDisposition pulumi.StringPtrInput
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [RFC2616 Content-Encoding](https://www.ietf.org/rfc/rfc2616.txt) for further details.
	ContentEncoding pulumi.StringPtrInput
	// The MD5 value of the content. Read [MD5](https://www.alibabacloud.com/help/doc-detail/31978.htm) for computing method.
	ContentMd5 pulumi.StringPtrInput
	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType pulumi.StringPtrInput
	// Specifies expire date for the the request/response. Read [RFC2616 Expires](https://www.ietf.org/rfc/rfc2616.txt) for further details.
	Expires pulumi.StringPtrInput
	// The name of the object once it is in the bucket.
	Key pulumi.StringInput
	// Specifies the primary key managed by KMS. This parameter is valid when the value of `serverSideEncryption` is set to KMS.
	KmsKeyId pulumi.StringPtrInput
	// Specifies server-side encryption of the object in OSS. Valid values are `AES256`, `KMS`. Default value is `AES256`.
	ServerSideEncryption pulumi.StringPtrInput
	// The path to the source file being uploaded to the bucket.
	Source pulumi.StringPtrInput
}

func (BucketObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketObjectArgs)(nil)).Elem()
}

type BucketObjectInput interface {
	pulumi.Input

	ToBucketObjectOutput() BucketObjectOutput
	ToBucketObjectOutputWithContext(ctx context.Context) BucketObjectOutput
}

func (BucketObject) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketObject)(nil)).Elem()
}

func (i BucketObject) ToBucketObjectOutput() BucketObjectOutput {
	return i.ToBucketObjectOutputWithContext(context.Background())
}

func (i BucketObject) ToBucketObjectOutputWithContext(ctx context.Context) BucketObjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketObjectOutput)
}

type BucketObjectOutput struct {
	*pulumi.OutputState
}

func (BucketObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketObjectOutput)(nil)).Elem()
}

func (o BucketObjectOutput) ToBucketObjectOutput() BucketObjectOutput {
	return o
}

func (o BucketObjectOutput) ToBucketObjectOutputWithContext(ctx context.Context) BucketObjectOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BucketObjectOutput{})
}
