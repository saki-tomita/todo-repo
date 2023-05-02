import { createApp } from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'
loadFonts()
import 'firebase/auth'
import router from './router'

createApp(App)
    .use(router)
    .use(vuetify)
    .mount('#app')

import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth"

const firebaseConfig =  {
    apiKey: "AIzaSyA0pqtZ0MS4IIG7pTG4sBszchCJLuH4jKU",
    authDomain: "ca-saki-tomita-test.firebaseapp.com",
    projectId: "ca-saki-tomita-test",
    storageBucket: "ca-saki-tomita-test.appspot.com",
    messagingSenderId: "630134509832",
    appId: "1:630134509832:web:c7fe8919a892f81619dcbd"
};

const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);