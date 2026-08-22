
import { BaseFeature } from './feature/base/BaseFeature'
import { TestFeature } from './feature/test/TestFeature'



const FEATURE_CLASS: Record<string, typeof BaseFeature> = {
   test: TestFeature,

}


class Config {

  makeFeature(this: any, fn: string) {
    const fc = FEATURE_CLASS[fn]
    const fi = new fc()
    // TODO: errors etc
    return fi
  }

  // False for a feature added at runtime via options.extend (station's
  // adopt path) - the constructor uses this to skip makeFeature for names
  // no generated class backs.
  hasFeature(this: any, fn: string) {
    return null != FEATURE_CLASS[fn]
  }


  main = {
    name: 'Statuspage',
        slug: "statuspage",
    version: "0.1.1",
    target: "ts",

  }


  feature = {
     test:     {
      "options": {
        "active": false
      }
    },

  }


  options = {
    base: "https://api.statuspage.io/v1",

    auth: {
      prefix: 'OAuth',
    },

    headers: {
      "content-type": "application/json"
    },

    entity: {
      
      component: {
      },

      component_group_uptime: {
      },

      group_component: {
      },

      incident: {
      },

      incident_postmortem: {
      },

      incident_subscriber: {
      },

      incident_template: {
      },

      incident_update: {
      },

      metric: {
      },

      metrics_provider: {
      },

      page: {
      },

      page_access_group: {
      },

      page_access_user: {
      },

      permission: {
      },

      postmortem: {
      },

      status_embed_config: {
      },

      subscriber: {
      },

      user: {
      },

    }
  }


