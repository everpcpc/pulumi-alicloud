# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import importlib
# Make subpackages available:
__all__ = ['actiontrail', 'adb', 'alikafka', 'apigateway', 'cas', 'cassandra', 'cdn', 'cen', 'cloudconnect', 'cms', 'config', 'cr', 'cs', 'datahub', 'ddos', 'dds', 'dms', 'dns', 'drds', 'eci', 'ecs', 'edas', 'elasticsearch', 'emr', 'ess', 'fc', 'gpdb', 'hbase', 'kms', 'kvstore', 'log', 'marketplace', 'maxcompute', 'mns', 'mongodb', 'nas', 'oss', 'ots', 'polardb', 'pvtz', 'ram', 'rds', 'resourcemanager', 'rocketmq', 'sag', 'slb', 'vpc', 'vpn', 'waf', 'yundun']
for pkg in __all__:
    if pkg != 'config':
        importlib.import_module(f'{__name__}.{pkg}')

# Export this package's modules as members:
from .get_account import *
from .get_caller_identity import *
from .get_file_crc64_checksum import *
from .get_regions import *
from .get_zones import *
from .provider import *
