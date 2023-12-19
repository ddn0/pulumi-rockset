# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetWorkspaceResult',
    'AwaitableGetWorkspaceResult',
    'get_workspace',
    'get_workspace_output',
]

@pulumi.output_type
class GetWorkspaceResult:
    """
    A collection of values returned by getWorkspace.
    """
    def __init__(__self__, collection_count=None, created_at=None, created_by=None, description=None, id=None, name=None):
        if collection_count and not isinstance(collection_count, int):
            raise TypeError("Expected argument 'collection_count' to be a int")
        pulumi.set(__self__, "collection_count", collection_count)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if created_by and not isinstance(created_by, str):
            raise TypeError("Expected argument 'created_by' to be a str")
        pulumi.set(__self__, "created_by", created_by)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="collectionCount")
    def collection_count(self) -> int:
        """
        Number of collections in the workspace.
        """
        return pulumi.get(self, "collection_count")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Created at in ISO-8601.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> str:
        """
        Created by.
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Workspace description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The workspace `name`.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Workspace name.
        """
        return pulumi.get(self, "name")


class AwaitableGetWorkspaceResult(GetWorkspaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkspaceResult(
            collection_count=self.collection_count,
            created_at=self.created_at,
            created_by=self.created_by,
            description=self.description,
            id=self.id,
            name=self.name)


def get_workspace(name: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkspaceResult:
    """
    Gets information about a workspace.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_rockset as rockset

    demo = rockset.get_workspace(name="demo")
    ```


    :param str name: Workspace name.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('rockset:index/getWorkspace:getWorkspace', __args__, opts=opts, typ=GetWorkspaceResult).value

    return AwaitableGetWorkspaceResult(
        collection_count=pulumi.get(__ret__, 'collection_count'),
        created_at=pulumi.get(__ret__, 'created_at'),
        created_by=pulumi.get(__ret__, 'created_by'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'))


@_utilities.lift_output_func(get_workspace)
def get_workspace_output(name: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWorkspaceResult]:
    """
    Gets information about a workspace.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_rockset as rockset

    demo = rockset.get_workspace(name="demo")
    ```


    :param str name: Workspace name.
    """
    ...