# {{classname}}

All URIs are relative to *https://api.blochq.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Customerdevicevalidation**](BillsPaymentsV11Api.md#Customerdevicevalidation) | **Get** /bills/customer/validate/{operatorID} | Customerdevicevalidation
[**Getoperatorproducts**](BillsPaymentsV11Api.md#Getoperatorproducts) | **Get** /bills/operators/{operatorID}/products | Getoperatorproducts
[**Getsupportedbills**](BillsPaymentsV11Api.md#Getsupportedbills) | **Get** /bills/supported | Getsupportedbills
[**Getsupportedoperators**](BillsPaymentsV11Api.md#Getsupportedoperators) | **Get** /bills/operators | Getsupportedoperators
[**MakePayment**](BillsPaymentsV11Api.md#MakePayment) | **Post** /bills/payment | MakePayment

# **Customerdevicevalidation**
> InlineResponse20013 Customerdevicevalidation(ctx, meterType, bill, deviceNumber, operatorID)
Customerdevicevalidation

Customer device validation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **meterType** | **string**| prepaid or postpaid | 
  **bill** | **string**|  | 
  **deviceNumber** | **string**|  | 
  **operatorID** | **string**|  | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getoperatorproducts**
> InlineResponse20012 Getoperatorproducts(ctx, bill, operatorID)
Getoperatorproducts

Get operator products

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bill** | **string**|  | 
  **operatorID** | **string**| Operator Id of the selected operator | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getsupportedbills**
> InlineResponse20010 Getsupportedbills(ctx, )
Getsupportedbills

Get supported bills

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getsupportedoperators**
> InlineResponse20011 Getsupportedoperators(ctx, bill)
Getsupportedoperators

Get supported operators

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bill** | **string**|  | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MakePayment**
> InlineResponse20014 MakePayment(ctx, body, bill)
MakePayment

Make Payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BillsPaymentBody**](BillsPaymentBody.md)|  | 
  **bill** | **string**|  | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

