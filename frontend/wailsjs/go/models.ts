export namespace main {
	
	export class GanttBarConfig {
	    id: string;
	    label: string;
	    hasHandles: boolean;
	
	    static createFrom(source: any = {}) {
	        return new GanttBarConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.label = source["label"];
	        this.hasHandles = source["hasHandles"];
	    }
	}
	export class Bar {
	    start: string;
	    end: string;
	    ganttBarConfig: GanttBarConfig;
	
	    static createFrom(source: any = {}) {
	        return new Bar(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.start = source["start"];
	        this.end = source["end"];
	        this.ganttBarConfig = this.convertValues(source["ganttBarConfig"], GanttBarConfig);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class Swimlane {
	    name: string;
	    start: string;
	    end: string;
	    items: Bar[];
	
	    static createFrom(source: any = {}) {
	        return new Swimlane(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.start = source["start"];
	        this.end = source["end"];
	        this.items = this.convertValues(source["items"], Bar);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

