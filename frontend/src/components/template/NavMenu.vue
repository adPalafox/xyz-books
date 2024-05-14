<template>
  <el-menu class="el-menu" mode="horizontal" :ellipsis="false" @select="handleSelect">
    <el-row justify="space-evenly" :gutter="10">
      <el-col :span="8" :md="8">
        <router-link to="/">
          <img style="width: 190px" src="../../assets/xyz-logo-text.svg" alt="XYZ Books logo" />
        </router-link>
      </el-col>
      <el-col :span="8" :md="8">
        <el-autocomplete
          v-model="searchTerm"
          :fetch-suggestions="fetchSuggestions"
          :trigger-on-focus="false"
          clearable
          class="inline-input w-50"
          placeholder="Search Books"
          @select="handleSelect"
        />
      </el-col>
    </el-row>
  </el-menu>
</template>

<script lang="ts" setup>
import { onMounted, ref, reactive } from 'vue'
import axios from 'axios' // Assuming Axios is installed

interface Book {
  title: string
  isbn_13: string
  publisher: string
  publication_year: string
  authors: string
}

const searchTerm = ref('')
const suggestions = ref<Book[]>([])
const selectedBook = ref<Book | null>(null) // To store selected book

const fetchSuggestions = async (queryString: string) => {
  if (!queryString.length) {
    suggestions.value = [] // Clear suggestions if empty string
    return
  }

  try {
    // const response = await axios.get(`http://your-backend-url/search`, {
    //   params: {
    //     query: queryString
    //   }
    // })
    // // Assuming "data" is the key in your backend response
    // suggestions.value = response.data.data.slice(0, 5) // Limit to top 5
    // suggestions
  } catch (error) {
    console.error('Error fetching suggestions:', error)
  }
}

const handleSelect = (item: Book) => {
  selectedBook.value = item
  // Perform further actions with selected book (e.g., navigate to details page)
}

const performFullSearch = async (queryString: string) => {
  // Implement your logic for performing the full search using the selected query
  // This could involve making a separate API call or filtering existing data
  // based on the searchTerm
}

onMounted(async () => {
  suggestions.value = [
    {
      title: 'The Lord of the Rings',
      isbn_13: '9780261102694',
      publisher: 'Houghton Mifflin Harcourt',
      publication_year: '1954',
      authors: 'J. R. R. Tolkien'
    },
    {
      title: 'Pride and Prejudice',
      isbn_13: '9780143995120',
      publisher: 'Penguin Classics',
      publication_year: '1813',
      authors: 'Jane Austen'
    },
    {
      title: 'To Kill a Mockingbird',
      isbn_13: '9780446310727',
      publisher: 'Harper Perennial Modern Classics',
      publication_year: '1960',
      authors: 'Harper Lee'
    },
    {
      title: 'One Hundred Years of Solitude',
      isbn_13: '9780307472708',
      publisher: 'Vintage Books',
      publication_year: '1967',
      authors: 'Gabriel García Márquez'
    },
    {
      title: 'The Great Gatsby',
      isbn_13: '9780743273565',
      publisher: 'Scribner',
      publication_year: '1925',
      authors: 'F. Scott Fitzgerald'
    }
  ]
})
</script>

<style scoped>
#search-wrap {
  width: 500px;
}
.el-menu {
  margin-top: 32px;
  display: block !important;
}
</style>
