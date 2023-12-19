// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about a workspace.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rockset from "@pulumi/rockset";
 *
 * const demo = rockset.getWorkspace({
 *     name: "demo",
 * });
 * ```
 */
export function getWorkspace(args: GetWorkspaceArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkspaceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("rockset:index/getWorkspace:getWorkspace", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getWorkspace.
 */
export interface GetWorkspaceArgs {
    /**
     * Workspace name.
     */
    name: string;
}

/**
 * A collection of values returned by getWorkspace.
 */
export interface GetWorkspaceResult {
    /**
     * Number of collections in the workspace.
     */
    readonly collectionCount: number;
    /**
     * Created at in ISO-8601.
     */
    readonly createdAt: string;
    /**
     * Created by.
     */
    readonly createdBy: string;
    /**
     * Workspace description.
     */
    readonly description: string;
    /**
     * The workspace `name`.
     */
    readonly id: string;
    /**
     * Workspace name.
     */
    readonly name: string;
}
/**
 * Gets information about a workspace.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rockset from "@pulumi/rockset";
 *
 * const demo = rockset.getWorkspace({
 *     name: "demo",
 * });
 * ```
 */
export function getWorkspaceOutput(args: GetWorkspaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWorkspaceResult> {
    return pulumi.output(args).apply((a: any) => getWorkspace(a, opts))
}

/**
 * A collection of arguments for invoking getWorkspace.
 */
export interface GetWorkspaceOutputArgs {
    /**
     * Workspace name.
     */
    name: pulumi.Input<string>;
}