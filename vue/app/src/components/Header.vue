<template>
  <v-row>

  <div class="header">

    <v-col cols="10">
      <div class="title">
        <h1>{{title}}</h1>
      </div>
    </v-col>

    <v-col cals="2">
      <div class="h_box">
        <v-btn @click="logout" class="links">ログアウト</v-btn>
      </div>
    </v-col>

  </div>
  </v-row>
</template>

<script>
export default {
  name: "header-todo",
  props: {
    title: String
  },
  methods: {
    logout(){
      this.$emit('log-out');
    }
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