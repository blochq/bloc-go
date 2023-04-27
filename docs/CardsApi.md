# {{classname}}

All URIs are relative to *https://api.blochq.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlockCard**](CardsApi.md#BlockCard) | **Put** /cards/block/{id} | BlockCard
[**ChangeCardPIN**](CardsApi.md#ChangeCardPIN) | **Put** /cards/change-pin/{cardID} | ChangeCardPIN
[**FreezeCard**](CardsApi.md#FreezeCard) | **Put** /cards/freeze/{cardID} | FreezeCard
[**FundCard**](CardsApi.md#FundCard) | **Post** /cards/fund/{cardID} | FundCard
[**GetCardSecureData**](CardsApi.md#GetCardSecureData) | **Get** /cards/secure-data/{cardID} | GetCardSecureData
[**GetCardbyId**](CardsApi.md#GetCardbyId) | **Get** /cards/{cardID} | GetCardbyId
[**GetCards**](CardsApi.md#GetCards) | **Get** /cards | GetCards
[**GetCustomerCards**](CardsApi.md#GetCustomerCards) | **Get** /cards/customer/{customerID} | GetCustomerCards
[**IssueVerveCardNGN**](CardsApi.md#IssueVerveCardNGN) | **Post** /cards | IssueVerveCard(NGN)
[**LinkFixedAccount**](CardsApi.md#LinkFixedAccount) | **Put** /cards/fixed-account/link | LinkFixedAccount
[**SimulateChargeCard**](CardsApi.md#SimulateChargeCard) | **Post** /cards/simulate/authorize | SimulateChargeCard
[**UnLinkFixedAccount**](CardsApi.md#UnLinkFixedAccount) | **Put** /cards/fixed-account/unlink/{cardID} | UnLinkFixedAccount
[**UnfreezeCard**](CardsApi.md#UnfreezeCard) | **Put** /cards/unfreeze/{id} | UnfreezeCard
[**WithdrawFromCard**](CardsApi.md#WithdrawFromCard) | **Post** /cards/withdraw/{cardID} | WithdrawFromCard

# **BlockCard**
> InlineResponse20026 BlockCard(ctx, body, id)
BlockCard

Block Card

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BlockIdBody**](BlockIdBody.md)|  | 
  **id** | **string**|  | 

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeCardPIN**
> InlineResponse20024 ChangeCardPIN(ctx, body, cardID)
ChangeCardPIN

Change Card PIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChangepinCardIdBody**](ChangepinCardIdBody.md)|  | 
  **cardID** | **string**|  | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FreezeCard**
> InlineResponse20025 FreezeCard(ctx, cardID)
FreezeCard

Freeze Card

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardID** | **string**|  | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FundCard**
> InlineResponse20027 FundCard(ctx, body, cardID)
FundCard

Fund Card

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FundCardIdBody**](FundCardIdBody.md)|  | 
  **cardID** | **string**|  | 

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardSecureData**
> InlineResponse20023 GetCardSecureData(ctx, cardID)
GetCardSecureData

Get Card Secure Data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardID** | **string**|  | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardbyId**
> InlineResponse20021 GetCardbyId(ctx, cardID)
GetCardbyId

Get Card by Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardID** | **string**|  | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCards**
> InlineResponse20019 GetCards(ctx, )
GetCards

Get Cards

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerCards**
> InlineResponse20022 GetCustomerCards(ctx, customerID)
GetCustomerCards

Get Customer Cards

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerID** | **string**|  | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IssueVerveCardNGN**
> InlineResponse20020 IssueVerveCardNGN(ctx, body)
IssueVerveCard(NGN)

Issue Verve Card (NGN)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CardsBody**](CardsBody.md)|  | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinkFixedAccount**
> InlineResponse20027 LinkFixedAccount(ctx, body)
LinkFixedAccount

Link Fixed Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FixedaccountLinkBody**](FixedaccountLinkBody.md)|  | 

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SimulateChargeCard**
> InlineResponse20029 SimulateChargeCard(ctx, body)
SimulateChargeCard

Simulate Charge Card

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SimulateAuthorizeBody**](SimulateAuthorizeBody.md)|  | 

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnLinkFixedAccount**
> InlineResponse20027 UnLinkFixedAccount(ctx, cardID)
UnLinkFixedAccount

UnLink Fixed Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardID** | **string**|  | 

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnfreezeCard**
> InlineResponse20026 UnfreezeCard(ctx, id)
UnfreezeCard

Unfreeze Card

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WithdrawFromCard**
> InlineResponse20028 WithdrawFromCard(ctx, body, cardID)
WithdrawFromCard

Withdraw From Card

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WithdrawCardIdBody**](WithdrawCardIdBody.md)|  | 
  **cardID** | **string**|  | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

