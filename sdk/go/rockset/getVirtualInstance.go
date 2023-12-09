// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rockset

import (
	"context"
	"reflect"

	"github.com/ddn0/pulumi-rockset/sdk/go/rockset/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualInstance(ctx *pulumi.Context, args *LookupVirtualInstanceArgs, opts ...pulumi.InvokeOption) (*LookupVirtualInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualInstanceResult
	err := ctx.Invoke("rockset:index/getVirtualInstance:getVirtualInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualInstance.
type LookupVirtualInstanceArgs struct {
	// Number of seconds without queries after which the Virtual Instance is suspended.
	AutoSuspendSeconds *int `pulumi:"autoSuspendSeconds"`
	// Virtual Instance current size.
	CurrentSize *string `pulumi:"currentSize"`
	// Virtual Instance name.
	Default *bool `pulumi:"default"`
	// Virtual Instance description.
	Description *string `pulumi:"description"`
	// Virtual Instance desired size.
	DesiredSize *string `pulumi:"desiredSize"`
	// When a Virtual Instance is resumed, it will remount all collections that were mounted when the Virtual Instance was suspended.
	EnableRemountOnResume *bool `pulumi:"enableRemountOnResume"`
	// Virtual Instance id.
	Id string `pulumi:"id"`
	// Number of seconds between data refreshes for mounts on this Virtual Instance.
	MountRefreshIntervalSeconds *int `pulumi:"mountRefreshIntervalSeconds"`
	// Virtual Instance name.
	Name *string `pulumi:"name"`
	// Virtual Instance state.
	State *string `pulumi:"state"`
}

// A collection of values returned by getVirtualInstance.
type LookupVirtualInstanceResult struct {
	// Number of seconds without queries after which the Virtual Instance is suspended.
	AutoSuspendSeconds *int `pulumi:"autoSuspendSeconds"`
	// Virtual Instance current size.
	CurrentSize *string `pulumi:"currentSize"`
	// Virtual Instance name.
	Default *bool `pulumi:"default"`
	// Virtual Instance description.
	Description *string `pulumi:"description"`
	// Virtual Instance desired size.
	DesiredSize *string `pulumi:"desiredSize"`
	// When a Virtual Instance is resumed, it will remount all collections that were mounted when the Virtual Instance was suspended.
	EnableRemountOnResume *bool `pulumi:"enableRemountOnResume"`
	// Virtual Instance id.
	Id string `pulumi:"id"`
	// Number of seconds between data refreshes for mounts on this Virtual Instance.
	MountRefreshIntervalSeconds *int `pulumi:"mountRefreshIntervalSeconds"`
	// Virtual Instance name.
	Name *string `pulumi:"name"`
	// Virtual Instance state.
	State *string `pulumi:"state"`
}

func LookupVirtualInstanceOutput(ctx *pulumi.Context, args LookupVirtualInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualInstanceResult, error) {
			args := v.(LookupVirtualInstanceArgs)
			r, err := LookupVirtualInstance(ctx, &args, opts...)
			var s LookupVirtualInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualInstanceResultOutput)
}

// A collection of arguments for invoking getVirtualInstance.
type LookupVirtualInstanceOutputArgs struct {
	// Number of seconds without queries after which the Virtual Instance is suspended.
	AutoSuspendSeconds pulumi.IntPtrInput `pulumi:"autoSuspendSeconds"`
	// Virtual Instance current size.
	CurrentSize pulumi.StringPtrInput `pulumi:"currentSize"`
	// Virtual Instance name.
	Default pulumi.BoolPtrInput `pulumi:"default"`
	// Virtual Instance description.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Virtual Instance desired size.
	DesiredSize pulumi.StringPtrInput `pulumi:"desiredSize"`
	// When a Virtual Instance is resumed, it will remount all collections that were mounted when the Virtual Instance was suspended.
	EnableRemountOnResume pulumi.BoolPtrInput `pulumi:"enableRemountOnResume"`
	// Virtual Instance id.
	Id pulumi.StringInput `pulumi:"id"`
	// Number of seconds between data refreshes for mounts on this Virtual Instance.
	MountRefreshIntervalSeconds pulumi.IntPtrInput `pulumi:"mountRefreshIntervalSeconds"`
	// Virtual Instance name.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Virtual Instance state.
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (LookupVirtualInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualInstance.
type LookupVirtualInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualInstanceResult)(nil)).Elem()
}

func (o LookupVirtualInstanceResultOutput) ToLookupVirtualInstanceResultOutput() LookupVirtualInstanceResultOutput {
	return o
}

func (o LookupVirtualInstanceResultOutput) ToLookupVirtualInstanceResultOutputWithContext(ctx context.Context) LookupVirtualInstanceResultOutput {
	return o
}

// Number of seconds without queries after which the Virtual Instance is suspended.
func (o LookupVirtualInstanceResultOutput) AutoSuspendSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualInstanceResult) *int { return v.AutoSuspendSeconds }).(pulumi.IntPtrOutput)
}

// Virtual Instance current size.
func (o LookupVirtualInstanceResultOutput) CurrentSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualInstanceResult) *string { return v.CurrentSize }).(pulumi.StringPtrOutput)
}

// Virtual Instance name.
func (o LookupVirtualInstanceResultOutput) Default() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualInstanceResult) *bool { return v.Default }).(pulumi.BoolPtrOutput)
}

// Virtual Instance description.
func (o LookupVirtualInstanceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualInstanceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Virtual Instance desired size.
func (o LookupVirtualInstanceResultOutput) DesiredSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualInstanceResult) *string { return v.DesiredSize }).(pulumi.StringPtrOutput)
}

// When a Virtual Instance is resumed, it will remount all collections that were mounted when the Virtual Instance was suspended.
func (o LookupVirtualInstanceResultOutput) EnableRemountOnResume() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualInstanceResult) *bool { return v.EnableRemountOnResume }).(pulumi.BoolPtrOutput)
}

// Virtual Instance id.
func (o LookupVirtualInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Number of seconds between data refreshes for mounts on this Virtual Instance.
func (o LookupVirtualInstanceResultOutput) MountRefreshIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualInstanceResult) *int { return v.MountRefreshIntervalSeconds }).(pulumi.IntPtrOutput)
}

// Virtual Instance name.
func (o LookupVirtualInstanceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualInstanceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Virtual Instance state.
func (o LookupVirtualInstanceResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualInstanceResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualInstanceResultOutput{})
}
