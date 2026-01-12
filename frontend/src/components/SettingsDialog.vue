<template>
    <div v-if="visible" class="modal-overlay">
        <div class="modal">
            <h3 class="modal-title">Settings</h3>

            <label>
                Plex Base URL
                <input type="text" v-model="localSettings.baseurl" placeholder="http://localhost:32400" />
            </label>

            <label>
                Plex Token
                <div style="display: flex;">
                    <input :type="showPassword ? 'text' : 'password'" v-model="localSettings.token"
                        placeholder="Your Plex Token" />
                    <button type="button" class="toggle-password" @click="showPassword = !showPassword">
                        <font-awesome-icon :icon="showPassword ? 'eye-slash' : 'eye'" :style="{ color: '#888' }"
                            fixed-width />
                    </button>
                </div>
            </label>

            <div class="modal-actions">
                <IconButton @buttonClick="handelTest" title="Test" :active=true icon="vial"
                    :showAnimation="isLoadingTest">
                </IconButton>
                <IconButton @buttonClick="validateAndSave" :active=true title="Save" icon="floppy-disk"
                    :showAnimation="false"> </IconButton>
                <IconButton @buttonClick="close" title="Cancel" :showAnimation="false" :active="false"> </IconButton>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { reactive, ref, watch } from "vue";
import { saveSettings, getSettings } from "@/service/ApiService";
import { useToast } from "vue-toastification";
import IconButton from "./IconButton.vue";

const toast = useToast();
const isLoadingTest = ref<boolean>(false)

const props = defineProps({
    visible: Boolean
});

const showPassword = ref(false)

const emit = defineEmits(["update:visible", "saved"]);

const localSettings = reactive({
    baseurl: "",
    token: ""
});

watch(() => props.visible, async (val) => {
    if (val) {
        try {
            const res = await getSettings();
            localSettings.baseurl = res.baseurl || "";
            localSettings.token = res.token || "";
        } catch (err) {
            toast.error("Could not load settings");
        }
    }
})

function close() {
    emit("update:visible", false);
}

async function handelTest() {
    if (!localSettings.baseurl || !localSettings.token) {
        toast.error("Please fill Base URL and Token first");
        return;
    }

    try {
        isLoadingTest.value = true;

        const response = await fetch(
            `${localSettings.baseurl}/identity`,
            {
                headers: {
                    "X-Plex-Token": localSettings.token,
                    "Accept": "application/json"
                }
            }
        );

        if (!response.ok) {
            toast.error("Response was fault: ${err.message}");
            throw new Error(`HTTP ${response.status}`);
        }

        const data = await response.json();

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
        await saveSettings(localSettings);
        toast.success("Settings saved!");
        close();
    } catch (err) {
        toast.error("Failed to save settings");
    }
}
</script>

<style scoped>
.modal-title {
    color: #e5a00d;
    margin-bottom: 0.75rem;
    padding-bottom: 0.5rem;
    border-bottom: 2px solid #e5a00d;
    font-weight: 600;
    letter-spacing: 0.03em;
}

.modal-overlay {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 50;
}

.modal {
    background: #181b20;
    color: #f9fafb;
    padding: 1.5rem;
    border-radius: 8px;
    width: 400px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.25);
}

label {
    display: block;
    margin-bottom: 1rem;
}

input[type="password"] {
    width: 100%;
    padding: 0.5rem 2.5rem 0.5rem 0.5rem;
    border-radius: 4px;
    border: none;
    margin-top: 0.25rem;
    box-sizing: border-box;
}

.toggle-password {
    background: none;
    border: none;
    font-size: 1rem;
    cursor: pointer;
    color: #888;
    padding: 0;
    width: 2rem;
    display: flex;
    align-items: center;
    justify-content: center;
}

input {
    width: 100%;
    padding: 0.5rem;
    border-radius: 4px;
    border: none;
    margin-top: 0.25rem;
}

.modal-actions {
    display: flex;
    justify-content: flex-end;
    gap: 0.5rem;
    margin-top: 1rem;
}

.eye-button {
    height: 100%;

}
</style>
