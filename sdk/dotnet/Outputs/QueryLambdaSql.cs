// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rockset.Outputs
{

    [OutputType]
    public sealed class QueryLambdaSql
    {
        public readonly ImmutableArray<Outputs.QueryLambdaSqlDefaultParameter> DefaultParameters;
        public readonly string Query;

        [OutputConstructor]
        private QueryLambdaSql(
            ImmutableArray<Outputs.QueryLambdaSqlDefaultParameter> defaultParameters,

            string query)
        {
            DefaultParameters = defaultParameters;
            Query = query;
        }
    }
}
