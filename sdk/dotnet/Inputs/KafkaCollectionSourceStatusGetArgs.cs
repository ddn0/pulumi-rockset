// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rockset.Inputs
{

    public sealed class KafkaCollectionSourceStatusGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("documentsProcessed")]
        public Input<int>? DocumentsProcessed { get; set; }

        [Input("lastConsumedTime")]
        public Input<string>? LastConsumedTime { get; set; }

        [Input("partitions")]
        private InputList<Inputs.KafkaCollectionSourceStatusPartitionGetArgs>? _partitions;
        public InputList<Inputs.KafkaCollectionSourceStatusPartitionGetArgs> Partitions
        {
            get => _partitions ?? (_partitions = new InputList<Inputs.KafkaCollectionSourceStatusPartitionGetArgs>());
            set => _partitions = value;
        }

        [Input("state")]
        public Input<string>? State { get; set; }

        public KafkaCollectionSourceStatusGetArgs()
        {
        }
        public static new KafkaCollectionSourceStatusGetArgs Empty => new KafkaCollectionSourceStatusGetArgs();
    }
}