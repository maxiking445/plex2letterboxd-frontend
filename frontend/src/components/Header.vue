<template>
    <SettingsDialog v-model:visible="showSettings" @saved="handleSettingsSaved" />
    <div class="app">
        <header class="app-header">
            <div class="header-left">
                <img alt="Vue logo" class="logo" src="../../assets/export.svg" width="28" height="28" />
                <h1 class="title">Plex2Letterboxd</h1>
            </div>
            <div class="header-right">
                <IconButton title="Import to Letterboxd" icon="arrow-up-right-from-square" :active=false
                    :show-animation=false @buttonClick="handleImport"></IconButton>
                <IconButton title="Start Export" icon="play" :show-animation=isLoading :active=true
                    @buttonClick="handleStart"></IconButton>
                <IconButton title="Settings" :icon="settingsStatusIcon" :status="settingsStatus" :active=false
                    @buttonClick="handleSetting" :show-animation=false></IconButton>

                <a class="github-link" href="https://github.com/maxiking445/plex2letterboxd-frontend" target="_blank"
                    rel="noopener noreferrer">
                    <img src="../../assets/github.svg" alt="GitHub" class="github-icon" />
                </a>

            </div>
        </header>
    </div>
</template>

<script setup lang="ts">
import { getSettings, startExport } from '@/service/ApiService';
import { computed, onMounted, ref } from 'vue';
import { useToast } from 'vue-toastification'
import SettingsDialog from './SettingsDialog.vue';
import IconButton from './IconButton.vue';

type SettingsStatus = 'checking' | 'ready' | 'required';

const showSettings = ref(false);
const toast = useToast()
var isLoading = ref(false)
const settingsStatus = ref<SettingsStatus>('checking')


const emit = defineEmits<{
    (e: 'finishedExport'): void
}>()

const handleStart = () => {
    isLoading.value = true
    startExport().then(res => {
        isLoading.value = false
        emit('finishedExport')
    }).catch(err => {
        isLoading.value = false
        toast.error("Check your settings. Something went wrong.")
    })
}

const settingsStatusIcon = computed(() => {
    if (settingsStatus.value === 'ready') return 'check';
    if (settingsStatus.value === 'checking') return 'spinner';
    return 'circle-exclamation';
})

const hasCompleteSettings = (settings: Awaited<ReturnType<typeof getSettings>>) => {
    return Boolean(
        settings.baseurl?.trim() &&
        settings.token?.trim() &&
        Array.isArray(settings.libarys) &&
        settings.libarys.length > 0
    );
}

const refreshSettingsStatus = async () => {
    settingsStatus.value = 'checking'
    try {
        const settings = await getSettings(false)
        settingsStatus.value = hasCompleteSettings(settings) ? 'ready' : 'required'
    } catch (err) {
        settingsStatus.value = 'required'
    }
}

const handleSettingsSaved = () => {
    refreshSettingsStatus()
}

const handleImport = () => {
    window.open('https://letterboxd.com/import/', '_blank')
}

const handleSetting = () => {
    showSettings.value = true;
}

onMounted(() => {
    refreshSettingsStatus()
})
</script>

<style scoped>
.header-right {
    display: flex;
    align-items: center;
    gap: 0.5rem;
}

.github-link {
    filter: invert(1);
    opacity: 0.5;
    transition: opacity var(--transition-fast);
}

.github-link:hover {
    opacity: 1;
}

.app-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0.75rem 1.5rem;
    border-bottom: 1px solid var(--border-subtle);
    background: var(--bg-primary);
    backdrop-filter: blur(12px);
    position: sticky;
    top: 0;
    z-index: 40;
}

.header-left {
    display: flex;
    align-items: center;
    gap: 0.6rem;
}

.logo {
    display: block;
}

.title {
    margin: 0;
    font-size: 1.05rem;
    font-weight: 600;
    color: var(--text-primary);
    letter-spacing: -0.01em;
}

.github-link {
    display: inline-flex;
    align-items: center;
    margin-left: 0.25rem;
}

.github-icon {
    width: 20px;
    height: 20px;
}
</style>
