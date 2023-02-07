import { createApp } from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'
//import "./assets/main.css";

loadFonts()

createApp(App)
  .use(vuetify)
  .mount('#app')

//const port = Number(process.env.PORT) || 8080;
//await app.listen(port, '0.0.0.0');

