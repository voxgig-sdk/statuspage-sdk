<?php
declare(strict_types=1);

// Statuspage SDK base feature

class StatuspageBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(StatuspageContext $ctx, array $options): void {}
    public function PostConstruct(StatuspageContext $ctx): void {}
    public function PostConstructEntity(StatuspageContext $ctx): void {}
    public function SetData(StatuspageContext $ctx): void {}
    public function GetData(StatuspageContext $ctx): void {}
    public function GetMatch(StatuspageContext $ctx): void {}
    public function SetMatch(StatuspageContext $ctx): void {}
    public function PrePoint(StatuspageContext $ctx): void {}
    public function PreSpec(StatuspageContext $ctx): void {}
    public function PreRequest(StatuspageContext $ctx): void {}
    public function PreResponse(StatuspageContext $ctx): void {}
    public function PreResult(StatuspageContext $ctx): void {}
    public function PreDone(StatuspageContext $ctx): void {}
    public function PreUnexpected(StatuspageContext $ctx): void {}
}
