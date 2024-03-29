// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rockset

import (
	"context"
	"reflect"

	"errors"
	"github.com/ddn0/pulumi-rockset/sdk/go/rockset/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a collection mount.
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
//			_, err := rockset.NewCollectionMount(ctx, "data", &rockset.CollectionMountArgs{
//				VirtualInstanceId: pulumi.Any(rockset_virtual_instance.Secondary.Id),
//				Path:              pulumi.String("commons.data"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
//
//	$ pulumi import rockset:index/collectionMount:CollectionMount demo 29e4a43c-fff4-4fe6-80e3-1ee57bc22e82
//
// ```
type CollectionMount struct {
	pulumi.CustomResourceState

	// ISO 8601 date when the mount was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// UNIX timestamp in milliseconds for most recent refresh. Not applicable for live mounts.
	LastRefreshTime pulumi.IntOutput `pulumi:"lastRefreshTime"`
	// Collection path to be mounted, in the form workspace.collection
	Path pulumi.StringOutput `pulumi:"path"`
	// RRN of this mount.
	Rrn pulumi.StringOutput `pulumi:"rrn"`
	// Mount state.
	State pulumi.StringOutput `pulumi:"state"`
	// Virtual Instance id
	VirtualInstanceId pulumi.StringOutput `pulumi:"virtualInstanceId"`
	// Virtual Instance RRN
	VirtualInstanceRrn pulumi.StringOutput `pulumi:"virtualInstanceRrn"`
}

// NewCollectionMount registers a new resource with the given unique name, arguments, and options.
func NewCollectionMount(ctx *pulumi.Context,
	name string, args *CollectionMountArgs, opts ...pulumi.ResourceOption) (*CollectionMount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.VirtualInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualInstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CollectionMount
	err := ctx.RegisterResource("rockset:index/collectionMount:CollectionMount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCollectionMount gets an existing CollectionMount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCollectionMount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CollectionMountState, opts ...pulumi.ResourceOption) (*CollectionMount, error) {
	var resource CollectionMount
	err := ctx.ReadResource("rockset:index/collectionMount:CollectionMount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CollectionMount resources.
type collectionMountState struct {
	// ISO 8601 date when the mount was created.
	CreatedAt *string `pulumi:"createdAt"`
	// UNIX timestamp in milliseconds for most recent refresh. Not applicable for live mounts.
	LastRefreshTime *int `pulumi:"lastRefreshTime"`
	// Collection path to be mounted, in the form workspace.collection
	Path *string `pulumi:"path"`
	// RRN of this mount.
	Rrn *string `pulumi:"rrn"`
	// Mount state.
	State *string `pulumi:"state"`
	// Virtual Instance id
	VirtualInstanceId *string `pulumi:"virtualInstanceId"`
	// Virtual Instance RRN
	VirtualInstanceRrn *string `pulumi:"virtualInstanceRrn"`
}

type CollectionMountState struct {
	// ISO 8601 date when the mount was created.
	CreatedAt pulumi.StringPtrInput
	// UNIX timestamp in milliseconds for most recent refresh. Not applicable for live mounts.
	LastRefreshTime pulumi.IntPtrInput
	// Collection path to be mounted, in the form workspace.collection
	Path pulumi.StringPtrInput
	// RRN of this mount.
	Rrn pulumi.StringPtrInput
	// Mount state.
	State pulumi.StringPtrInput
	// Virtual Instance id
	VirtualInstanceId pulumi.StringPtrInput
	// Virtual Instance RRN
	VirtualInstanceRrn pulumi.StringPtrInput
}

func (CollectionMountState) ElementType() reflect.Type {
	return reflect.TypeOf((*collectionMountState)(nil)).Elem()
}

type collectionMountArgs struct {
	// Collection path to be mounted, in the form workspace.collection
	Path string `pulumi:"path"`
	// Virtual Instance id
	VirtualInstanceId string `pulumi:"virtualInstanceId"`
}

// The set of arguments for constructing a CollectionMount resource.
type CollectionMountArgs struct {
	// Collection path to be mounted, in the form workspace.collection
	Path pulumi.StringInput
	// Virtual Instance id
	VirtualInstanceId pulumi.StringInput
}

func (CollectionMountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*collectionMountArgs)(nil)).Elem()
}

type CollectionMountInput interface {
	pulumi.Input

	ToCollectionMountOutput() CollectionMountOutput
	ToCollectionMountOutputWithContext(ctx context.Context) CollectionMountOutput
}

func (*CollectionMount) ElementType() reflect.Type {
	return reflect.TypeOf((**CollectionMount)(nil)).Elem()
}

func (i *CollectionMount) ToCollectionMountOutput() CollectionMountOutput {
	return i.ToCollectionMountOutputWithContext(context.Background())
}

func (i *CollectionMount) ToCollectionMountOutputWithContext(ctx context.Context) CollectionMountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectionMountOutput)
}

