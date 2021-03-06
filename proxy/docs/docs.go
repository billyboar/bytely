// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/urls": {
            "post": {
                "description": "This endpoint receives a URL and returns a shortened URL.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "receives a URL and returns a shortened URL.",
                "parameters": [
                    {
                        "description": "URL to shorten",
                        "name": "url",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/proxy.shortenURLRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/proxy.shortenURLResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/urls/{short_url}": {
            "get": {
                "description": "This endpoint returns click count for a shortened URL.",
                "produces": [
                    "application/json"
                ],
                "summary": "returns click count for a shortened URL.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Short URL to get stats",
                        "name": "short_url",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/proxy.getURLStatsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "This endpoint deletes a shortened URL.",
                "produces": [
                    "application/json"
                ],
                "summary": "deletes a shortened URL.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Short URL to delete",
                        "name": "short_url",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/{short_url}": {
            "get": {
                "description": "This endpoint find a correspoding URL for a short URL and redirects to it.",
                "summary": "find a correspoding URL for a short URL and redirects to it.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Short URL to redirect to the original URL",
                        "name": "short_url",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "proxy.getURLStatsResponse": {
            "type": "object",
            "properties": {
                "clicks": {
                    "type": "integer"
                }
            }
        },
        "proxy.shortenURLRequest": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "proxy.shortenURLResponse": {
            "type": "object",
            "properties": {
                "short_url": {
                    "type": "string"
                }
            }
        },
        "schema.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
