<template>
  <InputText autofocus v-on:keydown="onKey($event)" type="text" v-model="modelValue"/>
  <Button label="Save" icon="fa fa-solid fa-check" :onclick="onSave"/>
</template>

<script lang="ts" setup>
import {inject} from "vue";
import InputText from "primevue/inputtext";
import Button from 'primevue/button';

const dialogRef: any = inject('dialogRef');

const emit = defineEmits(['cancel', 'save'])
const modelValue = defineModel("barName")

modelValue.value = dialogRef.value.data.title

function onKey(event: KeyboardEvent) {
  if (event && event.key == 'Enter') {
    onSave()
  }
}

function onSave() {
  emit('save', {title: modelValue.value, barID: dialogRef.value.data.id});

  dialogRef.value.close();
}
</script>
