// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Wafv3.Outputs
{

    [OutputType]
    public sealed class DomainListen
    {
        /// <summary>
        /// The ID of the certificate to be added. This parameter is used only if the value of **https_ports** is not empty (indicating that the domain name uses the HTTPS protocol).
        /// </summary>
        public readonly string? CertId;
        /// <summary>
        /// The type of encryption suite to add. This parameter is used only if the value of **https_ports** is not empty (indicating that the domain name uses the HTTPS protocol). Value:
        /// - **1**: indicates that all encryption suites are added.
        /// - **2**: indicates that a strong encryption package is added. You can select this value only if the value of **tls_version** is `tlsv1.2`.
        /// - **99**: indicates that a custom encryption suite is added.
        /// </summary>
        public readonly int? CipherSuite;
        /// <summary>
        /// The specific custom encryption suite to add.
        /// </summary>
        public readonly ImmutableArray<string> CustomCiphers;
        /// <summary>
        /// Whether TSL1.3 version is supported. This parameter is used only if the value of **https_ports** is not empty (indicating that the domain name uses the HTTPS protocol). Value:
        /// - **true**: indicates that TSL1.3 is supported.
        /// - **false**: indicates that TSL1.3 is not supported.
        /// </summary>
        public readonly bool? EnableTlsv3;
        /// <summary>
        /// Whether to enable exclusive IP address. This parameter is used only when the value of **ipv6_enabled** is **false** (indicating that IPv6 is not enabled) and the value of **protection_resource** is **share** (indicating that a shared cluster is used). Value:
        /// - **true**: indicates that the exclusive IP address is enabled.
        /// - **false** (default): indicates that exclusive IP address is not enabled.
        /// </summary>
        public readonly bool? ExclusiveIp;
        /// <summary>
        /// Whether to enable the forced jump of HTTPS. This parameter is used only when the value of `https_ports` is not empty (indicating that the domain name uses HTTPS protocol) and the value of httports is empty (indicating that the domain name does not use HTTP protocol). Value:
        /// - **true**: indicates that HTTPS forced redirection is enabled.
        /// - **false**: indicates that HTTPS forced redirection is not enabled.
        /// </summary>
        public readonly bool? FocusHttps;
        /// <summary>
        /// Whether to turn on http2. This parameter is used only if the value of **https_ports** is not empty (indicating that the domain name uses the HTTPS protocol). Value:
        /// - **true:** indicates that HTTP2 is enabled.
        /// - **false** (default): indicates that HTTP2 is not enabled.
        /// </summary>
        public readonly bool? Http2Enabled;
        /// <summary>
        /// The listening port of the HTTP protocol.
        /// </summary>
        public readonly ImmutableArray<int> HttpPorts;
        /// <summary>
        /// The listening port of the HTTPS protocol.
        /// </summary>
        public readonly ImmutableArray<int> HttpsPorts;
        /// <summary>
        /// Whether IPv6 is turned on. Value:
        /// - **true**: indicates that IPv6 is enabled.
        /// - **false** (default): indicates that IPv6 is not enabled.
        /// </summary>
        public readonly bool? Ipv6Enabled;
        /// <summary>
        /// The type of protection resource to use. Value:
        /// - **share** (default): indicates that a shared cluster is used.
        /// - **gslb**: indicates that the shared cluster intelligent load balancing is used.
        /// </summary>
        public readonly string? ProtectionResource;
        /// <summary>
        /// The version of TLS to add. This parameter is used only if the value of **https_ports** is not empty (indicating that the domain name uses the HTTPS protocol). Value: **tlsv1**, **tlsv1.1**, **tlsv1.2**.
        /// </summary>
        public readonly string? TlsVersion;
        /// <summary>
        /// WAF obtains the real IP address of the client. Value:
        /// - **0** (default): indicates that the client has not forwarded the traffic to WAF through other layer -7 agents.
        /// - **1**: indicates that the first value of the X-Forwarded-For(XFF) field in the WAF read request header is used as the client IP address.
        /// - **2**: indicates that the custom field value set by you in the WAF read request header is used as the client IP address.
        /// </summary>
        public readonly int? XffHeaderMode;
        /// <summary>
        /// Set the list of custom fields used to obtain the client IP address.
        /// </summary>
        public readonly ImmutableArray<string> XffHeaders;

        [OutputConstructor]
        private DomainListen(
            string? certId,

            int? cipherSuite,

            ImmutableArray<string> customCiphers,

            bool? enableTlsv3,

            bool? exclusiveIp,

            bool? focusHttps,

            bool? http2Enabled,

            ImmutableArray<int> httpPorts,

            ImmutableArray<int> httpsPorts,

            bool? ipv6Enabled,

            string? protectionResource,

            string? tlsVersion,

            int? xffHeaderMode,

            ImmutableArray<string> xffHeaders)
        {
            CertId = certId;
            CipherSuite = cipherSuite;
            CustomCiphers = customCiphers;
            EnableTlsv3 = enableTlsv3;
            ExclusiveIp = exclusiveIp;
            FocusHttps = focusHttps;
            Http2Enabled = http2Enabled;
            HttpPorts = httpPorts;
            HttpsPorts = httpsPorts;
            Ipv6Enabled = ipv6Enabled;
            ProtectionResource = protectionResource;
            TlsVersion = tlsVersion;
            XffHeaderMode = xffHeaderMode;
            XffHeaders = xffHeaders;
        }
    }
}
