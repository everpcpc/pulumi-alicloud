// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.outputs;

import com.pulumi.alicloud.vpc.outputs.GetCommonBandwidthPackagesPackage;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetCommonBandwidthPackagesResult {
    /**
     * @return The name of bandwidth package.
     * 
     */
    private @Nullable String bandwidthPackageName;
    private @Nullable Boolean dryRun;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return (Optional) A list of Common Bandwidth Packages IDs.
     * 
     */
    private List<String> ids;
    private @Nullable Boolean includeReservationData;
    private @Nullable String nameRegex;
    /**
     * @return A list of Common Bandwidth Packages names.
     * 
     */
    private List<String> names;
    private @Nullable String outputFile;
    /**
     * @return A list of Common Bandwidth Packages. Each element contains the following attributes:
     * 
     */
    private List<GetCommonBandwidthPackagesPackage> packages;
    /**
     * @return The Id of resource group which the common bandwidth package belongs.
     * 
     */
    private @Nullable String resourceGroupId;
    /**
     * @return Status of the Common Bandwidth Package.
     * 
     */
    private @Nullable String status;

    private GetCommonBandwidthPackagesResult() {}
    /**
     * @return The name of bandwidth package.
     * 
     */
    public Optional<String> bandwidthPackageName() {
        return Optional.ofNullable(this.bandwidthPackageName);
    }
    public Optional<Boolean> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return (Optional) A list of Common Bandwidth Packages IDs.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public Optional<Boolean> includeReservationData() {
        return Optional.ofNullable(this.includeReservationData);
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return A list of Common Bandwidth Packages names.
     * 
     */
    public List<String> names() {
        return this.names;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return A list of Common Bandwidth Packages. Each element contains the following attributes:
     * 
     */
    public List<GetCommonBandwidthPackagesPackage> packages() {
        return this.packages;
    }
    /**
     * @return The Id of resource group which the common bandwidth package belongs.
     * 
     */
    public Optional<String> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }
    /**
     * @return Status of the Common Bandwidth Package.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCommonBandwidthPackagesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String bandwidthPackageName;
        private @Nullable Boolean dryRun;
        private String id;
        private List<String> ids;
        private @Nullable Boolean includeReservationData;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private List<GetCommonBandwidthPackagesPackage> packages;
        private @Nullable String resourceGroupId;
        private @Nullable String status;
        public Builder() {}
        public Builder(GetCommonBandwidthPackagesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bandwidthPackageName = defaults.bandwidthPackageName;
    	      this.dryRun = defaults.dryRun;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.includeReservationData = defaults.includeReservationData;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.packages = defaults.packages;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder bandwidthPackageName(@Nullable String bandwidthPackageName) {
            this.bandwidthPackageName = bandwidthPackageName;
            return this;
        }
        @CustomType.Setter
        public Builder dryRun(@Nullable Boolean dryRun) {
            this.dryRun = dryRun;
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
        public Builder includeReservationData(@Nullable Boolean includeReservationData) {
            this.includeReservationData = includeReservationData;
            return this;
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
        public Builder packages(List<GetCommonBandwidthPackagesPackage> packages) {
            this.packages = Objects.requireNonNull(packages);
            return this;
        }
        public Builder packages(GetCommonBandwidthPackagesPackage... packages) {
            return packages(List.of(packages));
        }
        @CustomType.Setter
        public Builder resourceGroupId(@Nullable String resourceGroupId) {
            this.resourceGroupId = resourceGroupId;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        public GetCommonBandwidthPackagesResult build() {
            final var o = new GetCommonBandwidthPackagesResult();
            o.bandwidthPackageName = bandwidthPackageName;
            o.dryRun = dryRun;
            o.id = id;
            o.ids = ids;
            o.includeReservationData = includeReservationData;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.packages = packages;
            o.resourceGroupId = resourceGroupId;
            o.status = status;
            return o;
        }
    }
}
