<template>
  <el-table :data="books" style="width: 100%">
    <el-table-column fixed prop="title" label="Title" width="150" />
    <el-table-column prop="isbn_13" label="ISBN 13" width="130" />
    <el-table-column prop="isbn_10" label="ISBN 10" width="120" />
    <el-table-column prop="list_price" label="Price" width="120" />
    <el-table-column prop="publisher" label="Publisher" width="600" />
    <el-table-column fixed="right" label="Operations" width="120">
      <template #default="scope">
        <el-button link type="primary" size="small" @click="handleClick(scope.row)">
          Details
        </el-button>
        <el-button link type="primary" size="small">Edit</el-button>
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

const handleClick = (selectedRow) => {
  console.log(selectedRow.isbn_13)
  router.push('/book/' + selectedRow.isbn_13)
}

console.log(props.books[0].authors)
</script>
