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

// Manages a collection with an Kinesis source attached.
type KinesisCollection struct {
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
	Sources KinesisCollectionSourceArrayOutput `pulumi:"sources"`
	// Wait until the collection is ready.
	WaitForCollection pulumi.BoolPtrOutput `pulumi:"waitForCollection"`
	// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
	WaitForDocuments pulumi.IntPtrOutput `pulumi:"waitForDocuments"`
	// The name of the workspace.
	Workspace pulumi.StringOutput `pulumi:"workspace"`
}

// NewKinesisCollection registers a new resource with the given unique name, arguments, and options.
func NewKinesisCollection(ctx *pulumi.Context,
	name string, args *KinesisCollectionArgs, opts ...pulumi.ResourceOption) (*KinesisCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Workspace == nil {
		return nil, errors.New("invalid value for required argument 'Workspace'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KinesisCollection
	err := ctx.RegisterResource("rockset:index/kinesisCollection:KinesisCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKinesisCollection gets an existing KinesisCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKinesisCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KinesisCollectionState, opts ...pulumi.ResourceOption) (*KinesisCollection, error) {
	var resource KinesisCollection
	err := ctx.ReadResource("rockset:index/kinesisCollection:KinesisCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KinesisCollection resources.
type kinesisCollectionState struct {
	// Text describing the collection.
	Description *string `pulumi:"description"`
	// Ingest transformation SQL query. Turns the collection into insertOnly mode.
	IngestTransformation *string `pulumi:"ingestTransformation"`
	// Unique identifier for the collection. Can contain alphanumeric or dash characters.
	Name *string `pulumi:"name"`
	// Number of seconds after which data is purged. Based on event time.
	RetentionSecs *int `pulumi:"retentionSecs"`
	// Defines a source for this collection.
	Sources []KinesisCollectionSource `pulumi:"sources"`
	// Wait until the collection is ready.
	WaitForCollection *bool `pulumi:"waitForCollection"`
	// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
	WaitForDocuments *int `pulumi:"waitForDocuments"`
	// The name of the workspace.
	Workspace *string `pulumi:"workspace"`
}

type KinesisCollectionState struct {
	// Text describing the collection.
	Description pulumi.StringPtrInput
	// Ingest transformation SQL query. Turns the collection into insertOnly mode.
	IngestTransformation pulumi.StringPtrInput
	// Unique identifier for the collection. Can contain alphanumeric or dash characters.
	Name pulumi.StringPtrInput
	// Number of seconds after which data is purged. Based on event time.
	RetentionSecs pulumi.IntPtrInput
	// Defines a source for this collection.
	Sources KinesisCollectionSourceArrayInput
	// Wait until the collection is ready.
	WaitForCollection pulumi.BoolPtrInput
	// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
	WaitForDocuments pulumi.IntPtrInput
	// The name of the workspace.
	Workspace pulumi.StringPtrInput
}

func (KinesisCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*kinesisCollectionState)(nil)).Elem()
}

type kinesisCollectionArgs struct {
	// Text describing the collection.
	Description *string `pulumi:"description"`
	// Ingest transformation SQL query. Turns the collection into insertOnly mode.
	IngestTransformation *string `pulumi:"ingestTransformation"`
	// Unique identifier for the collection. Can contain alphanumeric or dash characters.
	Name *string `pulumi:"name"`
	// Number of seconds after which data is purged. Based on event time.
	RetentionSecs *int `pulumi:"retentionSecs"`
	// Defines a source for this collection.
	Sources []KinesisCollectionSource `pulumi:"sources"`
	// Wait until the collection is ready.
	WaitForCollection *bool `pulumi:"waitForCollection"`
	// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
	WaitForDocuments *int `pulumi:"waitForDocuments"`
	// The name of the workspace.
	Workspace string `pulumi:"workspace"`
}

// The set of arguments for constructing a KinesisCollection resource.
type KinesisCollectionArgs struct {
	// Text describing the collection.
	Description pulumi.StringPtrInput
	// Ingest transformation SQL query. Turns the collection into insertOnly mode.
	IngestTransformation pulumi.StringPtrInput
	// Unique identifier for the collection. Can contain alphanumeric or dash characters.
	Name pulumi.StringPtrInput
	// Number of seconds after which data is purged. Based on event time.
	RetentionSecs pulumi.IntPtrInput
	// Defines a source for this collection.
	Sources KinesisCollectionSourceArrayInput
	// Wait until the collection is ready.
	WaitForCollection pulumi.BoolPtrInput
	// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
	WaitForDocuments pulumi.IntPtrInput
	// The name of the workspace.
	Workspace pulumi.StringInput
}

func (KinesisCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kinesisCollectionArgs)(nil)).Elem()
}

type KinesisCollectionInput interface {
	pulumi.Input

	ToKinesisCollectionOutput() KinesisCollectionOutput
	ToKinesisCollectionOutputWithContext(ctx context.Context) KinesisCollectionOutput
}

func (*KinesisCollection) ElementType() reflect.Type {
	return reflect.TypeOf((**KinesisCollection)(nil)).Elem()
}

func (i *KinesisCollection) ToKinesisCollectionOutput() KinesisCollectionOutput {
	return i.ToKinesisCollectionOutputWithContext(context.Background())
}

func (i *KinesisCollection) ToKinesisCollectionOutputWithContext(ctx context.Context) KinesisCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KinesisCollectionOutput)
}

// KinesisCollectionArrayInput is an input type that accepts KinesisCollectionArray and KinesisCollectionArrayOutput values.
// You can construct a concrete instance of `KinesisCollectionArrayInput` via:
//
//	KinesisCollectionArray{ KinesisCollectionArgs{...} }
type KinesisCollectionArrayInput interface {
	pulumi.Input

	ToKinesisCollectionArrayOutput() KinesisCollectionArrayOutput
	ToKinesisCollectionArrayOutputWithContext(context.Context) KinesisCollectionArrayOutput
}

