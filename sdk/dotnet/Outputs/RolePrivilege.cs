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
    public sealed class RolePrivilege
    {
        /// <summary>
        /// The action allowed by this privilege.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// Rockset cluster ID for which this action is allowed. Only valid for Workspace actions. Use '*ALL*' for actions which apply to all clusters.
        /// </summary>
        public readonly string? Cluster;
        /// <summary>
        /// The resource on which this action is allowed. Defaults to 'All' if not specified.
        /// </summary>
        public readonly string? ResourceName;

        [OutputConstructor]
        private RolePrivilege(
            string action,

            string? cluster,

            string? resourceName)
        {
            Action = action;
            Cluster = cluster;
            ResourceName = resourceName;
        }
    }
}
