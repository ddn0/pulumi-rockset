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

// Manages a collection with on or more S3 sources attached. Uses an S3 integration to access the S3 bucket. If no integration is provided, only data in public buckets are accessible.
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
//			sample, err := rockset.NewWorkspace(ctx, "sample", &rockset.WorkspaceArgs{
//				Description: pulumi.String("sample datasets"),
//			})
//			if err != nil {
//				return err
//			}
//			public, err := rockset.NewS3Integration(ctx, "public", &rockset.S3IntegrationArgs{
//				Description: pulumi.String("Integration to access Rockset's public datasets"),
//				AwsRoleArn:  pulumi.String("arn:aws:iam::469279130686:role/rockset-public-datasets"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rockset.NewS3Collection(ctx, "cities", &rockset.S3CollectionArgs{
//				Workspace: sample.Name,
//				Sources: rockset.S3CollectionSourceArray{
//					&rockset.S3CollectionSourceArgs{
//						Bucket:          pulumi.String("rockset-public-datasets"),
//						IntegrationName: public.Name,
//						Pattern:         pulumi.String("cities/*.json"),
//						Format:          pulumi.String("json"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type S3Collection struct {
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
	Sources S3CollectionSourceArrayOutput `pulumi:"sources"`
	// Wait until the collection is ready.
	WaitForCollection pulumi.BoolPtrOutput `pulumi:"waitForCollection"`
	// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
	WaitForDocuments pulumi.IntPtrOutput `pulumi:"waitForDocuments"`
	// The name of the workspace.
	Workspace pulumi.StringOutput `pulumi:"workspace"`
}

// NewS3Collection registers a new resource with the given unique name, arguments, and options.
func NewS3Collection(ctx *pulumi.Context,
	name string, args *S3CollectionArgs, opts ...pulumi.ResourceOption) (*S3Collection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Workspace == nil {
		return nil, errors.New("invalid value for required argument 'Workspace'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource S3Collection
	err := ctx.RegisterResource("rockset:index/s3Collection:S3Collection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetS3Collection gets an existing S3Collection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetS3Collection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *S3CollectionState, opts ...pulumi.ResourceOption) (*S3Collection, error) {
	var resource S3Collection
	err := ctx.ReadResource("rockset:index/s3Collection:S3Collection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering S3Collection resources.
type s3collectionState struct {
	// Text describing the collection.
	Description *string `pulumi:"description"`
	// Ingest transformation SQL query. Turns the collection into insertOnly mode.
	IngestTransformation *string `pulumi:"ingestTransformation"`
	// Unique identifier for the collection. Can contain alphanumeric or dash characters.
	Name *string `pulumi:"name"`
	// Number of seconds after which data is purged. Based on event time.
	RetentionSecs *int `pulumi:"retentionSecs"`
	// Defines a source for this collection.
	Sources []S3CollectionSource `pulumi:"sources"`
	// Wait until the collection is ready.
	WaitForCollection *bool `pulumi:"waitForCollection"`
	// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
	WaitForDocuments *int `pulumi:"waitForDocuments"`
	// The name of the workspace.
	Workspace *string `pulumi:"workspace"`
}

type S3CollectionState struct {
	// Text describing the collection.
	Description pulumi.StringPtrInput
	// Ingest transformation SQL query. Turns the collection into insertOnly mode.
	IngestTransformation pulumi.StringPtrInput
	// Unique identifier for the collection. Can contain alphanumeric or dash characters.
	Name pulumi.StringPtrInput
	// Number of seconds after which data is purged. Based on event time.
	RetentionSecs pulumi.IntPtrInput
	// Defines a source for this collection.
	Sources S3CollectionSourceArrayInput
	// Wait until the collection is ready.
	WaitForCollection pulumi.BoolPtrInput
	// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
	WaitForDocuments pulumi.IntPtrInput
	// The name of the workspace.
	Workspace pulumi.StringPtrInput
}

func (S3CollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*s3collectionState)(nil)).Elem()
}

type s3collectionArgs struct {
	// Text describing the collection.
	Description *string `pulumi:"description"`
	// Ingest transformation SQL query. Turns the collection into insertOnly mode.
	IngestTransformation *string `pulumi:"ingestTransformation"`
	// Unique identifier for the collection. Can contain alphanumeric or dash characters.
	Name *string `pulumi:"name"`
	// Number of seconds after which data is purged. Based on event time.
	RetentionSecs *int `pulumi:"retentionSecs"`
	// Defines a source for this collection.
	Sources []S3CollectionSource `pulumi:"sources"`
	// Wait until the collection is ready.
	WaitForCollection *bool `pulumi:"waitForCollection"`
	// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
	WaitForDocuments *int `pulumi:"waitForDocuments"`
	// The name of the workspace.
	Workspace string `pulumi:"workspace"`
}

// The set of arguments for constructing a S3Collection resource.
type S3CollectionArgs struct {
	// Text describing the collection.
	Description pulumi.StringPtrInput
	// Ingest transformation SQL query. Turns the collection into insertOnly mode.
	IngestTransformation pulumi.StringPtrInput
	// Unique identifier for the collection. Can contain alphanumeric or dash characters.
	Name pulumi.StringPtrInput
	// Number of seconds after which data is purged. Based on event time.
	RetentionSecs pulumi.IntPtrInput
	// Defines a source for this collection.
	Sources S3CollectionSourceArrayInput
	// Wait until the collection is ready.
	WaitForCollection pulumi.BoolPtrInput
	// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
	WaitForDocuments pulumi.IntPtrInput
	// The name of the workspace.
	Workspace pulumi.StringInput
}

func (S3CollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*s3collectionArgs)(nil)).Elem()
}

type S3CollectionInput interface {
	pulumi.Input

	ToS3CollectionOutput() S3CollectionOutput
	ToS3CollectionOutputWithContext(ctx context.Context) S3CollectionOutput
}

func (*S3Collection) ElementType() reflect.Type {
	return reflect.TypeOf((**S3Collection)(nil)).Elem()
}

func (i *S3Collection) ToS3CollectionOutput() S3CollectionOutput {
	return i.ToS3CollectionOutputWithContext(context.Background())
}

func (i *S3Collection) ToS3CollectionOutputWithContext(ctx context.Context) S3CollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3CollectionOutput)
}

