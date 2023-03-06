// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sae.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetNamespacesNamespace {
    /**
     * @return The ID of the Namespace.
     * 
     */
    private String id;
    /**
     * @return The Description of Namespace.
     * 
     */
    private String namespaceDescription;
    /**
     * @return The Id of Namespace.It can contain 2 to 32 characters.The value is in format {RegionId}:{namespace}.
     * 
     */
    private String namespaceId;
    /**
     * @return The Name of Namespace.
     * 
     */
    private String namespaceName;

    private GetNamespacesNamespace() {}
    /**
     * @return The ID of the Namespace.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The Description of Namespace.
     * 
     */
    public String namespaceDescription() {
        return this.namespaceDescription;
    }
    /**
     * @return The Id of Namespace.It can contain 2 to 32 characters.The value is in format {RegionId}:{namespace}.
     * 
     */
    public String namespaceId() {
        return this.namespaceId;
    }
    /**
     * @return The Name of Namespace.
     * 
     */
    public String namespaceName() {
        return this.namespaceName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetNamespacesNamespace defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String namespaceDescription;
        private String namespaceId;
        private String namespaceName;
        public Builder() {}
        public Builder(GetNamespacesNamespace defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.namespaceDescription = defaults.namespaceDescription;
    	      this.namespaceId = defaults.namespaceId;
    	      this.namespaceName = defaults.namespaceName;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder namespaceDescription(String namespaceDescription) {
            this.namespaceDescription = Objects.requireNonNull(namespaceDescription);
            return this;
        }
        @CustomType.Setter
        public Builder namespaceId(String namespaceId) {
            this.namespaceId = Objects.requireNonNull(namespaceId);
            return this;
        }
        @CustomType.Setter
        public Builder namespaceName(String namespaceName) {
            this.namespaceName = Objects.requireNonNull(namespaceName);
            return this;
        }
        public GetNamespacesNamespace build() {
            final var o = new GetNamespacesNamespace();
            o.id = id;
            o.namespaceDescription = namespaceDescription;
            o.namespaceId = namespaceId;
            o.namespaceName = namespaceName;
            return o;
        }
    }
}