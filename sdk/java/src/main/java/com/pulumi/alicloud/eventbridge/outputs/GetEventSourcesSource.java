// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eventbridge.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetEventSourcesSource {
    /**
     * @return The detail describe of event source.
     * 
     */
    private String description;
    /**
     * @return The code name of event source.
     * 
     */
    private String eventSourceName;
    /**
     * @return The config of external data source.
     * 
     */
    private Map<String,Object> externalSourceConfig;
    /**
     * @return The type of external data source.
     * 
     */
    private String externalSourceType;
    /**
     * @return The ID of the Event Source.
     * 
     */
    private String id;
    /**
     * @return Whether to connect to an external data source.
     * 
     */
    private Boolean linkedExternalSource;
    private String type;

    private GetEventSourcesSource() {}
    /**
     * @return The detail describe of event source.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The code name of event source.
     * 
     */
    public String eventSourceName() {
        return this.eventSourceName;
    }
    /**
     * @return The config of external data source.
     * 
     */
    public Map<String,Object> externalSourceConfig() {
        return this.externalSourceConfig;
    }
    /**
     * @return The type of external data source.
     * 
     */
    public String externalSourceType() {
        return this.externalSourceType;
    }
    /**
     * @return The ID of the Event Source.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Whether to connect to an external data source.
     * 
     */
    public Boolean linkedExternalSource() {
        return this.linkedExternalSource;
    }
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetEventSourcesSource defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String eventSourceName;
        private Map<String,Object> externalSourceConfig;
        private String externalSourceType;
        private String id;
        private Boolean linkedExternalSource;
        private String type;
        public Builder() {}
        public Builder(GetEventSourcesSource defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.eventSourceName = defaults.eventSourceName;
    	      this.externalSourceConfig = defaults.externalSourceConfig;
    	      this.externalSourceType = defaults.externalSourceType;
    	      this.id = defaults.id;
    	      this.linkedExternalSource = defaults.linkedExternalSource;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder eventSourceName(String eventSourceName) {
            this.eventSourceName = Objects.requireNonNull(eventSourceName);
            return this;
        }
        @CustomType.Setter
        public Builder externalSourceConfig(Map<String,Object> externalSourceConfig) {
            this.externalSourceConfig = Objects.requireNonNull(externalSourceConfig);
            return this;
        }
        @CustomType.Setter
        public Builder externalSourceType(String externalSourceType) {
            this.externalSourceType = Objects.requireNonNull(externalSourceType);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder linkedExternalSource(Boolean linkedExternalSource) {
            this.linkedExternalSource = Objects.requireNonNull(linkedExternalSource);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public GetEventSourcesSource build() {
            final var o = new GetEventSourcesSource();
            o.description = description;
            o.eventSourceName = eventSourceName;
            o.externalSourceConfig = externalSourceConfig;
            o.externalSourceType = externalSourceType;
            o.id = id;
            o.linkedExternalSource = linkedExternalSource;
            o.type = type;
            return o;
        }
    }
}