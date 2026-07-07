# Statuspage SDK configuration


def make_config():
    return {
        "main": {
            "name": "Statuspage",
        },
        "feature": {
            "test": {
        "options": {
          "active": False,
        },
      },
        },
        "options": {
            "base": "https://api.statuspage.io/v1",
            "auth": {
                "prefix": "OAuth",
            },
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "component": {},
                "component_group_uptime": {},
                "group_component": {},
                "incident": {},
                "incident_postmortem": {},
                "incident_subscriber": {},
                "incident_template": {},
                "incident_update": {},
                "metric": {},
                "metrics_provider": {},
                "page": {},
                "page_access_group": {},
                "page_access_user": {},
                "permission": {},
                "postmortem": {},
                "status_embed_config": {},
                "subscriber": {},
                "user": {},
            },
        },
        "entity": {
      "component": {
        "fields": [
          {
            "active": True,
            "name": "automation_email",
            "req": False,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "component",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "created_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "description",
            "req": False,
            "type": "`$STRING`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "group",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "group_id",
            "req": False,
            "type": "`$STRING`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "id",
            "req": False,
            "type": "`$STRING`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "major_outage",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "name",
            "req": False,
            "type": "`$STRING`",
            "index$": 8,
          },
          {
            "active": True,
            "name": "only_show_if_degraded",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 9,
          },
          {
            "active": True,
            "name": "page_id",
            "req": False,
            "type": "`$STRING`",
            "index$": 10,
          },
          {
            "active": True,
            "name": "partial_outage",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 11,
          },
          {
            "active": True,
            "name": "position",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 12,
          },
          {
            "active": True,
            "name": "range_end",
            "req": False,
            "type": "`$STRING`",
            "index$": 13,
          },
          {
            "active": True,
            "name": "range_start",
            "req": False,
            "type": "`$STRING`",
            "index$": 14,
          },
          {
            "active": True,
            "name": "related_event",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 15,
          },
          {
            "active": True,
            "name": "showcase",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 16,
          },
          {
            "active": True,
            "name": "start_date",
            "req": False,
            "type": "`$STRING`",
            "index$": 17,
          },
          {
            "active": True,
            "name": "status",
            "req": False,
            "type": "`$STRING`",
            "index$": 18,
          },
          {
            "active": True,
            "name": "updated_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 19,
          },
          {
            "active": True,
            "name": "uptime_percentage",
            "req": False,
            "type": "`$NUMBER`",
            "index$": 20,
          },
          {
            "active": True,
            "name": "warning",
            "req": False,
            "type": "`$STRING`",
            "index$": 21,
          },
        ],
        "name": "component",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "component_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/components/{component_id}/page_access_groups",
                "parts": [
                  "pages",
                  "{page_id}",
                  "components",
                  "{id}",
                  "page_access_groups",
                ],
                "rename": {
                  "param": {
                    "component_id": "id",
                  },
                },
                "select": {
                  "$action": "page_access_group",
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "component_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/components/{component_id}/page_access_users",
                "parts": [
                  "pages",
                  "{page_id}",
                  "components",
                  "{id}",
                  "page_access_users",
                ],
                "rename": {
                  "param": {
                    "component_id": "id",
                  },
                },
                "select": {
                  "$action": "page_access_user",
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/components",
                "parts": [
                  "pages",
                  "{page_id}",
                  "components",
                ],
                "select": {
                  "exist": [
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
            ],
            "key$": "create",
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_access_group_id",
                      "orig": "page_access_group_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "page",
                      "orig": "page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "per_page",
                      "orig": "per_page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_groups",
                  "{page_access_group_id}",
                  "components",
                ],
                "select": {
                  "exist": [
                    "page",
                    "page_access_group_id",
                    "page_id",
                    "per_page",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_access_user_id",
                      "orig": "page_access_user_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "page",
                      "orig": "page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "per_page",
                      "orig": "per_page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_users",
                  "{page_access_user_id}",
                  "components",
                ],
                "select": {
                  "exist": [
                    "page",
                    "page_access_user_id",
                    "page_id",
                    "per_page",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "page",
                      "orig": "page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "per_page",
                      "orig": "per_page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/components",
                "parts": [
                  "pages",
                  "{page_id}",
                  "components",
                ],
                "select": {
                  "exist": [
                    "page",
                    "page_id",
                    "per_page",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
            ],
            "key$": "list",
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "component_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "end",
                      "orig": "end",
                      "reqd": False,
                      "type": "Any",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "skip_related_event",
                      "orig": "skip_related_event",
                      "reqd": False,
                      "type": "`$BOOLEAN`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "start",
                      "orig": "start",
                      "reqd": False,
                      "type": "Any",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/components/{component_id}/uptime",
                "parts": [
                  "pages",
                  "{page_id}",
                  "components",
                  "{id}",
                  "uptime",
                ],
                "rename": {
                  "param": {
                    "component_id": "id",
                  },
                },
                "select": {
                  "$action": "uptime",
                  "exist": [
                    "end",
                    "id",
                    "page_id",
                    "skip_related_event",
                    "start",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "component_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/components/{component_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "components",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "component_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
            ],
            "key$": "load",
          },
          "patch": {
            "input": "data",
            "name": "patch",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "component_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/pages/{page_id}/components/{component_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "components",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "component_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "patch",
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "component_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/pages/{page_id}/components/{component_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "components",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "component_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "component_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/pages/{page_id}/components/{component_id}/page_access_groups",
                "parts": [
                  "pages",
                  "{page_id}",
                  "components",
                  "{id}",
                  "page_access_groups",
                ],
                "rename": {
                  "param": {
                    "component_id": "id",
                  },
                },
                "select": {
                  "$action": "page_access_group",
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "component_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/pages/{page_id}/components/{component_id}/page_access_users",
                "parts": [
                  "pages",
                  "{page_id}",
                  "components",
                  "{id}",
                  "page_access_users",
                ],
                "rename": {
                  "param": {
                    "component_id": "id",
                  },
                },
                "select": {
                  "$action": "page_access_user",
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
            ],
            "key$": "remove",
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "component_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "PUT",
                "orig": "/pages/{page_id}/components/{component_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "components",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "component_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "update",
          },
        },
        "relations": {
          "ancestors": [
            [
              "page",
            ],
            [
              "page",
              "page_access_group",
            ],
            [
              "page",
              "page_access_user",
            ],
          ],
        },
      },
      "component_group_uptime": {
        "fields": [
          {
            "active": True,
            "name": "id",
            "req": False,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "major_outage",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "name",
            "req": False,
            "type": "`$STRING`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "partial_outage",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "range_end",
            "req": False,
            "type": "`$STRING`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "range_start",
            "req": False,
            "type": "`$STRING`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "related_event",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "uptime_percentage",
            "req": False,
            "type": "`$NUMBER`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "warning",
            "req": False,
            "type": "`$STRING`",
            "index$": 8,
          },
        ],
        "name": "component_group_uptime",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "end",
                      "orig": "end",
                      "reqd": False,
                      "type": "Any",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "skip_related_event",
                      "orig": "skip_related_event",
                      "reqd": False,
                      "type": "`$BOOLEAN`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "start",
                      "orig": "start",
                      "reqd": False,
                      "type": "Any",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/component-groups/{id}/uptime",
                "parts": [
                  "pages",
                  "{page_id}",
                  "component-groups",
                  "{id}",
                  "uptime",
                ],
                "select": {
                  "exist": [
                    "end",
                    "id",
                    "page_id",
                    "skip_related_event",
                    "start",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
        },
        "relations": {
          "ancestors": [
            [
              "page",
            ],
          ],
        },
      },
      "group_component": {
        "fields": [
          {
            "active": True,
            "name": "component",
            "req": False,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "component_group",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "created_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "description",
            "req": False,
            "type": "`$STRING`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "id",
            "req": False,
            "type": "`$STRING`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "name",
            "req": False,
            "type": "`$STRING`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "page_id",
            "req": False,
            "type": "`$STRING`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "position",
            "req": False,
            "type": "`$STRING`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "updated_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 8,
          },
        ],
        "name": "group_component",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/component-groups",
                "parts": [
                  "pages",
                  "{page_id}",
                  "component-groups",
                ],
                "select": {
                  "exist": [
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "create",
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "page",
                      "orig": "page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "per_page",
                      "orig": "per_page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/component-groups",
                "parts": [
                  "pages",
                  "{page_id}",
                  "component-groups",
                ],
                "select": {
                  "exist": [
                    "page",
                    "page_id",
                    "per_page",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "list",
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/component-groups/{id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "component-groups",
                  "{id}",
                ],
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
          "patch": {
            "input": "data",
            "name": "patch",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/pages/{page_id}/component-groups/{id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "component-groups",
                  "{id}",
                ],
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "patch",
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/pages/{page_id}/component-groups/{id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "component-groups",
                  "{id}",
                ],
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "remove",
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "PUT",
                "orig": "/pages/{page_id}/component-groups/{id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "component-groups",
                  "{id}",
                ],
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "update",
          },
        },
        "relations": {
          "ancestors": [
            [
              "page",
            ],
          ],
        },
      },
      "incident": {
        "fields": [
          {
            "active": True,
            "name": "auto_transition_deliver_notifications_at_end",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "auto_transition_deliver_notifications_at_start",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "auto_transition_to_maintenance_state",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "auto_transition_to_operational_state",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "component",
            "req": False,
            "type": "`$ARRAY`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "created_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "id",
            "req": False,
            "type": "`$STRING`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "impact",
            "req": False,
            "type": "`$STRING`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "impact_override",
            "req": False,
            "type": "`$STRING`",
            "index$": 8,
          },
          {
            "active": True,
            "name": "incident",
            "op": {
              "patch": {
                "req": False,
                "type": "`$OBJECT`",
              },
              "update": {
                "req": False,
                "type": "`$OBJECT`",
              },
            },
            "req": True,
            "type": "`$OBJECT`",
            "index$": 9,
          },
          {
            "active": True,
            "name": "incident_impact",
            "req": False,
            "type": "`$ARRAY`",
            "index$": 10,
          },
          {
            "active": True,
            "name": "incident_update",
            "req": False,
            "type": "`$ARRAY`",
            "index$": 11,
          },
          {
            "active": True,
            "name": "metadata",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 12,
          },
          {
            "active": True,
            "name": "monitoring_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 13,
          },
          {
            "active": True,
            "name": "name",
            "req": False,
            "type": "`$STRING`",
            "index$": 14,
          },
          {
            "active": True,
            "name": "page_id",
            "req": False,
            "type": "`$STRING`",
            "index$": 15,
          },
          {
            "active": True,
            "name": "postmortem_body",
            "req": False,
            "type": "`$STRING`",
            "index$": 16,
          },
          {
            "active": True,
            "name": "postmortem_body_last_updated_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 17,
          },
          {
            "active": True,
            "name": "postmortem_ignored",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 18,
          },
          {
            "active": True,
            "name": "postmortem_notified_subscriber",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 19,
          },
          {
            "active": True,
            "name": "postmortem_notified_twitter",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 20,
          },
          {
            "active": True,
            "name": "postmortem_published_at",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 21,
          },
          {
            "active": True,
            "name": "reminder_interval",
            "req": False,
            "type": "`$STRING`",
            "index$": 22,
          },
          {
            "active": True,
            "name": "resolved_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 23,
          },
          {
            "active": True,
            "name": "scheduled_auto_completed",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 24,
          },
          {
            "active": True,
            "name": "scheduled_auto_in_progress",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 25,
          },
          {
            "active": True,
            "name": "scheduled_for",
            "req": False,
            "type": "`$STRING`",
            "index$": 26,
          },
          {
            "active": True,
            "name": "scheduled_remind_prior",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 27,
          },
          {
            "active": True,
            "name": "scheduled_reminded_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 28,
          },
          {
            "active": True,
            "name": "scheduled_until",
            "req": False,
            "type": "`$STRING`",
            "index$": 29,
          },
          {
            "active": True,
            "name": "shortlink",
            "req": False,
            "type": "`$STRING`",
            "index$": 30,
          },
          {
            "active": True,
            "name": "status",
            "req": False,
            "type": "`$STRING`",
            "index$": 31,
          },
          {
            "active": True,
            "name": "updated_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 32,
          },
        ],
        "name": "incident",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/incidents",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                ],
                "select": {
                  "exist": [
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "create",
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "limit",
                      "orig": "limit",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "page",
                      "orig": "page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "q",
                      "orig": "q",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/incidents",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                ],
                "select": {
                  "exist": [
                    "limit",
                    "page",
                    "page_id",
                    "q",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "example": 1,
                      "kind": "query",
                      "name": "page",
                      "orig": "page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "example": 100,
                      "kind": "query",
                      "name": "per_page",
                      "orig": "per_page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/incidents/active_maintenance",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                  "active_maintenance",
                ],
                "select": {
                  "$action": "active_maintenance",
                  "exist": [
                    "page",
                    "page_id",
                    "per_page",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "example": 1,
                      "kind": "query",
                      "name": "page",
                      "orig": "page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "example": 100,
                      "kind": "query",
                      "name": "per_page",
                      "orig": "per_page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/incidents/scheduled",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                  "scheduled",
                ],
                "select": {
                  "$action": "scheduled",
                  "exist": [
                    "page",
                    "page_id",
                    "per_page",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "example": 1,
                      "kind": "query",
                      "name": "page",
                      "orig": "page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "example": 100,
                      "kind": "query",
                      "name": "per_page",
                      "orig": "per_page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/incidents/unresolved",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                  "unresolved",
                ],
                "select": {
                  "$action": "unresolved",
                  "exist": [
                    "page",
                    "page_id",
                    "per_page",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 3,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "example": 1,
                      "kind": "query",
                      "name": "page",
                      "orig": "page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "example": 100,
                      "kind": "query",
                      "name": "per_page",
                      "orig": "per_page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/incidents/upcoming",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                  "upcoming",
                ],
                "select": {
                  "$action": "upcoming",
                  "exist": [
                    "page",
                    "page_id",
                    "per_page",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 4,
              },
            ],
            "key$": "list",
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "incident_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/incidents/{incident_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "incident_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
          "patch": {
            "input": "data",
            "name": "patch",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "incident_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/pages/{page_id}/incidents/{incident_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "incident_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "patch",
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "incident_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/pages/{page_id}/incidents/{incident_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "incident_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "remove",
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "incident_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "PUT",
                "orig": "/pages/{page_id}/incidents/{incident_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "incident_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "update",
          },
        },
        "relations": {
          "ancestors": [
            [
              "page",
            ],
          ],
        },
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
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "incident_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/pages/{page_id}/incidents/{incident_id}/postmortem",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                  "{id}",
                  "postmortem",
                ],
                "rename": {
                  "param": {
                    "incident_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "remove",
          },
        },
        "relations": {
          "ancestors": [
            [
              "page",
            ],
          ],
        },
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
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "incident_id",
                      "orig": "incident_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "subscriber_id",
                      "orig": "subscriber_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 2,
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}/resend_confirmation",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                  "{incident_id}",
                  "subscribers",
                  "{subscriber_id}",
                  "resend_confirmation",
                ],
                "select": {
                  "exist": [
                    "incident_id",
                    "page_id",
                    "subscriber_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "create",
          },
        },
        "relations": {
          "ancestors": [
            [
              "page",
              "incident",
              "subscriber",
            ],
          ],
        },
      },
      "incident_template": {
        "fields": [
          {
            "active": True,
            "name": "body",
            "req": False,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "component",
            "req": False,
            "type": "`$ARRAY`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "group_id",
            "req": False,
            "type": "`$STRING`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "id",
            "req": False,
            "type": "`$STRING`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "name",
            "req": False,
            "type": "`$STRING`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "should_send_notification",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "should_tweet",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "template",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "title",
            "req": False,
            "type": "`$STRING`",
            "index$": 8,
          },
          {
            "active": True,
            "name": "update_status",
            "req": False,
            "type": "`$STRING`",
            "index$": 9,
          },
        ],
        "name": "incident_template",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/incident_templates",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incident_templates",
                ],
                "select": {
                  "exist": [
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "create",
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "example": 1,
                      "kind": "query",
                      "name": "page",
                      "orig": "page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "example": 100,
                      "kind": "query",
                      "name": "per_page",
                      "orig": "per_page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/incident_templates",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incident_templates",
                ],
                "select": {
                  "exist": [
                    "page",
                    "page_id",
                    "per_page",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "list",
          },
        },
        "relations": {
          "ancestors": [
            [
              "page",
            ],
          ],
        },
      },
      "incident_update": {
        "fields": [
          {
            "active": True,
            "name": "affected_component",
            "req": False,
            "type": "`$ARRAY`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "body",
            "req": False,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "created_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "custom_tweet",
            "req": False,
            "type": "`$STRING`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "deliver_notification",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "display_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "id",
            "req": False,
            "type": "`$STRING`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "incident_id",
            "req": False,
            "type": "`$STRING`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "incident_update",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 8,
          },
          {
            "active": True,
            "name": "status",
            "req": False,
            "type": "`$STRING`",
            "index$": 9,
          },
          {
            "active": True,
            "name": "tweet_id",
            "req": False,
            "type": "`$STRING`",
            "index$": 10,
          },
          {
            "active": True,
            "name": "twitter_updated_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 11,
          },
          {
            "active": True,
            "name": "updated_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 12,
          },
          {
            "active": True,
            "name": "wants_twitter_update",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 13,
          },
        ],
        "name": "incident_update",
        "op": {
          "patch": {
            "input": "data",
            "name": "patch",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "incident_update_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "incident_id",
                      "orig": "incident_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                  "{incident_id}",
                  "incident_updates",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "incident_update_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "incident_id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "patch",
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "incident_update_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "incident_id",
                      "orig": "incident_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 2,
                    },
                  ],
                },
                "method": "PUT",
                "orig": "/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                  "{incident_id}",
                  "incident_updates",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "incident_update_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "incident_id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "update",
          },
        },
        "relations": {
          "ancestors": [
            [
              "page",
              "incident",
            ],
          ],
        },
      },
      "metric": {
        "fields": [
          {
            "active": True,
            "name": "backfill_percentage",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "backfilled",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "created_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "data",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "decimal_place",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "display",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "id",
            "req": False,
            "type": "`$STRING`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "last_fetched_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "metric",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 8,
          },
          {
            "active": True,
            "name": "metric_identifier",
            "req": False,
            "type": "`$STRING`",
            "index$": 9,
          },
          {
            "active": True,
            "name": "metrics_provider_id",
            "req": False,
            "type": "`$STRING`",
            "index$": 10,
          },
          {
            "active": True,
            "name": "most_recent_data_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 11,
          },
          {
            "active": True,
            "name": "name",
            "req": False,
            "type": "`$STRING`",
            "index$": 12,
          },
          {
            "active": True,
            "name": "reference_name",
            "req": False,
            "type": "`$STRING`",
            "index$": 13,
          },
          {
            "active": True,
            "name": "suffix",
            "req": False,
            "type": "`$STRING`",
            "index$": 14,
          },
          {
            "active": True,
            "name": "tooltip_description",
            "req": False,
            "type": "`$STRING`",
            "index$": 15,
          },
          {
            "active": True,
            "name": "updated_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 16,
          },
          {
            "active": True,
            "name": "y_axis_hidden",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 17,
          },
          {
            "active": True,
            "name": "y_axis_max",
            "req": False,
            "type": "`$NUMBER`",
            "index$": 18,
          },
          {
            "active": True,
            "name": "y_axis_min",
            "req": False,
            "type": "`$NUMBER`",
            "index$": 19,
          },
        ],
        "name": "metric",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "metric_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/metrics/{metric_id}/data",
                "parts": [
                  "pages",
                  "{page_id}",
                  "metrics",
                  "{id}",
                  "data",
                ],
                "rename": {
                  "param": {
                    "metric_id": "id",
                  },
                },
                "select": {
                  "$action": "data",
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "metrics_provider_id",
                      "orig": "metrics_provider_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics",
                "parts": [
                  "pages",
                  "{page_id}",
                  "metrics_providers",
                  "{metrics_provider_id}",
                  "metrics",
                ],
                "select": {
                  "exist": [
                    "metrics_provider_id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/metrics/data",
                "parts": [
                  "pages",
                  "{page_id}",
                  "metrics",
                  "data",
                ],
                "select": {
                  "$action": "data",
                  "exist": [
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
            ],
            "key$": "create",
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_access_user_id",
                      "orig": "page_access_user_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "page",
                      "orig": "page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "per_page",
                      "orig": "per_page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_users",
                  "{page_access_user_id}",
                  "metrics",
                ],
                "select": {
                  "exist": [
                    "page",
                    "page_access_user_id",
                    "page_id",
                    "per_page",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "list",
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "metrics_provider_id",
                      "orig": "metrics_provider_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "page",
                      "orig": "page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "per_page",
                      "orig": "per_page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics",
                "parts": [
                  "pages",
                  "{page_id}",
                  "metrics_providers",
                  "{metrics_provider_id}",
                  "metrics",
                ],
                "select": {
                  "exist": [
                    "metrics_provider_id",
                    "page",
                    "page_id",
                    "per_page",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "page",
                      "orig": "page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "per_page",
                      "orig": "per_page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/metrics",
                "parts": [
                  "pages",
                  "{page_id}",
                  "metrics",
                ],
                "select": {
                  "exist": [
                    "page",
                    "page_id",
                    "per_page",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "metric_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/metrics/{metric_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "metrics",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "metric_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
            ],
            "key$": "load",
          },
          "patch": {
            "input": "data",
            "name": "patch",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "metric_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/pages/{page_id}/metrics/{metric_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "metrics",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "metric_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "patch",
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "metric_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/pages/{page_id}/metrics/{metric_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "metrics",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "metric_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "metric_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/pages/{page_id}/metrics/{metric_id}/data",
                "parts": [
                  "pages",
                  "{page_id}",
                  "metrics",
                  "{id}",
                  "data",
                ],
                "rename": {
                  "param": {
                    "metric_id": "id",
                  },
                },
                "select": {
                  "$action": "data",
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
            ],
            "key$": "remove",
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "metric_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "PUT",
                "orig": "/pages/{page_id}/metrics/{metric_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "metrics",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "metric_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "update",
          },
        },
        "relations": {
          "ancestors": [
            [
              "page",
            ],
            [
              "page",
              "metrics_provider",
            ],
            [
              "page",
              "page_access_user",
            ],
          ],
        },
      },
      "metrics_provider": {
        "fields": [
          {
            "active": True,
            "name": "created_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "disabled",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "id",
            "req": False,
            "type": "`$STRING`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "last_revalidated_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "metric_base_uri",
            "req": False,
            "type": "`$STRING`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "metrics_provider",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "page_id",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "type",
            "req": False,
            "type": "`$STRING`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "updated_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 8,
          },
        ],
        "name": "metrics_provider",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/metrics_providers",
                "parts": [
                  "pages",
                  "{page_id}",
                  "metrics_providers",
                ],
                "select": {
                  "exist": [
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "create",
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/metrics_providers",
                "parts": [
                  "pages",
                  "{page_id}",
                  "metrics_providers",
                ],
                "select": {
                  "exist": [
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "list",
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "metrics_provider_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "metrics_providers",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "metrics_provider_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
          "patch": {
            "input": "data",
            "name": "patch",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "metrics_provider_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "metrics_providers",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "metrics_provider_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "patch",
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "metrics_provider_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "metrics_providers",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "metrics_provider_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "remove",
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "metrics_provider_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "PUT",
                "orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "metrics_providers",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "metrics_provider_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "update",
          },
        },
        "relations": {
          "ancestors": [
            [
              "page",
            ],
          ],
        },
      },
      "page": {
        "fields": [
          {
            "active": True,
            "name": "activity_score",
            "req": False,
            "type": "`$NUMBER`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "allow_email_subscriber",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "allow_incident_subscriber",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "allow_page_subscriber",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "allow_rss_atom_feed",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "allow_sms_subscriber",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "allow_webhook_subscriber",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "branding",
            "req": False,
            "type": "`$STRING`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "city",
            "req": False,
            "type": "`$STRING`",
            "index$": 8,
          },
          {
            "active": True,
            "name": "country",
            "req": False,
            "type": "`$STRING`",
            "index$": 9,
          },
          {
            "active": True,
            "name": "created_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 10,
          },
          {
            "active": True,
            "name": "css_blue",
            "req": False,
            "type": "`$STRING`",
            "index$": 11,
          },
          {
            "active": True,
            "name": "css_body_background_color",
            "req": False,
            "type": "`$STRING`",
            "index$": 12,
          },
          {
            "active": True,
            "name": "css_border_color",
            "req": False,
            "type": "`$STRING`",
            "index$": 13,
          },
          {
            "active": True,
            "name": "css_font_color",
            "req": False,
            "type": "`$STRING`",
            "index$": 14,
          },
          {
            "active": True,
            "name": "css_graph_color",
            "req": False,
            "type": "`$STRING`",
            "index$": 15,
          },
          {
            "active": True,
            "name": "css_green",
            "req": False,
            "type": "`$STRING`",
            "index$": 16,
          },
          {
            "active": True,
            "name": "css_light_font_color",
            "req": False,
            "type": "`$STRING`",
            "index$": 17,
          },
          {
            "active": True,
            "name": "css_link_color",
            "req": False,
            "type": "`$STRING`",
            "index$": 18,
          },
          {
            "active": True,
            "name": "css_no_data",
            "req": False,
            "type": "`$STRING`",
            "index$": 19,
          },
          {
            "active": True,
            "name": "css_orange",
            "req": False,
            "type": "`$STRING`",
            "index$": 20,
          },
          {
            "active": True,
            "name": "css_red",
            "req": False,
            "type": "`$STRING`",
            "index$": 21,
          },
          {
            "active": True,
            "name": "css_yellow",
            "req": False,
            "type": "`$STRING`",
            "index$": 22,
          },
          {
            "active": True,
            "name": "domain",
            "req": False,
            "type": "`$STRING`",
            "index$": 23,
          },
          {
            "active": True,
            "name": "email_logo",
            "req": False,
            "type": "`$STRING`",
            "index$": 24,
          },
          {
            "active": True,
            "name": "favicon_logo",
            "req": False,
            "type": "`$STRING`",
            "index$": 25,
          },
          {
            "active": True,
            "name": "headline",
            "req": False,
            "type": "`$STRING`",
            "index$": 26,
          },
          {
            "active": True,
            "name": "hero_cover",
            "req": False,
            "type": "`$STRING`",
            "index$": 27,
          },
          {
            "active": True,
            "name": "hidden_from_search",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 28,
          },
          {
            "active": True,
            "name": "id",
            "req": False,
            "type": "`$STRING`",
            "index$": 29,
          },
          {
            "active": True,
            "name": "ip_restriction",
            "req": False,
            "type": "`$STRING`",
            "index$": 30,
          },
          {
            "active": True,
            "name": "name",
            "req": False,
            "type": "`$STRING`",
            "index$": 31,
          },
          {
            "active": True,
            "name": "notifications_email_footer",
            "req": False,
            "type": "`$STRING`",
            "index$": 32,
          },
          {
            "active": True,
            "name": "notifications_from_email",
            "req": False,
            "type": "`$STRING`",
            "index$": 33,
          },
          {
            "active": True,
            "name": "page",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 34,
          },
          {
            "active": True,
            "name": "page_description",
            "req": False,
            "type": "`$STRING`",
            "index$": 35,
          },
          {
            "active": True,
            "name": "state",
            "req": False,
            "type": "`$STRING`",
            "index$": 36,
          },
          {
            "active": True,
            "name": "subdomain",
            "req": False,
            "type": "`$STRING`",
            "index$": 37,
          },
          {
            "active": True,
            "name": "support_url",
            "req": False,
            "type": "`$STRING`",
            "index$": 38,
          },
          {
            "active": True,
            "name": "time_zone",
            "req": False,
            "type": "`$STRING`",
            "index$": 39,
          },
          {
            "active": True,
            "name": "transactional_logo",
            "req": False,
            "type": "`$STRING`",
            "index$": 40,
          },
          {
            "active": True,
            "name": "twitter_logo",
            "req": False,
            "type": "`$STRING`",
            "index$": 41,
          },
          {
            "active": True,
            "name": "twitter_username",
            "req": False,
            "type": "`$STRING`",
            "index$": 42,
          },
          {
            "active": True,
            "name": "updated_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 43,
          },
          {
            "active": True,
            "name": "url",
            "req": False,
            "type": "`$STRING`",
            "index$": 44,
          },
          {
            "active": True,
            "name": "viewers_must_be_team_member",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 45,
          },
        ],
        "name": "page",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {},
                "method": "GET",
                "orig": "/pages",
                "parts": [
                  "pages",
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "list",
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}",
                "parts": [
                  "pages",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "page_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
          "patch": {
            "input": "data",
            "name": "patch",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/pages/{page_id}",
                "parts": [
                  "pages",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "page_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "patch",
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "PUT",
                "orig": "/pages/{page_id}",
                "parts": [
                  "pages",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "page_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "update",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "page_access_group": {
        "fields": [
          {
            "active": True,
            "name": "component_id",
            "op": {
              "create": {
                "req": True,
                "type": "`$ARRAY`",
              },
            },
            "req": False,
            "type": "`$ARRAY`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "created_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "external_identifier",
            "req": False,
            "type": "`$STRING`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "id",
            "req": False,
            "type": "`$STRING`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "metric_id",
            "req": False,
            "type": "`$ARRAY`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "name",
            "req": False,
            "type": "`$STRING`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "page_access_group",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "page_access_user_id",
            "req": False,
            "type": "`$ARRAY`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "page_id",
            "req": False,
            "type": "`$STRING`",
            "index$": 8,
          },
          {
            "active": True,
            "name": "updated_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 9,
          },
        ],
        "name": "page_access_group",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_group_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_groups",
                  "{id}",
                  "components",
                ],
                "rename": {
                  "param": {
                    "page_access_group_id": "id",
                  },
                },
                "select": {
                  "$action": "component",
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/page_access_groups",
                "parts": [
                  "pages",
                  "{id}",
                  "page_access_groups",
                ],
                "rename": {
                  "param": {
                    "page_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
            ],
            "key$": "create",
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "page",
                      "orig": "page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "per_page",
                      "orig": "per_page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/page_access_groups",
                "parts": [
                  "pages",
                  "{id}",
                  "page_access_groups",
                ],
                "rename": {
                  "param": {
                    "page_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page",
                    "per_page",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "list",
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_group_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_groups",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "page_access_group_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
          "patch": {
            "input": "data",
            "name": "patch",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_group_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_groups",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "page_access_group_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_group_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_groups",
                  "{id}",
                  "components",
                ],
                "rename": {
                  "param": {
                    "page_access_group_id": "id",
                  },
                },
                "select": {
                  "$action": "component",
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
            ],
            "key$": "patch",
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "component_id",
                      "orig": "component_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_group_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 2,
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components/{component_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_groups",
                  "{id}",
                  "components",
                  "{component_id}",
                ],
                "rename": {
                  "param": {
                    "page_access_group_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "component_id",
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_group_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_groups",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "page_access_group_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_group_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_groups",
                  "{id}",
                  "components",
                ],
                "rename": {
                  "param": {
                    "page_access_group_id": "id",
                  },
                },
                "select": {
                  "$action": "component",
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
            ],
            "key$": "remove",
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_group_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "PUT",
                "orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_groups",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "page_access_group_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_group_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PUT",
                "orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_groups",
                  "{id}",
                  "components",
                ],
                "rename": {
                  "param": {
                    "page_access_group_id": "id",
                  },
                },
                "select": {
                  "$action": "component",
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
            ],
            "key$": "update",
          },
        },
        "relations": {
          "ancestors": [
            [
              "page",
            ],
            [
              "page",
              "component",
            ],
          ],
        },
      },
      "page_access_user": {
        "fields": [
          {
            "active": True,
            "name": "component_id",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "created_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "email",
            "req": False,
            "type": "`$STRING`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "external_login",
            "req": False,
            "type": "`$STRING`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "id",
            "req": False,
            "type": "`$STRING`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "metric_id",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "page_access_group_id",
            "req": False,
            "type": "`$STRING`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "page_access_user",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "page_id",
            "req": False,
            "type": "`$STRING`",
            "index$": 8,
          },
          {
            "active": True,
            "name": "updated_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 9,
          },
        ],
        "name": "page_access_user",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_user_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_users",
                  "{id}",
                  "components",
                ],
                "rename": {
                  "param": {
                    "page_access_user_id": "id",
                  },
                },
                "select": {
                  "$action": "component",
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_user_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_users",
                  "{id}",
                  "metrics",
                ],
                "rename": {
                  "param": {
                    "page_access_user_id": "id",
                  },
                },
                "select": {
                  "$action": "metric",
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/page_access_users",
                "parts": [
                  "pages",
                  "{id}",
                  "page_access_users",
                ],
                "rename": {
                  "param": {
                    "page_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
            ],
            "key$": "create",
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "email",
                      "orig": "email",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "page",
                      "orig": "page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "per_page",
                      "orig": "per_page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/page_access_users",
                "parts": [
                  "pages",
                  "{id}",
                  "page_access_users",
                ],
                "rename": {
                  "param": {
                    "page_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "email",
                    "id",
                    "page",
                    "per_page",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "list",
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_user_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_users",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "page_access_user_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
          "patch": {
            "input": "data",
            "name": "patch",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_user_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_users",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "page_access_user_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_user_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_users",
                  "{id}",
                  "components",
                ],
                "rename": {
                  "param": {
                    "page_access_user_id": "id",
                  },
                },
                "select": {
                  "$action": "component",
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_user_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_users",
                  "{id}",
                  "metrics",
                ],
                "rename": {
                  "param": {
                    "page_access_user_id": "id",
                  },
                },
                "select": {
                  "$action": "metric",
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
            ],
            "key$": "patch",
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "component_id",
                      "orig": "component_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_user_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 2,
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components/{component_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_users",
                  "{id}",
                  "components",
                  "{component_id}",
                ],
                "rename": {
                  "param": {
                    "page_access_user_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "component_id",
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_user_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "metric_id",
                      "orig": "metric_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 2,
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics/{metric_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_users",
                  "{id}",
                  "metrics",
                  "{metric_id}",
                ],
                "rename": {
                  "param": {
                    "page_access_user_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "metric_id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_user_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_users",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "page_access_user_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_user_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_users",
                  "{id}",
                  "components",
                ],
                "rename": {
                  "param": {
                    "page_access_user_id": "id",
                  },
                },
                "select": {
                  "$action": "component",
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 3,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_user_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_users",
                  "{id}",
                  "metrics",
                ],
                "rename": {
                  "param": {
                    "page_access_user_id": "id",
                  },
                },
                "select": {
                  "$action": "metric",
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 4,
              },
            ],
            "key$": "remove",
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_user_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "PUT",
                "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_users",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "page_access_user_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_user_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PUT",
                "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_users",
                  "{id}",
                  "components",
                ],
                "rename": {
                  "param": {
                    "page_access_user_id": "id",
                  },
                },
                "select": {
                  "$action": "component",
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "page_access_user_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PUT",
                "orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
                "parts": [
                  "pages",
                  "{page_id}",
                  "page_access_users",
                  "{id}",
                  "metrics",
                ],
                "rename": {
                  "param": {
                    "page_access_user_id": "id",
                  },
                },
                "select": {
                  "$action": "metric",
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
            ],
            "key$": "update",
          },
        },
        "relations": {
          "ancestors": [
            [
              "page",
            ],
            [
              "page",
              "component",
            ],
            [
              "page",
              "metric",
            ],
          ],
        },
      },
      "permission": {
        "fields": [
          {
            "active": True,
            "name": "data",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "page",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 1,
          },
        ],
        "name": "permission",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "user_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "organization_id",
                      "orig": "organization_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "GET",
                "orig": "/organizations/{organization_id}/permissions/{user_id}",
                "parts": [
                  "organizations",
                  "{organization_id}",
                  "permissions",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "user_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "organization_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "user_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "organization_id",
                      "orig": "organization_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "PUT",
                "orig": "/organizations/{organization_id}/permissions/{user_id}",
                "parts": [
                  "organizations",
                  "{organization_id}",
                  "permissions",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "user_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "organization_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "update",
          },
        },
        "relations": {
          "ancestors": [
            [
              "organization",
            ],
          ],
        },
      },
      "postmortem": {
        "fields": [
          {
            "active": True,
            "name": "body",
            "req": False,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "body_draft",
            "req": False,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "body_draft_updated_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "body_updated_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "created_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "custom_tweet",
            "req": False,
            "type": "`$STRING`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "notify_subscriber",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "notify_twitter",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "postmortem",
            "op": {
              "update": {
                "req": False,
                "type": "`$OBJECT`",
              },
            },
            "req": True,
            "type": "`$OBJECT`",
            "index$": 8,
          },
          {
            "active": True,
            "name": "preview_key",
            "req": False,
            "type": "`$STRING`",
            "index$": 9,
          },
          {
            "active": True,
            "name": "published_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 10,
          },
          {
            "active": True,
            "name": "updated_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 11,
          },
        ],
        "name": "postmortem",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "incident_id",
                      "orig": "incident_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/incidents/{incident_id}/postmortem",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                  "{incident_id}",
                  "postmortem",
                ],
                "select": {
                  "exist": [
                    "incident_id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "incident_id",
                      "orig": "incident_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "PUT",
                "orig": "/pages/{page_id}/incidents/{incident_id}/postmortem",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                  "{incident_id}",
                  "postmortem",
                ],
                "select": {
                  "exist": [
                    "incident_id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "incident_id",
                      "orig": "incident_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PUT",
                "orig": "/pages/{page_id}/incidents/{incident_id}/postmortem/publish",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                  "{incident_id}",
                  "postmortem",
                  "publish",
                ],
                "select": {
                  "$action": "publish",
                  "exist": [
                    "incident_id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "incident_id",
                      "orig": "incident_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PUT",
                "orig": "/pages/{page_id}/incidents/{incident_id}/postmortem/revert",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                  "{incident_id}",
                  "postmortem",
                  "revert",
                ],
                "select": {
                  "$action": "revert",
                  "exist": [
                    "incident_id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
            ],
            "key$": "update",
          },
        },
        "relations": {
          "ancestors": [
            [
              "page",
              "incident",
            ],
          ],
        },
      },
      "status_embed_config": {
        "fields": [
          {
            "active": True,
            "name": "incident_background_color",
            "req": False,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "incident_text_color",
            "req": False,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "maintenance_background_color",
            "req": False,
            "type": "`$STRING`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "maintenance_text_color",
            "req": False,
            "type": "`$STRING`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "page_id",
            "req": False,
            "type": "`$STRING`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "position",
            "req": False,
            "type": "`$STRING`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "status_embed_config",
            "req": False,
            "type": "`$OBJECT`",
            "index$": 6,
          },
        ],
        "name": "status_embed_config",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/status_embed_config",
                "parts": [
                  "pages",
                  "{page_id}",
                  "status_embed_config",
                ],
                "select": {
                  "exist": [
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
          "patch": {
            "input": "data",
            "name": "patch",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/pages/{page_id}/status_embed_config",
                "parts": [
                  "pages",
                  "{page_id}",
                  "status_embed_config",
                ],
                "select": {
                  "exist": [
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "patch",
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "PUT",
                "orig": "/pages/{page_id}/status_embed_config",
                "parts": [
                  "pages",
                  "{page_id}",
                  "status_embed_config",
                ],
                "select": {
                  "exist": [
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "update",
          },
        },
        "relations": {
          "ancestors": [
            [
              "page",
            ],
          ],
        },
      },
      "subscriber": {
        "fields": [
          {
            "active": True,
            "name": "component",
            "req": False,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "component_id",
            "req": False,
            "type": "`$ARRAY`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "created_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "display_phone_number",
            "req": False,
            "type": "`$STRING`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "email",
            "req": False,
            "type": "`$STRING`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "endpoint",
            "req": False,
            "type": "`$STRING`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "id",
            "req": False,
            "type": "`$STRING`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "integration_partner",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "mode",
            "req": False,
            "type": "`$STRING`",
            "index$": 8,
          },
          {
            "active": True,
            "name": "obfuscated_channel_name",
            "req": False,
            "type": "`$STRING`",
            "index$": 9,
          },
          {
            "active": True,
            "name": "page_access_user_id",
            "req": False,
            "type": "`$STRING`",
            "index$": 10,
          },
          {
            "active": True,
            "name": "phone_country",
            "req": False,
            "type": "`$STRING`",
            "index$": 11,
          },
          {
            "active": True,
            "name": "phone_number",
            "req": False,
            "type": "`$STRING`",
            "index$": 12,
          },
          {
            "active": True,
            "name": "purge_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 13,
          },
          {
            "active": True,
            "name": "quarantined_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 14,
          },
          {
            "active": True,
            "name": "skip_confirmation_notification",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 15,
          },
          {
            "active": True,
            "name": "skip_unsubscription_notification",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 16,
          },
          {
            "active": True,
            "name": "slack",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 17,
          },
          {
            "active": True,
            "name": "sms",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 18,
          },
          {
            "active": True,
            "name": "state",
            "req": False,
            "type": "`$STRING`",
            "index$": 19,
          },
          {
            "active": True,
            "name": "subscriber",
            "op": {
              "create": {
                "req": True,
                "type": "`$STRING`",
              },
            },
            "req": False,
            "type": "`$OBJECT`",
            "index$": 20,
          },
          {
            "active": True,
            "name": "team",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 21,
          },
          {
            "active": True,
            "name": "type",
            "req": False,
            "type": "`$STRING`",
            "index$": 22,
          },
          {
            "active": True,
            "name": "webhook",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 23,
          },
          {
            "active": True,
            "name": "workspace_name",
            "req": False,
            "type": "`$STRING`",
            "index$": 24,
          },
        ],
        "name": "subscriber",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "subscriber_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/subscribers/{subscriber_id}/resend_confirmation",
                "parts": [
                  "pages",
                  "{page_id}",
                  "subscribers",
                  "{id}",
                  "resend_confirmation",
                ],
                "rename": {
                  "param": {
                    "subscriber_id": "id",
                  },
                },
                "select": {
                  "$action": "resend_confirmation",
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "incident_id",
                      "orig": "incident_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/incidents/{incident_id}/subscribers",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                  "{incident_id}",
                  "subscribers",
                ],
                "select": {
                  "exist": [
                    "incident_id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/subscribers",
                "parts": [
                  "pages",
                  "{page_id}",
                  "subscribers",
                ],
                "select": {
                  "exist": [
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/subscribers/reactivate",
                "parts": [
                  "pages",
                  "{page_id}",
                  "subscribers",
                  "reactivate",
                ],
                "select": {
                  "$action": "reactivate",
                  "exist": [
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 3,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/subscribers/resend_confirmation",
                "parts": [
                  "pages",
                  "{page_id}",
                  "subscribers",
                  "resend_confirmation",
                ],
                "select": {
                  "$action": "resend_confirmation",
                  "exist": [
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 4,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "POST",
                "orig": "/pages/{page_id}/subscribers/unsubscribe",
                "parts": [
                  "pages",
                  "{page_id}",
                  "subscribers",
                  "unsubscribe",
                ],
                "select": {
                  "$action": "unsubscribe",
                  "exist": [
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 5,
              },
            ],
            "key$": "create",
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "limit",
                      "orig": "limit",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "example": 0,
                      "kind": "query",
                      "name": "page",
                      "orig": "page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "q",
                      "orig": "q",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "example": "asc",
                      "kind": "query",
                      "name": "sort_direction",
                      "orig": "sort_direction",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "example": "primary",
                      "kind": "query",
                      "name": "sort_field",
                      "orig": "sort_field",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "example": "active",
                      "kind": "query",
                      "name": "state",
                      "orig": "state",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "type",
                      "orig": "type",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/subscribers",
                "parts": [
                  "pages",
                  "{page_id}",
                  "subscribers",
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
                    "type",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "incident_id",
                      "orig": "incident_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "page",
                      "orig": "page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "per_page",
                      "orig": "per_page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/incidents/{incident_id}/subscribers",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                  "{incident_id}",
                  "subscribers",
                ],
                "select": {
                  "exist": [
                    "incident_id",
                    "page",
                    "page_id",
                    "per_page",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "page",
                      "orig": "page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "per_page",
                      "orig": "per_page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/subscribers/unsubscribed",
                "parts": [
                  "pages",
                  "{page_id}",
                  "subscribers",
                  "unsubscribed",
                ],
                "select": {
                  "$action": "unsubscribed",
                  "exist": [
                    "page",
                    "page_id",
                    "per_page",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
            ],
            "key$": "list",
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "subscriber_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "incident_id",
                      "orig": "incident_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 2,
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                  "{incident_id}",
                  "subscribers",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "subscriber_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "incident_id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "example": "active",
                      "kind": "query",
                      "name": "state",
                      "orig": "state",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "type",
                      "orig": "type",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/subscribers/count",
                "parts": [
                  "pages",
                  "{page_id}",
                  "subscribers",
                  "count",
                ],
                "select": {
                  "$action": "count",
                  "exist": [
                    "page_id",
                    "state",
                    "type",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "subscriber_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/subscribers/{subscriber_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "subscribers",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "subscriber_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/pages/{page_id}/subscribers/histogram_by_state",
                "parts": [
                  "pages",
                  "{page_id}",
                  "subscribers",
                  "histogram_by_state",
                ],
                "select": {
                  "$action": "histogram_by_state",
                  "exist": [
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 3,
              },
            ],
            "key$": "load",
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "subscriber_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "incident_id",
                      "orig": "incident_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 2,
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "incidents",
                  "{incident_id}",
                  "subscribers",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "subscriber_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "incident_id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "subscriber_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "skip_unsubscription_notification",
                      "orig": "skip_unsubscription_notification",
                      "reqd": False,
                      "type": "`$BOOLEAN`",
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/pages/{page_id}/subscribers/{subscriber_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "subscribers",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "subscriber_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                    "skip_unsubscription_notification",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
            ],
            "key$": "remove",
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "subscriber_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "page_id",
                      "orig": "page_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/pages/{page_id}/subscribers/{subscriber_id}",
                "parts": [
                  "pages",
                  "{page_id}",
                  "subscribers",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "subscriber_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "page_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "update",
          },
        },
        "relations": {
          "ancestors": [
            [
              "page",
            ],
            [
              "page",
              "incident",
            ],
          ],
        },
      },
      "user": {
        "fields": [
          {
            "active": True,
            "name": "created_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "email",
            "req": False,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "first_name",
            "req": False,
            "type": "`$STRING`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "id",
            "req": False,
            "type": "`$STRING`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "last_name",
            "req": False,
            "type": "`$STRING`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "organization_id",
            "req": False,
            "type": "`$STRING`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "updated_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "user",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 7,
          },
        ],
        "name": "user",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "organization_id",
                      "orig": "organization_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "POST",
                "orig": "/organizations/{organization_id}/users",
                "parts": [
                  "organizations",
                  "{organization_id}",
                  "users",
                ],
                "select": {
                  "exist": [
                    "organization_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "create",
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "organization_id",
                      "orig": "organization_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "page",
                      "orig": "page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "per_page",
                      "orig": "per_page",
                      "reqd": False,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/organizations/{organization_id}/users",
                "parts": [
                  "organizations",
                  "{organization_id}",
                  "users",
                ],
                "select": {
                  "exist": [
                    "organization_id",
                    "page",
                    "per_page",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "list",
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "user_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                    {
                      "active": True,
                      "kind": "param",
                      "name": "organization_id",
                      "orig": "organization_id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 1,
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/organizations/{organization_id}/users/{user_id}",
                "parts": [
                  "organizations",
                  "{organization_id}",
                  "users",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "user_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                    "organization_id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "remove",
          },
        },
        "relations": {
          "ancestors": [
            [
              "organization",
            ],
          ],
        },
      },
    },
    }
