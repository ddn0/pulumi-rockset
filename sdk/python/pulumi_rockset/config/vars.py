# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('rockset')


class _ExportableConfig(types.ModuleType):
    @property
    def api_key(self) -> Optional[str]:
        """
        The API key used to access Rockset
        """
        return __config__.get('apiKey')

    @property
    def api_server(self) -> Optional[str]:
        """
        The API server for accessing Rockset
        """
        return __config__.get('apiServer')

