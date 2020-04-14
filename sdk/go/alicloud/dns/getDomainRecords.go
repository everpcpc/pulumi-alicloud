// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func GetDomainRecords(ctx *pulumi.Context, args *GetDomainRecordsArgs, opts ...pulumi.InvokeOption) (*GetDomainRecordsResult, error) {
	var rv GetDomainRecordsResult
	err := ctx.Invoke("alicloud:dns/getDomainRecords:getDomainRecords", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomainRecords.
type GetDomainRecordsArgs struct {
	DomainName      string   `pulumi:"domainName"`
	HostRecordRegex *string  `pulumi:"hostRecordRegex"`
	Ids             []string `pulumi:"ids"`
	IsLocked        *bool    `pulumi:"isLocked"`
	Line            *string  `pulumi:"line"`
	OutputFile      *string  `pulumi:"outputFile"`
	Status          *string  `pulumi:"status"`
	Type            *string  `pulumi:"type"`
	ValueRegex      *string  `pulumi:"valueRegex"`
}

// A collection of values returned by getDomainRecords.
type GetDomainRecordsResult struct {
	DomainName      string  `pulumi:"domainName"`
	HostRecordRegex *string `pulumi:"hostRecordRegex"`
	// id is the provider-assigned unique ID for this managed resource.
	Id         string                   `pulumi:"id"`
	Ids        []string                 `pulumi:"ids"`
	IsLocked   *bool                    `pulumi:"isLocked"`
	Line       *string                  `pulumi:"line"`
	OutputFile *string                  `pulumi:"outputFile"`
	Records    []GetDomainRecordsRecord `pulumi:"records"`
	Status     *string                  `pulumi:"status"`
	Type       *string                  `pulumi:"type"`
	Urls       []string                 `pulumi:"urls"`
	ValueRegex *string                  `pulumi:"valueRegex"`
}
