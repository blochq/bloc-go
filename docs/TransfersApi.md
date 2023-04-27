# {{classname}}

All URIs are relative to *https://api.blochq.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkTransfer**](TransfersApi.md#BulkTransfer) | **Post** /transfers/bulk | BulkTransfer
[**Internaltransfer**](TransfersApi.md#Internaltransfer) | **Post** /transfers/internal | Internaltransfer
[**TransferFromAFixedAccount**](TransfersApi.md#TransferFromAFixedAccount) | **Post** /transfers | TransferFromAFixedAccount
[**TransferFromOrganizationBalance**](TransfersApi.md#TransferFromOrganizationBalance) | **Post** /transfers/balance | TransferFromOrganizationBalance

# **BulkTransfer**
> InlineResponse20048 BulkTransfer(ctx, body)
BulkTransfer

Bulk Transfer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransfersBulkBody**](TransfersBulkBody.md)|  | 

### Return type

[**InlineResponse20048**](inline_response_200_48.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Internaltransfer**
> InlineResponse20047 Internaltransfer(ctx, body)
Internaltransfer

Internal transfer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransfersInternalBody**](TransfersInternalBody.md)|  | 

### Return type

[**InlineResponse20047**](inline_response_200_47.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransferFromAFixedAccount**
> InlineResponse20045 TransferFromAFixedAccount(ctx, body)
TransferFromAFixedAccount

Transfer From A Fixed Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransfersBody**](TransfersBody.md)|  | 

### Return type

[**InlineResponse20045**](inline_response_200_45.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransferFromOrganizationBalance**
> InlineResponse20046 TransferFromOrganizationBalance(ctx, body)
TransferFromOrganizationBalance

Transfer From Organization Balance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransfersBalanceBody**](TransfersBalanceBody.md)|  | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

