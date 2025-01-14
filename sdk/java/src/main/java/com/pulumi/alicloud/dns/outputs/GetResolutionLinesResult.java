// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dns.outputs;

import com.pulumi.alicloud.dns.outputs.GetResolutionLinesLine;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetResolutionLinesResult {
    private @Nullable String domainName;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String lang;
    /**
     * @return Line code.
     * 
     */
    private List<String> lineCodes;
    /**
     * @return A list of line display names.
     * 
     */
    private List<String> lineDisplayNames;
    private @Nullable List<String> lineNames;
    /**
     * @return A list of cloud resolution line. Each element contains the following attributes:
     * 
     */
    private List<GetResolutionLinesLine> lines;
    private @Nullable String outputFile;
    private @Nullable String userClientIp;

    private GetResolutionLinesResult() {}
    public Optional<String> domainName() {
        return Optional.ofNullable(this.domainName);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> lang() {
        return Optional.ofNullable(this.lang);
    }
    /**
     * @return Line code.
     * 
     */
    public List<String> lineCodes() {
        return this.lineCodes;
    }
    /**
     * @return A list of line display names.
     * 
     */
    public List<String> lineDisplayNames() {
        return this.lineDisplayNames;
    }
    public List<String> lineNames() {
        return this.lineNames == null ? List.of() : this.lineNames;
    }
    /**
     * @return A list of cloud resolution line. Each element contains the following attributes:
     * 
     */
    public List<GetResolutionLinesLine> lines() {
        return this.lines;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public Optional<String> userClientIp() {
        return Optional.ofNullable(this.userClientIp);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetResolutionLinesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String domainName;
        private String id;
        private @Nullable String lang;
        private List<String> lineCodes;
        private List<String> lineDisplayNames;
        private @Nullable List<String> lineNames;
        private List<GetResolutionLinesLine> lines;
        private @Nullable String outputFile;
        private @Nullable String userClientIp;
        public Builder() {}
        public Builder(GetResolutionLinesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.domainName = defaults.domainName;
    	      this.id = defaults.id;
    	      this.lang = defaults.lang;
    	      this.lineCodes = defaults.lineCodes;
    	      this.lineDisplayNames = defaults.lineDisplayNames;
    	      this.lineNames = defaults.lineNames;
    	      this.lines = defaults.lines;
    	      this.outputFile = defaults.outputFile;
    	      this.userClientIp = defaults.userClientIp;
        }

        @CustomType.Setter
        public Builder domainName(@Nullable String domainName) {
            this.domainName = domainName;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder lang(@Nullable String lang) {
            this.lang = lang;
            return this;
        }
        @CustomType.Setter
        public Builder lineCodes(List<String> lineCodes) {
            this.lineCodes = Objects.requireNonNull(lineCodes);
            return this;
        }
        public Builder lineCodes(String... lineCodes) {
            return lineCodes(List.of(lineCodes));
        }
        @CustomType.Setter
        public Builder lineDisplayNames(List<String> lineDisplayNames) {
            this.lineDisplayNames = Objects.requireNonNull(lineDisplayNames);
            return this;
        }
        public Builder lineDisplayNames(String... lineDisplayNames) {
            return lineDisplayNames(List.of(lineDisplayNames));
        }
        @CustomType.Setter
        public Builder lineNames(@Nullable List<String> lineNames) {
            this.lineNames = lineNames;
            return this;
        }
        public Builder lineNames(String... lineNames) {
            return lineNames(List.of(lineNames));
        }
        @CustomType.Setter
        public Builder lines(List<GetResolutionLinesLine> lines) {
            this.lines = Objects.requireNonNull(lines);
            return this;
        }
        public Builder lines(GetResolutionLinesLine... lines) {
            return lines(List.of(lines));
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder userClientIp(@Nullable String userClientIp) {
            this.userClientIp = userClientIp;
            return this;
        }
        public GetResolutionLinesResult build() {
            final var o = new GetResolutionLinesResult();
            o.domainName = domainName;
            o.id = id;
            o.lang = lang;
            o.lineCodes = lineCodes;
            o.lineDisplayNames = lineDisplayNames;
            o.lineNames = lineNames;
            o.lines = lines;
            o.outputFile = outputFile;
            o.userClientIp = userClientIp;
            return o;
        }
    }
}
