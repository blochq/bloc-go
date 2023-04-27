# {{classname}}

All URIs are relative to *https://api.blochq.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getoperatorproducts2**](TelecommunicationApi.md#Getoperatorproducts2) | **Get** /bills/telco/operators/{operatorID}/products | Getoperatorproducts2
[**Getsupportedoperators2**](TelecommunicationApi.md#Getsupportedoperators2) | **Get** /bills/telco/operators | Getsupportedoperators2
[**PostMakePayment**](TelecommunicationApi.md#PostMakePayment) | **Post** /bills/telco/payment | PostMakePayment

# **Getoperatorproducts2**
> InlineResponse20055 Getoperatorproducts2(ctx, operatorID)
Getoperatorproducts2

Get operator products

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorID** | **string**|  | 

### Return type

[**InlineResponse20055**](inline_response_200_55.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getsupportedoperators2**
> InlineResponse20054 Getsupportedoperators2(ctx, )
Getsupportedoperators2

Get supported operators

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20054**](inline_response_200_54.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMakePayment**
> InlineResponse20056 PostMakePayment(ctx, body)
PostMakePayment

Make Payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TelcoPaymentBody**](TelcoPaymentBody.md)|  | 

### Return type

[**InlineResponse20056**](inline_response_200_56.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

