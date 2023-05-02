<template>
  <div class="p_siginin">
  <div class="signin">
    <v-btn @click="authMethod">ログイン</v-btn>
    <div class="new_text">
      <p>新しいアカウントを作成しますか？
        <router-link to="/signup">新規登録</router-link>
      </p>
    </div>
  </div>
  </div>
</template>

<script>
import router from '../router'
import {
  getAuth,
  //onAuthStateChanged,
  GoogleAuthProvider,
  signInWithPopup,
} from "firebase/auth"
import {initializeApp} from "firebase/app";
import {reactive} from "vue";

export default {
  name: 'SignInTodo',
  props:{
  },

  data() {
    return{
      email_address: ''
    }
  },

  setup() {
    const authMethod = () => {

      const firebaseConfig = {
        apiKey: "AIzaSyA0pqtZ0MS4IIG7pTG4sBszchCJLuH4jKU",
        authDomain: "ca-saki-tomita-test.firebaseapp.com",
        projectId: "ca-saki-tomita-test",
        storageBucket: "ca-saki-tomita-test.appspot.com",
        messagingSenderId: "630134509832",
        appId: "1:630134509832:web:c7fe8919a892f81619dcbd"
      };
      initializeApp(firebaseConfig);

      const provider = new GoogleAuthProvider()
      const auth = getAuth();
      signInWithPopup(auth, provider).then((result) => {
        //console.log(result.user)
        const user = result.user
        data.email_address = user.email
        console.log(user.email)

        router.push({name: 'logapp', params:{usemail: user.email}})
        this.$router.push('/mains')
      }).catch((error) => {
        console.log(error)
      })
    }

    const data = reactive({
      email_address: ''
    })
    return {
      data, authMethod
    }
  }


//   mounted: {
//     onAuthStateChanged(auth, (user) => { // <4>
//   if (user) {
//     el.sectionSignin.style.display = 'none'
//     el.sectionUser.style.display = 'block'
//     el.sectionSignout.style.display = 'block'
//     el.uid.innerHTML = user.uid
//   } else {
//     el.sectionSignin.style.display = 'block'
//     el.sectionUser.style.display = 'none'
//     el.sectionSignout.style.display = 'none'
//   }
// })
// }
}


</script>

<style>
.p_siginin{
  display: flex;
  justify-content: center;
  padding: 50px 100px;
}
.signin{
  width: 600px;
  height: 300px;
  margin: 50px auto;
}

.new_text{
  margin: 20px 0;
}
</style>