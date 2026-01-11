<template>
    <div class="export-card">
        <div class="export-info">
            <div class="export-title">{{ title }}</div>
            <div class="export-meta">Date: {{ formattedDate }}</div>
            <div class="export-meta">Entries: {{ entries }}</div>
        </div>
        <div class="export-tile">

            <div>
                <button class="bin-btn" @click="removeExport">
                    <font-awesome-icon icon="trash" size="2x" />
                </button>

            </div>
            <div>
                <button class="download-btn" @click="downloadCSV">
                    <font-awesome-icon icon="download" size="2x" />
                </button>

            </div>
        </div>
    </div>

</template>

<script setup lang="ts">
import { computed } from 'vue';
import ApiService from '@/service/ApiService';
import { useToast } from 'vue-toastification';


const toast = useToast()
const props = defineProps<{
    title: string;
    date: string;
    entries: number;
}>();

const formattedDate = computed(() => {
    if (!props.date) return "";
    return new Intl.DateTimeFormat("en-EN", {}).format(new Date(props.date));
});

async function removeExport() {
    ApiService.removeExport(props.title)
}

async function downloadCSV() {
    try {
        const blob = await ApiService.loadExport(props.title);

        const url = window.URL.createObjectURL(blob);

        const a = document.createElement('a');
        a.href = url;
        a.download = props.title;
        document.body.appendChild(a);
        a.click();

        document.body.removeChild(a);
        window.URL.revokeObjectURL(url);
    } catch (err) {
        toast.error("Download failed")
        console.error("Download failed:", err);
    }
}

</script>

<style>
.export-card {
    display: flex;
    align-items: stretch;
    justify-content: end;

    width: 100%;
    padding: 16px 20px;

    background-color: #E5A00D;
    border: 1px solid #000000;
    border-radius: 14px;
}


.export-info {
    display: flex;
    flex-direction: column;
    gap: 4px;
}


.export-tile {
    display: flex;
    align-items: stretch;
    justify-content: end;

    width: 100%;
    padding: 16px 20px;

    background-color: #E5A00D;
}


.export-title {
    font-size: 1rem;
    font-weight: 800;
    color: #111827;
}

.export-meta {
    font-size: 0.85rem;
    font-weight: 800;
    color: #6b7280;
}

.download-btn {
    height: 100%;
    width: 72px;
    min-width: 72px;
    font-size: 15px;
    border: none;
    background-color: #E5A00D;

    cursor: pointer;

    display: flex;
    align-items: center;
    justify-content: center;

    transition: background-color 0.2s ease;
}

.download-btn:hover {
    font-size: 16px;
}

.download-btn:active {
    background-color: #fce305;
}

.bin-btn {
    height: 100%;
    width: 72px;
    min-width: 72px;
    font-size: 15px;
    border: none;
    background-color: #E5A00D;

    cursor: pointer;
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: end;

    transition: background-color 0.2s ease;
}

.bin-btn:hover {
    font-size: 16px;
}

.bin-btn:active {
    background-color: #fce305;
}
</style>