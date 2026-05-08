<template>
  <Transition name="modal">
    <div v-if="visible" class="modal-overlay" @click.self="close">
      <div class="modal">
        <h3 class="modal-title">Settings</h3>

        <div class="form-group">
          <label>Plex Base URL</label>
          <input
            type="text"
            v-model="localSettings.baseurl"
            placeholder="http://localhost:32400"
          />
        </div>

        <div class="form-group">
          <label>Libraries</label>
          <div class="token-input-wrapper">
            <VueSelect
              class="library-select"
              v-model="localSettings.libarys"
              :options="libraryOptions"
              :is-multi="true"
              placeholder="Select libraries"
            />
            <button type="button" class="refresh" @click="handleLoadLibaries">
              <font-awesome-icon icon="fa-refresh" />
            </button>
          </div>
        </div>

        <div class="form-group">
          <label>Plex Token</label>
          <div class="token-input-wrapper">
            <input
              :type="showPassword ? 'text' : 'password'"
              v-model="localSettings.token"
              placeholder="Your Plex Token"
            />
            <button
              type="button"
              class="toggle-password"
              @click="showPassword = !showPassword"
            >
              <font-awesome-icon
                :icon="showPassword ? 'eye-slash' : 'eye'"
                fixed-width
              />
            </button>
          </div>
        </div>

        <div class="modal-actions">
          <IconButton
            @buttonClick="handleTest"
            title="Test"
            :active="true"
            icon="vial"
            :showAnimation="isLoadingTest"
          >
          </IconButton>
          <IconButton
            @buttonClick="validateAndSave"
            :active="true"
            title="Save"
            icon="floppy-disk"
            :showAnimation="false"
          >
          </IconButton>
          <IconButton
            @buttonClick="close"
            title="Cancel"
            :showAnimation="false"
            :active="false"
          >
          </IconButton>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref, watch } from "vue";
import { saveSettings, getSettings } from "@/service/ApiService";
import { useToast } from "vue-toastification";
import IconButton from "./IconButton.vue";
import VueSelect from "vue3-select-component";

const toast = useToast();
const isLoadingTest = ref<boolean>(false);

const plexLibary = ref<PlexLibaryResponse>();

const props = defineProps({
  visible: Boolean,
});

const showPassword = ref(false);

const emit = defineEmits(["update:visible", "saved"]);

const localSettings = reactive({
  baseurl: "",
  token: "",
  libarys: [] as string[],
});

interface PlexLibaryResponse {
  MediaContainer: PlexLibaryMediaResponse;
}

interface PlexLibaryMediaResponse {
  Directory: PlexLibaryDirectoryResponse[];
}

interface PlexLibaryDirectoryResponse {
  title: string;
  type: string;
}

const libraryOptions = computed(
  () =>
    plexLibary.value?.MediaContainer.Directory.map((lib) => ({
      label: lib.title,
      value: lib.title,
    })) ?? [],
);

function initLibaries() {
  if (localSettings.baseurl != null && localSettings.token != null) {
    handleLoadLibaries();
  }
}

watch(
  () => props.visible,
  async (val) => {
    if (val) {
      try {
        const res = await getSettings();
        localSettings.baseurl = res.baseurl || "";
        localSettings.token = res.token || "";
        localSettings.libarys = Array.isArray(res.libarys)
          ? res.libarys
          : res.libarys
            ? res.libarys.split(",").map((s) => s.trim())
            : [];
        console.error(res);
        initLibaries();
      } catch (err) {
        console.error(err);
        toast.error("Could not load settings");
      }
    }
  },
);

function close() {
  emit("update:visible", false);
}

async function handleLoadLibaries() {
  if (!localSettings.baseurl || !localSettings.token) {
    toast.error("Please fill Base URL and Token first");
    return;
  }
  try {
    const response = await fetch(`${localSettings.baseurl}/library/sections`, {
      headers: {
        "X-Plex-Token": localSettings.token,
        Accept: "application/json",
      },
    });
    if (!response.ok) {
      toast.error("Response was fault: ${err.message}");
      throw new Error(`HTTP ${response.status}`);
    }
    plexLibary.value = await response.json();
  } catch (err: any) {
    toast.error(`Libaries could not be fetched: ${err.message}`);
  } finally {
    isLoadingTest.value = false;
  }
}

