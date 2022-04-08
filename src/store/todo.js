import { defineStore } from "pinia";
import axios from "axios";

export const useTodoStore = defineStore("todo", {
  state: () => {
    return {
      token:
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImFsaWZuMjciLCJSb2xlIjoiYWRtaW4iLCJJZCI6IjIiLCJpc3MiOiJnby1hdXRoLWp3dCIsImV4cCI6MTY0OTQ1ODY3OCwiaWF0IjoxNjQ5NDQ3ODc4fQ.3SeyNfSWJl_SqE3CSsYLZn1Adcxp8OfqJitrjvIT5mc",
      todos: [],
      filteredState: "all",
    };
  },
  getters: {
    getTodos: (state) => {
      return state.todos;
    },
  },
  actions: {
    setFilteredState(to) {
      this.filteredState = to;
    },
    async fetchTodos() {
      axios
        .get("http://localhost:4000/api/todo", {
          headers: {
            Authorization: "Bearer " + this.token,
          },
        })
        .then((response) => {
          if (response.status === 200) {
            this.todos = response.data.data;
          }
        })
        .catch(function (error) {
          console.log(error);
        });
    },
    async setDone(from, id) {
      axios
        .put(
          `http://localhost:4000/api/todo/${id}`,
          {
            completed: !from,
          },
          {
            headers: {
              Authorization: "Bearer " + this.token,
            },
            data: {},
          }
        )
        .then((response) => {
          if (response.status === 200) {
            this.fetchTodos;
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
});