// S3CollectionArrayInput is an input type that accepts S3CollectionArray and S3CollectionArrayOutput values.
// You can construct a concrete instance of `S3CollectionArrayInput` via:
//
//	S3CollectionArray{ S3CollectionArgs{...} }
type S3CollectionArrayInput interface {
	pulumi.Input

	ToS3CollectionArrayOutput() S3CollectionArrayOutput
	ToS3CollectionArrayOutputWithContext(context.Context) S3CollectionArrayOutput
}

type S3CollectionArray []S3CollectionInput

func (S3CollectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*S3Collection)(nil)).Elem()
}

func (i S3CollectionArray) ToS3CollectionArrayOutput() S3CollectionArrayOutput {
	return i.ToS3CollectionArrayOutputWithContext(context.Background())
}

func (i S3CollectionArray) ToS3CollectionArrayOutputWithContext(ctx context.Context) S3CollectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3CollectionArrayOutput)
}

// S3CollectionMapInput is an input type that accepts S3CollectionMap and S3CollectionMapOutput values.
// You can construct a concrete instance of `S3CollectionMapInput` via:
//
//	S3CollectionMap{ "key": S3CollectionArgs{...} }
type S3CollectionMapInput interface {
	pulumi.Input

	ToS3CollectionMapOutput() S3CollectionMapOutput
	ToS3CollectionMapOutputWithContext(context.Context) S3CollectionMapOutput
}

type S3CollectionMap map[string]S3CollectionInput

func (S3CollectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*S3Collection)(nil)).Elem()
}

func (i S3CollectionMap) ToS3CollectionMapOutput() S3CollectionMapOutput {
	return i.ToS3CollectionMapOutputWithContext(context.Background())
}

