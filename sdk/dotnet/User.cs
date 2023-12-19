// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rockset
{
    /// <summary>
    /// Manages a Rockset User.
    /// 
    /// First and last name can only be managed for users who have accepted the invite,
    /// i.e. when the state is ACCEPTED.
    /// </summary>
    [RocksetResourceType("rockset:index/user:User")]
    public partial class User : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ISO-8601 time of when the user was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Email address of the user. Also used to identify the user.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// User's first name. This can only be set once the state is ACTIVE, i.e after the user has accepted the invite.
        /// </summary>
        [Output("firstName")]
        public Output<string?> FirstName { get; private set; } = null!;

        /// <summary>
        /// User's last name. This can only be set once the state is ACTIVE, i.e after the user has accepted the invite.
        /// </summary>
        [Output("lastName")]
        public Output<string?> LastName { get; private set; } = null!;

        /// <summary>
        /// List of roles for the user. E.g. 'admin', 'member', 'read-only'.
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;

        /// <summary>
        /// State of the user, either NEW or ACTIVE.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("rockset:index/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("rockset:index/user:User", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Email address of the user. Also used to identify the user.
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        /// <summary>
        /// User's first name. This can only be set once the state is ACTIVE, i.e after the user has accepted the invite.
        /// </summary>
        [Input("firstName")]
        public Input<string>? FirstName { get; set; }

        /// <summary>
        /// User's last name. This can only be set once the state is ACTIVE, i.e after the user has accepted the invite.
        /// </summary>
        [Input("lastName")]
        public Input<string>? LastName { get; set; }

        [Input("roles", required: true)]
        private InputList<string>? _roles;

        /// <summary>
        /// List of roles for the user. E.g. 'admin', 'member', 'read-only'.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        public UserArgs()
        {
        }
        public static new UserArgs Empty => new UserArgs();
    }

    public sealed class UserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ISO-8601 time of when the user was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Email address of the user. Also used to identify the user.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// User's first name. This can only be set once the state is ACTIVE, i.e after the user has accepted the invite.
        /// </summary>
        [Input("firstName")]
        public Input<string>? FirstName { get; set; }

        /// <summary>
        /// User's last name. This can only be set once the state is ACTIVE, i.e after the user has accepted the invite.
        /// </summary>
        [Input("lastName")]
        public Input<string>? LastName { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// List of roles for the user. E.g. 'admin', 'member', 'read-only'.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// State of the user, either NEW or ACTIVE.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public UserState()
        {
        }
        public static new UserState Empty => new UserState();
    }
}