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

// Manages a collection with an GCS source attached.
type GcsCollection struct {
	pulumi.CustomResourceState

	// Text describing the collection.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Ingest transformation SQL query. Turns the collection into insertOnly mode.
	IngestTransformation pulumi.StringPtrOutput `pulumi:"ingestTransformation"`
	// Unique identifier for the collection. Can contain alphanumeric or dash characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of seconds after which data is purged. Based on event time.
	RetentionSecs pulumi.IntPtrOutput `pulumi:"retentionSecs"`
	// Defines a source for this collection.
	Sources GcsCollectionSourceArrayOutput `pulumi:"sources"`
	// Wait until the collection is ready.
	WaitForCollection pulumi.BoolPtrOutput `pulumi:"waitForCollection"`
	// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
	WaitForDocuments pulumi.IntPtrOutput `pulumi:"waitForDocuments"`
	// The name of the workspace.
	Workspace pulumi.StringOutput `pulumi:"workspace"`
}

// NewGcsCollection registers a new resource with the given unique name, arguments, and options.
func NewGcsCollection(ctx *pulumi.Context,
	name string, args *GcsCollectionArgs, opts ...pulumi.ResourceOption) (*GcsCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Workspace == nil {
		return nil, errors.New("invalid value for required argument 'Workspace'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GcsCollection
	err := ctx.RegisterResource("rockset:index/gcsCollection:GcsCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGcsCollection gets an existing GcsCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGcsCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GcsCollectionState, opts ...pulumi.ResourceOption) (*GcsCollection, error) {
	var resource GcsCollection
	err := ctx.ReadResource("rockset:index/gcsCollection:GcsCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GcsCollection resources.
type gcsCollectionState struct {
	// Text describing the collection.
	Description *string `pulumi:"description"`
	// Ingest transformation SQL query. Turns the collection into insertOnly mode.
	IngestTransformation *string `pulumi:"ingestTransformation"`
	// Unique identifier for the collection. Can contain alphanumeric or dash characters.
	Name *string `pulumi:"name"`
	// Number of seconds after which data is purged. Based on event time.
	RetentionSecs *int `pulumi:"retentionSecs"`
	// Defines a source for this collection.
	Sources []GcsCollectionSource `pulumi:"sources"`
	// Wait until the collection is ready.
	WaitForCollection *bool `pulumi:"waitForCollection"`
	// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
	WaitForDocuments *int `pulumi:"waitForDocuments"`
	// The name of the workspace.
	Workspace *string `pulumi:"workspace"`
}

type GcsCollectionState struct {
	// Text describing the collection.
	Description pulumi.StringPtrInput
	// Ingest transformation SQL query. Turns the collection into insertOnly mode.
	IngestTransformation pulumi.StringPtrInput
	// Unique identifier for the collection. Can contain alphanumeric or dash characters.
	Name pulumi.StringPtrInput
	// Number of seconds after which data is purged. Based on event time.
	RetentionSecs pulumi.IntPtrInput
	// Defines a source for this collection.
	Sources GcsCollectionSourceArrayInput
	// Wait until the collection is ready.
	WaitForCollection pulumi.BoolPtrInput
	// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
	WaitForDocuments pulumi.IntPtrInput
	// The name of the workspace.
	Workspace pulumi.StringPtrInput
}

func (GcsCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*gcsCollectionState)(nil)).Elem()
}

type gcsCollectionArgs struct {
	// Text describing the collection.
	Description *string `pulumi:"description"`
	// Ingest transformation SQL query. Turns the collection into insertOnly mode.
	IngestTransformation *string `pulumi:"ingestTransformation"`
	// Unique identifier for the collection. Can contain alphanumeric or dash characters.
	Name *string `pulumi:"name"`
	// Number of seconds after which data is purged. Based on event time.
	RetentionSecs *int `pulumi:"retentionSecs"`
	// Defines a source for this collection.
	Sources []GcsCollectionSource `pulumi:"sources"`
	// Wait until the collection is ready.
	WaitForCollection *bool `pulumi:"waitForCollection"`
	// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
	WaitForDocuments *int `pulumi:"waitForDocuments"`
	// The name of the workspace.
	Workspace string `pulumi:"workspace"`
}

// The set of arguments for constructing a GcsCollection resource.
type GcsCollectionArgs struct {
	// Text describing the collection.
	Description pulumi.StringPtrInput
	// Ingest transformation SQL query. Turns the collection into insertOnly mode.
	IngestTransformation pulumi.StringPtrInput
	// Unique identifier for the collection. Can contain alphanumeric or dash characters.
	Name pulumi.StringPtrInput
	// Number of seconds after which data is purged. Based on event time.
	RetentionSecs pulumi.IntPtrInput
	// Defines a source for this collection.
	Sources GcsCollectionSourceArrayInput
	// Wait until the collection is ready.
	WaitForCollection pulumi.BoolPtrInput
	// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
	WaitForDocuments pulumi.IntPtrInput
	// The name of the workspace.
	Workspace pulumi.StringInput
}

func (GcsCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gcsCollectionArgs)(nil)).Elem()
}

type GcsCollectionInput interface {
	pulumi.Input

	ToGcsCollectionOutput() GcsCollectionOutput
	ToGcsCollectionOutputWithContext(ctx context.Context) GcsCollectionOutput
}

func (*GcsCollection) ElementType() reflect.Type {
	return reflect.TypeOf((**GcsCollection)(nil)).Elem()
}

func (i *GcsCollection) ToGcsCollectionOutput() GcsCollectionOutput {
	return i.ToGcsCollectionOutputWithContext(context.Background())
}

func (i *GcsCollection) ToGcsCollectionOutputWithContext(ctx context.Context) GcsCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcsCollectionOutput)
}

