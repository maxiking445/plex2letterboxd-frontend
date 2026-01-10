<template>
    <SettingsDialog v-model:visible="showSettings" />
    <div class="app">
        <header class="app-header">
            <div class="header-left">
                <img alt="Vue logo" class="logo" src="../../assets/export.svg" width="32" height="32" />
                <h1 class="title">Plex2Letterboxd Exporter</h1>
            </div>
            <div class="header-right">
                <div class="additional-link">
                    <button :class="{ active: true }" @click="handleStart">
                        <font-awesome-icon icon="play" />
                        Start Export
                    </button>
                </div>

                <div class="additional-link">
                    <button :class="{ active: false }" @click="handleSetting">
                        <font-awesome-icon icon="gear" />
                        Settings
                    </button>
                </div>

                <a class="github-link" href="https://github.com/maxiking445/plex2letterboxd-frontend" target="_blank"
                    rel="noopener noreferrer">
                    <img src="../../assets/github.svg" alt="GitHub" class="github-icon" />
                </a>

            </div>
        </header>
    </div>
</template>

<script setup lang="ts">
import { startExport } from '@/service/ApiService';
import { ref } from 'vue';
import { useToast } from 'vue-toastification'
import SettingsDialog from './SettingsDialog.vue';

const showSettings = ref(false);
const toast = useToast()

const handleStart = () => {
    startExport().then(res => {
        console.log(res)
        // UPDATE TABLE
    }).catch(err => {
        toast.error("Check your config, something went wrong.")
    })
}

const handleSetting = () => {
    showSettings.value = true;
}



</script>
<style scoped>
.header-right {
    display: flex;
    align-items: center;
    gap: 0.5rem;
}




.github-link {
    filter: invert(1);
}

.app-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0.75rem 1.25rem;
    border-bottom: 1px solid #242731;
    background: #181b20;
}

.header-left {
    display: flex;
    align-items: center;
    gap: 0.5rem;
}

.logo {
    display: block;
}

.title {
    margin: 0;
    font-size: 1.1rem;
    font-weight: 600;
    color: #f9fafb;
}

.github-link {
    display: inline-flex;
    align-items: center;
}

.github-icon {
    width: 20px;
    height: 20px;
}

.additional-link button {
    padding: 0.3rem 0.8rem;
    font-weight: 600;
    border-radius: 0.25rem;
    border: none;
    background-color: #2d3138;
    color: #fff;
    cursor: pointer;
    transition: background-color 0.2s;
}

.additional-link button.active {
    background-color: #3bb54a;
    color: #181b20;
}

.additional-link button:hover {
    background-color: #44484f;
}
</style>