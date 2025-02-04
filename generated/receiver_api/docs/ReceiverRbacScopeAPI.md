# \ReceiverRbacScopeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IngestScopeRBAC**](ReceiverRbacScopeAPI.md#IngestScopeRBAC) | **Post** /stsAgent/rbac/scope | Create scope RBAC objects



## IngestScopeRBAC

> IngestScopeRBAC(ctx).RBACRequest(rBACRequest).Execute()

Create scope RBAC objects



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    rBACRequest := openapiclient.RBACRequest{RBACIncrementRequest: openapiclient.NewRBACIncrementRequest("Type_example", int64(123), "Cluster_example", []openapiclient.RbacDataChanges{openapiclient.RbacDataChanges{CreateRbacData: openapiclient.NewCreateRbacData("Type_example", openapiclient.RbacData{ClusterRole: openapiclient.NewClusterRole("Kind_example", *openapiclient.NewObjectMeta("Uid_example", "Name_example"))})}})} // RBACRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceiverRbacScopeAPI.IngestScopeRBAC(context.Background()).RBACRequest(rBACRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiverRbacScopeAPI.IngestScopeRBAC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestScopeRBACRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rBACRequest** | [**RBACRequest**](RBACRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

