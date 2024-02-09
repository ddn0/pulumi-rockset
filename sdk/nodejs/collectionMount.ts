// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a collection mount.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rockset from "@pulumi/rockset";
 *
 * const data = new rockset.CollectionMount("data", {
 *     virtualInstanceId: rockset_virtual_instance.secondary.id,
 *     path: "commons.data",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import rockset:index/collectionMount:CollectionMount demo 29e4a43c-fff4-4fe6-80e3-1ee57bc22e82
 * ```
 */
export class CollectionMount extends pulumi.CustomResource {
    /**
     * Get an existing CollectionMount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CollectionMountState, opts?: pulumi.CustomResourceOptions): CollectionMount {
        return new CollectionMount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rockset:index/collectionMount:CollectionMount';

    /**
     * Returns true if the given object is an instance of CollectionMount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CollectionMount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CollectionMount.__pulumiType;
    }

    /**
     * ISO 8601 date when the mount was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * UNIX timestamp in milliseconds for most recent refresh. Not applicable for live mounts.
     */
    public /*out*/ readonly lastRefreshTime!: pulumi.Output<number>;
    /**
     * Collection path to be mounted, in the form workspace.collection
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * RRN of this mount.
     */
    public /*out*/ readonly rrn!: pulumi.Output<string>;
    /**
     * Mount state.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Virtual Instance id
     */
    public readonly virtualInstanceId!: pulumi.Output<string>;
    /**
     * Virtual Instance RRN
     */
    public /*out*/ readonly virtualInstanceRrn!: pulumi.Output<string>;

    /**
     * Create a CollectionMount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CollectionMountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CollectionMountArgs | CollectionMountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CollectionMountState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["lastRefreshTime"] = state ? state.lastRefreshTime : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["rrn"] = state ? state.rrn : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["virtualInstanceId"] = state ? state.virtualInstanceId : undefined;
            resourceInputs["virtualInstanceRrn"] = state ? state.virtualInstanceRrn : undefined;
        } else {
            const args = argsOrState as CollectionMountArgs | undefined;
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            if ((!args || args.virtualInstanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'virtualInstanceId'");
            }
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["virtualInstanceId"] = args ? args.virtualInstanceId : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["lastRefreshTime"] = undefined /*out*/;
            resourceInputs["rrn"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["virtualInstanceRrn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CollectionMount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CollectionMount resources.
 */
export interface CollectionMountState {
    /**
     * ISO 8601 date when the mount was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * UNIX timestamp in milliseconds for most recent refresh. Not applicable for live mounts.
     */
    lastRefreshTime?: pulumi.Input<number>;
    /**
     * Collection path to be mounted, in the form workspace.collection
     */
    path?: pulumi.Input<string>;
    /**
     * RRN of this mount.
     */
    rrn?: pulumi.Input<string>;
    /**
     * Mount state.
     */
    state?: pulumi.Input<string>;
    /**
     * Virtual Instance id
     */
    virtualInstanceId?: pulumi.Input<string>;
    /**
     * Virtual Instance RRN
     */
    virtualInstanceRrn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CollectionMount resource.
 */
export interface CollectionMountArgs {
    /**
     * Collection path to be mounted, in the form workspace.collection
     */
    path: pulumi.Input<string>;
    /**
     * Virtual Instance id
     */
    virtualInstanceId: pulumi.Input<string>;
}
