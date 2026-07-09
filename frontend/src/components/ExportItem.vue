<template>
    <div class="export-card">
        <div class="accent-bar"></div>
        <div class="card-body">
            <div class="export-info">
                <div class="export-title">{{ title }}</div>
                <div class="export-details">
                    <span class="export-meta">{{ formattedDate }}</span>
                    <span class="export-separator">&middot;</span>
                    <span class="export-meta">{{ entries }} entries</span>
                </div>
            </div>
            <div class="export-actions">
                <button class="action-btn download-btn" @click="downloadCSV" title="Download CSV">
                    <font-awesome-icon icon="download" />
                </button>
                <button class="action-btn delete-btn" @click="removeExport" title="Delete export">
                    <font-awesome-icon icon="trash" />
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


const emit = defineEmits<{
    (e: 'removed', id: string): void
}>()

const formattedDate = computed(() => {
    if (!props.date) return "";
    return new Intl.DateTimeFormat("en-US", {}).format(new Date(props.date));
});

async function removeExport() {
    ApiService.removeExport(props.title).then(() => {
        toast.success("Export removed")
        emit('removed', props.title)
    })
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

<style scoped>
.export-card {
    display: flex;
    border-radius: var(--radius-md);
    background: var(--bg-surface);
    border: 1px solid var(--border-subtle);
    overflow: hidden;
    transition: background var(--transition-fast), border-color var(--transition-fast);
}

.export-card:hover {
    background: var(--bg-surface-hover);
    border-color: var(--border-medium);
}

.accent-bar {
    width: 3px;
    flex-shrink: 0;
    background: var(--plex-accent);
}

.card-body {
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 100%;
    padding: 0.85rem 1rem;
    gap: 1rem;
}

.export-info {
    display: flex;
    flex-direction: column;
    gap: 2px;
    min-width: 0;
}

.export-title {
    font-size: 0.9rem;
    font-weight: 600;
    color: var(--text-primary);
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}

.export-details {
    display: flex;
    align-items: center;
    gap: 0.4rem;
}

.export-meta {
    font-size: 0.78rem;
    color: var(--text-muted);
}

.export-separator {
    color: var(--text-muted);
    font-size: 0.7rem;
}

.export-actions {
    display: flex;
    align-items: center;
    gap: 0.25rem;
    flex-shrink: 0;
}

.action-btn {
    width: 34px;
    height: 34px;
    display: flex;
    align-items: center;
    justify-content: center;
    border: none;
    background: transparent;
    color: var(--text-secondary);
    cursor: pointer;
    border-radius: var(--radius-sm);
    font-size: 0.85rem;
    transition: all var(--transition-fast);
}

.download-btn:hover {
    background: var(--plex-accent-dim);
    color: var(--plex-accent);
}

.delete-btn:hover {
    background: rgba(239, 68, 68, 0.12);
    color: #ef4444;
}
</style>
