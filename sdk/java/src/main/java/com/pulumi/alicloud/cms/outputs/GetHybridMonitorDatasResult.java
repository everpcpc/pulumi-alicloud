// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms.outputs;

import com.pulumi.alicloud.cms.outputs.GetHybridMonitorDatasData;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetHybridMonitorDatasResult {
    private List<GetHybridMonitorDatasData> datas;
    private String end;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String namespace;
    private @Nullable String outputFile;
    private @Nullable String period;
    private String promSql;
    private String start;

    private GetHybridMonitorDatasResult() {}
    public List<GetHybridMonitorDatasData> datas() {
        return this.datas;
    }
    public String end() {
        return this.end;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String namespace() {
        return this.namespace;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public Optional<String> period() {
        return Optional.ofNullable(this.period);
    }
    public String promSql() {
        return this.promSql;
    }
    public String start() {
        return this.start;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetHybridMonitorDatasResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetHybridMonitorDatasData> datas;
        private String end;
        private String id;
        private String namespace;
        private @Nullable String outputFile;
        private @Nullable String period;
        private String promSql;
        private String start;
        public Builder() {}
        public Builder(GetHybridMonitorDatasResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.datas = defaults.datas;
    	      this.end = defaults.end;
    	      this.id = defaults.id;
    	      this.namespace = defaults.namespace;
    	      this.outputFile = defaults.outputFile;
    	      this.period = defaults.period;
    	      this.promSql = defaults.promSql;
    	      this.start = defaults.start;
        }

        @CustomType.Setter
        public Builder datas(List<GetHybridMonitorDatasData> datas) {
            this.datas = Objects.requireNonNull(datas);
            return this;
        }
        public Builder datas(GetHybridMonitorDatasData... datas) {
            return datas(List.of(datas));
        }
        @CustomType.Setter
        public Builder end(String end) {
            this.end = Objects.requireNonNull(end);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder namespace(String namespace) {
            this.namespace = Objects.requireNonNull(namespace);
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder period(@Nullable String period) {
            this.period = period;
            return this;
        }
        @CustomType.Setter
        public Builder promSql(String promSql) {
            this.promSql = Objects.requireNonNull(promSql);
            return this;
        }
        @CustomType.Setter
        public Builder start(String start) {
            this.start = Objects.requireNonNull(start);
            return this;
        }
        public GetHybridMonitorDatasResult build() {
            final var o = new GetHybridMonitorDatasResult();
            o.datas = datas;
            o.end = end;
            o.id = id;
            o.namespace = namespace;
            o.outputFile = outputFile;
            o.period = period;
            o.promSql = promSql;
            o.start = start;
            return o;
        }
    }
}
