// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can be used to fetch information about a specific user.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rockset from "@pulumi/rockset";
 *
 * const pme = rockset.getUser({
 *     email: "pme@rockset.com",
 * });
 * ```
 */
export function getUser(args?: GetUserArgs, opts?: pulumi.InvokeOptions): Promise<GetUserResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("rockset:index/getUser:getUser", {
        "email": args.email,
    }, opts);
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserArgs {
    /**
     * User email. If absent or blank, it gets the current user.
     */
    email?: string;
}

/**
 * A collection of values returned by getUser.
 */
export interface GetUserResult {
    /**
     * User email. If absent or blank, it gets the current user.
     */
    readonly email?: string;
    /**
     * User's first name.
     */
    readonly firstName: string;
    /**
     * The user ID, in the form of the `email`.
     */
    readonly id: string;
    /**
     * User's last name.
     */
    readonly lastName: string;
    /**
     * List of roles for the user. E.g. 'admin', 'member', 'read-only'.
     */
    readonly roles: string[];
    /**
     * State of the user, either NEW or ACTIVE.
     */
    readonly state: string;
}
/**
 * This data source can be used to fetch information about a specific user.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rockset from "@pulumi/rockset";
 *
 * const pme = rockset.getUser({
 *     email: "pme@rockset.com",
 * });
 * ```
 */
export function getUserOutput(args?: GetUserOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserResult> {
    return pulumi.output(args).apply((a: any) => getUser(a, opts))
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserOutputArgs {
    /**
     * User email. If absent or blank, it gets the current user.
     */
    email?: pulumi.Input<string>;
}
