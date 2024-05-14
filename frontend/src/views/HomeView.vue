<template>
  <div class="table-wrap" v-if="books.length > 0">
    <TableComponent v-bind:books="books" />
  </div>
  <div v-else>Loading books...</div>
  <div class="pagination-wrap">
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
  </div>
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
const pageSize1 = ref(10)
const pageSize2 = ref(50)
const pageSize3 = ref(100)
const pageSize4 = ref(500)
const small = ref(false)
const background = ref(false)
const disabled = ref(false)

async function fetchBooks(lengthInput, pageInput, sortInput, orderInput) {
  try {
    const baseUrl = 'http://localhost:8080/v1/api/book/list'
    const length = lengthInput
    const page = pageInput
    const sort = 'title'
    const order = 'asc'

    const formattedUrl = `${baseUrl}?length=${length}&page=${page}&sort=${sort}&order=${order}`
    const response = await axios.get(formattedUrl)
    books.value = response.data.data
    total_count.value = response.data.total_count
  } catch (error) {
    console.error('Error fetching books:', error)
  }
}

const handleSizeChange = (val) => {
  fetchBooks(val, currentPage1.value, 'title', 'asc')
}

const handleCurrentChange = (val) => {
  fetchBooks(pageSize1.value, val, 'title', 'asc')
}

fetchBooks(pageSize1.value, currentPage1.value, 'title', 'asc')
</script>

<style scoped>
.table-wrap {
  padding: 48px 32px;
  display: flex;
  justify-content: center;
}
.pagination-wrap {
  display: flex;
  justify-content: end;
  margin: 32px 136px;
}
</style>
