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
        { title: 'NO', align: 'center', key: 'ID' },
        { title: "タスク名", key: "Name", align: "center" },
        { title: "ラベル", key: "Label_v", align: "center" },
        { title: "終了期限", key: "Deadline", align: "center" },
        { title: "ステータス", key: "Status_v", align: "center" },
        { title: "優先度", key: "Rank_v", align: "center" },
        { title: "操作", key: "actions" }
      ],
       items: this.taskRes
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