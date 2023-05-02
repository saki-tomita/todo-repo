<template>
  <div class="signup">
    <table>
      <tr>
        <th>メールアドレス：</th>
      </tr>
      <tr>
        <td><input type="email" v-model="mailaddress"/></td>
      </tr>
      <tr>
        <th>パスワード：</th>
      </tr>
      <tr>
        <td><input type="password" v-model="password"/></td>
      </tr>
    </table>
    <button @click="signUp">登録</button>
  </div>
</template>

<script>
export default {
  name: 'SignUpTodo',
}

//  methods: {
//    signUp: function () {
//      firebase.auth().createUserWithEmailAndPassword(this.mailaddress, this.password)
//          .then(
//              console.log('Success regist.')
//          )
//          .catch(
//              console.log('Error regist')
//          )
//    }
//  },
//
//  mounted(){
//    firebase.auth().onAuthStateChanged(function(user) {
//      if (user) {
//        console.log('login');
//      } else {
//        console.log('logout');
//      }
//    });
//  }
</script>