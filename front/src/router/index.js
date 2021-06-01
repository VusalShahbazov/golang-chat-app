import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import {auth , noauth} from "../middleware/auth";

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home,
  },
  {
    path: '/logout',
    name: 'logout',
    component: () => import("../views/Dash/Pages/logout"),
  },
  {
    path: '/',
    component: ()=> import("../views/Dash/layout"),
    beforeEnter:auth,
    children: [
      {
        path: 'messages',
        name: 'messages',
        component: ()=> import("../views/Dash/Pages/messages"),
      },

    ]
  },
  {
    path: '/account',
    beforeEnter:noauth,
    component: () => import("../views/Auth/layout"),
    children:[
      {
        path: 'login',
        name: 'account-login',
        component: () => import("../views/Auth/Pages/login"),
      },
      {
        path: 'register',
        name: 'account-register',
        component: () => import("../views/Auth/Pages/register"),
      },
    ]
  },
  {
    path: '/errors',
    component: () => import("../views/Errors/layout.vue"),
    children: [
      {
        path: 'server-error',
        component: () => import("../views/Errors/Pages/server-error"),
        name:'server-error'
      },
      {
        path: 'not-found',
        component: () => import("../views/Errors/Pages/not-found"),
        name:'not-found'
      },
      {
        path: 'network-error',
        component: () => import("../views/Errors/Pages/network-error"),
        name:'network-error'
      },
    ]
  },
  {
    path: "*",
    component: () => import("../views/Errors/layout.vue"),
    children: [
      {
        path: '',
        component: () => import("../views/Errors/Pages/not-found"),
        name:'unexpected route'
      }
    ]
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
