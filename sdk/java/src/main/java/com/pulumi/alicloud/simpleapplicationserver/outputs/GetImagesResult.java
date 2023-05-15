// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.simpleapplicationserver.outputs;

import com.pulumi.alicloud.simpleapplicationserver.outputs.GetImagesImage;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetImagesResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String imageType;
    private List<GetImagesImage> images;
    private @Nullable String nameRegex;
    private List<String> names;
    private @Nullable String outputFile;
    private @Nullable String platform;

    private GetImagesResult() {}
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
    public Optional<String> imageType() {
        return Optional.ofNullable(this.imageType);
    }
    public List<GetImagesImage> images() {
        return this.images;
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    public List<String> names() {
        return this.names;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public Optional<String> platform() {
        return Optional.ofNullable(this.platform);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetImagesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> ids;
        private @Nullable String imageType;
        private List<GetImagesImage> images;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private @Nullable String platform;
        public Builder() {}
        public Builder(GetImagesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.imageType = defaults.imageType;
    	      this.images = defaults.images;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.platform = defaults.platform;
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
        public Builder imageType(@Nullable String imageType) {
            this.imageType = imageType;
            return this;
        }
        @CustomType.Setter
        public Builder images(List<GetImagesImage> images) {
            this.images = Objects.requireNonNull(images);
            return this;
        }
        public Builder images(GetImagesImage... images) {
            return images(List.of(images));
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
        public Builder platform(@Nullable String platform) {
            this.platform = platform;
            return this;
        }
        public GetImagesResult build() {
            final var o = new GetImagesResult();
            o.id = id;
            o.ids = ids;
            o.imageType = imageType;
            o.images = images;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.platform = platform;
            return o;
        }
    }
}
