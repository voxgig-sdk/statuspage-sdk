<?php
declare(strict_types=1);

// Statuspage SDK utility registration

require_once __DIR__ . '/../core/UtilityType.php';
require_once __DIR__ . '/Clean.php';
require_once __DIR__ . '/Done.php';
require_once __DIR__ . '/MakeError.php';
require_once __DIR__ . '/FeatureAdd.php';
require_once __DIR__ . '/FeatureHook.php';
require_once __DIR__ . '/FeatureInit.php';
require_once __DIR__ . '/Fetcher.php';
require_once __DIR__ . '/MakeFetchDef.php';
require_once __DIR__ . '/MakeContext.php';
require_once __DIR__ . '/MakeOptions.php';
require_once __DIR__ . '/MakeRequest.php';
require_once __DIR__ . '/MakeResponse.php';
require_once __DIR__ . '/MakeResult.php';
require_once __DIR__ . '/MakePoint.php';
require_once __DIR__ . '/MakeSpec.php';
require_once __DIR__ . '/MakeUrl.php';
require_once __DIR__ . '/Param.php';
require_once __DIR__ . '/PrepareAuth.php';
require_once __DIR__ . '/PrepareBody.php';
require_once __DIR__ . '/PrepareHeaders.php';
require_once __DIR__ . '/PrepareMethod.php';
require_once __DIR__ . '/PrepareParams.php';
require_once __DIR__ . '/PreparePath.php';
require_once __DIR__ . '/PrepareQuery.php';
require_once __DIR__ . '/ResultBasic.php';
require_once __DIR__ . '/ResultBody.php';
require_once __DIR__ . '/ResultHeaders.php';
require_once __DIR__ . '/TransformRequest.php';
require_once __DIR__ . '/TransformResponse.php';

StatuspageUtility::setRegistrar(function (StatuspageUtility $u): void {
    $u->clean = [StatuspageClean::class, 'call'];
    $u->done = [StatuspageDone::class, 'call'];
    $u->make_error = [StatuspageMakeError::class, 'call'];
    $u->feature_add = [StatuspageFeatureAdd::class, 'call'];
    $u->feature_hook = [StatuspageFeatureHook::class, 'call'];
    $u->feature_init = [StatuspageFeatureInit::class, 'call'];
    $u->fetcher = [StatuspageFetcher::class, 'call'];
    $u->make_fetch_def = [StatuspageMakeFetchDef::class, 'call'];
    $u->make_context = [StatuspageMakeContext::class, 'call'];
    $u->make_options = [StatuspageMakeOptions::class, 'call'];
    $u->make_request = [StatuspageMakeRequest::class, 'call'];
    $u->make_response = [StatuspageMakeResponse::class, 'call'];
    $u->make_result = [StatuspageMakeResult::class, 'call'];
    $u->make_point = [StatuspageMakePoint::class, 'call'];
    $u->make_spec = [StatuspageMakeSpec::class, 'call'];
    $u->make_url = [StatuspageMakeUrl::class, 'call'];
    $u->param = [StatuspageParam::class, 'call'];
    $u->prepare_auth = [StatuspagePrepareAuth::class, 'call'];
    $u->prepare_body = [StatuspagePrepareBody::class, 'call'];
    $u->prepare_headers = [StatuspagePrepareHeaders::class, 'call'];
    $u->prepare_method = [StatuspagePrepareMethod::class, 'call'];
    $u->prepare_params = [StatuspagePrepareParams::class, 'call'];
    $u->prepare_path = [StatuspagePreparePath::class, 'call'];
    $u->prepare_query = [StatuspagePrepareQuery::class, 'call'];
    $u->result_basic = [StatuspageResultBasic::class, 'call'];
    $u->result_body = [StatuspageResultBody::class, 'call'];
    $u->result_headers = [StatuspageResultHeaders::class, 'call'];
    $u->transform_request = [StatuspageTransformRequest::class, 'call'];
    $u->transform_response = [StatuspageTransformResponse::class, 'call'];
});
