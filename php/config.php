<?php
declare(strict_types=1);

// Statuspage SDK configuration

class StatuspageConfig
{
    /** @var array<string,mixed>|null */
    private static ?array $shared_config = null;

    /**
     * Return the process-wide config, built once on first use. The SDK reads
     * the config on every request and never writes to it, so one instance is
     * shared by every client rather than rebuilt per client.
     *
     * PHP arrays are copy-on-write, so callers that do mutate the result get
     * their own copy and cannot disturb the shared one.
     */
    public static function shared_config(): array
    {
        if (self::$shared_config === null) {
            self::$shared_config = self::make_config();
        }
        return self::$shared_config;
    }

    /**
     * Build a fresh, fully materialised config array. Every call rebuilds the
     * whole structure, so prefer shared_config unless you need a private copy.
     */
    public static function make_config(): array
    {
        return [
            "main" => [
                "name" => "Statuspage",
                "slug" => "statuspage",
                "version" => "0.0.2",
                "target" => "php",
            ],
            "feature" => [
                "debug" => [
          'options' => [
            'active' => false,
            'max' => 100,
            'redact' => [
              'authorization',
              'cookie',
              'set-cookie',
              'api-key',
              'apikey',
              'x-api-key',
              'idempotency-key',
            ],
          ],
          'optspec' => [
            'now' => '`$FUNCTION`',
            'onEntry' => '`$FUNCTION`',
          ],
          'strict' => false,
          'transport' => 'none',
        ],
                "idempotency" => [
          'options' => [
            'active' => false,
            'header' => 'Idempotency-Key',
            'methods' => [
              'POST',
              'PUT',
              'PATCH',
              'DELETE',
            ],
            'ops' => [
              'create',
              'update',
              'remove',
            ],
          ],
          'optspec' => [
            'keygen' => '`$FUNCTION`',
          ],
          'strict' => false,
          'transport' => 'none',
        ],
                "metrics" => [
          'options' => [
            'active' => false,
          ],
          'optspec' => [
            'now' => '`$FUNCTION`',
          ],
          'strict' => false,
          'transport' => 'none',
        ],
                "paging" => [
          'options' => [
            'active' => false,
            'afterVar' => 'after',
            'cursorParam' => 'cursor',
            'firstVar' => 'first',
            'limitParam' => 'limit',
            'pageParam' => 'page',
            'startPage' => 1,
          ],
          'optspec' => [
            'limit' => '`$NUMBER`',
            'ops' => '`$LIST`',
          ],
          'strict' => false,
          'transport' => 'none',
        ],
                "ratelimit" => [
          'options' => [
            'active' => false,
            'burst' => 5,
            'rate' => 5,
          ],
          'optspec' => [
            'now' => '`$FUNCTION`',
            'sleep' => '`$FUNCTION`',
          ],
          'strict' => false,
          'transport' => 'wrap',
        ],
                "retry" => [
          'options' => [
            'active' => false,
            'factor' => 2,
            'maxDelay' => 2000,
            'minDelay' => 50,
            'retries' => 2,
            'statuses' => [
              408,
              425,
              429,
              500,
              502,
              503,
              504,
            ],
          ],
          'optspec' => [
            'jitter' => '`$BOOLEAN`',
            'sleep' => '`$FUNCTION`',
          ],
          'strict' => false,
          'transport' => 'wrap',
        ],
                "test" => [
          'options' => [
            'active' => false,
          ],
          'optspec' => [
            'entity' => '`$MAP`',
            'net' => '`$MAP`',
          ],
          'strict' => false,
          'transport' => 'base',
        ],
                "timeout" => [
          'options' => [
            'active' => false,
            'ms' => 30000,
          ],
          'optspec' => [
            'clearTimer' => '`$FUNCTION`',
            'setTimer' => '`$FUNCTION`',
          ],
          'strict' => false,
          'transport' => 'wrap',
        ],
            ],
            "options" => [
                "base" => "https://api.statuspage.io/v1",
                "auth" => [
                    "prefix" => "OAuth",
                ],
                "headers" => [
          'content-type' => 'application/json',
        ],
                "entity" => [
                    "component" => [],
                    "component_group_uptime" => [],
                    "group_component" => [],
                    "incident" => [],
                    "incident_postmortem" => [],
                    "incident_subscriber" => [],
                    "incident_template" => [],
                    "incident_update" => [],
                    "metric" => [],
                    "metrics_provider" => [],
                    "page" => [],
                    "page_access_group" => [],
                    "page_access_user" => [],
                    "permission" => [],
                    "postmortem" => [],
                    "status_embed_config" => [],
                    "subscriber" => [],
                    "user" => [],
                ],
            ],
            "entity" => [
        'component' => [
          'fields' => [
            [
              'name' => 'automation_email',
              'short' => 'Requires a special feature flag to be enabled',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'component',
              'type' => '`$OBJECT`',
            ],
            [
              'format' => 'date-time',
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'description',
              'short' => 'More detailed description for component',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'group',
              'short' => 'Is this component a group',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'group_id',
              'short' => 'Component Group identifier',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'short' => 'Identifier for component',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'name',
              'short' => 'Display name for component',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'only_show_if_degraded',
              'short' => 'Requires a special feature flag to be enabled',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'page_id',
              'short' => 'Page identifier',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'int32',
              'name' => 'position',
              'short' => 'Order the component will appear on the page',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'showcase',
              'short' => 'Should this component be showcased',
              'type' => '`$BOOLEAN`',
            ],
            [
              'format' => 'date',
              'name' => 'start_date',
              'short' => 'The date this component started being used',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'status',
              'short' => 'Status of component',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'updated_at',
              'type' => '`$STRING`',
            ],
          ],
          'id' => [
            'field' => 'id',
            'name' => 'id',
          ],
          'name' => 'component',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/components/{component_id}/page_access_groups',
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'page_access_groups',
                    ],
                  ],
                  'select' => [
                    '$action' => 'page_access_group',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                    'page_access_groups',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/components/{component_id}/page_access_users',
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'page_access_users',
                    ],
                  ],
                  'select' => [
                    '$action' => 'page_access_user',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                    'page_access_users',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/components',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'component' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_access_group_id',
                        'orig' => 'page_access_group_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}/components',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_groups',
                    ],
                    [
                      'var' => 'page_access_group_id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page',
                      'page_access_group_id',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{page_access_group_id}',
                    'components',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_access_user_id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/components',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_users',
                    ],
                    [
                      'var' => 'page_access_user_id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page',
                      'page_access_user_id',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{page_access_user_id}',
                    'components',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/components',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'end',
                        'orig' => 'end',
                        'type' => 'Any',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'start',
                        'orig' => 'start',
                        'type' => 'Any',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/components/{component_id}/uptime',
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'uptime',
                    ],
                  ],
                  'select' => [
                    '$action' => 'uptime',
                    'exist' => [
                      'end',
                      'id',
                      'page_id',
                      'start',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.related_events`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                    'uptime',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/components/{component_id}',
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/components/{component_id}',
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'component' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/components/{component_id}',
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/components/{component_id}/page_access_groups',
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'page_access_groups',
                    ],
                  ],
                  'select' => [
                    '$action' => 'page_access_group',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                    'page_access_groups',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/components/{component_id}/page_access_users',
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'page_access_users',
                    ],
                  ],
                  'select' => [
                    '$action' => 'page_access_user',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                    'page_access_users',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/components/{component_id}',
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'component' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
              [
                'page',
                'page_access_group',
              ],
              [
                'page',
                'page_access_user',
              ],
            ],
          ],
        ],
        'component_group_uptime' => [
          'fields' => [
            [
              'name' => 'component_id',
              'short' => 'Component identifier',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'incidents',
              'short' => 'Related incidents',
              'type' => '`$OBJECT`',
            ],
          ],
          'id' => [
            'field' => 'id',
            'name' => 'id',
          ],
          'name' => 'component_group_uptime',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'end',
                        'orig' => 'end',
                        'type' => 'Any',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'start',
                        'orig' => 'start',
                        'type' => 'Any',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/component-groups/{id}/uptime',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'component-groups',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'uptime',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'end',
                      'id',
                      'page_id',
                      'start',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.related_events`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                    '{id}',
                    'uptime',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
            ],
          ],
        ],
        'group_component' => [
          'fields' => [
            [
              'name' => 'component_group',
              'req' => true,
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'components',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'description',
              'short' => 'Description of the component group.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'short' => 'Component Group Identifier',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'page_id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'position',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'updated_at',
              'type' => '`$STRING`',
            ],
          ],
          'id' => [
            'field' => 'id',
            'name' => 'id',
          ],
          'name' => 'group_component',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/component-groups',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'component-groups',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/component-groups',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'component-groups',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/component-groups/{id}',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'component-groups',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                    '{id}',
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/component-groups/{id}',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'component-groups',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                    '{id}',
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/component-groups/{id}',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'component-groups',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                    '{id}',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/component-groups/{id}',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'component-groups',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                    '{id}',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
            ],
          ],
        ],
        'incident' => [
          'fields' => [
            [
              'name' => 'auto_transition_deliver_notifications_at_end',
              'short' => 'Controls whether send notification when scheduled maintenances auto transition to completed.',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'auto_transition_deliver_notifications_at_start',
              'short' => 'Controls whether send notification when scheduled maintenances auto transition to started.',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'auto_transition_to_maintenance_state',
              'short' => 'Controls whether change components status to under_maintenance once scheduled maintenance is in progress.',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'auto_transition_to_operational_state',
              'short' => 'Controls whether change components status to operational once scheduled maintenance completes.',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'components',
              'short' => 'Incident components',
              'type' => '`$ARRAY`',
            ],
            [
              'format' => 'date-time',
              'name' => 'created_at',
              'short' => 'The timestamp when the incident was created at.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'short' => 'Incident Identifier',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'impact',
              'short' => 'The impact of the incident.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'impact_override',
              'short' => 'value to override calculated impact value',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'incident',
              'op' => [
                'patch' => [
                  'type' => '`$OBJECT`',
                ],
                'update' => [
                  'type' => '`$OBJECT`',
                ],
              ],
              'req' => true,
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'incident_updates',
              'short' => 'The incident updates for incident.',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'metadata',
              'short' => 'Metadata attached to the incident.',
              'type' => '`$OBJECT`',
            ],
            [
              'format' => 'date-time',
              'name' => 'monitoring_at',
              'short' => 'The timestamp when incident entered monitoring state.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'name',
              'short' => 'Incident Name.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'page_id',
              'short' => 'Incident Page Identifier',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'postmortem_body',
              'short' => 'Body of the Postmortem.',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'postmortem_body_last_updated_at',
              'short' => 'The timestamp when the incident postmortem body was last updated at.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'postmortem_ignored',
              'short' => 'Controls whether the incident will have postmortem.',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'postmortem_notified_subscribers',
              'short' => 'Indicates whether subscribers are already notificed about postmortem.',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'postmortem_notified_twitter',
              'short' => 'Controls whether to decide if notify postmortem on twitter.',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'postmortem_published_at',
              'short' => 'The timestamp when the postmortem was published.',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'reminder_intervals',
              'short' => 'Custom reminder intervals for unresolved/open incidents.',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'resolved_at',
              'short' => 'The timestamp when incident was resolved.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'scheduled_auto_completed',
              'short' => 'Controls whether the incident is scheduled to automatically change to complete.',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'scheduled_auto_in_progress',
              'short' => 'Controls whether the incident is scheduled to automatically change to in progress.',
              'type' => '`$BOOLEAN`',
            ],
            [
              'format' => 'date-time',
              'name' => 'scheduled_for',
              'short' => 'The timestamp the incident is scheduled for.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'scheduled_remind_prior',
              'short' => 'Controls whether to remind subscribers prior to scheduled incidents.',
              'type' => '`$BOOLEAN`',
            ],
            [
              'format' => 'date-time',
              'name' => 'scheduled_reminded_at',
              'short' => 'The timestamp when the scheduled incident reminder was sent at.',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'scheduled_until',
              'short' => 'The timestamp the incident is scheduled until.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'shortlink',
              'short' => 'Incident Shortlink.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'status',
              'short' => 'The incident status.',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'updated_at',
              'short' => 'The timestamp when the incident was updated at.',
              'type' => '`$STRING`',
            ],
          ],
          'id' => [
            'field' => 'id',
            'name' => 'id',
          ],
          'name' => 'incident',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/incidents',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'incident' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'limit',
                        'orig' => 'limit',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'q',
                        'orig' => 'q',
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incidents',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'limit',
                      'page',
                      'page_id',
                      'q',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'example' => 1,
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'example' => 100,
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incidents/active_maintenance',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                    [
                      'lit' => 'active_maintenance',
                    ],
                  ],
                  'select' => [
                    '$action' => 'active_maintenance',
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    'active_maintenance',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'example' => 1,
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'example' => 100,
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incidents/scheduled',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                    [
                      'lit' => 'scheduled',
                    ],
                  ],
                  'select' => [
                    '$action' => 'scheduled',
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    'scheduled',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'example' => 1,
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'example' => 100,
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incidents/unresolved',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                    [
                      'lit' => 'unresolved',
                    ],
                  ],
                  'select' => [
                    '$action' => 'unresolved',
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    'unresolved',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'example' => 1,
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'example' => 100,
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incidents/upcoming',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                    [
                      'lit' => 'upcoming',
                    ],
                  ],
                  'select' => [
                    '$action' => 'upcoming',
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    'upcoming',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}',
                  'rename' => [
                    'param' => [
                      'incident_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{id}',
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}',
                  'rename' => [
                    'param' => [
                      'incident_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'incident' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{id}',
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}',
                  'rename' => [
                    'param' => [
                      'incident_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{id}',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}',
                  'rename' => [
                    'param' => [
                      'incident_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'incident' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{id}',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
            ],
          ],
        ],
        'incident_postmortem' => [
          'fields' => [
            [
              'name' => 'id',
              'type' => '`$STRING`',
            ],
          ],
          'id' => [
            'field' => 'id',
            'name' => 'id',
          ],
          'name' => 'incident_postmortem',
          'op' => [
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/postmortem',
                  'rename' => [
                    'param' => [
                      'incident_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'postmortem',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{id}',
                    'postmortem',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
            ],
          ],
        ],
        'incident_subscriber' => [
          'fields' => [],
          'name' => 'incident_subscriber',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'subscriber_id',
                        'orig' => 'subscriber_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}/resend_confirmation',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                    [
                      'var' => 'incident_id',
                    ],
                    [
                      'lit' => 'subscribers',
                    ],
                    [
                      'var' => 'subscriber_id',
                    ],
                    [
                      'lit' => 'resend_confirmation',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'incident_id',
                      'page_id',
                      'subscriber_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'subscribers',
                    '{subscriber_id}',
                    'resend_confirmation',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
                'incident',
                'subscriber',
              ],
            ],
          ],
        ],
        'incident_template' => [
          'fields' => [
            [
              'name' => 'body',
              'short' => 'Body of the incident or maintenance update to be applied when selecting this template',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'components',
              'short' => 'Affected components',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'group_id',
              'short' => 'Identifier of Template Group this template belongs to',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'short' => 'Incident Template Identifier',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'name',
              'short' => 'Name of the template, as shown in the list on the "Templates" tab of the "Incidents" page',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'should_send_notifications',
              'short' => 'Whether the "deliver notifications" checkbox should be selected when selecting this template',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'should_tweet',
              'short' => 'Whether the "tweet update" checkbox should be selected when selecting this template',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'template',
              'req' => true,
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'title',
              'short' => 'Title to be applied to the incident or maintenance when selecting this template',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'update_status',
              'short' => 'The status the incident or maintenance should transition to when selecting this template',
              'type' => '`$STRING`',
            ],
          ],
          'id' => [
            'field' => 'id',
            'name' => 'id',
          ],
          'name' => 'incident_template',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/incident_templates',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incident_templates',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incident_templates',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'example' => 1,
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'example' => 100,
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incident_templates',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incident_templates',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incident_templates',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
            ],
          ],
        ],
        'incident_update' => [
          'fields' => [
            [
              'name' => 'affected_components',
              'short' => 'Affected components associated with the incident update.',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'body',
              'short' => 'Incident update body.',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'created_at',
              'short' => 'The timestamp when the incident update was created at.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'custom_tweet',
              'short' => 'An optional customized tweet message for incident postmortem.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'deliver_notifications',
              'short' => 'Controls whether to delivery notifications.',
              'type' => '`$BOOLEAN`',
            ],
            [
              'format' => 'date-time',
              'name' => 'display_at',
              'short' => 'Timestamp when incident update is happened.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'short' => 'Incident Update Identifier.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'incident_id',
              'short' => 'Incident Identifier.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'incident_update',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'status',
              'short' => 'The incident status.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'tweet_id',
              'short' => 'Tweet identifier associated to this incident update.',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'twitter_updated_at',
              'short' => 'The timestamp when twitter updated at.',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'updated_at',
              'short' => 'The timestamp when the incident update is updated.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'wants_twitter_update',
              'short' => 'Controls whether to create twitter update.',
              'type' => '`$BOOLEAN`',
            ],
          ],
          'id' => [
            'field' => 'id',
            'name' => 'id',
          ],
          'name' => 'incident_update',
          'op' => [
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'incident_update_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}',
                  'rename' => [
                    'param' => [
                      'incident_update_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                    [
                      'var' => 'incident_id',
                    ],
                    [
                      'lit' => 'incident_updates',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'incident_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'incident_update' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'incident_updates',
                    '{id}',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'incident_update_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}',
                  'rename' => [
                    'param' => [
                      'incident_update_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                    [
                      'var' => 'incident_id',
                    ],
                    [
                      'lit' => 'incident_updates',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'incident_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'incident_update' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'incident_updates',
                    '{id}',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
                'incident',
              ],
            ],
          ],
        ],
        'metric' => [
          'fields' => [
            [
              'format' => 'int32',
              'name' => 'backfill_percentage',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'backfilled',
              'type' => '`$BOOLEAN`',
            ],
            [
              'format' => 'date-time',
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'int32',
              'name' => 'decimal_places',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'display',
              'short' => 'Should the metric be displayed',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'id',
              'short' => 'Metric identifier',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'last_fetched_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'metric',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'metric_identifier',
              'short' => 'Metric Display identifier used to look up the metric data from the provider',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'metrics_provider_id',
              'short' => 'Metric Provider identifier',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'most_recent_data_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'name',
              'short' => 'Name of metric',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'reference_name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'suffix',
              'short' => 'Suffix to describe the units on the graph',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'tooltip_description',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'updated_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'y_axis_hidden',
              'short' => 'Should the values on the y axis be hidden on render',
              'type' => '`$BOOLEAN`',
            ],
            [
              'format' => 'float',
              'name' => 'y_axis_max',
              'type' => '`$NUMBER`',
            ],
            [
              'format' => 'float',
              'name' => 'y_axis_min',
              'type' => '`$NUMBER`',
            ],
          ],
          'id' => [
            'field' => 'id',
            'name' => 'id',
          ],
          'name' => 'metric',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'metric_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/metrics/{metric_id}/data',
                  'rename' => [
                    'param' => [
                      'metric_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'metrics',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'data',
                    ],
                  ],
                  'select' => [
                    '$action' => 'data',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.data`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    '{id}',
                    'data',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'metrics_provider_id',
                        'orig' => 'metrics_provider_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'metrics_providers',
                    ],
                    [
                      'var' => 'metrics_provider_id',
                    ],
                    [
                      'lit' => 'metrics',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'metrics_provider_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'metric' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                    '{metrics_provider_id}',
                    'metrics',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/metrics/data',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'metrics',
                    ],
                    [
                      'lit' => 'data',
                    ],
                  ],
                  'select' => [
                    '$action' => 'data',
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    'data',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_access_user_id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/metrics',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_users',
                    ],
                    [
                      'var' => 'page_access_user_id',
                    ],
                    [
                      'lit' => 'metrics',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page',
                      'page_access_user_id',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{page_access_user_id}',
                    'metrics',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'metrics_provider_id',
                        'orig' => 'metrics_provider_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'metrics_providers',
                    ],
                    [
                      'var' => 'metrics_provider_id',
                    ],
                    [
                      'lit' => 'metrics',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'metrics_provider_id',
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                    '{metrics_provider_id}',
                    'metrics',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/metrics',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'metrics',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'metric_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/metrics/{metric_id}',
                  'rename' => [
                    'param' => [
                      'metric_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'metrics',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    '{id}',
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'metric_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/metrics/{metric_id}',
                  'rename' => [
                    'param' => [
                      'metric_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'metrics',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'metric' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    '{id}',
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'metric_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/metrics/{metric_id}',
                  'rename' => [
                    'param' => [
                      'metric_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'metrics',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    '{id}',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'metric_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/metrics/{metric_id}/data',
                  'rename' => [
                    'param' => [
                      'metric_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'metrics',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'data',
                    ],
                  ],
                  'select' => [
                    '$action' => 'data',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    '{id}',
                    'data',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'metric_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/metrics/{metric_id}',
                  'rename' => [
                    'param' => [
                      'metric_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'metrics',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'metric' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    '{id}',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
              [
                'page',
                'metrics_provider',
              ],
              [
                'page',
                'page_access_user',
              ],
            ],
          ],
        ],
        'metrics_provider' => [
          'fields' => [
            [
              'format' => 'date-time',
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'disabled',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'id',
              'short' => 'Identifier for Metrics Provider',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'last_revalidated_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'metric_base_uri',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'metrics_provider',
              'type' => '`$OBJECT`',
            ],
            [
              'format' => 'int32',
              'name' => 'page_id',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'type',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'updated_at',
              'type' => '`$STRING`',
            ],
          ],
          'id' => [
            'field' => 'id',
            'name' => 'id',
          ],
          'name' => 'metrics_provider',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/metrics_providers',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'metrics_providers',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'metrics_provider' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/metrics_providers',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'metrics_providers',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'metrics_provider_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/metrics_providers/{metrics_provider_id}',
                  'rename' => [
                    'param' => [
                      'metrics_provider_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'metrics_providers',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                    '{id}',
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'metrics_provider_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/metrics_providers/{metrics_provider_id}',
                  'rename' => [
                    'param' => [
                      'metrics_provider_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'metrics_providers',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'metrics_provider' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                    '{id}',
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'metrics_provider_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/metrics_providers/{metrics_provider_id}',
                  'rename' => [
                    'param' => [
                      'metrics_provider_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'metrics_providers',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                    '{id}',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'metrics_provider_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/metrics_providers/{metrics_provider_id}',
                  'rename' => [
                    'param' => [
                      'metrics_provider_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'metrics_providers',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'metrics_provider' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                    '{id}',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
            ],
          ],
        ],
        'page' => [
          'fields' => [
            [
              'format' => 'float',
              'name' => 'activity_score',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'allow_email_subscribers',
              'short' => 'Can your users choose to receive notifications via email',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'allow_incident_subscribers',
              'short' => 'Can your users subscribe to notifications for a single incident',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'allow_page_subscribers',
              'short' => 'Can your users subscribe to all notifications on the page',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'allow_rss_atom_feeds',
              'short' => 'Can your users choose to access incident feeds via RSS/Atom (not functional on Audience-Specific pages)',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'allow_sms_subscribers',
              'short' => 'Can your users choose to receive notifications via SMS',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'allow_webhook_subscribers',
              'short' => 'Can your users choose to receive notifications via Webhooks',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'branding',
              'short' => 'The main template your statuspage will use',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'city',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'country',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'created_at',
              'short' => 'Timestamp the record was created',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_blues',
              'short' => 'CSS Color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_body_background_color',
              'short' => 'CSS Color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_border_color',
              'short' => 'CSS Color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_font_color',
              'short' => 'CSS Color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_graph_color',
              'short' => 'CSS Color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_greens',
              'short' => 'CSS Color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_light_font_color',
              'short' => 'CSS Color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_link_color',
              'short' => 'CSS Color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_no_data',
              'short' => 'CSS Color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_oranges',
              'short' => 'CSS Color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_reds',
              'short' => 'CSS Color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_yellows',
              'short' => 'CSS Color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'domain',
              'short' => 'CNAME alias for your status page',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'email_logo',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'favicon_logo',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'headline',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'hero_cover',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'hidden_from_search',
              'short' => 'Should your page hide itself from search engines',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'id',
              'short' => 'Page identifier',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'ip_restrictions',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'name',
              'short' => 'Name of your page to be displayed',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'notifications_email_footer',
              'short' => 'Allows you to customize the footer appearing on your notification emails.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'notifications_from_email',
              'short' => 'Allows you to customize the email address your page notifications come from',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'page',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'page_description',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'state',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'subdomain',
              'short' => 'Subdomain at which to access your status page',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'support_url',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'time_zone',
              'short' => 'Timezone configured for your page',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'transactional_logo',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'twitter_logo',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'twitter_username',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'updated_at',
              'short' => 'Timestamp the record was last updated',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'url',
              'short' => 'Website of your page.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'viewers_must_be_team_members',
              'type' => '`$BOOLEAN`',
            ],
          ],
          'id' => [
            'field' => 'id',
            'name' => 'id',
          ],
          'name' => 'page',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                  ],
                  'select' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}',
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{id}',
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}',
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'page' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{id}',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}',
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'page' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{id}',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'page_access_group' => [
          'fields' => [
            [
              'name' => 'component_ids',
              'type' => '`$ARRAY`',
            ],
            [
              'format' => 'date-time',
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'external_identifier',
              'short' => 'Associates group with external group.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'short' => 'Page Access Group Identifier',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'metric_ids',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'name',
              'short' => 'Name for this Group.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'page_access_group',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'page_access_user_ids',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'page_id',
              'short' => 'Page Identifier.',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'updated_at',
              'type' => '`$STRING`',
            ],
          ],
          'id' => [
            'field' => 'id',
            'name' => 'id',
          ],
          'name' => 'page_access_group',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}/components',
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_groups',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                    'components',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/page_access_groups',
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'page_access_groups',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'page_access_group' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{id}',
                    'page_access_groups',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/page_access_groups',
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'page_access_groups',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{id}',
                    'page_access_groups',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}',
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_groups',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}',
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_groups',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'page_access_group' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}/components',
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_groups',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                    'components',
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'component_id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}/components/{component_id}',
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_groups',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                    [
                      'var' => 'component_id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'component_id',
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                    'components',
                    '{component_id}',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}',
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_groups',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}/components',
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_groups',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                    'components',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}',
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_groups',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'page_access_group' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}/components',
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_groups',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                    'components',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
              [
                'page',
                'component',
              ],
            ],
          ],
        ],
        'page_access_user' => [
          'fields' => [
            [
              'format' => 'date-time',
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'email',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'external_login',
              'short' => 'IDP login user id.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'short' => 'Page Access User Identifier',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'page_access_group_id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'page_access_group_ids',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'page_access_user',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'page_id',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'updated_at',
              'type' => '`$STRING`',
            ],
          ],
          'id' => [
            'field' => 'id',
            'name' => 'id',
          ],
          'name' => 'page_access_user',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/components',
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_users',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'components',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/metrics',
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_users',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'metrics',
                    ],
                  ],
                  'select' => [
                    '$action' => 'metric',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'metrics',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/page_access_users',
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'page_access_users',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'page_access_user' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{id}',
                    'page_access_users',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'email',
                        'orig' => 'email',
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/page_access_users',
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'page_access_users',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'email',
                      'id',
                      'page',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{id}',
                    'page_access_users',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}',
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_users',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}',
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_users',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/components',
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_users',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'components',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/metrics',
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_users',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'metrics',
                    ],
                  ],
                  'select' => [
                    '$action' => 'metric',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'metrics',
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'component_id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/components/{component_id}',
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_users',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                    [
                      'var' => 'component_id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'component_id',
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'components',
                    '{component_id}',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'metric_id',
                        'orig' => 'metric_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/metrics/{metric_id}',
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_users',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'metrics',
                    ],
                    [
                      'var' => 'metric_id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'metric_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'metrics',
                    '{metric_id}',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}',
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_users',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/components',
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_users',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'components',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/metrics',
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_users',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'metrics',
                    ],
                  ],
                  'select' => [
                    '$action' => 'metric',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'metrics',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}',
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_users',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/components',
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_users',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'components',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/metrics',
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'page_access_users',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'metrics',
                    ],
                  ],
                  'select' => [
                    '$action' => 'metric',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'metrics',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
              [
                'page',
                'component',
              ],
              [
                'page',
                'metric',
              ],
            ],
          ],
        ],
        'permission' => [
          'fields' => [
            [
              'name' => 'id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'pages',
              'short' => 'Pages accessible by the user.',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'user_id',
              'short' => 'User identifier',
              'type' => '`$STRING`',
            ],
          ],
          'id' => [
            'field' => 'id',
            'name' => 'id',
          ],
          'name' => 'permission',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'organization_id',
                        'orig' => 'organization_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/organizations/{organization_id}/permissions/{user_id}',
                  'rename' => [
                    'param' => [
                      'user_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'organizations',
                    ],
                    [
                      'var' => 'organization_id',
                    ],
                    [
                      'lit' => 'permissions',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'organization_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.data`',
                  ],
                  'parts' => [
                    'organizations',
                    '{organization_id}',
                    'permissions',
                    '{id}',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'organization_id',
                        'orig' => 'organization_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/organizations/{organization_id}/permissions/{user_id}',
                  'rename' => [
                    'param' => [
                      'user_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'organizations',
                    ],
                    [
                      'var' => 'organization_id',
                    ],
                    [
                      'lit' => 'permissions',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'organization_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.data`',
                  ],
                  'parts' => [
                    'organizations',
                    '{organization_id}',
                    'permissions',
                    '{id}',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'organization',
              ],
            ],
          ],
        ],
        'postmortem' => [
          'fields' => [
            [
              'name' => 'body',
              'short' => 'Postmortem body',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'body_draft',
              'short' => 'Body draft',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'body_draft_updated_at',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'body_updated_at',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'custom_tweet',
              'short' => 'Custom tweet for Incident Postmortem',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'notify_subscribers',
              'short' => 'Should email subscribers be notified.',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'notify_twitter',
              'short' => 'Should Twitter followers be notified.',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'postmortem',
              'req' => true,
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'preview_key',
              'short' => 'Preview Key',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'published_at',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'updated_at',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'postmortem',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/postmortem',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                    [
                      'var' => 'incident_id',
                    ],
                    [
                      'lit' => 'postmortem',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'incident_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'postmortem',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/postmortem',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                    [
                      'var' => 'incident_id',
                    ],
                    [
                      'lit' => 'postmortem',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'incident_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'postmortem' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'postmortem',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/postmortem/publish',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                    [
                      'var' => 'incident_id',
                    ],
                    [
                      'lit' => 'postmortem',
                    ],
                    [
                      'lit' => 'publish',
                    ],
                  ],
                  'select' => [
                    '$action' => 'publish',
                    'exist' => [
                      'incident_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'postmortem' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'postmortem',
                    'publish',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/postmortem/revert',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                    [
                      'var' => 'incident_id',
                    ],
                    [
                      'lit' => 'postmortem',
                    ],
                    [
                      'lit' => 'revert',
                    ],
                  ],
                  'select' => [
                    '$action' => 'revert',
                    'exist' => [
                      'incident_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'postmortem',
                    'revert',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
                'incident',
              ],
            ],
          ],
        ],
        'status_embed_config' => [
          'fields' => [
            [
              'name' => 'incident_background_color',
              'short' => 'Color of status embed iframe background when displaying incident',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'incident_text_color',
              'short' => 'Color of status embed iframe text when displaying incident',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'maintenance_background_color',
              'short' => 'Color of status embed iframe background when displaying maintenance',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'maintenance_text_color',
              'short' => 'Color of status embed iframe text when displaying maintenance',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'page_id',
              'short' => 'Page identifier',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'position',
              'short' => 'Corner where status embed iframe will appear on page',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'status_embed_config',
              'type' => '`$OBJECT`',
            ],
          ],
          'name' => 'status_embed_config',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/status_embed_config',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'status_embed_config',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'status_embed_config',
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/status_embed_config',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'status_embed_config',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'status_embed_config' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'status_embed_config',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/status_embed_config',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'status_embed_config',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'status_embed_config' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'status_embed_config',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
            ],
          ],
        ],
        'subscriber' => [
          'fields' => [
            [
              'name' => 'component_ids',
              'short' => 'A list of component ids for which the subscriber should recieve updates for.',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'components',
              'short' => 'The components for which the subscriber has elected to receive updates.',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'display_phone_number',
              'short' => 'A formatted version of the phone_number and phone_country pair, nicely formatted for display.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'email',
              'short' => 'The email address to use to contact the subscriber.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'endpoint',
              'short' => 'The URL where a webhook subscriber elects to receive updates.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'short' => 'Subscriber Identifier',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'mode',
              'short' => 'The communication mode of the subscriber.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'obfuscated_channel_name',
              'short' => 'Obfuscated slack channel name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'page_access_user_id',
              'short' => 'The Page Access user this subscriber belongs to (only for audience-specific pages).',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'phone_country',
              'short' => 'The two-character country code representing the country of which the phone_number is a part.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'phone_number',
              'short' => 'The phone number used to contact an SMS subscriber',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'purge_at',
              'short' => 'The timestamp when a quarantined subscriber will be purged (unsubscribed).',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'quarantined_at',
              'short' => 'The timestamp when the subscriber was quarantined due to an issue reaching them.',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'skip_confirmation_notification',
              'short' => 'If this is true, do not notify the user with changes to their subscription.',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'subscriber',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'workspace_name',
              'short' => 'The workspace name of the slack subscriber.',
              'type' => '`$STRING`',
            ],
          ],
          'id' => [
            'field' => 'id',
            'name' => 'id',
          ],
          'name' => 'subscriber',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'subscriber_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/subscribers/{subscriber_id}/resend_confirmation',
                  'rename' => [
                    'param' => [
                      'subscriber_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'subscribers',
                    ],
                    [
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'resend_confirmation',
                    ],
                  ],
                  'select' => [
                    '$action' => 'resend_confirmation',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    '{id}',
                    'resend_confirmation',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/subscribers',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                    [
                      'var' => 'incident_id',
                    ],
                    [
                      'lit' => 'subscribers',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'incident_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'subscriber' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'subscribers',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/subscribers',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'subscribers',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'subscriber' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/subscribers/reactivate',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'subscribers',
                    ],
                    [
                      'lit' => 'reactivate',
                    ],
                  ],
                  'select' => [
                    '$action' => 'reactivate',
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    'reactivate',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/subscribers/resend_confirmation',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'subscribers',
                    ],
                    [
                      'lit' => 'resend_confirmation',
                    ],
                  ],
                  'select' => [
                    '$action' => 'resend_confirmation',
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    'resend_confirmation',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/subscribers/unsubscribe',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'subscribers',
                    ],
                    [
                      'lit' => 'unsubscribe',
                    ],
                  ],
                  'select' => [
                    '$action' => 'unsubscribe',
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    'unsubscribe',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'limit',
                        'orig' => 'limit',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'example' => 0,
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'q',
                        'orig' => 'q',
                        'type' => '`$STRING`',
                      ],
                      [
                        'example' => 'asc',
                        'kind' => 'query',
                        'name' => 'sort_direction',
                        'orig' => 'sort_direction',
                        'type' => '`$STRING`',
                      ],
                      [
                        'example' => 'primary',
                        'kind' => 'query',
                        'name' => 'sort_field',
                        'orig' => 'sort_field',
                        'type' => '`$STRING`',
                      ],
                      [
                        'example' => 'active',
                        'kind' => 'query',
                        'name' => 'state',
                        'orig' => 'state',
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'type',
                        'orig' => 'type',
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/subscribers',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'subscribers',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'limit',
                      'page',
                      'page_id',
                      'q',
                      'sort_direction',
                      'sort_field',
                      'state',
                      'type',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/subscribers',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                    [
                      'var' => 'incident_id',
                    ],
                    [
                      'lit' => 'subscribers',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'incident_id',
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'subscribers',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/subscribers/unsubscribed',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'subscribers',
                    ],
                    [
                      'lit' => 'unsubscribed',
                    ],
                  ],
                  'select' => [
                    '$action' => 'unsubscribed',
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    'unsubscribed',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'subscriber_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}',
                  'rename' => [
                    'param' => [
                      'subscriber_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                    [
                      'var' => 'incident_id',
                    ],
                    [
                      'lit' => 'subscribers',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'incident_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'subscribers',
                    '{id}',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'example' => 'active',
                        'kind' => 'query',
                        'name' => 'state',
                        'orig' => 'state',
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'type',
                        'orig' => 'type',
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/subscribers/count',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'subscribers',
                    ],
                    [
                      'lit' => 'count',
                    ],
                  ],
                  'select' => [
                    '$action' => 'count',
                    'exist' => [
                      'page_id',
                      'state',
                      'type',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    'count',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'subscriber_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/subscribers/{subscriber_id}',
                  'rename' => [
                    'param' => [
                      'subscriber_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'subscribers',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    '{id}',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/subscribers/histogram_by_state',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'subscribers',
                    ],
                    [
                      'lit' => 'histogram_by_state',
                    ],
                  ],
                  'select' => [
                    '$action' => 'histogram_by_state',
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    'histogram_by_state',
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'subscriber_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}',
                  'rename' => [
                    'param' => [
                      'subscriber_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'incidents',
                    ],
                    [
                      'var' => 'incident_id',
                    ],
                    [
                      'lit' => 'subscribers',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'incident_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'subscribers',
                    '{id}',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'subscriber_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'skip_unsubscription_notification',
                        'orig' => 'skip_unsubscription_notification',
                        'type' => '`$BOOLEAN`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/subscribers/{subscriber_id}',
                  'rename' => [
                    'param' => [
                      'subscriber_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'subscribers',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                      'skip_unsubscription_notification',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    '{id}',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'subscriber_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/subscribers/{subscriber_id}',
                  'rename' => [
                    'param' => [
                      'subscriber_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'page_id',
                    ],
                    [
                      'lit' => 'subscribers',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    '{id}',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
              [
                'page',
                'incident',
              ],
            ],
          ],
        ],
        'user' => [
          'fields' => [
            [
              'format' => 'date-time',
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'email',
              'short' => 'Email address for the team member',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'first_name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'short' => 'User identifier',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'last_name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'organization_id',
              'short' => 'Organization identifier',
              'type' => '`$STRING`',
            ],
            [
              'format' => 'date-time',
              'name' => 'updated_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'user',
              'req' => true,
              'type' => '`$OBJECT`',
            ],
          ],
          'id' => [
            'field' => 'id',
            'name' => 'id',
          ],
          'name' => 'user',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'organization_id',
                        'orig' => 'organization_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/organizations/{organization_id}/users',
                  'segments' => [
                    [
                      'lit' => 'organizations',
                    ],
                    [
                      'var' => 'organization_id',
                    ],
                    [
                      'lit' => 'users',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'organization_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'user' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'organizations',
                    '{organization_id}',
                    'users',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'organization_id',
                        'orig' => 'organization_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/organizations/{organization_id}/users',
                  'segments' => [
                    [
                      'lit' => 'organizations',
                    ],
                    [
                      'var' => 'organization_id',
                    ],
                    [
                      'lit' => 'users',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'organization_id',
                      'page',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'organizations',
                    '{organization_id}',
                    'users',
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'organization_id',
                        'orig' => 'organization_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/organizations/{organization_id}/users/{user_id}',
                  'rename' => [
                    'param' => [
                      'user_id' => 'id',
                    ],
                  ],
                  'segments' => [
                    [
                      'lit' => 'organizations',
                    ],
                    [
                      'var' => 'organization_id',
                    ],
                    [
                      'lit' => 'users',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'organization_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'parts' => [
                    'organizations',
                    '{organization_id}',
                    'users',
                    '{id}',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'organization',
              ],
            ],
          ],
        ],
      ],
        ];
    }


    public static function make_feature(string $name)
    {
        require_once __DIR__ . '/features.php';
        return StatuspageFeatures::make_feature($name);
    }
}
