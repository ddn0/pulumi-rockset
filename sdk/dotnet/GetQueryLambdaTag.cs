// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rockset
{
    public static class GetQueryLambdaTag
    {
        /// <summary>
        /// Deprecated. Use `rockset.QueryLambda` instead and specify the `tag`.
        /// </summary>
        public static Task<GetQueryLambdaTagResult> InvokeAsync(GetQueryLambdaTagArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetQueryLambdaTagResult>("rockset:index/getQueryLambdaTag:getQueryLambdaTag", args ?? new GetQueryLambdaTagArgs(), options.WithDefaults());

        /// <summary>
        /// Deprecated. Use `rockset.QueryLambda` instead and specify the `tag`.
        /// </summary>
        public static Output<GetQueryLambdaTagResult> Invoke(GetQueryLambdaTagInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetQueryLambdaTagResult>("rockset:index/getQueryLambdaTag:getQueryLambdaTag", args ?? new GetQueryLambdaTagInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetQueryLambdaTagArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the query lambda.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Tag name.
        /// </summary>
        [Input("tag", required: true)]
        public string Tag { get; set; } = null!;

        /// <summary>
        /// Workspace the query lambda resides in.
        /// </summary>
        [Input("workspace", required: true)]
        public string Workspace { get; set; } = null!;

        public GetQueryLambdaTagArgs()
        {
        }
        public static new GetQueryLambdaTagArgs Empty => new GetQueryLambdaTagArgs();
    }

    public sealed class GetQueryLambdaTagInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the query lambda.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Tag name.
        /// </summary>
        [Input("tag", required: true)]
        public Input<string> Tag { get; set; } = null!;

        /// <summary>
        /// Workspace the query lambda resides in.
        /// </summary>
        [Input("workspace", required: true)]
        public Input<string> Workspace { get; set; } = null!;

        public GetQueryLambdaTagInvokeArgs()
        {
        }
        public static new GetQueryLambdaTagInvokeArgs Empty => new GetQueryLambdaTagInvokeArgs();
    }


    [OutputType]
    public sealed class GetQueryLambdaTagResult
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
        public readonly string Tag;
        /// <summary>
        /// Query lambda tag version.
        /// </summary>
        public readonly string Version;
        /// <summary>
        /// Workspace the query lambda resides in.
        /// </summary>
        public readonly string Workspace;

        [OutputConstructor]
        private GetQueryLambdaTagResult(
            string description,

            string id,

            string lastExecuted,

            string name,

            string sql,

            string tag,

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
