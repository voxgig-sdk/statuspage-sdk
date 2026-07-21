<?php
declare(strict_types=1);

// Statuspage SDK utility: result_headers

class StatuspageResultHeaders
{
    public static function call(StatuspageContext $ctx): ?StatuspageResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
