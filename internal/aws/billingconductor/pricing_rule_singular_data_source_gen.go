// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package billingconductor

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_billingconductor_pricing_rule", pricingRuleDataSource)
}

// pricingRuleDataSource returns the Terraform awscc_billingconductor_pricing_rule data source.
// This Terraform data source corresponds to the CloudFormation AWS::BillingConductor::PricingRule resource.
func pricingRuleDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Pricing rule ARN",
		//	  "pattern": "arn:aws(-cn)?:billingconductor::[0-9]{12}:pricingrule/[a-zA-Z0-9]{10}",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Pricing rule ARN",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AssociatedPricingPlanCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of pricing plans associated with pricing rule",
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"associated_pricing_plan_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of pricing plans associated with pricing rule",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BillingEntity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The seller of services provided by AWS, their affiliates, or third-party providers selling services via AWS Marketplaces. Supported billing entities are AWS, AWS Marketplace, and AISPL.",
		//	  "enum": [
		//	    "AWS",
		//	    "AWS Marketplace",
		//	    "AISPL"
		//	  ],
		//	  "type": "string"
		//	}
		"billing_entity": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The seller of services provided by AWS, their affiliates, or third-party providers selling services via AWS Marketplaces. Supported billing entities are AWS, AWS Marketplace, and AISPL.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Creation timestamp in UNIX epoch time format",
		//	  "type": "integer"
		//	}
		"creation_time": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Creation timestamp in UNIX epoch time format",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Pricing rule description",
		//	  "maxLength": 1024,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Pricing rule description",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastModifiedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Latest modified timestamp in UNIX epoch time format",
		//	  "type": "integer"
		//	}
		"last_modified_time": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Latest modified timestamp in UNIX epoch time format",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ModifierPercentage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Pricing rule modifier percentage",
		//	  "minimum": 0,
		//	  "type": "number"
		//	}
		"modifier_percentage": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "Pricing rule modifier percentage",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Pricing rule name",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9_\\+=\\.\\-@]+",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Pricing rule name",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Operation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Operation which a SKU pricing rule is modifying",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^\\S+$",
		//	  "type": "string"
		//	}
		"operation": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Operation which a SKU pricing rule is modifying",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Scope
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A term used to categorize the granularity of a Pricing Rule.",
		//	  "enum": [
		//	    "GLOBAL",
		//	    "SERVICE",
		//	    "BILLING_ENTITY",
		//	    "SKU"
		//	  ],
		//	  "type": "string"
		//	}
		"scope": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A term used to categorize the granularity of a Pricing Rule.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Service
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The service which a pricing rule is applied on",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9\\.\\-]+",
		//	  "type": "string"
		//	}
		"service": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The service which a pricing rule is applied on",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tiering
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The set of tiering configurations for the pricing rule.",
		//	  "properties": {
		//	    "FreeTier": {
		//	      "additionalProperties": false,
		//	      "description": "The possible customizable free tier configurations.",
		//	      "properties": {
		//	        "Activated": {
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "required": [
		//	        "Activated"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tiering": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: FreeTier
				"free_tier": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Activated
						"activated": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The possible customizable free tier configurations.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The set of tiering configurations for the pricing rule.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One of MARKUP, DISCOUNT or TIERING that describes the behaviour of the pricing rule.",
		//	  "enum": [
		//	    "MARKUP",
		//	    "DISCOUNT",
		//	    "TIERING"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "One of MARKUP, DISCOUNT or TIERING that describes the behaviour of the pricing rule.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UsageType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The UsageType which a SKU pricing rule is modifying",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^\\S+$",
		//	  "type": "string"
		//	}
		"usage_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The UsageType which a SKU pricing rule is modifying",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::BillingConductor::PricingRule",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::BillingConductor::PricingRule").WithTerraformTypeName("awscc_billingconductor_pricing_rule")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"activated":                     "Activated",
		"arn":                           "Arn",
		"associated_pricing_plan_count": "AssociatedPricingPlanCount",
		"billing_entity":                "BillingEntity",
		"creation_time":                 "CreationTime",
		"description":                   "Description",
		"free_tier":                     "FreeTier",
		"key":                           "Key",
		"last_modified_time":            "LastModifiedTime",
		"modifier_percentage":           "ModifierPercentage",
		"name":                          "Name",
		"operation":                     "Operation",
		"scope":                         "Scope",
		"service":                       "Service",
		"tags":                          "Tags",
		"tiering":                       "Tiering",
		"type":                          "Type",
		"usage_type":                    "UsageType",
		"value":                         "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
