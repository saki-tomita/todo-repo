<template>
    <v-app>
        <Header
            title="ToDoApp"
            v-on:log-out="logout"
        />
        <div>{{u_email}}</div>
        <Search
            @get-search-res="getRes"
        />
        <!--      <template v-slot:activator="{ on }">-->
        <v-btn
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
              v-on:close-dialog="closeAndRefresh"
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
  import router from './router.js'
  import axios from "axios";
  
  //import { getAuth, onAuthStateChanged } from "firebase/auth";
  export default {
    name: 'App',
    // router: Router,
    props: {u_email: String},
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
        // router: Router,
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
        // let vm = this
        for(let i = 0; i < this.tasks.length; i++) {
          if(this.tasks[i].Label == "1") {
            this.tasks[i].Label_v = "仕事"
          }else if(this.tasks[i].Label == "2") {
            this.tasks[i].Label_v = "勉強"
          }else if(this.tasks[i].Label == "3") {
            this.tasks[i].Label_v = "遊び"
          }
          if(this.tasks[i].Rank == "1") {
            this.tasks[i].Rank_v = "★"
          }else if(this.tasks[i].Rank == "2") {
            this.tasks[i].Rank_v = "★★"
          }else if(this.tasks[i].Rank == "3") {
            this.tasks[i].Rank_v = "★★★"
          }
          if(this.tasks[i].Status == "1") {
            this.tasks[i].Status_v = "未着手"
          }else if(this.tasks[i].Status == "2") {
            this.tasks[i].Status_v = "着手中"
          }else if(this.tasks[i].Status == "3") {
            this.tasks[i].Status_v = "完了"
          }
        }

        console.log("tasks:")
        console.log(this.tasks);
      },
      add(){
        this.id = 0;
        this.task = {}
        this.task.ID = 0;
        this.dialog = true;
      },
      edit(item){
        console.log("edit...")
        console.log(item)
        // this.no = item.no;
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
      closeAndRefresh() {
        this.dialog=false
        axios
            .get('/task', {
              params: {
                email: 'first_guest@test.test'
              }
            })
            .then((res) => {
              this.tasks = res.data
              this.getRes(this.tasks)
            })
      },
      logout(){
        console.log('logout...')
        const auth = getAuth()
        signOut(auth)
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
          this.email = user.email;
          // User is signed in, see docs for a list of available properties
          // https://firebase.google.com/docs/reference/js/firebase.User
        } else {
          // User is signed out
          console.log("User is signed out");
          router.push({path:'/'})
        }
      });
    }
  }
  </script>
  
  <style>
  .main{
    padding: 8% 10%;
  }
  </style>