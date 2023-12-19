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
    /// Manage a Rockset Api Key.
    /// 
    /// Can be used together with roles to scope the actions the api key can take.
    /// </summary>
    [RocksetResourceType("rockset:index/apiKey:ApiKey")]
    public partial class ApiKey : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The resulting Rockset api key.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// Name of the api key.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The role the api key will use. If not specified, "All User Assigned Roles" will be used.
        /// </summary>
        [Output("role")]
        public Output<string?> Role { get; private set; } = null!;

        /// <summary>
        /// The user the key is created for.
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;


        /// <summary>
        /// Create a ApiKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApiKey(string name, ApiKeyArgs? args = null, CustomResourceOptions? options = null)
            : base("rockset:index/apiKey:ApiKey", name, args ?? new ApiKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApiKey(string name, Input<string> id, ApiKeyState? state = null, CustomResourceOptions? options = null)
            : base("rockset:index/apiKey:ApiKey", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "key",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ApiKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApiKey Get(string name, Input<string> id, ApiKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new ApiKey(name, id, state, options);
        }
    }

    public sealed class ApiKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the api key.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The role the api key will use. If not specified, "All User Assigned Roles" will be used.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public ApiKeyArgs()
        {
        }
        public static new ApiKeyArgs Empty => new ApiKeyArgs();
    }

    public sealed class ApiKeyState : global::Pulumi.ResourceArgs
    {
        [Input("key")]
        private Input<string>? _key;

        /// <summary>
        /// The resulting Rockset api key.
        /// </summary>
        public Input<string>? Key
        {
            get => _key;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _key = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Name of the api key.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The role the api key will use. If not specified, "All User Assigned Roles" will be used.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The user the key is created for.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public ApiKeyState()
        {
        }
        public static new ApiKeyState Empty => new ApiKeyState();
    }
}