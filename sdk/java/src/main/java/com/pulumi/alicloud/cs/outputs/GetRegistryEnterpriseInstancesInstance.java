// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cs.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetRegistryEnterpriseInstancesInstance {
    /**
     * @return The password that was used to log on to the registry.
     * 
     */
    private String authorizationToken;
    /**
     * @return ID of Container Registry Enterprise Edition instance.
     * 
     */
    private String id;
    /**
     * @return Name of Container Registry Enterprise Edition instance.
     * 
     */
    private String name;
    /**
     * @return The max number of namespaces that an instance can create.
     * 
     */
    private String namespaceQuota;
    /**
     * @return The number of namespaces already created.
     * 
     */
    private String namespaceUsage;
    /**
     * @return A list of domains for access on internet network.
     * 
     */
    private List<String> publicEndpoints;
    /**
     * @return Region of Container Registry Enterprise Edition instance.
     * 
     */
    private String region;
    /**
     * @return The max number of repos that an instance can create.
     * 
     */
    private String repoQuota;
    /**
     * @return The number of repos already created.
     * 
     */
    private String repoUsage;
    /**
     * @return Specification of Container Registry Enterprise Edition instance.
     * 
     */
    private String specification;
    /**
     * @return The username that was used to log on to the registry.
     * 
     */
    private String tempUsername;
    /**
     * @return A list of domains for access on vpc network.
     * 
     */
    private List<String> vpcEndpoints;

    private GetRegistryEnterpriseInstancesInstance() {}
    /**
     * @return The password that was used to log on to the registry.
     * 
     */
    public String authorizationToken() {
        return this.authorizationToken;
    }
    /**
     * @return ID of Container Registry Enterprise Edition instance.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Name of Container Registry Enterprise Edition instance.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The max number of namespaces that an instance can create.
     * 
     */
    public String namespaceQuota() {
        return this.namespaceQuota;
    }
    /**
     * @return The number of namespaces already created.
     * 
     */
    public String namespaceUsage() {
        return this.namespaceUsage;
    }
    /**
     * @return A list of domains for access on internet network.
     * 
     */
    public List<String> publicEndpoints() {
        return this.publicEndpoints;
    }
    /**
     * @return Region of Container Registry Enterprise Edition instance.
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return The max number of repos that an instance can create.
     * 
     */
    public String repoQuota() {
        return this.repoQuota;
    }
    /**
     * @return The number of repos already created.
     * 
     */
    public String repoUsage() {
        return this.repoUsage;
    }
    /**
     * @return Specification of Container Registry Enterprise Edition instance.
     * 
     */
    public String specification() {
        return this.specification;
    }
    /**
     * @return The username that was used to log on to the registry.
     * 
     */
    public String tempUsername() {
        return this.tempUsername;
    }
    /**
     * @return A list of domains for access on vpc network.
     * 
     */
    public List<String> vpcEndpoints() {
        return this.vpcEndpoints;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRegistryEnterpriseInstancesInstance defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String authorizationToken;
        private String id;
        private String name;
        private String namespaceQuota;
        private String namespaceUsage;
        private List<String> publicEndpoints;
        private String region;
        private String repoQuota;
        private String repoUsage;
        private String specification;
        private String tempUsername;
        private List<String> vpcEndpoints;
        public Builder() {}
        public Builder(GetRegistryEnterpriseInstancesInstance defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.authorizationToken = defaults.authorizationToken;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.namespaceQuota = defaults.namespaceQuota;
    	      this.namespaceUsage = defaults.namespaceUsage;
    	      this.publicEndpoints = defaults.publicEndpoints;
    	      this.region = defaults.region;
    	      this.repoQuota = defaults.repoQuota;
    	      this.repoUsage = defaults.repoUsage;
    	      this.specification = defaults.specification;
    	      this.tempUsername = defaults.tempUsername;
    	      this.vpcEndpoints = defaults.vpcEndpoints;
        }

        @CustomType.Setter
        public Builder authorizationToken(String authorizationToken) {
            this.authorizationToken = Objects.requireNonNull(authorizationToken);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder namespaceQuota(String namespaceQuota) {
            this.namespaceQuota = Objects.requireNonNull(namespaceQuota);
            return this;
        }
        @CustomType.Setter
        public Builder namespaceUsage(String namespaceUsage) {
            this.namespaceUsage = Objects.requireNonNull(namespaceUsage);
            return this;
        }
        @CustomType.Setter
        public Builder publicEndpoints(List<String> publicEndpoints) {
            this.publicEndpoints = Objects.requireNonNull(publicEndpoints);
            return this;
        }
        public Builder publicEndpoints(String... publicEndpoints) {
            return publicEndpoints(List.of(publicEndpoints));
        }
        @CustomType.Setter
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        @CustomType.Setter
        public Builder repoQuota(String repoQuota) {
            this.repoQuota = Objects.requireNonNull(repoQuota);
            return this;
        }
        @CustomType.Setter
        public Builder repoUsage(String repoUsage) {
            this.repoUsage = Objects.requireNonNull(repoUsage);
            return this;
        }
        @CustomType.Setter
        public Builder specification(String specification) {
            this.specification = Objects.requireNonNull(specification);
            return this;
        }
        @CustomType.Setter
        public Builder tempUsername(String tempUsername) {
            this.tempUsername = Objects.requireNonNull(tempUsername);
            return this;
        }
        @CustomType.Setter
        public Builder vpcEndpoints(List<String> vpcEndpoints) {
            this.vpcEndpoints = Objects.requireNonNull(vpcEndpoints);
            return this;
        }
        public Builder vpcEndpoints(String... vpcEndpoints) {
            return vpcEndpoints(List.of(vpcEndpoints));
        }
        public GetRegistryEnterpriseInstancesInstance build() {
            final var o = new GetRegistryEnterpriseInstancesInstance();
            o.authorizationToken = authorizationToken;
            o.id = id;
            o.name = name;
            o.namespaceQuota = namespaceQuota;
            o.namespaceUsage = namespaceUsage;
            o.publicEndpoints = publicEndpoints;
            o.region = region;
            o.repoQuota = repoQuota;
            o.repoUsage = repoUsage;
            o.specification = specification;
            o.tempUsername = tempUsername;
            o.vpcEndpoints = vpcEndpoints;
            return o;
        }
    }
}