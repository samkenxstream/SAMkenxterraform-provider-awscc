// Code generated by generators/resource/main.go; DO NOT EDIT.

package logs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_logs_log_group", logGroup)
}

// logGroup returns the Terraform aws_logs_log_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Logs::LogGroup resource type.
func logGroup(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The CloudWatch log group ARN.",
			     "type": "string"
			   }
			*/
			Description: `The CloudWatch log group ARN.`,
			Type:        types.StringType,
			Computed:    true,
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
			     "maxLength": 256,
			     "pattern": "^arn:[a-z0-9-]+:kms:[a-z0-9-]+:\\d{12}:(key|alias)/.+$",
			     "type": "string"
			   }
			*/
			Description: `The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.`,
			Type:        types.StringType,
			Optional:    true,
		},
		"log_group_name": {
			// Property: LogGroupName
			// PrimaryIdentifier: true
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.",
			     "maxLength": 512,
			     "minLength": 1,
			     "pattern": "^[.\\-_/#A-Za-z0-9]{1,512}$",
			     "type": "string"
			   }
			*/
			Description: `The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.`,
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// LogGroupName is a force-new attribute.
		},
		"retention_in_days": {
			// Property: RetentionInDays
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653.",
			     "enum": [
			       1,
			       3,
			       5,
			       7,
			       14,
			       30,
			       60,
			       90,
			       120,
			       150,
			       180,
			       365,
			       400,
			       545,
			       731,
			       1827,
			       3653
			     ],
			     "type": "integer"
			   }
			*/
			Description: `The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653.`,
			Type:        types.NumberType,
			Optional:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: `Resource schema for AWS::Logs::LogGroup`,
		Version:     1,
		Attributes:  attributes,
	}

	var features ResourceTypeFeatures

	features |= ResourceTypeHasUpdatableAttribute

	resourceType, err := NewResourceType(
		"AWS::Logs::LogGroup",      // CloudFormation type name
		"aws_logs_log_group",       // Terraform type name
		schema,                     // Terraform schema
		"/properties/LogGroupName", // Primary identifier property path (JSON Pointer)
		[]string{},                 // Write-only property paths (JSON Pointer)
		features,
	)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema for %s:\n\n%v", "aws_logs_log_group", schema)

	return resourceType, nil
}
