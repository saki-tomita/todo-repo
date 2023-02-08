import App from "@/App.vue";
import login from "./components/login.vue";
import {createRouter, createWebHistory} from 'vue-router'

const routes = [
    {
        path: '/',
        name: 'login',
        component: login
    },
    {
        path: '/app',
        name: 'app',
        component: App

    },
    {
        path: "/../App.vue",
        component: App,
        name: "logapp",
        props: true
    },
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})



export default router