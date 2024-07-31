<script setup>
import {reactive} from 'vue'
import {GetCDPNeighbors} from '../../wailsjs/go/interfaces/WailsApp'

const data = reactive({
  name: "",
  resultText: "Click to get CDP Neighbors from switch 👇",
})

function get_cdp_neighbors() {
  GetCDPNeighbors().then((res) => {
    data.resultText = res
  })
}

</script>

<template>
  <main class="main">
    <div id="result" class="result">{{ data.resultText }}</div>
    <div id="input" class="input-box">
      <button class="btn" @click="get_cdp_neighbors">Get CDP Neighbors</button>
    </div>
  </main>
</template>

<style scoped>

.main {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
}

.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 160px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
