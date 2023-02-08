import { createApp } from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'
//import "./assets/main.css";
import 'firebase/auth'
import router from './router'

loadFonts()

createApp(App)
    .use(router)
    .use(vuetify)
    .mount('#app')

//const port = Number(process.env.PORT) || 8080;
//await app.listen(port, '0.0.0.0');


//firebase
//import firebase from 'firebase'
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
//import { getAuth, onAuthStateChanged } from "firebase/auth";




//firebaseAuth = getAuth();