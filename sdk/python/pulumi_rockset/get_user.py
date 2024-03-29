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
    'GetUserResult',
    'AwaitableGetUserResult',
    'get_user',
    'get_user_output',
]

@pulumi.output_type
class GetUserResult:
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, email=None, first_name=None, id=None, last_name=None, roles=None, state=None):
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if first_name and not isinstance(first_name, str):
            raise TypeError("Expected argument 'first_name' to be a str")
        pulumi.set(__self__, "first_name", first_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_name and not isinstance(last_name, str):
            raise TypeError("Expected argument 'last_name' to be a str")
        pulumi.set(__self__, "last_name", last_name)
        if roles and not isinstance(roles, list):
            raise TypeError("Expected argument 'roles' to be a list")
        pulumi.set(__self__, "roles", roles)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def email(self) -> Optional[str]:
        """
        User email. If absent or blank, it gets the current user.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> str:
        """
        User's first name.
        """
        return pulumi.get(self, "first_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The user ID, in the form of the `email`.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> str:
        """
        User's last name.
        """
        return pulumi.get(self, "last_name")

    @property
    @pulumi.getter
    def roles(self) -> Sequence[str]:
        """
        List of roles for the user. E.g. 'admin', 'member', 'read-only'.
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        State of the user, either NEW or ACTIVE.
        """
        return pulumi.get(self, "state")


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            email=self.email,
            first_name=self.first_name,
            id=self.id,
            last_name=self.last_name,
            roles=self.roles,
            state=self.state)


def get_user(email: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    This data source can be used to fetch information about a specific user.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_rockset as rockset

    pme = rockset.get_user(email="pme@rockset.com")
    ```


    :param str email: User email. If absent or blank, it gets the current user.
    """
    __args__ = dict()
    __args__['email'] = email
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('rockset:index/getUser:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        email=pulumi.get(__ret__, 'email'),
        first_name=pulumi.get(__ret__, 'first_name'),
        id=pulumi.get(__ret__, 'id'),
        last_name=pulumi.get(__ret__, 'last_name'),
        roles=pulumi.get(__ret__, 'roles'),
        state=pulumi.get(__ret__, 'state'))


@_utilities.lift_output_func(get_user)
def get_user_output(email: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserResult]:
    """
    This data source can be used to fetch information about a specific user.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_rockset as rockset

    pme = rockset.get_user(email="pme@rockset.com")
    ```


    :param str email: User email. If absent or blank, it gets the current user.
    """
    ...
