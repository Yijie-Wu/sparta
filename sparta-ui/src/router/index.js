import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/apply',
    name: 'apply',
    component: () => import('../views/ApplyView.vue')
  },
  {
    path: '/baseline',
    name: 'baseline',
    component: () => import('../views/BaselineView.vue')
  },
  {
    path: '/blogs',
    name: 'blogs',
    component: () => import('../views/BlogsView.vue')
  },
  {
    path: '/admin',
    name: 'admin',
    component: () => import('../views/AdminView.vue')
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
