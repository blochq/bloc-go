# {{classname}}

All URIs are relative to *https://api.blochq.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseAccount**](AccountsApi.md#CloseAccount) | **Put** /accounts/{accountID}/close | CloseAccount
[**CreateCollectionAccount**](AccountsApi.md#CreateCollectionAccount) | **Post** /accounts/collections | CreateCollectionAccount
[**CreateFixedAccount**](AccountsApi.md#CreateFixedAccount) | **Post** /accounts | CreateFixedAccount
[**FreezeAccount**](AccountsApi.md#FreezeAccount) | **Put** /accounts/{accountID}/freeze | FreezeAccount
[**GetAccountbyaccountnumber**](AccountsApi.md#GetAccountbyaccountnumber) | **Get** /v1/accounts/number/{accountNumber} | GetAccountbyaccountnumber
[**GetAccountbyid**](AccountsApi.md#GetAccountbyid) | **Get** /accounts/{accountID} | GetAccountbyid
[**GetOrganisationDefaultAccounts**](AccountsApi.md#GetOrganisationDefaultAccounts) | **Get** /accounts/organization/default | GetOrganisationDefaultAccounts
[**Getaccounts**](AccountsApi.md#Getaccounts) | **Get** /accounts | Getaccounts
[**Getcustomeraccounts**](AccountsApi.md#Getcustomeraccounts) | **Get** /accounts/customers/accounts/{customerID} | Getcustomeraccounts
[**ReOpenAccount**](AccountsApi.md#ReOpenAccount) | **Put** /accounts/{accountID}/reopen | ReOpenAccount
[**UnfreezeAccount**](AccountsApi.md#UnfreezeAccount) | **Put** /accounts/{accountID}/unfreeze | UnfreezeAccount

# **CloseAccount**
> InlineResponse2008 CloseAccount(ctx, body, accountID)
CloseAccount

Close Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountIdCloseBody**](AccountIdCloseBody.md)|  | 
  **accountID** | **string**|  | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCollectionAccount**
> InlineResponse201 CreateCollectionAccount(ctx, body)
CreateCollectionAccount

Create Collection Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountsCollectionsBody**](AccountsCollectionsBody.md)|  | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFixedAccount**
> InlineResponse2001 CreateFixedAccount(ctx, body)
CreateFixedAccount

Create Fixed Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountsBody**](AccountsBody.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FreezeAccount**
> InlineResponse2006 FreezeAccount(ctx, body, accountID)
FreezeAccount

Freeze Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountIdFreezeBody**](AccountIdFreezeBody.md)|  | 
  **accountID** | **string**|  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountbyaccountnumber**
> InlineResponse2003 GetAccountbyaccountnumber(ctx, accountNumber)
GetAccountbyaccountnumber

Get Account by account number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountNumber** | **string**|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountbyid**
> InlineResponse2002 GetAccountbyid(ctx, accountID)
GetAccountbyid

Get Account by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountID** | **string**|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrganisationDefaultAccounts**
> InlineResponse2005 GetOrganisationDefaultAccounts(ctx, )
GetOrganisationDefaultAccounts

Get Organisation Default Accounts

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getaccounts**
> InlineResponse200 Getaccounts(ctx, )
Getaccounts

Get accounts

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcustomeraccounts**
> InlineResponse2004 Getcustomeraccounts(ctx, customerID)
Getcustomeraccounts

Get customer accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerID** | **string**|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReOpenAccount**
> InlineResponse2009 ReOpenAccount(ctx, body, accountID)
ReOpenAccount

ReOpen Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountIdReopenBody**](AccountIdReopenBody.md)|  | 
  **accountID** | **string**|  | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnfreezeAccount**
> InlineResponse2007 UnfreezeAccount(ctx, body, accountID)
UnfreezeAccount

Unfreeze Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountIdUnfreezeBody**](AccountIdUnfreezeBody.md)|  | 
  **accountID** | **string**|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

