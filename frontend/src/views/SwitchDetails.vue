<script setup>
import CDPNeighborsDetails from "../components/CDPNeighborsDetails.vue";
import IPInterfaceDetails from "../components/IPInterfaceDetails.vue";
import VLANDetails from "../components/VLANDetails.vue";
import { useRoute } from "vue-router";
import { ref, reactive } from "vue";
import { useRouter } from "vue-router";
import MegaMenu from "primevue/megamenu";

const route = useRoute();
const router = useRouter();

var address = ref(route.params.address.replace(/\-/g, ".").replace(/\=/g, ":"));
</script>

<template>
  <div>
    <MegaMenu :model="menubar" class="menubar">
      <template #start>
        <RouterLink to="/home">
          <p class="backButton">
            <i
              class="pi pi-angle-left"
              style="font-size: 3rem"
              @click="navigateToPage"
            ></i>
          </p>
        </RouterLink>
        <p class="title" style="">
          <template v-if="address">
            <h1 class="address">Switch {{ address }}</h1>
          </template>
        </p>
      </template>
      <template #end> </template>
    </MegaMenu>
    <CDPNeighborsDetails :host="address" />
    <IPInterfaceDetails :host="address" />
    <VLANDetails :host="address" />
  </div>
</template>
<style>
::v-deep .p-inplace-display {
  padding: 0 !important;
}

.pi-angle-left {
  color: var(--p-megamenu-color);
}
</style>
