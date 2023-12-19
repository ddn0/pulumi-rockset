// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rockset.Inputs
{

    public sealed class DynamodbCollectionSourceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AWS region name of DynamoDB table, by default us-west-2 is used.
        /// </summary>
        [Input("awsRegion")]
        public Input<string>? AwsRegion { get; set; }

        /// <summary>
        /// The name of the Rockset DynamoDB integration.
        /// </summary>
        [Input("integrationName", required: true)]
        public Input<string> IntegrationName { get; set; } = null!;

        /// <summary>
        /// Max RCU usage for scan.
        /// </summary>
        [Input("rcu")]
        public Input<int>? Rcu { get; set; }

        /// <summary>
        /// DynamoDB scan end time.
        /// </summary>
        [Input("scanEndTime")]
        public Input<string>? ScanEndTime { get; set; }

        /// <summary>
        /// Number of records inserted using scan.
        /// </summary>
        [Input("scanRecordsProcessed")]
        public Input<int>? ScanRecordsProcessed { get; set; }

        /// <summary>
        /// DynamoDB scan start time.
        /// </summary>
        [Input("scanStartTime")]
        public Input<string>? ScanStartTime { get; set; }

        /// <summary>
        /// Number of records in DynamoDB table at time of scan.
        /// </summary>
        [Input("scanTotalRecords")]
        public Input<int>? ScanTotalRecords { get; set; }

        /// <summary>
        /// State of current ingest for this table.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// ISO-8601 date when source was last processed.
        /// </summary>
        [Input("streamLastProcessedAt")]
        public Input<string>? StreamLastProcessedAt { get; set; }

        /// <summary>
        /// Name of DynamoDB table containing data.
        /// </summary>
        [Input("tableName", required: true)]
        public Input<string> TableName { get; set; } = null!;

        /// <summary>
        /// Whether the initial table scan should use the DynamoDB scan API. If false, export will be performed using an S3 bucket.
        /// </summary>
        [Input("useScanApi")]
        public Input<bool>? UseScanApi { get; set; }

        public DynamodbCollectionSourceGetArgs()
        {
        }
        public static new DynamodbCollectionSourceGetArgs Empty => new DynamodbCollectionSourceGetArgs();
    }
}