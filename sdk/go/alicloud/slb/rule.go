// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A forwarding rule is configured in `HTTP`/`HTTPS` listener and it used to listen a list of backend servers which in one specified virtual backend server group.
// You can add forwarding rules to a listener to forward requests based on the domain names or the URL in the request.
//
// > **NOTE:** One virtual backend server group can be attached in multiple forwarding rules.
//
// > **NOTE:** At least one "Domain" or "Url" must be specified when creating a new rule.
//
// > **NOTE:** Having the same 'Domain' and 'Url' rule can not be created repeatedly in the one listener.
//
// > **NOTE:** Rule only be created in the `HTTP` or `HTTPS` listener.
//
// > **NOTE:** Only rule's virtual server group can be modified.
//
// ## Import
//
// Load balancer forwarding rule can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:slb/rule:Rule example rule-abc123456
// ```
type Rule struct {
	pulumi.CustomResourceState

	// The cookie configured on the server. It is mandatory when `stickySession` is "on" and `stickySessionType` is "server". Otherwise, it will be ignored. Valid value：String in line with RFC 2965, with length being 1- 200. It only contains characters such as ASCII codes, English letters and digits instead of the comma, semicolon or spacing, and it cannot start with $.
	Cookie pulumi.StringPtrOutput `pulumi:"cookie"`
	// Cookie timeout. It is mandatory when `stickySession` is "on" and `stickySessionType` is "insert". Otherwise, it will be ignored. Valid value range: [1-86400] in seconds.
	CookieTimeout pulumi.IntPtrOutput `pulumi:"cookieTimeout"`
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation pulumi.BoolPtrOutput `pulumi:"deleteProtectionValidation"`
	// Domain name of the forwarding rule. It can contain letters a-z, numbers 0-9, hyphens (-), and periods (.),
	// and wildcard characters. The following two domain name formats are supported:
	// - Standard domain name: www.test.com
	// - Wildcard domain name: *.test.com. wildcard (*) must be the first character in the format of (*.)
	Domain pulumi.StringPtrOutput `pulumi:"domain"`
	// The listener frontend port which is used to launch the new forwarding rule. Valid range: [1-65535].
	FrontendPort pulumi.IntOutput `pulumi:"frontendPort"`
	// Whether to enable health check. Valid values are`on` and `off`. TCP and UDP listener's HealthCheck is always on, so it will be ignore when launching TCP or UDP listener. This parameter is required  and takes effect only when ListenerSync is set to off.
	HealthCheck pulumi.StringPtrOutput `pulumi:"healthCheck"`
	// Port used for health check. Valid value range: [1-65535]. Default to "None" means the backend server port is used.
	HealthCheckConnectPort pulumi.IntOutput `pulumi:"healthCheckConnectPort"`
	// Domain name used for health check. When it used to launch TCP listener, `healthCheckType` must be "http". Its length is limited to 1-80 and only characters such as letters, digits, ‘-‘ and ‘.’ are allowed. When it is not set or empty,  Server Load Balancer uses the private network IP address of each backend server as Domain used for health check.
	HealthCheckDomain pulumi.StringPtrOutput `pulumi:"healthCheckDomain"`
	// Regular health check HTTP status code. Multiple codes are segmented by “,”. It is required when `healthCheck` is on. Default to `http2xx`.  Valid values are: `http2xx`,  `http3xx`, `http4xx` and `http5xx`.
	HealthCheckHttpCode pulumi.StringPtrOutput `pulumi:"healthCheckHttpCode"`
	// Time interval of health checks. It is required when `healthCheck` is on. Valid value range: [1-50] in seconds. Default to 2.
	HealthCheckInterval pulumi.IntPtrOutput `pulumi:"healthCheckInterval"`
	// Maximum timeout of each health check response. It is required when `healthCheck` is on. Valid value range: [1-300] in seconds. Default to 5. Note: If `healthCheckTimeout` < `healthCheckInterval`, its will be replaced by `healthCheckInterval`.
	HealthCheckTimeout pulumi.IntPtrOutput `pulumi:"healthCheckTimeout"`
	// URI used for health check. When it used to launch TCP listener, `healthCheckType` must be "http". Its length is limited to 1-80 and it must start with /. Only characters such as letters, digits, ‘-’, ‘/’, ‘.’, ‘%’, ‘?’, #’ and ‘&’ are allowed.
	HealthCheckUri pulumi.StringPtrOutput `pulumi:"healthCheckUri"`
	// Threshold determining the result of the health check is success. It is required when `healthCheck` is on. Valid value range: [1-10] in seconds. Default to 3.
	HealthyThreshold pulumi.IntPtrOutput `pulumi:"healthyThreshold"`
	// Indicates whether a forwarding rule inherits the settings of a health check , session persistence, and scheduling algorithm from a listener. Default to on.
	ListenerSync pulumi.StringPtrOutput `pulumi:"listenerSync"`
	// The Load Balancer ID which is used to launch the new forwarding rule.
	LoadBalancerId pulumi.StringOutput `pulumi:"loadBalancerId"`
	// Name of the forwarding rule. Our plugin provides a default name: "tf-slb-rule".
	Name pulumi.StringOutput `pulumi:"name"`
	// Scheduling algorithm, Valid values are `wrr`, `rr` and `wlc`.  Default to "wrr". This parameter is required  and takes effect only when ListenerSync is set to off.
	Scheduler pulumi.StringPtrOutput `pulumi:"scheduler"`
	// ID of a virtual server group that will be forwarded.
	ServerGroupId pulumi.StringOutput `pulumi:"serverGroupId"`
	// Whether to enable session persistence, Valid values are `on` and `off`. Default to `off`. This parameter is required  and takes effect only when ListenerSync is set to off.
	StickySession pulumi.StringPtrOutput `pulumi:"stickySession"`
	// Mode for handling the cookie. If `stickySession` is "on", it is mandatory. Otherwise, it will be ignored. Valid values are `insert` and `server`. `insert` means it is inserted from Server Load Balancer; `server` means the Server Load Balancer learns from the backend server.
	StickySessionType pulumi.StringPtrOutput `pulumi:"stickySessionType"`
	// Threshold determining the result of the health check is fail. It is required when `healthCheck` is on. Valid value range: [1-10] in seconds. Default to 3.
	UnhealthyThreshold pulumi.IntPtrOutput `pulumi:"unhealthyThreshold"`
	// Domain of the forwarding rule. It must be 2-80 characters in length. Only letters a-z, numbers 0-9,
	// and characters '-' '/' '?' '%' '#' and '&' are allowed. URLs must be started with the character '/', but cannot be '/' alone.
	Url pulumi.StringPtrOutput `pulumi:"url"`
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FrontendPort == nil {
		return nil, errors.New("invalid value for required argument 'FrontendPort'")
	}
	if args.LoadBalancerId == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerId'")
	}
	if args.ServerGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ServerGroupId'")
	}
	var resource Rule
	err := ctx.RegisterResource("alicloud:slb/rule:Rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRule gets an existing Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleState, opts ...pulumi.ResourceOption) (*Rule, error) {
	var resource Rule
	err := ctx.ReadResource("alicloud:slb/rule:Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rule resources.
type ruleState struct {
	// The cookie configured on the server. It is mandatory when `stickySession` is "on" and `stickySessionType` is "server". Otherwise, it will be ignored. Valid value：String in line with RFC 2965, with length being 1- 200. It only contains characters such as ASCII codes, English letters and digits instead of the comma, semicolon or spacing, and it cannot start with $.
	Cookie *string `pulumi:"cookie"`
	// Cookie timeout. It is mandatory when `stickySession` is "on" and `stickySessionType` is "insert". Otherwise, it will be ignored. Valid value range: [1-86400] in seconds.
	CookieTimeout *int `pulumi:"cookieTimeout"`
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation *bool `pulumi:"deleteProtectionValidation"`
	// Domain name of the forwarding rule. It can contain letters a-z, numbers 0-9, hyphens (-), and periods (.),
	// and wildcard characters. The following two domain name formats are supported:
	// - Standard domain name: www.test.com
	// - Wildcard domain name: *.test.com. wildcard (*) must be the first character in the format of (*.)
	Domain *string `pulumi:"domain"`
	// The listener frontend port which is used to launch the new forwarding rule. Valid range: [1-65535].
	FrontendPort *int `pulumi:"frontendPort"`
	// Whether to enable health check. Valid values are`on` and `off`. TCP and UDP listener's HealthCheck is always on, so it will be ignore when launching TCP or UDP listener. This parameter is required  and takes effect only when ListenerSync is set to off.
	HealthCheck *string `pulumi:"healthCheck"`
	// Port used for health check. Valid value range: [1-65535]. Default to "None" means the backend server port is used.
	HealthCheckConnectPort *int `pulumi:"healthCheckConnectPort"`
	// Domain name used for health check. When it used to launch TCP listener, `healthCheckType` must be "http". Its length is limited to 1-80 and only characters such as letters, digits, ‘-‘ and ‘.’ are allowed. When it is not set or empty,  Server Load Balancer uses the private network IP address of each backend server as Domain used for health check.
	HealthCheckDomain *string `pulumi:"healthCheckDomain"`
	// Regular health check HTTP status code. Multiple codes are segmented by “,”. It is required when `healthCheck` is on. Default to `http2xx`.  Valid values are: `http2xx`,  `http3xx`, `http4xx` and `http5xx`.
	HealthCheckHttpCode *string `pulumi:"healthCheckHttpCode"`
	// Time interval of health checks. It is required when `healthCheck` is on. Valid value range: [1-50] in seconds. Default to 2.
	HealthCheckInterval *int `pulumi:"healthCheckInterval"`
	// Maximum timeout of each health check response. It is required when `healthCheck` is on. Valid value range: [1-300] in seconds. Default to 5. Note: If `healthCheckTimeout` < `healthCheckInterval`, its will be replaced by `healthCheckInterval`.
	HealthCheckTimeout *int `pulumi:"healthCheckTimeout"`
	// URI used for health check. When it used to launch TCP listener, `healthCheckType` must be "http". Its length is limited to 1-80 and it must start with /. Only characters such as letters, digits, ‘-’, ‘/’, ‘.’, ‘%’, ‘?’, #’ and ‘&’ are allowed.
	HealthCheckUri *string `pulumi:"healthCheckUri"`
	// Threshold determining the result of the health check is success. It is required when `healthCheck` is on. Valid value range: [1-10] in seconds. Default to 3.
	HealthyThreshold *int `pulumi:"healthyThreshold"`
	// Indicates whether a forwarding rule inherits the settings of a health check , session persistence, and scheduling algorithm from a listener. Default to on.
	ListenerSync *string `pulumi:"listenerSync"`
	// The Load Balancer ID which is used to launch the new forwarding rule.
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	// Name of the forwarding rule. Our plugin provides a default name: "tf-slb-rule".
	Name *string `pulumi:"name"`
	// Scheduling algorithm, Valid values are `wrr`, `rr` and `wlc`.  Default to "wrr". This parameter is required  and takes effect only when ListenerSync is set to off.
	Scheduler *string `pulumi:"scheduler"`
	// ID of a virtual server group that will be forwarded.
	ServerGroupId *string `pulumi:"serverGroupId"`
	// Whether to enable session persistence, Valid values are `on` and `off`. Default to `off`. This parameter is required  and takes effect only when ListenerSync is set to off.
	StickySession *string `pulumi:"stickySession"`
	// Mode for handling the cookie. If `stickySession` is "on", it is mandatory. Otherwise, it will be ignored. Valid values are `insert` and `server`. `insert` means it is inserted from Server Load Balancer; `server` means the Server Load Balancer learns from the backend server.
	StickySessionType *string `pulumi:"stickySessionType"`
	// Threshold determining the result of the health check is fail. It is required when `healthCheck` is on. Valid value range: [1-10] in seconds. Default to 3.
	UnhealthyThreshold *int `pulumi:"unhealthyThreshold"`
	// Domain of the forwarding rule. It must be 2-80 characters in length. Only letters a-z, numbers 0-9,
	// and characters '-' '/' '?' '%' '#' and '&' are allowed. URLs must be started with the character '/', but cannot be '/' alone.
	Url *string `pulumi:"url"`
}

type RuleState struct {
	// The cookie configured on the server. It is mandatory when `stickySession` is "on" and `stickySessionType` is "server". Otherwise, it will be ignored. Valid value：String in line with RFC 2965, with length being 1- 200. It only contains characters such as ASCII codes, English letters and digits instead of the comma, semicolon or spacing, and it cannot start with $.
	Cookie pulumi.StringPtrInput
	// Cookie timeout. It is mandatory when `stickySession` is "on" and `stickySessionType` is "insert". Otherwise, it will be ignored. Valid value range: [1-86400] in seconds.
	CookieTimeout pulumi.IntPtrInput
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation pulumi.BoolPtrInput
	// Domain name of the forwarding rule. It can contain letters a-z, numbers 0-9, hyphens (-), and periods (.),
	// and wildcard characters. The following two domain name formats are supported:
	// - Standard domain name: www.test.com
	// - Wildcard domain name: *.test.com. wildcard (*) must be the first character in the format of (*.)
	Domain pulumi.StringPtrInput
	// The listener frontend port which is used to launch the new forwarding rule. Valid range: [1-65535].
	FrontendPort pulumi.IntPtrInput
	// Whether to enable health check. Valid values are`on` and `off`. TCP and UDP listener's HealthCheck is always on, so it will be ignore when launching TCP or UDP listener. This parameter is required  and takes effect only when ListenerSync is set to off.
	HealthCheck pulumi.StringPtrInput
	// Port used for health check. Valid value range: [1-65535]. Default to "None" means the backend server port is used.
	HealthCheckConnectPort pulumi.IntPtrInput
	// Domain name used for health check. When it used to launch TCP listener, `healthCheckType` must be "http". Its length is limited to 1-80 and only characters such as letters, digits, ‘-‘ and ‘.’ are allowed. When it is not set or empty,  Server Load Balancer uses the private network IP address of each backend server as Domain used for health check.
	HealthCheckDomain pulumi.StringPtrInput
	// Regular health check HTTP status code. Multiple codes are segmented by “,”. It is required when `healthCheck` is on. Default to `http2xx`.  Valid values are: `http2xx`,  `http3xx`, `http4xx` and `http5xx`.
	HealthCheckHttpCode pulumi.StringPtrInput
	// Time interval of health checks. It is required when `healthCheck` is on. Valid value range: [1-50] in seconds. Default to 2.
	HealthCheckInterval pulumi.IntPtrInput
	// Maximum timeout of each health check response. It is required when `healthCheck` is on. Valid value range: [1-300] in seconds. Default to 5. Note: If `healthCheckTimeout` < `healthCheckInterval`, its will be replaced by `healthCheckInterval`.
	HealthCheckTimeout pulumi.IntPtrInput
	// URI used for health check. When it used to launch TCP listener, `healthCheckType` must be "http". Its length is limited to 1-80 and it must start with /. Only characters such as letters, digits, ‘-’, ‘/’, ‘.’, ‘%’, ‘?’, #’ and ‘&’ are allowed.
	HealthCheckUri pulumi.StringPtrInput
	// Threshold determining the result of the health check is success. It is required when `healthCheck` is on. Valid value range: [1-10] in seconds. Default to 3.
	HealthyThreshold pulumi.IntPtrInput
	// Indicates whether a forwarding rule inherits the settings of a health check , session persistence, and scheduling algorithm from a listener. Default to on.
	ListenerSync pulumi.StringPtrInput
	// The Load Balancer ID which is used to launch the new forwarding rule.
	LoadBalancerId pulumi.StringPtrInput
	// Name of the forwarding rule. Our plugin provides a default name: "tf-slb-rule".
	Name pulumi.StringPtrInput
	// Scheduling algorithm, Valid values are `wrr`, `rr` and `wlc`.  Default to "wrr". This parameter is required  and takes effect only when ListenerSync is set to off.
	Scheduler pulumi.StringPtrInput
	// ID of a virtual server group that will be forwarded.
	ServerGroupId pulumi.StringPtrInput
	// Whether to enable session persistence, Valid values are `on` and `off`. Default to `off`. This parameter is required  and takes effect only when ListenerSync is set to off.
	StickySession pulumi.StringPtrInput
	// Mode for handling the cookie. If `stickySession` is "on", it is mandatory. Otherwise, it will be ignored. Valid values are `insert` and `server`. `insert` means it is inserted from Server Load Balancer; `server` means the Server Load Balancer learns from the backend server.
	StickySessionType pulumi.StringPtrInput
	// Threshold determining the result of the health check is fail. It is required when `healthCheck` is on. Valid value range: [1-10] in seconds. Default to 3.
	UnhealthyThreshold pulumi.IntPtrInput
	// Domain of the forwarding rule. It must be 2-80 characters in length. Only letters a-z, numbers 0-9,
	// and characters '-' '/' '?' '%' '#' and '&' are allowed. URLs must be started with the character '/', but cannot be '/' alone.
	Url pulumi.StringPtrInput
}

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleState)(nil)).Elem()
}

