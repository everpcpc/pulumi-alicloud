// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.pvtz.outputs;

import com.pulumi.alicloud.pvtz.outputs.GetZoneRecordsRecord;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetZoneRecordsResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of Private Zone Record IDs.
     * 
     */
    private List<String> ids;
    private @Nullable String keyword;
    private @Nullable String lang;
    private @Nullable String outputFile;
    /**
     * @return A list of zone records. Each element contains the following attributes:
     * 
     */
    private List<GetZoneRecordsRecord> records;
    private @Nullable String searchMode;
    /**
     * @return Status of the Private Zone Record.
     * 
     */
    private @Nullable String status;
    private @Nullable String tag;
    private @Nullable String userClientIp;
    private String zoneId;

    private GetZoneRecordsResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list of Private Zone Record IDs.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> keyword() {
        return Optional.ofNullable(this.keyword);
    }
    public Optional<String> lang() {
        return Optional.ofNullable(this.lang);
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return A list of zone records. Each element contains the following attributes:
     * 
     */
    public List<GetZoneRecordsRecord> records() {
        return this.records;
    }
    public Optional<String> searchMode() {
        return Optional.ofNullable(this.searchMode);
    }
    /**
     * @return Status of the Private Zone Record.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    public Optional<String> tag() {
        return Optional.ofNullable(this.tag);
    }
    public Optional<String> userClientIp() {
        return Optional.ofNullable(this.userClientIp);
    }
    public String zoneId() {
        return this.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetZoneRecordsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> ids;
        private @Nullable String keyword;
        private @Nullable String lang;
        private @Nullable String outputFile;
        private List<GetZoneRecordsRecord> records;
        private @Nullable String searchMode;
        private @Nullable String status;
        private @Nullable String tag;
        private @Nullable String userClientIp;
        private String zoneId;
        public Builder() {}
        public Builder(GetZoneRecordsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.keyword = defaults.keyword;
    	      this.lang = defaults.lang;
    	      this.outputFile = defaults.outputFile;
    	      this.records = defaults.records;
    	      this.searchMode = defaults.searchMode;
    	      this.status = defaults.status;
    	      this.tag = defaults.tag;
    	      this.userClientIp = defaults.userClientIp;
    	      this.zoneId = defaults.zoneId;
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
        public Builder keyword(@Nullable String keyword) {
            this.keyword = keyword;
            return this;
        }
        @CustomType.Setter
        public Builder lang(@Nullable String lang) {
            this.lang = lang;
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder records(List<GetZoneRecordsRecord> records) {
            this.records = Objects.requireNonNull(records);
            return this;
        }
        public Builder records(GetZoneRecordsRecord... records) {
            return records(List.of(records));
        }
        @CustomType.Setter
        public Builder searchMode(@Nullable String searchMode) {
            this.searchMode = searchMode;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder tag(@Nullable String tag) {
            this.tag = tag;
            return this;
        }
        @CustomType.Setter
        public Builder userClientIp(@Nullable String userClientIp) {
            this.userClientIp = userClientIp;
            return this;
        }
        @CustomType.Setter
        public Builder zoneId(String zoneId) {
            this.zoneId = Objects.requireNonNull(zoneId);
            return this;
        }
        public GetZoneRecordsResult build() {
            final var o = new GetZoneRecordsResult();
            o.id = id;
            o.ids = ids;
            o.keyword = keyword;
            o.lang = lang;
            o.outputFile = outputFile;
            o.records = records;
            o.searchMode = searchMode;
            o.status = status;
            o.tag = tag;
            o.userClientIp = userClientIp;
            o.zoneId = zoneId;
            return o;
        }
    }
}
