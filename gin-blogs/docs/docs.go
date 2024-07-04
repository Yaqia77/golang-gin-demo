// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
  "swagger": "2.0",
  "info": {
    "description": "This is a sample server for a blog.",
    "version": "1.0.0",
    "title": "Blog API"
  },
  "host": "localhost:8080",
  "basePath": "/api",
  "schemes": [
    "http"
  ],
  "paths": {
    "/register": {
      "post": {
        "summary": "Register a new user",
        "description": "Register a new user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "in": "body",
            "name": "user",
            "description": "User object",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Bad request"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/login": {
      "post": {
        "summary": "Login user",
        "description": "Login user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "in": "body",
            "name": "user",
            "description": "User object",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Bad request"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/blog": {
      "get": {
        "summary": "Get all blogs",
        "description": "Get all blogs",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Blog"
              }
            }
          },
          "404": {
            "description": "Not found"
          }
        }
      },
      "post": {
        "summary": "Create a blog",
        "description": "Create a new blog",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "in": "body",
            "name": "blog",
            "description": "Blog object",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Blog"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Blog"
            }
          },
          "400": {
            "description": "Bad request"
          },
          "409": {
            "description": "Conflict"
          }
        }
      }
    },
    "/blog/{id}": {
      "get": {
        "summary": "Get a blog by ID",
        "description": "Get a blog by ID",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "description": "Blog ID",
            "required": true,
            "type": "integer"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Blog"
            }
          },
          "400": {
            "description": "Bad request"
          },
          "404": {
            "description": "Not found"
          }
        }
      },
      "put": {
        "summary": "Update a blog",
        "description": "Update an existing blog",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "description": "Blog ID",
            "required": true,
            "type": "integer"
          },
          {
            "in": "body",
            "name": "blog",
            "description": "Blog object",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Blog"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Blog"
            }
          },
          "400": {
            "description": "Bad request"
          },
          "404": {
            "description": "Not found"
          }
        }
      },
      "delete": {
        "summary": "Delete a blog",
        "description": "Delete a blog by ID",
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "description": "Blog ID",
            "required": true,
            "type": "integer"
          }
        ],
        "responses": {
          "204": {
            "description": "successful operation"
          },
          "400": {
            "description": "Bad request"
          },
          "404": {
            "description": "Not found"
          }
        }
      }
    }
  },
  "definitions": {
    "User": {
      "type": "object",
      "required": [
        "username",
        "password"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "readOnly": true
        },
        "username": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "Blog": {
      "type": "object",
      "required": [
        "title",
        "content"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "readOnly": true
        },
        "title": {
          "type": "string"
        },
        "content": {
          "type": "string"
        },
        "created_at": {
          "type": "string",
          "format": "date-time"
        },
        "updated_at": {
          "type": "string",
          "format": "date-time"
        }
      }
    }
  }
}
`

var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/",
	Schemes:          []string{"http"},
	Title:            "MyApp API",
	Description:      "This is a sample server for MyApp.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
