// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rockset.Inputs
{

    public sealed class MongodbCollectionSourceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// MongoDB collection name of the target collection.
        /// </summary>
        [Input("collectionName", required: true)]
        public Input<string> CollectionName { get; set; } = null!;

        /// <summary>
        /// MongoDB database name containing the target collection.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name of the Rockset MongoDB integration.
        /// </summary>
        [Input("integrationName", required: true)]
        public Input<string> IntegrationName { get; set; } = null!;

        /// <summary>
        /// Whether to get the full document from the MongoDB change stream to enable multi-field expression transformations.
        /// Selecting this option will increase load on your upstream MongoDB database.
        /// </summary>
        [Input("retrieveFullDocument")]
        public Input<bool>? RetrieveFullDocument { get; set; }

        /// <summary>
        /// MongoDB scan end time.
        /// </summary>
        [Input("scanEndTime")]
        public Input<string>? ScanEndTime { get; set; }

        /// <summary>
        /// Number of records inserted using scan.
        /// </summary>
        [Input("scanRecordsProcessed")]
        public Input<int>? ScanRecordsProcessed { get; set; }

        /// <summary>
        /// MongoDB scan start time.
        /// </summary>
        [Input("scanStartTime")]
        public Input<string>? ScanStartTime { get; set; }

        /// <summary>
        /// Number of records in MongoDB table at time of scan.
        /// </summary>
        [Input("scanTotalRecords")]
        public Input<int>? ScanTotalRecords { get; set; }

        /// <summary>
        /// State of current ingest for this table.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// ISO-8601 date when delete from source was last processed.
        /// </summary>
        [Input("streamLastDeleteProcessedAt")]
        public Input<string>? StreamLastDeleteProcessedAt { get; set; }

        /// <summary>
        /// ISO-8601 date when new insert from source was last processed.
        /// </summary>
        [Input("streamLastInsertProcessedAt")]
        public Input<string>? StreamLastInsertProcessedAt { get; set; }

        /// <summary>
        /// ISO-8601 date when update from source was last processed.
        /// </summary>
        [Input("streamLastUpdateProcessedAt")]
        public Input<string>? StreamLastUpdateProcessedAt { get; set; }

        /// <summary>
        /// Number of new records deleted using stream.
        /// </summary>
        [Input("streamRecordsDeleted")]
        public Input<int>? StreamRecordsDeleted { get; set; }

        /// <summary>
        /// Number of new records inserted using stream.
        /// </summary>
        [Input("streamRecordsInserted")]
        public Input<int>? StreamRecordsInserted { get; set; }

        /// <summary>
        /// Number of new records updated using stream.
        /// </summary>
        [Input("streamRecordsUpdated")]
        public Input<int>? StreamRecordsUpdated { get; set; }

        public MongodbCollectionSourceGetArgs()
        {
        }
        public static new MongodbCollectionSourceGetArgs Empty => new MongodbCollectionSourceGetArgs();
    }
}
