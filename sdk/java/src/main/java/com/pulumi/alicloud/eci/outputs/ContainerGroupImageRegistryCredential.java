// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eci.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ContainerGroupImageRegistryCredential {
    /**
     * @return The password used to log on to the image repository. It is required when `image_registry_credential` is configured.
     * 
     */
    private String password;
    /**
     * @return The address of the image repository. It is required when `image_registry_credential` is configured.
     * 
     */
    private String server;
    /**
     * @return The username used to log on to the image repository. It is required when `image_registry_credential` is configured.
     * 
     */
    private String userName;

    private ContainerGroupImageRegistryCredential() {}
    /**
     * @return The password used to log on to the image repository. It is required when `image_registry_credential` is configured.
     * 
     */
    public String password() {
        return this.password;
    }
    /**
     * @return The address of the image repository. It is required when `image_registry_credential` is configured.
     * 
     */
    public String server() {
        return this.server;
    }
    /**
     * @return The username used to log on to the image repository. It is required when `image_registry_credential` is configured.
     * 
     */
    public String userName() {
        return this.userName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ContainerGroupImageRegistryCredential defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String password;
        private String server;
        private String userName;
        public Builder() {}
        public Builder(ContainerGroupImageRegistryCredential defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.password = defaults.password;
    	      this.server = defaults.server;
    	      this.userName = defaults.userName;
        }

        @CustomType.Setter
        public Builder password(String password) {
            this.password = Objects.requireNonNull(password);
            return this;
        }
        @CustomType.Setter
        public Builder server(String server) {
            this.server = Objects.requireNonNull(server);
            return this;
        }
        @CustomType.Setter
        public Builder userName(String userName) {
            this.userName = Objects.requireNonNull(userName);
            return this;
        }
        public ContainerGroupImageRegistryCredential build() {
            final var o = new ContainerGroupImageRegistryCredential();
            o.password = password;
            o.server = server;
            o.userName = userName;
            return o;
        }
    }
}