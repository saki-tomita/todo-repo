<template>
    <v-container>

      <div id="maintbl">
        <v-container class="container">
          <v-data-table :headers="headers" :items="items">
            <template v-slot:[`item.actions`]="{ item }">
                <!-- 編集ボタン -->
                <v-icon @click="onClickEditItem(item.value)">
                  mdi-pencil
                </v-icon>
                <!-- 削除ボタン -->
                <v-icon @click="onClickDeleteItem(item.value)">
                  mdi-trash-can
                </v-icon>
            </template>
          </v-data-table>
        </v-container>
      </div>
    </v-container>
</template>

<script>
export default {
  name: 'MainTable',
  props: ['taskRes'],
  data () {
    return {

      headers: [
        { title: 'NO', align: 'center', key: 'no' },
        { title: "タスク名", key: "Name", align: "center" },
        { title: "ラベル", key: "Label", align: "center" },
        { title: "終了期限", key: "Deadline", align: "center" },
        { title: "ステータス", key: "Status", align: "center" },
        { title: "優先度", key: "Rank", align: "center" },
        { title: "操作", key: "actions" }
      ],
      //  items: this.taskRes
      items: [
        {no:"1", Name:"Toeic勉強", Label:"勉強", Deadline:"2023-01-01", Status:"着手中", Rank:"★"},
        {no:"2", Name:"Toeic勉強", Label:"勉強", Deadline:"2023-01-01", Status:"着手中", Rank:"★"}
      ]
    }

},

  methods: {
    onClickEditItem: function(item) {
      console.log(item)
      this.$emit('open-edit-dialog', item);
    },
    onClickDeleteItem: function(item) {
      console.log(item.Name);
      this.$emit('open-delete-dialog', item);
    },
  },

  computed: {
    list() {
      return []
    }
  }
}
</script>

<style scoped>

</style>