package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkStackProps struct {
	awscdk.StackProps
}

func NewCdkStack(scope constructs.Construct, id string, props *CdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Define S3 bucket.
	bucketName := "apollo"
	s3Bucket := awss3.NewBucket(stack, jsii.String("S3Bucket"), &awss3.BucketProps{
		BucketName:        jsii.String(bucketName),
		AutoDeleteObjects: jsii.Bool(true),
		BlockPublicAccess: awss3.NewBlockPublicAccess(&awss3.BlockPublicAccessOptions{
			BlockPublicAcls:       jsii.Bool(false),
			BlockPublicPolicy:     jsii.Bool(false),
			IgnorePublicAcls:      jsii.Bool(false),
			RestrictPublicBuckets: jsii.Bool(false),
		}),
		AccessControl: awss3.BucketAccessControl_BUCKET_OWNER_FULL_CONTROL,
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	// Define bucket policy to allow public read access
	s3Bucket.AddToResourcePolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: jsii.Strings("s3:GetObject"),
		Effect:  awsiam.Effect_ALLOW,
		Principals: &[]awsiam.IPrincipal{
			awsiam.NewAnyPrincipal(),
		},
		Resources: jsii.Strings(*s3Bucket.BucketArn() + "/*"),
	}))

	// Define dynamoDB tables.
	productTable := awsdynamodb.NewTable(stack, jsii.String("ProductTable"), &awsdynamodb.TableProps{
		TableName:     jsii.String("products"),
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("uuid"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		SortKey: &awsdynamodb.Attribute{
			Name: jsii.String("category_uuid"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})

	// categoryTable := awsdynamodb.NewTable(stack, jsii.String("CategoryTable"), &awsdynamodb.TableProps{
	// 	TableName:     jsii.String("category"),
	// 	RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	// 	PartitionKey: &awsdynamodb.Attribute{
	// 		Name: jsii.String("uuid"),
	// 		Type: awsdynamodb.AttributeType_STRING,
	// 	},
	// })

	// Define lambda function.
	lambdaFunc := awslambda.NewFunction(stack, jsii.String("LambdaFunction"), &awslambda.FunctionProps{
		FunctionName: jsii.String("products"),
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Architecture: awslambda.Architecture_ARM_64(),
		Code:         awslambda.AssetCode_FromAsset(jsii.String("./lambda/bin/apollo.zip"), nil),
		Handler:      jsii.String("main"),
		MemorySize:   jsii.Number(128),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(60)),
		Environment: &map[string]*string{
			"LOG_LEVEL": jsii.String("info"),
			"BUCKET":    jsii.String(bucketName),
		},
	})

	// Grant lambda function read/write access to the table.
	productTable.GrantReadWriteData(lambdaFunc)
	s3Bucket.GrantReadWrite(lambdaFunc, nil)

	// Define API Gateway REST API.
	restapi := awsapigateway.NewRestApi(stack, jsii.String("ApiGateway"), &awsapigateway.RestApiProps{
		RestApiName: jsii.String("apollo"),
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowHeaders: jsii.Strings("Content-Type", "Authorization"),
			AllowMethods: jsii.Strings("GET", "POST", "DELETE", "PUT", "OPTIONS"),
			AllowOrigins: jsii.Strings("*"),
		},
		DeployOptions: &awsapigateway.StageOptions{
			StageName: jsii.String("dev"),
		},
	})

	// Integrate the lambda function to the API Gateway.
	integration := awsapigateway.NewLambdaIntegration(lambdaFunc, nil)

	// Define the API Gateway routes.
	apiResource := restapi.Root().AddResource(jsii.String("api"), nil)
	v1Resource := apiResource.AddResource(jsii.String("v1"), nil)
	productsResource := v1Resource.AddResource(jsii.String("products"), nil)
	productsResource.AddMethod(jsii.String("POST"), integration, nil)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCdkStack(app, "ApolloStack", &CdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
