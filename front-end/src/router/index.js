import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import FeedView from '../views/FeedView.vue'

const routes = [
  {path: '/', name: 'feed', component: FeedView},
  {path: '/category/:name',name:'category', component: () => import( '@/views/CategoryView.vue')},
  {path: '/post/:id', name:'post',component: () => import( '@/views/PostView.vue')},
  {path: '/newpost', name: 'newpost', component: () => import( '@/views/NewPostView.vue')},
  {path: '/profile', name: 'profile', component: () => import('@/views/ProfileView.vue')},
  {path: '/login',name: 'login', component: LoginView},
  {path: '/signup',name: 'signup', component: () => import('@/views/SignupView.vue')},
  {path: '/logout',name: 'logout', component: LoginView},
  {path: '/chat',name: 'chat', component: () => import('@/views/ChatView.vue')},
  {path: '/chat/:id',name: 'chatroom', component: () => import('@/views/ChatView.vue')},
  {path: '/:catchAll(.*)',name: 'notFound', component: () => import('@/views/NotFoundView.vue')}
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
