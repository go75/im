<script lang="ts" setup>
import {reactive} from 'vue'
import {ref} from 'vue'

var name = ref('')

const emit = defineEmits(['page', "msg"])
function emitNew() {
  emit('page', "new")
}
function create() {
    emit('msg', 'new')
}
</script>

<template>
    <input type="text" placeholder="群聊名称">
    <br>
    <input type="text" placeholder="群聊简介">
    <br>
    <button >确认创建</button>
    <button @click="emitNew">关闭</button>
</template>

<style scoped>
  .start {
    background-image: url(https://img95.699pic.com/photo/40140/9844.gif_wh300.gif!/gifto/false);
    background-repeat: no-repeat;
    background-size:cover;
    color: white;
    height: 100%;
    width: 100%;
  }
  .input {
    width: 250px;
    height: 40px;
    margin-top: 50px;
    margin-bottom: 50px;
    text-align: center;
  }
  button {
    width: 80px;
    height: 40px;
    border-radius: 10px;
    border: 0;
    color: white;
    text-align: center;
    font-size: 20px;
    font-weight: bold;
    background-color: lightskyblue;
  }
</style>
