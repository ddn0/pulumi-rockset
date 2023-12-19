# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['GcsIntegrationArgs', 'GcsIntegration']

@pulumi.input_type
class GcsIntegrationArgs:
    def __init__(__self__, *,
                 service_account_key: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GcsIntegration resource.
        :param pulumi.Input[str] service_account_key: The GCP service account key JSON.
        :param pulumi.Input[str] description: Text describing the integration.
        :param pulumi.Input[str] name: Unique identifier for the integration. Can contain alphanumeric or dash characters.
        """
        pulumi.set(__self__, "service_account_key", service_account_key)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="serviceAccountKey")
    def service_account_key(self) -> pulumi.Input[str]:
        """
        The GCP service account key JSON.
        """
        return pulumi.get(self, "service_account_key")

    @service_account_key.setter
    def service_account_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_account_key", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Text describing the integration.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier for the integration. Can contain alphanumeric or dash characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _GcsIntegrationState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 service_account_key: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GcsIntegration resources.
        :param pulumi.Input[str] description: Text describing the integration.
        :param pulumi.Input[str] name: Unique identifier for the integration. Can contain alphanumeric or dash characters.
        :param pulumi.Input[str] service_account_key: The GCP service account key JSON.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if service_account_key is not None:
            pulumi.set(__self__, "service_account_key", service_account_key)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Text describing the integration.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier for the integration. Can contain alphanumeric or dash characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="serviceAccountKey")
    def service_account_key(self) -> Optional[pulumi.Input[str]]:
        """
        The GCP service account key JSON.
        """
        return pulumi.get(self, "service_account_key")

    @service_account_key.setter
    def service_account_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_key", value)


class GcsIntegration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 service_account_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Rockset GCS Integration.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Text describing the integration.
        :param pulumi.Input[str] name: Unique identifier for the integration. Can contain alphanumeric or dash characters.
        :param pulumi.Input[str] service_account_key: The GCP service account key JSON.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GcsIntegrationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Rockset GCS Integration.

        :param str resource_name: The name of the resource.
        :param GcsIntegrationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GcsIntegrationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 service_account_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GcsIntegrationArgs.__new__(GcsIntegrationArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if service_account_key is None and not opts.urn:
                raise TypeError("Missing required property 'service_account_key'")
            __props__.__dict__["service_account_key"] = service_account_key
        super(GcsIntegration, __self__).__init__(
            'rockset:index/gcsIntegration:GcsIntegration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            service_account_key: Optional[pulumi.Input[str]] = None) -> 'GcsIntegration':
        """
        Get an existing GcsIntegration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Text describing the integration.
        :param pulumi.Input[str] name: Unique identifier for the integration. Can contain alphanumeric or dash characters.
        :param pulumi.Input[str] service_account_key: The GCP service account key JSON.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GcsIntegrationState.__new__(_GcsIntegrationState)

        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["service_account_key"] = service_account_key
        return GcsIntegration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Text describing the integration.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Unique identifier for the integration. Can contain alphanumeric or dash characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceAccountKey")
    def service_account_key(self) -> pulumi.Output[str]:
        """
        The GCP service account key JSON.
        """
        return pulumi.get(self, "service_account_key")
