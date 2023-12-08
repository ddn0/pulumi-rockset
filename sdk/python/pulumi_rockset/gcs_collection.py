# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['GcsCollectionArgs', 'GcsCollection']

@pulumi.input_type
class GcsCollectionArgs:
    def __init__(__self__, *,
                 workspace: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 ingest_transformation: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retention_secs: Optional[pulumi.Input[int]] = None,
                 sources: Optional[pulumi.Input[Sequence[pulumi.Input['GcsCollectionSourceArgs']]]] = None,
                 wait_for_collection: Optional[pulumi.Input[bool]] = None,
                 wait_for_documents: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a GcsCollection resource.
        :param pulumi.Input[str] workspace: The name of the workspace.
        :param pulumi.Input[str] description: Text describing the collection.
        :param pulumi.Input[str] ingest_transformation: Ingest transformation SQL query. Turns the collection into insert_only mode.
        :param pulumi.Input[str] name: Unique identifier for the collection. Can contain alphanumeric or dash characters.
        :param pulumi.Input[int] retention_secs: Number of seconds after which data is purged. Based on event time.
        :param pulumi.Input[Sequence[pulumi.Input['GcsCollectionSourceArgs']]] sources: Defines a source for this collection.
        :param pulumi.Input[bool] wait_for_collection: Wait until the collection is ready.
        :param pulumi.Input[int] wait_for_documents: Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
        """
        pulumi.set(__self__, "workspace", workspace)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ingest_transformation is not None:
            pulumi.set(__self__, "ingest_transformation", ingest_transformation)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if retention_secs is not None:
            pulumi.set(__self__, "retention_secs", retention_secs)
        if sources is not None:
            pulumi.set(__self__, "sources", sources)
        if wait_for_collection is not None:
            pulumi.set(__self__, "wait_for_collection", wait_for_collection)
        if wait_for_documents is not None:
            pulumi.set(__self__, "wait_for_documents", wait_for_documents)

    @property
    @pulumi.getter
    def workspace(self) -> pulumi.Input[str]:
        """
        The name of the workspace.
        """
        return pulumi.get(self, "workspace")

    @workspace.setter
    def workspace(self, value: pulumi.Input[str]):
        pulumi.set(self, "workspace", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Text describing the collection.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ingestTransformation")
    def ingest_transformation(self) -> Optional[pulumi.Input[str]]:
        """
        Ingest transformation SQL query. Turns the collection into insert_only mode.
        """
        return pulumi.get(self, "ingest_transformation")

    @ingest_transformation.setter
    def ingest_transformation(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ingest_transformation", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier for the collection. Can contain alphanumeric or dash characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="retentionSecs")
    def retention_secs(self) -> Optional[pulumi.Input[int]]:
        """
        Number of seconds after which data is purged. Based on event time.
        """
        return pulumi.get(self, "retention_secs")

    @retention_secs.setter
    def retention_secs(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_secs", value)

    @property
    @pulumi.getter
    def sources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GcsCollectionSourceArgs']]]]:
        """
        Defines a source for this collection.
        """
        return pulumi.get(self, "sources")

    @sources.setter
    def sources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GcsCollectionSourceArgs']]]]):
        pulumi.set(self, "sources", value)

    @property
    @pulumi.getter(name="waitForCollection")
    def wait_for_collection(self) -> Optional[pulumi.Input[bool]]:
        """
        Wait until the collection is ready.
        """
        return pulumi.get(self, "wait_for_collection")

    @wait_for_collection.setter
    def wait_for_collection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_for_collection", value)

    @property
    @pulumi.getter(name="waitForDocuments")
    def wait_for_documents(self) -> Optional[pulumi.Input[int]]:
        """
        Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
        """
        return pulumi.get(self, "wait_for_documents")

    @wait_for_documents.setter
    def wait_for_documents(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "wait_for_documents", value)


@pulumi.input_type
class _GcsCollectionState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 ingest_transformation: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retention_secs: Optional[pulumi.Input[int]] = None,
                 sources: Optional[pulumi.Input[Sequence[pulumi.Input['GcsCollectionSourceArgs']]]] = None,
                 wait_for_collection: Optional[pulumi.Input[bool]] = None,
                 wait_for_documents: Optional[pulumi.Input[int]] = None,
                 workspace: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GcsCollection resources.
        :param pulumi.Input[str] description: Text describing the collection.
        :param pulumi.Input[str] ingest_transformation: Ingest transformation SQL query. Turns the collection into insert_only mode.
        :param pulumi.Input[str] name: Unique identifier for the collection. Can contain alphanumeric or dash characters.
        :param pulumi.Input[int] retention_secs: Number of seconds after which data is purged. Based on event time.
        :param pulumi.Input[Sequence[pulumi.Input['GcsCollectionSourceArgs']]] sources: Defines a source for this collection.
        :param pulumi.Input[bool] wait_for_collection: Wait until the collection is ready.
        :param pulumi.Input[int] wait_for_documents: Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
        :param pulumi.Input[str] workspace: The name of the workspace.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ingest_transformation is not None:
            pulumi.set(__self__, "ingest_transformation", ingest_transformation)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if retention_secs is not None:
            pulumi.set(__self__, "retention_secs", retention_secs)
        if sources is not None:
            pulumi.set(__self__, "sources", sources)
        if wait_for_collection is not None:
            pulumi.set(__self__, "wait_for_collection", wait_for_collection)
        if wait_for_documents is not None:
            pulumi.set(__self__, "wait_for_documents", wait_for_documents)
        if workspace is not None:
            pulumi.set(__self__, "workspace", workspace)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Text describing the collection.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ingestTransformation")
    def ingest_transformation(self) -> Optional[pulumi.Input[str]]:
        """
        Ingest transformation SQL query. Turns the collection into insert_only mode.
        """
        return pulumi.get(self, "ingest_transformation")

    @ingest_transformation.setter
    def ingest_transformation(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ingest_transformation", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier for the collection. Can contain alphanumeric or dash characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="retentionSecs")
    def retention_secs(self) -> Optional[pulumi.Input[int]]:
        """
        Number of seconds after which data is purged. Based on event time.
        """
        return pulumi.get(self, "retention_secs")

    @retention_secs.setter
    def retention_secs(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_secs", value)

    @property
    @pulumi.getter
    def sources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GcsCollectionSourceArgs']]]]:
        """
        Defines a source for this collection.
        """
        return pulumi.get(self, "sources")

    @sources.setter
    def sources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GcsCollectionSourceArgs']]]]):
        pulumi.set(self, "sources", value)

    @property
    @pulumi.getter(name="waitForCollection")
    def wait_for_collection(self) -> Optional[pulumi.Input[bool]]:
        """
        Wait until the collection is ready.
        """
        return pulumi.get(self, "wait_for_collection")

    @wait_for_collection.setter
    def wait_for_collection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_for_collection", value)

    @property
    @pulumi.getter(name="waitForDocuments")
    def wait_for_documents(self) -> Optional[pulumi.Input[int]]:
        """
        Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
        """
        return pulumi.get(self, "wait_for_documents")

    @wait_for_documents.setter
    def wait_for_documents(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "wait_for_documents", value)

    @property
    @pulumi.getter
    def workspace(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the workspace.
        """
        return pulumi.get(self, "workspace")

    @workspace.setter
    def workspace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workspace", value)


class GcsCollection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ingest_transformation: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retention_secs: Optional[pulumi.Input[int]] = None,
                 sources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GcsCollectionSourceArgs']]]]] = None,
                 wait_for_collection: Optional[pulumi.Input[bool]] = None,
                 wait_for_documents: Optional[pulumi.Input[int]] = None,
                 workspace: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a collection with an GCS source attached.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Text describing the collection.
        :param pulumi.Input[str] ingest_transformation: Ingest transformation SQL query. Turns the collection into insert_only mode.
        :param pulumi.Input[str] name: Unique identifier for the collection. Can contain alphanumeric or dash characters.
        :param pulumi.Input[int] retention_secs: Number of seconds after which data is purged. Based on event time.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GcsCollectionSourceArgs']]]] sources: Defines a source for this collection.
        :param pulumi.Input[bool] wait_for_collection: Wait until the collection is ready.
        :param pulumi.Input[int] wait_for_documents: Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
        :param pulumi.Input[str] workspace: The name of the workspace.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GcsCollectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a collection with an GCS source attached.

        :param str resource_name: The name of the resource.
        :param GcsCollectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GcsCollectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ingest_transformation: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retention_secs: Optional[pulumi.Input[int]] = None,
                 sources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GcsCollectionSourceArgs']]]]] = None,
                 wait_for_collection: Optional[pulumi.Input[bool]] = None,
                 wait_for_documents: Optional[pulumi.Input[int]] = None,
                 workspace: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GcsCollectionArgs.__new__(GcsCollectionArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["ingest_transformation"] = ingest_transformation
            __props__.__dict__["name"] = name
            __props__.__dict__["retention_secs"] = retention_secs
            __props__.__dict__["sources"] = sources
            __props__.__dict__["wait_for_collection"] = wait_for_collection
            __props__.__dict__["wait_for_documents"] = wait_for_documents
            if workspace is None and not opts.urn:
                raise TypeError("Missing required property 'workspace'")
            __props__.__dict__["workspace"] = workspace
        super(GcsCollection, __self__).__init__(
            'rockset:index/gcsCollection:GcsCollection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            ingest_transformation: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            retention_secs: Optional[pulumi.Input[int]] = None,
            sources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GcsCollectionSourceArgs']]]]] = None,
            wait_for_collection: Optional[pulumi.Input[bool]] = None,
            wait_for_documents: Optional[pulumi.Input[int]] = None,
            workspace: Optional[pulumi.Input[str]] = None) -> 'GcsCollection':
        """
        Get an existing GcsCollection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Text describing the collection.
        :param pulumi.Input[str] ingest_transformation: Ingest transformation SQL query. Turns the collection into insert_only mode.
        :param pulumi.Input[str] name: Unique identifier for the collection. Can contain alphanumeric or dash characters.
        :param pulumi.Input[int] retention_secs: Number of seconds after which data is purged. Based on event time.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GcsCollectionSourceArgs']]]] sources: Defines a source for this collection.
        :param pulumi.Input[bool] wait_for_collection: Wait until the collection is ready.
        :param pulumi.Input[int] wait_for_documents: Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
        :param pulumi.Input[str] workspace: The name of the workspace.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GcsCollectionState.__new__(_GcsCollectionState)

        __props__.__dict__["description"] = description
        __props__.__dict__["ingest_transformation"] = ingest_transformation
        __props__.__dict__["name"] = name
        __props__.__dict__["retention_secs"] = retention_secs
        __props__.__dict__["sources"] = sources
        __props__.__dict__["wait_for_collection"] = wait_for_collection
        __props__.__dict__["wait_for_documents"] = wait_for_documents
        __props__.__dict__["workspace"] = workspace
        return GcsCollection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Text describing the collection.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ingestTransformation")
    def ingest_transformation(self) -> pulumi.Output[Optional[str]]:
        """
        Ingest transformation SQL query. Turns the collection into insert_only mode.
        """
        return pulumi.get(self, "ingest_transformation")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Unique identifier for the collection. Can contain alphanumeric or dash characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="retentionSecs")
    def retention_secs(self) -> pulumi.Output[Optional[int]]:
        """
        Number of seconds after which data is purged. Based on event time.
        """
        return pulumi.get(self, "retention_secs")

    @property
    @pulumi.getter
    def sources(self) -> pulumi.Output[Optional[Sequence['outputs.GcsCollectionSource']]]:
        """
        Defines a source for this collection.
        """
        return pulumi.get(self, "sources")

    @property
    @pulumi.getter(name="waitForCollection")
    def wait_for_collection(self) -> pulumi.Output[Optional[bool]]:
        """
        Wait until the collection is ready.
        """
        return pulumi.get(self, "wait_for_collection")

    @property
    @pulumi.getter(name="waitForDocuments")
    def wait_for_documents(self) -> pulumi.Output[Optional[int]]:
        """
        Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
        """
        return pulumi.get(self, "wait_for_documents")

    @property
    @pulumi.getter
    def workspace(self) -> pulumi.Output[str]:
        """
        The name of the workspace.
        """
        return pulumi.get(self, "workspace")

