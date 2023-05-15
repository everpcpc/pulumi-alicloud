// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection.outputs;

import com.pulumi.alicloud.threatdetection.outputs.GetHoneypotNodesNode;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetHoneypotNodesResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of Honeypot Node IDs.
     * 
     */
    private List<String> ids;
    private @Nullable String nameRegex;
    /**
     * @return A list of Honeypot Node names.
     * 
     */
    private List<String> names;
    /**
     * @return Honeypot management node id.
     * 
     */
    private @Nullable String nodeId;
    /**
     * @return Management node name.
     * 
     */
    private @Nullable String nodeName;
    /**
     * @return A list of Honeypot Node Entries. Each element contains the following attributes:
     * 
     */
    private List<GetHoneypotNodesNode> nodes;
    private @Nullable String outputFile;
    private @Nullable Integer pageNumber;
    private @Nullable Integer pageSize;

    private GetHoneypotNodesResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list of Honeypot Node IDs.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return A list of Honeypot Node names.
     * 
     */
    public List<String> names() {
        return this.names;
    }
    /**
     * @return Honeypot management node id.
     * 
     */
    public Optional<String> nodeId() {
        return Optional.ofNullable(this.nodeId);
    }
    /**
     * @return Management node name.
     * 
     */
    public Optional<String> nodeName() {
        return Optional.ofNullable(this.nodeName);
    }
    /**
     * @return A list of Honeypot Node Entries. Each element contains the following attributes:
     * 
     */
    public List<GetHoneypotNodesNode> nodes() {
        return this.nodes;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public Optional<Integer> pageNumber() {
        return Optional.ofNullable(this.pageNumber);
    }
    public Optional<Integer> pageSize() {
        return Optional.ofNullable(this.pageSize);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetHoneypotNodesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> ids;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String nodeId;
        private @Nullable String nodeName;
        private List<GetHoneypotNodesNode> nodes;
        private @Nullable String outputFile;
        private @Nullable Integer pageNumber;
        private @Nullable Integer pageSize;
        public Builder() {}
        public Builder(GetHoneypotNodesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.nodeId = defaults.nodeId;
    	      this.nodeName = defaults.nodeName;
    	      this.nodes = defaults.nodes;
    	      this.outputFile = defaults.outputFile;
    	      this.pageNumber = defaults.pageNumber;
    	      this.pageSize = defaults.pageSize;
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
        public Builder nodeId(@Nullable String nodeId) {
            this.nodeId = nodeId;
            return this;
        }
        @CustomType.Setter
        public Builder nodeName(@Nullable String nodeName) {
            this.nodeName = nodeName;
            return this;
        }
        @CustomType.Setter
        public Builder nodes(List<GetHoneypotNodesNode> nodes) {
            this.nodes = Objects.requireNonNull(nodes);
            return this;
        }
        public Builder nodes(GetHoneypotNodesNode... nodes) {
            return nodes(List.of(nodes));
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder pageNumber(@Nullable Integer pageNumber) {
            this.pageNumber = pageNumber;
            return this;
        }
        @CustomType.Setter
        public Builder pageSize(@Nullable Integer pageSize) {
            this.pageSize = pageSize;
            return this;
        }
        public GetHoneypotNodesResult build() {
            final var o = new GetHoneypotNodesResult();
            o.id = id;
            o.ids = ids;
            o.nameRegex = nameRegex;
            o.names = names;
            o.nodeId = nodeId;
            o.nodeName = nodeName;
            o.nodes = nodes;
            o.outputFile = outputFile;
            o.pageNumber = pageNumber;
            o.pageSize = pageSize;
            return o;
        }
    }
}
