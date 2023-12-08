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
    /// Manages a Rockset MongoDB Integration.
    /// </summary>
    [RocksetResourceType("rockset:index/mongodbIntegration:MongodbIntegration")]
    public partial class MongodbIntegration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// MongoDB connection URI string. The password is scrubbed from the URI when fetched by the API so this field is NOT set on imports and reads.
        /// </summary>
        [Output("connectionUri")]
        public Output<string> ConnectionUri { get; private set; } = null!;

        /// <summary>
        /// Text describing the integration.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Unique identifier for the integration. Can contain alphanumeric or dash characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a MongodbIntegration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MongodbIntegration(string name, MongodbIntegrationArgs args, CustomResourceOptions? options = null)
            : base("rockset:index/mongodbIntegration:MongodbIntegration", name, args ?? new MongodbIntegrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MongodbIntegration(string name, Input<string> id, MongodbIntegrationState? state = null, CustomResourceOptions? options = null)
            : base("rockset:index/mongodbIntegration:MongodbIntegration", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "connectionUri",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MongodbIntegration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MongodbIntegration Get(string name, Input<string> id, MongodbIntegrationState? state = null, CustomResourceOptions? options = null)
        {
            return new MongodbIntegration(name, id, state, options);
        }
    }

    public sealed class MongodbIntegrationArgs : global::Pulumi.ResourceArgs
    {
        [Input("connectionUri", required: true)]
        private Input<string>? _connectionUri;

        /// <summary>
        /// MongoDB connection URI string. The password is scrubbed from the URI when fetched by the API so this field is NOT set on imports and reads.
        /// </summary>
        public Input<string>? ConnectionUri
        {
            get => _connectionUri;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _connectionUri = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Text describing the integration.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Unique identifier for the integration. Can contain alphanumeric or dash characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public MongodbIntegrationArgs()
        {
        }
        public static new MongodbIntegrationArgs Empty => new MongodbIntegrationArgs();
    }

    public sealed class MongodbIntegrationState : global::Pulumi.ResourceArgs
    {
        [Input("connectionUri")]
        private Input<string>? _connectionUri;

        /// <summary>
        /// MongoDB connection URI string. The password is scrubbed from the URI when fetched by the API so this field is NOT set on imports and reads.
        /// </summary>
        public Input<string>? ConnectionUri
        {
            get => _connectionUri;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _connectionUri = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Text describing the integration.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Unique identifier for the integration. Can contain alphanumeric or dash characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public MongodbIntegrationState()
        {
        }
        public static new MongodbIntegrationState Empty => new MongodbIntegrationState();
    }
}
