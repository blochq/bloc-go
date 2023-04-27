# {{classname}}

All URIs are relative to *https://api.blochq.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCardDispute**](DisputesApi.md#CreateCardDispute) | **Post** /cards/dispute | CreateCardDispute
[**GetCardDispute**](DisputesApi.md#GetCardDispute) | **Get** /cards/dispute/{disputeID} | GetCardDispute
[**GetCardDisputeReasons**](DisputesApi.md#GetCardDisputeReasons) | **Get** /cards/dispute/reasons | GetCardDisputeReasons
[**GetCardDisputes**](DisputesApi.md#GetCardDisputes) | **Get** /cards/dispute | GetCardDisputes
[**UpdateCardDispute**](DisputesApi.md#UpdateCardDispute) | **Put** /cards/dispute/{disputeID} | UpdateCardDispute

# **CreateCardDispute**
> InlineResponse20017 CreateCardDispute(ctx, body)
CreateCardDispute

Create Card Dispute

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CardsDisputeBody**](CardsDisputeBody.md)|  | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardDispute**
> InlineResponse20018 GetCardDispute(ctx, disputeID)
GetCardDispute

Get Card Dispute

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **disputeID** | **string**|  | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardDisputeReasons**
> InlineResponse20015 GetCardDisputeReasons(ctx, )
GetCardDisputeReasons

Get Card Dispute Reasons

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardDisputes**
> InlineResponse20016 GetCardDisputes(ctx, )
GetCardDisputes

Get Card Disputes

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCardDispute**
> InlineResponse20018 UpdateCardDispute(ctx, body, disputeID)
UpdateCardDispute

Update Card Dispute

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DisputeDisputeIdBody**](DisputeDisputeIdBody.md)|  | 
  **disputeID** | **string**|  | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

