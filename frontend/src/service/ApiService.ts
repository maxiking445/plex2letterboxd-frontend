import {
  Configuration,
  DataApi,
  ModelsSettings,
  SettingsApi,
} from "@/generated/api";

const configuration = new Configuration({
  basePath: "/api",
});

const dataAPI = new DataApi(configuration);

const settingsAPI = new SettingsApi(configuration);

export async function startExport(): Promise<string> {
  return await dataAPI.executeScriptGet();
}

export async function saveSettings(model: ModelsSettings): Promise<string> {
  return await settingsAPI.settingsSavePost({ settings: model });
}


export async function getSettings(): Promise<ModelsSettings> {
  return await settingsAPI.settingsGet();
}