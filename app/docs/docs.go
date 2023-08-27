// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/databases": {
            "get": {
                "description": "list database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "db"
                ],
                "summary": "list Database",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.DB"
                            }
                        }
                    }
                }
            }
        },
        "/databases/{dbName}": {
            "get": {
                "description": "db Get",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "db"
                ],
                "summary": "db Get",
                "parameters": [
                    {
                        "type": "string",
                        "description": "database customer name",
                        "name": "dbName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.DB"
                        }
                    }
                }
            },
            "put": {
                "description": "db CreateOrUpdate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "db"
                ],
                "summary": "db CreateOrUpdate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "database customer name",
                        "name": "dbName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "database CreateOrUpdate request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.DBReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "db Remove",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "db"
                ],
                "summary": "db Remove",
                "parameters": [
                    {
                        "type": "string",
                        "description": "database customer name",
                        "name": "dbName",
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
                    }
                }
            }
        },
        "/execute/runSQL": {
            "get": {
                "description": "execute sql cmd",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "execute"
                ],
                "summary": "execute sql cmd",
                "parameters": [
                    {
                        "description": "Run SQL request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.RunSQLReq"
                        }
                    }
                ],
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
        "/tasks": {
            "get": {
                "description": "list cron tasks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "list tasks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.Task"
                            }
                        }
                    }
                }
            }
        },
        "/tasks/{taskName}": {
            "get": {
                "description": "task Get",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "task Get",
                "parameters": [
                    {
                        "type": "string",
                        "description": "task name",
                        "name": "taskName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Task"
                        }
                    }
                }
            },
            "put": {
                "description": "create or update cron task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "create or update task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "cron task name",
                        "name": "taskName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "create or update cron task request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.TaskReq"
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
                    }
                }
            },
            "delete": {
                "description": "remove cron task by task name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "remove task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "cron task name",
                        "name": "taskName",
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
                    }
                }
            }
        }
    },
    "definitions": {
        "types.DB": {
            "type": "object",
            "required": [
                "addr",
                "customerName",
                "databaseName",
                "dbType",
                "password",
                "user"
            ],
            "properties": {
                "addr": {
                    "type": "string"
                },
                "customerName": {
                    "type": "string"
                },
                "databaseName": {
                    "type": "string"
                },
                "dbType": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "types.DBReq": {
            "type": "object",
            "required": [
                "addr",
                "databaseName",
                "dbType",
                "password",
                "user"
            ],
            "properties": {
                "addr": {
                    "type": "string"
                },
                "databaseName": {
                    "type": "string"
                },
                "dbType": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "types.RunSQLReq": {
            "type": "object",
            "required": [
                "SQLStr",
                "dbCustomerName"
            ],
            "properties": {
                "SQLStr": {
                    "type": "string"
                },
                "dbCustomerName": {
                    "type": "string"
                }
            }
        },
        "types.Task": {
            "type": "object",
            "required": [
                "SQLStr",
                "cronStr",
                "dbName"
            ],
            "properties": {
                "SQLStr": {
                    "type": "string"
                },
                "cronStr": {
                    "type": "string"
                },
                "dbName": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "types.TaskReq": {
            "type": "object",
            "required": [
                "SQLStr",
                "cronStr",
                "dbName"
            ],
            "properties": {
                "SQLStr": {
                    "type": "string"
                },
                "cronStr": {
                    "type": "string"
                },
                "dbName": {
                    "type": "string"
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "APP",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
