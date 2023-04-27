# {{classname}}

All URIs are relative to *https://api.blochq.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCustomerDetails**](TelevisionApi.md#GetCustomerDetails) | **Get** /bills/television/customers/{operatorID}/{productNumber} | GetCustomerDetails
[**Getoperatorproducts3**](TelevisionApi.md#Getoperatorproducts3) | **Get** /bills/television/operators/{operatorID}/products | Getoperatorproducts3
[**Getsupportedoperators3**](TelevisionApi.md#Getsupportedoperators3) | **Get** /bills/television/operators | Getsupportedoperators3
[**MakeTelevisionPayment**](TelevisionApi.md#MakeTelevisionPayment) | **Post** /bills/television/payment | MakeTelevisionPayment

# **GetCustomerDetails**
> InlineResponse20059 GetCustomerDetails(ctx, operatorID, productNumber)
GetCustomerDetails

Get Customer Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorID** | **string**|  | 
  **productNumber** | **int32**|  | 

### Return type

[**InlineResponse20059**](inline_response_200_59.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getoperatorproducts3**
> InlineResponse20058 Getoperatorproducts3(ctx, operatorID)
Getoperatorproducts3

Get operator products

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorID** | **string**|  | 

### Return type

[**InlineResponse20058**](inline_response_200_58.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getsupportedoperators3**
> InlineResponse20057 Getsupportedoperators3(ctx, )
Getsupportedoperators3

Get supported operators

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20057**](inline_response_200_57.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MakeTelevisionPayment**
> InlineResponse20060 MakeTelevisionPayment(ctx, body)
MakeTelevisionPayment

Make Television Payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TelevisionPaymentBody**](TelevisionPaymentBody.md)|  | 

### Return type

[**InlineResponse20060**](inline_response_200_60.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

