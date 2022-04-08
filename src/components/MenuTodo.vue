<script setup>
import { computed } from "@vue/reactivity";
import { ref, watch } from "vue";
import { useTodoStore } from "../store/todo";

const filteredState = ref("all")

const todoStore = useTodoStore()

watch(filteredState, () => {
  todoStore.setFilteredState(filteredState.value)
  console.log("Berubah")
})

const activeTodos = computed(() => {
  const todos = todoStore.getTodos.filter((todo) => {
    if (todo.completed !== true) {
      return todo
    }
  })

  return todos.length
})
</script>

<template>
  <section id="footer-todo" class="flex items-center px-3 py-2 font-base">
    <p>{{ activeTodos }} items left</p>
    <div class="flex-1 flex justify-center gap-x-2">
      <button
        :class="{ 'border-2': filteredState === 'all' }"
        @click="filteredState = 'all'"
        class="px-2 py-1 rounded hover:border-2"
      >All</button>
      <button
        :class="{ 'border-2': filteredState === 'active' }"
        @click="filteredState = 'active'"
        class="px-2 py-1 rounded hover:border-2"
      >Active</button>
      <button
        :class="{ 'border-2': filteredState === 'completed' }"
        @click="filteredState = 'completed'"
        class="px-2 py-1 rounded hover:border-2"
      >Completed</button>
    </div>
    <button class="px-2 py-1 rounded hover:text-red-500 transition-all">Clear completed</button>
  </section>
</template>