async function handleTest() {
  handleLoadLibaries();
  if (!localSettings.baseurl || !localSettings.token) {
    toast.error("Please fill Base URL and Token first");
    return;
  }

  try {
    isLoadingTest.value = true;

    const response = await fetch(`${localSettings.baseurl}/identity`, {
      headers: {
        "X-Plex-Token": localSettings.token,
        Accept: "application/json",
      },
    });

    if (!response.ok) {
      toast.error("Response was fault: ${err.message}");
      throw new Error(`HTTP ${response.status}`);
    }
    toast.success(`Connection successful!`);
  } catch (err: any) {
    toast.error(`Connection failed: ${err.message}`);
  } finally {
    isLoadingTest.value = false;
  }
}

async function validateAndSave() {
  if (!localSettings.baseurl || !localSettings.token) {
    toast.error("Base URL and Token are required");
    return;
  }

  try {
    await saveSettings({
      token: localSettings.token,
      baseurl: localSettings.baseurl,
      libarys: localSettings.libarys,
    });
    toast.success("Settings saved!");
    close();
  } catch (err) {
    toast.error("Failed to save settings");
  }
}
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  inset: 0;
  background: var(--bg-overlay);
  backdrop-filter: blur(4px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 50;
}

.modal {
  background: var(--bg-primary);
  color: var(--text-primary);
  padding: 1.75rem;
  border-radius: var(--radius-lg);
  width: 420px;
  max-width: calc(100vw - 2rem);
  box-shadow: var(--shadow-lg);
  border: 1px solid var(--border-subtle);
}

.modal-title {
  color: var(--plex-accent);
  margin-bottom: 1.25rem;
  padding-bottom: 0.6rem;
  border-bottom: 2px solid var(--plex-accent);
  font-weight: 600;
  font-size: 1.1rem;
  letter-spacing: 0.02em;
}

.form-group {
  margin-bottom: 1rem;
}

.library-select {
  --vs-min-height: 2.45rem;
  --vs-padding: 0.35rem 0.75rem;
  --vs-border: 1px solid var(--border-medium);
  --vs-border-radius: var(--radius-sm);
  --vs-font-size: 0.9rem;
  --vs-font-family: inherit;
  --vs-text-color: var(--text-primary);
  --vs-placeholder-color: var(--text-muted);
  --vs-background-color: var(--bg-surface);
  --vs-outline-width: 0;
  --vs-outline-color: transparent;
  --vs-menu-offset-top: 0.35rem;
  --vs-menu-height: 13rem;
  --vs-menu-border: 1px solid var(--border-medium);
  --vs-menu-background-color: var(--bg-surface);
  --vs-menu-box-shadow: var(--shadow-lg);
  --vs-menu-z-index: 80;
  --vs-option-padding: 0.55rem 0.7rem;
  --vs-option-font-size: 0.88rem;
  --vs-option-text-color: var(--text-primary);
  --vs-option-focused-text-color: var(--text-primary);
  --vs-option-selected-text-color: var(--text-primary);
  --vs-option-disabled-text-color: var(--text-muted);
  --vs-option-background-color: transparent;
  --vs-option-focused-background-color: var(--bg-surface-hover);
  --vs-option-selected-background-color: var(--plex-accent-dim);
  --vs-option-disabled-background-color: transparent;
  --vs-multi-value-margin: 0.15rem;
  --vs-multi-value-border-radius: var(--radius-sm);
  --vs-multi-value-background-color: rgba(229, 160, 13, 0.14);
  --vs-multi-value-label-padding: 0.24rem 0.25rem 0.24rem 0.5rem;
  --vs-multi-value-label-font-size: 0.78rem;
  --vs-multi-value-label-font-weight: 500;
  --vs-multi-value-label-line-height: 1.15;
  --vs-multi-value-label-text-color: var(--text-primary);
  --vs-multi-value-delete-padding: 0 0.32rem;
  --vs-multi-value-delete-hover-background-color: rgba(255, 255, 255, 0.08);
  --vs-multi-value-xmark-size: 0.9rem;
  --vs-multi-value-xmark-color: var(--text-secondary);
  --vs-multi-value-xmark-hover-color: var(--text-primary);
  --vs-indicator-icon-size: 1rem;
  --vs-indicator-icon-color: var(--text-secondary);
  --vs-spinner-color: var(--plex-accent);
}

