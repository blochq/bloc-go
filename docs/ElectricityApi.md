# {{classname}}

All URIs are relative to *https://api.blochq.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Customervalidation**](ElectricityApi.md#Customervalidation) | **Get** /bills/electricity/customers/{operatorID}/{productNumber} | Customervalidation
[**Getoperatorproducts1**](ElectricityApi.md#Getoperatorproducts1) | **Get** /bills/electricity/operators/{operatorID}/products | Getoperatorproducts1
[**Getsupportedoperators1**](ElectricityApi.md#Getsupportedoperators1) | **Get** /bills/electricity/operators | Getsupportedoperators1
[**Makeelectricitypayment**](ElectricityApi.md#Makeelectricitypayment) | **Post** /bills/electricity/payment | Makeelectricitypayment

# **Customervalidation**
> InlineResponse20052 Customervalidation(ctx, meterType, operatorID, productNumber)
Customervalidation

Customer validation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **meterType** | **string**|  | 
  **operatorID** | **string**|  | 
  **productNumber** | **string**|  | 

### Return type

[**InlineResponse20052**](inline_response_200_52.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getoperatorproducts1**
> InlineResponse20051 Getoperatorproducts1(ctx, operatorID)
Getoperatorproducts1

Get operator products

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorID** | **string**|  | 

### Return type

[**InlineResponse20051**](inline_response_200_51.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getsupportedoperators1**
> InlineResponse20011 Getsupportedoperators1(ctx, )
Getsupportedoperators1

Get supported operators

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Makeelectricitypayment**
> InlineResponse20053 Makeelectricitypayment(ctx, body)
Makeelectricitypayment

Make electricity payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ElectricityPaymentBody**](ElectricityPaymentBody.md)|  | 

### Return type

[**InlineResponse20053**](inline_response_200_53.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

