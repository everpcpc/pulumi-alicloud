// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dns.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AddressPoolAddressArgs extends com.pulumi.resources.ResourceArgs {

    public static final AddressPoolAddressArgs Empty = new AddressPoolAddressArgs();

    /**
     * The address lists of the Address Pool. See the following `Block address`.
     * 
     */
    @Import(name="address", required=true)
    private Output<String> address;

    /**
     * @return The address lists of the Address Pool. See the following `Block address`.
     * 
     */
    public Output<String> address() {
        return this.address;
    }

    /**
     * The source region of the address. expressed as a JSON string. The structure is as follows:
     * 
     */
    @Import(name="attributeInfo", required=true)
    private Output<String> attributeInfo;

    /**
     * @return The source region of the address. expressed as a JSON string. The structure is as follows:
     * 
     */
    public Output<String> attributeInfo() {
        return this.attributeInfo;
    }

    /**
     * The weight of the address. **NOTE:** The attribute is valid when the attribute `lba_strategy` is `RATIO`.
     * 
     */
    @Import(name="lbaWeight")
    private @Nullable Output<Integer> lbaWeight;

    /**
     * @return The weight of the address. **NOTE:** The attribute is valid when the attribute `lba_strategy` is `RATIO`.
     * 
     */
    public Optional<Output<Integer>> lbaWeight() {
        return Optional.ofNullable(this.lbaWeight);
    }

    /**
     * The type of the address. Valid values:`SMART`, `ONLINE` and `OFFLINE`.
     * 
     */
    @Import(name="mode", required=true)
    private Output<String> mode;

    /**
     * @return The type of the address. Valid values:`SMART`, `ONLINE` and `OFFLINE`.
     * 
     */
    public Output<String> mode() {
        return this.mode;
    }

    /**
     * The description of the address.
     * 
     */
    @Import(name="remark")
    private @Nullable Output<String> remark;

    /**
     * @return The description of the address.
     * 
     */
    public Optional<Output<String>> remark() {
        return Optional.ofNullable(this.remark);
    }

    private AddressPoolAddressArgs() {}

    private AddressPoolAddressArgs(AddressPoolAddressArgs $) {
        this.address = $.address;
        this.attributeInfo = $.attributeInfo;
        this.lbaWeight = $.lbaWeight;
        this.mode = $.mode;
        this.remark = $.remark;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AddressPoolAddressArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AddressPoolAddressArgs $;

        public Builder() {
            $ = new AddressPoolAddressArgs();
        }

        public Builder(AddressPoolAddressArgs defaults) {
            $ = new AddressPoolAddressArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param address The address lists of the Address Pool. See the following `Block address`.
         * 
         * @return builder
         * 
         */
        public Builder address(Output<String> address) {
            $.address = address;
            return this;
        }

        /**
         * @param address The address lists of the Address Pool. See the following `Block address`.
         * 
         * @return builder
         * 
         */
        public Builder address(String address) {
            return address(Output.of(address));
        }

        /**
         * @param attributeInfo The source region of the address. expressed as a JSON string. The structure is as follows:
         * 
         * @return builder
         * 
         */
        public Builder attributeInfo(Output<String> attributeInfo) {
            $.attributeInfo = attributeInfo;
            return this;
        }

        /**
         * @param attributeInfo The source region of the address. expressed as a JSON string. The structure is as follows:
         * 
         * @return builder
         * 
         */
        public Builder attributeInfo(String attributeInfo) {
            return attributeInfo(Output.of(attributeInfo));
        }

        /**
         * @param lbaWeight The weight of the address. **NOTE:** The attribute is valid when the attribute `lba_strategy` is `RATIO`.
         * 
         * @return builder
         * 
         */
        public Builder lbaWeight(@Nullable Output<Integer> lbaWeight) {
            $.lbaWeight = lbaWeight;
            return this;
        }

        /**
         * @param lbaWeight The weight of the address. **NOTE:** The attribute is valid when the attribute `lba_strategy` is `RATIO`.
         * 
         * @return builder
         * 
         */
        public Builder lbaWeight(Integer lbaWeight) {
            return lbaWeight(Output.of(lbaWeight));
        }

        /**
         * @param mode The type of the address. Valid values:`SMART`, `ONLINE` and `OFFLINE`.
         * 
         * @return builder
         * 
         */
        public Builder mode(Output<String> mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param mode The type of the address. Valid values:`SMART`, `ONLINE` and `OFFLINE`.
         * 
         * @return builder
         * 
         */
        public Builder mode(String mode) {
            return mode(Output.of(mode));
        }

        /**
         * @param remark The description of the address.
         * 
         * @return builder
         * 
         */
        public Builder remark(@Nullable Output<String> remark) {
            $.remark = remark;
            return this;
        }

        /**
         * @param remark The description of the address.
         * 
         * @return builder
         * 
         */
        public Builder remark(String remark) {
            return remark(Output.of(remark));
        }

        public AddressPoolAddressArgs build() {
            $.address = Objects.requireNonNull($.address, "expected parameter 'address' to be non-null");
            $.attributeInfo = Objects.requireNonNull($.attributeInfo, "expected parameter 'attributeInfo' to be non-null");
            $.mode = Objects.requireNonNull($.mode, "expected parameter 'mode' to be non-null");
            return $;
        }
    }

}