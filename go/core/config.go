package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "Statuspage",
			"slug": "statuspage",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
			},
		},
		"options": map[string]any{
			"base": "https://api.statuspage.io/v1",
			"auth": map[string]any{
				"prefix": "OAuth",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"component": map[string]any{},
				"component_group_uptime": map[string]any{},
				"group_component": map[string]any{},
				"incident": map[string]any{},
				"incident_postmortem": map[string]any{},
				"incident_subscriber": map[string]any{},
				"incident_template": map[string]any{},
				"incident_update": map[string]any{},
				"metric": map[string]any{},
				"metrics_provider": map[string]any{},
				"page": map[string]any{},
				"page_access_group": map[string]any{},
				"page_access_user": map[string]any{},
				"permission": map[string]any{},
				"postmortem": map[string]any{},
				"status_embed_config": map[string]any{},
				"subscriber": map[string]any{},
				"user": map[string]any{},
			},
		},
		"entity": map[string]any{
			"component": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "automation_email",
						"short": "Requires a special feature flag to be enabled",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "component",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "description",
						"short": "More detailed description for component",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "group",
						"short": "Is this component a group",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "group_id",
						"short": "Component Group identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Incident identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"short": "Display name for component",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "only_show_if_degraded",
						"short": "Requires a special feature flag to be enabled",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "page_id",
						"short": "Page identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "position",
						"short": "Order the component will appear on the page",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "showcase",
						"short": "Should this component be showcased",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "start_date",
						"short": "The date this component started being used",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"short": "Status of component",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$STRING`",
					},
				},
				"name": "component",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/components/{component_id}/page_access_groups",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
									"{id}",
									"page_access_groups",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"component_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "page_access_group",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/components/{component_id}/page_access_users",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
									"{id}",
									"page_access_users",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"component_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "page_access_user",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"component": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_access_group_id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{page_access_group_id}",
									"components",
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_access_group_id",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_access_user_id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{page_access_user_id}",
									"components",
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_access_user_id",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "end",
											"orig": "end",
											"type": "Any",
										},
										map[string]any{
											"kind": "query",
											"name": "start",
											"orig": "start",
											"type": "Any",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/components/{component_id}/uptime",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
									"{id}",
									"uptime",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"component_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "uptime",
									"exist": []any{
										"end",
										"id",
										"page_id",
										"start",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.related_events`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/components/{component_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"component_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/components/{component_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"component_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"component": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/components/{component_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"component_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/components/{component_id}/page_access_groups",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
									"{id}",
									"page_access_groups",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"component_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "page_access_group",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/components/{component_id}/page_access_users",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
									"{id}",
									"page_access_users",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"component_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "page_access_user",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/components/{component_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"component_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"component": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
						[]any{
							"page",
							"page_access_group",
						},
						[]any{
							"page",
							"page_access_user",
						},
					},
				},
			},
			"component_group_uptime": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "component_id",
						"short": "Component identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "incidents",
						"short": "Related incidents",
						"type": "`$OBJECT`",
					},
				},
				"name": "component_group_uptime",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "end",
											"orig": "end",
											"type": "Any",
										},
										map[string]any{
											"kind": "query",
											"name": "start",
											"orig": "start",
											"type": "Any",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/component-groups/{id}/uptime",
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
									"{id}",
									"uptime",
								},
								"select": map[string]any{
									"exist": []any{
										"end",
										"id",
										"page_id",
										"start",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.related_events`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
					},
				},
			},
			"group_component": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "component_group",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "components",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "description",
						"short": "Description of the component group.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Component Group Identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "position",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$STRING`",
					},
				},
				"name": "group_component",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/component-groups",
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/component-groups",
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/component-groups/{id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/component-groups/{id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/component-groups/{id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/component-groups/{id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
					},
				},
			},
			"incident": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "auto_transition_deliver_notifications_at_end",
						"short": "Controls whether send notification when scheduled maintenances auto transition to completed.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "auto_transition_deliver_notifications_at_start",
						"short": "Controls whether send notification when scheduled maintenances auto transition to started.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "auto_transition_to_maintenance_state",
						"short": "Controls whether change components status to under_maintenance once scheduled maintenance is in progress.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "auto_transition_to_operational_state",
						"short": "Controls whether change components status to operational once scheduled maintenance completes.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "components",
						"short": "Incident components",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "created_at",
						"short": "The timestamp when the incident was created at.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Incident Identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "impact",
						"short": "The impact of the incident.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "impact_override",
						"short": "value to override calculated impact value",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "incident",
						"op": map[string]any{
							"patch": map[string]any{
								"type": "`$OBJECT`",
							},
							"update": map[string]any{
								"type": "`$OBJECT`",
							},
						},
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "incident_updates",
						"short": "The incident updates for incident.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "metadata",
						"short": "Metadata attached to the incident.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "monitoring_at",
						"short": "The timestamp when incident entered monitoring state.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"short": "Incident Name.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page_id",
						"short": "Incident Page Identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "postmortem_body",
						"short": "Body of the Postmortem.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "postmortem_body_last_updated_at",
						"short": "The timestamp when the incident postmortem body was last updated at.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "postmortem_ignored",
						"short": "Controls whether the incident will have postmortem.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "postmortem_notified_subscribers",
						"short": "Indicates whether subscribers are already notificed about postmortem.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "postmortem_notified_twitter",
						"short": "Controls whether to decide if notify postmortem on twitter.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "postmortem_published_at",
						"short": "The timestamp when the postmortem was published.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "reminder_intervals",
						"short": "Custom reminder intervals for unresolved/open incidents.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "resolved_at",
						"short": "The timestamp when incident was resolved.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "scheduled_auto_completed",
						"short": "Controls whether the incident is scheduled to automatically change to complete.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "scheduled_auto_in_progress",
						"short": "Controls whether the incident is scheduled to automatically change to in progress.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "scheduled_for",
						"short": "The timestamp the incident is scheduled for.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "scheduled_remind_prior",
						"short": "Controls whether to remind subscribers prior to scheduled incidents.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "scheduled_reminded_at",
						"short": "The timestamp when the scheduled incident reminder was sent at.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "scheduled_until",
						"short": "The timestamp the incident is scheduled until.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shortlink",
						"short": "Incident Shortlink.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"short": "The incident status.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "The timestamp when the incident was updated at.",
						"type": "`$STRING`",
					},
				},
				"name": "incident",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/incidents",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"incident": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
								},
								"select": map[string]any{
									"exist": []any{
										"limit",
										"page",
										"page_id",
										"q",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 100,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/active_maintenance",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"active_maintenance",
								},
								"select": map[string]any{
									"$action": "active_maintenance",
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 100,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/scheduled",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"scheduled",
								},
								"select": map[string]any{
									"$action": "scheduled",
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 100,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/unresolved",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"unresolved",
								},
								"select": map[string]any{
									"$action": "unresolved",
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 100,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/upcoming",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"upcoming",
								},
								"select": map[string]any{
									"$action": "upcoming",
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/{incident_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"incident_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/incidents/{incident_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"incident_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"incident": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/incidents/{incident_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"incident_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/incidents/{incident_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"incident_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"incident": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
					},
				},
			},
			"incident_postmortem": map[string]any{
				"fields": []any{},
				"name": "incident_postmortem",
				"op": map[string]any{
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/incidents/{incident_id}/postmortem",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{id}",
									"postmortem",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"incident_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
					},
				},
			},
			"incident_subscriber": map[string]any{
				"fields": []any{},
				"name": "incident_subscriber",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "subscriber_id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}/resend_confirmation",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"subscribers",
									"{subscriber_id}",
									"resend_confirmation",
								},
								"select": map[string]any{
									"exist": []any{
										"incident_id",
										"page_id",
										"subscriber_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
							"incident",
							"subscriber",
						},
					},
				},
			},
			"incident_template": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "body",
						"short": "Body of the incident or maintenance update to be applied when selecting this template",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "components",
						"short": "Affected components",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "group_id",
						"short": "Identifier of Template Group this template belongs to",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Incident Template Identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"short": "Name of the template, as shown in the list on the \"Templates\" tab of the \"Incidents\" page",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "should_send_notifications",
						"short": "Whether the \"deliver notifications\" checkbox should be selected when selecting this template",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "should_tweet",
						"short": "Whether the \"tweet update\" checkbox should be selected when selecting this template",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "template",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "title",
						"short": "Title to be applied to the incident or maintenance when selecting this template",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "update_status",
						"short": "The status the incident or maintenance should transition to when selecting this template",
						"type": "`$STRING`",
					},
				},
				"name": "incident_template",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/incident_templates",
								"parts": []any{
									"pages",
									"{page_id}",
									"incident_templates",
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 100,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incident_templates",
								"parts": []any{
									"pages",
									"{page_id}",
									"incident_templates",
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
					},
				},
			},
			"incident_update": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "affected_components",
						"short": "Affected components associated with the incident update.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "body",
						"short": "Incident update body.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "created_at",
						"short": "The timestamp when the incident update was created at.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "custom_tweet",
						"short": "An optional customized tweet message for incident postmortem.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "deliver_notifications",
						"short": "Controls whether to delivery notifications.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "display_at",
						"short": "Timestamp when incident update is happened.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Incident Update Identifier.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "incident_id",
						"short": "Incident Identifier.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "incident_update",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "status",
						"short": "The incident status.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tweet_id",
						"short": "Tweet identifier associated to this incident update.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "twitter_updated_at",
						"short": "The timestamp when twitter updated at.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "The timestamp when the incident update is updated.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "wants_twitter_update",
						"short": "Controls whether to create twitter update.",
						"type": "`$BOOLEAN`",
					},
				},
				"name": "incident_update",
				"op": map[string]any{
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "incident_update_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"incident_updates",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"incident_update_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"incident_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"incident_update": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "incident_update_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"incident_updates",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"incident_update_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"incident_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"incident_update": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
							"incident",
						},
					},
				},
			},
			"metric": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "backfill_percentage",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "backfilled",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "data",
						"req": true,
						"short": "Add data points to metrics",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "decimal_places",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "display",
						"short": "Should the metric be displayed",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "id",
						"short": "Metric identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "last_fetched_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "metric",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "metric_identifier",
						"short": "Metric Display identifier used to look up the metric data from the provider",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "metrics_provider_id",
						"short": "Metric Provider identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "most_recent_data_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"short": "Name of metric",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "reference_name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "suffix",
						"short": "Suffix to describe the units on the graph",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tooltip_description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "y_axis_hidden",
						"short": "Should the values on the y axis be hidden on render",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "y_axis_max",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "y_axis_min",
						"type": "`$NUMBER`",
					},
				},
				"name": "metric",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/metrics/{metric_id}/data",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics",
									"{id}",
									"data",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"metric_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "data",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "metrics_provider_id",
											"orig": "metrics_provider_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics_providers",
									"{metrics_provider_id}",
									"metrics",
								},
								"select": map[string]any{
									"exist": []any{
										"metrics_provider_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"metric": "`reqdata`",
									},
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/metrics/data",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics",
									"data",
								},
								"select": map[string]any{
									"$action": "data",
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_access_user_id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{page_access_user_id}",
									"metrics",
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_access_user_id",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "metrics_provider_id",
											"orig": "metrics_provider_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics_providers",
									"{metrics_provider_id}",
									"metrics",
								},
								"select": map[string]any{
									"exist": []any{
										"metrics_provider_id",
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/metrics",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics",
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/metrics/{metric_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"metric_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/metrics/{metric_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"metric_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"metric": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/metrics/{metric_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"metric_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/metrics/{metric_id}/data",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics",
									"{id}",
									"data",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"metric_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "data",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/metrics/{metric_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"metric_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"metric": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
						[]any{
							"page",
							"metrics_provider",
						},
						[]any{
							"page",
							"page_access_user",
						},
					},
				},
			},
			"metrics_provider": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "disabled",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "id",
						"short": "Identifier for Metrics Provider",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "last_revalidated_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "metric_base_uri",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "metrics_provider",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "page_id",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$STRING`",
					},
				},
				"name": "metrics_provider",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/metrics_providers",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics_providers",
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"metrics_provider": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/metrics_providers",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics_providers",
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "metrics_provider_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics_providers",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"metrics_provider_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "metrics_provider_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics_providers",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"metrics_provider_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"metrics_provider": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "metrics_provider_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics_providers",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"metrics_provider_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "metrics_provider_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics_providers",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"metrics_provider_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"metrics_provider": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
					},
				},
			},
			"page": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "activity_score",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "allow_email_subscribers",
						"short": "Can your users choose to receive notifications via email",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "allow_incident_subscribers",
						"short": "Can your users subscribe to notifications for a single incident",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "allow_page_subscribers",
						"short": "Can your users subscribe to all notifications on the page",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "allow_rss_atom_feeds",
						"short": "Can your users choose to access incident feeds via RSS/Atom (not functional on Audience-Specific pages)",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "allow_sms_subscribers",
						"short": "Can your users choose to receive notifications via SMS",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "allow_webhook_subscribers",
						"short": "Can your users choose to receive notifications via Webhooks",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "branding",
						"short": "The main template your statuspage will use",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "city",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "country",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "created_at",
						"short": "Timestamp the record was created",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_blues",
						"short": "CSS Color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_body_background_color",
						"short": "CSS Color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_border_color",
						"short": "CSS Color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_font_color",
						"short": "CSS Color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_graph_color",
						"short": "CSS Color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_greens",
						"short": "CSS Color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_light_font_color",
						"short": "CSS Color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_link_color",
						"short": "CSS Color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_no_data",
						"short": "CSS Color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_oranges",
						"short": "CSS Color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_reds",
						"short": "CSS Color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_yellows",
						"short": "CSS Color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "domain",
						"short": "CNAME alias for your status page",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "email_logo",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "favicon_logo",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "headline",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "hero_cover",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "hidden_from_search",
						"short": "Should your page hide itself from search engines",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "id",
						"short": "Page identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ip_restrictions",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"short": "Name of your page to be displayed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "notifications_email_footer",
						"short": "Allows you to customize the footer appearing on your notification emails.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "notifications_from_email",
						"short": "Allows you to customize the email address your page notifications come from",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "page_description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "state",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "subdomain",
						"short": "Subdomain at which to access your status page",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "support_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "time_zone",
						"short": "Timezone configured for your page",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "transactional_logo",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "twitter_logo",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "twitter_username",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"short": "Timestamp the record was last updated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "url",
						"short": "Website of your page.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "viewers_must_be_team_members",
						"type": "`$BOOLEAN`",
					},
				},
				"name": "page",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/pages",
								"parts": []any{
									"pages",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}",
								"parts": []any{
									"pages",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}",
								"parts": []any{
									"pages",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"page": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}",
								"parts": []any{
									"pages",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"page": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"page_access_group": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "component_ids",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$ARRAY`",
							},
						},
						"short": "List of components codes to set on the page access group",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "external_identifier",
						"short": "Associates group with external group.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Page Access Group Identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "metric_ids",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "name",
						"short": "Name for this Group.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page_access_group",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "page_access_user_ids",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "page_id",
						"short": "Page Identifier.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$STRING`",
					},
				},
				"name": "page_access_group",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{id}",
									"components",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_group_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/page_access_groups",
								"parts": []any{
									"pages",
									"{id}",
									"page_access_groups",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"page_access_group": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_groups",
								"parts": []any{
									"pages",
									"{id}",
									"page_access_groups",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_group_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_group_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"page_access_group": "`reqdata`",
									},
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{id}",
									"components",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_group_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "component_id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components/{component_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{id}",
									"components",
									"{component_id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_group_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"component_id",
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_group_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{id}",
									"components",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_group_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_group_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"page_access_group": "`reqdata`",
									},
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{id}",
									"components",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_group_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
						[]any{
							"page",
							"component",
						},
					},
				},
			},
			"page_access_user": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "component_ids",
						"req": true,
						"short": "List of component codes to allow access to",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "email",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "external_login",
						"short": "IDP login user id.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Page Access User Identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "metric_ids",
						"req": true,
						"short": "List of metrics to add",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "page_access_group_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page_access_group_ids",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page_access_user",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "page_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$STRING`",
					},
				},
				"name": "page_access_user",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
									"components",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
									"metrics",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "metric",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/page_access_users",
								"parts": []any{
									"pages",
									"{id}",
									"page_access_users",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"page_access_user": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "email",
											"orig": "email",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_users",
								"parts": []any{
									"pages",
									"{id}",
									"page_access_users",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"email",
										"id",
										"page",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
									"components",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
									"metrics",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "metric",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "component_id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components/{component_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
									"components",
									"{component_id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"component_id",
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "metric_id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics/{metric_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
									"metrics",
									"{metric_id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"metric_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
									"components",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
									"metrics",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "metric",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
									"components",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
									"metrics",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "metric",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
						[]any{
							"page",
							"component",
						},
						[]any{
							"page",
							"metric",
						},
					},
				},
			},
			"permission": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "pages",
						"short": "Pages accessible by the user.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "user_id",
						"short": "User identifier",
						"type": "`$STRING`",
					},
				},
				"name": "permission",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "organization_id",
											"orig": "organization_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/organizations/{organization_id}/permissions/{user_id}",
								"parts": []any{
									"organizations",
									"{organization_id}",
									"permissions",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"user_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"organization_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "organization_id",
											"orig": "organization_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/organizations/{organization_id}/permissions/{user_id}",
								"parts": []any{
									"organizations",
									"{organization_id}",
									"permissions",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"user_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"organization_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"organization",
						},
					},
				},
			},
			"postmortem": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "body",
						"short": "Postmortem body",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "body_draft",
						"short": "Body draft",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "body_draft_updated_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "body_updated_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "custom_tweet",
						"short": "Custom tweet for Incident Postmortem",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "notify_subscribers",
						"short": "Should email subscribers be notified.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "notify_twitter",
						"short": "Should Twitter followers be notified.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "postmortem",
						"op": map[string]any{
							"update": map[string]any{
								"type": "`$OBJECT`",
							},
						},
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "preview_key",
						"short": "Preview Key",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "published_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$STRING`",
					},
				},
				"name": "postmortem",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/{incident_id}/postmortem",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"postmortem",
								},
								"select": map[string]any{
									"exist": []any{
										"incident_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/incidents/{incident_id}/postmortem",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"postmortem",
								},
								"select": map[string]any{
									"exist": []any{
										"incident_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"postmortem": "`reqdata`",
									},
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/incidents/{incident_id}/postmortem/publish",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"postmortem",
									"publish",
								},
								"select": map[string]any{
									"$action": "publish",
									"exist": []any{
										"incident_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"postmortem": "`reqdata`",
									},
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/incidents/{incident_id}/postmortem/revert",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"postmortem",
									"revert",
								},
								"select": map[string]any{
									"$action": "revert",
									"exist": []any{
										"incident_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
							"incident",
						},
					},
				},
			},
			"status_embed_config": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "incident_background_color",
						"short": "Color of status embed iframe background when displaying incident",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "incident_text_color",
						"short": "Color of status embed iframe text when displaying incident",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "maintenance_background_color",
						"short": "Color of status embed iframe background when displaying maintenance",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "maintenance_text_color",
						"short": "Color of status embed iframe text when displaying maintenance",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page_id",
						"short": "Page identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "position",
						"short": "Corner where status embed iframe will appear on page",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status_embed_config",
						"type": "`$OBJECT`",
					},
				},
				"name": "status_embed_config",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/status_embed_config",
								"parts": []any{
									"pages",
									"{page_id}",
									"status_embed_config",
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/status_embed_config",
								"parts": []any{
									"pages",
									"{page_id}",
									"status_embed_config",
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"status_embed_config": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/status_embed_config",
								"parts": []any{
									"pages",
									"{page_id}",
									"status_embed_config",
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"status_embed_config": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
					},
				},
			},
			"subscriber": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "component_ids",
						"short": "A list of component ids for which the subscriber should recieve updates for.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "components",
						"short": "The components for which the subscriber has elected to receive updates.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "display_phone_number",
						"short": "A formatted version of the phone_number and phone_country pair, nicely formatted for display.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "email",
						"short": "The email address to use to contact the subscriber.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "endpoint",
						"short": "The URL where a webhook subscriber elects to receive updates.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Subscriber Identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "integration_partner",
						"short": "The number of integration partners found by the query.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "mode",
						"short": "The communication mode of the subscriber.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "obfuscated_channel_name",
						"short": "Obfuscated slack channel name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page_access_user_id",
						"short": "The Page Access user this subscriber belongs to (only for audience-specific pages).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "phone_country",
						"short": "The two-character country code representing the country of which the phone_number is a part.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "phone_number",
						"short": "The phone number used to contact an SMS subscriber",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "purge_at",
						"short": "The timestamp when a quarantined subscriber will be purged (unsubscribed).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "quarantined_at",
						"short": "The timestamp when the subscriber was quarantined due to an issue reaching them.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "skip_confirmation_notification",
						"short": "If this is true, do not notify the user with changes to their subscription.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "skip_unsubscription_notification",
						"short": "If skip_unsubscription_notification is true, the subscribers do not receive any notifications when they are unsubscribed.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "slack",
						"short": "The number of Slack subscribers found by the query.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "sms",
						"short": "The number of Webhook subscribers found by the query.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "state",
						"short": "If this is present, only unsubscribe subscribers in this state.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "subscriber",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "subscribers",
						"req": true,
						"short": "The array of quarantined subscriber codes to reactivate, or \"all\" to reactivate all quarantined subscribers.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "teams",
						"short": "The number of MS teams subscribers found by the query.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "type",
						"short": "If this is present, only reactivate subscribers of this type.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "webhook",
						"short": "The number of SMS subscribers found by the query.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "workspace_name",
						"short": "The workspace name of the slack subscriber.",
						"type": "`$STRING`",
					},
				},
				"name": "subscriber",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/subscribers/{subscriber_id}/resend_confirmation",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"{id}",
									"resend_confirmation",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"subscriber_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "resend_confirmation",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/incidents/{incident_id}/subscribers",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"subscribers",
								},
								"select": map[string]any{
									"exist": []any{
										"incident_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"subscriber": "`reqdata`",
									},
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/subscribers",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"subscriber": "`reqdata`",
									},
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/subscribers/reactivate",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"reactivate",
								},
								"select": map[string]any{
									"$action": "reactivate",
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/subscribers/resend_confirmation",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"resend_confirmation",
								},
								"select": map[string]any{
									"$action": "resend_confirmation",
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/subscribers/unsubscribe",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"unsubscribe",
								},
								"select": map[string]any{
									"$action": "unsubscribe",
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "asc",
											"kind": "query",
											"name": "sort_direction",
											"orig": "sort_direction",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "primary",
											"kind": "query",
											"name": "sort_field",
											"orig": "sort_field",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "active",
											"kind": "query",
											"name": "state",
											"orig": "state",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "type",
											"orig": "type",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/subscribers",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
								},
								"select": map[string]any{
									"exist": []any{
										"limit",
										"page",
										"page_id",
										"q",
										"sort_direction",
										"sort_field",
										"state",
										"type",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/{incident_id}/subscribers",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"subscribers",
								},
								"select": map[string]any{
									"exist": []any{
										"incident_id",
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/subscribers/unsubscribed",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"unsubscribed",
								},
								"select": map[string]any{
									"$action": "unsubscribed",
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"subscribers",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"subscriber_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"incident_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": "active",
											"kind": "query",
											"name": "state",
											"orig": "state",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "type",
											"orig": "type",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/subscribers/count",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"count",
								},
								"select": map[string]any{
									"$action": "count",
									"exist": []any{
										"page_id",
										"state",
										"type",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/subscribers/{subscriber_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"subscriber_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/subscribers/histogram_by_state",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"histogram_by_state",
								},
								"select": map[string]any{
									"$action": "histogram_by_state",
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"subscribers",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"subscriber_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"incident_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "skip_unsubscription_notification",
											"orig": "skip_unsubscription_notification",
											"type": "`$BOOLEAN`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/subscribers/{subscriber_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"subscriber_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
										"skip_unsubscription_notification",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/subscribers/{subscriber_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"subscriber_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
						[]any{
							"page",
							"incident",
						},
					},
				},
			},
			"user": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "email",
						"short": "Email address for the team member",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "first_name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "User identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "last_name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "organization_id",
						"short": "Organization identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "user",
						"req": true,
						"type": "`$OBJECT`",
					},
				},
				"name": "user",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "organization_id",
											"orig": "organization_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/organizations/{organization_id}/users",
								"parts": []any{
									"organizations",
									"{organization_id}",
									"users",
								},
								"select": map[string]any{
									"exist": []any{
										"organization_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"user": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "organization_id",
											"orig": "organization_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/organizations/{organization_id}/users",
								"parts": []any{
									"organizations",
									"{organization_id}",
									"users",
								},
								"select": map[string]any{
									"exist": []any{
										"organization_id",
										"page",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "organization_id",
											"orig": "organization_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/organizations/{organization_id}/users/{user_id}",
								"parts": []any{
									"organizations",
									"{organization_id}",
									"users",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"user_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"organization_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"organization",
						},
					},
				},
			},
		},
	}
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
