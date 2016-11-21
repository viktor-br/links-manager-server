# \DefaultApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserLoginPost**](DefaultApi.md#UserLoginPost) | **Post** /user/login | Authentication
[**UserPut**](DefaultApi.md#UserPut) | **Put** /user | User creation


# **UserLoginPost**
> UserLoginPost($body)

Authentication

Endpoint verifies user credentials and returns authentication token.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserAuth**](UserAuth.md)| Body JSON | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserPut**
> UserPut($body, $xAuthToken)

User creation

Endpoint creates new user


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**User**](User.md)| body JSON | 
 **xAuthToken** | **string**| Authentication token | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

