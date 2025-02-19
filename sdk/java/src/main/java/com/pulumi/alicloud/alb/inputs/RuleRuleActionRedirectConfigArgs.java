// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RuleRuleActionRedirectConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final RuleRuleActionRedirectConfigArgs Empty = new RuleRuleActionRedirectConfigArgs();

    /**
     * The host name of the destination to which requests are redirected within ALB.  Valid values:  The host name must be 3 to 128 characters in length, and can contain letters, digits, hyphens (-), periods (.), asterisks (*), and question marks (?). The host name must contain at least one period (.), and cannot start or end with a period (.). The rightmost domain label can contain only letters, asterisks (*) and question marks (?) and cannot contain digits or hyphens (-). Other domain labels cannot start or end with a hyphen (-). You can include asterisks (*) and question marks (?) anywhere in a domain label. Default value: ${host}. You cannot use this value with other characters at the same time.
     * 
     */
    @Import(name="host")
    private @Nullable Output<String> host;

    /**
     * @return The host name of the destination to which requests are redirected within ALB.  Valid values:  The host name must be 3 to 128 characters in length, and can contain letters, digits, hyphens (-), periods (.), asterisks (*), and question marks (?). The host name must contain at least one period (.), and cannot start or end with a period (.). The rightmost domain label can contain only letters, asterisks (*) and question marks (?) and cannot contain digits or hyphens (-). Other domain labels cannot start or end with a hyphen (-). You can include asterisks (*) and question marks (?) anywhere in a domain label. Default value: ${host}. You cannot use this value with other characters at the same time.
     * 
     */
    public Optional<Output<String>> host() {
        return Optional.ofNullable(this.host);
    }

    /**
     * The redirect method. Valid values:301, 302, 303, 307, and 308.
     * 
     */
    @Import(name="httpCode")
    private @Nullable Output<String> httpCode;

    /**
     * @return The redirect method. Valid values:301, 302, 303, 307, and 308.
     * 
     */
    public Optional<Output<String>> httpCode() {
        return Optional.ofNullable(this.httpCode);
    }

    /**
     * The path to which requests are to be redirected within ALB.  Valid values: The path must be 1 to 128 characters in length, and start with a forward slash (/). The path can contain letters, digits, asterisks (*), question marks (?)and the following special characters: $ - _ . + / &amp; ~ @ :. It cannot contain the following special characters: &#34; % # ; ! ( ) [ ] ^ , ”. The path is case-sensitive.  Default value: ${path}. This value can be used only once. You can use it with a valid string.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The path to which requests are to be redirected within ALB.  Valid values: The path must be 1 to 128 characters in length, and start with a forward slash (/). The path can contain letters, digits, asterisks (*), question marks (?)and the following special characters: $ - _ . + / &amp; ~ @ :. It cannot contain the following special characters: &#34; % # ; ! ( ) [ ] ^ , ”. The path is case-sensitive.  Default value: ${path}. This value can be used only once. You can use it with a valid string.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * The port of the destination to which requests are redirected.  Valid values: 1 to 63335.  Default value: ${port}. You cannot use this value together with other characters at the same time.
     * 
     */
    @Import(name="port")
    private @Nullable Output<String> port;

    /**
     * @return The port of the destination to which requests are redirected.  Valid values: 1 to 63335.  Default value: ${port}. You cannot use this value together with other characters at the same time.
     * 
     */
    public Optional<Output<String>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * The protocol of the requests to be redirected.  Valid values: HTTP and HTTPS.  Default value: ${protocol}. You cannot use this value together with other characters at the same time.  Note HTTPS listeners can redirect only HTTPS requests.
     * 
     */
    @Import(name="protocol")
    private @Nullable Output<String> protocol;

    /**
     * @return The protocol of the requests to be redirected.  Valid values: HTTP and HTTPS.  Default value: ${protocol}. You cannot use this value together with other characters at the same time.  Note HTTPS listeners can redirect only HTTPS requests.
     * 
     */
    public Optional<Output<String>> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    /**
     * The query string of the request to be redirected within ALB.  The query string must be 1 to 128 characters in length, can contain letters and printable characters. It cannot contain the following special characters: # [ ] { } \ | &lt; &gt; &amp;.  Default value: ${query}. This value can be used only once. You can use it with a valid string.
     * 
     */
    @Import(name="query")
    private @Nullable Output<String> query;

    /**
     * @return The query string of the request to be redirected within ALB.  The query string must be 1 to 128 characters in length, can contain letters and printable characters. It cannot contain the following special characters: # [ ] { } \ | &lt; &gt; &amp;.  Default value: ${query}. This value can be used only once. You can use it with a valid string.
     * 
     */
    public Optional<Output<String>> query() {
        return Optional.ofNullable(this.query);
    }

    private RuleRuleActionRedirectConfigArgs() {}

    private RuleRuleActionRedirectConfigArgs(RuleRuleActionRedirectConfigArgs $) {
        this.host = $.host;
        this.httpCode = $.httpCode;
        this.path = $.path;
        this.port = $.port;
        this.protocol = $.protocol;
        this.query = $.query;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RuleRuleActionRedirectConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RuleRuleActionRedirectConfigArgs $;

        public Builder() {
            $ = new RuleRuleActionRedirectConfigArgs();
        }

        public Builder(RuleRuleActionRedirectConfigArgs defaults) {
            $ = new RuleRuleActionRedirectConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param host The host name of the destination to which requests are redirected within ALB.  Valid values:  The host name must be 3 to 128 characters in length, and can contain letters, digits, hyphens (-), periods (.), asterisks (*), and question marks (?). The host name must contain at least one period (.), and cannot start or end with a period (.). The rightmost domain label can contain only letters, asterisks (*) and question marks (?) and cannot contain digits or hyphens (-). Other domain labels cannot start or end with a hyphen (-). You can include asterisks (*) and question marks (?) anywhere in a domain label. Default value: ${host}. You cannot use this value with other characters at the same time.
         * 
         * @return builder
         * 
         */
        public Builder host(@Nullable Output<String> host) {
            $.host = host;
            return this;
        }

        /**
         * @param host The host name of the destination to which requests are redirected within ALB.  Valid values:  The host name must be 3 to 128 characters in length, and can contain letters, digits, hyphens (-), periods (.), asterisks (*), and question marks (?). The host name must contain at least one period (.), and cannot start or end with a period (.). The rightmost domain label can contain only letters, asterisks (*) and question marks (?) and cannot contain digits or hyphens (-). Other domain labels cannot start or end with a hyphen (-). You can include asterisks (*) and question marks (?) anywhere in a domain label. Default value: ${host}. You cannot use this value with other characters at the same time.
         * 
         * @return builder
         * 
         */
        public Builder host(String host) {
            return host(Output.of(host));
        }

        /**
         * @param httpCode The redirect method. Valid values:301, 302, 303, 307, and 308.
         * 
         * @return builder
         * 
         */
        public Builder httpCode(@Nullable Output<String> httpCode) {
            $.httpCode = httpCode;
            return this;
        }

        /**
         * @param httpCode The redirect method. Valid values:301, 302, 303, 307, and 308.
         * 
         * @return builder
         * 
         */
        public Builder httpCode(String httpCode) {
            return httpCode(Output.of(httpCode));
        }

        /**
         * @param path The path to which requests are to be redirected within ALB.  Valid values: The path must be 1 to 128 characters in length, and start with a forward slash (/). The path can contain letters, digits, asterisks (*), question marks (?)and the following special characters: $ - _ . + / &amp; ~ @ :. It cannot contain the following special characters: &#34; % # ; ! ( ) [ ] ^ , ”. The path is case-sensitive.  Default value: ${path}. This value can be used only once. You can use it with a valid string.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The path to which requests are to be redirected within ALB.  Valid values: The path must be 1 to 128 characters in length, and start with a forward slash (/). The path can contain letters, digits, asterisks (*), question marks (?)and the following special characters: $ - _ . + / &amp; ~ @ :. It cannot contain the following special characters: &#34; % # ; ! ( ) [ ] ^ , ”. The path is case-sensitive.  Default value: ${path}. This value can be used only once. You can use it with a valid string.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param port The port of the destination to which requests are redirected.  Valid values: 1 to 63335.  Default value: ${port}. You cannot use this value together with other characters at the same time.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<String> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The port of the destination to which requests are redirected.  Valid values: 1 to 63335.  Default value: ${port}. You cannot use this value together with other characters at the same time.
         * 
         * @return builder
         * 
         */
        public Builder port(String port) {
            return port(Output.of(port));
        }

        /**
         * @param protocol The protocol of the requests to be redirected.  Valid values: HTTP and HTTPS.  Default value: ${protocol}. You cannot use this value together with other characters at the same time.  Note HTTPS listeners can redirect only HTTPS requests.
         * 
         * @return builder
         * 
         */
        public Builder protocol(@Nullable Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol The protocol of the requests to be redirected.  Valid values: HTTP and HTTPS.  Default value: ${protocol}. You cannot use this value together with other characters at the same time.  Note HTTPS listeners can redirect only HTTPS requests.
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        /**
         * @param query The query string of the request to be redirected within ALB.  The query string must be 1 to 128 characters in length, can contain letters and printable characters. It cannot contain the following special characters: # [ ] { } \ | &lt; &gt; &amp;.  Default value: ${query}. This value can be used only once. You can use it with a valid string.
         * 
         * @return builder
         * 
         */
        public Builder query(@Nullable Output<String> query) {
            $.query = query;
            return this;
        }

        /**
         * @param query The query string of the request to be redirected within ALB.  The query string must be 1 to 128 characters in length, can contain letters and printable characters. It cannot contain the following special characters: # [ ] { } \ | &lt; &gt; &amp;.  Default value: ${query}. This value can be used only once. You can use it with a valid string.
         * 
         * @return builder
         * 
         */
        public Builder query(String query) {
            return query(Output.of(query));
        }

        public RuleRuleActionRedirectConfigArgs build() {
            return $;
        }
    }

}
