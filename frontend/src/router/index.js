import { createWebHistory, createRouter } from 'vue-router'

import HomeView from '@/views/HomeView.vue'
import BookView from '@/views/BookView.vue'
import EditBook from '@/views/EditBook.vue'

const routes = [
  { path: '/', component: HomeView },
  { path: '/book', component: BookView },
  { path: '/book/edit', component: EditBook },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router;
