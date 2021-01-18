# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .application import *
from .cluster import *
from .edge_kubernetes import *
from .get_ack_service import *
from .get_edge_kubernetes_clusters import *
from .get_kubernetes_clusters import *
from .get_managed_kubernetes_clusters import *
from .get_registry_enterprise_instances import *
from .get_registry_enterprise_namespaces import *
from .get_registry_enterprise_repos import *
from .get_registry_enterprise_sync_rules import *
from .get_serverless_kubernetes_clusters import *
from .kubernetes import *
from .kubernetes_autoscaler import *
from .managed_kubernetes import *
from .node_pool import *
from .registry_enterprise_namespace import *
from .registry_enterprise_repo import *
from .registry_enterprise_sync_rule import *
from .serverless_kubernetes import *
from .swarm import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "alicloud:cs/application:Application":
                return Application(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:cs/cluster:Cluster":
                return Cluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:cs/edgeKubernetes:EdgeKubernetes":
                return EdgeKubernetes(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:cs/kubernetes:Kubernetes":
                return Kubernetes(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:cs/kubernetesAutoscaler:KubernetesAutoscaler":
                return KubernetesAutoscaler(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:cs/managedKubernetes:ManagedKubernetes":
                return ManagedKubernetes(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:cs/nodePool:NodePool":
                return NodePool(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:cs/registryEnterpriseNamespace:RegistryEnterpriseNamespace":
                return RegistryEnterpriseNamespace(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:cs/registryEnterpriseRepo:RegistryEnterpriseRepo":
                return RegistryEnterpriseRepo(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:cs/registryEnterpriseSyncRule:RegistryEnterpriseSyncRule":
                return RegistryEnterpriseSyncRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:cs/serverlessKubernetes:ServerlessKubernetes":
                return ServerlessKubernetes(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:cs/swarm:Swarm":
                return Swarm(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("alicloud", "cs/application", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "cs/cluster", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "cs/edgeKubernetes", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "cs/kubernetes", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "cs/kubernetesAutoscaler", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "cs/managedKubernetes", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "cs/nodePool", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "cs/registryEnterpriseNamespace", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "cs/registryEnterpriseRepo", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "cs/registryEnterpriseSyncRule", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "cs/serverlessKubernetes", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "cs/swarm", _module_instance)

_register_module()
