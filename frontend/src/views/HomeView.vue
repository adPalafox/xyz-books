<template>
  <div v-if="books.length > 0">
    <TableComponent v-bind:books="books" />
  </div>
  <div v-else>Loading books...</div>
  <el-pagination
    v-model:current-page="currentPage1"
    v-model:page-size="pageSize1"
    :page-sizes="[10, 100, 500]"
    :small="small"
    :disabled="disabled"
    :background="background"
    layout="total, sizes, prev, pager, next, jumper"
    :total="total_count"
    @size-change="handleSizeChange"
    @current-change="handleCurrentChange"
  />
</template>

<script setup>
import TableComponent from '@/components/TableComponent.vue'
import axios from 'axios'
import { ref } from 'vue'

const books = ref([])
const total_count = ref(1)

const currentPage1 = ref(1)
const currentPage2 = ref(2)
const currentPage3 = ref(3)
const currentPage4 = ref(4)
const pageSize1 = ref(1)
const pageSize2 = ref(10)
const pageSize3 = ref(100)
const pageSize4 = ref(500)
const small = ref(false)
const background = ref(false)
const disabled = ref(false)

async function fetchBooks(lengthInput, pageInput, sortInput, orderInput) {
  try {
    console.log('fetching books')
    
    const baseUrl = 'http://localhost:8080/v1/api/book/list'
    const length = lengthInput // Number of items per page
    const page = pageInput // Current page number
    const sort = 'title' // Sort by field (e.g., 'title', 'author')
    const order = 'asc' // Sort order (e.g., 'asc', 'desc')

    const formattedUrl = `${baseUrl}?length=${length}&page=${page}&sort=${sort}&order=${order}`
    const response = await axios.get(
      formattedUrl
    )
    books.value = response.data.data
    total_count.value = response.data.total_count
  } catch (error) {
    console.error('Error fetching books:', error)
  }
}

const handleSizeChange = (val) => {
  console.log(`${val} items per page`)
  fetchBooks(val, currentPage1.value, 'title', 'asc')
}

const handleCurrentChange = (val) => {
  // console.log('handle')
  // console.log(typeof val === 'number')
  // console.log(`current page: ${val}`)
  // console.log(typeof val) // number
  // console.log(val.target)
  fetchBooks(pageSize1.value, val, 'title', 'asc')
}

fetchBooks(pageSize1.value, currentPage1.value, 'title', 'asc')
</script>
