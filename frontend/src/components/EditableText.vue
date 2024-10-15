<script setup lang="ts">
import {ref} from 'vue'
import { defineEmits } from "vue";

const emit = defineEmits(['update:modelValue', 'save'])

const edit = ref(false)
const props = defineProps(['value'])

const modelValue = ref(props.value)

function updateModel(event: any) {
  if (event.target) {
    emit('update:modelValue', event.target.value)
  }
}

const vFocus = {
  mounted: (el:HTMLInputElement) => el.focus()
}
</script>


<template>
  <span>
    <input class="editable"
           type="text"
           v-if="edit"
           @blur="$emit('save', ($event.target as HTMLInputElement).value); edit=false"
           @keyup.enter="$emit('save', ($event.target as HTMLInputElement).value); edit=false"
           @keyup.esc="edit = false"
           @input="updateModel($event)"
           v-model="modelValue"
           v-focus
    />
    <span v-else @dblclick="edit = true;">
      {{ modelValue }}
    </span>
  </span>
</template>


<style>
.editable {
  width: 100%;
}
</style>
