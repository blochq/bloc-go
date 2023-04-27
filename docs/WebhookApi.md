# {{classname}}

All URIs are relative to *https://api.blochq.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebhook**](WebhookApi.md#GetWebhook) | **Get** /webhooks | GetWebhook
[**SetWebhook**](WebhookApi.md#SetWebhook) | **Post** /webhooks | SetWebhook

# **GetWebhook**
> InlineResponse20049 GetWebhook(ctx, )
GetWebhook

Get Webhook

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20049**](inline_response_200_49.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetWebhook**
> InlineResponse20050 SetWebhook(ctx, body)
SetWebhook

Set Webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebhooksBody**](WebhooksBody.md)|  | 

### Return type

[**InlineResponse20050**](inline_response_200_50.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

