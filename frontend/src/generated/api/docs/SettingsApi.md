# SettingsApi

All URIs are relative to */api*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**settingsGet**](SettingsApi.md#settingsget) | **GET** /settings | returns current setings |
| [**settingsSavePost**](SettingsApi.md#settingssavepost) | **POST** /settings/save | saves settings |



## settingsGet

> ModelsSettings settingsGet()

returns current setings

### Example

```ts
import {
  Configuration,
  SettingsApi,
} from '';
import type { SettingsGetRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new SettingsApi();

  try {
    const data = await api.settingsGet();
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

[**ModelsSettings**](ModelsSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successfully returned settings |  -  |
| **400** | Settings may missing or are incomplete |  -  |
| **500** | Unknown error during execution |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## settingsSavePost

> string settingsSavePost(settings)

saves settings

Receives baseurl and token and stores them in config.ini

### Example

```ts
import {
  Configuration,
  SettingsApi,
} from '';
import type { SettingsSavePostRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new SettingsApi();

  const body = {
    // ModelsSettings | Settings payload
    settings: ...,
  } satisfies SettingsSavePostRequest;

  try {
    const data = await api.settingsSavePost(body);
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
| **settings** | [ModelsSettings](ModelsSettings.md) | Settings payload | |

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successfully saved settings |  -  |
| **400** | Settings may be missing or are incomplete |  -  |
| **500** | Unknown error during execution |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

