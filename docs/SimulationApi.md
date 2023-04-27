# {{classname}}

All URIs are relative to *https://api.blochq.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Creditaccount**](SimulationApi.md#Creditaccount) | **Post** /accounts/credit/manual | Creditaccount
[**Debitaccount**](SimulationApi.md#Debitaccount) | **Post** /accounts/debit/manual | Debitaccount

# **Creditaccount**
> InlineResponse20042 Creditaccount(ctx, body)
Creditaccount

Credit account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreditManualBody**](CreditManualBody.md)|  | 

### Return type

[**InlineResponse20042**](inline_response_200_42.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Debitaccount**
> InlineResponse20043 Debitaccount(ctx, body)
Debitaccount

Debit account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DebitManualBody**](DebitManualBody.md)|  | 

### Return type

[**InlineResponse20043**](inline_response_200_43.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

