<template>
  <div v-if="books.length > 0">
    <TableComponent v-bind:books="books" />
  </div>
  <div v-else>Loading books...</div>
  <el-pagination background layout="prev, pager, next" :total="30" />
</template>

<script setup>
import TableComponent from '@/components/TableComponent.vue'
import axios from 'axios'
import { ref } from 'vue'

const books = ref([])

async function fetchBooks() {
  try {
    console.log('fetching books')
    const response = await axios.get(
      'http://localhost:8080/v1/api/book/list?length=30&page=1&sort=title&order=asc'
    )
    books.value = response.data.data
    console.log(books.value.data)
    console.log(typeof books.value)
  } catch (error) {
    console.error('Error fetching books:', error)
  }
}

fetchBooks()
</script>
