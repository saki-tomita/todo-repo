<template>
  <v-app>
    <v-main>
       <div class="route_box">
        <router-view/>
       </div>
      <Header
          v-if="usemail"
          title="ToDoApp"
          v-on:log-out="logout"
      />
      <div>{{usemail}}</div>
      <Search
          v-if="usemail"
          @get-search-res="getRes"
      />
<!--      <template v-slot:activator="{ on }">-->
        <v-btn
            v-if="usemail"
            icon color="indigo" v-bind="attrs" v-on="on" @click="add()">
          <v-icon>mdi-plus</v-icon>
        </v-btn>
<!--      </template>-->
      <v-dialog
          v-model="dialog"
          persistent
          max-width="600px"
      >
        <AddTask
            v-if="usemail"
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
            v-if="usemail"
            v-on:close-dialog="deldialog=false"
            :task="task"
        ></DelTask>
      </v-dialog>
      <MainTable
          v-if="usemail"
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

import "firebase/auth";
//import {auth} from "@/main";
require('firebase/auth')
import {onAuthStateChanged, getAuth, signOut} from "firebase/auth";
import {initializeApp} from "firebase/app";


//import { getAuth, onAuthStateChanged } from "firebase/auth";

export default {
  name: 'App',
  props: {usemail: String},

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
        },
      email: ""
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
    },
    logout(){
      signOut(this.auth)
    }
  },


  setup(){
    const firebaseConfig =  {
      apiKey: "AIzaSyA0pqtZ0MS4IIG7pTG4sBszchCJLuH4jKU",
      authDomain: "ca-saki-tomita-test.firebaseapp.com",
      projectId: "ca-saki-tomita-test",
      storageBucket: "ca-saki-tomita-test.appspot.com",
      messagingSenderId: "630134509832",
      appId: "1:630134509832:web:c7fe8919a892f81619dcbd"
    };

    const app = initializeApp(firebaseConfig);
    const auth = getAuth(app);

    onAuthStateChanged(auth, (user) => {
      if (user) {
        console.log("User is signed in")
        console.log(user);
        this.email = this.usemail;
        // User is signed in, see docs for a list of available properties
        // https://firebase.google.com/docs/reference/js/firebase.User


      } else {
        // User is signed out
        console.log("User is signed out");
      }
    });
  }
}
</script>

<style>
.route_box{
  padding: 10px 100px;
}

</style>