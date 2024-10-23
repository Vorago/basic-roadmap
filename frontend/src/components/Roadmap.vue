<template>
  <ContextMenu ref="menuRef" :model="menuItems"/>
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
        <span @contextmenu.prevent="showContextForSwimlane($event, index, sl.name)">
            {{ sl.name }}
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
      label: 'Move Up',
      disabled: true
    },
    {
      label: 'Move Down',
      disabled: true
    }
  ]
  menuRef.value.show(e);
}

function showContextForSwimlane(e: Event, index: number, title: string) {
  menuItems.value = [
    {
      label: 'Rename',
      command: (e) => showSwimlaneRenameModal(e, index, title)
    },
    {
      label: 'Delete',
      command: () => store.removeSwimlane(index)
    },
    {
      label: 'Add Bar',
      command: () => store.addBarToSwimlane(index)
    }
  ]
  menuRef.value.show(e);
}


function showSwimlaneRenameModal(e: MenuItemCommandEvent, index: number, title: string) {
  dialog.open(modalInput, {
        props: {
          header: "Edit Swimlane Title",
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
          title: title,
          id: index
        },
        onClose: (e) => {
        },
        emits: {
          onSave: (e: any) => {
            console.log('saving')
            console.log(e)
            store.renameSwimlane(index, e.title)
          }
        }
      }
  )
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
