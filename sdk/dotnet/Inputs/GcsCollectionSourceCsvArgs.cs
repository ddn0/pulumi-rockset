// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rockset.Inputs
{

    public sealed class GcsCollectionSourceCsvArgs : global::Pulumi.ResourceArgs
    {
        [Input("columnNames")]
        private InputList<string>? _columnNames;

        /// <summary>
        /// The names of the columns.
        /// </summary>
        public InputList<string> ColumnNames
        {
            get => _columnNames ?? (_columnNames = new InputList<string>());
            set => _columnNames = value;
        }

        [Input("columnTypes")]
        private InputList<string>? _columnTypes;

        /// <summary>
        /// The types of the columns.
        /// </summary>
        public InputList<string> ColumnTypes
        {
            get => _columnTypes ?? (_columnTypes = new InputList<string>());
            set => _columnTypes = value;
        }

        /// <summary>
        /// Can be one of: UTF-8, ISO*8859*1, UTF-16.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        /// <summary>
        /// Escape character removes any special meaning from the character that follows it. Defaults to backslash.
        /// </summary>
        [Input("escapeChar")]
        public Input<string>? EscapeChar { get; set; }

        /// <summary>
        /// If the first line in every object specifies the column names.
        /// </summary>
        [Input("firstLineAsColumnNames")]
        public Input<bool>? FirstLineAsColumnNames { get; set; }

        /// <summary>
        /// Character within which a cell value is enclosed. Defaults to double quote.
        /// </summary>
        [Input("quoteChar")]
        public Input<string>? QuoteChar { get; set; }

        /// <summary>
        /// A single character that is the column separator.
        /// </summary>
        [Input("separator")]
        public Input<string>? Separator { get; set; }

        public GcsCollectionSourceCsvArgs()
        {
        }
        public static new GcsCollectionSourceCsvArgs Empty => new GcsCollectionSourceCsvArgs();
    }
}
