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
    /// Manages a collection with an Kinesis source attached.
    /// </summary>
    [RocksetResourceType("rockset:index/kinesisCollection:KinesisCollection")]
    public partial class KinesisCollection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Text describing the collection.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Ingest transformation SQL query. Turns the collection into insert_only mode.
        /// </summary>
        [Output("ingestTransformation")]
        public Output<string?> IngestTransformation { get; private set; } = null!;

        /// <summary>
        /// Unique identifier for the collection. Can contain alphanumeric or dash characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Number of seconds after which data is purged. Based on event time.
        /// </summary>
        [Output("retentionSecs")]
        public Output<int?> RetentionSecs { get; private set; } = null!;

        /// <summary>
        /// Defines a source for this collection.
        /// </summary>
        [Output("sources")]
        public Output<ImmutableArray<Outputs.KinesisCollectionSource>> Sources { get; private set; } = null!;

        /// <summary>
        /// Wait until the collection is ready.
        /// </summary>
        [Output("waitForCollection")]
        public Output<bool?> WaitForCollection { get; private set; } = null!;

        /// <summary>
        /// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
        /// </summary>
        [Output("waitForDocuments")]
        public Output<int?> WaitForDocuments { get; private set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Output("workspace")]
        public Output<string> Workspace { get; private set; } = null!;


        /// <summary>
        /// Create a KinesisCollection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KinesisCollection(string name, KinesisCollectionArgs args, CustomResourceOptions? options = null)
            : base("rockset:index/kinesisCollection:KinesisCollection", name, args ?? new KinesisCollectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KinesisCollection(string name, Input<string> id, KinesisCollectionState? state = null, CustomResourceOptions? options = null)
            : base("rockset:index/kinesisCollection:KinesisCollection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KinesisCollection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KinesisCollection Get(string name, Input<string> id, KinesisCollectionState? state = null, CustomResourceOptions? options = null)
        {
            return new KinesisCollection(name, id, state, options);
        }
    }

    public sealed class KinesisCollectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Text describing the collection.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Ingest transformation SQL query. Turns the collection into insert_only mode.
        /// </summary>
        [Input("ingestTransformation")]
        public Input<string>? IngestTransformation { get; set; }

        /// <summary>
        /// Unique identifier for the collection. Can contain alphanumeric or dash characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Number of seconds after which data is purged. Based on event time.
        /// </summary>
        [Input("retentionSecs")]
        public Input<int>? RetentionSecs { get; set; }

        [Input("sources")]
        private InputList<Inputs.KinesisCollectionSourceArgs>? _sources;

        /// <summary>
        /// Defines a source for this collection.
        /// </summary>
        public InputList<Inputs.KinesisCollectionSourceArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.KinesisCollectionSourceArgs>());
            set => _sources = value;
        }

        /// <summary>
        /// Wait until the collection is ready.
        /// </summary>
        [Input("waitForCollection")]
        public Input<bool>? WaitForCollection { get; set; }

        /// <summary>
        /// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
        /// </summary>
        [Input("waitForDocuments")]
        public Input<int>? WaitForDocuments { get; set; }

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspace", required: true)]
        public Input<string> Workspace { get; set; } = null!;

        public KinesisCollectionArgs()
        {
        }
        public static new KinesisCollectionArgs Empty => new KinesisCollectionArgs();
    }

    public sealed class KinesisCollectionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Text describing the collection.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Ingest transformation SQL query. Turns the collection into insert_only mode.
        /// </summary>
        [Input("ingestTransformation")]
        public Input<string>? IngestTransformation { get; set; }

        /// <summary>
        /// Unique identifier for the collection. Can contain alphanumeric or dash characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Number of seconds after which data is purged. Based on event time.
        /// </summary>
        [Input("retentionSecs")]
        public Input<int>? RetentionSecs { get; set; }

        [Input("sources")]
        private InputList<Inputs.KinesisCollectionSourceGetArgs>? _sources;

        /// <summary>
        /// Defines a source for this collection.
        /// </summary>
        public InputList<Inputs.KinesisCollectionSourceGetArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.KinesisCollectionSourceGetArgs>());
            set => _sources = value;
        }

        /// <summary>
        /// Wait until the collection is ready.
        /// </summary>
        [Input("waitForCollection")]
        public Input<bool>? WaitForCollection { get; set; }

        /// <summary>
        /// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
        /// </summary>
        [Input("waitForDocuments")]
        public Input<int>? WaitForDocuments { get; set; }

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspace")]
        public Input<string>? Workspace { get; set; }

        public KinesisCollectionState()
        {
        }
        public static new KinesisCollectionState Empty => new KinesisCollectionState();
    }
}