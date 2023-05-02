// import Vue from 'vue'
// import { createApp } from 'vue'
import App from "./App.vue";
import login from "./components/login.vue";
import Mains from "./Mains.vue";
import {createRouter, createWebHistory} from 'vue-router'
// import Router from 'vue-router'
// import VueRouter from "vue-router";

// createApp(App)
//     .use(VueRouter)


const routes = [
    {
        path: '/',
        name: 'login',
        component: login
    },
    // {
    //     path: '/',
    //     name: 'Mains',
    //     component: Mains

    // },
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
    {
        path: "/mains",
        component: Mains,
        name: "logapp",
        props: true
    },
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})
// const router = new VueRouter({
//     routes
//     })
export default router

// export default new Router({
//     mode: 'history',
//     base: process.env.BASE_URL,
//     routes: routes
// })