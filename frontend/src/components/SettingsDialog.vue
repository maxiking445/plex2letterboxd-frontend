<template>
    <div v-if="visible" class="modal-overlay">
        <div class="modal">
            <h3>Settings</h3>

            <label>
                Plex Base URL
                <input type="text" v-model="localSettings.baseurl" placeholder="http://localhost:32400" />
            </label>

            <label>
                Plex Token
                <input type="password" v-model="localSettings.token" placeholder="Your Plex Token" />
            </label>

            <div class="modal-actions">
                <button class="save-btn">Test</button>
                <button @click="validateAndSave" class="save-btn">Save</button>
                <button @click="close" class="cancel-btn">Cancel</button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { reactive, watch } from "vue";
import { saveSettings, getSettings } from "@/service/ApiService";
import { useToast } from "vue-toastification";

const toast = useToast();

const props = defineProps({
    visible: Boolean
});

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

.save-btn {
    background-color: #3bb54a;
    color: #181b20;
    padding: 0.5rem 1rem;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}

.cancel-btn {
    background-color: #2d3138;
    color: #fff;
    padding: 0.5rem 1rem;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}

.save-btn:hover {
    background-color: #2faa3e;
}

.cancel-btn:hover {
    background-color: #44484f;
}
</style>
