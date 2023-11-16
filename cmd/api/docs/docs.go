// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Return a json content with a welcome message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "initial"
                ],
                "summary": "Gets a json welcome message",
                "responses": {
                    "200": {
                        "description": "Demo data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Something unidentified has occurred"
                    }
                }
            }
        },
        "/analyze": {
            "post": {
                "description": "Analyze a transactions file and send an email with generated data",
                "tags": [
                    "transaction_analyzer"
                ],
                "summary": "Analyze a transactions file",
                "parameters": [
                    {
                        "description": "Data to process request",
                        "name": "analyze_data_input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/txanalizer.AnalyzeAccountTransactionsInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Confirmation message",
                        "schema": {
                            "$ref": "#/definitions/echo.Map"
                        }
                    },
                    "500": {
                        "description": "Something unidentified has occurred"
                    }
                }
            }
        }
    },
    "definitions": {
        "echo.Map": {
            "type": "object",
            "additionalProperties": true
        },
        "txanalizer.AnalyzeAccountTransactionsInput": {
            "type": "object",
            "properties": {
                "send_to": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "stori_challenge",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