  entity = {
    "component": {
      "fields": [
        {
          "name": "automation_email",
          "short": "Requires a special feature flag to be enabled",
          "type": "`$STRING`"
        },
        {
          "name": "component",
          "type": "`$OBJECT`"
        },
        {
          "name": "created_at",
          "type": "`$STRING`"
        },
        {
          "name": "description",
          "short": "More detailed description for component",
          "type": "`$STRING`"
        },
        {
          "name": "group",
          "short": "Is this component a group",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "group_id",
          "short": "Component Group identifier",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "short": "Incident identifier",
          "type": "`$STRING`"
        },
        {
          "name": "name",
          "short": "Display name for component",
          "type": "`$STRING`"
        },
        {
          "name": "only_show_if_degraded",
          "short": "Requires a special feature flag to be enabled",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "page_id",
          "short": "Page identifier",
          "type": "`$STRING`"
        },
        {
          "name": "position",
          "short": "Order the component will appear on the page",
          "type": "`$INTEGER`"
        },
        {
          "name": "showcase",
          "short": "Should this component be showcased",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "start_date",
          "short": "The date this component started being used",
          "type": "`$STRING`"
        },
        {
          "name": "status",
          "short": "Status of component",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$STRING`"
        }
      ],
      "name": "component",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "component_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/components/{component_id}/page_access_groups",
              "parts": [
                "pages",
                "{page_id}",
                "components",
                "{id}",
                "page_access_groups"
              ],
              "rename": {
                "param": {
                  "component_id": "id"
                }
              },
              "select": {
                "$action": "page_access_group",
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "component_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/components/{component_id}/page_access_users",
              "parts": [
                "pages",
                "{page_id}",
                "components",
                "{id}",
                "page_access_users"
              ],
              "rename": {
                "param": {
                  "component_id": "id"
                }
              },
              "select": {
                "$action": "page_access_user",
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/components",
              "parts": [
                "pages",
                "{page_id}",
                "components"
              ],
              "select": {
                "exist": [
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "component": "`reqdata`"
                },
                "res": "`body`"
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_access_group_id",
                    "orig": "page_access_group_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "per_page",
                    "orig": "per_page",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_groups",
                "{page_access_group_id}",
                "components"
              ],
              "select": {
                "exist": [
                  "page",
                  "page_access_group_id",
                  "page_id",
                  "per_page"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_access_user_id",
                    "orig": "page_access_user_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "per_page",
                    "orig": "per_page",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_users",
                "{page_access_user_id}",
                "components"
              ],
              "select": {
                "exist": [
                  "page",
                  "page_access_user_id",
                  "page_id",
                  "per_page"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "per_page",
                    "orig": "per_page",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/components",
              "parts": [
                "pages",
                "{page_id}",
                "components"
              ],
              "select": {
                "exist": [
                  "page",
                  "page_id",
                  "per_page"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "component_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "end",
                    "orig": "end",
                    "type": "Any"
                  },
                  {
                    "kind": "query",
                    "name": "start",
                    "orig": "start",
                    "type": "Any"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/components/{component_id}/uptime",
              "parts": [
                "pages",
                "{page_id}",
                "components",
                "{id}",
                "uptime"
              ],
              "rename": {
                "param": {
                  "component_id": "id"
                }
              },
              "select": {
                "$action": "uptime",
                "exist": [
                  "end",
                  "id",
                  "page_id",
                  "start"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.related_events`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "component_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/components/{component_id}",
              "parts": [
                "pages",
                "{page_id}",
                "components",
                "{id}"
              ],
              "rename": {
                "param": {
                  "component_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "patch": {
          "input": "data",
          "name": "patch",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "component_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PATCH",
              "orig": "/pages/{page_id}/components/{component_id}",
              "parts": [
                "pages",
                "{page_id}",
                "components",
                "{id}"
              ],
              "rename": {
                "param": {
                  "component_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "component": "`reqdata`"
                },
                "res": "`body`"
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "component_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "DELETE",
              "orig": "/pages/{page_id}/components/{component_id}",
              "parts": [
                "pages",
                "{page_id}",
                "components",
                "{id}"
              ],
              "rename": {
                "param": {
                  "component_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "component_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "DELETE",
              "orig": "/pages/{page_id}/components/{component_id}/page_access_groups",
              "parts": [
                "pages",
                "{page_id}",
                "components",
                "{id}",
                "page_access_groups"
              ],
              "rename": {
                "param": {
                  "component_id": "id"
                }
              },
              "select": {
                "$action": "page_access_group",
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "component_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "DELETE",
              "orig": "/pages/{page_id}/components/{component_id}/page_access_users",
              "parts": [
                "pages",
                "{page_id}",
                "components",
                "{id}",
                "page_access_users"
              ],
              "rename": {
                "param": {
                  "component_id": "id"
                }
              },
              "select": {
                "$action": "page_access_user",
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "component_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PUT",
              "orig": "/pages/{page_id}/components/{component_id}",
              "parts": [
                "pages",
                "{page_id}",
                "components",
                "{id}"
              ],
              "rename": {
                "param": {
                  "component_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "component": "`reqdata`"
                },
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "page"
          ],
          [
            "page",
            "page_access_group"
          ],
          [
            "page",
            "page_access_user"
          ]
        ]
      }
    },
    "component_group_uptime": {
      "fields": [
        {
          "name": "component_id",
          "short": "Component identifier",
          "type": "`$STRING`"
        },
        {
          "name": "incidents",
          "short": "Related incidents",
          "type": "`$OBJECT`"
        }
      ],
      "name": "component_group_uptime",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "end",
                    "orig": "end",
                    "type": "Any"
                  },
                  {
                    "kind": "query",
                    "name": "start",
                    "orig": "start",
                    "type": "Any"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/component-groups/{id}/uptime",
              "parts": [
                "pages",
                "{page_id}",
                "component-groups",
                "{id}",
                "uptime"
              ],
              "select": {
                "exist": [
                  "end",
                  "id",
                  "page_id",
                  "start"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.related_events`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "page"
          ]
        ]
      }
    },
    "group_component": {
      "fields": [
        {
          "name": "component_group",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "components",
          "type": "`$STRING`"
        },
        {
          "name": "created_at",
          "type": "`$STRING`"
        },
        {
          "name": "description",
          "short": "Description of the component group.",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "short": "Component Group Identifier",
          "type": "`$STRING`"
        },
        {
          "name": "name",
          "type": "`$STRING`"
        },
        {
          "name": "page_id",
          "type": "`$STRING`"
        },
        {
          "name": "position",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$STRING`"
        }
      ],
      "name": "group_component",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/component-groups",
              "parts": [
                "pages",
                "{page_id}",
                "component-groups"
              ],
              "select": {
                "exist": [
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "per_page",
                    "orig": "per_page",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/component-groups",
              "parts": [
                "pages",
                "{page_id}",
                "component-groups"
              ],
              "select": {
                "exist": [
                  "page",
                  "page_id",
                  "per_page"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/component-groups/{id}",
              "parts": [
                "pages",
                "{page_id}",
                "component-groups",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "patch": {
          "input": "data",
          "name": "patch",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PATCH",
              "orig": "/pages/{page_id}/component-groups/{id}",
              "parts": [
                "pages",
                "{page_id}",
                "component-groups",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "DELETE",
              "orig": "/pages/{page_id}/component-groups/{id}",
              "parts": [
                "pages",
                "{page_id}",
                "component-groups",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PUT",
              "orig": "/pages/{page_id}/component-groups/{id}",
              "parts": [
                "pages",
                "{page_id}",
                "component-groups",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "page"
          ]
        ]
      }
    },
    "incident": {
      "fields": [
        {
          "name": "auto_transition_deliver_notifications_at_end",
          "short": "Controls whether send notification when scheduled maintenances auto transition to completed.",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "auto_transition_deliver_notifications_at_start",
          "short": "Controls whether send notification when scheduled maintenances auto transition to started.",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "auto_transition_to_maintenance_state",
          "short": "Controls whether change components status to under_maintenance once scheduled maintenance is in progress.",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "auto_transition_to_operational_state",
          "short": "Controls whether change components status to operational once scheduled maintenance completes.",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "components",
          "short": "Incident components",
          "type": "`$ARRAY`"
        },
        {
          "name": "created_at",
          "short": "The timestamp when the incident was created at.",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "short": "Incident Identifier",
          "type": "`$STRING`"
        },
        {
          "name": "impact",
          "short": "The impact of the incident.",
          "type": "`$STRING`"
        },
        {
          "name": "impact_override",
          "short": "value to override calculated impact value",
          "type": "`$STRING`"
        },
        {
          "name": "incident",
          "op": {
            "patch": {
              "type": "`$OBJECT`"
            },
            "update": {
              "type": "`$OBJECT`"
            }
          },
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "incident_updates",
          "short": "The incident updates for incident.",
          "type": "`$ARRAY`"
        },
        {
          "name": "metadata",
          "short": "Metadata attached to the incident.",
          "type": "`$OBJECT`"
        },
        {
          "name": "monitoring_at",
          "short": "The timestamp when incident entered monitoring state.",
          "type": "`$STRING`"
        },
        {
          "name": "name",
          "short": "Incident Name.",
          "type": "`$STRING`"
        },
        {
          "name": "page_id",
          "short": "Incident Page Identifier",
          "type": "`$STRING`"
        },
        {
          "name": "postmortem_body",
          "short": "Body of the Postmortem.",
          "type": "`$STRING`"
        },
        {
          "name": "postmortem_body_last_updated_at",
          "short": "The timestamp when the incident postmortem body was last updated at.",
          "type": "`$STRING`"
        },
        {
          "name": "postmortem_ignored",
          "short": "Controls whether the incident will have postmortem.",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "postmortem_notified_subscribers",
          "short": "Indicates whether subscribers are already notificed about postmortem.",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "postmortem_notified_twitter",
          "short": "Controls whether to decide if notify postmortem on twitter.",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "postmortem_published_at",
          "short": "The timestamp when the postmortem was published.",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "reminder_intervals",
          "short": "Custom reminder intervals for unresolved/open incidents.",
          "type": "`$STRING`"
        },
        {
          "name": "resolved_at",
          "short": "The timestamp when incident was resolved.",
          "type": "`$STRING`"
        },
        {
          "name": "scheduled_auto_completed",
          "short": "Controls whether the incident is scheduled to automatically change to complete.",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "scheduled_auto_in_progress",
          "short": "Controls whether the incident is scheduled to automatically change to in progress.",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "scheduled_for",
          "short": "The timestamp the incident is scheduled for.",
          "type": "`$STRING`"
        },
        {
          "name": "scheduled_remind_prior",
          "short": "Controls whether to remind subscribers prior to scheduled incidents.",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "scheduled_reminded_at",
          "short": "The timestamp when the scheduled incident reminder was sent at.",
          "type": "`$STRING`"
        },
        {
          "name": "scheduled_until",
          "short": "The timestamp the incident is scheduled until.",
          "type": "`$STRING`"
        },
        {
          "name": "shortlink",
          "short": "Incident Shortlink.",
          "type": "`$STRING`"
        },
        {
          "name": "status",
          "short": "The incident status.",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "short": "The timestamp when the incident was updated at.",
          "type": "`$STRING`"
        }
      ],
      "name": "incident",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/incidents",
              "parts": [
                "pages",
                "{page_id}",
                "incidents"
              ],
              "select": {
                "exist": [
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "incident": "`reqdata`"
                },
                "res": "`body`"
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "limit",
                    "orig": "limit",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/incidents",
              "parts": [
                "pages",
                "{page_id}",
                "incidents"
              ],
              "select": {
                "exist": [
                  "limit",
                  "page",
                  "page_id",
                  "q"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 100,
                    "kind": "query",
                    "name": "per_page",
                    "orig": "per_page",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/incidents/active_maintenance",
              "parts": [
                "pages",
                "{page_id}",
                "incidents",
                "active_maintenance"
              ],
              "select": {
                "$action": "active_maintenance",
                "exist": [
                  "page",
                  "page_id",
                  "per_page"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 100,
                    "kind": "query",
                    "name": "per_page",
                    "orig": "per_page",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/incidents/scheduled",
              "parts": [
                "pages",
                "{page_id}",
                "incidents",
                "scheduled"
              ],
              "select": {
                "$action": "scheduled",
                "exist": [
                  "page",
                  "page_id",
                  "per_page"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 100,
                    "kind": "query",
                    "name": "per_page",
                    "orig": "per_page",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/incidents/unresolved",
              "parts": [
                "pages",
                "{page_id}",
                "incidents",
                "unresolved"
              ],
              "select": {
                "$action": "unresolved",
                "exist": [
                  "page",
                  "page_id",
                  "per_page"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 100,
                    "kind": "query",
                    "name": "per_page",
                    "orig": "per_page",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/incidents/upcoming",
              "parts": [
                "pages",
                "{page_id}",
                "incidents",
                "upcoming"
              ],
              "select": {
                "$action": "upcoming",
                "exist": [
                  "page",
                  "page_id",
                  "per_page"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "incident_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/incidents/{incident_id}",
              "parts": [
                "pages",
                "{page_id}",
                "incidents",
                "{id}"
              ],
              "rename": {
                "param": {
                  "incident_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "patch": {
          "input": "data",
          "name": "patch",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "incident_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PATCH",
              "orig": "/pages/{page_id}/incidents/{incident_id}",
              "parts": [
                "pages",
                "{page_id}",
                "incidents",
                "{id}"
              ],
              "rename": {
                "param": {
                  "incident_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "incident": "`reqdata`"
                },
                "res": "`body`"
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "incident_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "DELETE",
              "orig": "/pages/{page_id}/incidents/{incident_id}",
              "parts": [
                "pages",
                "{page_id}",
                "incidents",
                "{id}"
              ],
              "rename": {
                "param": {
                  "incident_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "incident_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PUT",
              "orig": "/pages/{page_id}/incidents/{incident_id}",
              "parts": [
                "pages",
                "{page_id}",
                "incidents",
                "{id}"
              ],
              "rename": {
                "param": {
                  "incident_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "incident": "`reqdata`"
                },
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "page"
          ]
        ]
      }
    },
    "incident_postmortem": {
      "fields": [],
      "name": "incident_postmortem",
      "op": {
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "incident_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "DELETE",
              "orig": "/pages/{page_id}/incidents/{incident_id}/postmortem",
              "parts": [
                "pages",
                "{page_id}",
                "incidents",
                "{id}",
                "postmortem"
              ],
              "rename": {
                "param": {
                  "incident_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "page"
          ]
        ]
      }
    },
    "incident_subscriber": {
      "fields": [],
      "name": "incident_subscriber",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "incident_id",
                    "orig": "incident_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "subscriber_id",
                    "orig": "subscriber_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}/resend_confirmation",
              "parts": [
                "pages",
                "{page_id}",
                "incidents",
                "{incident_id}",
                "subscribers",
                "{subscriber_id}",
                "resend_confirmation"
              ],
              "select": {
                "exist": [
                  "incident_id",
                  "page_id",
                  "subscriber_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "page",
            "incident",
            "subscriber"
          ]
        ]
      }
    },
    "incident_template": {
      "fields": [
        {
          "name": "body",
          "short": "Body of the incident or maintenance update to be applied when selecting this template",
          "type": "`$STRING`"
        },
        {
          "name": "components",
          "short": "Affected components",
          "type": "`$ARRAY`"
        },
        {
          "name": "group_id",
          "short": "Identifier of Template Group this template belongs to",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "short": "Incident Template Identifier",
          "type": "`$STRING`"
        },
        {
          "name": "name",
          "short": "Name of the template, as shown in the list on the \"Templates\" tab of the \"Incidents\" page",
          "type": "`$STRING`"
        },
        {
          "name": "should_send_notifications",
          "short": "Whether the \"deliver notifications\" checkbox should be selected when selecting this template",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "should_tweet",
          "short": "Whether the \"tweet update\" checkbox should be selected when selecting this template",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "template",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "title",
          "short": "Title to be applied to the incident or maintenance when selecting this template",
          "type": "`$STRING`"
        },
        {
          "name": "update_status",
          "short": "The status the incident or maintenance should transition to when selecting this template",
          "type": "`$STRING`"
        }
      ],
      "name": "incident_template",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/incident_templates",
              "parts": [
                "pages",
                "{page_id}",
                "incident_templates"
              ],
              "select": {
                "exist": [
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 100,
                    "kind": "query",
                    "name": "per_page",
                    "orig": "per_page",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/incident_templates",
              "parts": [
                "pages",
                "{page_id}",
                "incident_templates"
              ],
              "select": {
                "exist": [
                  "page",
                  "page_id",
                  "per_page"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "page"
          ]
        ]
      }
    },
    "incident_update": {
      "fields": [
        {
          "name": "affected_components",
          "short": "Affected components associated with the incident update.",
          "type": "`$ARRAY`"
        },
        {
          "name": "body",
          "short": "Incident update body.",
          "type": "`$STRING`"
        },
        {
          "name": "created_at",
          "short": "The timestamp when the incident update was created at.",
          "type": "`$STRING`"
        },
        {
          "name": "custom_tweet",
          "short": "An optional customized tweet message for incident postmortem.",
          "type": "`$STRING`"
        },
        {
          "name": "deliver_notifications",
          "short": "Controls whether to delivery notifications.",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "display_at",
          "short": "Timestamp when incident update is happened.",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "short": "Incident Update Identifier.",
          "type": "`$STRING`"
        },
        {
          "name": "incident_id",
          "short": "Incident Identifier.",
          "type": "`$STRING`"
        },
        {
          "name": "incident_update",
          "type": "`$OBJECT`"
        },
        {
          "name": "status",
          "short": "The incident status.",
          "type": "`$STRING`"
        },
        {
          "name": "tweet_id",
          "short": "Tweet identifier associated to this incident update.",
          "type": "`$STRING`"
        },
        {
          "name": "twitter_updated_at",
          "short": "The timestamp when twitter updated at.",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "short": "The timestamp when the incident update is updated.",
          "type": "`$STRING`"
        },
        {
          "name": "wants_twitter_update",
          "short": "Controls whether to create twitter update.",
          "type": "`$BOOLEAN`"
        }
      ],
      "name": "incident_update",
      "op": {
        "patch": {
          "input": "data",
          "name": "patch",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "incident_update_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "incident_id",
                    "orig": "incident_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PATCH",
              "orig": "/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}",
              "parts": [
                "pages",
                "{page_id}",
                "incidents",
                "{incident_id}",
                "incident_updates",
                "{id}"
              ],
              "rename": {
                "param": {
                  "incident_update_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "incident_id",
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "incident_update": "`reqdata`"
                },
                "res": "`body`"
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "incident_update_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "incident_id",
                    "orig": "incident_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PUT",
              "orig": "/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}",
              "parts": [
                "pages",
                "{page_id}",
                "incidents",
                "{incident_id}",
                "incident_updates",
                "{id}"
              ],
              "rename": {
                "param": {
                  "incident_update_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "incident_id",
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "incident_update": "`reqdata`"
                },
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "page",
            "incident"
          ]
        ]
      }
    },
    "metric": {
      "fields": [
        {
          "name": "backfill_percentage",
          "type": "`$INTEGER`"
        },
        {
          "name": "backfilled",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "created_at",
          "type": "`$STRING`"
        },
        {
          "name": "data",
          "req": true,
          "short": "Add data points to metrics",
          "type": "`$OBJECT`"
        },
        {
          "name": "decimal_places",
          "type": "`$INTEGER`"
        },
        {
          "name": "display",
          "short": "Should the metric be displayed",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "id",
          "short": "Metric identifier",
          "type": "`$STRING`"
        },
        {
          "name": "last_fetched_at",
          "type": "`$STRING`"
        },
        {
          "name": "metric",
          "type": "`$OBJECT`"
        },
        {
          "name": "metric_identifier",
          "short": "Metric Display identifier used to look up the metric data from the provider",
          "type": "`$STRING`"
        },
        {
          "name": "metrics_provider_id",
          "short": "Metric Provider identifier",
          "type": "`$STRING`"
        },
        {
          "name": "most_recent_data_at",
          "type": "`$STRING`"
        },
        {
          "name": "name",
          "short": "Name of metric",
          "type": "`$STRING`"
        },
        {
          "name": "reference_name",
          "type": "`$STRING`"
        },
        {
          "name": "suffix",
          "short": "Suffix to describe the units on the graph",
          "type": "`$STRING`"
        },
        {
          "name": "tooltip_description",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$STRING`"
        },
        {
          "name": "y_axis_hidden",
          "short": "Should the values on the y axis be hidden on render",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "y_axis_max",
          "type": "`$NUMBER`"
        },
        {
          "name": "y_axis_min",
          "type": "`$NUMBER`"
        }
      ],
      "name": "metric",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "metric_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/metrics/{metric_id}/data",
              "parts": [
                "pages",
                "{page_id}",
                "metrics",
                "{id}",
                "data"
              ],
              "rename": {
                "param": {
                  "metric_id": "id"
                }
              },
              "select": {
                "$action": "data",
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.data`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "metrics_provider_id",
                    "orig": "metrics_provider_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics",
              "parts": [
                "pages",
                "{page_id}",
                "metrics_providers",
                "{metrics_provider_id}",
                "metrics"
              ],
              "select": {
                "exist": [
                  "metrics_provider_id",
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "metric": "`reqdata`"
                },
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/metrics/data",
              "parts": [
                "pages",
                "{page_id}",
                "metrics",
                "data"
              ],
              "select": {
                "$action": "data",
                "exist": [
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_access_user_id",
                    "orig": "page_access_user_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "per_page",
                    "orig": "per_page",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_users",
                "{page_access_user_id}",
                "metrics"
              ],
              "select": {
                "exist": [
                  "page",
                  "page_access_user_id",
                  "page_id",
                  "per_page"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "metrics_provider_id",
                    "orig": "metrics_provider_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "per_page",
                    "orig": "per_page",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics",
              "parts": [
                "pages",
                "{page_id}",
                "metrics_providers",
                "{metrics_provider_id}",
                "metrics"
              ],
              "select": {
                "exist": [
                  "metrics_provider_id",
                  "page",
                  "page_id",
                  "per_page"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "per_page",
                    "orig": "per_page",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/metrics",
              "parts": [
                "pages",
                "{page_id}",
                "metrics"
              ],
              "select": {
                "exist": [
                  "page",
                  "page_id",
                  "per_page"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "metric_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/metrics/{metric_id}",
              "parts": [
                "pages",
                "{page_id}",
                "metrics",
                "{id}"
              ],
              "rename": {
                "param": {
                  "metric_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "patch": {
          "input": "data",
          "name": "patch",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "metric_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PATCH",
              "orig": "/pages/{page_id}/metrics/{metric_id}",
              "parts": [
                "pages",
                "{page_id}",
                "metrics",
                "{id}"
              ],
              "rename": {
                "param": {
                  "metric_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "metric": "`reqdata`"
                },
                "res": "`body`"
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "metric_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "DELETE",
              "orig": "/pages/{page_id}/metrics/{metric_id}",
              "parts": [
                "pages",
                "{page_id}",
                "metrics",
                "{id}"
              ],
              "rename": {
                "param": {
                  "metric_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "metric_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "DELETE",
              "orig": "/pages/{page_id}/metrics/{metric_id}/data",
              "parts": [
                "pages",
                "{page_id}",
                "metrics",
                "{id}",
                "data"
              ],
              "rename": {
                "param": {
                  "metric_id": "id"
                }
              },
              "select": {
                "$action": "data",
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "metric_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PUT",
              "orig": "/pages/{page_id}/metrics/{metric_id}",
              "parts": [
                "pages",
                "{page_id}",
                "metrics",
                "{id}"
              ],
              "rename": {
                "param": {
                  "metric_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "metric": "`reqdata`"
                },
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "page"
          ],
          [
            "page",
            "metrics_provider"
          ],
          [
            "page",
            "page_access_user"
          ]
        ]
      }
    },
    "metrics_provider": {
      "fields": [
        {
          "name": "created_at",
          "type": "`$STRING`"
        },
        {
          "name": "disabled",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "id",
          "short": "Identifier for Metrics Provider",
          "type": "`$STRING`"
        },
        {
          "name": "last_revalidated_at",
          "type": "`$STRING`"
        },
        {
          "name": "metric_base_uri",
          "type": "`$STRING`"
        },
        {
          "name": "metrics_provider",
          "type": "`$OBJECT`"
        },
        {
          "name": "page_id",
          "type": "`$INTEGER`"
        },
        {
          "name": "type",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$STRING`"
        }
      ],
      "name": "metrics_provider",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/metrics_providers",
              "parts": [
                "pages",
                "{page_id}",
                "metrics_providers"
              ],
              "select": {
                "exist": [
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "metrics_provider": "`reqdata`"
                },
                "res": "`body`"
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/metrics_providers",
              "parts": [
                "pages",
                "{page_id}",
                "metrics_providers"
              ],
              "select": {
                "exist": [
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "metrics_provider_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
              "parts": [
                "pages",
                "{page_id}",
                "metrics_providers",
                "{id}"
              ],
              "rename": {
                "param": {
                  "metrics_provider_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "patch": {
          "input": "data",
          "name": "patch",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "metrics_provider_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PATCH",
              "orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
              "parts": [
                "pages",
                "{page_id}",
                "metrics_providers",
                "{id}"
              ],
              "rename": {
                "param": {
                  "metrics_provider_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "metrics_provider": "`reqdata`"
                },
                "res": "`body`"
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "metrics_provider_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "DELETE",
              "orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
              "parts": [
                "pages",
                "{page_id}",
                "metrics_providers",
                "{id}"
              ],
              "rename": {
                "param": {
                  "metrics_provider_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "metrics_provider_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PUT",
              "orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
              "parts": [
                "pages",
                "{page_id}",
                "metrics_providers",
                "{id}"
              ],
              "rename": {
                "param": {
                  "metrics_provider_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "metrics_provider": "`reqdata`"
                },
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "page"
          ]
        ]
      }
    },
    "page": {
      "fields": [
        {
          "name": "activity_score",
          "type": "`$NUMBER`"
        },
        {
          "name": "allow_email_subscribers",
          "short": "Can your users choose to receive notifications via email",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "allow_incident_subscribers",
          "short": "Can your users subscribe to notifications for a single incident",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "allow_page_subscribers",
          "short": "Can your users subscribe to all notifications on the page",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "allow_rss_atom_feeds",
          "short": "Can your users choose to access incident feeds via RSS/Atom (not functional on Audience-Specific pages)",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "allow_sms_subscribers",
          "short": "Can your users choose to receive notifications via SMS",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "allow_webhook_subscribers",
          "short": "Can your users choose to receive notifications via Webhooks",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "branding",
          "short": "The main template your statuspage will use",
          "type": "`$STRING`"
        },
        {
          "name": "city",
          "type": "`$STRING`"
        },
        {
          "name": "country",
          "type": "`$STRING`"
        },
        {
          "name": "created_at",
          "short": "Timestamp the record was created",
          "type": "`$STRING`"
        },
        {
          "name": "css_blues",
          "short": "CSS Color",
          "type": "`$STRING`"
        },
        {
          "name": "css_body_background_color",
          "short": "CSS Color",
          "type": "`$STRING`"
        },
        {
          "name": "css_border_color",
          "short": "CSS Color",
          "type": "`$STRING`"
        },
        {
          "name": "css_font_color",
          "short": "CSS Color",
          "type": "`$STRING`"
        },
        {
          "name": "css_graph_color",
          "short": "CSS Color",
          "type": "`$STRING`"
        },
        {
          "name": "css_greens",
          "short": "CSS Color",
          "type": "`$STRING`"
        },
        {
          "name": "css_light_font_color",
          "short": "CSS Color",
          "type": "`$STRING`"
        },
        {
          "name": "css_link_color",
          "short": "CSS Color",
          "type": "`$STRING`"
        },
        {
          "name": "css_no_data",
          "short": "CSS Color",
          "type": "`$STRING`"
        },
        {
          "name": "css_oranges",
          "short": "CSS Color",
          "type": "`$STRING`"
        },
        {
          "name": "css_reds",
          "short": "CSS Color",
          "type": "`$STRING`"
        },
        {
          "name": "css_yellows",
          "short": "CSS Color",
          "type": "`$STRING`"
        },
        {
          "name": "domain",
          "short": "CNAME alias for your status page",
          "type": "`$STRING`"
        },
        {
          "name": "email_logo",
          "type": "`$STRING`"
        },
        {
          "name": "favicon_logo",
          "type": "`$STRING`"
        },
        {
          "name": "headline",
          "type": "`$STRING`"
        },
        {
          "name": "hero_cover",
          "type": "`$STRING`"
        },
        {
          "name": "hidden_from_search",
          "short": "Should your page hide itself from search engines",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "id",
          "short": "Page identifier",
          "type": "`$STRING`"
        },
        {
          "name": "ip_restrictions",
          "type": "`$STRING`"
        },
        {
          "name": "name",
          "short": "Name of your page to be displayed",
          "type": "`$STRING`"
        },
        {
          "name": "notifications_email_footer",
          "short": "Allows you to customize the footer appearing on your notification emails.",
          "type": "`$STRING`"
        },
        {
          "name": "notifications_from_email",
          "short": "Allows you to customize the email address your page notifications come from",
          "type": "`$STRING`"
        },
        {
          "name": "page",
          "type": "`$OBJECT`"
        },
        {
          "name": "page_description",
          "type": "`$STRING`"
        },
        {
          "name": "state",
          "type": "`$STRING`"
        },
        {
          "name": "subdomain",
          "short": "Subdomain at which to access your status page",
          "type": "`$STRING`"
        },
        {
          "name": "support_url",
          "type": "`$STRING`"
        },
        {
          "name": "time_zone",
          "short": "Timezone configured for your page",
          "type": "`$STRING`"
        },
        {
          "name": "transactional_logo",
          "type": "`$STRING`"
        },
        {
          "name": "twitter_logo",
          "type": "`$STRING`"
        },
        {
          "name": "twitter_username",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "short": "Timestamp the record was last updated",
          "type": "`$STRING`"
        },
        {
          "name": "url",
          "short": "Website of your page.",
          "type": "`$STRING`"
        },
        {
          "name": "viewers_must_be_team_members",
          "type": "`$BOOLEAN`"
        }
      ],
      "name": "page",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/pages",
              "parts": [
                "pages"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}",
              "parts": [
                "pages",
                "{id}"
              ],
              "rename": {
                "param": {
                  "page_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "patch": {
          "input": "data",
          "name": "patch",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PATCH",
              "orig": "/pages/{page_id}",
              "parts": [
                "pages",
                "{id}"
              ],
              "rename": {
                "param": {
                  "page_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": {
                  "page": "`reqdata`"
                },
                "res": "`body`"
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PUT",
              "orig": "/pages/{page_id}",
              "parts": [
                "pages",
                "{id}"
              ],
              "rename": {
                "param": {
                  "page_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": {
                  "page": "`reqdata`"
                },
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "page_access_group": {
      "fields": [
        {
          "name": "component_ids",
          "op": {
            "create": {
              "req": true,
              "type": "`$ARRAY`"
            }
          },
          "short": "List of components codes to set on the page access group",
          "type": "`$ARRAY`"
        },
        {
          "name": "created_at",
          "type": "`$STRING`"
        },
        {
          "name": "external_identifier",
          "short": "Associates group with external group.",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "short": "Page Access Group Identifier",
          "type": "`$STRING`"
        },
        {
          "name": "metric_ids",
          "type": "`$ARRAY`"
        },
        {
          "name": "name",
          "short": "Name for this Group.",
          "type": "`$STRING`"
        },
        {
          "name": "page_access_group",
          "type": "`$OBJECT`"
        },
        {
          "name": "page_access_user_ids",
          "type": "`$ARRAY`"
        },
        {
          "name": "page_id",
          "short": "Page Identifier.",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$STRING`"
        }
      ],
      "name": "page_access_group",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_group_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_groups",
                "{id}",
                "components"
              ],
              "rename": {
                "param": {
                  "page_access_group_id": "id"
                }
              },
              "select": {
                "$action": "component",
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/page_access_groups",
              "parts": [
                "pages",
                "{id}",
                "page_access_groups"
              ],
              "rename": {
                "param": {
                  "page_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": {
                  "page_access_group": "`reqdata`"
                },
                "res": "`body`"
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "per_page",
                    "orig": "per_page",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/page_access_groups",
              "parts": [
                "pages",
                "{id}",
                "page_access_groups"
              ],
              "rename": {
                "param": {
                  "page_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page",
                  "per_page"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_group_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_groups",
                "{id}"
              ],
              "rename": {
                "param": {
                  "page_access_group_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "patch": {
          "input": "data",
          "name": "patch",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_group_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PATCH",
              "orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_groups",
                "{id}"
              ],
              "rename": {
                "param": {
                  "page_access_group_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "page_access_group": "`reqdata`"
                },
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_group_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PATCH",
              "orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_groups",
                "{id}",
                "components"
              ],
              "rename": {
                "param": {
                  "page_access_group_id": "id"
                }
              },
              "select": {
                "$action": "component",
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "component_id",
                    "orig": "component_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_group_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "DELETE",
              "orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components/{component_id}",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_groups",
                "{id}",
                "components",
                "{component_id}"
              ],
              "rename": {
                "param": {
                  "page_access_group_id": "id"
                }
              },
              "select": {
                "exist": [
                  "component_id",
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_group_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "DELETE",
              "orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_groups",
                "{id}"
              ],
              "rename": {
                "param": {
                  "page_access_group_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_group_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "DELETE",
              "orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_groups",
                "{id}",
                "components"
              ],
              "rename": {
                "param": {
                  "page_access_group_id": "id"
                }
              },
              "select": {
                "$action": "component",
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_group_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PUT",
              "orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_groups",
                "{id}"
              ],
              "rename": {
                "param": {
                  "page_access_group_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "page_access_group": "`reqdata`"
                },
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_group_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PUT",
              "orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_groups",
                "{id}",
                "components"
              ],
              "rename": {
                "param": {
                  "page_access_group_id": "id"
                }
              },
              "select": {
                "$action": "component",
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "page"
          ],
          [
            "page",
            "component"
          ]
        ]
      }
    },
    "page_access_user": {
      "fields": [
        {
          "name": "component_ids",
          "req": true,
          "short": "List of component codes to allow access to",
          "type": "`$ARRAY`"
        },
        {
          "name": "created_at",
          "type": "`$STRING`"
        },
        {
          "name": "email",
          "type": "`$STRING`"
        },
        {
          "name": "external_login",
          "short": "IDP login user id.",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "short": "Page Access User Identifier",
          "type": "`$STRING`"
        },
        {
          "name": "metric_ids",
          "req": true,
          "short": "List of metrics to add",
          "type": "`$ARRAY`"
        },
        {
          "name": "page_access_group_id",
          "type": "`$STRING`"
        },
        {
          "name": "page_access_group_ids",
          "type": "`$STRING`"
        },
        {
          "name": "page_access_user",
          "type": "`$OBJECT`"
        },
        {
          "name": "page_id",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$STRING`"
        }
      ],
      "name": "page_access_user",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_user_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_users",
                "{id}",
                "components"
              ],
              "rename": {
                "param": {
                  "page_access_user_id": "id"
                }
              },
              "select": {
                "$action": "component",
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_user_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_users",
                "{id}",
                "metrics"
              ],
              "rename": {
                "param": {
                  "page_access_user_id": "id"
                }
              },
              "select": {
                "$action": "metric",
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/page_access_users",
              "parts": [
                "pages",
                "{id}",
                "page_access_users"
              ],
              "rename": {
                "param": {
                  "page_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": {
                  "page_access_user": "`reqdata`"
                },
                "res": "`body`"
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "email",
                    "orig": "email",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "per_page",
                    "orig": "per_page",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/page_access_users",
              "parts": [
                "pages",
                "{id}",
                "page_access_users"
              ],
              "rename": {
                "param": {
                  "page_id": "id"
                }
              },
              "select": {
                "exist": [
                  "email",
                  "id",
                  "page",
                  "per_page"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_user_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_users",
                "{id}"
              ],
              "rename": {
                "param": {
                  "page_access_user_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "patch": {
          "input": "data",
          "name": "patch",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_user_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PATCH",
              "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_users",
                "{id}"
              ],
              "rename": {
                "param": {
                  "page_access_user_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_user_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PATCH",
              "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_users",
                "{id}",
                "components"
              ],
              "rename": {
                "param": {
                  "page_access_user_id": "id"
                }
              },
              "select": {
                "$action": "component",
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_user_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PATCH",
              "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_users",
                "{id}",
                "metrics"
              ],
              "rename": {
                "param": {
                  "page_access_user_id": "id"
                }
              },
              "select": {
                "$action": "metric",
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "component_id",
                    "orig": "component_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_user_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "DELETE",
              "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components/{component_id}",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_users",
                "{id}",
                "components",
                "{component_id}"
              ],
              "rename": {
                "param": {
                  "page_access_user_id": "id"
                }
              },
              "select": {
                "exist": [
                  "component_id",
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_user_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "metric_id",
                    "orig": "metric_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "DELETE",
              "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics/{metric_id}",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_users",
                "{id}",
                "metrics",
                "{metric_id}"
              ],
              "rename": {
                "param": {
                  "page_access_user_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "metric_id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_user_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "DELETE",
              "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_users",
                "{id}"
              ],
              "rename": {
                "param": {
                  "page_access_user_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_user_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "DELETE",
              "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_users",
                "{id}",
                "components"
              ],
              "rename": {
                "param": {
                  "page_access_user_id": "id"
                }
              },
              "select": {
                "$action": "component",
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_user_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "DELETE",
              "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_users",
                "{id}",
                "metrics"
              ],
              "rename": {
                "param": {
                  "page_access_user_id": "id"
                }
              },
              "select": {
                "$action": "metric",
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_user_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PUT",
              "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_users",
                "{id}"
              ],
              "rename": {
                "param": {
                  "page_access_user_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_user_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PUT",
              "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_users",
                "{id}",
                "components"
              ],
              "rename": {
                "param": {
                  "page_access_user_id": "id"
                }
              },
              "select": {
                "$action": "component",
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "page_access_user_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PUT",
              "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
              "parts": [
                "pages",
                "{page_id}",
                "page_access_users",
                "{id}",
                "metrics"
              ],
              "rename": {
                "param": {
                  "page_access_user_id": "id"
                }
              },
              "select": {
                "$action": "metric",
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "page"
          ],
          [
            "page",
            "component"
          ],
          [
            "page",
            "metric"
          ]
        ]
      }
    },
    "permission": {
      "fields": [
        {
          "name": "pages",
          "short": "Pages accessible by the user.",
          "type": "`$OBJECT`"
        },
        {
          "name": "user_id",
          "short": "User identifier",
          "type": "`$STRING`"
        }
      ],
      "name": "permission",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "user_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "organization_id",
                    "orig": "organization_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/organizations/{organization_id}/permissions/{user_id}",
              "parts": [
                "organizations",
                "{organization_id}",
                "permissions",
                "{id}"
              ],
              "rename": {
                "param": {
                  "user_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "organization_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.data`"
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "user_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "organization_id",
                    "orig": "organization_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PUT",
              "orig": "/organizations/{organization_id}/permissions/{user_id}",
              "parts": [
                "organizations",
                "{organization_id}",
                "permissions",
                "{id}"
              ],
              "rename": {
                "param": {
                  "user_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "organization_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.data`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "organization"
          ]
        ]
      }
    },
    "postmortem": {
      "fields": [
        {
          "name": "body",
          "short": "Postmortem body",
          "type": "`$STRING`"
        },
        {
          "name": "body_draft",
          "short": "Body draft",
          "type": "`$STRING`"
        },
        {
          "name": "body_draft_updated_at",
          "type": "`$STRING`"
        },
        {
          "name": "body_updated_at",
          "type": "`$STRING`"
        },
        {
          "name": "created_at",
          "type": "`$STRING`"
        },
        {
          "name": "custom_tweet",
          "short": "Custom tweet for Incident Postmortem",
          "type": "`$STRING`"
        },
        {
          "name": "notify_subscribers",
          "short": "Should email subscribers be notified.",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "notify_twitter",
          "short": "Should Twitter followers be notified.",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "postmortem",
          "op": {
            "update": {
              "type": "`$OBJECT`"
            }
          },
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "preview_key",
          "short": "Preview Key",
          "type": "`$STRING`"
        },
        {
          "name": "published_at",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$STRING`"
        }
      ],
      "name": "postmortem",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "incident_id",
                    "orig": "incident_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/incidents/{incident_id}/postmortem",
              "parts": [
                "pages",
                "{page_id}",
                "incidents",
                "{incident_id}",
                "postmortem"
              ],
              "select": {
                "exist": [
                  "incident_id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "incident_id",
                    "orig": "incident_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PUT",
              "orig": "/pages/{page_id}/incidents/{incident_id}/postmortem",
              "parts": [
                "pages",
                "{page_id}",
                "incidents",
                "{incident_id}",
                "postmortem"
              ],
              "select": {
                "exist": [
                  "incident_id",
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "postmortem": "`reqdata`"
                },
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "incident_id",
                    "orig": "incident_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PUT",
              "orig": "/pages/{page_id}/incidents/{incident_id}/postmortem/publish",
              "parts": [
                "pages",
                "{page_id}",
                "incidents",
                "{incident_id}",
                "postmortem",
                "publish"
              ],
              "select": {
                "$action": "publish",
                "exist": [
                  "incident_id",
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "postmortem": "`reqdata`"
                },
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "incident_id",
                    "orig": "incident_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PUT",
              "orig": "/pages/{page_id}/incidents/{incident_id}/postmortem/revert",
              "parts": [
                "pages",
                "{page_id}",
                "incidents",
                "{incident_id}",
                "postmortem",
                "revert"
              ],
              "select": {
                "$action": "revert",
                "exist": [
                  "incident_id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "page",
            "incident"
          ]
        ]
      }
    },
    "status_embed_config": {
      "fields": [
        {
          "name": "incident_background_color",
          "short": "Color of status embed iframe background when displaying incident",
          "type": "`$STRING`"
        },
        {
          "name": "incident_text_color",
          "short": "Color of status embed iframe text when displaying incident",
          "type": "`$STRING`"
        },
        {
          "name": "maintenance_background_color",
          "short": "Color of status embed iframe background when displaying maintenance",
          "type": "`$STRING`"
        },
        {
          "name": "maintenance_text_color",
          "short": "Color of status embed iframe text when displaying maintenance",
          "type": "`$STRING`"
        },
        {
          "name": "page_id",
          "short": "Page identifier",
          "type": "`$STRING`"
        },
        {
          "name": "position",
          "short": "Corner where status embed iframe will appear on page",
          "type": "`$STRING`"
        },
        {
          "name": "status_embed_config",
          "type": "`$OBJECT`"
        }
      ],
      "name": "status_embed_config",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/status_embed_config",
              "parts": [
                "pages",
                "{page_id}",
                "status_embed_config"
              ],
              "select": {
                "exist": [
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "patch": {
          "input": "data",
          "name": "patch",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PATCH",
              "orig": "/pages/{page_id}/status_embed_config",
              "parts": [
                "pages",
                "{page_id}",
                "status_embed_config"
              ],
              "select": {
                "exist": [
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "status_embed_config": "`reqdata`"
                },
                "res": "`body`"
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PUT",
              "orig": "/pages/{page_id}/status_embed_config",
              "parts": [
                "pages",
                "{page_id}",
                "status_embed_config"
              ],
              "select": {
                "exist": [
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "status_embed_config": "`reqdata`"
                },
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "page"
          ]
        ]
      }
    },
    "subscriber": {
      "fields": [
        {
          "name": "component_ids",
          "short": "A list of component ids for which the subscriber should recieve updates for.",
          "type": "`$ARRAY`"
        },
        {
          "name": "components",
          "short": "The components for which the subscriber has elected to receive updates.",
          "type": "`$STRING`"
        },
        {
          "name": "created_at",
          "type": "`$STRING`"
        },
        {
          "name": "display_phone_number",
          "short": "A formatted version of the phone_number and phone_country pair, nicely formatted for display.",
          "type": "`$STRING`"
        },
        {
          "name": "email",
          "short": "The email address to use to contact the subscriber.",
          "type": "`$STRING`"
        },
        {
          "name": "endpoint",
          "short": "The URL where a webhook subscriber elects to receive updates.",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "short": "Subscriber Identifier",
          "type": "`$STRING`"
        },
        {
          "name": "integration_partner",
          "short": "The number of integration partners found by the query.",
          "type": "`$INTEGER`"
        },
        {
          "name": "mode",
          "short": "The communication mode of the subscriber.",
          "type": "`$STRING`"
        },
        {
          "name": "obfuscated_channel_name",
          "short": "Obfuscated slack channel name",
          "type": "`$STRING`"
        },
        {
          "name": "page_access_user_id",
          "short": "The Page Access user this subscriber belongs to (only for audience-specific pages).",
          "type": "`$STRING`"
        },
        {
          "name": "phone_country",
          "short": "The two-character country code representing the country of which the phone_number is a part.",
          "type": "`$STRING`"
        },
        {
          "name": "phone_number",
          "short": "The phone number used to contact an SMS subscriber",
          "type": "`$STRING`"
        },
        {
          "name": "purge_at",
          "short": "The timestamp when a quarantined subscriber will be purged (unsubscribed).",
          "type": "`$STRING`"
        },
        {
          "name": "quarantined_at",
          "short": "The timestamp when the subscriber was quarantined due to an issue reaching them.",
          "type": "`$STRING`"
        },
        {
          "name": "skip_confirmation_notification",
          "short": "If this is true, do not notify the user with changes to their subscription.",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "skip_unsubscription_notification",
          "short": "If skip_unsubscription_notification is true, the subscribers do not receive any notifications when they are unsubscribed.",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "slack",
          "short": "The number of Slack subscribers found by the query.",
          "type": "`$INTEGER`"
        },
        {
          "name": "sms",
          "short": "The number of Webhook subscribers found by the query.",
          "type": "`$INTEGER`"
        },
        {
          "name": "state",
          "short": "If this is present, only unsubscribe subscribers in this state.",
          "type": "`$STRING`"
        },
        {
          "name": "subscriber",
          "type": "`$OBJECT`"
        },
        {
          "name": "subscribers",
          "req": true,
          "short": "The array of quarantined subscriber codes to reactivate, or \"all\" to reactivate all quarantined subscribers.",
          "type": "`$STRING`"
        },
        {
          "name": "teams",
          "short": "The number of MS teams subscribers found by the query.",
          "type": "`$INTEGER`"
        },
        {
          "name": "type",
          "short": "If this is present, only reactivate subscribers of this type.",
          "type": "`$STRING`"
        },
        {
          "name": "webhook",
          "short": "The number of SMS subscribers found by the query.",
          "type": "`$INTEGER`"
        },
        {
          "name": "workspace_name",
          "short": "The workspace name of the slack subscriber.",
          "type": "`$STRING`"
        }
      ],
      "name": "subscriber",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "subscriber_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/subscribers/{subscriber_id}/resend_confirmation",
              "parts": [
                "pages",
                "{page_id}",
                "subscribers",
                "{id}",
                "resend_confirmation"
              ],
              "rename": {
                "param": {
                  "subscriber_id": "id"
                }
              },
              "select": {
                "$action": "resend_confirmation",
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "incident_id",
                    "orig": "incident_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/incidents/{incident_id}/subscribers",
              "parts": [
                "pages",
                "{page_id}",
                "incidents",
                "{incident_id}",
                "subscribers"
              ],
              "select": {
                "exist": [
                  "incident_id",
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "subscriber": "`reqdata`"
                },
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/subscribers",
              "parts": [
                "pages",
                "{page_id}",
                "subscribers"
              ],
              "select": {
                "exist": [
                  "page_id"
                ]
              },
              "transform": {
                "req": {
                  "subscriber": "`reqdata`"
                },
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/subscribers/reactivate",
              "parts": [
                "pages",
                "{page_id}",
                "subscribers",
                "reactivate"
              ],
              "select": {
                "$action": "reactivate",
                "exist": [
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/subscribers/resend_confirmation",
              "parts": [
                "pages",
                "{page_id}",
                "subscribers",
                "resend_confirmation"
              ],
              "select": {
                "$action": "resend_confirmation",
                "exist": [
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/pages/{page_id}/subscribers/unsubscribe",
              "parts": [
                "pages",
                "{page_id}",
                "subscribers",
                "unsubscribe"
              ],
              "select": {
                "$action": "unsubscribe",
                "exist": [
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "limit",
                    "orig": "limit",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "asc",
                    "kind": "query",
                    "name": "sort_direction",
                    "orig": "sort_direction",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "primary",
                    "kind": "query",
                    "name": "sort_field",
                    "orig": "sort_field",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "active",
                    "kind": "query",
                    "name": "state",
                    "orig": "state",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "type",
                    "orig": "type",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/subscribers",
              "parts": [
                "pages",
                "{page_id}",
                "subscribers"
              ],
              "select": {
                "exist": [
                  "limit",
                  "page",
                  "page_id",
                  "q",
                  "sort_direction",
                  "sort_field",
                  "state",
                  "type"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "incident_id",
                    "orig": "incident_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "per_page",
                    "orig": "per_page",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/incidents/{incident_id}/subscribers",
              "parts": [
                "pages",
                "{page_id}",
                "incidents",
                "{incident_id}",
                "subscribers"
              ],
              "select": {
                "exist": [
                  "incident_id",
                  "page",
                  "page_id",
                  "per_page"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "per_page",
                    "orig": "per_page",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/subscribers/unsubscribed",
              "parts": [
                "pages",
                "{page_id}",
                "subscribers",
                "unsubscribed"
              ],
              "select": {
                "$action": "unsubscribed",
                "exist": [
                  "page",
                  "page_id",
                  "per_page"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "subscriber_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "incident_id",
                    "orig": "incident_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}",
              "parts": [
                "pages",
                "{page_id}",
                "incidents",
                "{incident_id}",
                "subscribers",
                "{id}"
              ],
              "rename": {
                "param": {
                  "subscriber_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "incident_id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "example": "active",
                    "kind": "query",
                    "name": "state",
                    "orig": "state",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "type",
                    "orig": "type",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/subscribers/count",
              "parts": [
                "pages",
                "{page_id}",
                "subscribers",
                "count"
              ],
              "select": {
                "$action": "count",
                "exist": [
                  "page_id",
                  "state",
                  "type"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "subscriber_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/subscribers/{subscriber_id}",
              "parts": [
                "pages",
                "{page_id}",
                "subscribers",
                "{id}"
              ],
              "rename": {
                "param": {
                  "subscriber_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/pages/{page_id}/subscribers/histogram_by_state",
              "parts": [
                "pages",
                "{page_id}",
                "subscribers",
                "histogram_by_state"
              ],
              "select": {
                "$action": "histogram_by_state",
                "exist": [
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "subscriber_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "incident_id",
                    "orig": "incident_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "DELETE",
              "orig": "/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}",
              "parts": [
                "pages",
                "{page_id}",
                "incidents",
                "{incident_id}",
                "subscribers",
                "{id}"
              ],
              "rename": {
                "param": {
                  "subscriber_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "incident_id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "subscriber_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "skip_unsubscription_notification",
                    "orig": "skip_unsubscription_notification",
                    "type": "`$BOOLEAN`"
                  }
                ]
              },
              "kind": "http",
              "method": "DELETE",
              "orig": "/pages/{page_id}/subscribers/{subscriber_id}",
              "parts": [
                "pages",
                "{page_id}",
                "subscribers",
                "{id}"
              ],
              "rename": {
                "param": {
                  "subscriber_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id",
                  "skip_unsubscription_notification"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "subscriber_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "page_id",
                    "orig": "page_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "PATCH",
              "orig": "/pages/{page_id}/subscribers/{subscriber_id}",
              "parts": [
                "pages",
                "{page_id}",
                "subscribers",
                "{id}"
              ],
              "rename": {
                "param": {
                  "subscriber_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "page_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "page"
          ],
          [
            "page",
            "incident"
          ]
        ]
      }
    },
    "user": {
      "fields": [
        {
          "name": "created_at",
          "type": "`$STRING`"
        },
        {
          "name": "email",
          "short": "Email address for the team member",
          "type": "`$STRING`"
        },
        {
          "name": "first_name",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "short": "User identifier",
          "type": "`$STRING`"
        },
        {
          "name": "last_name",
          "type": "`$STRING`"
        },
        {
          "name": "organization_id",
          "short": "Organization identifier",
          "type": "`$STRING`"
        },
        {
          "name": "updated_at",
          "type": "`$STRING`"
        },
        {
          "name": "user",
          "req": true,
          "type": "`$OBJECT`"
        }
      ],
      "name": "user",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "organization_id",
                    "orig": "organization_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "POST",
              "orig": "/organizations/{organization_id}/users",
              "parts": [
                "organizations",
                "{organization_id}",
                "users"
              ],
              "select": {
                "exist": [
                  "organization_id"
                ]
              },
              "transform": {
                "req": {
                  "user": "`reqdata`"
                },
                "res": "`body`"
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "organization_id",
                    "orig": "organization_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "per_page",
                    "orig": "per_page",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/organizations/{organization_id}/users",
              "parts": [
                "organizations",
                "{organization_id}",
                "users"
              ],
              "select": {
                "exist": [
                  "organization_id",
                  "page",
                  "per_page"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "user_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "organization_id",
                    "orig": "organization_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "DELETE",
              "orig": "/organizations/{organization_id}/users/{user_id}",
              "parts": [
                "organizations",
                "{organization_id}",
                "users",
                "{id}"
              ],
              "rename": {
                "param": {
                  "user_id": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "organization_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "organization"
          ]
        ]
      }
    }
  }
}


const config = new Config()

export {
  config
}

