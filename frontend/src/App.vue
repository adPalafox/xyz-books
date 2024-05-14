<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <NavMenu />
        <el-page-header v-if="route.path.startsWith('/book')" @back="goBack">
          <template #content>
            <span class="text-large font-600 mr-3"> {{ route.name }} </span>
          </template>
        </el-page-header>
      </el-header>
      <el-main>
        <RouterView />
      </el-main>
      <el-footer>
        <Footer />
      </el-footer>
    </el-container>
  </div>
</template>

<script setup>
import NavMenu from '@/components/template/NavMenu.vue'
import Footer from '@/components/template/Footer.vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()

const goBack = () => {
  if (route.path.includes('/edit')) {
    router.back()
  }
  router.push('/')
}
</script>
