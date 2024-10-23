import {reactive} from 'vue'
import {GetSwimlanes, LoadFromFile, UpdateSwimlanes} from "../wailsjs/go/main/App.js";
import {main} from "../wailsjs/go/models";
import { useFileDialog } from '@vueuse/core'

class GBC {
    public id!: string;
    public label!: string;
    public hasHandles!: boolean;
}

class Bar {
    public start!: string;
    public end!: string;
    public ganttBarConfig!: GBC;

    constructor(id: string) {
        this.ganttBarConfig = new GBC()
        this.ganttBarConfig.id = id
        this.start = '2024-10-01'
        this.end = '2024-10-10'
        this.ganttBarConfig.label = 'bar'
        this.ganttBarConfig.hasHandles = true
    }
}

class Swimlane {
    public name!: string;
    public start!: string;
    public end!: string;
    public items!: Bar[];
}


class Store {
    public swimlanes!: Swimlane[]
    public addSwimlane() {
        let sw = new Swimlane()
        sw.items = []
        sw.name = "new"
        this.swimlanes.push(sw)
    }
    public removeSwimlane(index: number) {
        this.swimlanes.splice(index, 1)
    }
    public init(swimlanes: Swimlane[]) {
        this.swimlanes = swimlanes
    }

    public updateBar(bars: Bar[], id: string, newTitle: string) {
        bars.forEach((bar, index) => {
            if (bar.ganttBarConfig.id === id) {
                bar.ganttBarConfig.label = newTitle
                return
            }
        })
    }

    public updateBarTitle(id: string, title: string){
        this.swimlanes.forEach(sl => this.updateBar(sl.items, id, title))
    }
    public addBarToSwimlane(index: number) {
        this.swimlanes[index].items.push(new Bar('idid'))
    }
    public renameSwimlane(index: number, name: string) {
        this.swimlanes[index].name = name
    }
    public saveToBackend() {
        UpdateSwimlanes(this.swimlanes.map(main.Swimlane.createFrom))
            .catch((err) => console.log(err))
    }

    public loadFromBackend() {
        LoadFromFile().then(s => {
            store.init(s)
        })
    }
}

export const store: Store = reactive(new Store())

