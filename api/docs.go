// Code generated by swaggo/swag. DO NOT EDIT.

package api

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
        "/v1/items/{id}/price": {
            "get": {
                "description": "Endpoint to get OSRS Item by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Provides an endpoint to get OSRS Item by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Runescape Item Price",
                        "schema": {
                            "$ref": "#/definitions/osrs-items-api_pkg_oldschoolrs.ItemPriceRecord"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "osrs-items-api_pkg_oldschoolrs.ItemPriceRecord": {
            "type": "object",
            "properties": {
                "average": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/osrs-items-api_pkg_oldschoolrs.Price"
                    }
                },
                "daily": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/osrs-items-api_pkg_oldschoolrs.Price"
                    }
                }
            }
        },
        "osrs-items-api_pkg_oldschoolrs.Price": {
            "type": "object",
            "properties": {
                "timestamp": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
