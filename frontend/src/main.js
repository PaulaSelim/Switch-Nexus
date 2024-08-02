import {createApp} from 'vue'
import App from './App.vue'
import './style.css';

import PrimeVue from 'primevue/config';
import Aura from '@primevue/themes/aura';

createApp(App).mount('#app')

app.use(PrimeVue, {
    theme: {
        preset: Aura
    }
});