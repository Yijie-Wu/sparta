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
    path: '/notifications',
    name: 'notifications',
    component: () => import('../views/NotificationView.vue')
  },
  {
    path: '/reviews',
    name: 'reviews',
    component: () => import('../views/ReviewView.vue')
  },
  {
    path: '/admin/applies',
    name: 'admin-applies',
    component: () => import('../views/ApplyManagerView.vue')
  },
  {
    path: '/admin/users',
    name: 'admin-users',
    component: () => import('../views/UserManagerView.vue')
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
