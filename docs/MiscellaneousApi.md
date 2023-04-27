# {{classname}}

All URIs are relative to *https://api.blochq.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getlistofbanks**](MiscellaneousApi.md#Getlistofbanks) | **Get** /banks | Getlistofbanks
[**ResolveAccount**](MiscellaneousApi.md#ResolveAccount) | **Get** /resolve-account | ResolveAccount

# **Getlistofbanks**
> InlineResponse20040 Getlistofbanks(ctx, )
Getlistofbanks

Get list of banks

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResolveAccount**
> InlineResponse20041 ResolveAccount(ctx, accountNumber, bankCode)
ResolveAccount

Resolve Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountNumber** | **string**|  | 
  **bankCode** | **string**|  | 

### Return type

[**InlineResponse20041**](inline_response_200_41.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

