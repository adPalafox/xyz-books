<template>
  <el-card style="width: 1200px">
    <el-row justify="space-evenly" :gutter="8">
      <el-col :span="8" :md="4">
        <div style="image-wrap">
          <el-image
            style="width: 100%; height: 240px; margin-top: 72px"
            :src="book.image_url"
            fit="cover"
          />
        </div>
      </el-col>
      <el-col :span="16" :md="18">
        <el-descriptions class="margin-top" :title="book.title" :column="1" :size="size" border>
          <template #extra>
            <router-link
              :to="{
                name: 'Edit',
                params: {
                  isbn_13: book.isbn_13,
                  isbn_10: book.isbn_10 || '',
                  id: book.id,
                  title: book.title,
                  publisher: book.publisher,
                  publication_year: book.publication_year,
                  edition: book.edition || '',
                  list_price: book.list_price,
                  image_url: book.image_url || ''
                }
              }"
            >
              <el-button type="primary"> Edit Book </el-button>
            </router-link>
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
            <ul>
              <li v-for="author in book.authors" :key="author.id">
                {{ author.first_name + ' ' + author.last_name }}
              </li>
            </ul>
          </el-descriptions-item>
        </el-descriptions>
      </el-col>
    </el-row>
  </el-card>
</template>

<script setup>
import { Iphone, Location, OfficeBuilding, Tickets } from '@element-plus/icons-vue'
const props = defineProps({
  book: Object
})

const book = props.book

const size = 'large'
const iconStyle = 'default'
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
.image-wrap {
  margin-top: 64px;
}
</style>
