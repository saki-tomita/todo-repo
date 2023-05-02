<template>
   <v-card>
      <v-card-title>
        <span class="text-h5">タスク追加</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col
                cols="6"
            >
              <v-text-field
                  label="タスク名*"
                  required
                  v-model="SetVal.Name"
              ></v-text-field>
            </v-col>
            <v-col
                cols="6"
            >
              <v-select
                  label="ラベル*"
                  :items="labels"
                  v-model="SetVal.Label"
                  return-object>
              </v-select>
<!--              <v-text-field-->
<!--                  label="ラベル*"-->
<!--                  required-->
<!--                  v-model="SetVal.label"-->
<!--              ></v-text-field>-->
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="6">
              <Datepicker
                  v-model="SetVal.Deadline"
                  :enable-time-picker="false"
              ></Datepicker>
            </v-col>
            <v-col cols="6">
              <v-select
                  label="優先度*"
                  :items="ranks"
                  v-model="SetVal.Rank"
                  return-object>
              </v-select>
<!--              <v-text-field-->
<!--                  label="優先*"-->
<!--                  required-->
<!--                  v-model="SetVal.rank"-->
<!--              ></v-text-field>-->
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-text-field
                  label="メモ*"
                  required
                  v-model="SetVal.Memo"
              ></v-text-field>
            </v-col>
          </v-row>
        </v-container>
        <small>*indicates required field</small>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn
            color="blue darken-1"
            text
            @click="close()"
        >
          Close
        </v-btn>
        <v-btn
            color="blue darken-1"
            text
            @click="addTask()"
        >
          Save
        </v-btn>
      </v-card-actions>
    </v-card>
</template>

<script>
import axios from "axios";
import Datepicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css'

export default {
  name: "AddTask",
  props: ['task'],
  components: { Datepicker },

  data() {
    return {
      dialog: false,
      label: { title: '勉強', key: '1' },
      labels: [
        { title: '勉強', key: '1' },
        { title: '仕事', key: '2' },
        { title: '遊び', key: '3' }
      ],
      // SetVal: {
      //   name: "",
      //   label: "",
      //   deadline: "",
      //   rank: 1,
      //   memo.js: ""
      // }
      SetVal: this.task,
      ranks: [
        { title: '★', key: '1' },
        { title: '★★', key: '2' },
        { title: '★★★', key: '3' }
      ],
      rank: { title: '★', key: '1' }
    }
  },
  methods: {
    close(){
      this.$emit('close-dialog');
    },
    addTask(){
      if (this.task.ID == 0) {
        console.log('Add new task.')
        axios
            .post('/task', {
              id: 0,
              name: this.SetVal.Name,
              Status : 1,
              Rank : Number(this.SetVal.Rank.key),
              Deadline : this.SetVal.Deadline,
              Label : Number(this.SetVal.Label.key),
              Memo : this.SetVal.Memo,
              UserID : 1,
              DelFlg : 0,
              CreatedAt : "",
              CreatedUser : "",
              UpdatedAt : "",
              UpdatedUser : ""
            })
            .then(function (res) {
              console.log(res.data)
            })
        this.SetVal = {}
        this.$emit('close-dialog');
      } else {
        console.log('Edit task ID: ' + this.task.ID)
        axios
            .put('/task/'+this.SetVal.ID, {
              ID: Number(this.SetVal.ID),
              name: this.SetVal.Name,
              Status : 1,
              Rank : Number(this.SetVal.Rank.key),
              Deadline : this.SetVal.Deadline,
              Label : Number(this.SetVal.Label.key),
              Memo : this.SetVal.Memo,
              UserID : 1,
              DelFlg : 0,
              CreatedAt : "",
              CreatedUser : "",
              UpdatedAt : "",
              UpdatedUser : ""
            })
            .then(function (res) {
              console.log(res.data)
            })
        this.SetVal = {}
        this.$emit('close-dialog');
      }

    }
  }
}
</script>

<style scoped>

</style>