type ruleArgs struct {
	// The cookie configured on the server. It is mandatory when `stickySession` is "on" and `stickySessionType` is "server". Otherwise, it will be ignored. Valid value：String in line with RFC 2965, with length being 1- 200. It only contains characters such as ASCII codes, English letters and digits instead of the comma, semicolon or spacing, and it cannot start with $.
	Cookie *string `pulumi:"cookie"`
	// Cookie timeout. It is mandatory when `stickySession` is "on" and `stickySessionType` is "insert". Otherwise, it will be ignored. Valid value range: [1-86400] in seconds.
	CookieTimeout *int `pulumi:"cookieTimeout"`
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation *bool `pulumi:"deleteProtectionValidation"`
	// Domain name of the forwarding rule. It can contain letters a-z, numbers 0-9, hyphens (-), and periods (.),
	// and wildcard characters. The following two domain name formats are supported:
	// - Standard domain name: www.test.com
	// - Wildcard domain name: *.test.com. wildcard (*) must be the first character in the format of (*.)
	Domain *string `pulumi:"domain"`
	// The listener frontend port which is used to launch the new forwarding rule. Valid range: [1-65535].
	FrontendPort int `pulumi:"frontendPort"`
	// Whether to enable health check. Valid values are`on` and `off`. TCP and UDP listener's HealthCheck is always on, so it will be ignore when launching TCP or UDP listener. This parameter is required  and takes effect only when ListenerSync is set to off.
	HealthCheck *string `pulumi:"healthCheck"`
	// Port used for health check. Valid value range: [1-65535]. Default to "None" means the backend server port is used.
	HealthCheckConnectPort *int `pulumi:"healthCheckConnectPort"`
	// Domain name used for health check. When it used to launch TCP listener, `healthCheckType` must be "http". Its length is limited to 1-80 and only characters such as letters, digits, ‘-‘ and ‘.’ are allowed. When it is not set or empty,  Server Load Balancer uses the private network IP address of each backend server as Domain used for health check.
	HealthCheckDomain *string `pulumi:"healthCheckDomain"`
	// Regular health check HTTP status code. Multiple codes are segmented by “,”. It is required when `healthCheck` is on. Default to `http2xx`.  Valid values are: `http2xx`,  `http3xx`, `http4xx` and `http5xx`.
	HealthCheckHttpCode *string `pulumi:"healthCheckHttpCode"`
	// Time interval of health checks. It is required when `healthCheck` is on. Valid value range: [1-50] in seconds. Default to 2.
	HealthCheckInterval *int `pulumi:"healthCheckInterval"`
	// Maximum timeout of each health check response. It is required when `healthCheck` is on. Valid value range: [1-300] in seconds. Default to 5. Note: If `healthCheckTimeout` < `healthCheckInterval`, its will be replaced by `healthCheckInterval`.
	HealthCheckTimeout *int `pulumi:"healthCheckTimeout"`
	// URI used for health check. When it used to launch TCP listener, `healthCheckType` must be "http". Its length is limited to 1-80 and it must start with /. Only characters such as letters, digits, ‘-’, ‘/’, ‘.’, ‘%’, ‘?’, #’ and ‘&’ are allowed.
	HealthCheckUri *string `pulumi:"healthCheckUri"`
	// Threshold determining the result of the health check is success. It is required when `healthCheck` is on. Valid value range: [1-10] in seconds. Default to 3.
	HealthyThreshold *int `pulumi:"healthyThreshold"`
	// Indicates whether a forwarding rule inherits the settings of a health check , session persistence, and scheduling algorithm from a listener. Default to on.
	ListenerSync *string `pulumi:"listenerSync"`
	// The Load Balancer ID which is used to launch the new forwarding rule.
	LoadBalancerId string `pulumi:"loadBalancerId"`
	// Name of the forwarding rule. Our plugin provides a default name: "tf-slb-rule".
	Name *string `pulumi:"name"`
	// Scheduling algorithm, Valid values are `wrr`, `rr` and `wlc`.  Default to "wrr". This parameter is required  and takes effect only when ListenerSync is set to off.
	Scheduler *string `pulumi:"scheduler"`
	// ID of a virtual server group that will be forwarded.
	ServerGroupId string `pulumi:"serverGroupId"`
	// Whether to enable session persistence, Valid values are `on` and `off`. Default to `off`. This parameter is required  and takes effect only when ListenerSync is set to off.
	StickySession *string `pulumi:"stickySession"`
	// Mode for handling the cookie. If `stickySession` is "on", it is mandatory. Otherwise, it will be ignored. Valid values are `insert` and `server`. `insert` means it is inserted from Server Load Balancer; `server` means the Server Load Balancer learns from the backend server.
	StickySessionType *string `pulumi:"stickySessionType"`
	// Threshold determining the result of the health check is fail. It is required when `healthCheck` is on. Valid value range: [1-10] in seconds. Default to 3.
	UnhealthyThreshold *int `pulumi:"unhealthyThreshold"`
	// Domain of the forwarding rule. It must be 2-80 characters in length. Only letters a-z, numbers 0-9,
	// and characters '-' '/' '?' '%' '#' and '&' are allowed. URLs must be started with the character '/', but cannot be '/' alone.
	Url *string `pulumi:"url"`
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// The cookie configured on the server. It is mandatory when `stickySession` is "on" and `stickySessionType` is "server". Otherwise, it will be ignored. Valid value：String in line with RFC 2965, with length being 1- 200. It only contains characters such as ASCII codes, English letters and digits instead of the comma, semicolon or spacing, and it cannot start with $.
	Cookie pulumi.StringPtrInput
	// Cookie timeout. It is mandatory when `stickySession` is "on" and `stickySessionType` is "insert". Otherwise, it will be ignored. Valid value range: [1-86400] in seconds.
	CookieTimeout pulumi.IntPtrInput
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation pulumi.BoolPtrInput
	// Domain name of the forwarding rule. It can contain letters a-z, numbers 0-9, hyphens (-), and periods (.),
	// and wildcard characters. The following two domain name formats are supported:
	// - Standard domain name: www.test.com
	// - Wildcard domain name: *.test.com. wildcard (*) must be the first character in the format of (*.)
	Domain pulumi.StringPtrInput
	// The listener frontend port which is used to launch the new forwarding rule. Valid range: [1-65535].
	FrontendPort pulumi.IntInput
	// Whether to enable health check. Valid values are`on` and `off`. TCP and UDP listener's HealthCheck is always on, so it will be ignore when launching TCP or UDP listener. This parameter is required  and takes effect only when ListenerSync is set to off.
	HealthCheck pulumi.StringPtrInput
	// Port used for health check. Valid value range: [1-65535]. Default to "None" means the backend server port is used.
	HealthCheckConnectPort pulumi.IntPtrInput
	// Domain name used for health check. When it used to launch TCP listener, `healthCheckType` must be "http". Its length is limited to 1-80 and only characters such as letters, digits, ‘-‘ and ‘.’ are allowed. When it is not set or empty,  Server Load Balancer uses the private network IP address of each backend server as Domain used for health check.
	HealthCheckDomain pulumi.StringPtrInput
	// Regular health check HTTP status code. Multiple codes are segmented by “,”. It is required when `healthCheck` is on. Default to `http2xx`.  Valid values are: `http2xx`,  `http3xx`, `http4xx` and `http5xx`.
	HealthCheckHttpCode pulumi.StringPtrInput
	// Time interval of health checks. It is required when `healthCheck` is on. Valid value range: [1-50] in seconds. Default to 2.
	HealthCheckInterval pulumi.IntPtrInput
	// Maximum timeout of each health check response. It is required when `healthCheck` is on. Valid value range: [1-300] in seconds. Default to 5. Note: If `healthCheckTimeout` < `healthCheckInterval`, its will be replaced by `healthCheckInterval`.
	HealthCheckTimeout pulumi.IntPtrInput
	// URI used for health check. When it used to launch TCP listener, `healthCheckType` must be "http". Its length is limited to 1-80 and it must start with /. Only characters such as letters, digits, ‘-’, ‘/’, ‘.’, ‘%’, ‘?’, #’ and ‘&’ are allowed.
	HealthCheckUri pulumi.StringPtrInput
	// Threshold determining the result of the health check is success. It is required when `healthCheck` is on. Valid value range: [1-10] in seconds. Default to 3.
	HealthyThreshold pulumi.IntPtrInput
	// Indicates whether a forwarding rule inherits the settings of a health check , session persistence, and scheduling algorithm from a listener. Default to on.
	ListenerSync pulumi.StringPtrInput
	// The Load Balancer ID which is used to launch the new forwarding rule.
	LoadBalancerId pulumi.StringInput
	// Name of the forwarding rule. Our plugin provides a default name: "tf-slb-rule".
	Name pulumi.StringPtrInput
	// Scheduling algorithm, Valid values are `wrr`, `rr` and `wlc`.  Default to "wrr". This parameter is required  and takes effect only when ListenerSync is set to off.
	Scheduler pulumi.StringPtrInput
	// ID of a virtual server group that will be forwarded.
	ServerGroupId pulumi.StringInput
	// Whether to enable session persistence, Valid values are `on` and `off`. Default to `off`. This parameter is required  and takes effect only when ListenerSync is set to off.
	StickySession pulumi.StringPtrInput
	// Mode for handling the cookie. If `stickySession` is "on", it is mandatory. Otherwise, it will be ignored. Valid values are `insert` and `server`. `insert` means it is inserted from Server Load Balancer; `server` means the Server Load Balancer learns from the backend server.
	StickySessionType pulumi.StringPtrInput
	// Threshold determining the result of the health check is fail. It is required when `healthCheck` is on. Valid value range: [1-10] in seconds. Default to 3.
	UnhealthyThreshold pulumi.IntPtrInput
	// Domain of the forwarding rule. It must be 2-80 characters in length. Only letters a-z, numbers 0-9,
	// and characters '-' '/' '?' '%' '#' and '&' are allowed. URLs must be started with the character '/', but cannot be '/' alone.
	Url pulumi.StringPtrInput
}

