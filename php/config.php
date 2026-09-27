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
                "version" => "0.1.1",
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
              'title' => 'Automation Email',
              'type' => '`$STRING`',
              'short' => 'Requires a special feature flag to be enabled',
            ],
            [
              'name' => 'component',
              'title' => 'Component',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'created_at',
              'title' => 'Created At',
              'type' => '`$STRING`',
              'format' => 'date-time',
            ],
            [
              'name' => 'description',
              'title' => 'Description',
              'type' => '`$STRING`',
              'short' => 'More detailed description for component',
            ],
            [
              'name' => 'group',
              'title' => 'Group',
              'type' => '`$BOOLEAN`',
              'short' => 'Is this component a group',
            ],
            [
              'name' => 'group_id',
              'title' => 'Group Id',
              'type' => '`$STRING`',
              'short' => 'Component Group identifier',
            ],
            [
              'name' => 'id',
              'title' => 'Id',
              'type' => '`$STRING`',
              'short' => 'Identifier for component',
            ],
            [
              'name' => 'name',
              'title' => 'Name',
              'type' => '`$STRING`',
              'short' => 'Display name for component',
            ],
            [
              'name' => 'only_show_if_degraded',
              'title' => 'Only Show If Degraded',
              'type' => '`$BOOLEAN`',
              'short' => 'Requires a special feature flag to be enabled',
            ],
            [
              'name' => 'page_id',
              'title' => 'Page Id',
              'type' => '`$STRING`',
              'short' => 'Page identifier',
            ],
            [
              'name' => 'position',
              'title' => 'Position',
              'type' => '`$INTEGER`',
              'short' => 'Order the component will appear on the page',
              'format' => 'int32',
            ],
            [
              'name' => 'showcase',
              'title' => 'Showcase',
              'type' => '`$BOOLEAN`',
              'short' => 'Should this component be showcased',
            ],
            [
              'name' => 'start_date',
              'title' => 'Start Date',
              'type' => '`$STRING`',
              'short' => 'The date this component started being used',
              'format' => 'date',
            ],
            [
              'name' => 'status',
              'title' => 'Status',
              'type' => '`$STRING`',
              'short' => 'Status of component',
            ],
            [
              'name' => 'updated_at',
              'title' => 'Updated At',
              'type' => '`$STRING`',
              'format' => 'date-time',
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
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/components/{component_id}/page_access_groups',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                    'page_access_groups',
                  ],
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'component_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'page_access_group',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/components/{component_id}/page_access_users',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                    'page_access_users',
                  ],
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'component_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'page_access_user',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => [
                      'component' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{page_access_group_id}',
                    'components',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_access_group_id',
                        'orig' => 'page_access_group_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
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
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{page_access_user_id}',
                    'components',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_access_user_id',
                        'orig' => 'page_access_user_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
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
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/components/{component_id}/uptime',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                    'uptime',
                  ],
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.related_events`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'component_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'end',
                        'orig' => 'end',
                        'type' => 'Any',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'start',
                        'orig' => 'start',
                        'type' => 'Any',
                        'kind' => 'query',
                      ],
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
                ],
                [
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/components/{component_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'component_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/components/{component_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'component' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'component_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/components/{component_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'component_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/components/{component_id}/page_access_groups',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                    'page_access_groups',
                  ],
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'component_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'page_access_group',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/components/{component_id}/page_access_users',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                    'page_access_users',
                  ],
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'component_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'page_access_user',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/components/{component_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'component' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'component_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                '$.main.kit.entity.page',
              ],
              [
                '$.main.kit.entity.page',
                '$.main.kit.entity.page_access_group',
              ],
              [
                '$.main.kit.entity.page',
                '$.main.kit.entity.page_access_user',
              ],
            ],
          ],
        ],
        'component_group_uptime' => [
          'fields' => [
            [
              'name' => 'component_id',
              'title' => 'Component Id',
              'type' => '`$STRING`',
              'short' => 'Component identifier',
            ],
            [
              'name' => 'id',
              'title' => 'Id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'incidents',
              'title' => 'Incidents',
              'type' => '`$OBJECT`',
              'short' => 'Related incidents',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                    '{id}',
                    'uptime',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.related_events`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'end',
                        'orig' => 'end',
                        'type' => 'Any',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'start',
                        'orig' => 'start',
                        'type' => 'Any',
                        'kind' => 'query',
                      ],
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
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                '$.main.kit.entity.page',
              ],
            ],
          ],
        ],
        'group_component' => [
          'fields' => [
            [
              'name' => 'component_group',
              'title' => 'Component Group',
              'type' => '`$OBJECT`',
              'req' => true,
            ],
            [
              'name' => 'components',
              'title' => 'Components',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'created_at',
              'title' => 'Created At',
              'type' => '`$STRING`',
              'format' => 'date-time',
            ],
            [
              'name' => 'description',
              'title' => 'Description',
              'type' => '`$STRING`',
              'short' => 'Description of the component group.',
            ],
            [
              'name' => 'id',
              'title' => 'Id',
              'type' => '`$STRING`',
              'short' => 'Component Group Identifier',
            ],
            [
              'name' => 'name',
              'title' => 'Name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'page_id',
              'title' => 'Page Id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'position',
              'title' => 'Position',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'updated_at',
              'title' => 'Updated At',
              'type' => '`$STRING`',
              'format' => 'date-time',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                    '{id}',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                    '{id}',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                    '{id}',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                    '{id}',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                '$.main.kit.entity.page',
              ],
            ],
          ],
        ],
        'incident' => [
          'fields' => [
            [
              'name' => 'auto_transition_deliver_notifications_at_end',
              'title' => 'Auto Transition Deliver Notifications At End',
              'type' => '`$BOOLEAN`',
              'short' => 'Controls whether send notification when scheduled maintenances auto transition to completed.',
            ],
            [
              'name' => 'auto_transition_deliver_notifications_at_start',
              'title' => 'Auto Transition Deliver Notifications At Start',
              'type' => '`$BOOLEAN`',
              'short' => 'Controls whether send notification when scheduled maintenances auto transition to started.',
            ],
            [
              'name' => 'auto_transition_to_maintenance_state',
              'title' => 'Auto Transition To Maintenance State',
              'type' => '`$BOOLEAN`',
              'short' => 'Controls whether change components status to under_maintenance once scheduled maintenance is in progress.',
            ],
            [
              'name' => 'auto_transition_to_operational_state',
              'title' => 'Auto Transition To Operational State',
              'type' => '`$BOOLEAN`',
              'short' => 'Controls whether change components status to operational once scheduled maintenance completes.',
            ],
            [
              'name' => 'components',
              'title' => 'Components',
              'type' => '`$ARRAY`',
              'short' => 'Incident components',
            ],
            [
              'name' => 'created_at',
              'title' => 'Created At',
              'type' => '`$STRING`',
              'short' => 'The timestamp when the incident was created at.',
              'format' => 'date-time',
            ],
            [
              'name' => 'id',
              'title' => 'Id',
              'type' => '`$STRING`',
              'short' => 'Incident Identifier',
            ],
            [
              'name' => 'impact',
              'title' => 'Impact',
              'type' => '`$STRING`',
              'short' => 'The impact of the incident.',
            ],
            [
              'name' => 'impact_override',
              'title' => 'Impact Override',
              'type' => '`$STRING`',
              'short' => 'value to override calculated impact value',
            ],
            [
              'name' => 'incident',
              'title' => 'Incident',
              'type' => '`$OBJECT`',
              'req' => true,
              'op' => [
                'patch' => [
                  'type' => '`$OBJECT`',
                ],
                'update' => [
                  'type' => '`$OBJECT`',
                ],
              ],
            ],
            [
              'name' => 'incident_updates',
              'title' => 'Incident Updates',
              'type' => '`$ARRAY`',
              'short' => 'The incident updates for incident.',
            ],
            [
              'name' => 'metadata',
              'title' => 'Metadata',
              'type' => '`$OBJECT`',
              'short' => 'Metadata attached to the incident.',
            ],
            [
              'name' => 'monitoring_at',
              'title' => 'Monitoring At',
              'type' => '`$STRING`',
              'short' => 'The timestamp when incident entered monitoring state.',
              'format' => 'date-time',
            ],
            [
              'name' => 'name',
              'title' => 'Name',
              'type' => '`$STRING`',
              'short' => 'Incident Name.',
            ],
            [
              'name' => 'page_id',
              'title' => 'Page Id',
              'type' => '`$STRING`',
              'short' => 'Incident Page Identifier',
            ],
            [
              'name' => 'postmortem_body',
              'title' => 'Postmortem Body',
              'type' => '`$STRING`',
              'short' => 'Body of the Postmortem.',
            ],
            [
              'name' => 'postmortem_body_last_updated_at',
              'title' => 'Postmortem Body Last Updated At',
              'type' => '`$STRING`',
              'short' => 'The timestamp when the incident postmortem body was last updated at.',
              'format' => 'date-time',
            ],
            [
              'name' => 'postmortem_ignored',
              'title' => 'Postmortem Ignored',
              'type' => '`$BOOLEAN`',
              'short' => 'Controls whether the incident will have postmortem.',
            ],
            [
              'name' => 'postmortem_notified_subscribers',
              'title' => 'Postmortem Notified Subscribers',
              'type' => '`$BOOLEAN`',
              'short' => 'Indicates whether subscribers are already notificed about postmortem.',
            ],
            [
              'name' => 'postmortem_notified_twitter',
              'title' => 'Postmortem Notified Twitter',
              'type' => '`$BOOLEAN`',
              'short' => 'Controls whether to decide if notify postmortem on twitter.',
            ],
            [
              'name' => 'postmortem_published_at',
              'title' => 'Postmortem Published At',
              'type' => '`$BOOLEAN`',
              'short' => 'The timestamp when the postmortem was published.',
            ],
            [
              'name' => 'reminder_intervals',
              'title' => 'Reminder Intervals',
              'type' => '`$STRING`',
              'short' => 'Custom reminder intervals for unresolved/open incidents.',
            ],
            [
              'name' => 'resolved_at',
              'title' => 'Resolved At',
              'type' => '`$STRING`',
              'short' => 'The timestamp when incident was resolved.',
              'format' => 'date-time',
            ],
            [
              'name' => 'scheduled_auto_completed',
              'title' => 'Scheduled Auto Completed',
              'type' => '`$BOOLEAN`',
              'short' => 'Controls whether the incident is scheduled to automatically change to complete.',
            ],
            [
              'name' => 'scheduled_auto_in_progress',
              'title' => 'Scheduled Auto In Progress',
              'type' => '`$BOOLEAN`',
              'short' => 'Controls whether the incident is scheduled to automatically change to in progress.',
            ],
            [
              'name' => 'scheduled_for',
              'title' => 'Scheduled For',
              'type' => '`$STRING`',
              'short' => 'The timestamp the incident is scheduled for.',
              'format' => 'date-time',
            ],
            [
              'name' => 'scheduled_remind_prior',
              'title' => 'Scheduled Remind Prior',
              'type' => '`$BOOLEAN`',
              'short' => 'Controls whether to remind subscribers prior to scheduled incidents.',
            ],
            [
              'name' => 'scheduled_reminded_at',
              'title' => 'Scheduled Reminded At',
              'type' => '`$STRING`',
              'short' => 'The timestamp when the scheduled incident reminder was sent at.',
              'format' => 'date-time',
            ],
            [
              'name' => 'scheduled_until',
              'title' => 'Scheduled Until',
              'type' => '`$STRING`',
              'short' => 'The timestamp the incident is scheduled until.',
              'format' => 'date-time',
            ],
            [
              'name' => 'shortlink',
              'title' => 'Shortlink',
              'type' => '`$STRING`',
              'short' => 'Incident Shortlink.',
            ],
            [
              'name' => 'status',
              'title' => 'Status',
              'type' => '`$STRING`',
              'short' => 'The incident status.',
            ],
            [
              'name' => 'updated_at',
              'title' => 'Updated At',
              'type' => '`$STRING`',
              'short' => 'The timestamp when the incident was updated at.',
              'format' => 'date-time',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => [
                      'incident' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'limit',
                        'orig' => 'limit',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'q',
                        'orig' => 'q',
                        'type' => '`$STRING`',
                        'kind' => 'query',
                      ],
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
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    'active_maintenance',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                        'example' => 1,
                      ],
                      [
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                        'example' => 100,
                      ],
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
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    'scheduled',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                        'example' => 1,
                      ],
                      [
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                        'example' => 100,
                      ],
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
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    'unresolved',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                        'example' => 1,
                      ],
                      [
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                        'example' => 100,
                      ],
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
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    'upcoming',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                        'example' => 1,
                      ],
                      [
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                        'example' => 100,
                      ],
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
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'incident_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'incident_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'incident_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'incident' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'incident_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'incident_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'incident_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'incident_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'incident' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'incident_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                '$.main.kit.entity.page',
              ],
            ],
          ],
        ],
        'incident_postmortem' => [
          'fields' => [
            [
              'name' => 'id',
              'title' => 'Id',
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
                  'kind' => 'http',
                  'method' => 'DELETE',
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
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'postmortem',
                    ],
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{id}',
                    'postmortem',
                  ],
                  'rename' => [
                    'param' => [
                      'incident_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'incident_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                '$.main.kit.entity.page',
              ],
            ],
          ],
        ],
        'incident_template' => [
          'fields' => [
            [
              'name' => 'body',
              'title' => 'Body',
              'type' => '`$STRING`',
              'short' => 'Body of the incident or maintenance update to be applied when selecting this template',
            ],
            [
              'name' => 'components',
              'title' => 'Components',
              'type' => '`$ARRAY`',
              'short' => 'Affected components',
            ],
            [
              'name' => 'group_id',
              'title' => 'Group Id',
              'type' => '`$STRING`',
              'short' => 'Identifier of Template Group this template belongs to',
            ],
            [
              'name' => 'id',
              'title' => 'Id',
              'type' => '`$STRING`',
              'short' => 'Incident Template Identifier',
            ],
            [
              'name' => 'name',
              'title' => 'Name',
              'type' => '`$STRING`',
              'short' => 'Name of the template, as shown in the list on the "Templates" tab of the "Incidents" page',
            ],
            [
              'name' => 'should_send_notifications',
              'title' => 'Should Send Notifications',
              'type' => '`$BOOLEAN`',
              'short' => 'Whether the "deliver notifications" checkbox should be selected when selecting this template',
            ],
            [
              'name' => 'should_tweet',
              'title' => 'Should Tweet',
              'type' => '`$BOOLEAN`',
              'short' => 'Whether the "tweet update" checkbox should be selected when selecting this template',
            ],
            [
              'name' => 'template',
              'title' => 'Template',
              'type' => '`$OBJECT`',
              'req' => true,
            ],
            [
              'name' => 'title',
              'title' => 'Title',
              'type' => '`$STRING`',
              'short' => 'Title to be applied to the incident or maintenance when selecting this template',
            ],
            [
              'name' => 'update_status',
              'title' => 'Update Status',
              'type' => '`$STRING`',
              'short' => 'The status the incident or maintenance should transition to when selecting this template',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incident_templates',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incident_templates',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                        'example' => 1,
                      ],
                      [
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                        'example' => 100,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                '$.main.kit.entity.page',
              ],
            ],
          ],
        ],
        'incident_update' => [
          'fields' => [
            [
              'name' => 'affected_components',
              'title' => 'Affected Components',
              'type' => '`$ARRAY`',
              'short' => 'Affected components associated with the incident update.',
            ],
            [
              'name' => 'body',
              'title' => 'Body',
              'type' => '`$STRING`',
              'short' => 'Incident update body.',
            ],
            [
              'name' => 'created_at',
              'title' => 'Created At',
              'type' => '`$STRING`',
              'short' => 'The timestamp when the incident update was created at.',
              'format' => 'date-time',
            ],
            [
              'name' => 'custom_tweet',
              'title' => 'Custom Tweet',
              'type' => '`$STRING`',
              'short' => 'An optional customized tweet message for incident postmortem.',
            ],
            [
              'name' => 'deliver_notifications',
              'title' => 'Deliver Notifications',
              'type' => '`$BOOLEAN`',
              'short' => 'Controls whether to delivery notifications.',
            ],
            [
              'name' => 'display_at',
              'title' => 'Display At',
              'type' => '`$STRING`',
              'short' => 'Timestamp when incident update is happened.',
              'format' => 'date-time',
            ],
            [
              'name' => 'id',
              'title' => 'Id',
              'type' => '`$STRING`',
              'short' => 'Incident Update Identifier.',
            ],
            [
              'name' => 'incident_id',
              'title' => 'Incident Id',
              'type' => '`$STRING`',
              'short' => 'Incident Identifier.',
            ],
            [
              'name' => 'incident_update',
              'title' => 'Incident Update',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'status',
              'title' => 'Status',
              'type' => '`$STRING`',
              'short' => 'The incident status.',
            ],
            [
              'name' => 'tweet_id',
              'title' => 'Tweet Id',
              'type' => '`$STRING`',
              'short' => 'Tweet identifier associated to this incident update.',
            ],
            [
              'name' => 'twitter_updated_at',
              'title' => 'Twitter Updated At',
              'type' => '`$STRING`',
              'short' => 'The timestamp when twitter updated at.',
              'format' => 'date-time',
            ],
            [
              'name' => 'updated_at',
              'title' => 'Updated At',
              'type' => '`$STRING`',
              'short' => 'The timestamp when the incident update is updated.',
              'format' => 'date-time',
            ],
            [
              'name' => 'wants_twitter_update',
              'title' => 'Wants Twitter Update',
              'type' => '`$BOOLEAN`',
              'short' => 'Controls whether to create twitter update.',
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
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'incident_updates',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'incident_update_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'incident_update' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'incident_update_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'incident_id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'incident_updates',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'incident_update_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'incident_update' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'incident_update_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'incident_id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                '$.main.kit.entity.page',
                '$.main.kit.entity.incident',
              ],
            ],
          ],
        ],
        'metric' => [
          'fields' => [
            [
              'name' => 'backfill_percentage',
              'title' => 'Backfill Percentage',
              'type' => '`$INTEGER`',
              'format' => 'int32',
            ],
            [
              'name' => 'backfilled',
              'title' => 'Backfilled',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'created_at',
              'title' => 'Created At',
              'type' => '`$STRING`',
              'format' => 'date-time',
            ],
            [
              'name' => 'decimal_places',
              'title' => 'Decimal Places',
              'type' => '`$INTEGER`',
              'format' => 'int32',
            ],
            [
              'name' => 'display',
              'title' => 'Display',
              'type' => '`$BOOLEAN`',
              'short' => 'Should the metric be displayed',
            ],
            [
              'name' => 'id',
              'title' => 'Id',
              'type' => '`$STRING`',
              'short' => 'Metric identifier',
            ],
            [
              'name' => 'last_fetched_at',
              'title' => 'Last Fetched At',
              'type' => '`$STRING`',
              'format' => 'date-time',
            ],
            [
              'name' => 'metric',
              'title' => 'Metric',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'metric_identifier',
              'title' => 'Metric Identifier',
              'type' => '`$STRING`',
              'short' => 'Metric Display identifier used to look up the metric data from the provider',
            ],
            [
              'name' => 'metrics_provider_id',
              'title' => 'Metrics Provider Id',
              'type' => '`$STRING`',
              'short' => 'Metric Provider identifier',
            ],
            [
              'name' => 'most_recent_data_at',
              'title' => 'Most Recent Data At',
              'type' => '`$STRING`',
              'format' => 'date-time',
            ],
            [
              'name' => 'name',
              'title' => 'Name',
              'type' => '`$STRING`',
              'short' => 'Name of metric',
            ],
            [
              'name' => 'reference_name',
              'title' => 'Reference Name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'suffix',
              'title' => 'Suffix',
              'type' => '`$STRING`',
              'short' => 'Suffix to describe the units on the graph',
            ],
            [
              'name' => 'tooltip_description',
              'title' => 'Tooltip Description',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'updated_at',
              'title' => 'Updated At',
              'type' => '`$STRING`',
              'format' => 'date-time',
            ],
            [
              'name' => 'y_axis_hidden',
              'title' => 'Y Axis Hidden',
              'type' => '`$BOOLEAN`',
              'short' => 'Should the values on the y axis be hidden on render',
            ],
            [
              'name' => 'y_axis_max',
              'title' => 'Y Axis Max',
              'type' => '`$NUMBER`',
              'format' => 'float',
            ],
            [
              'name' => 'y_axis_min',
              'title' => 'Y Axis Min',
              'type' => '`$NUMBER`',
              'format' => 'float',
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
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/metrics/{metric_id}/data',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    '{id}',
                    'data',
                  ],
                  'rename' => [
                    'param' => [
                      'metric_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.data`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'metric_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'data',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                    '{metrics_provider_id}',
                    'metrics',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => [
                      'metric' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'metrics_provider_id',
                        'orig' => 'metrics_provider_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'metrics_provider_id',
                      'page_id',
                    ],
                  ],
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    'data',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'data',
                    'exist' => [
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{page_access_user_id}',
                    'metrics',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_access_user_id',
                        'orig' => 'page_access_user_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
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
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                    '{metrics_provider_id}',
                    'metrics',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'metrics_provider_id',
                        'orig' => 'metrics_provider_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
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
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/metrics/{metric_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'metric_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'metric_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/metrics/{metric_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'metric_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'metric' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'metric_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/metrics/{metric_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'metric_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'metric_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/metrics/{metric_id}/data',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    '{id}',
                    'data',
                  ],
                  'rename' => [
                    'param' => [
                      'metric_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'metric_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'data',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/metrics/{metric_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'metric_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'metric' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'metric_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                '$.main.kit.entity.page',
              ],
              [
                '$.main.kit.entity.page',
                '$.main.kit.entity.metrics_provider',
              ],
              [
                '$.main.kit.entity.page',
                '$.main.kit.entity.page_access_user',
              ],
            ],
          ],
        ],
        'metrics_provider' => [
          'fields' => [
            [
              'name' => 'created_at',
              'title' => 'Created At',
              'type' => '`$STRING`',
              'format' => 'date-time',
            ],
            [
              'name' => 'disabled',
              'title' => 'Disabled',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'id',
              'title' => 'Id',
              'type' => '`$STRING`',
              'short' => 'Identifier for Metrics Provider',
            ],
            [
              'name' => 'last_revalidated_at',
              'title' => 'Last Revalidated At',
              'type' => '`$STRING`',
              'format' => 'date-time',
            ],
            [
              'name' => 'metric_base_uri',
              'title' => 'Metric Base Uri',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'metrics_provider',
              'title' => 'Metrics Provider',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'page_id',
              'title' => 'Page Id',
              'type' => '`$INTEGER`',
              'format' => 'int32',
            ],
            [
              'name' => 'type',
              'title' => 'Type',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'updated_at',
              'title' => 'Updated At',
              'type' => '`$STRING`',
              'format' => 'date-time',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => [
                      'metrics_provider' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/metrics_providers/{metrics_provider_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'metrics_provider_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'metrics_provider_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/metrics_providers/{metrics_provider_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'metrics_provider_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'metrics_provider' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'metrics_provider_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/metrics_providers/{metrics_provider_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'metrics_provider_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'metrics_provider_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/metrics_providers/{metrics_provider_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'metrics_provider_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'metrics_provider' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'metrics_provider_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                '$.main.kit.entity.page',
              ],
            ],
          ],
        ],
        'page' => [
          'fields' => [
            [
              'name' => 'activity_score',
              'title' => 'Activity Score',
              'type' => '`$NUMBER`',
              'format' => 'float',
            ],
            [
              'name' => 'allow_email_subscribers',
              'title' => 'Allow Email Subscribers',
              'type' => '`$BOOLEAN`',
              'short' => 'Can your users choose to receive notifications via email',
            ],
            [
              'name' => 'allow_incident_subscribers',
              'title' => 'Allow Incident Subscribers',
              'type' => '`$BOOLEAN`',
              'short' => 'Can your users subscribe to notifications for a single incident',
            ],
            [
              'name' => 'allow_page_subscribers',
              'title' => 'Allow Page Subscribers',
              'type' => '`$BOOLEAN`',
              'short' => 'Can your users subscribe to all notifications on the page',
            ],
            [
              'name' => 'allow_rss_atom_feeds',
              'title' => 'Allow Rss Atom Feeds',
              'type' => '`$BOOLEAN`',
              'short' => 'Can your users choose to access incident feeds via RSS/Atom (not functional on Audience-Specific pages)',
            ],
            [
              'name' => 'allow_sms_subscribers',
              'title' => 'Allow Sms Subscribers',
              'type' => '`$BOOLEAN`',
              'short' => 'Can your users choose to receive notifications via SMS',
            ],
            [
              'name' => 'allow_webhook_subscribers',
              'title' => 'Allow Webhook Subscribers',
              'type' => '`$BOOLEAN`',
              'short' => 'Can your users choose to receive notifications via Webhooks',
            ],
            [
              'name' => 'branding',
              'title' => 'Branding',
              'type' => '`$STRING`',
              'short' => 'The main template your statuspage will use',
            ],
            [
              'name' => 'city',
              'title' => 'City',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'country',
              'title' => 'Country',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'created_at',
              'title' => 'Created At',
              'type' => '`$STRING`',
              'short' => 'Timestamp the record was created',
              'format' => 'date-time',
            ],
            [
              'name' => 'css_blues',
              'title' => 'Css Blues',
              'type' => '`$STRING`',
              'short' => 'CSS Color',
            ],
            [
              'name' => 'css_body_background_color',
              'title' => 'Css Body Background Color',
              'type' => '`$STRING`',
              'short' => 'CSS Color',
            ],
            [
              'name' => 'css_border_color',
              'title' => 'Css Border Color',
              'type' => '`$STRING`',
              'short' => 'CSS Color',
            ],
            [
              'name' => 'css_font_color',
              'title' => 'Css Font Color',
              'type' => '`$STRING`',
              'short' => 'CSS Color',
            ],
            [
              'name' => 'css_graph_color',
              'title' => 'Css Graph Color',
              'type' => '`$STRING`',
              'short' => 'CSS Color',
            ],
            [
              'name' => 'css_greens',
              'title' => 'Css Greens',
              'type' => '`$STRING`',
              'short' => 'CSS Color',
            ],
            [
              'name' => 'css_light_font_color',
              'title' => 'Css Light Font Color',
              'type' => '`$STRING`',
              'short' => 'CSS Color',
            ],
            [
              'name' => 'css_link_color',
              'title' => 'Css Link Color',
              'type' => '`$STRING`',
              'short' => 'CSS Color',
            ],
            [
              'name' => 'css_no_data',
              'title' => 'Css No Data',
              'type' => '`$STRING`',
              'short' => 'CSS Color',
            ],
            [
              'name' => 'css_oranges',
              'title' => 'Css Oranges',
              'type' => '`$STRING`',
              'short' => 'CSS Color',
            ],
            [
              'name' => 'css_reds',
              'title' => 'Css Reds',
              'type' => '`$STRING`',
              'short' => 'CSS Color',
            ],
            [
              'name' => 'css_yellows',
              'title' => 'Css Yellows',
              'type' => '`$STRING`',
              'short' => 'CSS Color',
            ],
            [
              'name' => 'domain',
              'title' => 'Domain',
              'type' => '`$STRING`',
              'short' => 'CNAME alias for your status page',
            ],
            [
              'name' => 'email_logo',
              'title' => 'Email Logo',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'favicon_logo',
              'title' => 'Favicon Logo',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'headline',
              'title' => 'Headline',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'hero_cover',
              'title' => 'Hero Cover',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'hidden_from_search',
              'title' => 'Hidden From Search',
              'type' => '`$BOOLEAN`',
              'short' => 'Should your page hide itself from search engines',
            ],
            [
              'name' => 'id',
              'title' => 'Id',
              'type' => '`$STRING`',
              'short' => 'Page identifier',
            ],
            [
              'name' => 'ip_restrictions',
              'title' => 'Ip Restrictions',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'name',
              'title' => 'Name',
              'type' => '`$STRING`',
              'short' => 'Name of your page to be displayed',
            ],
            [
              'name' => 'notifications_email_footer',
              'title' => 'Notifications Email Footer',
              'type' => '`$STRING`',
              'short' => 'Allows you to customize the footer appearing on your notification emails.',
            ],
            [
              'name' => 'notifications_from_email',
              'title' => 'Notifications From Email',
              'type' => '`$STRING`',
              'short' => 'Allows you to customize the email address your page notifications come from',
            ],
            [
              'name' => 'page',
              'title' => 'Page',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'page_description',
              'title' => 'Page Description',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'state',
              'title' => 'State',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'subdomain',
              'title' => 'Subdomain',
              'type' => '`$STRING`',
              'short' => 'Subdomain at which to access your status page',
            ],
            [
              'name' => 'support_url',
              'title' => 'Support Url',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'time_zone',
              'title' => 'Time Zone',
              'type' => '`$STRING`',
              'short' => 'Timezone configured for your page',
            ],
            [
              'name' => 'transactional_logo',
              'title' => 'Transactional Logo',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'twitter_logo',
              'title' => 'Twitter Logo',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'twitter_username',
              'title' => 'Twitter Username',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'updated_at',
              'title' => 'Updated At',
              'type' => '`$STRING`',
              'short' => 'Timestamp the record was last updated',
              'format' => 'date-time',
            ],
            [
              'name' => 'url',
              'title' => 'Url',
              'type' => '`$STRING`',
              'short' => 'Website of your page.',
            ],
            [
              'name' => 'viewers_must_be_team_members',
              'title' => 'Viewers Must Be Team Members',
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
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                  ],
                  'parts' => [
                    'pages',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [],
                  'select' => [],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'parts' => [
                    'pages',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'parts' => [
                    'pages',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'page' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}',
                  'segments' => [
                    [
                      'lit' => 'pages',
                    ],
                    [
                      'var' => 'id',
                    ],
                  ],
                  'parts' => [
                    'pages',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'page' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
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
              'title' => 'Component Ids',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'created_at',
              'title' => 'Created At',
              'type' => '`$STRING`',
              'format' => 'date-time',
            ],
            [
              'name' => 'external_identifier',
              'title' => 'External Identifier',
              'type' => '`$STRING`',
              'short' => 'Associates group with external group.',
            ],
            [
              'name' => 'id',
              'title' => 'Id',
              'type' => '`$STRING`',
              'short' => 'Page Access Group Identifier',
            ],
            [
              'name' => 'metric_ids',
              'title' => 'Metric Ids',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'name',
              'title' => 'Name',
              'type' => '`$STRING`',
              'short' => 'Name for this Group.',
            ],
            [
              'name' => 'page_access_group',
              'title' => 'Page Access Group',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'page_access_user_ids',
              'title' => 'Page Access User Ids',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'page_id',
              'title' => 'Page Id',
              'type' => '`$STRING`',
              'short' => 'Page Identifier.',
            ],
            [
              'name' => 'updated_at',
              'title' => 'Updated At',
              'type' => '`$STRING`',
              'format' => 'date-time',
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
                  'kind' => 'http',
                  'method' => 'POST',
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
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                    'components',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/page_access_groups',
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
                  'parts' => [
                    'pages',
                    '{id}',
                    'page_access_groups',
                  ],
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'page_access_group' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/page_access_groups',
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
                  'parts' => [
                    'pages',
                    '{id}',
                    'page_access_groups',
                  ],
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page',
                      'per_page',
                    ],
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'page_access_group' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'PATCH',
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
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                    'components',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}/components/{component_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                    'components',
                    '{component_id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'component_id',
                        'orig' => 'component_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'component_id',
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'DELETE',
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
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                    'components',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'page_access_group' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'PUT',
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
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                    'components',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                '$.main.kit.entity.page',
              ],
              [
                '$.main.kit.entity.page',
                '$.main.kit.entity.component',
              ],
            ],
          ],
        ],
        'page_access_user' => [
          'fields' => [
            [
              'name' => 'created_at',
              'title' => 'Created At',
              'type' => '`$STRING`',
              'format' => 'date-time',
            ],
            [
              'name' => 'email',
              'title' => 'Email',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'external_login',
              'title' => 'External Login',
              'type' => '`$STRING`',
              'short' => 'IDP login user id.',
            ],
            [
              'name' => 'id',
              'title' => 'Id',
              'type' => '`$STRING`',
              'short' => 'Page Access User Identifier',
            ],
            [
              'name' => 'page_access_group_id',
              'title' => 'Page Access Group Id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'page_access_group_ids',
              'title' => 'Page Access Group Ids',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'page_access_user',
              'title' => 'Page Access User',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'page_id',
              'title' => 'Page Id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'updated_at',
              'title' => 'Updated At',
              'type' => '`$STRING`',
              'format' => 'date-time',
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
                  'kind' => 'http',
                  'method' => 'POST',
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
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'components',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'POST',
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
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'metrics',
                    ],
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'metrics',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'metric',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/page_access_users',
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
                  'parts' => [
                    'pages',
                    '{id}',
                    'page_access_users',
                  ],
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'page_access_user' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/page_access_users',
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
                  'parts' => [
                    'pages',
                    '{id}',
                    'page_access_users',
                  ],
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'email',
                        'orig' => 'email',
                        'type' => '`$STRING`',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
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
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'PATCH',
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
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'components',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'PATCH',
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
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'metrics',
                    ],
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'metrics',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'metric',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/components/{component_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'components',
                    '{component_id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'component_id',
                        'orig' => 'component_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'component_id',
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/metrics/{metric_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'metrics',
                    '{metric_id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'metric_id',
                        'orig' => 'metric_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'metric_id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'DELETE',
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
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'components',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'DELETE',
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
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'metrics',
                    ],
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'metrics',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'metric',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'PUT',
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
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'components',
                    ],
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'components',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'PUT',
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
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'metrics',
                    ],
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'metrics',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'metric',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                '$.main.kit.entity.page',
              ],
              [
                '$.main.kit.entity.page',
                '$.main.kit.entity.component',
              ],
              [
                '$.main.kit.entity.page',
                '$.main.kit.entity.metric',
              ],
            ],
          ],
        ],
        'permission' => [
          'fields' => [
            [
              'name' => 'id',
              'title' => 'Id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'pages',
              'title' => 'Pages',
              'type' => '`$OBJECT`',
              'short' => 'Pages accessible by the user.',
            ],
            [
              'name' => 'user_id',
              'title' => 'User Id',
              'type' => '`$STRING`',
              'short' => 'User identifier',
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
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/organizations/{organization_id}/permissions/{user_id}',
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
                  'parts' => [
                    'organizations',
                    '{organization_id}',
                    'permissions',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'user_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.data`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'user_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'organization_id',
                        'orig' => 'organization_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'organization_id',
                    ],
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/organizations/{organization_id}/permissions/{user_id}',
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
                  'parts' => [
                    'organizations',
                    '{organization_id}',
                    'permissions',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'user_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.data`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'user_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'organization_id',
                        'orig' => 'organization_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'organization_id',
                    ],
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'postmortem' => [
          'fields' => [
            [
              'name' => 'body',
              'title' => 'Body',
              'type' => '`$STRING`',
              'short' => 'Postmortem body',
            ],
            [
              'name' => 'body_draft',
              'title' => 'Body Draft',
              'type' => '`$STRING`',
              'short' => 'Body draft',
            ],
            [
              'name' => 'body_draft_updated_at',
              'title' => 'Body Draft Updated At',
              'type' => '`$STRING`',
              'format' => 'date-time',
            ],
            [
              'name' => 'body_updated_at',
              'title' => 'Body Updated At',
              'type' => '`$STRING`',
              'format' => 'date-time',
            ],
            [
              'name' => 'created_at',
              'title' => 'Created At',
              'type' => '`$STRING`',
              'format' => 'date-time',
            ],
            [
              'name' => 'custom_tweet',
              'title' => 'Custom Tweet',
              'type' => '`$STRING`',
              'short' => 'Custom tweet for Incident Postmortem',
            ],
            [
              'name' => 'notify_subscribers',
              'title' => 'Notify Subscribers',
              'type' => '`$BOOLEAN`',
              'short' => 'Should email subscribers be notified.',
            ],
            [
              'name' => 'notify_twitter',
              'title' => 'Notify Twitter',
              'type' => '`$BOOLEAN`',
              'short' => 'Should Twitter followers be notified.',
            ],
            [
              'name' => 'postmortem',
              'title' => 'Postmortem',
              'type' => '`$OBJECT`',
              'req' => true,
            ],
            [
              'name' => 'preview_key',
              'title' => 'Preview Key',
              'type' => '`$STRING`',
              'short' => 'Preview Key',
            ],
            [
              'name' => 'published_at',
              'title' => 'Published At',
              'type' => '`$STRING`',
              'format' => 'date-time',
            ],
            [
              'name' => 'updated_at',
              'title' => 'Updated At',
              'type' => '`$STRING`',
              'format' => 'date-time',
            ],
          ],
          'name' => 'postmortem',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'postmortem',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'incident_id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'postmortem',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => [
                      'postmortem' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'incident_id',
                      'page_id',
                    ],
                  ],
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'postmortem',
                    'publish',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => [
                      'postmortem' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'publish',
                    'exist' => [
                      'incident_id',
                      'page_id',
                    ],
                  ],
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'postmortem',
                    'revert',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'revert',
                    'exist' => [
                      'incident_id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                '$.main.kit.entity.page',
                '$.main.kit.entity.incident',
              ],
            ],
          ],
        ],
        'status_embed_config' => [
          'fields' => [
            [
              'name' => 'incident_background_color',
              'title' => 'Incident Background Color',
              'type' => '`$STRING`',
              'short' => 'Color of status embed iframe background when displaying incident',
            ],
            [
              'name' => 'incident_text_color',
              'title' => 'Incident Text Color',
              'type' => '`$STRING`',
              'short' => 'Color of status embed iframe text when displaying incident',
            ],
            [
              'name' => 'maintenance_background_color',
              'title' => 'Maintenance Background Color',
              'type' => '`$STRING`',
              'short' => 'Color of status embed iframe background when displaying maintenance',
            ],
            [
              'name' => 'maintenance_text_color',
              'title' => 'Maintenance Text Color',
              'type' => '`$STRING`',
              'short' => 'Color of status embed iframe text when displaying maintenance',
            ],
            [
              'name' => 'page_id',
              'title' => 'Page Id',
              'type' => '`$STRING`',
              'short' => 'Page identifier',
            ],
            [
              'name' => 'position',
              'title' => 'Position',
              'type' => '`$STRING`',
              'short' => 'Corner where status embed iframe will appear on page',
            ],
            [
              'name' => 'status_embed_config',
              'title' => 'Status Embed Config',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'status_embed_config',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'status_embed_config',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => [
                      'status_embed_config' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'status_embed_config',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => [
                      'status_embed_config' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                '$.main.kit.entity.page',
              ],
            ],
          ],
        ],
        'subscriber' => [
          'fields' => [
            [
              'name' => 'component_ids',
              'title' => 'Component Ids',
              'type' => '`$ARRAY`',
              'short' => 'A list of component ids for which the subscriber should recieve updates for.',
            ],
            [
              'name' => 'components',
              'title' => 'Components',
              'type' => '`$STRING`',
              'short' => 'The components for which the subscriber has elected to receive updates.',
            ],
            [
              'name' => 'created_at',
              'title' => 'Created At',
              'type' => '`$STRING`',
              'format' => 'date-time',
            ],
            [
              'name' => 'display_phone_number',
              'title' => 'Display Phone Number',
              'type' => '`$STRING`',
              'short' => 'A formatted version of the phone_number and phone_country pair, nicely formatted for display.',
            ],
            [
              'name' => 'email',
              'title' => 'Email',
              'type' => '`$STRING`',
              'short' => 'The email address to use to contact the subscriber.',
            ],
            [
              'name' => 'endpoint',
              'title' => 'Endpoint',
              'type' => '`$STRING`',
              'short' => 'The URL where a webhook subscriber elects to receive updates.',
            ],
            [
              'name' => 'id',
              'title' => 'Id',
              'type' => '`$STRING`',
              'short' => 'Subscriber Identifier',
            ],
            [
              'name' => 'mode',
              'title' => 'Mode',
              'type' => '`$STRING`',
              'short' => 'The communication mode of the subscriber.',
            ],
            [
              'name' => 'obfuscated_channel_name',
              'title' => 'Obfuscated Channel Name',
              'type' => '`$STRING`',
              'short' => 'Obfuscated slack channel name',
            ],
            [
              'name' => 'page_access_user_id',
              'title' => 'Page Access User Id',
              'type' => '`$STRING`',
              'short' => 'The Page Access user this subscriber belongs to (only for audience-specific pages).',
            ],
            [
              'name' => 'phone_country',
              'title' => 'Phone Country',
              'type' => '`$STRING`',
              'short' => 'The two-character country code representing the country of which the phone_number is a part.',
            ],
            [
              'name' => 'phone_number',
              'title' => 'Phone Number',
              'type' => '`$STRING`',
              'short' => 'The phone number used to contact an SMS subscriber',
            ],
            [
              'name' => 'purge_at',
              'title' => 'Purge At',
              'type' => '`$STRING`',
              'short' => 'The timestamp when a quarantined subscriber will be purged (unsubscribed).',
              'format' => 'date-time',
            ],
            [
              'name' => 'quarantined_at',
              'title' => 'Quarantined At',
              'type' => '`$STRING`',
              'short' => 'The timestamp when the subscriber was quarantined due to an issue reaching them.',
              'format' => 'date-time',
            ],
            [
              'name' => 'skip_confirmation_notification',
              'title' => 'Skip Confirmation Notification',
              'type' => '`$BOOLEAN`',
              'short' => 'If this is true, do not notify the user with changes to their subscription.',
            ],
            [
              'name' => 'subscriber',
              'title' => 'Subscriber',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'workspace_name',
              'title' => 'Workspace Name',
              'type' => '`$STRING`',
              'short' => 'The workspace name of the slack subscriber.',
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
                      'var' => 'id',
                    ],
                    [
                      'lit' => 'resend_confirmation',
                    ],
                  ],
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'subscribers',
                    '{id}',
                    'resend_confirmation',
                  ],
                  'rename' => [
                    'param' => [
                      'subscriber_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'subscriber_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'resend_confirmation',
                    'exist' => [
                      'id',
                      'incident_id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/subscribers/{subscriber_id}/resend_confirmation',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    '{id}',
                    'resend_confirmation',
                  ],
                  'rename' => [
                    'param' => [
                      'subscriber_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'subscriber_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'resend_confirmation',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'subscribers',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => [
                      'subscriber' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'incident_id',
                      'page_id',
                    ],
                  ],
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => [
                      'subscriber' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    'reactivate',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'reactivate',
                    'exist' => [
                      'page_id',
                    ],
                  ],
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    'resend_confirmation',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'resend_confirmation',
                    'exist' => [
                      'page_id',
                    ],
                  ],
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    'unsubscribe',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'unsubscribe',
                    'exist' => [
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'limit',
                        'orig' => 'limit',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                        'example' => 0,
                      ],
                      [
                        'name' => 'q',
                        'orig' => 'q',
                        'type' => '`$STRING`',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'sort_direction',
                        'orig' => 'sort_direction',
                        'type' => '`$STRING`',
                        'kind' => 'query',
                        'example' => 'asc',
                      ],
                      [
                        'name' => 'sort_field',
                        'orig' => 'sort_field',
                        'type' => '`$STRING`',
                        'kind' => 'query',
                        'example' => 'primary',
                      ],
                      [
                        'name' => 'state',
                        'orig' => 'state',
                        'type' => '`$STRING`',
                        'kind' => 'query',
                        'example' => 'active',
                      ],
                      [
                        'name' => 'type',
                        'orig' => 'type',
                        'type' => '`$STRING`',
                        'kind' => 'query',
                      ],
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
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'subscribers',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
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
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    'unsubscribed',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
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
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'subscribers',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'subscriber_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'subscriber_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'incident_id',
                      'page_id',
                    ],
                  ],
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    'count',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'state',
                        'orig' => 'state',
                        'type' => '`$STRING`',
                        'kind' => 'query',
                        'example' => 'active',
                      ],
                      [
                        'name' => 'type',
                        'orig' => 'type',
                        'type' => '`$STRING`',
                        'kind' => 'query',
                      ],
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
                ],
                [
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/subscribers/{subscriber_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'subscriber_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'subscriber_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
                [
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    'histogram_by_state',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    '$action' => 'histogram_by_state',
                    'exist' => [
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'subscribers',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'subscriber_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'subscriber_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'incident_id',
                      'page_id',
                    ],
                  ],
                ],
                [
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/subscribers/{subscriber_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'subscriber_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'subscriber_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'skip_unsubscription_notification',
                        'orig' => 'skip_unsubscription_notification',
                        'type' => '`$BOOLEAN`',
                        'kind' => 'query',
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                      'skip_unsubscription_notification',
                    ],
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/subscribers/{subscriber_id}',
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
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'subscriber_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'subscriber_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                '$.main.kit.entity.page',
              ],
              [
                '$.main.kit.entity.page',
                '$.main.kit.entity.incident',
              ],
            ],
          ],
        ],
        'user' => [
          'fields' => [
            [
              'name' => 'created_at',
              'title' => 'Created At',
              'type' => '`$STRING`',
              'format' => 'date-time',
            ],
            [
              'name' => 'email',
              'title' => 'Email',
              'type' => '`$STRING`',
              'short' => 'Email address for the team member',
            ],
            [
              'name' => 'first_name',
              'title' => 'First Name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'title' => 'Id',
              'type' => '`$STRING`',
              'short' => 'User identifier',
            ],
            [
              'name' => 'last_name',
              'title' => 'Last Name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'organization_id',
              'title' => 'Organization Id',
              'type' => '`$STRING`',
              'short' => 'Organization identifier',
            ],
            [
              'name' => 'updated_at',
              'title' => 'Updated At',
              'type' => '`$STRING`',
              'format' => 'date-time',
            ],
            [
              'name' => 'user',
              'title' => 'User',
              'type' => '`$OBJECT`',
              'req' => true,
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
                  'parts' => [
                    'organizations',
                    '{organization_id}',
                    'users',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => [
                      'user' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'organization_id',
                        'orig' => 'organization_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'organization_id',
                    ],
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
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
                  'parts' => [
                    'organizations',
                    '{organization_id}',
                    'users',
                  ],
                  'rename' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'organization_id',
                        'orig' => 'organization_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                    'query' => [
                      [
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
                      [
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                        'kind' => 'query',
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'organization_id',
                      'page',
                      'per_page',
                    ],
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/organizations/{organization_id}/users/{user_id}',
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
                  'parts' => [
                    'organizations',
                    '{organization_id}',
                    'users',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'user_id' => 'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'args' => [
                    'params' => [
                      [
                        'name' => 'id',
                        'orig' => 'user_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                      [
                        'name' => 'organization_id',
                        'orig' => 'organization_id',
                        'type' => '`$STRING`',
                        'kind' => 'param',
                        'reqd' => true,
                      ],
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'organization_id',
                    ],
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
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
