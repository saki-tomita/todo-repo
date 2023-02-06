<template>
  <v-app>
    <v-main>

      <Header title="ToDoApp"/>
      <Search
          @get-search-res="getRes"
      />
<!--      <template v-slot:activator="{ on }">-->
        <v-btn icon color="indigo" v-bind="attrs" v-on="on" @click="add()">
          <v-icon>mdi-plus</v-icon>
        </v-btn>
<!--      </template>-->
      <v-dialog
          v-model="dialog"
          persistent
          max-width="600px"
      >
        <AddTask
            v-on:close-dialog="dialog=false"
            :task="task"
        ></AddTask>
      </v-dialog>
      <v-dialog
          v-model="deldialog"
          persistent
          max-width="400px"
      >
        <DelTask
            v-on:close-dialog="deldialog=false"
            :task="task"
        ></DelTask>
      </v-dialog>
      <MainTable
          v-on:open-edit-dialog="edit"
          v-on:open-delete-dialog="del"
          :taskRes="tasks"
      />
    </v-main>
  </v-app>
</template>

<script>
import Header from "@/components/Header.vue";
import AddTask from "@/components/dialog/AddTask.vue";
import DelTask from "@/components/dialog/DelTask.vue"
import Search from "@/components/Search.vue";
import MainTable from "@/components/MainTable.vue";

export default {
  name: 'App',

  components: {
    Header,
    AddTask,
    DelTask,
    Search,
    MainTable,
  },

  data() {
    return{
      tasks: [],
      id: 0,
      dialog: false,
      deldialog: false,
      task: {
          Name: "",
          Label: "",
          Deadline: "",
          Rank: 1,
          Memo: ""
        }
    }
  },

  methods: {
    getRes(value) {
      console.log("value:")
      console.log(value);
      this.tasks.splice(0, this.tasks.length)
      this.tasks.push(...value);
    },
    add(){
      this.id = 0;
      this.task = {}
      this.task.ID = 0;
      this.dialog = true;
    },
    edit(item){
      console.log("edit...")
      this.id = item.id;
      this.task = {}
      Object.assign(this.task, item)
      this.$nextTick(function() {
        this.dialog = true;
      });
    },
    del(item){
      this.id = item.id;
      this.task = {}
      Object.assign(this.task, item)
      this.deldialog = true;
    }
  }
}
</script>

<style>


</style>