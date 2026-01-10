<template>
  <Header />
  <div class="app-container">
    <div class="content">
      <ExportColumn />
    </div>
  </div>
</template>

<script setup>
import ExportColumn from './components/ExportColumn.vue';
import Header from './components/Header.vue';
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
