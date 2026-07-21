<?php
declare(strict_types=1);

// Statuspage SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class StatuspageFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new StatuspageBaseFeature();
            case "test":
                return new StatuspageTestFeature();
            default:
                return new StatuspageBaseFeature();
        }
    }
}