// GcsCollectionArrayInput is an input type that accepts GcsCollectionArray and GcsCollectionArrayOutput values.
// You can construct a concrete instance of `GcsCollectionArrayInput` via:
//
//	GcsCollectionArray{ GcsCollectionArgs{...} }
type GcsCollectionArrayInput interface {
	pulumi.Input

	ToGcsCollectionArrayOutput() GcsCollectionArrayOutput
	ToGcsCollectionArrayOutputWithContext(context.Context) GcsCollectionArrayOutput
}

type GcsCollectionArray []GcsCollectionInput

func (GcsCollectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GcsCollection)(nil)).Elem()
}

func (i GcsCollectionArray) ToGcsCollectionArrayOutput() GcsCollectionArrayOutput {
	return i.ToGcsCollectionArrayOutputWithContext(context.Background())
}

func (i GcsCollectionArray) ToGcsCollectionArrayOutputWithContext(ctx context.Context) GcsCollectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcsCollectionArrayOutput)
}

// GcsCollectionMapInput is an input type that accepts GcsCollectionMap and GcsCollectionMapOutput values.
// You can construct a concrete instance of `GcsCollectionMapInput` via:
//
//	GcsCollectionMap{ "key": GcsCollectionArgs{...} }
type GcsCollectionMapInput interface {
	pulumi.Input

	ToGcsCollectionMapOutput() GcsCollectionMapOutput
	ToGcsCollectionMapOutputWithContext(context.Context) GcsCollectionMapOutput
}

type GcsCollectionMap map[string]GcsCollectionInput

func (GcsCollectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GcsCollection)(nil)).Elem()
}

func (i GcsCollectionMap) ToGcsCollectionMapOutput() GcsCollectionMapOutput {
	return i.ToGcsCollectionMapOutputWithContext(context.Background())
}

func (i GcsCollectionMap) ToGcsCollectionMapOutputWithContext(ctx context.Context) GcsCollectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcsCollectionMapOutput)
}

type GcsCollectionOutput struct{ *pulumi.OutputState }

func (GcsCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GcsCollection)(nil)).Elem()
}

func (o GcsCollectionOutput) ToGcsCollectionOutput() GcsCollectionOutput {
	return o
}

func (o GcsCollectionOutput) ToGcsCollectionOutputWithContext(ctx context.Context) GcsCollectionOutput {
	return o
}

// Text describing the collection.
func (o GcsCollectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GcsCollection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Ingest transformation SQL query. Turns the collection into insertOnly mode.
func (o GcsCollectionOutput) IngestTransformation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GcsCollection) pulumi.StringPtrOutput { return v.IngestTransformation }).(pulumi.StringPtrOutput)
}

// Unique identifier for the collection. Can contain alphanumeric or dash characters.
func (o GcsCollectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GcsCollection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of seconds after which data is purged. Based on event time.
func (o GcsCollectionOutput) RetentionSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GcsCollection) pulumi.IntPtrOutput { return v.RetentionSecs }).(pulumi.IntPtrOutput)
}

// Defines a source for this collection.
func (o GcsCollectionOutput) Sources() GcsCollectionSourceArrayOutput {
	return o.ApplyT(func(v *GcsCollection) GcsCollectionSourceArrayOutput { return v.Sources }).(GcsCollectionSourceArrayOutput)
}

// Wait until the collection is ready.
func (o GcsCollectionOutput) WaitForCollection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GcsCollection) pulumi.BoolPtrOutput { return v.WaitForCollection }).(pulumi.BoolPtrOutput)
}

// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
func (o GcsCollectionOutput) WaitForDocuments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GcsCollection) pulumi.IntPtrOutput { return v.WaitForDocuments }).(pulumi.IntPtrOutput)
}

// The name of the workspace.
func (o GcsCollectionOutput) Workspace() pulumi.StringOutput {
	return o.ApplyT(func(v *GcsCollection) pulumi.StringOutput { return v.Workspace }).(pulumi.StringOutput)
}

type GcsCollectionArrayOutput struct{ *pulumi.OutputState }

func (GcsCollectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GcsCollection)(nil)).Elem()
}

func (o GcsCollectionArrayOutput) ToGcsCollectionArrayOutput() GcsCollectionArrayOutput {
	return o
}

func (o GcsCollectionArrayOutput) ToGcsCollectionArrayOutputWithContext(ctx context.Context) GcsCollectionArrayOutput {
	return o
}

func (o GcsCollectionArrayOutput) Index(i pulumi.IntInput) GcsCollectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GcsCollection {
		return vs[0].([]*GcsCollection)[vs[1].(int)]
	}).(GcsCollectionOutput)
}

type GcsCollectionMapOutput struct{ *pulumi.OutputState }

func (GcsCollectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GcsCollection)(nil)).Elem()
}

func (o GcsCollectionMapOutput) ToGcsCollectionMapOutput() GcsCollectionMapOutput {
	return o
}

func (o GcsCollectionMapOutput) ToGcsCollectionMapOutputWithContext(ctx context.Context) GcsCollectionMapOutput {
	return o
}

func (o GcsCollectionMapOutput) MapIndex(k pulumi.StringInput) GcsCollectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GcsCollection {
		return vs[0].(map[string]*GcsCollection)[vs[1].(string)]
	}).(GcsCollectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GcsCollectionInput)(nil)).Elem(), &GcsCollection{})
	pulumi.RegisterInputType(reflect.TypeOf((*GcsCollectionArrayInput)(nil)).Elem(), GcsCollectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GcsCollectionMapInput)(nil)).Elem(), GcsCollectionMap{})
	pulumi.RegisterOutputType(GcsCollectionOutput{})
	pulumi.RegisterOutputType(GcsCollectionArrayOutput{})
	pulumi.RegisterOutputType(GcsCollectionMapOutput{})
}
