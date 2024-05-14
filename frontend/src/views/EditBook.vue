<template>
  <div class="form-wrap">
    <el-card shadow="never" style="width: 1200px; padding: 42px 36px">
      <el-form
        ref="ruleFormRef"
        :model="ruleForm"
        :rules="rules"
        label-position="right"
        style="max-width: 800px; margin: 0 auto"
        label-width="auto"
        class="demo-ruleForm"
        :size="formSize"
        status-icon
      >
        <el-form-item label="Title" prop="title">
          <el-input v-model="ruleForm.title" />
        </el-form-item>
        <el-form-item label="ISBN 13" prop="isbn_13">
          <el-input v-model="ruleForm.isbn_13" />
        </el-form-item>
        <el-form-item label="ISBN 10" prop="isbn_10">
          <el-input v-model="ruleForm.isbn_10" />
        </el-form-item>
        <el-form-item label="Publication Year" prop="publication_year">
          <el-input v-model="ruleForm.publication_year" />
        </el-form-item>
        <el-form-item label="Edition" prop="edition">
          <el-input v-model="ruleForm.edition" />
        </el-form-item>
        <el-form-item label="Price" prop="price">
          <el-input v-model="ruleForm.price" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm(ruleFormRef)"> Create </el-button>
          <el-button @click="resetForm(ruleFormRef)">Reset</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script lang="ts" setup>
import { ComponentSize, FormInstance, FormRules, ElNotification } from 'element-plus'
import axios from 'axios'
import { reactive, ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()

interface BookForm {
  id: string
  title: string
  isbn_13: string
  isbn_10: string
  publisher: string
  publication_year: string
  edition?: string
  price: string
}

const formSize = ref<ComponentSize>('large')
const ruleFormRef = ref<FormInstance>()
const ruleForm = reactive<BookForm>({
  id: route.params.id as string,
  title: route.params.title as string,
  isbn_13: route.params.isbn_13 as string,
  isbn_10: route.params.isbn_10 as string,
  publisher: route.params.publisher as string,
  publication_year: route.params.publication_year as string,
  edition: route.params.edition as string,
  price: route.params.list_price as string
})

const rules = reactive<FormRules<BookForm>>({
  title: [{ required: true, message: 'Please input book title', trigger: 'blur' }],
  isbn_13: [
    {
      required: true,
      message: 'Please enter a valid isbn 13',
      trigger: 'change'
    }
  ],
  isbn_10: [
    {
      required: false,
      message: 'Please enter a valid isbn 10',
      trigger: 'change'
    }
  ],
  publisher: [
    {
      required: true,
      message: 'Please select publisher',
      trigger: 'blur'
    }
  ],
  publication_year: [
    {
      required: true,
      message: 'Please enter valid year',
      trigger: 'change'
    }
  ],
  edition: [
    {
      required: false
    }
  ],
  price: [
    {
      required: true,
      message: 'Please enter a valid price',
      trigger: 'change'
    }
  ]
})

const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate(async (valid, fields) => {
    if (valid) {
      try {
        const { isbn_13 } = route.params
        const editedBookData = {
          id: parseInt(ruleForm.id, 10),
          title: ruleForm.title,
          isbn_13: ruleForm.isbn_13,
          isbn_10: ruleForm.isbn_10,
          list_price: parseFloat(ruleForm.price),
          publication_year: parseInt(ruleForm.publication_year, 10),
          publisher: ruleForm.publisher,
          edition: ruleForm.edition
        }

        const response = await axios.patch(
          `http://localhost:8080/v1/api/book/${isbn_13}`,
          editedBookData
        )

        if (response.status === 200) {
          const success = (message: string) => {
            ElNotification({
              title: 'Success',
              message: message,
              type: 'success'
            })
          }
          success('Book updated successfully!')
          router.push('/book/' + ruleForm.isbn_13)
        } else {
          const fail = (message: string) => {
            ElNotification({
              title: 'Error',
              message: message,
              type: 'error'
            })
          }
          fail('Error updating book')
        }
      } catch (error) {
        console.error('Error submitting form:', error)
        const fail = (message: string) => {
          ElNotification({
            title: 'Error',
            message: message,
            type: 'error'
          })
        }
        fail('Error updating book')
      }
    } else {
      console.log('error submit!', fields)
      const fail = (message: string) => {
        ElNotification({
          title: 'Error',
          message: message,
          type: 'error'
        })
      }
      fail('Error updating book')
    }
  })
}

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
}
</script>

<style scoped>
.form-wrap {
  margin: 64px 32px;
  display: flex;
  justify-content: center;
}
</style>
