<script setup>
import SwitchListItem from "../components/SwitchListItem.vue";
import { reactive } from "vue";
import { useRouter } from "vue-router";
import MegaMenu from "primevue/megamenu";
import MySwitchService from "../service/MySwitchService";
import { ref, onMounted } from "vue";

const router = useRouter();
var switches = reactive([]);

function navigateToPage() {
  router.push("/settings/");
}

onMounted(() => {
  MySwitchService.getMySwitches().then((data) => {
    switches.push(...data);
  });
});

</script>

<template>
  <div>
    <MegaMenu
      :model="menubar"
      class="menubar"
    >
      <template #start>
        <p class="title">
            <h1 class="Available Switches">Available Switches</h1>
        </p>
      </template>
      <template #end>
        <i class="pi pi-cog" style="font-size: 3rem" @click="navigateToPage"></i>
      </template>
    </MegaMenu>
  </div>
  <div class="switch-list-container">
    <div v-for="switchitem in switches">
      <SwitchListItem
        class="switch-list-item"
        :ipAddress="switchitem.IPaddress"
        :switchStatus="switchitem.switchStatus"
        :name="switchitem.name"
      />
    </div>
  </div>

</template>

<style>
.switch-list-container {
  display: flex;
  flex-wrap: wrap;
  padding: 20px;
  width: 100vw;
  justify-content: center;
  margin-top: 50px;
}

.switch-list-item {
  display: flex;
  margin: 10px;
}

.title {
  font-size: 1.5rem;
  font-weight: 500;
  padding-left: 30px;
}

.pi-cog {
  padding-right: 30px; 
  cursor: pointer;
  transition: transform 0.2s ease-in-out, box-shadow 0.2s ease-in-out;
}

.pi-cog:hover { 
  transform: translateY(-2px);
}

.menubar{
  height: 80px; 
  margin: 50px 86px; 
  margin-bottom: 150px; 
}
</style>
