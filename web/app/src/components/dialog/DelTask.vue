<template>
   <v-card>
      <v-card-title>
        <span class="text-h5">Confirm</span>
      </v-card-title>
      <v-card-text class="pa-4">
        <p>タスク削除しますか？</p>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn
            color="blue darken-1"
            text
            @click="DelTask()"
        >
          OK
        </v-btn>
        <v-btn
            color="blue darken-1"
            text
            @click="cancel()"
        >
          キャンセル
        </v-btn>
      </v-card-actions>
    </v-card>
</template>

<script>
import axios from "axios";

export default {
  name: "DelTask",
  props: ['task'],

  data() {
    return {
      dialog: false,
      SetVal: this.task,
    }
  },
  methods: {
    cancel(){
      this.$emit('close-dialog');
    },
    DelTask(){
        console.log('Delete task.')
        axios
            .put('/task/del/'+this.SetVal.ID, {
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
</script>

<style scoped>

</style>