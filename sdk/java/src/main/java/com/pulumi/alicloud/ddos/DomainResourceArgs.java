// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ddos;

import com.pulumi.alicloud.ddos.inputs.DomainResourceProxyTypeArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DomainResourceArgs extends com.pulumi.resources.ResourceArgs {

    public static final DomainResourceArgs Empty = new DomainResourceArgs();

    /**
     * The domain name of the website that you want to add to the instance.
     * 
     */
    @Import(name="domain", required=true)
    private Output<String> domain;

    /**
     * @return The domain name of the website that you want to add to the instance.
     * 
     */
    public Output<String> domain() {
        return this.domain;
    }

    /**
     * The advanced HTTPS settings. This parameter takes effect only when the value of ProxyType includes https. This parameter is a string that contains a JSON struct. The JSON struct includes the following fields:
     * 
     */
    @Import(name="httpsExt")
    private @Nullable Output<String> httpsExt;

    /**
     * @return The advanced HTTPS settings. This parameter takes effect only when the value of ProxyType includes https. This parameter is a string that contains a JSON struct. The JSON struct includes the following fields:
     * 
     */
    public Optional<Output<String>> httpsExt() {
        return Optional.ofNullable(this.httpsExt);
    }

    @Import(name="instanceIds", required=true)
    private Output<List<String>> instanceIds;

    public Output<List<String>> instanceIds() {
        return this.instanceIds;
    }

    /**
     * Protocol type and port number information.
     * 
     */
    @Import(name="proxyTypes", required=true)
    private Output<List<DomainResourceProxyTypeArgs>> proxyTypes;

    /**
     * @return Protocol type and port number information.
     * 
     */
    public Output<List<DomainResourceProxyTypeArgs>> proxyTypes() {
        return this.proxyTypes;
    }

    /**
     * the IP address. This field is required and must be a string array.
     * 
     */
    @Import(name="realServers", required=true)
    private Output<List<String>> realServers;

    /**
     * @return the IP address. This field is required and must be a string array.
     * 
     */
    public Output<List<String>> realServers() {
        return this.realServers;
    }

    /**
     * The address type of the origin server. Valid values: `0`: IP address. `1`: domain name. Use the domain name of the origin server if you deploy proxies, such as Web Application Firewall (WAF), between the origin server and the Anti-DDoS Pro or Anti-DDoS Premium instance. If you use the domain name, you must enter the address of the proxy, such as the CNAME of WAF.
     * 
     */
    @Import(name="rsType", required=true)
    private Output<Integer> rsType;

    /**
     * @return The address type of the origin server. Valid values: `0`: IP address. `1`: domain name. Use the domain name of the origin server if you deploy proxies, such as Web Application Firewall (WAF), between the origin server and the Anti-DDoS Pro or Anti-DDoS Premium instance. If you use the domain name, you must enter the address of the proxy, such as the CNAME of WAF.
     * 
     */
    public Output<Integer> rsType() {
        return this.rsType;
    }

    private DomainResourceArgs() {}

    private DomainResourceArgs(DomainResourceArgs $) {
        this.domain = $.domain;
        this.httpsExt = $.httpsExt;
        this.instanceIds = $.instanceIds;
        this.proxyTypes = $.proxyTypes;
        this.realServers = $.realServers;
        this.rsType = $.rsType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DomainResourceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DomainResourceArgs $;

        public Builder() {
            $ = new DomainResourceArgs();
        }

        public Builder(DomainResourceArgs defaults) {
            $ = new DomainResourceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param domain The domain name of the website that you want to add to the instance.
         * 
         * @return builder
         * 
         */
        public Builder domain(Output<String> domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param domain The domain name of the website that you want to add to the instance.
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        /**
         * @param httpsExt The advanced HTTPS settings. This parameter takes effect only when the value of ProxyType includes https. This parameter is a string that contains a JSON struct. The JSON struct includes the following fields:
         * 
         * @return builder
         * 
         */
        public Builder httpsExt(@Nullable Output<String> httpsExt) {
            $.httpsExt = httpsExt;
            return this;
        }

        /**
         * @param httpsExt The advanced HTTPS settings. This parameter takes effect only when the value of ProxyType includes https. This parameter is a string that contains a JSON struct. The JSON struct includes the following fields:
         * 
         * @return builder
         * 
         */
        public Builder httpsExt(String httpsExt) {
            return httpsExt(Output.of(httpsExt));
        }

        public Builder instanceIds(Output<List<String>> instanceIds) {
            $.instanceIds = instanceIds;
            return this;
        }

        public Builder instanceIds(List<String> instanceIds) {
            return instanceIds(Output.of(instanceIds));
        }

        public Builder instanceIds(String... instanceIds) {
            return instanceIds(List.of(instanceIds));
        }

        /**
         * @param proxyTypes Protocol type and port number information.
         * 
         * @return builder
         * 
         */
        public Builder proxyTypes(Output<List<DomainResourceProxyTypeArgs>> proxyTypes) {
            $.proxyTypes = proxyTypes;
            return this;
        }

        /**
         * @param proxyTypes Protocol type and port number information.
         * 
         * @return builder
         * 
         */
        public Builder proxyTypes(List<DomainResourceProxyTypeArgs> proxyTypes) {
            return proxyTypes(Output.of(proxyTypes));
        }

        /**
         * @param proxyTypes Protocol type and port number information.
         * 
         * @return builder
         * 
         */
        public Builder proxyTypes(DomainResourceProxyTypeArgs... proxyTypes) {
            return proxyTypes(List.of(proxyTypes));
        }

        /**
         * @param realServers the IP address. This field is required and must be a string array.
         * 
         * @return builder
         * 
         */
        public Builder realServers(Output<List<String>> realServers) {
            $.realServers = realServers;
            return this;
        }

        /**
         * @param realServers the IP address. This field is required and must be a string array.
         * 
         * @return builder
         * 
         */
        public Builder realServers(List<String> realServers) {
            return realServers(Output.of(realServers));
        }

        /**
         * @param realServers the IP address. This field is required and must be a string array.
         * 
         * @return builder
         * 
         */
        public Builder realServers(String... realServers) {
            return realServers(List.of(realServers));
        }

        /**
         * @param rsType The address type of the origin server. Valid values: `0`: IP address. `1`: domain name. Use the domain name of the origin server if you deploy proxies, such as Web Application Firewall (WAF), between the origin server and the Anti-DDoS Pro or Anti-DDoS Premium instance. If you use the domain name, you must enter the address of the proxy, such as the CNAME of WAF.
         * 
         * @return builder
         * 
         */
        public Builder rsType(Output<Integer> rsType) {
            $.rsType = rsType;
            return this;
        }

        /**
         * @param rsType The address type of the origin server. Valid values: `0`: IP address. `1`: domain name. Use the domain name of the origin server if you deploy proxies, such as Web Application Firewall (WAF), between the origin server and the Anti-DDoS Pro or Anti-DDoS Premium instance. If you use the domain name, you must enter the address of the proxy, such as the CNAME of WAF.
         * 
         * @return builder
         * 
         */
        public Builder rsType(Integer rsType) {
            return rsType(Output.of(rsType));
        }

        public DomainResourceArgs build() {
            $.domain = Objects.requireNonNull($.domain, "expected parameter 'domain' to be non-null");
            $.instanceIds = Objects.requireNonNull($.instanceIds, "expected parameter 'instanceIds' to be non-null");
            $.proxyTypes = Objects.requireNonNull($.proxyTypes, "expected parameter 'proxyTypes' to be non-null");
            $.realServers = Objects.requireNonNull($.realServers, "expected parameter 'realServers' to be non-null");
            $.rsType = Objects.requireNonNull($.rsType, "expected parameter 'rsType' to be non-null");
            return $;
        }
    }

}