func (i S3CollectionMap) ToS3CollectionMapOutputWithContext(ctx context.Context) S3CollectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3CollectionMapOutput)
}

type S3CollectionOutput struct{ *pulumi.OutputState }

func (S3CollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**S3Collection)(nil)).Elem()
}

func (o S3CollectionOutput) ToS3CollectionOutput() S3CollectionOutput {
	return o
}

func (o S3CollectionOutput) ToS3CollectionOutputWithContext(ctx context.Context) S3CollectionOutput {
	return o
}

// Text describing the collection.
func (o S3CollectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *S3Collection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Ingest transformation SQL query. Turns the collection into insertOnly mode.
func (o S3CollectionOutput) IngestTransformation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *S3Collection) pulumi.StringPtrOutput { return v.IngestTransformation }).(pulumi.StringPtrOutput)
}

// Unique identifier for the collection. Can contain alphanumeric or dash characters.
func (o S3CollectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *S3Collection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of seconds after which data is purged. Based on event time.
func (o S3CollectionOutput) RetentionSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *S3Collection) pulumi.IntPtrOutput { return v.RetentionSecs }).(pulumi.IntPtrOutput)
}

// Defines a source for this collection.
func (o S3CollectionOutput) Sources() S3CollectionSourceArrayOutput {
	return o.ApplyT(func(v *S3Collection) S3CollectionSourceArrayOutput { return v.Sources }).(S3CollectionSourceArrayOutput)
}

// Wait until the collection is ready.
func (o S3CollectionOutput) WaitForCollection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *S3Collection) pulumi.BoolPtrOutput { return v.WaitForCollection }).(pulumi.BoolPtrOutput)
}

// Wait until the collection has documents. The default is to wait for 0 documents, which means it doesn't wait.
func (o S3CollectionOutput) WaitForDocuments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *S3Collection) pulumi.IntPtrOutput { return v.WaitForDocuments }).(pulumi.IntPtrOutput)
}

// The name of the workspace.
func (o S3CollectionOutput) Workspace() pulumi.StringOutput {
	return o.ApplyT(func(v *S3Collection) pulumi.StringOutput { return v.Workspace }).(pulumi.StringOutput)
}

type S3CollectionArrayOutput struct{ *pulumi.OutputState }

func (S3CollectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*S3Collection)(nil)).Elem()
}

func (o S3CollectionArrayOutput) ToS3CollectionArrayOutput() S3CollectionArrayOutput {
	return o
}

func (o S3CollectionArrayOutput) ToS3CollectionArrayOutputWithContext(ctx context.Context) S3CollectionArrayOutput {
	return o
}

func (o S3CollectionArrayOutput) Index(i pulumi.IntInput) S3CollectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *S3Collection {
		return vs[0].([]*S3Collection)[vs[1].(int)]
	}).(S3CollectionOutput)
}

type S3CollectionMapOutput struct{ *pulumi.OutputState }

func (S3CollectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*S3Collection)(nil)).Elem()
}

func (o S3CollectionMapOutput) ToS3CollectionMapOutput() S3CollectionMapOutput {
	return o
}

func (o S3CollectionMapOutput) ToS3CollectionMapOutputWithContext(ctx context.Context) S3CollectionMapOutput {
	return o
}

func (o S3CollectionMapOutput) MapIndex(k pulumi.StringInput) S3CollectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *S3Collection {
		return vs[0].(map[string]*S3Collection)[vs[1].(string)]
	}).(S3CollectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*S3CollectionInput)(nil)).Elem(), &S3Collection{})
	pulumi.RegisterInputType(reflect.TypeOf((*S3CollectionArrayInput)(nil)).Elem(), S3CollectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*S3CollectionMapInput)(nil)).Elem(), S3CollectionMap{})
	pulumi.RegisterOutputType(S3CollectionOutput{})
	pulumi.RegisterOutputType(S3CollectionArrayOutput{})
	pulumi.RegisterOutputType(S3CollectionMapOutput{})
}
