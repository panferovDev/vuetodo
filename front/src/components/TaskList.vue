<template>
  <div class="p-grid p-jc-center">
    <div class="p-col-6">
      <ul>
        <TaskItem
          v-for="task of tasks"
          v-bind:task="task"
          v-bind:key="task.id"
          v-on:onDone="onDone"
          v-on:onDelete="onDelete"
        />
      </ul>
    </div>
  </div>
</template>

<script>
import TaskItem from "./TaskItem";
export default {
  name: "TaskList",
  components: {
    TaskItem,
  },
  props: ["tasks", "onDone", "onDelete"],
};
</script>

<style scoped>
ul {
  list-style: none;
  margin: 0px;
  padding: 0px;
}
</style>