// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb;

import com.pulumi.alicloud.alb.inputs.ListenerAccessLogTracingConfigArgs;
import com.pulumi.alicloud.alb.inputs.ListenerAclConfigArgs;
import com.pulumi.alicloud.alb.inputs.ListenerCertificatesArgs;
import com.pulumi.alicloud.alb.inputs.ListenerDefaultActionArgs;
import com.pulumi.alicloud.alb.inputs.ListenerQuicConfigArgs;
import com.pulumi.alicloud.alb.inputs.ListenerXForwardedForConfigArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ListenerArgs extends com.pulumi.resources.ResourceArgs {

    public static final ListenerArgs Empty = new ListenerArgs();

    /**
     * Indicates whether the access log has a custom header field. Valid values: true and false. Default value: false.
     * 
     * &gt; **NOTE:** Only Instances outside the Security Group to Access the Log Switch **accesslogenabled** Open, in Order to Set This Parameter to the **True**.
     * 
     */
    @Import(name="accessLogRecordCustomizedHeadersEnabled")
    private @Nullable Output<Boolean> accessLogRecordCustomizedHeadersEnabled;

    /**
     * @return Indicates whether the access log has a custom header field. Valid values: true and false. Default value: false.
     * 
     * &gt; **NOTE:** Only Instances outside the Security Group to Access the Log Switch **accesslogenabled** Open, in Order to Set This Parameter to the **True**.
     * 
     */
    public Optional<Output<Boolean>> accessLogRecordCustomizedHeadersEnabled() {
        return Optional.ofNullable(this.accessLogRecordCustomizedHeadersEnabled);
    }

    /**
     * Xtrace Configuration Information. See the following `Block access_log_tracing_config`.
     * 
     */
    @Import(name="accessLogTracingConfig")
    private @Nullable Output<ListenerAccessLogTracingConfigArgs> accessLogTracingConfig;

    /**
     * @return Xtrace Configuration Information. See the following `Block access_log_tracing_config`.
     * 
     */
    public Optional<Output<ListenerAccessLogTracingConfigArgs>> accessLogTracingConfig() {
        return Optional.ofNullable(this.accessLogTracingConfig);
    }

    /**
     * The configurations of the access control lists (ACLs). See the following `Block acl_config`. **NOTE:** Field `acl_config` has been deprecated from provider version 1.163.0, and it will be removed in the future version. Please use the new resource `alicloud.alb.ListenerAclAttachment`.,
     * 
     * @deprecated
     * Field &#39;acl_config&#39; has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_alb_listener_acl_attachment&#39;.
     * 
     */
    @Deprecated /* Field 'acl_config' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_alb_listener_acl_attachment'. */
    @Import(name="aclConfig")
    private @Nullable Output<ListenerAclConfigArgs> aclConfig;

    /**
     * @return The configurations of the access control lists (ACLs). See the following `Block acl_config`. **NOTE:** Field `acl_config` has been deprecated from provider version 1.163.0, and it will be removed in the future version. Please use the new resource `alicloud.alb.ListenerAclAttachment`.,
     * 
     * @deprecated
     * Field &#39;acl_config&#39; has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_alb_listener_acl_attachment&#39;.
     * 
     */
    @Deprecated /* Field 'acl_config' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_alb_listener_acl_attachment'. */
    public Optional<Output<ListenerAclConfigArgs>> aclConfig() {
        return Optional.ofNullable(this.aclConfig);
    }

    /**
     * The default certificate of the Listener. See the following `Block certificates`. **NOTE:** When `listener_protocol` is `HTTPS`, The default certificate must be set one。
     * 
     */
    @Import(name="certificates")
    private @Nullable Output<ListenerCertificatesArgs> certificates;

    /**
     * @return The default certificate of the Listener. See the following `Block certificates`. **NOTE:** When `listener_protocol` is `HTTPS`, The default certificate must be set one。
     * 
     */
    public Optional<Output<ListenerCertificatesArgs>> certificates() {
        return Optional.ofNullable(this.certificates);
    }

    /**
     * The Default Rule Action List. See the following `Block default_actions`.
     * 
     */
    @Import(name="defaultActions")
    private @Nullable Output<List<ListenerDefaultActionArgs>> defaultActions;

    /**
     * @return The Default Rule Action List. See the following `Block default_actions`.
     * 
     */
    public Optional<Output<List<ListenerDefaultActionArgs>>> defaultActions() {
        return Optional.ofNullable(this.defaultActions);
    }

    /**
     * The dry run.
     * 
     */
    @Import(name="dryRun")
    private @Nullable Output<Boolean> dryRun;

    /**
     * @return The dry run.
     * 
     */
    public Optional<Output<Boolean>> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }

    /**
     * Whether to Enable Gzip Compression, as a Specific File Type on a Compression. Valid values: `false`, `true`. Default Value: `true`. .
     * 
     */
    @Import(name="gzipEnabled")
    private @Nullable Output<Boolean> gzipEnabled;

    /**
     * @return Whether to Enable Gzip Compression, as a Specific File Type on a Compression. Valid values: `false`, `true`. Default Value: `true`. .
     * 
     */
    public Optional<Output<Boolean>> gzipEnabled() {
        return Optional.ofNullable(this.gzipEnabled);
    }

    /**
     * Whether to Enable HTTP/2 Features. Valid Values: `True` Or `False`. Default Value: `True`.
     * 
     * &gt; **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
     * 
     */
    @Import(name="http2Enabled")
    private @Nullable Output<Boolean> http2Enabled;

    /**
     * @return Whether to Enable HTTP/2 Features. Valid Values: `True` Or `False`. Default Value: `True`.
     * 
     * &gt; **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
     * 
     */
    public Optional<Output<Boolean>> http2Enabled() {
        return Optional.ofNullable(this.http2Enabled);
    }

    /**
     * Specify the Connection Idle Timeout Value: `1` to `60`. Unit: Seconds.
     * 
     */
    @Import(name="idleTimeout")
    private @Nullable Output<Integer> idleTimeout;

    /**
     * @return Specify the Connection Idle Timeout Value: `1` to `60`. Unit: Seconds.
     * 
     */
    public Optional<Output<Integer>> idleTimeout() {
        return Optional.ofNullable(this.idleTimeout);
    }

    /**
     * The description of the listener. The description must be 2 to 256 characters in length. The name can contain only the characters in the following string: `/^([^\x00-\xff]|[\w.,;/@-]){2,256}$/`.
     * 
     */
    @Import(name="listenerDescription")
    private @Nullable Output<String> listenerDescription;

    /**
     * @return The description of the listener. The description must be 2 to 256 characters in length. The name can contain only the characters in the following string: `/^([^\x00-\xff]|[\w.,;/@-]){2,256}$/`.
     * 
     */
    public Optional<Output<String>> listenerDescription() {
        return Optional.ofNullable(this.listenerDescription);
    }

    /**
     * The ALB Instance Front-End, and Those of the Ports Used. Value: `1` to `65535`.
     * 
     */
    @Import(name="listenerPort", required=true)
    private Output<Integer> listenerPort;

    /**
     * @return The ALB Instance Front-End, and Those of the Ports Used. Value: `1` to `65535`.
     * 
     */
    public Output<Integer> listenerPort() {
        return this.listenerPort;
    }

    /**
     * Snooping Protocols. Valid Values: `HTTP`, `HTTPS` Or `QUIC`.
     * 
     */
    @Import(name="listenerProtocol", required=true)
    private Output<String> listenerProtocol;

    /**
     * @return Snooping Protocols. Valid Values: `HTTP`, `HTTPS` Or `QUIC`.
     * 
     */
    public Output<String> listenerProtocol() {
        return this.listenerProtocol;
    }

    /**
     * The ALB Instance Id.
     * 
     */
    @Import(name="loadBalancerId", required=true)
    private Output<String> loadBalancerId;

    /**
     * @return The ALB Instance Id.
     * 
     */
    public Output<String> loadBalancerId() {
        return this.loadBalancerId;
    }

    /**
     * Configuration Associated with the QuIC Listening. See the following `Block quic_config`.
     * 
     */
    @Import(name="quicConfig")
    private @Nullable Output<ListenerQuicConfigArgs> quicConfig;

    /**
     * @return Configuration Associated with the QuIC Listening. See the following `Block quic_config`.
     * 
     */
    public Optional<Output<ListenerQuicConfigArgs>> quicConfig() {
        return Optional.ofNullable(this.quicConfig);
    }

    /**
     * The Specified Request Timeout Time. Value: `1` to `180`. Unit: Seconds. Default Value: `60`. If the Timeout Time Within the Back-End Server Has Not Answered the ALB Will Give up Waiting, the Client Returns the HTTP 504 Error Code.
     * 
     */
    @Import(name="requestTimeout")
    private @Nullable Output<Integer> requestTimeout;

    /**
     * @return The Specified Request Timeout Time. Value: `1` to `180`. Unit: Seconds. Default Value: `60`. If the Timeout Time Within the Back-End Server Has Not Answered the ALB Will Give up Waiting, the Client Returns the HTTP 504 Error Code.
     * 
     */
    public Optional<Output<Integer>> requestTimeout() {
        return Optional.ofNullable(this.requestTimeout);
    }

    /**
     * Security Policy.
     * 
     * &gt; **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
     * 
     */
    @Import(name="securityPolicyId")
    private @Nullable Output<String> securityPolicyId;

    /**
     * @return Security Policy.
     * 
     * &gt; **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
     * 
     */
    public Optional<Output<String>> securityPolicyId() {
        return Optional.ofNullable(this.securityPolicyId);
    }

    /**
     * The state of the listener. Valid Values: `Running` Or `Stopped`. Valid values: `Running`: The listener is running. `Stopped`: The listener is stopped.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The state of the listener. Valid Values: `Running` Or `Stopped`. Valid values: `Running`: The listener is running. `Stopped`: The listener is stopped.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The `x_forward_for` Related Attribute Configuration. See the following `Block x_forwarded_for_config`. **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
     * 
     */
    @Import(name="xForwardedForConfig")
    private @Nullable Output<ListenerXForwardedForConfigArgs> xForwardedForConfig;

    /**
     * @return The `x_forward_for` Related Attribute Configuration. See the following `Block x_forwarded_for_config`. **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
     * 
     */
    public Optional<Output<ListenerXForwardedForConfigArgs>> xForwardedForConfig() {
        return Optional.ofNullable(this.xForwardedForConfig);
    }

    private ListenerArgs() {}

    private ListenerArgs(ListenerArgs $) {
        this.accessLogRecordCustomizedHeadersEnabled = $.accessLogRecordCustomizedHeadersEnabled;
        this.accessLogTracingConfig = $.accessLogTracingConfig;
        this.aclConfig = $.aclConfig;
        this.certificates = $.certificates;
        this.defaultActions = $.defaultActions;
        this.dryRun = $.dryRun;
        this.gzipEnabled = $.gzipEnabled;
        this.http2Enabled = $.http2Enabled;
        this.idleTimeout = $.idleTimeout;
        this.listenerDescription = $.listenerDescription;
        this.listenerPort = $.listenerPort;
        this.listenerProtocol = $.listenerProtocol;
        this.loadBalancerId = $.loadBalancerId;
        this.quicConfig = $.quicConfig;
        this.requestTimeout = $.requestTimeout;
        this.securityPolicyId = $.securityPolicyId;
        this.status = $.status;
        this.xForwardedForConfig = $.xForwardedForConfig;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ListenerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ListenerArgs $;

        public Builder() {
            $ = new ListenerArgs();
        }

        public Builder(ListenerArgs defaults) {
            $ = new ListenerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessLogRecordCustomizedHeadersEnabled Indicates whether the access log has a custom header field. Valid values: true and false. Default value: false.
         * 
         * &gt; **NOTE:** Only Instances outside the Security Group to Access the Log Switch **accesslogenabled** Open, in Order to Set This Parameter to the **True**.
         * 
         * @return builder
         * 
         */
        public Builder accessLogRecordCustomizedHeadersEnabled(@Nullable Output<Boolean> accessLogRecordCustomizedHeadersEnabled) {
            $.accessLogRecordCustomizedHeadersEnabled = accessLogRecordCustomizedHeadersEnabled;
            return this;
        }

        /**
         * @param accessLogRecordCustomizedHeadersEnabled Indicates whether the access log has a custom header field. Valid values: true and false. Default value: false.
         * 
         * &gt; **NOTE:** Only Instances outside the Security Group to Access the Log Switch **accesslogenabled** Open, in Order to Set This Parameter to the **True**.
         * 
         * @return builder
         * 
         */
        public Builder accessLogRecordCustomizedHeadersEnabled(Boolean accessLogRecordCustomizedHeadersEnabled) {
            return accessLogRecordCustomizedHeadersEnabled(Output.of(accessLogRecordCustomizedHeadersEnabled));
        }

        /**
         * @param accessLogTracingConfig Xtrace Configuration Information. See the following `Block access_log_tracing_config`.
         * 
         * @return builder
         * 
         */
        public Builder accessLogTracingConfig(@Nullable Output<ListenerAccessLogTracingConfigArgs> accessLogTracingConfig) {
            $.accessLogTracingConfig = accessLogTracingConfig;
            return this;
        }

        /**
         * @param accessLogTracingConfig Xtrace Configuration Information. See the following `Block access_log_tracing_config`.
         * 
         * @return builder
         * 
         */
        public Builder accessLogTracingConfig(ListenerAccessLogTracingConfigArgs accessLogTracingConfig) {
            return accessLogTracingConfig(Output.of(accessLogTracingConfig));
        }

        /**
         * @param aclConfig The configurations of the access control lists (ACLs). See the following `Block acl_config`. **NOTE:** Field `acl_config` has been deprecated from provider version 1.163.0, and it will be removed in the future version. Please use the new resource `alicloud.alb.ListenerAclAttachment`.,
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;acl_config&#39; has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_alb_listener_acl_attachment&#39;.
         * 
         */
        @Deprecated /* Field 'acl_config' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_alb_listener_acl_attachment'. */
        public Builder aclConfig(@Nullable Output<ListenerAclConfigArgs> aclConfig) {
            $.aclConfig = aclConfig;
            return this;
        }

        /**
         * @param aclConfig The configurations of the access control lists (ACLs). See the following `Block acl_config`. **NOTE:** Field `acl_config` has been deprecated from provider version 1.163.0, and it will be removed in the future version. Please use the new resource `alicloud.alb.ListenerAclAttachment`.,
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;acl_config&#39; has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_alb_listener_acl_attachment&#39;.
         * 
         */
        @Deprecated /* Field 'acl_config' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_alb_listener_acl_attachment'. */
        public Builder aclConfig(ListenerAclConfigArgs aclConfig) {
            return aclConfig(Output.of(aclConfig));
        }

        /**
         * @param certificates The default certificate of the Listener. See the following `Block certificates`. **NOTE:** When `listener_protocol` is `HTTPS`, The default certificate must be set one。
         * 
         * @return builder
         * 
         */
        public Builder certificates(@Nullable Output<ListenerCertificatesArgs> certificates) {
            $.certificates = certificates;
            return this;
        }

        /**
         * @param certificates The default certificate of the Listener. See the following `Block certificates`. **NOTE:** When `listener_protocol` is `HTTPS`, The default certificate must be set one。
         * 
         * @return builder
         * 
         */
        public Builder certificates(ListenerCertificatesArgs certificates) {
            return certificates(Output.of(certificates));
        }

        /**
         * @param defaultActions The Default Rule Action List. See the following `Block default_actions`.
         * 
         * @return builder
         * 
         */
        public Builder defaultActions(@Nullable Output<List<ListenerDefaultActionArgs>> defaultActions) {
            $.defaultActions = defaultActions;
            return this;
        }

        /**
         * @param defaultActions The Default Rule Action List. See the following `Block default_actions`.
         * 
         * @return builder
         * 
         */
        public Builder defaultActions(List<ListenerDefaultActionArgs> defaultActions) {
            return defaultActions(Output.of(defaultActions));
        }

        /**
         * @param defaultActions The Default Rule Action List. See the following `Block default_actions`.
         * 
         * @return builder
         * 
         */
        public Builder defaultActions(ListenerDefaultActionArgs... defaultActions) {
            return defaultActions(List.of(defaultActions));
        }

        /**
         * @param dryRun The dry run.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(@Nullable Output<Boolean> dryRun) {
            $.dryRun = dryRun;
            return this;
        }

        /**
         * @param dryRun The dry run.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(Boolean dryRun) {
            return dryRun(Output.of(dryRun));
        }

        /**
         * @param gzipEnabled Whether to Enable Gzip Compression, as a Specific File Type on a Compression. Valid values: `false`, `true`. Default Value: `true`. .
         * 
         * @return builder
         * 
         */
        public Builder gzipEnabled(@Nullable Output<Boolean> gzipEnabled) {
            $.gzipEnabled = gzipEnabled;
            return this;
        }

        /**
         * @param gzipEnabled Whether to Enable Gzip Compression, as a Specific File Type on a Compression. Valid values: `false`, `true`. Default Value: `true`. .
         * 
         * @return builder
         * 
         */
        public Builder gzipEnabled(Boolean gzipEnabled) {
            return gzipEnabled(Output.of(gzipEnabled));
        }

        /**
         * @param http2Enabled Whether to Enable HTTP/2 Features. Valid Values: `True` Or `False`. Default Value: `True`.
         * 
         * &gt; **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
         * 
         * @return builder
         * 
         */
        public Builder http2Enabled(@Nullable Output<Boolean> http2Enabled) {
            $.http2Enabled = http2Enabled;
            return this;
        }

        /**
         * @param http2Enabled Whether to Enable HTTP/2 Features. Valid Values: `True` Or `False`. Default Value: `True`.
         * 
         * &gt; **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
         * 
         * @return builder
         * 
         */
        public Builder http2Enabled(Boolean http2Enabled) {
            return http2Enabled(Output.of(http2Enabled));
        }

        /**
         * @param idleTimeout Specify the Connection Idle Timeout Value: `1` to `60`. Unit: Seconds.
         * 
         * @return builder
         * 
         */
        public Builder idleTimeout(@Nullable Output<Integer> idleTimeout) {
            $.idleTimeout = idleTimeout;
            return this;
        }

        /**
         * @param idleTimeout Specify the Connection Idle Timeout Value: `1` to `60`. Unit: Seconds.
         * 
         * @return builder
         * 
         */
        public Builder idleTimeout(Integer idleTimeout) {
            return idleTimeout(Output.of(idleTimeout));
        }

        /**
         * @param listenerDescription The description of the listener. The description must be 2 to 256 characters in length. The name can contain only the characters in the following string: `/^([^\x00-\xff]|[\w.,;/@-]){2,256}$/`.
         * 
         * @return builder
         * 
         */
        public Builder listenerDescription(@Nullable Output<String> listenerDescription) {
            $.listenerDescription = listenerDescription;
            return this;
        }

        /**
         * @param listenerDescription The description of the listener. The description must be 2 to 256 characters in length. The name can contain only the characters in the following string: `/^([^\x00-\xff]|[\w.,;/@-]){2,256}$/`.
         * 
         * @return builder
         * 
         */
        public Builder listenerDescription(String listenerDescription) {
            return listenerDescription(Output.of(listenerDescription));
        }

        /**
         * @param listenerPort The ALB Instance Front-End, and Those of the Ports Used. Value: `1` to `65535`.
         * 
         * @return builder
         * 
         */
        public Builder listenerPort(Output<Integer> listenerPort) {
            $.listenerPort = listenerPort;
            return this;
        }

        /**
         * @param listenerPort The ALB Instance Front-End, and Those of the Ports Used. Value: `1` to `65535`.
         * 
         * @return builder
         * 
         */
        public Builder listenerPort(Integer listenerPort) {
            return listenerPort(Output.of(listenerPort));
        }

        /**
         * @param listenerProtocol Snooping Protocols. Valid Values: `HTTP`, `HTTPS` Or `QUIC`.
         * 
         * @return builder
         * 
         */
        public Builder listenerProtocol(Output<String> listenerProtocol) {
            $.listenerProtocol = listenerProtocol;
            return this;
        }

        /**
         * @param listenerProtocol Snooping Protocols. Valid Values: `HTTP`, `HTTPS` Or `QUIC`.
         * 
         * @return builder
         * 
         */
        public Builder listenerProtocol(String listenerProtocol) {
            return listenerProtocol(Output.of(listenerProtocol));
        }

        /**
         * @param loadBalancerId The ALB Instance Id.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerId(Output<String> loadBalancerId) {
            $.loadBalancerId = loadBalancerId;
            return this;
        }

        /**
         * @param loadBalancerId The ALB Instance Id.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerId(String loadBalancerId) {
            return loadBalancerId(Output.of(loadBalancerId));
        }

        /**
         * @param quicConfig Configuration Associated with the QuIC Listening. See the following `Block quic_config`.
         * 
         * @return builder
         * 
         */
        public Builder quicConfig(@Nullable Output<ListenerQuicConfigArgs> quicConfig) {
            $.quicConfig = quicConfig;
            return this;
        }

        /**
         * @param quicConfig Configuration Associated with the QuIC Listening. See the following `Block quic_config`.
         * 
         * @return builder
         * 
         */
        public Builder quicConfig(ListenerQuicConfigArgs quicConfig) {
            return quicConfig(Output.of(quicConfig));
        }

        /**
         * @param requestTimeout The Specified Request Timeout Time. Value: `1` to `180`. Unit: Seconds. Default Value: `60`. If the Timeout Time Within the Back-End Server Has Not Answered the ALB Will Give up Waiting, the Client Returns the HTTP 504 Error Code.
         * 
         * @return builder
         * 
         */
        public Builder requestTimeout(@Nullable Output<Integer> requestTimeout) {
            $.requestTimeout = requestTimeout;
            return this;
        }

        /**
         * @param requestTimeout The Specified Request Timeout Time. Value: `1` to `180`. Unit: Seconds. Default Value: `60`. If the Timeout Time Within the Back-End Server Has Not Answered the ALB Will Give up Waiting, the Client Returns the HTTP 504 Error Code.
         * 
         * @return builder
         * 
         */
        public Builder requestTimeout(Integer requestTimeout) {
            return requestTimeout(Output.of(requestTimeout));
        }

        /**
         * @param securityPolicyId Security Policy.
         * 
         * &gt; **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(@Nullable Output<String> securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        /**
         * @param securityPolicyId Security Policy.
         * 
         * &gt; **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(String securityPolicyId) {
            return securityPolicyId(Output.of(securityPolicyId));
        }

        /**
         * @param status The state of the listener. Valid Values: `Running` Or `Stopped`. Valid values: `Running`: The listener is running. `Stopped`: The listener is stopped.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The state of the listener. Valid Values: `Running` Or `Stopped`. Valid values: `Running`: The listener is running. `Stopped`: The listener is stopped.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param xForwardedForConfig The `x_forward_for` Related Attribute Configuration. See the following `Block x_forwarded_for_config`. **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
         * 
         * @return builder
         * 
         */
        public Builder xForwardedForConfig(@Nullable Output<ListenerXForwardedForConfigArgs> xForwardedForConfig) {
            $.xForwardedForConfig = xForwardedForConfig;
            return this;
        }

        /**
         * @param xForwardedForConfig The `x_forward_for` Related Attribute Configuration. See the following `Block x_forwarded_for_config`. **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
         * 
         * @return builder
         * 
         */
        public Builder xForwardedForConfig(ListenerXForwardedForConfigArgs xForwardedForConfig) {
            return xForwardedForConfig(Output.of(xForwardedForConfig));
        }

        public ListenerArgs build() {
            $.listenerPort = Objects.requireNonNull($.listenerPort, "expected parameter 'listenerPort' to be non-null");
            $.listenerProtocol = Objects.requireNonNull($.listenerProtocol, "expected parameter 'listenerProtocol' to be non-null");
            $.loadBalancerId = Objects.requireNonNull($.loadBalancerId, "expected parameter 'loadBalancerId' to be non-null");
            return $;
        }
    }

}
