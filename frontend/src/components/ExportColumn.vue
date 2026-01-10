<template>
  <div class="export-panel">
    <h3>Exports</h3>
    <div class="export-list">
      <div v-if="exports.length === 0" class="no-data">No exports yet</div>
      <div v-for="(item, index) in exports" :key="index" class="export-item">
        <span class="name">{{ item.name }}</span>
        <span class="date">{{ item.date }}</span>
        <span class="count">{{ item.count }} entries</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineProps } from "vue";

const props = defineProps({
  exports: {
    type: Array,
    required: true,
    default: () => [],
  },
});
</script>

<style scoped>
.export-panel {
  color: #f9fafb;
  padding: 1rem;
  border-radius: 8px;
  width: 100%;
  max-width: 600px;
}

.export-panel h3 {
  margin-top: 0;
  margin-bottom: 0.75rem;
  font-size: 1.2rem;
  border-bottom: 1px solid #2d3138;
  padding-bottom: 0.5rem;
}

.export-list {
  max-height: 300px;
  overflow-y: auto;
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

.name {
  flex: 2;
  font-weight: 600;
}

.date {
  flex: 1;
  text-align: center;
  color: #aaa;
}

.count {
  flex: 1;
  text-align: right;
  font-weight: 500;
}

.no-data {
  text-align: center;
  color: #888;
  padding: 1rem 0;
}
</style>
