// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-12-07 12:01:31.151973445 +0700 +07 m=+0.041302554

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/GetAllAuctionHistory": {
            "get": {
                "description": "Get tất cả dữ liệu ở bảng AuctionHistoryAll trả về  Json",
                "consumes": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AuctionHistory"
                        }
                    }
                }
            }
        },
        "/GetAllCategories": {
            "get": {
                "description": "Get tất cả dữ liệu ở bảng Categories trả về  Json",
                "consumes": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Categorie"
                        }
                    }
                }
            }
        },
        "/GetAllProduct": {
            "get": {
                "description": "Get tất cả dữ liệu ở bảng GetProductAll trả về  Json",
                "consumes": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Product"
                        }
                    }
                }
            }
        },
        "/GetAllUser": {
            "get": {
                "description": "Get tất cả dữ liệu ở bảng GetUserAll trả về  Json",
                "consumes": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AuctionHistory": {
            "type": "object",
            "properties": {
                "biddingPrice": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "productId": {
                    "type": "integer"
                },
                "timeToBid": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "model.Categorie": {
            "type": "object",
            "properties": {
                "NameCategories": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.Product": {
            "type": "object",
            "required": [
                "CategoriesID",
                "Description",
                "EndTime",
                "PriceNow",
                "PriceStart",
                "ProductImage",
                "ProductName",
                "StartTime",
                "Status"
            ],
            "properties": {
                "CategoriesID": {
                    "type": "integer"
                },
                "Description": {
                    "type": "string"
                },
                "EndTime": {
                    "type": "string"
                },
                "PriceNow": {
                    "type": "integer"
                },
                "PriceStart": {
                    "type": "integer"
                },
                "ProductImage": {
                    "type": "string"
                },
                "ProductName": {
                    "type": "string"
                },
                "StartTime": {
                    "type": "string"
                },
                "Status": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.User": {
            "type": "object",
            "required": [
                "Email",
                "FisrtName",
                "Id",
                "LastName",
                "Password",
                "Phone",
                "Role",
                "Wallet"
            ],
            "properties": {
                "Email": {
                    "type": "string"
                },
                "FisrtName": {
                    "type": "string"
                },
                "Id": {
                    "type": "integer"
                },
                "LastName": {
                    "type": "string"
                },
                "Password": {
                    "type": "string"
                },
                "Phone": {
                    "type": "string"
                },
                "Role": {
                    "type": "boolean"
                },
                "Wallet": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "petstore.swagger.io",
	BasePath:    "/v2",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample server Petstore server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
