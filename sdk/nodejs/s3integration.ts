// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a Rockset S3 Integration.
 */
export class S3Integration extends pulumi.CustomResource {
    /**
     * Get an existing S3Integration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: S3IntegrationState, opts?: pulumi.CustomResourceOptions): S3Integration {
        return new S3Integration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rockset:index/s3Integration:S3Integration';

    /**
     * Returns true if the given object is an instance of S3Integration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is S3Integration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === S3Integration.__pulumiType;
    }

    /**
     * The AWS Role Arn to use for this integration.
     */
    public readonly awsRoleArn!: pulumi.Output<string>;
    /**
     * Text describing the integration.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Unique identifier for the integration. Can contain alphanumeric or dash characters.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a S3Integration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: S3IntegrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: S3IntegrationArgs | S3IntegrationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as S3IntegrationState | undefined;
            resourceInputs["awsRoleArn"] = state ? state.awsRoleArn : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as S3IntegrationArgs | undefined;
            if ((!args || args.awsRoleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'awsRoleArn'");
            }
            resourceInputs["awsRoleArn"] = args ? args.awsRoleArn : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(S3Integration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering S3Integration resources.
 */
export interface S3IntegrationState {
    /**
     * The AWS Role Arn to use for this integration.
     */
    awsRoleArn?: pulumi.Input<string>;
    /**
     * Text describing the integration.
     */
    description?: pulumi.Input<string>;
    /**
     * Unique identifier for the integration. Can contain alphanumeric or dash characters.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a S3Integration resource.
 */
export interface S3IntegrationArgs {
    /**
     * The AWS Role Arn to use for this integration.
     */
    awsRoleArn: pulumi.Input<string>;
    /**
     * Text describing the integration.
     */
    description?: pulumi.Input<string>;
    /**
     * Unique identifier for the integration. Can contain alphanumeric or dash characters.
     */
    name?: pulumi.Input<string>;
}
