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
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Null = Pulumi.Null;
    /// using Rockset = Pulumi.Rockset;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var rocksetApiserver = config.Get("rocksetApiserver") ?? "https://api.usw2a1.rockset.com";
    ///     var kafkaConnect = config.Get("kafkaConnect") ?? "localhost:8083";
    ///     var topics = config.RequireObject&lt;dynamic&gt;("topics");
    ///     var maxTasks = config.GetDouble("maxTasks") ?? 10;
    ///     var local = new Rockset.KafkaIntegration("local", new()
    ///     {
    ///         Description = "Integration to ingest from a local kafka.",
    ///         KafkaDataFormat = "JSON",
    ///         KafkaTopicNames = new[]
    ///         {
    ///             "foo",
    ///             "bar",
    ///         },
    ///         WaitForIntegration = false,
    ///     });
    /// 
    ///     var configure_kafka_connect = new Null.Resource("configure-kafka-connect");
    /// 
    /// });
    /// ```
    /// </summary>
    [RocksetResourceType("rockset:index/kafkaIntegration:KafkaIntegration")]
    public partial class KafkaIntegration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Kafka bootstrap server url(s). Required only for V3 integration.
        /// </summary>
        [Output("bootstrapServers")]
        public Output<string?> BootstrapServers { get; private set; } = null!;

        /// <summary>
        /// Kafka connection string.
        /// </summary>
        [Output("connectionString")]
        public Output<string> ConnectionString { get; private set; } = null!;

        /// <summary>
        /// Text describing the integration.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The format of the Kafka topics being tailed. Possible values: JSON, AVRO.
        /// </summary>
        [Output("kafkaDataFormat")]
        public Output<string?> KafkaDataFormat { get; private set; } = null!;

        /// <summary>
        /// Kafka topics to tail.
        /// </summary>
        [Output("kafkaTopicNames")]
        public Output<ImmutableArray<string>> KafkaTopicNames { get; private set; } = null!;

        /// <summary>
        /// Unique identifier for the integration. Can contain alphanumeric or dash characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Kafka configuration for schema registry. Required only for V3 integration.
        /// </summary>
        [Output("schemaRegistryConfig")]
        public Output<ImmutableDictionary<string, string>?> SchemaRegistryConfig { get; private set; } = null!;

        /// <summary>
        /// Kafka security configurations. Required only for V3 integration.
        /// </summary>
        [Output("securityConfig")]
        public Output<ImmutableDictionary<string, string>?> SecurityConfig { get; private set; } = null!;

        /// <summary>
        /// Use v3 for Confluent Cloud.
        /// </summary>
        [Output("useV3")]
        public Output<bool?> UseV3 { get; private set; } = null!;

        /// <summary>
        /// Wait until the integration is active.
        /// </summary>
        [Output("waitForIntegration")]
        public Output<bool?> WaitForIntegration { get; private set; } = null!;


        /// <summary>
        /// Create a KafkaIntegration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KafkaIntegration(string name, KafkaIntegrationArgs? args = null, CustomResourceOptions? options = null)
            : base("rockset:index/kafkaIntegration:KafkaIntegration", name, args ?? new KafkaIntegrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KafkaIntegration(string name, Input<string> id, KafkaIntegrationState? state = null, CustomResourceOptions? options = null)
            : base("rockset:index/kafkaIntegration:KafkaIntegration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KafkaIntegration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KafkaIntegration Get(string name, Input<string> id, KafkaIntegrationState? state = null, CustomResourceOptions? options = null)
        {
            return new KafkaIntegration(name, id, state, options);
        }
    }

    public sealed class KafkaIntegrationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Kafka bootstrap server url(s). Required only for V3 integration.
        /// </summary>
        [Input("bootstrapServers")]
        public Input<string>? BootstrapServers { get; set; }

        /// <summary>
        /// Kafka connection string.
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

        /// <summary>
        /// Text describing the integration.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The format of the Kafka topics being tailed. Possible values: JSON, AVRO.
        /// </summary>
        [Input("kafkaDataFormat")]
        public Input<string>? KafkaDataFormat { get; set; }

        [Input("kafkaTopicNames")]
        private InputList<string>? _kafkaTopicNames;

        /// <summary>
        /// Kafka topics to tail.
        /// </summary>
        public InputList<string> KafkaTopicNames
        {
            get => _kafkaTopicNames ?? (_kafkaTopicNames = new InputList<string>());
            set => _kafkaTopicNames = value;
        }

        /// <summary>
        /// Unique identifier for the integration. Can contain alphanumeric or dash characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("schemaRegistryConfig")]
        private InputMap<string>? _schemaRegistryConfig;

        /// <summary>
        /// Kafka configuration for schema registry. Required only for V3 integration.
        /// </summary>
        public InputMap<string> SchemaRegistryConfig
        {
            get => _schemaRegistryConfig ?? (_schemaRegistryConfig = new InputMap<string>());
            set => _schemaRegistryConfig = value;
        }

        [Input("securityConfig")]
        private InputMap<string>? _securityConfig;

        /// <summary>
        /// Kafka security configurations. Required only for V3 integration.
        /// </summary>
        public InputMap<string> SecurityConfig
        {
            get => _securityConfig ?? (_securityConfig = new InputMap<string>());
            set => _securityConfig = value;
        }

        /// <summary>
        /// Use v3 for Confluent Cloud.
        /// </summary>
        [Input("useV3")]
        public Input<bool>? UseV3 { get; set; }

        /// <summary>
        /// Wait until the integration is active.
        /// </summary>
        [Input("waitForIntegration")]
        public Input<bool>? WaitForIntegration { get; set; }

        public KafkaIntegrationArgs()
        {
        }
        public static new KafkaIntegrationArgs Empty => new KafkaIntegrationArgs();
    }

    public sealed class KafkaIntegrationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Kafka bootstrap server url(s). Required only for V3 integration.
        /// </summary>
        [Input("bootstrapServers")]
        public Input<string>? BootstrapServers { get; set; }

        /// <summary>
        /// Kafka connection string.
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

        /// <summary>
        /// Text describing the integration.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The format of the Kafka topics being tailed. Possible values: JSON, AVRO.
        /// </summary>
        [Input("kafkaDataFormat")]
        public Input<string>? KafkaDataFormat { get; set; }

        [Input("kafkaTopicNames")]
        private InputList<string>? _kafkaTopicNames;

        /// <summary>
        /// Kafka topics to tail.
        /// </summary>
        public InputList<string> KafkaTopicNames
        {
            get => _kafkaTopicNames ?? (_kafkaTopicNames = new InputList<string>());
            set => _kafkaTopicNames = value;
        }

        /// <summary>
        /// Unique identifier for the integration. Can contain alphanumeric or dash characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("schemaRegistryConfig")]
        private InputMap<string>? _schemaRegistryConfig;

        /// <summary>
        /// Kafka configuration for schema registry. Required only for V3 integration.
        /// </summary>
        public InputMap<string> SchemaRegistryConfig
        {
            get => _schemaRegistryConfig ?? (_schemaRegistryConfig = new InputMap<string>());
            set => _schemaRegistryConfig = value;
        }

        [Input("securityConfig")]
        private InputMap<string>? _securityConfig;

        /// <summary>
        /// Kafka security configurations. Required only for V3 integration.
        /// </summary>
        public InputMap<string> SecurityConfig
        {
            get => _securityConfig ?? (_securityConfig = new InputMap<string>());
            set => _securityConfig = value;
        }

        /// <summary>
        /// Use v3 for Confluent Cloud.
        /// </summary>
        [Input("useV3")]
        public Input<bool>? UseV3 { get; set; }

        /// <summary>
        /// Wait until the integration is active.
        /// </summary>
        [Input("waitForIntegration")]
        public Input<bool>? WaitForIntegration { get; set; }

        public KafkaIntegrationState()
        {
        }
        public static new KafkaIntegrationState Empty => new KafkaIntegrationState();
    }
}
