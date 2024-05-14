import { createWebHistory, createRouter } from 'vue-router'

import HomeView from '@/views/HomeView.vue'
import BookView from '@/views/BookView.vue'
import EditBook from '@/views/EditBook.vue'

const routes = [
  { name: 'Home', path: '/', component: HomeView },
  { name: 'Book', path: '/book/:isbn_13', component: BookView },
  { name: 'Edit', 
    path: '/book/:isbn_13/edit/:isbn_10/:id/:title/:publisher_id/:publisher/:publication_year/:edition?/:list_price', 
    component: EditBook },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router;
