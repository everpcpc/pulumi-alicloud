// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a key pair resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/key_pair.html.markdown.
type KeyPair struct {
	s *pulumi.ResourceState
}

// NewKeyPair registers a new resource with the given unique name, arguments, and options.
func NewKeyPair(ctx *pulumi.Context,
	name string, args *KeyPairArgs, opts ...pulumi.ResourceOpt) (*KeyPair, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["keyFile"] = nil
		inputs["keyName"] = nil
		inputs["keyNamePrefix"] = nil
		inputs["publicKey"] = nil
		inputs["resourceGroupId"] = nil
		inputs["tags"] = nil
	} else {
		inputs["keyFile"] = args.KeyFile
		inputs["keyName"] = args.KeyName
		inputs["keyNamePrefix"] = args.KeyNamePrefix
		inputs["publicKey"] = args.PublicKey
		inputs["resourceGroupId"] = args.ResourceGroupId
		inputs["tags"] = args.Tags
	}
	inputs["fingerPrint"] = nil
	s, err := ctx.RegisterResource("alicloud:ecs/keyPair:KeyPair", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &KeyPair{s: s}, nil
}

// GetKeyPair gets an existing KeyPair resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyPair(ctx *pulumi.Context,
	name string, id pulumi.ID, state *KeyPairState, opts ...pulumi.ResourceOpt) (*KeyPair, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["fingerPrint"] = state.FingerPrint
		inputs["keyFile"] = state.KeyFile
		inputs["keyName"] = state.KeyName
		inputs["keyNamePrefix"] = state.KeyNamePrefix
		inputs["publicKey"] = state.PublicKey
		inputs["resourceGroupId"] = state.ResourceGroupId
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("alicloud:ecs/keyPair:KeyPair", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &KeyPair{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *KeyPair) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *KeyPair) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *KeyPair) FingerPrint() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["fingerPrint"])
}

// The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
func (r *KeyPair) KeyFile() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["keyFile"])
}

// The key pair's name. It is the only in one Alicloud account.
func (r *KeyPair) KeyName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["keyName"])
}

func (r *KeyPair) KeyNamePrefix() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["keyNamePrefix"])
}

// You can import an existing public key and using Alicloud key pair to manage it.
func (r *KeyPair) PublicKey() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["publicKey"])
}

// The Id of resource group which the key pair belongs.
func (r *KeyPair) ResourceGroupId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupId"])
}

func (r *KeyPair) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering KeyPair resources.
type KeyPairState struct {
	FingerPrint interface{}
	// The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
	KeyFile interface{}
	// The key pair's name. It is the only in one Alicloud account.
	KeyName interface{}
	KeyNamePrefix interface{}
	// You can import an existing public key and using Alicloud key pair to manage it.
	PublicKey interface{}
	// The Id of resource group which the key pair belongs.
	ResourceGroupId interface{}
	Tags interface{}
}

// The set of arguments for constructing a KeyPair resource.
type KeyPairArgs struct {
	// The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
	KeyFile interface{}
	// The key pair's name. It is the only in one Alicloud account.
	KeyName interface{}
	KeyNamePrefix interface{}
	// You can import an existing public key and using Alicloud key pair to manage it.
	PublicKey interface{}
	// The Id of resource group which the key pair belongs.
	ResourceGroupId interface{}
	Tags interface{}
}
