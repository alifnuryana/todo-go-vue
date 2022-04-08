<script setup>
import { defineStore } from "pinia";
import { onMounted } from "vue";
import InputTodo from "./components/InputTodo.vue";
import ListTodo from "./components/ListTodo.vue";
import MenuTodo from "./components/MenuTodo.vue";
import { useTodoStore } from "./store/todo";

const todoStore = useTodoStore()

onMounted(() => {
  todoStore.fetchTodos();
})

</script>

<template>
  <section class="max-w-lg w-full mx-auto my-5">
    <h1 class="text-2xl font-bold text-red-500 my-5">Simple Todo</h1>
    <section id="todo" class="shadow-lg border-2 rounded">
      <InputTodo />
      <ListTodo />
      <MenuTodo />
    </section>
  </section>
</template>