func (RuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleArgs)(nil)).Elem()
}

type RuleInput interface {
	pulumi.Input

	ToRuleOutput() RuleOutput
	ToRuleOutputWithContext(ctx context.Context) RuleOutput
}

func (*Rule) ElementType() reflect.Type {
	return reflect.TypeOf((*Rule)(nil))
}

func (i *Rule) ToRuleOutput() RuleOutput {
	return i.ToRuleOutputWithContext(context.Background())
}

func (i *Rule) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleOutput)
}

func (i *Rule) ToRulePtrOutput() RulePtrOutput {
	return i.ToRulePtrOutputWithContext(context.Background())
}

func (i *Rule) ToRulePtrOutputWithContext(ctx context.Context) RulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulePtrOutput)
}

type RulePtrInput interface {
	pulumi.Input

	ToRulePtrOutput() RulePtrOutput
	ToRulePtrOutputWithContext(ctx context.Context) RulePtrOutput
}

type rulePtrType RuleArgs

func (*rulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil))
}

func (i *rulePtrType) ToRulePtrOutput() RulePtrOutput {
	return i.ToRulePtrOutputWithContext(context.Background())
}

func (i *rulePtrType) ToRulePtrOutputWithContext(ctx context.Context) RulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulePtrOutput)
}

