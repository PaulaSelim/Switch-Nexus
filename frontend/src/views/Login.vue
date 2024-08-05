<script setup>
import { reactive } from "vue";
import InputText from "primevue/inputtext";
import Button from "primevue/button";
import "primeicons/primeicons.css";

const formData = reactive({
  username: "",
  password: "",
});

const handleLogin = () => {
  // Implement login logic here
  console.log("Login clicked", formData);
};

const handleForgotPassword = () => {
  // forgot password logic here
  console.log("Forgot Password clicked");
};
</script>

<template>
  <div class="login-container">
    <div class="login-card">
      <div class="logo-text">Switch Nexus</div>
      <div class="input-wrapper">
        <InputText
          v-model="formData.username"
          placeholder="USERNAME"
          class="input-field"
        />
        <i class="pi pi-user input-icon"></i>
      </div>
      <div class="input-wrapper">
        <InputText
          type="password"
          v-model="formData.password"
          placeholder="PASSWORD"
          class="input-field"
        />
        <i class="pi pi-lock input-icon"></i>
      </div>
      <Button @click="handleLogin" label="LOGIN" class="login-button" />
      <a href="#" class="forgot-password" @click="handleForgotPassword"
        >Forgot password?</a
      >
    </div>
  </div>
</template>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 95vh;
}

.login-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 300px;
}

.logo-text {
  font-size: 64px;
  font-weight: bold;
  color: black;
  margin-bottom: 40px;
  margin-top: -50px;
  text-align: center;
  line-height: 4;
  width: 500px;
}

.input-wrapper {
  position: relative;
  margin-bottom: 20px;
  border-radius: 4px;
  transition: border-color 0.2s, transform 0.2s ease-in-out,
    box-shadow 0.2s ease-in-out;
  width: 100%;
}

.input-wrapper:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.25);
}

.input-icon {
  position: absolute;
  left: 10px;
  top: 50%;
  transform: translateY(-50%);
  color: black;
}

.input-field {
  width: 250px;
  padding: 12px 12px 12px 35px;
  border-radius: 4px;
  font-size: 16px;
  border: 2px solid black;
}

.login-button {
  display: flex;
  align-items: center;
  margin-bottom: 5px;
  margin-top: 40px;
}

.login-button {
  width: 100%;
  padding: 12px;
  border-radius: 4px;
  border: none;
  font-size: 16px;
  font-weight: 600;
  background-color: #cfcfcf;
  transition: transform 0.2s ease-in-out, box-shadow 0.2s ease-in-out;
  cursor: pointer;
  display: flex;
  justify-content: center;
  align-items: center;
  color: #000;
  text-decoration: none;
  /* align-self: flex-end; */
}

.login-button:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.25);
}

.forgot-password {
  color: #000;
  font-size: 14px;
  text-decoration: none;
  cursor: pointer;
  align-self: flex-end;
  margin-top: 5px;
}

.forgot-password:hover {
  text-decoration: underline;
}
</style>
