// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a basic collection with no sources. Usually used for the write api.
 */
export class Collection extends pulumi.CustomResource {
    /**
     * Get an existing Collection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CollectionState, opts?: pulumi.CustomResourceOptions): Collection {
        return new Collection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rockset:index/collection:Collection';

    /**
     * Returns true if the given object is an instance of Collection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Collection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Collection.__pulumiType;
    }

    /**
     * Text describing the collection.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Ingest transformation SQL query. Turns the collection into insertOnly mode.
     */
    public readonly ingestTransformation!: pulumi.Output<string | undefined>;
    /**
     * Unique identifier for the collection. Can contain alphanumeric or dash characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Number of seconds after which data is purged. Based on event time.
     */
    public readonly retentionSecs!: pulumi.Output<number | undefined>;
    /**
     * Wait until the collection is ready.
     */
    public readonly waitForCollection!: pulumi.Output<boolean | undefined>;
    /**
     * Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
     */
    public readonly waitForDocuments!: pulumi.Output<number | undefined>;
    /**
     * The name of the workspace.
     */
    public readonly workspace!: pulumi.Output<string>;

    /**
     * Create a Collection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CollectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CollectionArgs | CollectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CollectionState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["ingestTransformation"] = state ? state.ingestTransformation : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["retentionSecs"] = state ? state.retentionSecs : undefined;
            resourceInputs["waitForCollection"] = state ? state.waitForCollection : undefined;
            resourceInputs["waitForDocuments"] = state ? state.waitForDocuments : undefined;
            resourceInputs["workspace"] = state ? state.workspace : undefined;
        } else {
            const args = argsOrState as CollectionArgs | undefined;
            if ((!args || args.workspace === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspace'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ingestTransformation"] = args ? args.ingestTransformation : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["retentionSecs"] = args ? args.retentionSecs : undefined;
            resourceInputs["waitForCollection"] = args ? args.waitForCollection : undefined;
            resourceInputs["waitForDocuments"] = args ? args.waitForDocuments : undefined;
            resourceInputs["workspace"] = args ? args.workspace : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Collection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Collection resources.
 */
export interface CollectionState {
    /**
     * Text describing the collection.
     */
    description?: pulumi.Input<string>;
    /**
     * Ingest transformation SQL query. Turns the collection into insertOnly mode.
     */
    ingestTransformation?: pulumi.Input<string>;
    /**
     * Unique identifier for the collection. Can contain alphanumeric or dash characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Number of seconds after which data is purged. Based on event time.
     */
    retentionSecs?: pulumi.Input<number>;
    /**
     * Wait until the collection is ready.
     */
    waitForCollection?: pulumi.Input<boolean>;
    /**
     * Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
     */
    waitForDocuments?: pulumi.Input<number>;
    /**
     * The name of the workspace.
     */
    workspace?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Collection resource.
 */
export interface CollectionArgs {
    /**
     * Text describing the collection.
     */
    description?: pulumi.Input<string>;
    /**
     * Ingest transformation SQL query. Turns the collection into insertOnly mode.
     */
    ingestTransformation?: pulumi.Input<string>;
    /**
     * Unique identifier for the collection. Can contain alphanumeric or dash characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Number of seconds after which data is purged. Based on event time.
     */
    retentionSecs?: pulumi.Input<number>;
    /**
     * Wait until the collection is ready.
     */
    waitForCollection?: pulumi.Input<boolean>;
    /**
     * Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
     */
    waitForDocuments?: pulumi.Input<number>;
    /**
     * The name of the workspace.
     */
    workspace: pulumi.Input<string>;
}
