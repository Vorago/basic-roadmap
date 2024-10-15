<template>
  <ContextMenu ref="menuRef" :model="menuItems" />
  <g-gantt-chart
      chart-start="2024-10-01"
      chart-end="2024-12-01"
      precision="week"
      bar-start="start"
      bar-end="end"
      date-format="YYYY-MM-DD"
      @contextmenu-bar="showContextForBar($event.bar, $event.e, $event.datetime)"
  >
    <span v-for="(sl, index) in store.swimlanes" style="display: flex; flex-direction: row">
      <span class="swimlane">
        <button class="btn" @click="store.removeSwimlane(index)">-</button>
        <button class="btn" @click="store.addBarToSwimlane(index)">+</button>
        <span @contextmenu.prevent="showContextForSwimlane($event)">
            <EditableText
                :value="sl.name"
                @save="(name:string) => store.renameSwimlane(index, name)"
            />
        </span>
      </span>

      <g-gantt-row :label="sl.name" :bars="sl.items"/>
    </span>

  </g-gantt-chart>
</template>

<script lang="ts" setup>
import {GanttBarObject, GGanttChart, GGanttRow} from "@infectoone/vue-ganttastic";
import {GetSwimlanes} from "../../wailsjs/go/main/App.js";
import {store} from "../state"
import EditableText from "./EditableText.vue";

import {defineAsyncComponent, ref} from "vue";
import {MenuItem, MenuItemCommandEvent} from "primevue/menuitem";
import {useDialog} from "primevue/usedialog";
import ContextMenu from "primevue/contextmenu";

const menuRef = ref();
const modalInput = defineAsyncComponent(() => import('./ModalInput.vue'));
const dialog = useDialog()

const menuItems = ref([] as MenuItem[])

function showContextForBar(bar: GanttBarObject, e: Event, datetime: any) {
  const barID = bar.ganttBarConfig.id
  const barLabel = bar.ganttBarConfig.label ? bar.ganttBarConfig.label : ""
  menuItems.value = [
    {
      label: 'Rename',
      command: (e) => showModal(e, barID, barLabel),
    },
    {
      label: 'Move Up'
    },
    {
      label: 'Move Down'
    }
  ]
  menuRef.value.show(e);
}

function showContextForSwimlane(e: Event) {
  menuItems.value = [
    {
      label: 'Rename'
    },
    {
      label: 'Delete'
    }
  ]
  menuRef.value.show(e);
}


function showModal(e: MenuItemCommandEvent, barID: string, barLabel: string) {
  dialog.open(modalInput, {
        props: {
          header: "Edit Bar Title",
          style: {
            width: '50vw',
          },
          breakpoints: {
            '960px': '75vw',
            '640px': '90vw'
          },
          modal: true,
        },
        data: {
          // title: bar.ganttBarConfig.label ? bar.ganttBarConfig.label : "",
          title: barLabel,
          // id: bar.ganttBarConfig.id
          id: barID
        },
        onClose: (e) => {
        },
        emits: {
          onSave: (e: any) => {
            store.updateBarTitle(e.barID, e.title)
          }
        }
      }
  )
}

GetSwimlanes().then(s => {
  store.init(s)
})

</script>

<style>
.g-timeaxis {
  margin-left: 150px;
}

.swimlane {
  display: flex;
  flex: 0 0 150px;
  color: rgb(64, 64, 64);
  border-top: 2px solid #eaeaea;
  border-right: 2px solid #eaeaea;
}

.swimlane .btn {
  height: 100%;
}

.swimlane span {
  height: 100%;
  width: 100%;
  display: inline-flex;
  align-items: center;
  justify-content: center;
}

.g-gantt-chart {
  z-index: 0;
}
</style>
