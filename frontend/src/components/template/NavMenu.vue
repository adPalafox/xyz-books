<template>
  <el-menu class="el-menu" mode="horizontal" :ellipsis="false" @select="handleSelect">
    <el-row justify="space-evenly" :gutter="10">
      <el-col :span="8" :md="8">
        <router-link to="/">
          <img style="width: 190px" src="../../assets/xyz-logo-text.svg" alt="XYZ Books logo" />
        </router-link>
      </el-col>
      <!-- I am sorry I have no time to implement search-->
    </el-row>
  </el-menu>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

interface Book {
  title: string
  isbn_13: string
  publisher: string
  publication_year: string
  authors: string
}

const selectedBook = ref<Book | null>(null)

const handleSelect = (item: Book) => {
  selectedBook.value = item
}
</script>

<style scoped>
.el-menu {
  margin-top: 32px;
  display: block !important;
}
</style>
