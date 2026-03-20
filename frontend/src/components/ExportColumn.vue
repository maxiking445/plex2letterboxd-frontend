<template>
  <div class="export-panel">
    <div class="panel-header">
      <h3>Exports</h3>
    </div>
    <div class="export-list">
      <div v-if="exports.length === 0" class="no-data">
        <font-awesome-icon icon="download" class="no-data-icon" />
        <span>No exports yet</span>
        <span class="no-data-hint">Start an export to see it here</span>
      </div>
      <TransitionGroup name="export" tag="div">
        <div v-for="exp in exports" :key="exp.name" class="export-item">
          <ExportItem :title="exp.name" :entries="exp.exportItemCount" :date="exp.date" @removed="handleRemoved" />
        </div>
      </TransitionGroup>
    </div>
  </div>
</template>

<script setup lang="ts">
import ExportItem from "./ExportItem.vue";
import { loadExports } from '@/service/ApiService';
import { ModelsExport } from "@/generated/api";
import { onMounted, ref, watch } from "vue";

const exports = ref<ModelsExport[]>([]);

const props = defineProps<{
  refreshTrigger: Number;
}>();

onMounted(async () => {
  try {
    exports.value = await loadExports();
  } catch (err) {
    console.error("Error during load:", err);
  }
});

watch(
  () => props.refreshTrigger,
  async (newVal, oldVal) => {
    if (newVal !== oldVal) {
      console.log("TRIGGERT", newVal, oldVal)
      try {
        exports.value = await loadExports()
      } catch (err) {
        console.error("Reload failed:", err)
      }
    }
  }
)

function handleRemoved(id: string) {
  exports.value = exports.value.filter(e => e.name !== id)
}

</script>

<style scoped>
.export-panel {
  color: var(--text-primary);
  padding: 1.5rem 0;
  width: 100%;
  max-width: 600px;
  height: calc(100vh - 60px);
  display: flex;
  flex-direction: column;
}

.panel-header {
  flex-shrink: 0;
  margin-bottom: 1rem;
}

.panel-header h3 {
  margin: 0;
  font-size: 1.1rem;
  font-weight: 600;
  color: var(--text-secondary);
  text-transform: uppercase;
  letter-spacing: 0.06em;
  font-size: 0.8rem;
}

.export-list {
  overflow-y: auto;
  flex-grow: 1;
  padding-right: 0.25rem;
}

.export-list::-webkit-scrollbar {
  width: 6px;
}

.export-list::-webkit-scrollbar-track {
  background: transparent;
  border-radius: 3px;
}

.export-list::-webkit-scrollbar-thumb {
  background-color: rgba(255, 255, 255, 0.1);
  border-radius: 3px;
}

.export-list::-webkit-scrollbar-thumb:hover {
  background-color: rgba(255, 255, 255, 0.2);
}

.export-item {
  margin-bottom: 0.6rem;
}

.no-data {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.5rem;
  padding: 3rem 1rem;
  color: var(--text-muted);
}

.no-data-icon {
  font-size: 2rem;
  opacity: 0.3;
  margin-bottom: 0.25rem;
}

.no-data-hint {
  font-size: 0.8rem;
  opacity: 0.6;
}

/* List transitions */
.export-enter-active {
  transition: all 0.3s ease;
}

.export-leave-active {
  transition: all 0.2s ease;
}

.export-enter-from {
  opacity: 0;
  transform: translateY(-10px);
}

.export-leave-to {
  opacity: 0;
  transform: translateX(20px);
}
</style>
