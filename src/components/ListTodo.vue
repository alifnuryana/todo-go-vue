<script setup>
import { computed } from "@vue/reactivity";
import { onMounted, ref } from "vue";
import { useTodoStore } from "../store/todo";
import CheckIcon from "./icon/CheckIcon.vue";
import CloseIcon from "./icon/CloseIcon.vue";

const todoStore = useTodoStore()

const filteredTodos = computed(() => {
  switch (todoStore.filteredState) {
    case "active":
      return todoStore.todos.filter((todo) => {
        if (todo.completed !== true) {
          return todo
        }
      })
    case "completed":
      return todoStore.todos.filter((todo) => {
        if (todo.completed === true) {
          return todo
        }
      })
    default:
      return todoStore.todos
  }
})

</script>

<template>
  <section id="body-todo">
    <ul>
      <li
        v-for="todo in filteredTodos"
        :key="todo.ID"
        class="flex items-center px-3 py-3 group border-b-2"
      >
        <div
          @click="todoStore.setDone(todo.completed, todo.ID)"
          class="w-7 h-7 border-2 border-slate-700 rounded-full hover:cursor-pointer flex justify-center items-center"
        >
          <CheckIcon
            v-show="todo.completed"
            class="w-5 h-5 transition-all text-green-500 hover:cursor-pointer"
          />
        </div>
        <p
          class="flex-1 ml-3 text-lg hover:cursor-pointer"
          v-bind:class="{ 'line-through': todo.completed }"
        >{{ todo.text }}</p>
        <CloseIcon
          class="hover:cursor-pointer group-hover:opacity-100 transition-all opacity-0 text-red-500"
        />
      </li>
    </ul>
  </section>
</template>