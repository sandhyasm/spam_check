<script setup>
import { ref } from 'vue'

defineProps({
  msg: String,
})

const count = ref(0)
</script>

<template>
  <h1>{{ msg }}</h1>
  <div class="card">
    <div>
      <label for="name">Enter Name:  </label>
      <input type="text">
    </div>
    <br>
    <div>
      <label for="email">Enter Email:  </label>
      <input type="text">
    </div>
    <br>
    <div>
      <label for="name">Enter Feedback:  </label>
      <input class="is-box" type="text">
    </div>
    <br>
    <button type="button" class="is-button" @click="">Send</button>
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
</style>
