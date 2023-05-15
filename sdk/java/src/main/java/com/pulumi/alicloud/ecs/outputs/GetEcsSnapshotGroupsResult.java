// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.outputs;

import com.pulumi.alicloud.ecs.outputs.GetEcsSnapshotGroupsGroup;
import com.pulumi.core.annotations.CustomType;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetEcsSnapshotGroupsResult {
    private List<GetEcsSnapshotGroupsGroup> groups;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String instanceId;
    private @Nullable String nameRegex;
    private List<String> names;
    private @Nullable String outputFile;
    private @Nullable String snapshotGroupName;
    private @Nullable String status;
    private @Nullable Map<String,Object> tags;

    private GetEcsSnapshotGroupsResult() {}
    public List<GetEcsSnapshotGroupsGroup> groups() {
        return this.groups;
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
    public Optional<String> instanceId() {
        return Optional.ofNullable(this.instanceId);
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
    public Optional<String> snapshotGroupName() {
        return Optional.ofNullable(this.snapshotGroupName);
    }
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    public Map<String,Object> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetEcsSnapshotGroupsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetEcsSnapshotGroupsGroup> groups;
        private String id;
        private List<String> ids;
        private @Nullable String instanceId;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private @Nullable String snapshotGroupName;
        private @Nullable String status;
        private @Nullable Map<String,Object> tags;
        public Builder() {}
        public Builder(GetEcsSnapshotGroupsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.groups = defaults.groups;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.instanceId = defaults.instanceId;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.snapshotGroupName = defaults.snapshotGroupName;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
        }

        @CustomType.Setter
        public Builder groups(List<GetEcsSnapshotGroupsGroup> groups) {
            this.groups = Objects.requireNonNull(groups);
            return this;
        }
        public Builder groups(GetEcsSnapshotGroupsGroup... groups) {
            return groups(List.of(groups));
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
        public Builder instanceId(@Nullable String instanceId) {
            this.instanceId = instanceId;
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
        public Builder snapshotGroupName(@Nullable String snapshotGroupName) {
            this.snapshotGroupName = snapshotGroupName;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder tags(@Nullable Map<String,Object> tags) {
            this.tags = tags;
            return this;
        }
        public GetEcsSnapshotGroupsResult build() {
            final var o = new GetEcsSnapshotGroupsResult();
            o.groups = groups;
            o.id = id;
            o.ids = ids;
            o.instanceId = instanceId;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.snapshotGroupName = snapshotGroupName;
            o.status = status;
            o.tags = tags;
            return o;
        }
    }
}
