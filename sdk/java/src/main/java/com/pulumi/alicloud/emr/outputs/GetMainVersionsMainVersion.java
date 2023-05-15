// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.emr.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetMainVersionsMainVersion {
    /**
     * @return A list of cluster types the emr cluster supported. Possible values: `HADOOP`, `ZOOKEEPER`, `KAFKA`, `DRUID`.
     * 
     */
    private List<String> clusterTypes;
    /**
     * @return The version of the emr cluster instance. Possible values: `EMR-4.0.0`, `EMR-3.23.0`, `EMR-3.22.0`.
     * 
     */
    private String emrVersion;
    /**
     * @return The image id of the emr cluster instance.
     * 
     */
    private String imageId;

    private GetMainVersionsMainVersion() {}
    /**
     * @return A list of cluster types the emr cluster supported. Possible values: `HADOOP`, `ZOOKEEPER`, `KAFKA`, `DRUID`.
     * 
     */
    public List<String> clusterTypes() {
        return this.clusterTypes;
    }
    /**
     * @return The version of the emr cluster instance. Possible values: `EMR-4.0.0`, `EMR-3.23.0`, `EMR-3.22.0`.
     * 
     */
    public String emrVersion() {
        return this.emrVersion;
    }
    /**
     * @return The image id of the emr cluster instance.
     * 
     */
    public String imageId() {
        return this.imageId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMainVersionsMainVersion defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> clusterTypes;
        private String emrVersion;
        private String imageId;
        public Builder() {}
        public Builder(GetMainVersionsMainVersion defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterTypes = defaults.clusterTypes;
    	      this.emrVersion = defaults.emrVersion;
    	      this.imageId = defaults.imageId;
        }

        @CustomType.Setter
        public Builder clusterTypes(List<String> clusterTypes) {
            this.clusterTypes = Objects.requireNonNull(clusterTypes);
            return this;
        }
        public Builder clusterTypes(String... clusterTypes) {
            return clusterTypes(List.of(clusterTypes));
        }
        @CustomType.Setter
        public Builder emrVersion(String emrVersion) {
            this.emrVersion = Objects.requireNonNull(emrVersion);
            return this;
        }
        @CustomType.Setter
        public Builder imageId(String imageId) {
            this.imageId = Objects.requireNonNull(imageId);
            return this;
        }
        public GetMainVersionsMainVersion build() {
            final var o = new GetMainVersionsMainVersion();
            o.clusterTypes = clusterTypes;
            o.emrVersion = emrVersion;
            o.imageId = imageId;
            return o;
        }
    }
}
