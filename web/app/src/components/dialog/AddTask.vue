<template>
   <v-card>
      <v-card-title>
        <span class="text-h5">タスク追加</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col
                cols="12"
            >
              <v-radio-group inline v-model="SetVal.Status">
                <v-radio label="未着手" value="1"/>
                <v-radio label="着手中" value="2"/>
                <v-radio label="完了" value="3"/>
              </v-radio-group>
            </v-col>
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
                  v-model="s_label"
                  return-object>
              </v-select>
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
                  v-model="s_rank"
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
      label: { title: '仕事', key: '1' },
      labels: [
        { title: '仕事', key: '1' },
        { title: '勉強', key: '2' },
        { title: '遊び', key: '3' }
      ],
      SetVal: this.task,
      s_label: this.task.ID == 0 ? '' : { title:this.task.Label_v, key:this.task.Label },
      s_rank: this.task.ID == 0 ? '' : { title:this.task.Rank_v, key:this.task.Rank },
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
              Rank : Number(this.s_rank.key),
              Deadline : this.SetVal.Deadline,
              Label : Number(this.s_label.key),
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
              Status : Number(this.SetVal.Status),
              Rank : Number(this.s_rank.key),
              Deadline : this.SetVal.Deadline,
              Label : Number(this.s_label.key),
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