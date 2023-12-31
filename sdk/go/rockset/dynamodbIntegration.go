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

// Manages a Rockset DynamoDB Integration.
type DynamodbIntegration struct {
	pulumi.CustomResourceState

	// The AWS Role Arn to use for this integration.
	AwsRoleArn pulumi.StringOutput `pulumi:"awsRoleArn"`
	// Text describing the integration.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Unique identifier for the integration. Can contain alphanumeric or dash characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// AWS S3 bucket name used for exporting the DynamoDB tables.
	S3ExportBucketName pulumi.StringOutput `pulumi:"s3ExportBucketName"`
}

// NewDynamodbIntegration registers a new resource with the given unique name, arguments, and options.
func NewDynamodbIntegration(ctx *pulumi.Context,
	name string, args *DynamodbIntegrationArgs, opts ...pulumi.ResourceOption) (*DynamodbIntegration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AwsRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'AwsRoleArn'")
	}
	if args.S3ExportBucketName == nil {
		return nil, errors.New("invalid value for required argument 'S3ExportBucketName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DynamodbIntegration
	err := ctx.RegisterResource("rockset:index/dynamodbIntegration:DynamodbIntegration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDynamodbIntegration gets an existing DynamodbIntegration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDynamodbIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DynamodbIntegrationState, opts ...pulumi.ResourceOption) (*DynamodbIntegration, error) {
	var resource DynamodbIntegration
	err := ctx.ReadResource("rockset:index/dynamodbIntegration:DynamodbIntegration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DynamodbIntegration resources.
type dynamodbIntegrationState struct {
	// The AWS Role Arn to use for this integration.
	AwsRoleArn *string `pulumi:"awsRoleArn"`
	// Text describing the integration.
	Description *string `pulumi:"description"`
	// Unique identifier for the integration. Can contain alphanumeric or dash characters.
	Name *string `pulumi:"name"`
	// AWS S3 bucket name used for exporting the DynamoDB tables.
	S3ExportBucketName *string `pulumi:"s3ExportBucketName"`
}

type DynamodbIntegrationState struct {
	// The AWS Role Arn to use for this integration.
	AwsRoleArn pulumi.StringPtrInput
	// Text describing the integration.
	Description pulumi.StringPtrInput
	// Unique identifier for the integration. Can contain alphanumeric or dash characters.
	Name pulumi.StringPtrInput
	// AWS S3 bucket name used for exporting the DynamoDB tables.
	S3ExportBucketName pulumi.StringPtrInput
}

func (DynamodbIntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*dynamodbIntegrationState)(nil)).Elem()
}

type dynamodbIntegrationArgs struct {
	// The AWS Role Arn to use for this integration.
	AwsRoleArn string `pulumi:"awsRoleArn"`
	// Text describing the integration.
	Description *string `pulumi:"description"`
	// Unique identifier for the integration. Can contain alphanumeric or dash characters.
	Name *string `pulumi:"name"`
	// AWS S3 bucket name used for exporting the DynamoDB tables.
	S3ExportBucketName string `pulumi:"s3ExportBucketName"`
}

// The set of arguments for constructing a DynamodbIntegration resource.
type DynamodbIntegrationArgs struct {
	// The AWS Role Arn to use for this integration.
	AwsRoleArn pulumi.StringInput
	// Text describing the integration.
	Description pulumi.StringPtrInput
	// Unique identifier for the integration. Can contain alphanumeric or dash characters.
	Name pulumi.StringPtrInput
	// AWS S3 bucket name used for exporting the DynamoDB tables.
	S3ExportBucketName pulumi.StringInput
}

func (DynamodbIntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dynamodbIntegrationArgs)(nil)).Elem()
}

type DynamodbIntegrationInput interface {
	pulumi.Input

	ToDynamodbIntegrationOutput() DynamodbIntegrationOutput
	ToDynamodbIntegrationOutputWithContext(ctx context.Context) DynamodbIntegrationOutput
}

func (*DynamodbIntegration) ElementType() reflect.Type {
	return reflect.TypeOf((**DynamodbIntegration)(nil)).Elem()
}

func (i *DynamodbIntegration) ToDynamodbIntegrationOutput() DynamodbIntegrationOutput {
	return i.ToDynamodbIntegrationOutputWithContext(context.Background())
}

func (i *DynamodbIntegration) ToDynamodbIntegrationOutputWithContext(ctx context.Context) DynamodbIntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynamodbIntegrationOutput)
}

// DynamodbIntegrationArrayInput is an input type that accepts DynamodbIntegrationArray and DynamodbIntegrationArrayOutput values.
// You can construct a concrete instance of `DynamodbIntegrationArrayInput` via:
//
//	DynamodbIntegrationArray{ DynamodbIntegrationArgs{...} }
type DynamodbIntegrationArrayInput interface {
	pulumi.Input

	ToDynamodbIntegrationArrayOutput() DynamodbIntegrationArrayOutput
	ToDynamodbIntegrationArrayOutputWithContext(context.Context) DynamodbIntegrationArrayOutput
}

