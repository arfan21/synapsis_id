// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.synapsis.id"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/carts": {
            "get": {
                "description": "Get Cart By Customer ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Get Cart By Customer ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "With the bearer started",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/github_com_arfan21_synapsis_id_internal_model.GetCartResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create Cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Create Cart",
                "parameters": [
                    {
                        "type": "string",
                        "description": "With the bearer started",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Payload Create Cart Request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_synapsis_id_internal_model.CreateCartRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                        }
                    },
                    "400": {
                        "description": "Error validation field",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "errors": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.ErrValidationResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/customers/login": {
            "post": {
                "description": "Login Customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Login Customer",
                "parameters": [
                    {
                        "description": "Payload Customer Login Request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_synapsis_id_internal_model.CustomerLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/github_com_arfan21_synapsis_id_internal_model.CustomerLoginResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Error validation field",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "errors": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.ErrValidationResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/customers/refresh-token": {
            "post": {
                "description": "Refresh Token Customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Refresh Token Customer",
                "parameters": [
                    {
                        "description": "Payload Customer Refresh Token Request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_synapsis_id_internal_model.CustomerRefreshTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/github_com_arfan21_synapsis_id_internal_model.CustomerLoginResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Error validation field",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "errors": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.ErrValidationResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/customers/register": {
            "post": {
                "description": "Register Customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Register Customer",
                "parameters": [
                    {
                        "description": "Payload Customer Register Request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_synapsis_id_internal_model.CustomerRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                        }
                    },
                    "400": {
                        "description": "Error validation field",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "errors": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.ErrValidationResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/products": {
            "get": {
                "description": "Get Products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get Products",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name of product",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Category id of product",
                        "name": "category_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.PaginationResponse"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "data": {
                                                            "$ref": "#/definitions/github_com_arfan21_synapsis_id_internal_model.GetProductResponse"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create Product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Create Product",
                "parameters": [
                    {
                        "description": "Payload Create Product Request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_synapsis_id_internal_model.ProductCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                        }
                    },
                    "400": {
                        "description": "Error validation field",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "errors": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.ErrValidationResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/products/categories": {
            "get": {
                "description": "Get Product Categories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get Product Categories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/github_com_arfan21_synapsis_id_internal_model.GetCategoryResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_arfan21_synapsis_id_internal_model.CreateCartRequest": {
            "type": "object",
            "required": [
                "customer_id",
                "product_id"
            ],
            "properties": {
                "customer_id": {
                    "type": "string"
                },
                "product_id": {
                    "type": "string"
                }
            }
        },
        "github_com_arfan21_synapsis_id_internal_model.CustomerLoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 8
                }
            }
        },
        "github_com_arfan21_synapsis_id_internal_model.CustomerLoginResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_in": {
                    "type": "integer"
                },
                "expires_in_refresh_token": {
                    "type": "integer"
                },
                "refresh_token": {
                    "type": "string"
                },
                "token_type": {
                    "type": "string"
                }
            }
        },
        "github_com_arfan21_synapsis_id_internal_model.CustomerRefreshTokenRequest": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "github_com_arfan21_synapsis_id_internal_model.CustomerRegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "fullname",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 8
                }
            }
        },
        "github_com_arfan21_synapsis_id_internal_model.GetCartResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "product_id": {
                    "type": "string"
                },
                "product_name": {
                    "type": "string"
                },
                "product_price": {
                    "type": "string"
                },
                "product_stok": {
                    "type": "integer"
                }
            }
        },
        "github_com_arfan21_synapsis_id_internal_model.GetCategoryResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "string": {
                    "type": "string"
                }
            }
        },
        "github_com_arfan21_synapsis_id_internal_model.GetProductResponse": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "string"
                },
                "category_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "owner_id": {
                    "type": "string"
                },
                "owner_name": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "stok": {
                    "type": "integer"
                }
            }
        },
        "github_com_arfan21_synapsis_id_internal_model.ProductCreateRequest": {
            "type": "object",
            "required": [
                "category_id",
                "customer_id",
                "name",
                "price",
                "stok"
            ],
            "properties": {
                "category_id": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "stok": {
                    "type": "integer"
                }
            }
        },
        "github_com_arfan21_synapsis_id_pkg_pkgutil.ErrValidationResponse": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "github_com_arfan21_synapsis_id_pkg_pkgutil.HTTPResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {},
                "errors": {
                    "type": "array",
                    "items": {}
                },
                "message": {
                    "type": "string",
                    "example": "Success"
                },
                "status": {
                    "type": "string",
                    "example": "OK"
                }
            }
        },
        "github_com_arfan21_synapsis_id_pkg_pkgutil.PaginationResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "limit": {
                    "type": "integer",
                    "example": 10
                },
                "page": {
                    "type": "integer",
                    "example": 1
                },
                "total_data": {
                    "type": "integer",
                    "example": 1
                },
                "total_page": {
                    "type": "integer",
                    "example": 1
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8888",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Synapsis ID API",
	Description:      "This is a sample server cell for Synapsis ID Test API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