// CollectionMountArrayInput is an input type that accepts CollectionMountArray and CollectionMountArrayOutput values.
// You can construct a concrete instance of `CollectionMountArrayInput` via:
//
//	CollectionMountArray{ CollectionMountArgs{...} }
type CollectionMountArrayInput interface {
	pulumi.Input

	ToCollectionMountArrayOutput() CollectionMountArrayOutput
	ToCollectionMountArrayOutputWithContext(context.Context) CollectionMountArrayOutput
}

type CollectionMountArray []CollectionMountInput

func (CollectionMountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CollectionMount)(nil)).Elem()
}

func (i CollectionMountArray) ToCollectionMountArrayOutput() CollectionMountArrayOutput {
	return i.ToCollectionMountArrayOutputWithContext(context.Background())
}

func (i CollectionMountArray) ToCollectionMountArrayOutputWithContext(ctx context.Context) CollectionMountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectionMountArrayOutput)
}

// CollectionMountMapInput is an input type that accepts CollectionMountMap and CollectionMountMapOutput values.
// You can construct a concrete instance of `CollectionMountMapInput` via:
//
//	CollectionMountMap{ "key": CollectionMountArgs{...} }
type CollectionMountMapInput interface {
	pulumi.Input

	ToCollectionMountMapOutput() CollectionMountMapOutput
	ToCollectionMountMapOutputWithContext(context.Context) CollectionMountMapOutput
}

type CollectionMountMap map[string]CollectionMountInput

func (CollectionMountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CollectionMount)(nil)).Elem()
}

func (i CollectionMountMap) ToCollectionMountMapOutput() CollectionMountMapOutput {
	return i.ToCollectionMountMapOutputWithContext(context.Background())
}

func (i CollectionMountMap) ToCollectionMountMapOutputWithContext(ctx context.Context) CollectionMountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectionMountMapOutput)
}

type CollectionMountOutput struct{ *pulumi.OutputState }

func (CollectionMountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CollectionMount)(nil)).Elem()
}

func (o CollectionMountOutput) ToCollectionMountOutput() CollectionMountOutput {
	return o
}

func (o CollectionMountOutput) ToCollectionMountOutputWithContext(ctx context.Context) CollectionMountOutput {
	return o
}

// ISO 8601 date when the mount was created.
func (o CollectionMountOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CollectionMount) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// UNIX timestamp in milliseconds for most recent refresh. Not applicable for live mounts.
func (o CollectionMountOutput) LastRefreshTime() pulumi.IntOutput {
	return o.ApplyT(func(v *CollectionMount) pulumi.IntOutput { return v.LastRefreshTime }).(pulumi.IntOutput)
}

// Collection path to be mounted, in the form workspace.collection
func (o CollectionMountOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *CollectionMount) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// RRN of this mount.
func (o CollectionMountOutput) Rrn() pulumi.StringOutput {
	return o.ApplyT(func(v *CollectionMount) pulumi.StringOutput { return v.Rrn }).(pulumi.StringOutput)
}

// Mount state.
func (o CollectionMountOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *CollectionMount) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Virtual Instance id
func (o CollectionMountOutput) VirtualInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CollectionMount) pulumi.StringOutput { return v.VirtualInstanceId }).(pulumi.StringOutput)
}

// Virtual Instance RRN
func (o CollectionMountOutput) VirtualInstanceRrn() pulumi.StringOutput {
	return o.ApplyT(func(v *CollectionMount) pulumi.StringOutput { return v.VirtualInstanceRrn }).(pulumi.StringOutput)
}

type CollectionMountArrayOutput struct{ *pulumi.OutputState }

func (CollectionMountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CollectionMount)(nil)).Elem()
}

func (o CollectionMountArrayOutput) ToCollectionMountArrayOutput() CollectionMountArrayOutput {
	return o
}

func (o CollectionMountArrayOutput) ToCollectionMountArrayOutputWithContext(ctx context.Context) CollectionMountArrayOutput {
	return o
}

func (o CollectionMountArrayOutput) Index(i pulumi.IntInput) CollectionMountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CollectionMount {
		return vs[0].([]*CollectionMount)[vs[1].(int)]
	}).(CollectionMountOutput)
}

type CollectionMountMapOutput struct{ *pulumi.OutputState }

func (CollectionMountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CollectionMount)(nil)).Elem()
}

func (o CollectionMountMapOutput) ToCollectionMountMapOutput() CollectionMountMapOutput {
	return o
}

func (o CollectionMountMapOutput) ToCollectionMountMapOutputWithContext(ctx context.Context) CollectionMountMapOutput {
	return o
}

func (o CollectionMountMapOutput) MapIndex(k pulumi.StringInput) CollectionMountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CollectionMount {
		return vs[0].(map[string]*CollectionMount)[vs[1].(string)]
	}).(CollectionMountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CollectionMountInput)(nil)).Elem(), &CollectionMount{})
	pulumi.RegisterInputType(reflect.TypeOf((*CollectionMountArrayInput)(nil)).Elem(), CollectionMountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CollectionMountMapInput)(nil)).Elem(), CollectionMountMap{})
	pulumi.RegisterOutputType(CollectionMountOutput{})
	pulumi.RegisterOutputType(CollectionMountArrayOutput{})
	pulumi.RegisterOutputType(CollectionMountMapOutput{})
}
