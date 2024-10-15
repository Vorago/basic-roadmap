<template>
  <div style="z-index: 1200">
    <Menubar :model="items"  />
  </div>
</template>

<script lang="ts" setup>
import {store} from "../state"
import 'font-awesome/css/font-awesome.min.css';
import Menubar from 'primevue/menubar';
import {MenuItem} from "primevue/menuitem";


let items: MenuItem[]
items = [
  {
    label: 'File',
    items: [
      {
        label: 'Save',
        command: () => store.saveToBackend()
      },
      {
        label: 'Load',
        command: () => store.loadFromBackend()
      }
    ]
  },
  {
    label: 'Edit',
    items: [
      {
        label: "Add Swimlane",
        command: () => store.addSwimlane()
      }
    ]
  }
]

</script>

