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
        "license": {
            "name": "WTFPL 2.0",
            "url": "http://www.wtfpl.net/txt/copying/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/users": {
            "get": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "Search users by some filters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Search users",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page size",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetUsersResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new user with set password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "New user data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.NewEntityResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Logs user in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Log in",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/logout": {
            "post": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "Logs user out",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Log out",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "Get user info by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get full user info",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "Update user information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Edit user info",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated user data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.EditUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "Delete user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/{id}/password": {
            "put": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "Change user password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Change user password",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "New password data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PasswordData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreateUserRequest": {
            "type": "object",
            "required": [
                "birthday",
                "country",
                "currency",
                "gender",
                "language",
                "login",
                "name",
                "password",
                "timezone"
            ],
            "properties": {
                "about": {
                    "description": "Short user self-description",
                    "type": "string",
                    "maxLength": 1000,
                    "example": "Some useful and interesting info"
                },
                "birthday": {
                    "description": "User birthday",
                    "type": "string",
                    "example": "2022-01-01"
                },
                "country": {
                    "description": "User country code",
                    "type": "string",
                    "example": "RU"
                },
                "country_region": {
                    "description": "User country region code",
                    "type": "string",
                    "example": "RU-PER"
                },
                "currency": {
                    "description": "User preferred currency code",
                    "type": "string",
                    "example": "RUB"
                },
                "email": {
                    "description": "User main e-mail",
                    "type": "string",
                    "maxLength": 320,
                    "example": "example@example.com"
                },
                "gender": {
                    "description": "User gender",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Gender"
                        }
                    ],
                    "example": "MALE"
                },
                "language": {
                    "description": "User preferred language code",
                    "type": "string",
                    "example": "ru_RU"
                },
                "login": {
                    "description": "User login",
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 3,
                    "example": "user"
                },
                "name": {
                    "description": "User full name",
                    "type": "string",
                    "maxLength": 1500,
                    "example": "User Usersson"
                },
                "password": {
                    "description": "User password",
                    "type": "string",
                    "minLength": 5,
                    "example": "password"
                },
                "phone": {
                    "description": "User main phone number",
                    "type": "string",
                    "example": "+78005553535"
                },
                "timezone": {
                    "description": "User timezone",
                    "type": "string",
                    "example": "UTC"
                }
            }
        },
        "models.EditUserRequest": {
            "type": "object",
            "required": [
                "birthday",
                "country",
                "currency",
                "gender",
                "language",
                "login",
                "name",
                "timezone"
            ],
            "properties": {
                "about": {
                    "description": "Short user self-description",
                    "type": "string",
                    "maxLength": 1000,
                    "example": "Some useful and interesting info"
                },
                "birthday": {
                    "description": "User birthday",
                    "type": "string",
                    "example": "2022-01-01"
                },
                "country": {
                    "description": "User country code",
                    "type": "string",
                    "example": "RU"
                },
                "country_region": {
                    "description": "User country region code",
                    "type": "string",
                    "example": "RU-PER"
                },
                "currency": {
                    "description": "User preferred currency code",
                    "type": "string",
                    "example": "RUB"
                },
                "email": {
                    "description": "User main e-mail",
                    "type": "string",
                    "maxLength": 320,
                    "example": "example@example.com"
                },
                "gender": {
                    "description": "User gender",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Gender"
                        }
                    ],
                    "example": "MALE"
                },
                "language": {
                    "description": "User preferred language code",
                    "type": "string",
                    "example": "ru_RU"
                },
                "login": {
                    "description": "User login",
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 3,
                    "example": "user"
                },
                "name": {
                    "description": "User full name",
                    "type": "string",
                    "maxLength": 1500,
                    "example": "User Usersson"
                },
                "phone": {
                    "description": "User main phone number",
                    "type": "string",
                    "example": "+78005553535"
                },
                "timezone": {
                    "description": "User timezone",
                    "type": "string",
                    "example": "UTC"
                }
            }
        },
        "models.Gender": {
            "type": "string",
            "enum": [
                "",
                "MALE",
                "FEMALE"
            ],
            "x-enum-varnames": [
                "Undefined",
                "Male",
                "Female"
            ]
        },
        "models.GetUsersResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "description": "Found items",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.UserItem"
                    }
                },
                "page_number": {
                    "description": "Number of returned page",
                    "type": "integer"
                },
                "page_size": {
                    "description": "Max page size",
                    "type": "integer"
                },
                "total_items": {
                    "description": "Total count of found items",
                    "type": "integer"
                }
            }
        },
        "models.NewEntityResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.PasswordData": {
            "description": "Password definition",
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "password": {
                    "description": "User password",
                    "type": "string",
                    "minLength": 5,
                    "example": "password"
                }
            }
        },
        "models.UserData": {
            "description": "User full info",
            "type": "object",
            "required": [
                "birthday",
                "country",
                "currency",
                "gender",
                "language",
                "login",
                "name",
                "timezone"
            ],
            "properties": {
                "about": {
                    "description": "Short user self-description",
                    "type": "string",
                    "maxLength": 1000,
                    "example": "Some useful and interesting info"
                },
                "birthday": {
                    "description": "User birthday",
                    "type": "string",
                    "example": "2022-01-01"
                },
                "country": {
                    "description": "User country code",
                    "type": "string",
                    "example": "RU"
                },
                "country_region": {
                    "description": "User country region code",
                    "type": "string",
                    "example": "RU-PER"
                },
                "currency": {
                    "description": "User preferred currency code",
                    "type": "string",
                    "example": "RUB"
                },
                "email": {
                    "description": "User main e-mail",
                    "type": "string",
                    "maxLength": 320,
                    "example": "example@example.com"
                },
                "gender": {
                    "description": "User gender",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Gender"
                        }
                    ],
                    "example": "MALE"
                },
                "language": {
                    "description": "User preferred language code",
                    "type": "string",
                    "example": "ru_RU"
                },
                "login": {
                    "description": "User login",
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 3,
                    "example": "user"
                },
                "name": {
                    "description": "User full name",
                    "type": "string",
                    "maxLength": 1500,
                    "example": "User Usersson"
                },
                "phone": {
                    "description": "User main phone number",
                    "type": "string",
                    "example": "+78005553535"
                },
                "timezone": {
                    "description": "User timezone",
                    "type": "string",
                    "example": "UTC"
                }
            }
        },
        "models.UserItem": {
            "description": "User short info item",
            "type": "object",
            "required": [
                "birthday",
                "gender",
                "login",
                "name"
            ],
            "properties": {
                "birthday": {
                    "description": "User birthday",
                    "type": "string",
                    "example": "2022-01-01"
                },
                "gender": {
                    "description": "User gender",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Gender"
                        }
                    ],
                    "example": "MALE"
                },
                "id": {
                    "description": "User ID",
                    "type": "integer"
                },
                "login": {
                    "description": "User login",
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 3,
                    "example": "user"
                },
                "name": {
                    "description": "User full name",
                    "type": "string",
                    "maxLength": 1500,
                    "example": "User Usersson"
                }
            }
        }
    },
    "tags": [
        {
            "description": "User service",
            "name": "users"
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Social Nyetwork Swagger",
	Description:      "This is a Social Nyetwork server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