// RuleArrayInput is an input type that accepts RuleArray and RuleArrayOutput values.
// You can construct a concrete instance of `RuleArrayInput` via:
//
//          RuleArray{ RuleArgs{...} }
type RuleArrayInput interface {
	pulumi.Input

	ToRuleArrayOutput() RuleArrayOutput
	ToRuleArrayOutputWithContext(context.Context) RuleArrayOutput
}

type RuleArray []RuleInput

func (RuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rule)(nil)).Elem()
}

func (i RuleArray) ToRuleArrayOutput() RuleArrayOutput {
	return i.ToRuleArrayOutputWithContext(context.Background())
}

func (i RuleArray) ToRuleArrayOutputWithContext(ctx context.Context) RuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleArrayOutput)
}

// RuleMapInput is an input type that accepts RuleMap and RuleMapOutput values.
// You can construct a concrete instance of `RuleMapInput` via:
//
//          RuleMap{ "key": RuleArgs{...} }
type RuleMapInput interface {
	pulumi.Input

	ToRuleMapOutput() RuleMapOutput
	ToRuleMapOutputWithContext(context.Context) RuleMapOutput
}

type RuleMap map[string]RuleInput

func (RuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rule)(nil)).Elem()
}

func (i RuleMap) ToRuleMapOutput() RuleMapOutput {
	return i.ToRuleMapOutputWithContext(context.Background())
}

