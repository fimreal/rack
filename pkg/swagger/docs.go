// Package swagger GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package swagger

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
        "/hc": {
            "get": {
                "description": "healthcheck",
                "produces": [
                    "text/plain"
                ],
                "summary": "健康检查",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "healthcheck",
                "produces": [
                    "text/plain"
                ],
                "summary": "健康检查",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ip": {
            "get": {
                "description": "Return client ip",
                "produces": [
                    "text/plain"
                ],
                "summary": "Return client ip",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ipinfo": {
            "get": {
                "description": "Describe client ip",
                "produces": [
                    "application/json"
                ],
                "summary": "Describe client ip",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1.",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "rack",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
