// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rockset

import (
	"context"
	"reflect"

	"github.com/ddn0/pulumi-rockset/sdk/go/rockset/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a workspace.
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
//			_, err := rockset.LookupWorkspace(ctx, &rockset.LookupWorkspaceArgs{
//				Name: "demo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceResult
	err := ctx.Invoke("rockset:index/getWorkspace:getWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWorkspace.
type LookupWorkspaceArgs struct {
	// Workspace name.
	Name string `pulumi:"name"`
}

// A collection of values returned by getWorkspace.
type LookupWorkspaceResult struct {
	// Number of collections in the workspace.
	CollectionCount int `pulumi:"collectionCount"`
	// Created at in ISO-8601.
	CreatedAt string `pulumi:"createdAt"`
	// Created by.
	CreatedBy string `pulumi:"createdBy"`
	// Workspace description.
	Description string `pulumi:"description"`
	// The workspace `name`.
	Id string `pulumi:"id"`
	// Workspace name.
	Name string `pulumi:"name"`
}

func LookupWorkspaceOutput(ctx *pulumi.Context, args LookupWorkspaceOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceResult, error) {
			args := v.(LookupWorkspaceArgs)
			r, err := LookupWorkspace(ctx, &args, opts...)
			var s LookupWorkspaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceResultOutput)
}

// A collection of arguments for invoking getWorkspace.
type LookupWorkspaceOutputArgs struct {
	// Workspace name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupWorkspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceArgs)(nil)).Elem()
}

// A collection of values returned by getWorkspace.
type LookupWorkspaceResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceResult)(nil)).Elem()
}

func (o LookupWorkspaceResultOutput) ToLookupWorkspaceResultOutput() LookupWorkspaceResultOutput {
	return o
}

func (o LookupWorkspaceResultOutput) ToLookupWorkspaceResultOutputWithContext(ctx context.Context) LookupWorkspaceResultOutput {
	return o
}

// Number of collections in the workspace.
func (o LookupWorkspaceResultOutput) CollectionCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) int { return v.CollectionCount }).(pulumi.IntOutput)
}

// Created at in ISO-8601.
func (o LookupWorkspaceResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Created by.
func (o LookupWorkspaceResultOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.CreatedBy }).(pulumi.StringOutput)
}

// Workspace description.
func (o LookupWorkspaceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Description }).(pulumi.StringOutput)
}

// The workspace `name`.
func (o LookupWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Workspace name.
func (o LookupWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceResultOutput{})
}
