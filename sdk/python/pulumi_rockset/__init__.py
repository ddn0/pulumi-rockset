# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .alias import *
from .api_key import *
from .collection import *
from .collection_mount import *
from .dynamodb_collection import *
from .dynamodb_integration import *
from .gcs_collection import *
from .gcs_integration import *
from .get_account import *
from .get_query_lambda import *
from .get_query_lambda_tag import *
from .get_user import *
from .get_virtual_instance import *
from .get_workspace import *
from .kafka_collection import *
from .kafka_integration import *
from .kinesis_collection import *
from .kinesis_integration import *
from .mongodb_collection import *
from .mongodb_integration import *
from .provider import *
from .query_lambda import *
from .query_lambda_tag import *
from .role import *
from .s3_collection import *
from .s3_integration import *
from .user import *
from .view import *
from .virtual_instance import *
from .workspace import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_rockset.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_rockset.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "rockset",
  "mod": "index/alias",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/alias:Alias": "Alias"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/apiKey",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/apiKey:ApiKey": "ApiKey"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/collection",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/collection:Collection": "Collection"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/collectionMount",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/collectionMount:CollectionMount": "CollectionMount"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/dynamodbCollection",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/dynamodbCollection:DynamodbCollection": "DynamodbCollection"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/dynamodbIntegration",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/dynamodbIntegration:DynamodbIntegration": "DynamodbIntegration"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/gcsCollection",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/gcsCollection:GcsCollection": "GcsCollection"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/gcsIntegration",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/gcsIntegration:GcsIntegration": "GcsIntegration"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/kafkaCollection",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/kafkaCollection:KafkaCollection": "KafkaCollection"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/kafkaIntegration",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/kafkaIntegration:KafkaIntegration": "KafkaIntegration"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/kinesisCollection",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/kinesisCollection:KinesisCollection": "KinesisCollection"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/kinesisIntegration",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/kinesisIntegration:KinesisIntegration": "KinesisIntegration"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/mongodbCollection",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/mongodbCollection:MongodbCollection": "MongodbCollection"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/mongodbIntegration",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/mongodbIntegration:MongodbIntegration": "MongodbIntegration"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/queryLambda",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/queryLambda:QueryLambda": "QueryLambda"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/queryLambdaTag",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/queryLambdaTag:QueryLambdaTag": "QueryLambdaTag"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/role",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/role:Role": "Role"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/s3Collection",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/s3Collection:S3Collection": "S3Collection"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/s3Integration",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/s3Integration:S3Integration": "S3Integration"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/user",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/user:User": "User"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/view",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/view:View": "View"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/virtualInstance",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/virtualInstance:VirtualInstance": "VirtualInstance"
  }
 },
 {
  "pkg": "rockset",
  "mod": "index/workspace",
  "fqn": "pulumi_rockset",
  "classes": {
   "rockset:index/workspace:Workspace": "Workspace"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "rockset",
  "token": "pulumi:providers:rockset",
  "fqn": "pulumi_rockset",
  "class": "Provider"
 }
]
"""
)
