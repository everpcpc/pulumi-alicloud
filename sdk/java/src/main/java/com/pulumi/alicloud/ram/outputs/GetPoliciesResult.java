// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ram.outputs;

import com.pulumi.alicloud.ram.outputs.GetPoliciesPolicy;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetPoliciesResult {
    private @Nullable Boolean enableDetails;
    private @Nullable String groupName;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String nameRegex;
    /**
     * @return A list of ram group names.
     * 
     */
    private List<String> names;
    private @Nullable String outputFile;
    /**
     * @return A list of policies. Each element contains the following attributes:
     * 
     */
    private List<GetPoliciesPolicy> policies;
    private @Nullable String roleName;
    /**
     * @return Type of the policy.
     * 
     */
    private @Nullable String type;
    private @Nullable String userName;

    private GetPoliciesResult() {}
    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }
    public Optional<String> groupName() {
        return Optional.ofNullable(this.groupName);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return A list of ram group names.
     * 
     */
    public List<String> names() {
        return this.names;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return A list of policies. Each element contains the following attributes:
     * 
     */
    public List<GetPoliciesPolicy> policies() {
        return this.policies;
    }
    public Optional<String> roleName() {
        return Optional.ofNullable(this.roleName);
    }
    /**
     * @return Type of the policy.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }
    public Optional<String> userName() {
        return Optional.ofNullable(this.userName);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPoliciesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean enableDetails;
        private @Nullable String groupName;
        private String id;
        private List<String> ids;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private List<GetPoliciesPolicy> policies;
        private @Nullable String roleName;
        private @Nullable String type;
        private @Nullable String userName;
        public Builder() {}
        public Builder(GetPoliciesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enableDetails = defaults.enableDetails;
    	      this.groupName = defaults.groupName;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.policies = defaults.policies;
    	      this.roleName = defaults.roleName;
    	      this.type = defaults.type;
    	      this.userName = defaults.userName;
        }

        @CustomType.Setter
        public Builder enableDetails(@Nullable Boolean enableDetails) {
            this.enableDetails = enableDetails;
            return this;
        }
        @CustomType.Setter
        public Builder groupName(@Nullable String groupName) {
            this.groupName = groupName;
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
        public Builder nameRegex(@Nullable String nameRegex) {
            this.nameRegex = nameRegex;
            return this;
        }
        @CustomType.Setter
        public Builder names(List<String> names) {
            this.names = Objects.requireNonNull(names);
            return this;
        }
        public Builder names(String... names) {
            return names(List.of(names));
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder policies(List<GetPoliciesPolicy> policies) {
            this.policies = Objects.requireNonNull(policies);
            return this;
        }
        public Builder policies(GetPoliciesPolicy... policies) {
            return policies(List.of(policies));
        }
        @CustomType.Setter
        public Builder roleName(@Nullable String roleName) {
            this.roleName = roleName;
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }
        @CustomType.Setter
        public Builder userName(@Nullable String userName) {
            this.userName = userName;
            return this;
        }
        public GetPoliciesResult build() {
            final var o = new GetPoliciesResult();
            o.enableDetails = enableDetails;
            o.groupName = groupName;
            o.id = id;
            o.ids = ids;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.policies = policies;
            o.roleName = roleName;
            o.type = type;
            o.userName = userName;
            return o;
        }
    }
}
