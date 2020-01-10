# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import importlib
# Make subpackages available:
__all__ = ['actiontrail', 'alikafka', 'apigateway', 'cas', 'cdn', 'cen', 'cloudconnect', 'cms', 'config', 'cr', 'cs', 'datahub', 'ddos', 'dds', 'dns', 'drds', 'ecs', 'elasticsearch', 'emr', 'ess', 'fc', 'gpdb', 'kms', 'kvstore', 'log', 'marketplace', 'mns', 'mongodb', 'nas', 'oss', 'ots', 'polardb', 'pvtz', 'ram', 'rds', 'rocketmq', 'sag', 'slb', 'vpc', 'vpn', 'yundun']
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
