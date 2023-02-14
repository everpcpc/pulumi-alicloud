// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr.outputs;

import com.pulumi.alicloud.hbr.outputs.GetHanaBackupClientsHanaBackupClient;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetHanaBackupClientsResult {
    /**
     * @return The ID of the backup client.
     * 
     */
    private @Nullable String clientId;
    /**
     * @return The ID of the SAP HANA instance.
     * 
     */
    private @Nullable String clusterId;
    /**
     * @return A list of Hana Backup Clients. Each element contains the following attributes:
     * 
     */
    private List<GetHanaBackupClientsHanaBackupClient> hanaBackupClients;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String outputFile;
    private @Nullable Integer pageNumber;
    private @Nullable Integer pageSize;
    /**
     * @return The status of the backup client.
     * 
     */
    private @Nullable String status;
    /**
     * @return The ID of the backup vault.
     * 
     */
    private String vaultId;

    private GetHanaBackupClientsResult() {}
    /**
     * @return The ID of the backup client.
     * 
     */
    public Optional<String> clientId() {
        return Optional.ofNullable(this.clientId);
    }
    /**
     * @return The ID of the SAP HANA instance.
     * 
     */
    public Optional<String> clusterId() {
        return Optional.ofNullable(this.clusterId);
    }
    /**
     * @return A list of Hana Backup Clients. Each element contains the following attributes:
     * 
     */
    public List<GetHanaBackupClientsHanaBackupClient> hanaBackupClients() {
        return this.hanaBackupClients;
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
    public Optional<Integer> pageNumber() {
        return Optional.ofNullable(this.pageNumber);
    }
    public Optional<Integer> pageSize() {
        return Optional.ofNullable(this.pageSize);
    }
    /**
     * @return The status of the backup client.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    /**
     * @return The ID of the backup vault.
     * 
     */
    public String vaultId() {
        return this.vaultId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetHanaBackupClientsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String clientId;
        private @Nullable String clusterId;
        private List<GetHanaBackupClientsHanaBackupClient> hanaBackupClients;
        private String id;
        private List<String> ids;
        private @Nullable String outputFile;
        private @Nullable Integer pageNumber;
        private @Nullable Integer pageSize;
        private @Nullable String status;
        private String vaultId;
        public Builder() {}
        public Builder(GetHanaBackupClientsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clientId = defaults.clientId;
    	      this.clusterId = defaults.clusterId;
    	      this.hanaBackupClients = defaults.hanaBackupClients;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.outputFile = defaults.outputFile;
    	      this.pageNumber = defaults.pageNumber;
    	      this.pageSize = defaults.pageSize;
    	      this.status = defaults.status;
    	      this.vaultId = defaults.vaultId;
        }

        @CustomType.Setter
        public Builder clientId(@Nullable String clientId) {
            this.clientId = clientId;
            return this;
        }
        @CustomType.Setter
        public Builder clusterId(@Nullable String clusterId) {
            this.clusterId = clusterId;
            return this;
        }
        @CustomType.Setter
        public Builder hanaBackupClients(List<GetHanaBackupClientsHanaBackupClient> hanaBackupClients) {
            this.hanaBackupClients = Objects.requireNonNull(hanaBackupClients);
            return this;
        }
        public Builder hanaBackupClients(GetHanaBackupClientsHanaBackupClient... hanaBackupClients) {
            return hanaBackupClients(List.of(hanaBackupClients));
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
        public Builder pageNumber(@Nullable Integer pageNumber) {
            this.pageNumber = pageNumber;
            return this;
        }
        @CustomType.Setter
        public Builder pageSize(@Nullable Integer pageSize) {
            this.pageSize = pageSize;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder vaultId(String vaultId) {
            this.vaultId = Objects.requireNonNull(vaultId);
            return this;
        }
        public GetHanaBackupClientsResult build() {
            final var o = new GetHanaBackupClientsResult();
            o.clientId = clientId;
            o.clusterId = clusterId;
            o.hanaBackupClients = hanaBackupClients;
            o.id = id;
            o.ids = ids;
            o.outputFile = outputFile;
            o.pageNumber = pageNumber;
            o.pageSize = pageSize;
            o.status = status;
            o.vaultId = vaultId;
            return o;
        }
    }
}
