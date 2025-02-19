// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.nas.outputs;

import com.pulumi.alicloud.nas.outputs.GetAccessRulesRule;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetAccessRulesResult {
    private String accessGroupName;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of rule IDs, Each element set to `access_rule_id` (Each element formats as `&lt;access_group_name&gt;:&lt;access_rule_id&gt;` before 1.53.0).
     * 
     */
    private List<String> ids;
    private @Nullable String outputFile;
    /**
     * @return A list of AccessRules. Each element contains the following attributes:
     * 
     */
    private List<GetAccessRulesRule> rules;
    /**
     * @return RWAccess of the AccessRule.
     * 
     */
    private @Nullable String rwAccess;
    /**
     * @return SourceCidrIp of the AccessRule.
     * 
     */
    private @Nullable String sourceCidrIp;
    /**
     * @return UserAccess of the AccessRule
     * 
     */
    private @Nullable String userAccess;

    private GetAccessRulesResult() {}
    public String accessGroupName() {
        return this.accessGroupName;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list of rule IDs, Each element set to `access_rule_id` (Each element formats as `&lt;access_group_name&gt;:&lt;access_rule_id&gt;` before 1.53.0).
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return A list of AccessRules. Each element contains the following attributes:
     * 
     */
    public List<GetAccessRulesRule> rules() {
        return this.rules;
    }
    /**
     * @return RWAccess of the AccessRule.
     * 
     */
    public Optional<String> rwAccess() {
        return Optional.ofNullable(this.rwAccess);
    }
    /**
     * @return SourceCidrIp of the AccessRule.
     * 
     */
    public Optional<String> sourceCidrIp() {
        return Optional.ofNullable(this.sourceCidrIp);
    }
    /**
     * @return UserAccess of the AccessRule
     * 
     */
    public Optional<String> userAccess() {
        return Optional.ofNullable(this.userAccess);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAccessRulesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accessGroupName;
        private String id;
        private List<String> ids;
        private @Nullable String outputFile;
        private List<GetAccessRulesRule> rules;
        private @Nullable String rwAccess;
        private @Nullable String sourceCidrIp;
        private @Nullable String userAccess;
        public Builder() {}
        public Builder(GetAccessRulesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessGroupName = defaults.accessGroupName;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.outputFile = defaults.outputFile;
    	      this.rules = defaults.rules;
    	      this.rwAccess = defaults.rwAccess;
    	      this.sourceCidrIp = defaults.sourceCidrIp;
    	      this.userAccess = defaults.userAccess;
        }

        @CustomType.Setter
        public Builder accessGroupName(String accessGroupName) {
            this.accessGroupName = Objects.requireNonNull(accessGroupName);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ids(List<String> ids) {
            this.ids = Objects.requireNonNull(ids);
            return this;
        }
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder rules(List<GetAccessRulesRule> rules) {
            this.rules = Objects.requireNonNull(rules);
            return this;
        }
        public Builder rules(GetAccessRulesRule... rules) {
            return rules(List.of(rules));
        }
        @CustomType.Setter
        public Builder rwAccess(@Nullable String rwAccess) {
            this.rwAccess = rwAccess;
            return this;
        }
        @CustomType.Setter
        public Builder sourceCidrIp(@Nullable String sourceCidrIp) {
            this.sourceCidrIp = sourceCidrIp;
            return this;
        }
        @CustomType.Setter
        public Builder userAccess(@Nullable String userAccess) {
            this.userAccess = userAccess;
            return this;
        }
        public GetAccessRulesResult build() {
            final var o = new GetAccessRulesResult();
            o.accessGroupName = accessGroupName;
            o.id = id;
            o.ids = ids;
            o.outputFile = outputFile;
            o.rules = rules;
            o.rwAccess = rwAccess;
            o.sourceCidrIp = sourceCidrIp;
            o.userAccess = userAccess;
            return o;
        }
    }
}
