// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rockset
{
    public static class GetUser
    {
        /// <summary>
        /// This data source can be used to fetch information about a specific user.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Rockset = Pulumi.Rockset;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var pme = Rockset.GetUser.Invoke(new()
        ///     {
        ///         Email = "pme@rockset.com",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUserResult> InvokeAsync(GetUserArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("rockset:index/getUser:getUser", args ?? new GetUserArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can be used to fetch information about a specific user.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Rockset = Pulumi.Rockset;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var pme = Rockset.GetUser.Invoke(new()
        ///     {
        ///         Email = "pme@rockset.com",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetUserResult> Invoke(GetUserInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserResult>("rockset:index/getUser:getUser", args ?? new GetUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// User email. If absent or blank, it gets the current user.
        /// </summary>
        [Input("email")]
        public string? Email { get; set; }

        public GetUserArgs()
        {
        }
        public static new GetUserArgs Empty => new GetUserArgs();
    }

    public sealed class GetUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// User email. If absent or blank, it gets the current user.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        public GetUserInvokeArgs()
        {
        }
        public static new GetUserInvokeArgs Empty => new GetUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserResult
    {
        /// <summary>
        /// User email. If absent or blank, it gets the current user.
        /// </summary>
        public readonly string? Email;
        /// <summary>
        /// User's first name.
        /// </summary>
        public readonly string FirstName;
        /// <summary>
        /// The user ID, in the form of the `email`.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// User's last name.
        /// </summary>
        public readonly string LastName;
        /// <summary>
        /// List of roles for the user. E.g. 'admin', 'member', 'read-only'.
        /// </summary>
        public readonly ImmutableArray<string> Roles;
        /// <summary>
        /// State of the user, either NEW or ACTIVE.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetUserResult(
            string? email,

            string firstName,

            string id,

            string lastName,

            ImmutableArray<string> roles,

            string state)
        {
            Email = email;
            FirstName = firstName;
            Id = id;
            LastName = lastName;
            Roles = roles;
            State = state;
        }
    }
}
