<template>
    <Transition name="modal">
        <div v-if="visible" class="modal-overlay" @click.self="close">
            <div class="modal">
                <h3 class="modal-title">Settings</h3>

                <div class="form-group">
                    <label>Plex Base URL</label>
                    <input type="text" v-model="localSettings.baseurl" placeholder="http://localhost:32400" />
                </div>

                <div class="form-group">
                    <label>Libraries</label>
                    <input type="text" v-model="localSettings.libarys" placeholder="Movies,Series" />
                </div>

                <div class="form-group">
                    <label>Plex Token</label>
                    <div class="token-input-wrapper">
                        <input :type="showPassword ? 'text' : 'password'" v-model="localSettings.token"
                            placeholder="Your Plex Token" />
                        <button type="button" class="toggle-password" @click="showPassword = !showPassword">
                            <font-awesome-icon :icon="showPassword ? 'eye-slash' : 'eye'" fixed-width />
                        </button>
                    </div>
                </div>

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
    </Transition>
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
    token: "",
    libarys: ""
});

watch(() => props.visible, async (val) => {
    if (val) {
        try {
            const res = await getSettings();
            localSettings.baseurl = res.baseurl || "";
            localSettings.token = res.token || "";
            localSettings.libarys = Array.isArray(res.libarys)
                ? res.libarys.join(', ')
                : (res.libarys || '');
            console.error(res)
        } catch (err) {
            console.error(err)
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
        await saveSettings({
            token: localSettings.token,
            baseurl: localSettings.baseurl,
            libarys: localSettings.libarys
                ? localSettings.libarys.split(",").map(s => s.trim())
                : []
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
    transition: transform 0.2s ease, opacity 0.2s ease;
}

.modal-leave-active {
    transition: opacity 0.15s ease;
}

.modal-leave-active .modal {
    transition: transform 0.15s ease, opacity 0.15s ease;
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
</style>
