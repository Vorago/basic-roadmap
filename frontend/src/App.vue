<script lang="ts" setup>
import Roadmap from './components/Roadmap.vue'
import Controls from "./components/Controls.vue";
import DynamicDialog from "primevue/dynamicdialog";
</script>

<template>
  <DynamicDialog />
  <Controls/>
  <Roadmap/>
</template>

