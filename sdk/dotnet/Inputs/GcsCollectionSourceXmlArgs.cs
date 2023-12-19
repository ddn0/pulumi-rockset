// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rockset.Inputs
{

    public sealed class GcsCollectionSourceXmlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Tag to differentiate between attributes and elements.
        /// </summary>
        [Input("attributePrefix")]
        public Input<string>? AttributePrefix { get; set; }

        /// <summary>
        /// Tags with which documents are identified
        /// </summary>
        [Input("docTag")]
        public Input<string>? DocTag { get; set; }

        /// <summary>
        /// Can be one of: UTF-8, ISO*8859*1, UTF-16.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        /// <summary>
        /// Tag until which xml is ignored.
        /// </summary>
        [Input("rootTag")]
        public Input<string>? RootTag { get; set; }

        /// <summary>
        /// Tag used for the value when there are attributes in the element having no child.
        /// </summary>
        [Input("valueTag")]
        public Input<string>? ValueTag { get; set; }

        public GcsCollectionSourceXmlArgs()
        {
        }
        public static new GcsCollectionSourceXmlArgs Empty => new GcsCollectionSourceXmlArgs();
    }
}