// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rockset

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/ddn0/pulumi-rockset/sdk/go/rockset/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "rockset:index/alias:Alias":
		r = &Alias{}
	case "rockset:index/apiKey:ApiKey":
		r = &ApiKey{}
	case "rockset:index/collection:Collection":
		r = &Collection{}
	case "rockset:index/collectionMount:CollectionMount":
		r = &CollectionMount{}
	case "rockset:index/dynamodbCollection:DynamodbCollection":
		r = &DynamodbCollection{}
	case "rockset:index/dynamodbIntegration:DynamodbIntegration":
		r = &DynamodbIntegration{}
	case "rockset:index/gcsCollection:GcsCollection":
		r = &GcsCollection{}
	case "rockset:index/gcsIntegration:GcsIntegration":
		r = &GcsIntegration{}
	case "rockset:index/kafkaCollection:KafkaCollection":
		r = &KafkaCollection{}
	case "rockset:index/kafkaIntegration:KafkaIntegration":
		r = &KafkaIntegration{}
	case "rockset:index/kinesisCollection:KinesisCollection":
		r = &KinesisCollection{}
	case "rockset:index/kinesisIntegration:KinesisIntegration":
		r = &KinesisIntegration{}
	case "rockset:index/mongodbCollection:MongodbCollection":
		r = &MongodbCollection{}
	case "rockset:index/mongodbIntegration:MongodbIntegration":
		r = &MongodbIntegration{}
	case "rockset:index/queryLambda:QueryLambda":
		r = &QueryLambda{}
	case "rockset:index/queryLambdaTag:QueryLambdaTag":
		r = &QueryLambdaTag{}
	case "rockset:index/role:Role":
		r = &Role{}
	case "rockset:index/s3Collection:S3Collection":
		r = &S3Collection{}
	case "rockset:index/s3Integration:S3Integration":
		r = &S3Integration{}
	case "rockset:index/user:User":
		r = &User{}
	case "rockset:index/view:View":
		r = &View{}
	case "rockset:index/virtualInstance:VirtualInstance":
		r = &VirtualInstance{}
	case "rockset:index/workspace:Workspace":
		r = &Workspace{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:rockset" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"rockset",
		"index/alias",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/apiKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/collection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/collectionMount",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/dynamodbCollection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/dynamodbIntegration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/gcsCollection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/gcsIntegration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/kafkaCollection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/kafkaIntegration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/kinesisCollection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/kinesisIntegration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/mongodbCollection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/mongodbIntegration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/queryLambda",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/queryLambdaTag",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/role",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/s3Collection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/s3Integration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/view",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/virtualInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rockset",
		"index/workspace",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"rockset",
		&pkg{version},
	)
}
