<template>
      <v-container fluid fill-height>
        <v-row>
          <v-col cols="4">
            <v-text-field v-model="searchInput" class="searchInput" solo/>
          </v-col>
          <v-col cols="5">
            <v-radio-group inline v-model="searchStatus">
              <v-radio label="未着手" value="1"/>
              <v-radio label="着手中" value="2"/>
              <v-radio label="完了" value="3"/>
            </v-radio-group>
          </v-col>
        </v-row>
        <v-row class="justify-end">
          <v-btn v-on:click="search">検索</v-btn>
        </v-row>
      </v-container>

</template>

<script>
import axios from 'axios'

export default {
  name: 'search-task',
  props: {usemail: String},
  data(){
    return {
      searchInput: null,
      searchStatus: '',
      searchMemo: false,
      tasks:[]
    }
  },
  mounted() {
        axios
            .get('/task', {
              params: {
                email: 'first_guest@test.test'
              }
            })
            .then((res) => {
              this.tasks = res.data
              this.$emit('get-search-res', this.tasks)
            })
  },
  methods:{
    search: function(){
      if(this.searchStatus == '') {
        this.searchStatus = '1'
      }
      let params = {'title': this.searchInput, 'status': this.searchStatus}
      if(this.searchInput == '' || this.searchInput == null) {
        axios
            .get('/task', {
              params: {
                email: 'first_guest@test.test'
              }
            })
            .then((res) => {
              this.tasks = res.data
              this.$emit('get-search-res', this.tasks)
            })
      } else {
        axios
            .get('/task/1', {
              params: {
                name: params.title,
                status: params.status
              }
            })
            .then((res) => {
              console.log(res)
              this.tasks = res.data
              console.log(this.tasks)
              this.$emit('get-search-res', this.tasks)
            })
      }
    },
  }
}

</script>

<style>


</style>