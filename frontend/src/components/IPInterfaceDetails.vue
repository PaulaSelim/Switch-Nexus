<script setup>
import { reactive } from "vue";
import { GetIPInterface } from "../../wailsjs/go/wails_interfaces/WailsInterface";
import Inplace from "primevue/inplace";
import Button from "primevue/button";

const data = reactive({
  resultText: "Getting IP Interface Details ...",
});

const props = defineProps({
  host: String,
});

function get_IP_Interface() {
    GetIPInterface(props.host).then((res) => {
    data.resultText = res;
  });
}
</script>

<template>
  <main class="main">
    <Inplace>
      <template #display>
        <div class="input-box">
          <Button class="btn" @click="get_IP_Interface">
            Get IP Interface Details
          </Button>
        </div>
      </template>
      <template #content>
        <div class="result">{{ data.resultText }}</div>
      </template>
    </Inplace>
  </main>
</template>

<style scoped>
.main {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.btn {
  width: 200px;
  padding: 12px;
  transition: transform 0.2s ease-in-out, box-shadow 0.2s ease-in-out;
}

.btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.25);
}

.result {
  font-family: "sans-serif", monospace;
  background-color: #000;
  color: white;
  padding: 20px;
  border-radius: 5px;
  margin-top: 20px;
  white-space: pre-wrap;
  width: 100%;
  max-width: 800px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
}

::v-deep .p-inplace-display {
  padding: 0 !important;
}
</style>
