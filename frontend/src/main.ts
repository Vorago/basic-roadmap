import {createApp} from 'vue'
import App from './App.vue'
import './style.css';
import ganttastic from '@infectoone/vue-ganttastic'
import PrimeVue from "primevue/config";
import Aura from '@primevue/themes/aura';
import DialogService from 'primevue/dialogservice';

createApp(App)
    .use(ganttastic)
    .use(DialogService)
    .use(PrimeVue, {
        theme: {
            preset: Aura,
            options: {
                darkModeSelector: false
            }
        }
    })
    .mount('#app')
