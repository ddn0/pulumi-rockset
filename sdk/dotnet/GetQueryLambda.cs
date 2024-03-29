// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rockset
{
    public static class GetQueryLambda
    {
        /// <summary>
        /// Gets information about a query lambda. The `tag` defaults to `latest`.
        /// </summary>
        public static Task<GetQueryLambdaResult> InvokeAsync(GetQueryLambdaArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetQueryLambdaResult>("rockset:index/getQueryLambda:getQueryLambda", args ?? new GetQueryLambdaArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a query lambda. The `tag` defaults to `latest`.
        /// </summary>
        public static Output<GetQueryLambdaResult> Invoke(GetQueryLambdaInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetQueryLambdaResult>("rockset:index/getQueryLambda:getQueryLambda", args ?? new GetQueryLambdaInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetQueryLambdaArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the query lambda.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Tag name.
        /// </summary>
        [Input("tag")]
        public string? Tag { get; set; }

        /// <summary>
        /// Workspace the query lambda resides in.
        /// </summary>
        [Input("workspace", required: true)]
        public string Workspace { get; set; } = null!;

        public GetQueryLambdaArgs()
        {
        }
        public static new GetQueryLambdaArgs Empty => new GetQueryLambdaArgs();
    }

    public sealed class GetQueryLambdaInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the query lambda.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Tag name.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// Workspace the query lambda resides in.
        /// </summary>
        [Input("workspace", required: true)]
        public Input<string> Workspace { get; set; } = null!;

        public GetQueryLambdaInvokeArgs()
        {
        }
        public static new GetQueryLambdaInvokeArgs Empty => new GetQueryLambdaInvokeArgs();
    }


    [OutputType]
    public sealed class GetQueryLambdaResult
    {
        /// <summary>
        /// Description of the query lambda.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Last time the query lambda was executed.
        /// </summary>
        public readonly string LastExecuted;
        /// <summary>
        /// Name of the query lambda.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Query lambda SQL.
        /// </summary>
        public readonly string Sql;
        /// <summary>
        /// Tag name.
        /// </summary>
        public readonly string? Tag;
        /// <summary>
        /// Query lambda tag version.
        /// </summary>
        public readonly string Version;
        /// <summary>
        /// Workspace the query lambda resides in.
        /// </summary>
        public readonly string Workspace;

        [OutputConstructor]
        private GetQueryLambdaResult(
            string description,

            string id,

            string lastExecuted,

            string name,

            string sql,

            string? tag,

            string version,

            string workspace)
        {
            Description = description;
            Id = id;
            LastExecuted = lastExecuted;
            Name = name;
            Sql = sql;
            Tag = tag;
            Version = version;
            Workspace = workspace;
        }
    }
}
