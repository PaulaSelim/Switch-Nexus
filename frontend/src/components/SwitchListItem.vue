<script setup>
import { useRouter } from "vue-router";
import Badge from "primevue/badge";
import { GetStatus } from "../../wailsjs/go/wails_interfaces/WailsInterface";
import { onMounted, reactive } from "vue";

const router = useRouter();

function navigateToPage(ipAddress) {
  const safeIPAddress = ipAddress.replace(/\./g, "-").replace(/:/g, "=");
  router.push("/switch/" + safeIPAddress);
}

const data = reactive({
  status: "contrast",
});

const props = defineProps({
  ipAddress: String,
  switchStatus: String,
  name: String,
});

function get_status() {
  var addr = props.ipAddress.split(":")[0];
  GetStatus(addr)
    .then((res) => {
      if (res == "success") data.status = "success";
    })
    .catch((err) => {
      data.status = "danger";
    });
}

onMounted(() => {
  get_status();
});
</script>

<template>
  <div class="card" @click="navigateToPage(ipAddress)">
    <div class="card-icon">
      <i class="pi pi-server" style="font-size: 2.5rem"></i>
    </div>
    <div class="card-content">
      <h4 class="ip-address">Host: {{ ipAddress }}</h4>
      <h1 class="name">{{ name }}</h1>
      <p class="badge">
        status: <Badge class="badge-dot" value="" :severity="data.status" />
      </p>
    </div>
  </div>
</template>

<style scoped>
.card {
  width: fit-content;
  border-radius: 6px;
  display: flex;
  color: var(--primary-color);
  transition: transform 0.2s ease-in-out, box-shadow 0.2s ease-in-out;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  cursor: pointer;
  padding: 10px;
  background-color: var(--card-bg);
}

.ip-address {
  font-size: 12px;
  font-weight: 500;
  margin: 0;
  color: var(--p-surface-500);
}

.card:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.25);
}

.card-content {
  min-width: 200px;
}

.card-icon {
  display: flex;
  justify-content: center;
  align-items: center;
  /* padding: 20px; */
  margin-right: 10px;
  color: var(--p-surface-700);
}

.name {
  font-size: 24px;
  font-weight: 700;
  margin: 5px -1px;
}

.badge {
  font-size: 14px;
  font-weight: 500;
  margin: 0;
  color: var(--p-surface-500);
  float: inline-end;
  display: flex;
}

.badge-dot {
  align-self: center;
  margin-left: 4px;
}
</style>
