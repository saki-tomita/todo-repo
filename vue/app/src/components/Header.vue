<template>
  <div class="header">
    <div class="title">
      <h1>{{title}}</h1>
    </div>
    <div class="h_box">
      <v-btn to="../App" class="links">Todo</v-btn>
      <v-btn to="xxx" class="links">ログアウト</v-btn>
    </div>
  </div>
</template>

<script>
export default {
  name: "header-todo",
  props: {
    title: String
  }
}
</script>


<style scoped>
.h_box{
  text-align: left;
  margin: 0 2%;
}
.links{
  margin: 2%;
}

</style>