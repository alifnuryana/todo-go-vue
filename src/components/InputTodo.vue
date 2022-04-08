<script setup>
import { computed } from "@vue/reactivity";
import { ref } from "vue";
import ChevronDoubleDownIcon from "./icon/ChevronDoubleDownIcon.vue";

const newTodo = ref("")

const isTyping = computed(() => {
  return newTodo.value.length > 0 ? true : false;
})

</script>

<template>
  <section id="header-todo" class="flex items-center border-b-2">
    <ChevronDoubleDownIcon class="hover:cursor-pointer ml-3 text-slate-700" />
    <input
      class="flex-1 p-3 outline-none text-xl"
      :class="{ 'italic': !isTyping }"
      placeholder="What needs to be done?"
      type="text"
      v-model="newTodo"
      name="newTodo"
      id="newTodo"
    />
  </section>
</template>