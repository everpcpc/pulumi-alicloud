// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.nas.outputs;

import com.pulumi.alicloud.nas.outputs.GetFilesetsFileset;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetFilesetsResult {
    private String fileSystemId;
    private List<GetFilesetsFileset> filesets;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String outputFile;
    private @Nullable String status;

    private GetFilesetsResult() {}
    public String fileSystemId() {
        return this.fileSystemId;
    }
    public List<GetFilesetsFileset> filesets() {
        return this.filesets;
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
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetFilesetsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String fileSystemId;
        private List<GetFilesetsFileset> filesets;
        private String id;
        private List<String> ids;
        private @Nullable String outputFile;
        private @Nullable String status;
        public Builder() {}
        public Builder(GetFilesetsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.fileSystemId = defaults.fileSystemId;
    	      this.filesets = defaults.filesets;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.outputFile = defaults.outputFile;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder fileSystemId(String fileSystemId) {
            this.fileSystemId = Objects.requireNonNull(fileSystemId);
            return this;
        }
        @CustomType.Setter
        public Builder filesets(List<GetFilesetsFileset> filesets) {
            this.filesets = Objects.requireNonNull(filesets);
            return this;
        }
        public Builder filesets(GetFilesetsFileset... filesets) {
            return filesets(List.of(filesets));
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
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        public GetFilesetsResult build() {
            final var o = new GetFilesetsResult();
            o.fileSystemId = fileSystemId;
            o.filesets = filesets;
            o.id = id;
            o.ids = ids;
            o.outputFile = outputFile;
            o.status = status;
            return o;
        }
    }
}
