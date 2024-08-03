import { createApp } from "vue";
import App from "./App.vue";
import "./style.css";
import { router } from "./router";

import PrimeVue from "primevue/config";
import Aura from "@primevue/themes/aura";

import Button from "primevue/button";
import Card from "primevue/card";
import Chip from "primevue/chip";

const app = createApp(App);

app.use(router);
app.mount("#app");

app.use(PrimeVue, {
  theme: {
    preset: Aura,
  },
});

app.component("Button", Button);
app.component("Card", Card);
app.component("Chip", Chip);