type DynamodbIntegrationArray []DynamodbIntegrationInput

func (DynamodbIntegrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DynamodbIntegration)(nil)).Elem()
}

func (i DynamodbIntegrationArray) ToDynamodbIntegrationArrayOutput() DynamodbIntegrationArrayOutput {
	return i.ToDynamodbIntegrationArrayOutputWithContext(context.Background())
}

func (i DynamodbIntegrationArray) ToDynamodbIntegrationArrayOutputWithContext(ctx context.Context) DynamodbIntegrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynamodbIntegrationArrayOutput)
}

// DynamodbIntegrationMapInput is an input type that accepts DynamodbIntegrationMap and DynamodbIntegrationMapOutput values.
// You can construct a concrete instance of `DynamodbIntegrationMapInput` via:
//
//	DynamodbIntegrationMap{ "key": DynamodbIntegrationArgs{...} }
type DynamodbIntegrationMapInput interface {
	pulumi.Input

	ToDynamodbIntegrationMapOutput() DynamodbIntegrationMapOutput
	ToDynamodbIntegrationMapOutputWithContext(context.Context) DynamodbIntegrationMapOutput
}

type DynamodbIntegrationMap map[string]DynamodbIntegrationInput

func (DynamodbIntegrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DynamodbIntegration)(nil)).Elem()
}

func (i DynamodbIntegrationMap) ToDynamodbIntegrationMapOutput() DynamodbIntegrationMapOutput {
	return i.ToDynamodbIntegrationMapOutputWithContext(context.Background())
}

func (i DynamodbIntegrationMap) ToDynamodbIntegrationMapOutputWithContext(ctx context.Context) DynamodbIntegrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynamodbIntegrationMapOutput)
}

type DynamodbIntegrationOutput struct{ *pulumi.OutputState }

func (DynamodbIntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DynamodbIntegration)(nil)).Elem()
}

func (o DynamodbIntegrationOutput) ToDynamodbIntegrationOutput() DynamodbIntegrationOutput {
	return o
}

func (o DynamodbIntegrationOutput) ToDynamodbIntegrationOutputWithContext(ctx context.Context) DynamodbIntegrationOutput {
	return o
}

// The AWS Role Arn to use for this integration.
func (o DynamodbIntegrationOutput) AwsRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *DynamodbIntegration) pulumi.StringOutput { return v.AwsRoleArn }).(pulumi.StringOutput)
}

// Text describing the integration.
func (o DynamodbIntegrationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DynamodbIntegration) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Unique identifier for the integration. Can contain alphanumeric or dash characters.
func (o DynamodbIntegrationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DynamodbIntegration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// AWS S3 bucket name used for exporting the DynamoDB tables.
func (o DynamodbIntegrationOutput) S3ExportBucketName() pulumi.StringOutput {
	return o.ApplyT(func(v *DynamodbIntegration) pulumi.StringOutput { return v.S3ExportBucketName }).(pulumi.StringOutput)
}

type DynamodbIntegrationArrayOutput struct{ *pulumi.OutputState }

func (DynamodbIntegrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DynamodbIntegration)(nil)).Elem()
}

func (o DynamodbIntegrationArrayOutput) ToDynamodbIntegrationArrayOutput() DynamodbIntegrationArrayOutput {
	return o
}

func (o DynamodbIntegrationArrayOutput) ToDynamodbIntegrationArrayOutputWithContext(ctx context.Context) DynamodbIntegrationArrayOutput {
	return o
}

func (o DynamodbIntegrationArrayOutput) Index(i pulumi.IntInput) DynamodbIntegrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DynamodbIntegration {
		return vs[0].([]*DynamodbIntegration)[vs[1].(int)]
	}).(DynamodbIntegrationOutput)
}

type DynamodbIntegrationMapOutput struct{ *pulumi.OutputState }

func (DynamodbIntegrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DynamodbIntegration)(nil)).Elem()
}

func (o DynamodbIntegrationMapOutput) ToDynamodbIntegrationMapOutput() DynamodbIntegrationMapOutput {
	return o
}

func (o DynamodbIntegrationMapOutput) ToDynamodbIntegrationMapOutputWithContext(ctx context.Context) DynamodbIntegrationMapOutput {
	return o
}

func (o DynamodbIntegrationMapOutput) MapIndex(k pulumi.StringInput) DynamodbIntegrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DynamodbIntegration {
		return vs[0].(map[string]*DynamodbIntegration)[vs[1].(string)]
	}).(DynamodbIntegrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DynamodbIntegrationInput)(nil)).Elem(), &DynamodbIntegration{})
	pulumi.RegisterInputType(reflect.TypeOf((*DynamodbIntegrationArrayInput)(nil)).Elem(), DynamodbIntegrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DynamodbIntegrationMapInput)(nil)).Elem(), DynamodbIntegrationMap{})
	pulumi.RegisterOutputType(DynamodbIntegrationOutput{})
	pulumi.RegisterOutputType(DynamodbIntegrationArrayOutput{})
	pulumi.RegisterOutputType(DynamodbIntegrationMapOutput{})
}
