# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['S3IntegrationArgs', 'S3Integration']

@pulumi.input_type
class S3IntegrationArgs:
    def __init__(__self__, *,
                 aws_role_arn: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a S3Integration resource.
        :param pulumi.Input[str] aws_role_arn: The AWS Role Arn to use for this integration.
        :param pulumi.Input[str] description: Text describing the integration.
        :param pulumi.Input[str] name: Unique identifier for the integration. Can contain alphanumeric or dash characters.
        """
        pulumi.set(__self__, "aws_role_arn", aws_role_arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="awsRoleArn")
    def aws_role_arn(self) -> pulumi.Input[str]:
        """
        The AWS Role Arn to use for this integration.
        """
        return pulumi.get(self, "aws_role_arn")

    @aws_role_arn.setter
    def aws_role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "aws_role_arn", value)

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
class _S3IntegrationState:
    def __init__(__self__, *,
                 aws_role_arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering S3Integration resources.
        :param pulumi.Input[str] aws_role_arn: The AWS Role Arn to use for this integration.
        :param pulumi.Input[str] description: Text describing the integration.
        :param pulumi.Input[str] name: Unique identifier for the integration. Can contain alphanumeric or dash characters.
        """
        if aws_role_arn is not None:
            pulumi.set(__self__, "aws_role_arn", aws_role_arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="awsRoleArn")
    def aws_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS Role Arn to use for this integration.
        """
        return pulumi.get(self, "aws_role_arn")

    @aws_role_arn.setter
    def aws_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_role_arn", value)

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


class S3Integration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_role_arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Rockset S3 Integration.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_role_arn: The AWS Role Arn to use for this integration.
        :param pulumi.Input[str] description: Text describing the integration.
        :param pulumi.Input[str] name: Unique identifier for the integration. Can contain alphanumeric or dash characters.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: S3IntegrationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Rockset S3 Integration.

        :param str resource_name: The name of the resource.
        :param S3IntegrationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(S3IntegrationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_role_arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = S3IntegrationArgs.__new__(S3IntegrationArgs)

            if aws_role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'aws_role_arn'")
            __props__.__dict__["aws_role_arn"] = aws_role_arn
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
        super(S3Integration, __self__).__init__(
            'rockset:index/s3Integration:S3Integration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            aws_role_arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'S3Integration':
        """
        Get an existing S3Integration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_role_arn: The AWS Role Arn to use for this integration.
        :param pulumi.Input[str] description: Text describing the integration.
        :param pulumi.Input[str] name: Unique identifier for the integration. Can contain alphanumeric or dash characters.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _S3IntegrationState.__new__(_S3IntegrationState)

        __props__.__dict__["aws_role_arn"] = aws_role_arn
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        return S3Integration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="awsRoleArn")
    def aws_role_arn(self) -> pulumi.Output[str]:
        """
        The AWS Role Arn to use for this integration.
        """
        return pulumi.get(self, "aws_role_arn")

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

