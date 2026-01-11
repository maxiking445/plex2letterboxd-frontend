import {
  Configuration,
  DataApi,
  ModelsExport,
  ModelsSettings,
  SettingsApi,
} from "@/generated/api";
import { useToast } from "vue-toastification";

const toast = useToast();

const configuration = new Configuration({
  basePath: "/api",
});

const dataAPI = new DataApi(configuration);

const settingsAPI = new SettingsApi(configuration);

export default {
  loadExport,
  removeExport,
  startExport,
};

export async function startExport(): Promise<string> {
  try {
    return await dataAPI.executeScriptGet();
  } catch (err) {
    console.error("startExport API failed:", err);
    toast.error("Export failed!");
    throw err;
  }
}

export async function saveSettings(model: ModelsSettings): Promise<string> {
  try {
    return await settingsAPI.settingsSavePost({ settings: model });
  } catch (err) {
    console.error("saveSettings API failed:", err);
    toast.error("Saving settings failed!");
    throw err;
  }
}

export async function loadExports(): Promise<ModelsExport[]> {
  try {
    return await dataAPI.exportsGet();
  } catch (err) {
    console.error("loadExports API failed:", err);
    toast.error("Loading exports failed!");
    throw err;
  }
}

export async function loadExport(id: string): Promise<Blob> {
  try {
    return await dataAPI.exportsIdGet({ id });
  } catch (err) {
    console.error(`loadExport API failed for id: ${id}`, err);
    toast.error(`Downloading export ${id} failed!`);
    throw err;
  }
}

export async function getSettings(): Promise<ModelsSettings> {
  try {
    return await settingsAPI.settingsGet();
  } catch (err) {
    console.error("getSettings API failed:", err);
    toast.error("Loading settings failed!");
    throw err;
  }
}

export async function removeExport(id: string): Promise<string> {
  try {
    return await dataAPI.exportsRemoveIdDelete({ id });
  } catch (err) {
    console.error("removeExport API failed:", err);
    toast.error("Removing Export failed!");
    throw err;
  }
}
