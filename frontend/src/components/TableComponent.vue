<template>
  <el-table :data="books" style="width: 90%" :table-layout="auto">
    <el-table-column fixed prop="title" label="Title" width="160" />
    <el-table-column prop="isbn_13" label="ISBN 13" width="160" />
    <el-table-column prop="isbn_10" label="ISBN 10" width="160" />
    <el-table-column prop="list_price" label="Price" width="120" />
    <el-table-column prop="publisher" label="Publisher" show-overflow-tooptip="true" />
    <el-table-column fixed="right" label="Operation" width="120">
      <template #default="scope">
        <el-button link type="primary" size="large" @click="handleView(scope.row)">
          Details
        </el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const props = defineProps({
  books: Object
})

const handleView = (selectedRow) => {
  router.push('/book/' + selectedRow.isbn_13)
}
</script>