func (i RuleMap) ToRuleMapOutputWithContext(ctx context.Context) RuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleMapOutput)
}

type RuleOutput struct{ *pulumi.OutputState }

func (RuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Rule)(nil))
}

func (o RuleOutput) ToRuleOutput() RuleOutput {
	return o
}

func (o RuleOutput) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return o
}

func (o RuleOutput) ToRulePtrOutput() RulePtrOutput {
	return o.ToRulePtrOutputWithContext(context.Background())
}

func (o RuleOutput) ToRulePtrOutputWithContext(ctx context.Context) RulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Rule) *Rule {
		return &v
	}).(RulePtrOutput)
}

type RulePtrOutput struct{ *pulumi.OutputState }

func (RulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil))
}

func (o RulePtrOutput) ToRulePtrOutput() RulePtrOutput {
	return o
}

func (o RulePtrOutput) ToRulePtrOutputWithContext(ctx context.Context) RulePtrOutput {
	return o
}

func (o RulePtrOutput) Elem() RuleOutput {
	return o.ApplyT(func(v *Rule) Rule {
		if v != nil {
			return *v
		}
		var ret Rule
		return ret
	}).(RuleOutput)
}

type RuleArrayOutput struct{ *pulumi.OutputState }

func (RuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Rule)(nil))
}

