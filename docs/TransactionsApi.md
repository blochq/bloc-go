# {{classname}}

All URIs are relative to *https://api.blochq.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllTransactions**](TransactionsApi.md#GetAllTransactions) | **Get** /transactions | GetAllTransactions
[**GetTransactionbyReference**](TransactionsApi.md#GetTransactionbyReference) | **Get** /transactions/reference/{reference} | GetTransactionbyReference

# **GetAllTransactions**
> AllTransactions GetAllTransactions(ctx, perPage, startDate, endDate, status, drcr, paymentMethod, source, limit, cardId)
GetAllTransactions

Get All Transactions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **perPage** | **string**|  | 
  **startDate** | **string**|  | 
  **endDate** | **string**|  | 
  **status** | **string**|  | 
  **drcr** | **string**|  | 
  **paymentMethod** | **string**|  | 
  **source** | **string**|  | 
  **limit** | **string**|  | 
  **cardId** | **string**|  | 

### Return type

[**AllTransactions**](AllTransactions.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionbyReference**
> InlineResponse20044 GetTransactionbyReference(ctx, reference)
GetTransactionbyReference

Get Transaction by Reference

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reference** | **string**|  | 

### Return type

[**InlineResponse20044**](inline_response_200_44.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

