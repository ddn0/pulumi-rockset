// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rockset

import (
	"context"
	"reflect"

	"github.com/ddn0/pulumi-rockset/sdk/go/rockset/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can be used to fetch information about a specific user.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ddn0/pulumi-rockset/sdk/go/rockset"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rockset.LookupUser(ctx, &rockset.LookupUserArgs{
//				Email: pulumi.StringRef("pme@rockset.com"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserResult
	err := ctx.Invoke("rockset:index/getUser:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUser.
type LookupUserArgs struct {
	// User email. If absent or blank, it gets the current user.
	Email *string `pulumi:"email"`
}

// A collection of values returned by getUser.
type LookupUserResult struct {
	// User email. If absent or blank, it gets the current user.
	Email *string `pulumi:"email"`
	// User's first name.
	FirstName string `pulumi:"firstName"`
	// The user ID, in the form of the `email`.
	Id string `pulumi:"id"`
	// User's last name.
	LastName string `pulumi:"lastName"`
	// List of roles for the user. E.g. 'admin', 'member', 'read-only'.
	Roles []string `pulumi:"roles"`
	// State of the user, either NEW or ACTIVE.
	State string `pulumi:"state"`
}

func LookupUserOutput(ctx *pulumi.Context, args LookupUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserResult, error) {
			args := v.(LookupUserArgs)
			r, err := LookupUser(ctx, &args, opts...)
			var s LookupUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserResultOutput)
}

// A collection of arguments for invoking getUser.
type LookupUserOutputArgs struct {
	// User email. If absent or blank, it gets the current user.
	Email pulumi.StringPtrInput `pulumi:"email"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}

// A collection of values returned by getUser.
type LookupUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserResult)(nil)).Elem()
}

func (o LookupUserResultOutput) ToLookupUserResultOutput() LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToLookupUserResultOutputWithContext(ctx context.Context) LookupUserResultOutput {
	return o
}

// User email. If absent or blank, it gets the current user.
func (o LookupUserResultOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.Email }).(pulumi.StringPtrOutput)
}

// User's first name.
func (o LookupUserResultOutput) FirstName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.FirstName }).(pulumi.StringOutput)
}

// The user ID, in the form of the `email`.
func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// User's last name.
func (o LookupUserResultOutput) LastName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.LastName }).(pulumi.StringOutput)
}

// List of roles for the user. E.g. 'admin', 'member', 'read-only'.
func (o LookupUserResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

// State of the user, either NEW or ACTIVE.
func (o LookupUserResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}
