<template>
  <div class="book-card-wrap">
    <BookCard v-if="book" :book="book" />
    <div v-else>Loading books...</div>
  </div>
</template>

<script setup>
import BookCard from '@/components/BookCard.vue'
import { useRoute } from 'vue-router'
import axios from 'axios'
import { ref } from 'vue'

const route = useRoute()

const book = ref(null)

async function fetchBook(isbn13) {
  try {
    console.log('fetching book ' + isbn13)
    const response = await axios.get('http://localhost:8080/v1/api/book/' + isbn13)
    book.value = response.data.data
    console.log(book.value)
    console.log(typeof book.value)
  } catch (error) {
    console.error('Error fetching books:', error)
  }
}

fetchBook(route.params.isbn_13)
</script>

<style scoped>
.book-card-wrap {
  margin: 64px 32px;
  display: flex;
  justify-content: center;
}
</style>