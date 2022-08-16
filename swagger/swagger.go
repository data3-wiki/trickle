package swagger

import (
	"fmt"

	"github.com/dereference-xyz/trickle/model"
	"github.com/go-openapi/jsonreference"
	"github.com/go-openapi/spec"
)

// Generate a SwaggerV2 spec based on the given program type.
func Generate(programType *model.ProgramType) (*spec.Swagger, error) {
	paths := make(map[string]spec.PathItem)
	for _, accountType := range programType.AccountTypes {
		endpoint := fmt.Sprintf("/api/v1/solana/account/read/%s", accountType.Name)
		response, err := generateResponse(accountType)
		if err != nil {
			return nil, err
		}
		paths[endpoint] = spec.PathItem{
			PathItemProps: spec.PathItemProps{
				Get: &spec.Operation{
					OperationProps: spec.OperationProps{
						Produces: []string{"application/json"},
						Responses: &spec.Responses{
							ResponsesProps: spec.ResponsesProps{
								StatusCodeResponses: map[int]spec.Response{
									200: *response,
								},
							},
						},
						Parameters: generateParameters(accountType),
						Tags:       []string{"account"},
					},
				},
			},
		}
	}

	return &spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Swagger: "2.0",
			Paths: &spec.Paths{
				Paths: paths,
			},
			Definitions: generateDefinitions(programType),
		},
	}, nil
}

// Generate Swagger model definitions based on the program type.
func generateDefinitions(programType *model.ProgramType) spec.Definitions {
	defs := make(map[string]spec.Schema)
	for _, accountType := range programType.AccountTypes {
		properties := make(map[string]spec.Schema)
		for _, propertyType := range accountType.PropertyTypes {
			properties[propertyType.Name] = *typeOf(swaggerType(propertyType.DataType))
		}
		defs[accountType.Name] = spec.Schema{
			SchemaProps: spec.SchemaProps{
				Properties: properties,
			},
		}
	}
	return defs
}

// Generate Swagger response based on account type.
func generateResponse(accountType *model.AccountType) (*spec.Response, error) {
	accountRef, err := refOf(accountType.Name)
	if err != nil {
		return nil, err
	}

	return &spec.Response{
		ResponseProps: spec.ResponseProps{
			Description: "Array of account data for accounts matching the field predicates (if specified).",
			Schema: objectOf(
				spec.SchemaProperties{
					"accounts": *arrayOf(objectOf(
						spec.SchemaProperties{
							"type": *typeOf("string"),
							"data": *accountRef,
						})),
				}),
		},
	}, nil
}

// Generate Swagger parameters based on account type.
func generateParameters(accountType *model.AccountType) []spec.Parameter {
	parameters := []spec.Parameter{}
	for _, propertyType := range accountType.PropertyTypes {
		parameters = append(parameters, spec.Parameter{
			ParamProps: spec.ParamProps{
				In:       "query",
				Name:     propertyType.Name,
				Schema:   typeOf(swaggerType(propertyType.DataType)),
				Required: false,
			},
		})
	}
	return parameters
}

// Helper function to return a Swagger primitive type.
func typeOf(name string) *spec.Schema {
	return &spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type: []string{name},
		},
	}
}

// Helper function to return a Swagger array type.
func arrayOf(schema *spec.Schema) *spec.Schema {
	return &spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type: []string{"array"},
			Items: &spec.SchemaOrArray{
				Schema: schema,
			},
		},
	}
}

// Helper function to return a Swagger object type.
func objectOf(properties spec.SchemaProperties) *spec.Schema {
	return &spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type:       []string{"object"},
			Properties: properties,
		},
	}
}

// Helper function to return a Swagger ref.
func refOf(definition string) (*spec.Schema, error) {
	accountRef, err := jsonreference.New(fmt.Sprintf("#/definitions/%s", definition))
	if err != nil {
		return nil, err
	}
	return &spec.Schema{
		SchemaProps: spec.SchemaProps{
			Ref: spec.Ref{
				Ref: accountRef,
			},
		},
	}, nil
}
