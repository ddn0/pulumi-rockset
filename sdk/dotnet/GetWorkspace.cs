// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rockset
{
    public static class GetWorkspace
    {
        /// <summary>
        /// Gets information about a workspace.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Rockset = Pulumi.Rockset;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var demo = Rockset.GetWorkspace.Invoke(new()
        ///     {
        ///         Name = "demo",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetWorkspaceResult> InvokeAsync(GetWorkspaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWorkspaceResult>("rockset:index/getWorkspace:getWorkspace", args ?? new GetWorkspaceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a workspace.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Rockset = Pulumi.Rockset;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var demo = Rockset.GetWorkspace.Invoke(new()
        ///     {
        ///         Name = "demo",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetWorkspaceResult> Invoke(GetWorkspaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWorkspaceResult>("rockset:index/getWorkspace:getWorkspace", args ?? new GetWorkspaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWorkspaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Workspace name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetWorkspaceArgs()
        {
        }
        public static new GetWorkspaceArgs Empty => new GetWorkspaceArgs();
    }

    public sealed class GetWorkspaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Workspace name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetWorkspaceInvokeArgs()
        {
        }
        public static new GetWorkspaceInvokeArgs Empty => new GetWorkspaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetWorkspaceResult
    {
        /// <summary>
        /// Number of collections in the workspace.
        /// </summary>
        public readonly int CollectionCount;
        /// <summary>
        /// Created at in ISO-8601.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Created by.
        /// </summary>
        public readonly string CreatedBy;
        /// <summary>
        /// Workspace description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The workspace `name`.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Workspace name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetWorkspaceResult(
            int collectionCount,

            string createdAt,

            string createdBy,

            string description,

            string id,

            string name)
        {
            CollectionCount = collectionCount;
            CreatedAt = createdAt;
            CreatedBy = createdBy;
            Description = description;
            Id = id;
            Name = name;
        }
    }
}
