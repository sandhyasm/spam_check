<script setup>
import axios from 'axios'
import { ref } from 'vue'
const userName = ref('')
const email =  ref('')
const message = ref('')
const response = ref('')
const error = ref('')
defineProps({
  msg: String,
})

const checkSpamMessage = async () => {
  try {
    const res = await axios.post('http://localhost:8080/api/spam-check', {userName: userName.value, email: email.value, message: message.value})
    if (res.data.apiErr!='') {
      error.value = res.data.apiErr
    } else {
      response.value = res.data.spamResponse
    }
  } catch (error) {
    error.value = error
  }
}

const count = ref(0)
</script>

<template>
  <h1>{{ msg }}</h1>
  <div class="card">
    <div>
      <label for="name">Enter Name:  </label>
      <input type="text" v-model="userName">
    </div>
    <br>
    <div>
      <label for="email">Enter Email:  </label>
      <input type="text" v-model="email">
    </div>
    <br>
    <div>
      <label for="name">Enter Message:  </label>
      <input class="is-box" type="text" v-model="message">
    </div>
    <br>
    <button type="button" class="is-button" @click="checkSpamMessage">Send</button>
  </div>
  <div>
    <p v-if="error" class="is-error"> {{ error }}</p>
  </div>
</template>
<style scoped>
.read-the-docs {
  color: #888;
}
.is-flex {
  display: flex;
}
.is-button{
  background-color: #04AA6D; /* Green */
  border: none;
  color: white;
  padding: 15px 32px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
}
.is-box{
  width: 250px;
  height: 100px;
}
.is-error {
  color: red;
}
</style>
