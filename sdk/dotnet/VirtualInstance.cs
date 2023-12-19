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
    /// Manages a Rockset Virtual Instance. To be able to create a new Virtual Instance,
    /// The main virtual instance must use a dedicated instance to create a secondary virtual instance,
    /// which must be SMALL or larger. To enable live mount, the secondary virtual instance must be MEDIUM or larger.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Rockset = Pulumi.Rockset;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var query = new Rockset.VirtualInstance("query", new()
    ///     {
    ///         Description = "vi for executing query lambdas",
    ///         Size = "MEDIUM",
    ///         RemountOnResume = true,
    ///     });
    /// 
    ///     var patch = new Rockset.CollectionMount("patch", new()
    ///     {
    ///         VirtualInstanceId = query.Id,
    ///         Path = "commons.data",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import rockset:index/virtualInstance:VirtualInstance query 29e4a43c-fff4-4fe6-80e3-1ee57bc22e82
    /// ```
    /// </summary>
    [RocksetResourceType("rockset:index/virtualInstance:VirtualInstance")]
    public partial class VirtualInstance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Number of seconds without queries after which the Virtual Instance is suspended.
        /// </summary>
        [Output("autoSuspendSeconds")]
        public Output<int?> AutoSuspendSeconds { get; private set; } = null!;

        /// <summary>
        /// Current size of the virtual instance.
        /// </summary>
        [Output("currentSize")]
        public Output<string> CurrentSize { get; private set; } = null!;

        /// <summary>
        /// Is this Virtual Instance the default.
        /// </summary>
        [Output("default")]
        public Output<bool> Default { get; private set; } = null!;

        /// <summary>
        /// Description of the virtual instance.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Desired size of the virtual instance.
        /// </summary>
        [Output("desiredSize")]
        public Output<string> DesiredSize { get; private set; } = null!;

        /// <summary>
        /// Is monitoring enabled for this Virtual Instance.
        /// </summary>
        [Output("monitoringEnabled")]
        public Output<bool> MonitoringEnabled { get; private set; } = null!;

        /// <summary>
        /// Number of seconds between data refreshes for mounts on this Virtual Instance. A value of 0 means continuous refresh and a value of null means never refresh.
        /// </summary>
        [Output("mountRefreshIntervalSeconds")]
        public Output<int?> MountRefreshIntervalSeconds { get; private set; } = null!;

        /// <summary>
        /// Name of the virtual instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// When a Virtual Instance is resumed, remount all collections that were mounted when the Virtual Instance was suspended.
        /// </summary>
        [Output("remountOnResume")]
        public Output<bool?> RemountOnResume { get; private set; } = null!;

        /// <summary>
        /// RRN of this Virtual Instance.
        /// </summary>
        [Output("rrn")]
        public Output<string> Rrn { get; private set; } = null!;

        /// <summary>
        /// Requested virtual instance size. Note that this field is called type in the API documentation.
        /// </summary>
        [Output("size")]
        public Output<string> Size { get; private set; } = null!;

        /// <summary>
        /// Virtual Instance state.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualInstance(string name, VirtualInstanceArgs args, CustomResourceOptions? options = null)
            : base("rockset:index/virtualInstance:VirtualInstance", name, args ?? new VirtualInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualInstance(string name, Input<string> id, VirtualInstanceState? state = null, CustomResourceOptions? options = null)
            : base("rockset:index/virtualInstance:VirtualInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualInstance Get(string name, Input<string> id, VirtualInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualInstance(name, id, state, options);
        }
    }

    public sealed class VirtualInstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of seconds without queries after which the Virtual Instance is suspended.
        /// </summary>
        [Input("autoSuspendSeconds")]
        public Input<int>? AutoSuspendSeconds { get; set; }

        /// <summary>
        /// Description of the virtual instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Number of seconds between data refreshes for mounts on this Virtual Instance. A value of 0 means continuous refresh and a value of null means never refresh.
        /// </summary>
        [Input("mountRefreshIntervalSeconds")]
        public Input<int>? MountRefreshIntervalSeconds { get; set; }

        /// <summary>
        /// Name of the virtual instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// When a Virtual Instance is resumed, remount all collections that were mounted when the Virtual Instance was suspended.
        /// </summary>
        [Input("remountOnResume")]
        public Input<bool>? RemountOnResume { get; set; }

        /// <summary>
        /// Requested virtual instance size. Note that this field is called type in the API documentation.
        /// </summary>
        [Input("size", required: true)]
        public Input<string> Size { get; set; } = null!;

        public VirtualInstanceArgs()
        {
        }
        public static new VirtualInstanceArgs Empty => new VirtualInstanceArgs();
    }

    public sealed class VirtualInstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of seconds without queries after which the Virtual Instance is suspended.
        /// </summary>
        [Input("autoSuspendSeconds")]
        public Input<int>? AutoSuspendSeconds { get; set; }

        /// <summary>
        /// Current size of the virtual instance.
        /// </summary>
        [Input("currentSize")]
        public Input<string>? CurrentSize { get; set; }

        /// <summary>
        /// Is this Virtual Instance the default.
        /// </summary>
        [Input("default")]
        public Input<bool>? Default { get; set; }

        /// <summary>
        /// Description of the virtual instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Desired size of the virtual instance.
        /// </summary>
        [Input("desiredSize")]
        public Input<string>? DesiredSize { get; set; }

        /// <summary>
        /// Is monitoring enabled for this Virtual Instance.
        /// </summary>
        [Input("monitoringEnabled")]
        public Input<bool>? MonitoringEnabled { get; set; }

        /// <summary>
        /// Number of seconds between data refreshes for mounts on this Virtual Instance. A value of 0 means continuous refresh and a value of null means never refresh.
        /// </summary>
        [Input("mountRefreshIntervalSeconds")]
        public Input<int>? MountRefreshIntervalSeconds { get; set; }

        /// <summary>
        /// Name of the virtual instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// When a Virtual Instance is resumed, remount all collections that were mounted when the Virtual Instance was suspended.
        /// </summary>
        [Input("remountOnResume")]
        public Input<bool>? RemountOnResume { get; set; }

        /// <summary>
        /// RRN of this Virtual Instance.
        /// </summary>
        [Input("rrn")]
        public Input<string>? Rrn { get; set; }

        /// <summary>
        /// Requested virtual instance size. Note that this field is called type in the API documentation.
        /// </summary>
        [Input("size")]
        public Input<string>? Size { get; set; }

        /// <summary>
        /// Virtual Instance state.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public VirtualInstanceState()
        {
        }
        public static new VirtualInstanceState Empty => new VirtualInstanceState();
    }
}