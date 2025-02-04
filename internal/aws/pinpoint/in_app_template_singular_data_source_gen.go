// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_pinpoint_in_app_template", inAppTemplateDataSource)
}

// inAppTemplateDataSource returns the Terraform awscc_pinpoint_in_app_template data source.
// This Terraform data source corresponds to the CloudFormation AWS::Pinpoint::InAppTemplate resource.
func inAppTemplateDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Content
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "BackgroundColor": {
		//	        "type": "string"
		//	      },
		//	      "BodyConfig": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Alignment": {
		//	            "enum": [
		//	              "LEFT",
		//	              "CENTER",
		//	              "RIGHT"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "Body": {
		//	            "type": "string"
		//	          },
		//	          "TextColor": {
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "HeaderConfig": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Alignment": {
		//	            "enum": [
		//	              "LEFT",
		//	              "CENTER",
		//	              "RIGHT"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "Header": {
		//	            "type": "string"
		//	          },
		//	          "TextColor": {
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "ImageUrl": {
		//	        "type": "string"
		//	      },
		//	      "PrimaryBtn": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Android": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "ButtonAction": {
		//	                "enum": [
		//	                  "LINK",
		//	                  "DEEP_LINK",
		//	                  "CLOSE"
		//	                ],
		//	                "type": "string"
		//	              },
		//	              "Link": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "DefaultConfig": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "BackgroundColor": {
		//	                "type": "string"
		//	              },
		//	              "BorderRadius": {
		//	                "type": "integer"
		//	              },
		//	              "ButtonAction": {
		//	                "enum": [
		//	                  "LINK",
		//	                  "DEEP_LINK",
		//	                  "CLOSE"
		//	                ],
		//	                "type": "string"
		//	              },
		//	              "Link": {
		//	                "type": "string"
		//	              },
		//	              "Text": {
		//	                "type": "string"
		//	              },
		//	              "TextColor": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "IOS": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "ButtonAction": {
		//	                "enum": [
		//	                  "LINK",
		//	                  "DEEP_LINK",
		//	                  "CLOSE"
		//	                ],
		//	                "type": "string"
		//	              },
		//	              "Link": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "Web": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "ButtonAction": {
		//	                "enum": [
		//	                  "LINK",
		//	                  "DEEP_LINK",
		//	                  "CLOSE"
		//	                ],
		//	                "type": "string"
		//	              },
		//	              "Link": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "SecondaryBtn": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Android": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "ButtonAction": {
		//	                "enum": [
		//	                  "LINK",
		//	                  "DEEP_LINK",
		//	                  "CLOSE"
		//	                ],
		//	                "type": "string"
		//	              },
		//	              "Link": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "DefaultConfig": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "BackgroundColor": {
		//	                "type": "string"
		//	              },
		//	              "BorderRadius": {
		//	                "type": "integer"
		//	              },
		//	              "ButtonAction": {
		//	                "enum": [
		//	                  "LINK",
		//	                  "DEEP_LINK",
		//	                  "CLOSE"
		//	                ],
		//	                "type": "string"
		//	              },
		//	              "Link": {
		//	                "type": "string"
		//	              },
		//	              "Text": {
		//	                "type": "string"
		//	              },
		//	              "TextColor": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "IOS": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "ButtonAction": {
		//	                "enum": [
		//	                  "LINK",
		//	                  "DEEP_LINK",
		//	                  "CLOSE"
		//	                ],
		//	                "type": "string"
		//	              },
		//	              "Link": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "Web": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "ButtonAction": {
		//	                "enum": [
		//	                  "LINK",
		//	                  "DEEP_LINK",
		//	                  "CLOSE"
		//	                ],
		//	                "type": "string"
		//	              },
		//	              "Link": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          }
		//	        },
		//	        "type": "object"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"content": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: BackgroundColor
					"background_color": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: BodyConfig
					"body_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Alignment
							"alignment": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Body
							"body": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: TextColor
							"text_color": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: HeaderConfig
					"header_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Alignment
							"alignment": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Header
							"header": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: TextColor
							"text_color": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: ImageUrl
					"image_url": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: PrimaryBtn
					"primary_btn": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Android
							"android": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: ButtonAction
									"button_action": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: Link
									"link": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: DefaultConfig
							"default_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: BackgroundColor
									"background_color": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: BorderRadius
									"border_radius": schema.Int64Attribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: ButtonAction
									"button_action": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: Link
									"link": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: Text
									"text": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: TextColor
									"text_color": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: IOS
							"ios": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: ButtonAction
									"button_action": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: Link
									"link": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Web
							"web": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: ButtonAction
									"button_action": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: Link
									"link": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: SecondaryBtn
					"secondary_btn": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Android
							"android": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: ButtonAction
									"button_action": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: Link
									"link": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: DefaultConfig
							"default_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: BackgroundColor
									"background_color": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: BorderRadius
									"border_radius": schema.Int64Attribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: ButtonAction
									"button_action": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: Link
									"link": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: Text
									"text": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: TextColor
									"text_color": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: IOS
							"ios": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: ButtonAction
									"button_action": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: Link
									"link": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Web
							"web": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: ButtonAction
									"button_action": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: Link
									"link": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CustomConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "object"
		//	}
		"custom_config": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Layout
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "BOTTOM_BANNER",
		//	    "TOP_BANNER",
		//	    "OVERLAYS",
		//	    "MOBILE_FEED",
		//	    "MIDDLE_BANNER",
		//	    "CAROUSEL"
		//	  ],
		//	  "type": "string"
		//	}
		"layout": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "object"
		//	}
		"tags": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TemplateDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"template_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: TemplateName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"template_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Pinpoint::InAppTemplate",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Pinpoint::InAppTemplate").WithTerraformTypeName("awscc_pinpoint_in_app_template")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alignment":            "Alignment",
		"android":              "Android",
		"arn":                  "Arn",
		"background_color":     "BackgroundColor",
		"body":                 "Body",
		"body_config":          "BodyConfig",
		"border_radius":        "BorderRadius",
		"button_action":        "ButtonAction",
		"content":              "Content",
		"custom_config":        "CustomConfig",
		"default_config":       "DefaultConfig",
		"header":               "Header",
		"header_config":        "HeaderConfig",
		"image_url":            "ImageUrl",
		"ios":                  "IOS",
		"layout":               "Layout",
		"link":                 "Link",
		"primary_btn":          "PrimaryBtn",
		"secondary_btn":        "SecondaryBtn",
		"tags":                 "Tags",
		"template_description": "TemplateDescription",
		"template_name":        "TemplateName",
		"text":                 "Text",
		"text_color":           "TextColor",
		"web":                  "Web",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
