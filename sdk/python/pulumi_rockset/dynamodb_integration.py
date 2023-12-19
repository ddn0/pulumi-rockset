# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DynamodbIntegrationArgs', 'DynamodbIntegration']

@pulumi.input_type
class DynamodbIntegrationArgs:
    def __init__(__self__, *,
                 aws_role_arn: pulumi.Input[str],
                 s3_export_bucket_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DynamodbIntegration resource.
        :param pulumi.Input[str] aws_role_arn: The AWS Role Arn to use for this integration.
        :param pulumi.Input[str] s3_export_bucket_name: AWS S3 bucket name used for exporting the DynamoDB tables.
        :param pulumi.Input[str] description: Text describing the integration.
        :param pulumi.Input[str] name: Unique identifier for the integration. Can contain alphanumeric or dash characters.
        """
        pulumi.set(__self__, "aws_role_arn", aws_role_arn)
        pulumi.set(__self__, "s3_export_bucket_name", s3_export_bucket_name)
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
    @pulumi.getter(name="s3ExportBucketName")
    def s3_export_bucket_name(self) -> pulumi.Input[str]:
        """
        AWS S3 bucket name used for exporting the DynamoDB tables.
        """
        return pulumi.get(self, "s3_export_bucket_name")

    @s3_export_bucket_name.setter
    def s3_export_bucket_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "s3_export_bucket_name", value)

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
class _DynamodbIntegrationState:
    def __init__(__self__, *,
                 aws_role_arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 s3_export_bucket_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DynamodbIntegration resources.
        :param pulumi.Input[str] aws_role_arn: The AWS Role Arn to use for this integration.
        :param pulumi.Input[str] description: Text describing the integration.
        :param pulumi.Input[str] name: Unique identifier for the integration. Can contain alphanumeric or dash characters.
        :param pulumi.Input[str] s3_export_bucket_name: AWS S3 bucket name used for exporting the DynamoDB tables.
        """
        if aws_role_arn is not None:
            pulumi.set(__self__, "aws_role_arn", aws_role_arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if s3_export_bucket_name is not None:
            pulumi.set(__self__, "s3_export_bucket_name", s3_export_bucket_name)

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

    @property
    @pulumi.getter(name="s3ExportBucketName")
    def s3_export_bucket_name(self) -> Optional[pulumi.Input[str]]:
        """
        AWS S3 bucket name used for exporting the DynamoDB tables.
        """
        return pulumi.get(self, "s3_export_bucket_name")

    @s3_export_bucket_name.setter
    def s3_export_bucket_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s3_export_bucket_name", value)


class DynamodbIntegration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_role_arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 s3_export_bucket_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Rockset DynamoDB Integration.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_role_arn: The AWS Role Arn to use for this integration.
        :param pulumi.Input[str] description: Text describing the integration.
        :param pulumi.Input[str] name: Unique identifier for the integration. Can contain alphanumeric or dash characters.
        :param pulumi.Input[str] s3_export_bucket_name: AWS S3 bucket name used for exporting the DynamoDB tables.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DynamodbIntegrationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Rockset DynamoDB Integration.

        :param str resource_name: The name of the resource.
        :param DynamodbIntegrationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DynamodbIntegrationArgs, pulumi.ResourceOptions, *args, **kwargs)
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
                 s3_export_bucket_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DynamodbIntegrationArgs.__new__(DynamodbIntegrationArgs)

            if aws_role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'aws_role_arn'")
            __props__.__dict__["aws_role_arn"] = aws_role_arn
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if s3_export_bucket_name is None and not opts.urn:
                raise TypeError("Missing required property 's3_export_bucket_name'")
            __props__.__dict__["s3_export_bucket_name"] = s3_export_bucket_name
        super(DynamodbIntegration, __self__).__init__(
            'rockset:index/dynamodbIntegration:DynamodbIntegration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            aws_role_arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            s3_export_bucket_name: Optional[pulumi.Input[str]] = None) -> 'DynamodbIntegration':
        """
        Get an existing DynamodbIntegration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_role_arn: The AWS Role Arn to use for this integration.
        :param pulumi.Input[str] description: Text describing the integration.
        :param pulumi.Input[str] name: Unique identifier for the integration. Can contain alphanumeric or dash characters.
        :param pulumi.Input[str] s3_export_bucket_name: AWS S3 bucket name used for exporting the DynamoDB tables.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DynamodbIntegrationState.__new__(_DynamodbIntegrationState)

        __props__.__dict__["aws_role_arn"] = aws_role_arn
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["s3_export_bucket_name"] = s3_export_bucket_name
        return DynamodbIntegration(resource_name, opts=opts, __props__=__props__)

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

    @property
    @pulumi.getter(name="s3ExportBucketName")
    def s3_export_bucket_name(self) -> pulumi.Output[str]:
        """
        AWS S3 bucket name used for exporting the DynamoDB tables.
        """
        return pulumi.get(self, "s3_export_bucket_name")
