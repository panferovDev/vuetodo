<template>
  <li>
    <span v-bind:class="{ done: task.done }">
      <Checkbox
        class="check"
        v-on:change="$emit('onDone', task.id, task.done)"
        v-model="checked"
        :binary="true"
      />
      {{ task.title }}
    </span>

    <Button
      icon="pi pi-trash"
      iconPos="right"
      class="p-button-outlined p-button-rounded p-button-danger"
      v-on:click="$emit('onDelete', task.id, task.done)"
    />
  </li>
</template>

<script>
export default {
  props: ["task"],
  data() {
    return {
      checked: this.task.done,
    };
  },
};
</script>

<style scoped>
li {
  display: flex;
  justify-content: space-between;
  height: 60px;
  align-items: center;
  margin-bottom: 1rem;
  border-bottom: 1px solid rgb(228, 225, 225);
}
.check {
  margin-right: 1rem;
}
.done {
  text-decoration: line-through;
}
</style>