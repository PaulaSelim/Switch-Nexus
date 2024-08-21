import { createApp } from "vue";
import App from "./App.vue";
import "./style.css";
import { router } from "./router";

import PrimeVue from "primevue/config";
import Aura from "@primevue/themes/aura";
import "primeicons/primeicons.css";

import ConfirmationService from 'primevue/confirmationservice'
import DialogService from 'primevue/dialogservice'
import ToastService from 'primevue/toastservice';

import { FilterMatchMode } from "@primevue/core/api";
import { useToast } from "primevue/usetoast";
import Dialog from "primevue/dialog";
import Button from "primevue/button";
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import Toolbar from "primevue/toolbar";
import InputText from "primevue/inputtext";
import IconField from "primevue/iconfield";
import InputIcon from "primevue/inputicon";
import MegaMenu from "primevue/megamenu";

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

app.use(ConfirmationService);
app.use(ToastService);
app.use(DialogService);

app.use(FilterMatchMode);
app.use(useToast);
app.use(Dialog);
app.use(Button);
app.use(DataTable);
app.use(Column);
app.use(Toolbar);
app.use(InputText);
app.use(IconField);
app.use(InputIcon);
app.use(MegaMenu);


export default app;
