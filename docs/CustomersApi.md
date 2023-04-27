# {{classname}}

All URIs are relative to *https://api.blochq.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Createcustomers**](CustomersApi.md#Createcustomers) | **Post** /customers | Createcustomers
[**Getcustomerbyid**](CustomersApi.md#Getcustomerbyid) | **Get** /customers/{customerID} | Getcustomerbyid
[**Getcustomers**](CustomersApi.md#Getcustomers) | **Get** /customers | Getcustomers
[**MeansofIdentification**](CustomersApi.md#MeansofIdentification) | **Get** /customers/means-of-id | MeansofIdentification
[**RevalidateCustomerKYC**](CustomersApi.md#RevalidateCustomerKYC) | **Get** /customers/kyc/revalidate/{customerID} | RevalidateCustomerKYC
[**Updatecustomer**](CustomersApi.md#Updatecustomer) | **Put** /customers/{customerID} | Updatecustomer
[**UpgradeCustomerToKYCT1**](CustomersApi.md#UpgradeCustomerToKYCT1) | **Put** /customers/upgrade/t1/{customerID} | UpgradeCustomerToKYCT1
[**UpgradeCustomerToKYCT2**](CustomersApi.md#UpgradeCustomerToKYCT2) | **Put** /customers/upgrade/t2/{customerID} | UpgradeCustomerToKYCT2
[**UpgradeCustomerToKYCT3**](CustomersApi.md#UpgradeCustomerToKYCT3) | **Put** /customers/upgrade/t3/{customerID} | UpgradeCustomerToKYCT3

# **Createcustomers**
> InlineResponse20032 Createcustomers(ctx, body)
Createcustomers

Create customers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomersBody**](CustomersBody.md)|  | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcustomerbyid**
> InlineResponse20036 Getcustomerbyid(ctx, customerID)
Getcustomerbyid

Get customer by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerID** | **string**|  | 

### Return type

[**InlineResponse20036**](inline_response_200_36.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getcustomers**
> InlineResponse20031 Getcustomers(ctx, )
Getcustomers

Get customers

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20031**](inline_response_200_31.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MeansofIdentification**
> InlineResponse20038 MeansofIdentification(ctx, )
MeansofIdentification

Means of Identification

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20038**](inline_response_200_38.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevalidateCustomerKYC**
> InlineResponse20039 RevalidateCustomerKYC(ctx, customerID)
RevalidateCustomerKYC

Revalidate Customer KYC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerID** | **string**|  | 

### Return type

[**InlineResponse20039**](inline_response_200_39.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Updatecustomer**
> InlineResponse20037 Updatecustomer(ctx, body, customerID)
Updatecustomer

Update customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomersCustomerIdBody**](CustomersCustomerIdBody.md)|  | 
  **customerID** | **string**|  | 

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeCustomerToKYCT1**
> InlineResponse20033 UpgradeCustomerToKYCT1(ctx, body, customerID)
UpgradeCustomerToKYCT1

Upgrade Customer To  KYC T1

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**T1CustomerIdBody**](T1CustomerIdBody.md)|  | 
  **customerID** | **string**|  | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeCustomerToKYCT2**
> InlineResponse20034 UpgradeCustomerToKYCT2(ctx, body, customerID)
UpgradeCustomerToKYCT2

Upgrade Customer To KYC T2

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**T2CustomerIdBody**](T2CustomerIdBody.md)|  | 
  **customerID** | **string**|  | 

### Return type

[**InlineResponse20034**](inline_response_200_34.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeCustomerToKYCT3**
> InlineResponse20035 UpgradeCustomerToKYCT3(ctx, body, customerID)
UpgradeCustomerToKYCT3

Upgrade Customer To KYC T3

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**T3CustomerIdBody**](T3CustomerIdBody.md)|  | 
  **customerID** | **string**|  | 

### Return type

[**InlineResponse20035**](inline_response_200_35.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

