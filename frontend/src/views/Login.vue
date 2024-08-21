<script setup>
import { reactive } from "vue";
import InputText from "primevue/inputtext";
import Button from "primevue/button";
import { useRouter } from "vue-router";

const router = useRouter();

const formData = reactive({
  username: "",
  password: "",
});

const data = reactive({
  username: "",
  password: "",
  response: "",
});

const handleLogin = () => {
  if (data.username === "admin" && data.password === "admin") {
    router.push("/home");
  } else {
    data.response = "Invalid username or password";
    return;
  }
};
</script>

<template>
  <div class="login-container">
    <div class="login-card">
      <div class="logo-text">Switch Nexus</div>
      <div class="input-wrapper">
        <i class="pi pi-user input-icon"></i>
        <InputText
          v-model="data.username"
          placeholder="USERNAME"
          class="input-field"
        />
      </div>
      <div class="input-wrapper">
        <i class="pi pi-lock input-icon"></i>
        <InputText
          type="password"
          v-model="data.password"
          placeholder="PASSWORD"
          class="input-field"
        />
      </div>
      <p class="response">{{ data.response }}</p>
      <Button @click="handleLogin" label="LOGIN" class="login-button" />
    </div>
  </div>
</template>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.login-card {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.logo-text {
  font-size: 32px;
  font-weight: bold;
  margin-bottom: 25px;
}

.input-wrapper {
  margin-bottom: 20px;
  border-radius: 4px;
  transition: border-color 0.2s, transform 0.2s ease-in-out,
    box-shadow 0.2s ease-in-out;
}

.input-wrapper:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.25);
}

.input-icon {
  position: absolute;
  padding: 5px 10px;
  transform: translateY(9px);
}

.input-field {
  width: 100%;
  padding: 12px 12px 12px 35px;
  border-radius: 4px;
  font-size: 16px;
  border: 2px solid black;
}

.login-button {
  width: 100%;
  padding: 12px;
  transition: transform 0.2s ease-in-out, box-shadow 0.2s ease-in-out;
}

.login-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.25);
}

.response {
  color: red;
}
</style>
