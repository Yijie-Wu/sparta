import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/Main/HomeView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/apply',
    name: 'apply',
    component: () => import('../views/Main/ApplyView.vue')
  },
  {
    path: '/baseline',
    name: 'baseline',
    component: () => import('../views/Main/BaselineView.vue')
  },
  {
    path: '/blogs',
    name: 'blogs',
    component: () => import('../views/Main/BlogsView.vue')
  },
  {
    path: '/notifications',
    name: 'notifications',
    component: () => import('../views/User/NotificationView.vue')
  },
  {
    path: '/reviews',
    name: 'reviews',
    component: () => import('../views/Main/ReviewView.vue')
  },
  {
    path: '/admin/applies',
    name: 'admin-applies',
    component: () => import('../views/Admin/ApplyManagerView.vue')
  },
  {
    path: '/admin/users',
    name: 'admin-users',
    component: () => import('../views/Admin/UserManagerView.vue')
  },
  {
    path: '/admin/projects',
    name: 'admin-projects',
    component: () => import('../views/Admin/ProjectsManagerView')
  },
  {
    path: '/admin/flows',
    name: 'admin-flows',
    component: () => import('../views/Admin/FlowsManagerView')
  },
  {
    path: '/admin/spiders',
    name: 'admin-spiders',
    component: () => import('../views/Admin/SpiderManagerView')
  },
  {
    path: '/admin/schedulers',
    name: 'admin-schedulers',
    component: () => import('../views/Admin/ScheduleManagerView')
  },
  {
    path: '/admin/web/settings',
    name: 'admin-web-settings',
    component: () => import('../views/Admin/WebSettingsView')
  },
  {
    path: '/user/applies',
    name: 'user-applies',
    component: () => import('../views/User/YourApplyView')
  },
  {
    path: '/user/flows',
    name: 'user-flows',
    component: () => import('../views/User/YourFlowsView')
  },
  {
    path: '/user/blogs',
    name: 'user-blogs',
    component: () => import('../views/User/YourBlogsView')
  },
  {
    path: '/user/profile',
    name: 'user-profile',
    component: () => import('../views/User/YourProfileView')
  },
  {
    path: '/auth/login',
    name: 'auth-login',
    component: () => import('../views/Auth/LoginView')
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
