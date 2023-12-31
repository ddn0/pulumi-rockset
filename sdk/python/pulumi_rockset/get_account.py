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

__all__ = [
    'GetAccountResult',
    'AwaitableGetAccountResult',
    'get_account',
    'get_account_output',
]

@pulumi.output_type
class GetAccountResult:
    """
    A collection of values returned by getAccount.
    """
    def __init__(__self__, account_id=None, clusters=None, external_id=None, id=None, organization=None, rockset_user=None):
        if account_id and not isinstance(account_id, str):
            raise TypeError("Expected argument 'account_id' to be a str")
        pulumi.set(__self__, "account_id", account_id)
        if clusters and not isinstance(clusters, list):
            raise TypeError("Expected argument 'clusters' to be a list")
        pulumi.set(__self__, "clusters", clusters)
        if external_id and not isinstance(external_id, str):
            raise TypeError("Expected argument 'external_id' to be a str")
        pulumi.set(__self__, "external_id", external_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if organization and not isinstance(organization, str):
            raise TypeError("Expected argument 'organization' to be a str")
        pulumi.set(__self__, "organization", organization)
        if rockset_user and not isinstance(rockset_user, str):
            raise TypeError("Expected argument 'rockset_user' to be a str")
        pulumi.set(__self__, "rockset_user", rockset_user)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> str:
        """
        The AWS account ID to reference in AWS policies.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter
    def clusters(self) -> Sequence['outputs.GetAccountClusterResult']:
        """
        The Rockset clusters available to this API key.
        """
        return pulumi.get(self, "clusters")

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> str:
        """
        The external ID to use in AWS trust policies.
        """
        return pulumi.get(self, "external_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def organization(self) -> str:
        """
        The name of the organization for the API key.
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter(name="rocksetUser")
    def rockset_user(self) -> str:
        """
        The name of the Rockset user used for AWS trust policies.
        """
        return pulumi.get(self, "rockset_user")


class AwaitableGetAccountResult(GetAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountResult(
            account_id=self.account_id,
            clusters=self.clusters,
            external_id=self.external_id,
            id=self.id,
            organization=self.organization,
            rockset_user=self.rockset_user)


def get_account(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccountResult:
    """
    Gets information about the Rockset deployment for the specified api server.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('rockset:index/getAccount:getAccount', __args__, opts=opts, typ=GetAccountResult).value

    return AwaitableGetAccountResult(
        account_id=pulumi.get(__ret__, 'account_id'),
        clusters=pulumi.get(__ret__, 'clusters'),
        external_id=pulumi.get(__ret__, 'external_id'),
        id=pulumi.get(__ret__, 'id'),
        organization=pulumi.get(__ret__, 'organization'),
        rockset_user=pulumi.get(__ret__, 'rockset_user'))


@_utilities.lift_output_func(get_account)
def get_account_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccountResult]:
    """
    Gets information about the Rockset deployment for the specified api server.
    """
    ...
