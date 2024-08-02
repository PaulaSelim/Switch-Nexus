import { createMemoryHistory, createRouter } from "vue-router";

import Home from "./views/Home.vue";
import Login from "./views/Login.vue";
import SwitchList from "./views/SwitchList.vue";
import SwitchDetails from "./views/SwitchDetails.vue";
import Settings from "./views/Settings.vue";

const routes = [
  { path: "/", component: Home },
  { path: "/login", component: Login },
  { path: "/switchlist", component: SwitchList },
  { path: "/switchdetails", component: SwitchDetails },
  { path: "/settings", component: Settings },
];

const router = createRouter({
  history: createMemoryHistory(),
  routes,
});

export { router };