.library-select :deep(.control) {
  transition:
    border-color var(--transition-fast),
    background-color var(--transition-fast),
    box-shadow var(--transition-fast);
}

.library-select :deep(.control:hover) {
  background-color: var(--bg-surface-hover);
}

.library-select :deep(.control.focused) {
  border-color: var(--plex-accent);
  box-shadow: 0 0 0 3px var(--plex-accent-dim);
}

.library-select :deep(.value-container) {
  gap: 0.25rem;
}

.library-select :deep(.search-input::placeholder) {
  color: var(--text-muted);
}

.library-select :deep(.menu) {
  padding: 0.35rem;
}

.library-select :deep(.menu-option) {
  border-radius: var(--radius-sm);
  transition:
    background-color var(--transition-fast),
    color var(--transition-fast);
}

.library-select :deep(.menu-option + .menu-option) {
  margin-top: 0.15rem;
}

.library-select :deep(.multi-value) {
  border: 1px solid rgba(229, 160, 13, 0.25);
}

.form-group label {
  display: block;
  font-size: 0.8rem;
  font-weight: 500;
  color: var(--text-secondary);
  margin-bottom: 0.35rem;
  text-transform: uppercase;
  letter-spacing: 0.04em;
}

.form-group input {
  width: 100%;
  padding: 0.55rem 0.75rem;
  border-radius: var(--radius-sm);
  border: 1px solid var(--border-medium);
  background: var(--bg-surface);
  color: var(--text-primary);
  font-size: 0.9rem;
  font-family: inherit;
  outline: none;
  transition: border-color var(--transition-fast);
}

.form-group input::placeholder {
  color: var(--text-muted);
}

.form-group input:focus {
  border-color: var(--plex-accent);
}

.token-input-wrapper {
  display: flex;
  align-items: stretch;
  gap: 0;
  position: relative;
}

.token-input-wrapper input {
  padding-right: 2.5rem;
}

.refresh {
  right: 0;
  top: 0;
  bottom: 0;
  width: 2.5rem;
  background: none;
  border: none;
  cursor: pointer;
  color: var(--text-muted);
  display: flex;
  align-items: center;
  justify-content: center;
  transition: color var(--transition-fast);
}

.toggle-password {
  position: absolute;
  right: 0;
  top: 0;
  bottom: 0;
  width: 2.5rem;
  background: none;
  border: none;
  cursor: pointer;
  color: var(--text-muted);
  display: flex;
  align-items: center;
  justify-content: center;
  transition: color var(--transition-fast);
}

.toggle-password:hover {
  color: var(--text-secondary);
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 0.5rem;
  margin-top: 1.5rem;
  padding-top: 1rem;
  border-top: 1px solid var(--border-subtle);
}

/* Modal transition */
.modal-enter-active {
  transition: opacity 0.2s ease;
}

.modal-enter-active .modal {
  transition:
    transform 0.2s ease,
    opacity 0.2s ease;
}

.modal-leave-active {
  transition: opacity 0.15s ease;
}

.modal-leave-active .modal {
  transition:
    transform 0.15s ease,
    opacity 0.15s ease;
}

.modal-enter-from {
  opacity: 0;
}

.modal-enter-from .modal {
  transform: scale(0.95);
  opacity: 0;
}

.modal-leave-to {
  opacity: 0;
}

.modal-leave-to .modal {
  transform: scale(0.95);
  opacity: 0;
}

:deep(.vue-select-header) {
  min-height: 42px;
  border-radius: var(--radius-sm);
  border: 1px solid var(--border-medium);
  background: var(--bg-surface);
  color: var(--text-primary);
}

:deep(.vue-select-focused .vue-select-header) {
  border-color: var(--plex-accent);
}

:deep(.vue-select-dropdown) {
  background: var(--bg-primary);
  border: 1px solid var(--border-medium);
}

:deep(.vue-select-option) {
  color: var(--text-primary);
}

:deep(.vue-select-option-selected) {
  background: var(--plex-accent);
  color: white;
}

:deep(.vue-select-tag) {
  background: var(--plex-accent);
  color: white;
  border-radius: 999px;
  padding: 0.1rem 0.5rem;
}
</style>
