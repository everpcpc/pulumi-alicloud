// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.directmail;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class DomainArgs extends com.pulumi.resources.ResourceArgs {

    public static final DomainArgs Empty = new DomainArgs();

    /**
     * Domain, length `1` to `50`, including numbers or capitals or lowercase letters or `.` or `-`
     * 
     */
    @Import(name="domainName", required=true)
    private Output<String> domainName;

    /**
     * @return Domain, length `1` to `50`, including numbers or capitals or lowercase letters or `.` or `-`
     * 
     */
    public Output<String> domainName() {
        return this.domainName;
    }

    private DomainArgs() {}

    private DomainArgs(DomainArgs $) {
        this.domainName = $.domainName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DomainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DomainArgs $;

        public Builder() {
            $ = new DomainArgs();
        }

        public Builder(DomainArgs defaults) {
            $ = new DomainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param domainName Domain, length `1` to `50`, including numbers or capitals or lowercase letters or `.` or `-`
         * 
         * @return builder
         * 
         */
        public Builder domainName(Output<String> domainName) {
            $.domainName = domainName;
            return this;
        }

        /**
         * @param domainName Domain, length `1` to `50`, including numbers or capitals or lowercase letters or `.` or `-`
         * 
         * @return builder
         * 
         */
        public Builder domainName(String domainName) {
            return domainName(Output.of(domainName));
        }

        public DomainArgs build() {
            $.domainName = Objects.requireNonNull($.domainName, "expected parameter 'domainName' to be non-null");
            return $;
        }
    }

}
