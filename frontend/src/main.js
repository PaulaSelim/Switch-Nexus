import { createApp } from "vue";
import App from "./App.vue";
import "./style.css";
import { router } from "./router";

import PrimeVue from "primevue/config";
import Aura from "@primevue/themes/aura";

createApp(App).use(router).mount("#app");

app.use(PrimeVue, {
  theme: {
    preset: Aura,
  },
});
