module github.com/pulumi/pulumi-alicloud/provider

go 1.13

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v0.0.0-20181010200618-458213699411
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	k8s.io/api => k8s.io/api v0.0.0-20191121015604-11707872ac1c
	k8s.io/client-go => k8s.io/client-go v0.0.0-20191029021442-5f2132fc4383
)

require (
	github.com/hashicorp/terraform-plugin-sdk v1.4.1
	github.com/pulumi/pulumi-terraform-bridge v1.8.4
	github.com/pulumi/pulumi/sdk v1.13.1
	github.com/terraform-providers/terraform-provider-alicloud v1.75.0
)