<template>
  <div>
    <div class="p-grid p-jc-center">
      <div class="p-col-12 p-text-center">
        <h2>Task list</h2>
      </div>
      <div class="p-col-6">
        <form @submit.prevent="onSubmit">
          <div class="p-formgroup-inline p-ai-center">
            <div class="p-col">
              <label for="newtask" class="p-sr-only">Новая задача</label>
              <InputText
                id="newtask"
                type="text"
                placeholder="Add new task"
                class="taskText"
                v-model="newTask"
              />
            </div>
            <Button
              type="submit"
              class="p-button-raised p-button-outlined"
              label="add task"
            />
          </div>
        </form>
      </div>
    </div>
    <TaskList v-bind:tasks="tasks" v-bind:onDone="onDone" :onDelete="onDelete" />
    <Toast/>
  </div>
</template>

<script>
import TaskList from "./TaskList.vue";
import axios from "axios";
export default {
  mounted() {
    axios.post("/op", {}).then((res) => (this.tasks = res.data.all));
  },
  name: "TaskMain",
  components: { TaskList },
  props: {
    msg: String,
  },
  methods: {
    onSubmit() {
      if (this.newTask != "") {
        axios
          .create("create", { task: this.newTask })
          .then((res) => (this.tasks = res.data.all));
      }
      this.newTask = "";
    },
    onDone(id, done) {
      axios
        .patch("update", { id, done })
        .then((res) => (this.tasks = res.data.all));
    },
    onDelete(id, done){
      if(done){
        axios.delete("delete", {id, done})
        .then((res) => (this.tasks = res.data.all));
      } else {
        this.$toast.add({severity:'warn', summary:'Attention!',detail:'Complete the task first!', life: 3000 })
      }
    }
  },
  data() {
    return {
      tasks: [],
      newTask: "",
    };
  },
};
</script>

<style scoped>
.taskText {
  width: 100%;
}
</style>