type KinesisCollectionArray []KinesisCollectionInput

func (KinesisCollectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KinesisCollection)(nil)).Elem()
}

func (i KinesisCollectionArray) ToKinesisCollectionArrayOutput() KinesisCollectionArrayOutput {
	return i.ToKinesisCollectionArrayOutputWithContext(context.Background())
}

func (i KinesisCollectionArray) ToKinesisCollectionArrayOutputWithContext(ctx context.Context) KinesisCollectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KinesisCollectionArrayOutput)
}

// KinesisCollectionMapInput is an input type that accepts KinesisCollectionMap and KinesisCollectionMapOutput values.
// You can construct a concrete instance of `KinesisCollectionMapInput` via:
//
//	KinesisCollectionMap{ "key": KinesisCollectionArgs{...} }
type KinesisCollectionMapInput interface {
	pulumi.Input

	ToKinesisCollectionMapOutput() KinesisCollectionMapOutput
	ToKinesisCollectionMapOutputWithContext(context.Context) KinesisCollectionMapOutput
}

type KinesisCollectionMap map[string]KinesisCollectionInput

func (KinesisCollectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KinesisCollection)(nil)).Elem()
}

func (i KinesisCollectionMap) ToKinesisCollectionMapOutput() KinesisCollectionMapOutput {
	return i.ToKinesisCollectionMapOutputWithContext(context.Background())
}

func (i KinesisCollectionMap) ToKinesisCollectionMapOutputWithContext(ctx context.Context) KinesisCollectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KinesisCollectionMapOutput)
}

type KinesisCollectionOutput struct{ *pulumi.OutputState }

func (KinesisCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KinesisCollection)(nil)).Elem()
}

func (o KinesisCollectionOutput) ToKinesisCollectionOutput() KinesisCollectionOutput {
	return o
}

func (o KinesisCollectionOutput) ToKinesisCollectionOutputWithContext(ctx context.Context) KinesisCollectionOutput {
	return o
}

// Text describing the collection.
func (o KinesisCollectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KinesisCollection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Ingest transformation SQL query. Turns the collection into insertOnly mode.
func (o KinesisCollectionOutput) IngestTransformation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KinesisCollection) pulumi.StringPtrOutput { return v.IngestTransformation }).(pulumi.StringPtrOutput)
}

// Unique identifier for the collection. Can contain alphanumeric or dash characters.
func (o KinesisCollectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KinesisCollection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of seconds after which data is purged. Based on event time.
func (o KinesisCollectionOutput) RetentionSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KinesisCollection) pulumi.IntPtrOutput { return v.RetentionSecs }).(pulumi.IntPtrOutput)
}

// Defines a source for this collection.
func (o KinesisCollectionOutput) Sources() KinesisCollectionSourceArrayOutput {
	return o.ApplyT(func(v *KinesisCollection) KinesisCollectionSourceArrayOutput { return v.Sources }).(KinesisCollectionSourceArrayOutput)
}

// Wait until the collection is ready.
func (o KinesisCollectionOutput) WaitForCollection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KinesisCollection) pulumi.BoolPtrOutput { return v.WaitForCollection }).(pulumi.BoolPtrOutput)
}

// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
func (o KinesisCollectionOutput) WaitForDocuments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KinesisCollection) pulumi.IntPtrOutput { return v.WaitForDocuments }).(pulumi.IntPtrOutput)
}

// The name of the workspace.
func (o KinesisCollectionOutput) Workspace() pulumi.StringOutput {
	return o.ApplyT(func(v *KinesisCollection) pulumi.StringOutput { return v.Workspace }).(pulumi.StringOutput)
}

type KinesisCollectionArrayOutput struct{ *pulumi.OutputState }

func (KinesisCollectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KinesisCollection)(nil)).Elem()
}

func (o KinesisCollectionArrayOutput) ToKinesisCollectionArrayOutput() KinesisCollectionArrayOutput {
	return o
}

func (o KinesisCollectionArrayOutput) ToKinesisCollectionArrayOutputWithContext(ctx context.Context) KinesisCollectionArrayOutput {
	return o
}

func (o KinesisCollectionArrayOutput) Index(i pulumi.IntInput) KinesisCollectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KinesisCollection {
		return vs[0].([]*KinesisCollection)[vs[1].(int)]
	}).(KinesisCollectionOutput)
}

type KinesisCollectionMapOutput struct{ *pulumi.OutputState }

func (KinesisCollectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KinesisCollection)(nil)).Elem()
}

func (o KinesisCollectionMapOutput) ToKinesisCollectionMapOutput() KinesisCollectionMapOutput {
	return o
}

func (o KinesisCollectionMapOutput) ToKinesisCollectionMapOutputWithContext(ctx context.Context) KinesisCollectionMapOutput {
	return o
}

func (o KinesisCollectionMapOutput) MapIndex(k pulumi.StringInput) KinesisCollectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KinesisCollection {
		return vs[0].(map[string]*KinesisCollection)[vs[1].(string)]
	}).(KinesisCollectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KinesisCollectionInput)(nil)).Elem(), &KinesisCollection{})
	pulumi.RegisterInputType(reflect.TypeOf((*KinesisCollectionArrayInput)(nil)).Elem(), KinesisCollectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KinesisCollectionMapInput)(nil)).Elem(), KinesisCollectionMap{})
	pulumi.RegisterOutputType(KinesisCollectionOutput{})
	pulumi.RegisterOutputType(KinesisCollectionArrayOutput{})
	pulumi.RegisterOutputType(KinesisCollectionMapOutput{})
}
