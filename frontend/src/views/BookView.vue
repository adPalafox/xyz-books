<template>
  <BookCard />
  <button @click="goToEdit">Go Edit this book</button>
</template>

<script lang="ts" setup>
import BookCard from '@/components/BookCard.vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()

const goToEdit = () => {
  router.push('/book/edit')
}
</script>
