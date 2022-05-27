# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ServiceMeshLoadBalancerArgs',
    'ServiceMeshMeshConfigArgs',
    'ServiceMeshMeshConfigAccessLogArgs',
    'ServiceMeshMeshConfigAuditArgs',
    'ServiceMeshMeshConfigKialiArgs',
    'ServiceMeshMeshConfigOpaArgs',
    'ServiceMeshMeshConfigPilotArgs',
    'ServiceMeshMeshConfigProxyArgs',
    'ServiceMeshMeshConfigSidecarInjectorArgs',
    'ServiceMeshNetworkArgs',
]

@pulumi.input_type
class ServiceMeshLoadBalancerArgs:
    def __init__(__self__, *,
                 api_server_loadbalancer_id: Optional[pulumi.Input[str]] = None,
                 api_server_public_eip: Optional[pulumi.Input[bool]] = None,
                 pilot_public_eip: Optional[pulumi.Input[bool]] = None,
                 pilot_public_loadbalancer_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] api_server_public_eip: Whether to use the IP address of a public network exposed the API Server.
        :param pulumi.Input[bool] pilot_public_eip: Whether to use the IP address of a public network exposure the Istio Pilot.
        """
        if api_server_loadbalancer_id is not None:
            pulumi.set(__self__, "api_server_loadbalancer_id", api_server_loadbalancer_id)
        if api_server_public_eip is not None:
            pulumi.set(__self__, "api_server_public_eip", api_server_public_eip)
        if pilot_public_eip is not None:
            pulumi.set(__self__, "pilot_public_eip", pilot_public_eip)
        if pilot_public_loadbalancer_id is not None:
            pulumi.set(__self__, "pilot_public_loadbalancer_id", pilot_public_loadbalancer_id)

    @property
    @pulumi.getter(name="apiServerLoadbalancerId")
    def api_server_loadbalancer_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "api_server_loadbalancer_id")

    @api_server_loadbalancer_id.setter
    def api_server_loadbalancer_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_server_loadbalancer_id", value)

    @property
    @pulumi.getter(name="apiServerPublicEip")
    def api_server_public_eip(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to use the IP address of a public network exposed the API Server.
        """
        return pulumi.get(self, "api_server_public_eip")

    @api_server_public_eip.setter
    def api_server_public_eip(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "api_server_public_eip", value)

    @property
    @pulumi.getter(name="pilotPublicEip")
    def pilot_public_eip(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to use the IP address of a public network exposure the Istio Pilot.
        """
        return pulumi.get(self, "pilot_public_eip")

    @pilot_public_eip.setter
    def pilot_public_eip(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "pilot_public_eip", value)

    @property
    @pulumi.getter(name="pilotPublicLoadbalancerId")
    def pilot_public_loadbalancer_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "pilot_public_loadbalancer_id")

    @pilot_public_loadbalancer_id.setter
    def pilot_public_loadbalancer_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pilot_public_loadbalancer_id", value)


@pulumi.input_type
class ServiceMeshMeshConfigArgs:
    def __init__(__self__, *,
                 access_log: Optional[pulumi.Input['ServiceMeshMeshConfigAccessLogArgs']] = None,
                 audit: Optional[pulumi.Input['ServiceMeshMeshConfigAuditArgs']] = None,
                 customized_zipkin: Optional[pulumi.Input[bool]] = None,
                 enable_locality_lb: Optional[pulumi.Input[bool]] = None,
                 kiali: Optional[pulumi.Input['ServiceMeshMeshConfigKialiArgs']] = None,
                 opa: Optional[pulumi.Input['ServiceMeshMeshConfigOpaArgs']] = None,
                 outbound_traffic_policy: Optional[pulumi.Input[str]] = None,
                 pilot: Optional[pulumi.Input['ServiceMeshMeshConfigPilotArgs']] = None,
                 proxy: Optional[pulumi.Input['ServiceMeshMeshConfigProxyArgs']] = None,
                 sidecar_injector: Optional[pulumi.Input['ServiceMeshMeshConfigSidecarInjectorArgs']] = None,
                 telemetry: Optional[pulumi.Input[bool]] = None,
                 tracing: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input['ServiceMeshMeshConfigAccessLogArgs'] access_log: The configuration of the access logging.
        :param pulumi.Input['ServiceMeshMeshConfigAuditArgs'] audit: The configuration of the audit. See the following `Block audit`.
        :param pulumi.Input[bool] customized_zipkin: Whether to enable the use of a custom zipkin.
        :param pulumi.Input[bool] enable_locality_lb: The enable locality lb.
        :param pulumi.Input['ServiceMeshMeshConfigKialiArgs'] kiali: The configuration of the Kiali. See the following `Block kiali`.
        :param pulumi.Input['ServiceMeshMeshConfigOpaArgs'] opa: The open-door policy of agent (OPA) plug-in information. See the following `Block opa`.
        :param pulumi.Input[str] outbound_traffic_policy: The policy of the Out to the traffic. Valid values: `ALLOW_ANY` and `REGISTRY_ONLY`.
        :param pulumi.Input['ServiceMeshMeshConfigPilotArgs'] pilot: The configuration of the Link trace sampling. See the following `Block pilot`.
        :param pulumi.Input['ServiceMeshMeshConfigProxyArgs'] proxy: The configuration of the Proxy. See the following `Block proxy`.
        :param pulumi.Input['ServiceMeshMeshConfigSidecarInjectorArgs'] sidecar_injector: The configuration of the Sidecar injector. See the following `Block sidecar_injector`.
        :param pulumi.Input[bool] telemetry: Whether to enable acquisition Prometheus metrics (it is recommended that you use [Alibaba Cloud Prometheus monitoring](https://arms.console.aliyun.com/).
        :param pulumi.Input[bool] tracing: Whether to enable link trace (you need to have [Alibaba Cloud link tracking service](https://tracing-analysis.console.aliyun.com/).
        """
        if access_log is not None:
            pulumi.set(__self__, "access_log", access_log)
        if audit is not None:
            pulumi.set(__self__, "audit", audit)
        if customized_zipkin is not None:
            pulumi.set(__self__, "customized_zipkin", customized_zipkin)
        if enable_locality_lb is not None:
            pulumi.set(__self__, "enable_locality_lb", enable_locality_lb)
        if kiali is not None:
            pulumi.set(__self__, "kiali", kiali)
        if opa is not None:
            pulumi.set(__self__, "opa", opa)
        if outbound_traffic_policy is not None:
            pulumi.set(__self__, "outbound_traffic_policy", outbound_traffic_policy)
        if pilot is not None:
            pulumi.set(__self__, "pilot", pilot)
        if proxy is not None:
            pulumi.set(__self__, "proxy", proxy)
        if sidecar_injector is not None:
            pulumi.set(__self__, "sidecar_injector", sidecar_injector)
        if telemetry is not None:
            pulumi.set(__self__, "telemetry", telemetry)
        if tracing is not None:
            pulumi.set(__self__, "tracing", tracing)

    @property
    @pulumi.getter(name="accessLog")
    def access_log(self) -> Optional[pulumi.Input['ServiceMeshMeshConfigAccessLogArgs']]:
        """
        The configuration of the access logging.
        """
        return pulumi.get(self, "access_log")

    @access_log.setter
    def access_log(self, value: Optional[pulumi.Input['ServiceMeshMeshConfigAccessLogArgs']]):
        pulumi.set(self, "access_log", value)

    @property
    @pulumi.getter
    def audit(self) -> Optional[pulumi.Input['ServiceMeshMeshConfigAuditArgs']]:
        """
        The configuration of the audit. See the following `Block audit`.
        """
        return pulumi.get(self, "audit")

    @audit.setter
    def audit(self, value: Optional[pulumi.Input['ServiceMeshMeshConfigAuditArgs']]):
        pulumi.set(self, "audit", value)

    @property
    @pulumi.getter(name="customizedZipkin")
    def customized_zipkin(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable the use of a custom zipkin.
        """
        return pulumi.get(self, "customized_zipkin")

    @customized_zipkin.setter
    def customized_zipkin(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "customized_zipkin", value)

    @property
    @pulumi.getter(name="enableLocalityLb")
    def enable_locality_lb(self) -> Optional[pulumi.Input[bool]]:
        """
        The enable locality lb.
        """
        return pulumi.get(self, "enable_locality_lb")

    @enable_locality_lb.setter
    def enable_locality_lb(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_locality_lb", value)

    @property
    @pulumi.getter
    def kiali(self) -> Optional[pulumi.Input['ServiceMeshMeshConfigKialiArgs']]:
        """
        The configuration of the Kiali. See the following `Block kiali`.
        """
        return pulumi.get(self, "kiali")

    @kiali.setter
    def kiali(self, value: Optional[pulumi.Input['ServiceMeshMeshConfigKialiArgs']]):
        pulumi.set(self, "kiali", value)

    @property
    @pulumi.getter
    def opa(self) -> Optional[pulumi.Input['ServiceMeshMeshConfigOpaArgs']]:
        """
        The open-door policy of agent (OPA) plug-in information. See the following `Block opa`.
        """
        return pulumi.get(self, "opa")

    @opa.setter
    def opa(self, value: Optional[pulumi.Input['ServiceMeshMeshConfigOpaArgs']]):
        pulumi.set(self, "opa", value)

    @property
    @pulumi.getter(name="outboundTrafficPolicy")
    def outbound_traffic_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The policy of the Out to the traffic. Valid values: `ALLOW_ANY` and `REGISTRY_ONLY`.
        """
        return pulumi.get(self, "outbound_traffic_policy")

    @outbound_traffic_policy.setter
    def outbound_traffic_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "outbound_traffic_policy", value)

    @property
    @pulumi.getter
    def pilot(self) -> Optional[pulumi.Input['ServiceMeshMeshConfigPilotArgs']]:
        """
        The configuration of the Link trace sampling. See the following `Block pilot`.
        """
        return pulumi.get(self, "pilot")

    @pilot.setter
    def pilot(self, value: Optional[pulumi.Input['ServiceMeshMeshConfigPilotArgs']]):
        pulumi.set(self, "pilot", value)

    @property
    @pulumi.getter
    def proxy(self) -> Optional[pulumi.Input['ServiceMeshMeshConfigProxyArgs']]:
        """
        The configuration of the Proxy. See the following `Block proxy`.
        """
        return pulumi.get(self, "proxy")

    @proxy.setter
    def proxy(self, value: Optional[pulumi.Input['ServiceMeshMeshConfigProxyArgs']]):
        pulumi.set(self, "proxy", value)

    @property
    @pulumi.getter(name="sidecarInjector")
    def sidecar_injector(self) -> Optional[pulumi.Input['ServiceMeshMeshConfigSidecarInjectorArgs']]:
        """
        The configuration of the Sidecar injector. See the following `Block sidecar_injector`.
        """
        return pulumi.get(self, "sidecar_injector")

    @sidecar_injector.setter
    def sidecar_injector(self, value: Optional[pulumi.Input['ServiceMeshMeshConfigSidecarInjectorArgs']]):
        pulumi.set(self, "sidecar_injector", value)

    @property
    @pulumi.getter
    def telemetry(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable acquisition Prometheus metrics (it is recommended that you use [Alibaba Cloud Prometheus monitoring](https://arms.console.aliyun.com/).
        """
        return pulumi.get(self, "telemetry")

    @telemetry.setter
    def telemetry(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "telemetry", value)

    @property
    @pulumi.getter
    def tracing(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable link trace (you need to have [Alibaba Cloud link tracking service](https://tracing-analysis.console.aliyun.com/).
        """
        return pulumi.get(self, "tracing")

    @tracing.setter
    def tracing(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "tracing", value)


@pulumi.input_type
class ServiceMeshMeshConfigAccessLogArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] enabled: Whether to enable Service grid audit.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable Service grid audit.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class ServiceMeshMeshConfigAuditArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] enabled: Whether to enable Service grid audit.
        :param pulumi.Input[str] project: The Service grid audit that to the project.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable Service grid audit.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The Service grid audit that to the project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class ServiceMeshMeshConfigKialiArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] enabled: Whether to enable Service grid audit.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable Service grid audit.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class ServiceMeshMeshConfigOpaArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 limit_cpu: Optional[pulumi.Input[str]] = None,
                 limit_memory: Optional[pulumi.Input[str]] = None,
                 log_level: Optional[pulumi.Input[str]] = None,
                 request_cpu: Optional[pulumi.Input[str]] = None,
                 request_memory: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] enabled: Whether to enable Service grid audit.
        :param pulumi.Input[str] limit_cpu: The CPU resource  of the limitsOPA proxy container.
        :param pulumi.Input[str] limit_memory: The memory resource limit of the OPA proxy container.
        :param pulumi.Input[str] log_level: The log level of the OPA proxy container .
        :param pulumi.Input[str] request_cpu: The CPU resource request of the OPA proxy container.
        :param pulumi.Input[str] request_memory: The memory resource request of the OPA proxy container.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if limit_cpu is not None:
            pulumi.set(__self__, "limit_cpu", limit_cpu)
        if limit_memory is not None:
            pulumi.set(__self__, "limit_memory", limit_memory)
        if log_level is not None:
            pulumi.set(__self__, "log_level", log_level)
        if request_cpu is not None:
            pulumi.set(__self__, "request_cpu", request_cpu)
        if request_memory is not None:
            pulumi.set(__self__, "request_memory", request_memory)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable Service grid audit.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="limitCpu")
    def limit_cpu(self) -> Optional[pulumi.Input[str]]:
        """
        The CPU resource  of the limitsOPA proxy container.
        """
        return pulumi.get(self, "limit_cpu")

    @limit_cpu.setter
    def limit_cpu(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "limit_cpu", value)

    @property
    @pulumi.getter(name="limitMemory")
    def limit_memory(self) -> Optional[pulumi.Input[str]]:
        """
        The memory resource limit of the OPA proxy container.
        """
        return pulumi.get(self, "limit_memory")

    @limit_memory.setter
    def limit_memory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "limit_memory", value)

    @property
    @pulumi.getter(name="logLevel")
    def log_level(self) -> Optional[pulumi.Input[str]]:
        """
        The log level of the OPA proxy container .
        """
        return pulumi.get(self, "log_level")

    @log_level.setter
    def log_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_level", value)

    @property
    @pulumi.getter(name="requestCpu")
    def request_cpu(self) -> Optional[pulumi.Input[str]]:
        """
        The CPU resource request of the OPA proxy container.
        """
        return pulumi.get(self, "request_cpu")

    @request_cpu.setter
    def request_cpu(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_cpu", value)

    @property
    @pulumi.getter(name="requestMemory")
    def request_memory(self) -> Optional[pulumi.Input[str]]:
        """
        The memory resource request of the OPA proxy container.
        """
        return pulumi.get(self, "request_memory")

    @request_memory.setter
    def request_memory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_memory", value)


@pulumi.input_type
class ServiceMeshMeshConfigPilotArgs:
    def __init__(__self__, *,
                 http10_enabled: Optional[pulumi.Input[bool]] = None,
                 trace_sampling: Optional[pulumi.Input[float]] = None):
        """
        :param pulumi.Input[bool] http10_enabled: Whether to support the HTTP1.0.
        :param pulumi.Input[float] trace_sampling: The  percentage of the Link trace sampling.
        """
        if http10_enabled is not None:
            pulumi.set(__self__, "http10_enabled", http10_enabled)
        if trace_sampling is not None:
            pulumi.set(__self__, "trace_sampling", trace_sampling)

    @property
    @pulumi.getter(name="http10Enabled")
    def http10_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to support the HTTP1.0.
        """
        return pulumi.get(self, "http10_enabled")

    @http10_enabled.setter
    def http10_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "http10_enabled", value)

    @property
    @pulumi.getter(name="traceSampling")
    def trace_sampling(self) -> Optional[pulumi.Input[float]]:
        """
        The  percentage of the Link trace sampling.
        """
        return pulumi.get(self, "trace_sampling")

    @trace_sampling.setter
    def trace_sampling(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "trace_sampling", value)


@pulumi.input_type
class ServiceMeshMeshConfigProxyArgs:
    def __init__(__self__, *,
                 limit_cpu: Optional[pulumi.Input[str]] = None,
                 limit_memory: Optional[pulumi.Input[str]] = None,
                 request_cpu: Optional[pulumi.Input[str]] = None,
                 request_memory: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] limit_cpu: The CPU resource  of the limitsOPA proxy container.
        :param pulumi.Input[str] limit_memory: The memory resource limit of the OPA proxy container.
        :param pulumi.Input[str] request_cpu: The CPU resource request of the OPA proxy container.
        :param pulumi.Input[str] request_memory: The memory resource request of the OPA proxy container.
        """
        if limit_cpu is not None:
            pulumi.set(__self__, "limit_cpu", limit_cpu)
        if limit_memory is not None:
            pulumi.set(__self__, "limit_memory", limit_memory)
        if request_cpu is not None:
            pulumi.set(__self__, "request_cpu", request_cpu)
        if request_memory is not None:
            pulumi.set(__self__, "request_memory", request_memory)

    @property
    @pulumi.getter(name="limitCpu")
    def limit_cpu(self) -> Optional[pulumi.Input[str]]:
        """
        The CPU resource  of the limitsOPA proxy container.
        """
        return pulumi.get(self, "limit_cpu")

    @limit_cpu.setter
    def limit_cpu(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "limit_cpu", value)

    @property
    @pulumi.getter(name="limitMemory")
    def limit_memory(self) -> Optional[pulumi.Input[str]]:
        """
        The memory resource limit of the OPA proxy container.
        """
        return pulumi.get(self, "limit_memory")

    @limit_memory.setter
    def limit_memory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "limit_memory", value)

    @property
    @pulumi.getter(name="requestCpu")
    def request_cpu(self) -> Optional[pulumi.Input[str]]:
        """
        The CPU resource request of the OPA proxy container.
        """
        return pulumi.get(self, "request_cpu")

    @request_cpu.setter
    def request_cpu(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_cpu", value)

    @property
    @pulumi.getter(name="requestMemory")
    def request_memory(self) -> Optional[pulumi.Input[str]]:
        """
        The memory resource request of the OPA proxy container.
        """
        return pulumi.get(self, "request_memory")

    @request_memory.setter
    def request_memory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_memory", value)


@pulumi.input_type
class ServiceMeshMeshConfigSidecarInjectorArgs:
    def __init__(__self__, *,
                 auto_injection_policy_enabled: Optional[pulumi.Input[bool]] = None,
                 enable_namespaces_by_default: Optional[pulumi.Input[bool]] = None,
                 limit_cpu: Optional[pulumi.Input[str]] = None,
                 limit_memory: Optional[pulumi.Input[str]] = None,
                 request_cpu: Optional[pulumi.Input[str]] = None,
                 request_memory: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] auto_injection_policy_enabled: Whether to enable by Pod Annotations automatic injection Sidecar.
        :param pulumi.Input[bool] enable_namespaces_by_default: Whether it is the all namespaces you turn on the auto injection capabilities.
        :param pulumi.Input[str] limit_cpu: The CPU resource  of the limitsOPA proxy container.
        :param pulumi.Input[str] limit_memory: The memory resource limit of the OPA proxy container.
        :param pulumi.Input[str] request_cpu: The CPU resource request of the OPA proxy container.
        :param pulumi.Input[str] request_memory: The memory resource request of the OPA proxy container.
        """
        if auto_injection_policy_enabled is not None:
            pulumi.set(__self__, "auto_injection_policy_enabled", auto_injection_policy_enabled)
        if enable_namespaces_by_default is not None:
            pulumi.set(__self__, "enable_namespaces_by_default", enable_namespaces_by_default)
        if limit_cpu is not None:
            pulumi.set(__self__, "limit_cpu", limit_cpu)
        if limit_memory is not None:
            pulumi.set(__self__, "limit_memory", limit_memory)
        if request_cpu is not None:
            pulumi.set(__self__, "request_cpu", request_cpu)
        if request_memory is not None:
            pulumi.set(__self__, "request_memory", request_memory)

    @property
    @pulumi.getter(name="autoInjectionPolicyEnabled")
    def auto_injection_policy_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable by Pod Annotations automatic injection Sidecar.
        """
        return pulumi.get(self, "auto_injection_policy_enabled")

    @auto_injection_policy_enabled.setter
    def auto_injection_policy_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_injection_policy_enabled", value)

    @property
    @pulumi.getter(name="enableNamespacesByDefault")
    def enable_namespaces_by_default(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether it is the all namespaces you turn on the auto injection capabilities.
        """
        return pulumi.get(self, "enable_namespaces_by_default")

    @enable_namespaces_by_default.setter
    def enable_namespaces_by_default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_namespaces_by_default", value)

    @property
    @pulumi.getter(name="limitCpu")
    def limit_cpu(self) -> Optional[pulumi.Input[str]]:
        """
        The CPU resource  of the limitsOPA proxy container.
        """
        return pulumi.get(self, "limit_cpu")

    @limit_cpu.setter
    def limit_cpu(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "limit_cpu", value)

    @property
    @pulumi.getter(name="limitMemory")
    def limit_memory(self) -> Optional[pulumi.Input[str]]:
        """
        The memory resource limit of the OPA proxy container.
        """
        return pulumi.get(self, "limit_memory")

    @limit_memory.setter
    def limit_memory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "limit_memory", value)

    @property
    @pulumi.getter(name="requestCpu")
    def request_cpu(self) -> Optional[pulumi.Input[str]]:
        """
        The CPU resource request of the OPA proxy container.
        """
        return pulumi.get(self, "request_cpu")

    @request_cpu.setter
    def request_cpu(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_cpu", value)

    @property
    @pulumi.getter(name="requestMemory")
    def request_memory(self) -> Optional[pulumi.Input[str]]:
        """
        The memory resource request of the OPA proxy container.
        """
        return pulumi.get(self, "request_memory")

    @request_memory.setter
    def request_memory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_memory", value)


@pulumi.input_type
class ServiceMeshNetworkArgs:
    def __init__(__self__, *,
                 vpc_id: pulumi.Input[str],
                 vswitche_list: pulumi.Input[str]):
        """
        :param pulumi.Input[str] vpc_id: The ID of the VPC.
        :param pulumi.Input[str] vswitche_list: The list of Virtual Switch.
        """
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "vswitche_list", vswitche_list)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The ID of the VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vswitcheList")
    def vswitche_list(self) -> pulumi.Input[str]:
        """
        The list of Virtual Switch.
        """
        return pulumi.get(self, "vswitche_list")

    @vswitche_list.setter
    def vswitche_list(self, value: pulumi.Input[str]):
        pulumi.set(self, "vswitche_list", value)


