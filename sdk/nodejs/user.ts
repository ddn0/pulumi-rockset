// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a Rockset User.
 *
 * First and last name can only be managed for users who have accepted the invite,
 * i.e. when the state is ACCEPTED.
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rockset:index/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * The ISO-8601 time of when the user was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Email address of the user. Also used to identify the user.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * User's first name. This can only be set once the state is ACTIVE, i.e after the user has accepted the invite.
     */
    public readonly firstName!: pulumi.Output<string | undefined>;
    /**
     * User's last name. This can only be set once the state is ACTIVE, i.e after the user has accepted the invite.
     */
    public readonly lastName!: pulumi.Output<string | undefined>;
    /**
     * List of roles for the user. E.g. 'admin', 'member', 'read-only'.
     */
    public readonly roles!: pulumi.Output<string[]>;
    /**
     * State of the user, either NEW or ACTIVE.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["firstName"] = state ? state.firstName : undefined;
            resourceInputs["lastName"] = state ? state.lastName : undefined;
            resourceInputs["roles"] = state ? state.roles : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.roles === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roles'");
            }
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["firstName"] = args ? args.firstName : undefined;
            resourceInputs["lastName"] = args ? args.lastName : undefined;
            resourceInputs["roles"] = args ? args.roles : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * The ISO-8601 time of when the user was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Email address of the user. Also used to identify the user.
     */
    email?: pulumi.Input<string>;
    /**
     * User's first name. This can only be set once the state is ACTIVE, i.e after the user has accepted the invite.
     */
    firstName?: pulumi.Input<string>;
    /**
     * User's last name. This can only be set once the state is ACTIVE, i.e after the user has accepted the invite.
     */
    lastName?: pulumi.Input<string>;
    /**
     * List of roles for the user. E.g. 'admin', 'member', 'read-only'.
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * State of the user, either NEW or ACTIVE.
     */
    state?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * Email address of the user. Also used to identify the user.
     */
    email: pulumi.Input<string>;
    /**
     * User's first name. This can only be set once the state is ACTIVE, i.e after the user has accepted the invite.
     */
    firstName?: pulumi.Input<string>;
    /**
     * User's last name. This can only be set once the state is ACTIVE, i.e after the user has accepted the invite.
     */
    lastName?: pulumi.Input<string>;
    /**
     * List of roles for the user. E.g. 'admin', 'member', 'read-only'.
     */
    roles: pulumi.Input<pulumi.Input<string>[]>;
}
