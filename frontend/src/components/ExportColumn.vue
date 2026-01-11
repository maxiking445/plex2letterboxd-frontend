<template>
  <div class="export-panel">
    <h3>Exports</h3>
    <div class="export-list">
      <div v-if="exports.length === 0" class="no-data">No exports yet</div>
      <div v-for="exp in exports" :key="exp.name" class="export-item">
        <ExportItem :title="exp.name" :entries="exp.exportItemCount" :date="exp.date" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import ExportItem from "./ExportItem.vue";
import { loadExports } from '@/service/ApiService';
import { ModelsExport } from "@/generated/api";
import { onMounted, ref } from "vue";

const exports = ref<ModelsExport[]>([]);

onMounted(async () => {
  try {
    const res = await loadExports();
    exports.value = res;
  } catch (err) {
    console.error("Error during load:", err);
  }
});
</script>

<style scoped>
.export-panel {
  color: #f9fafb;
  padding: 1rem;
  border-radius: 8px;
  width: 100%;
  max-width: 600px;


  height: calc(100vh - 2rem);
  display: flex;
  flex-direction: column;
}

.export-panel h3 {
  margin: 0 0 0.75rem 0;
  font-size: 1.2rem;
  border-bottom: 1px solid #2d3138;
  padding-bottom: 0.5rem;
  flex-shrink: 0;
}

.export-list {
  overflow-y: auto;
  flex-grow: 1;
  padding-right: 0.25rem;
}

.export-list::-webkit-scrollbar {
  width: 8px;
}

.export-list::-webkit-scrollbar-track {
  background: #1f2228;
  border-radius: 4px;
}

.export-list::-webkit-scrollbar-thumb {
  background-color: #555;
  border-radius: 4px;
  border: 2px solid #1f2228;
}

.export-item {
  display: flex;
  justify-content: space-between;
  padding: 0.5rem 0.25rem;
  border-bottom: 1px solid #2d3138;
}

.export-item:last-child {
  border-bottom: none;
}

.no-data {
  text-align: center;
  color: #888;
  padding: 1rem 0;
}
</style>
