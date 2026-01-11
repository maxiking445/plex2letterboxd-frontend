# DataApi

All URIs are relative to */api*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**executeScriptGet**](DataApi.md#executescriptget) | **GET** /executeScript | executes script |
| [**exportsGet**](DataApi.md#exportsget) | **GET** /exports | Returns all existing exports |
| [**exportsIdGet**](DataApi.md#exportsidget) | **GET** /exports/{id} | returns CSV of export |
| [**exportsRemoveIdDelete**](DataApi.md#exportsremoveiddelete) | **DELETE** /exports/remove/{id} | Deletes a specific export |



## executeScriptGet

> string executeScriptGet()

executes script

executes plex2letterbox python script with current settings

### Example

```ts
import {
  Configuration,
  DataApi,
} from '';
import type { ExecuteScriptGetRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new DataApi();

  try {
    const data = await api.executeScriptGet();
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters

This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Sucessfully executed export |  -  |
| **400** | Settings may missing or are incomplete |  -  |
| **500** | Unknown error during execution |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## exportsGet

> Array&lt;ModelsExport&gt; exportsGet()

Returns all existing exports

Returns all existing exports under /data

### Example

```ts
import {
  Configuration,
  DataApi,
} from '';
import type { ExportsGetRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new DataApi();

  try {
    const data = await api.exportsGet();
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**Array&lt;ModelsExport&gt;**](ModelsExport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successfully retrieved exports |  -  |
| **500** | Unknown error during execution |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## exportsIdGet

> Blob exportsIdGet(id)

returns CSV of export

reads CSV

### Example

```ts
import {
  Configuration,
  DataApi,
} from '';
import type { ExportsIdGetRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new DataApi();

  const body = {
    // string | Export ID (Filename)
    id: id_example,
  } satisfies ExportsIdGetRequest;

  try {
    const data = await api.exportsIdGet(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | `string` | Export ID (Filename) | [Defaults to `undefined`] |

### Return type

**Blob**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/octet-stream`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | CSV File |  -  |
| **404** | File not found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## exportsRemoveIdDelete

> string exportsRemoveIdDelete(id)

Deletes a specific export

Deletes the CSV export file with the given ID

### Example

```ts
import {
  Configuration,
  DataApi,
} from '';
import type { ExportsRemoveIdDeleteRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new DataApi();

  const body = {
    // string | Export ID (Filename)
    id: id_example,
  } satisfies ExportsRemoveIdDeleteRequest;

  try {
    const data = await api.exportsRemoveIdDelete(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | `string` | Export ID (Filename) | [Defaults to `undefined`] |

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Export successfully deleted |  -  |
| **404** | File not found |  -  |
| **500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

