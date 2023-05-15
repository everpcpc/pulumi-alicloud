// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dcdn.outputs;

import com.pulumi.alicloud.dcdn.outputs.GetDomainsDomain;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetDomainsResult {
    private @Nullable String changeEndTime;
    private @Nullable String changeStartTime;
    private @Nullable Boolean checkDomainShow;
    private @Nullable String domainSearchType;
    /**
     * @return A list of domains. Each element contains the following attributes:
     * 
     */
    private List<GetDomainsDomain> domains;
    private @Nullable Boolean enableDetails;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list ids of DCDN Domain.
     * 
     */
    private List<String> ids;
    private @Nullable String nameRegex;
    /**
     * @return A list of DCDN Domain names.
     * 
     */
    private List<String> names;
    private @Nullable String outputFile;
    /**
     * @return The ID of the resource group.
     * 
     */
    private @Nullable String resourceGroupId;
    private @Nullable String securityToken;
    /**
     * @return The status of DCDN Domain. Valid values: `online`, `offline`, `check_failed`, `checking`, `configure_failed`, `configuring`.
     * 
     */
    private @Nullable String status;

    private GetDomainsResult() {}
    public Optional<String> changeEndTime() {
        return Optional.ofNullable(this.changeEndTime);
    }
    public Optional<String> changeStartTime() {
        return Optional.ofNullable(this.changeStartTime);
    }
    public Optional<Boolean> checkDomainShow() {
        return Optional.ofNullable(this.checkDomainShow);
    }
    public Optional<String> domainSearchType() {
        return Optional.ofNullable(this.domainSearchType);
    }
    /**
     * @return A list of domains. Each element contains the following attributes:
     * 
     */
    public List<GetDomainsDomain> domains() {
        return this.domains;
    }
    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list ids of DCDN Domain.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return A list of DCDN Domain names.
     * 
     */
    public List<String> names() {
        return this.names;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return The ID of the resource group.
     * 
     */
    public Optional<String> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }
    public Optional<String> securityToken() {
        return Optional.ofNullable(this.securityToken);
    }
    /**
     * @return The status of DCDN Domain. Valid values: `online`, `offline`, `check_failed`, `checking`, `configure_failed`, `configuring`.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDomainsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String changeEndTime;
        private @Nullable String changeStartTime;
        private @Nullable Boolean checkDomainShow;
        private @Nullable String domainSearchType;
        private List<GetDomainsDomain> domains;
        private @Nullable Boolean enableDetails;
        private String id;
        private List<String> ids;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private @Nullable String resourceGroupId;
        private @Nullable String securityToken;
        private @Nullable String status;
        public Builder() {}
        public Builder(GetDomainsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.changeEndTime = defaults.changeEndTime;
    	      this.changeStartTime = defaults.changeStartTime;
    	      this.checkDomainShow = defaults.checkDomainShow;
    	      this.domainSearchType = defaults.domainSearchType;
    	      this.domains = defaults.domains;
    	      this.enableDetails = defaults.enableDetails;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.securityToken = defaults.securityToken;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder changeEndTime(@Nullable String changeEndTime) {
            this.changeEndTime = changeEndTime;
            return this;
        }
        @CustomType.Setter
        public Builder changeStartTime(@Nullable String changeStartTime) {
            this.changeStartTime = changeStartTime;
            return this;
        }
        @CustomType.Setter
        public Builder checkDomainShow(@Nullable Boolean checkDomainShow) {
            this.checkDomainShow = checkDomainShow;
            return this;
        }
        @CustomType.Setter
        public Builder domainSearchType(@Nullable String domainSearchType) {
            this.domainSearchType = domainSearchType;
            return this;
        }
        @CustomType.Setter
        public Builder domains(List<GetDomainsDomain> domains) {
            this.domains = Objects.requireNonNull(domains);
            return this;
        }
        public Builder domains(GetDomainsDomain... domains) {
            return domains(List.of(domains));
        }
        @CustomType.Setter
        public Builder enableDetails(@Nullable Boolean enableDetails) {
            this.enableDetails = enableDetails;
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
        public Builder resourceGroupId(@Nullable String resourceGroupId) {
            this.resourceGroupId = resourceGroupId;
            return this;
        }
        @CustomType.Setter
        public Builder securityToken(@Nullable String securityToken) {
            this.securityToken = securityToken;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        public GetDomainsResult build() {
            final var o = new GetDomainsResult();
            o.changeEndTime = changeEndTime;
            o.changeStartTime = changeStartTime;
            o.checkDomainShow = checkDomainShow;
            o.domainSearchType = domainSearchType;
            o.domains = domains;
            o.enableDetails = enableDetails;
            o.id = id;
            o.ids = ids;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.resourceGroupId = resourceGroupId;
            o.securityToken = securityToken;
            o.status = status;
            return o;
        }
    }
}
