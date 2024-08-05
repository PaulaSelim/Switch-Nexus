import { createApp } from "vue";
import App from "./App.vue";
import "./style.css";
import { router } from "./router";

import PrimeVue from "primevue/config";
import Aura from "@primevue/themes/aura";
import "primeicons/primeicons.css";

const app = createApp(App);

app.use(router);
app.mount("#app");

app.use(PrimeVue, {
  // Default theme configuration
  theme: {
    preset: Aura,
    options: {
      prefix: "p",
      darkModeSelector: "light",
      cssLayer: false,
    },
  },
});

export default app;
