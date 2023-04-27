# {{classname}}

All URIs are relative to *https://api.blochq.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Createcheckout**](CheckoutApi.md#Createcheckout) | **Post** /checkout/new | createcheckout

# **Createcheckout**
> InlineResponse20030 Createcheckout(ctx, body)
createcheckout

create checkout

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CheckoutNewBody**](CheckoutNewBody.md)|  | 

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

