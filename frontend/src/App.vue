<template>
  <Header @finished-export="handleFinishedExport" />
  <div class="app-container">
    <div class="content">
      <ExportColumn :key="refreshTrigger" :refresh-trigger="refreshTrigger" />
    </div>
  </div>
</template>

<script setup lang="ts">
import ExportColumn from './components/ExportColumn.vue';
import Header from './components/Header.vue';
import { ref } from "vue";

const refreshTrigger = ref<number>(10);

function handleFinishedExport() {
  refreshTrigger.value++;
}
</script>

<style scoped>
.app-container {
  display: flex;
  justify-content: center;
  align-items: top;
  height: 100vh;
  background-color: #181b20;
}

.content {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
  width: 100%;
  max-width: 800px;
}
</style>