func (o RuleArrayOutput) ToRuleArrayOutput() RuleArrayOutput {
	return o
}

func (o RuleArrayOutput) ToRuleArrayOutputWithContext(ctx context.Context) RuleArrayOutput {
	return o
}

func (o RuleArrayOutput) Index(i pulumi.IntInput) RuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Rule {
		return vs[0].([]Rule)[vs[1].(int)]
	}).(RuleOutput)
}

type RuleMapOutput struct{ *pulumi.OutputState }

func (RuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Rule)(nil))
}

func (o RuleMapOutput) ToRuleMapOutput() RuleMapOutput {
	return o
}

func (o RuleMapOutput) ToRuleMapOutputWithContext(ctx context.Context) RuleMapOutput {
	return o
}

func (o RuleMapOutput) MapIndex(k pulumi.StringInput) RuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Rule {
		return vs[0].(map[string]Rule)[vs[1].(string)]
	}).(RuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleInput)(nil)).Elem(), &Rule{})
	pulumi.RegisterInputType(reflect.TypeOf((*RulePtrInput)(nil)).Elem(), &Rule{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleArrayInput)(nil)).Elem(), RuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleMapInput)(nil)).Elem(), RuleMap{})
	pulumi.RegisterOutputType(RuleOutput{})
	pulumi.RegisterOutputType(RulePtrOutput{})
	pulumi.RegisterOutputType(RuleArrayOutput{})
	pulumi.RegisterOutputType(RuleMapOutput{})
}
