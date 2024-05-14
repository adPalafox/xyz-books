<template>
  <el-card style="max-width: 1200px">
    <!-- <template #header>
      <div class="card-header">
        <span>Book Details</span>
      </div>
    </template> -->

    <el-descriptions class="margin-top" :title="book.title" :column="2" :size="size" border>
      <template #extra>
        <el-button type="primary">
          <router-link
            :to="{
              name: 'Edit',
              params: {
                isbn_13: book.isbn_13,
                isbn_10: book.isbn_10,
                id: book.id,
                title: book.title,
                publisher_id: book.publisher_id,
                publisher: book.publisher,
                publication_year: book.publication_year,
                edition: book.edition || '',
                list_price: book.list_price
              }
            }"
            >Edit Book</router-link
          >
        </el-button>
      </template>
      <el-descriptions-item>
        <template #label>
          <div class="cell-item">
            <el-icon :style="iconStyle">
              <iphone />
            </el-icon>
            ISBN 13
          </div>
        </template>
        {{ book.isbn_13 }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <div class="cell-item">
            <el-icon :style="iconStyle">
              <iphone />
            </el-icon>
            ISBN 10
          </div>
        </template>
        {{ book.isbn_10 }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <div class="cell-item">
            <el-icon :style="iconStyle">
              <location />
            </el-icon>
            Publisher
          </div>
        </template>
        {{ book.publisher }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <div class="cell-item">
            <el-icon :style="iconStyle">
              <location />
            </el-icon>
            Publication Year
          </div>
        </template>
        {{ book.publication_year }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <div class="cell-item">
            <el-icon :style="iconStyle">
              <location />
            </el-icon>
            Edition
          </div>
        </template>
        <template v-if="book.edition">
          {{ book.edition }}
        </template>
        <el-tag v-else :size="size">none</el-tag>
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <div class="cell-item">
            <el-icon :style="iconStyle">
              <tickets />
            </el-icon>
            Price
          </div>
        </template>
        <el-tag :size="size">{{ book.list_price }}</el-tag>
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <div class="cell-item">
            <el-icon :style="iconStyle">
              <office-building />
            </el-icon>
            Authors
          </div>
        </template>
        No.1188, Wuzhong Avenue, Wuzhong District, Suzhou, Jiangsu Province
      </el-descriptions-item>
    </el-descriptions>

    <template #footer>Last Edited on: </template>
  </el-card>
</template>

<script setup>
import { computed } from 'vue'
import { Iphone, Location, OfficeBuilding, Tickets, User } from '@element-plus/icons-vue'
const props = defineProps({
  book: Object
})

const book = props.book

console.log(book)

const size = 'large'
const iconStyle = 'default'

const blockMargin = computed(() => {
  const marginMap = {
    large: '32px',
    default: '28px',
    small: '24px'
  }
  return {
    marginTop: marginMap['small'] || marginMap.default
  }
})
</script>

<style scoped>
.el-descriptions {
  margin-top: 20px;
}
.cell-item {
  display: flex;
  align-items: center;
}
.margin-top {
  margin-top: 20px;
}
</style>
