<template>
  <div>
    <el-form
      ref="ruleFormRef"
      style="max-width: 600px"
      :model="ruleForm"
      :rules="rules"
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
      <el-form-item label="Publisher" prop="publisher">
        <el-select v-model="ruleForm.publisher" placeholder="Publisher">
          <el-option label="Zone one" value="shanghai" />
          <el-option label="Zone two" value="beijing" />
        </el-select>
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
  </div>
</template>

<script lang="ts" setup>
import { ComponentSize, FormInstance, FormRules } from 'element-plus'
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
  publisher_id: string
  publisher: string
  publication_year: string
  edition: string
  price: string
}

const book = route.params
console.log(book)

const formSize = ref<ComponentSize>('default')
const ruleFormRef = ref<FormInstance>()
const ruleForm = reactive<BookForm>({
  id: route.params.id as string,
  title: route.params.title as string,
  isbn_13: route.params.isbn_13 as string,
  isbn_10: route.params.isbn_10 as string,
  publisher_id: route.params.publisher_id as string,
  publisher: route.params.publisher as string,
  publication_year: route.params.publication_year as string,
  edition: route.params.edition as string,
  price: route.params.list_price as string
})

const locationOptions = ['Home', 'Company', 'School']

const rules = reactive<FormRules<BookForm>>({
  title: [{ required: true, message: 'Please input book title', trigger: 'blur' }],
  isbn_13: [
    {
      required: true,
      message: 'Please enter a valid isbn 13',
      trigger: 'change'
      // validator: (rule, value: string, callback: (result: boolean) => void) => {
      //   if (value.length === 0) {
      //     callback(true) // Allow empty ISBN 13
      //     return
      //   }

      //   // Basic ISBN-13 validation (13 digits, check digit)
      //   const sum = value
      //     .split('')
      //     .reverse()
      //     .reduce((acc, digit, i) => acc + parseInt(digit, 10) * (i % 2 === 0 ? 1 : 3), 0)
      //   const checkDigit = (10 - (sum % 10)) % 10
      //   const valid = checkDigit === parseInt(value.charAt(12), 10)

      //   if (valid) {
      //     callback(true)
      //   } else {
      //     console.log(checkDigit)
      //     console.log(value)
      //     callback(new Error('Invalid ISBN-13 format'))
      //   }
      // }
    }
  ],
  isbn_10: [
    {
      required: false,
      message: 'Please enter a valid isbn 10',
      trigger: 'change'
      // validator: (rule, value: string, callback: (result: boolean) => void) => {
      //   if (value.length === 0) {
      //     callback(true) // Allow empty ISBN 10
      //     return
      //   }

      //   // Basic ISBN-10 validation (10 digits, check digit)
      //   // Replace with a more comprehensive ISBN-10 validation library if needed
      //   const sum = value
      //     .split('')
      //     .reverse()
      //     .slice(0, 9)
      //     .reduce((acc, digit, i) => acc + parseInt(digit, 10) * (i + 1), 0)
      //   const checkDigit = (11 - (sum % 11)) % 11
      //   const validCheckDigit = checkDigit === 10 ? 'X' : checkDigit.toString()
      //   const valid = validCheckDigit === value.charAt(9).toUpperCase()

      //   if (valid) {
      //     callback(true)
      //   } else {
      //     console.log(checkDigit)
      //     console.log(value)
      //     callback(new Error('Invalid ISBN-10 format'))
      //   }
      // }
    }
  ],
  publisher: [
    {
      required: true,
      message: 'Please select Activity count',
      trigger: 'blur'
    }
  ],
  publication_year: [
    {
      type: 'date',
      required: true,
      message: 'Please pick a date',
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
      message: 'Please select a location',
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
          list_price: parseFloat(ruleForm.price), // Assuming 'price' is intended for list_price
          publisher_id: parseInt(ruleForm.publisher_id, 10),
          publication_year: parseInt(ruleForm.publication_year, 10),
          publisher: ruleForm.publisher
        }

        const response = await axios.patch(
          `http://localhost:8080/v1/api/book/${isbn_13}`,
          editedBookData
        )

        if (response.status === 200) {
          console.log('Book updated successfully!', response.data)
          router.push('/book/' + ruleForm.isbn_13)
        } else {
          console.error('Error updating book:', response.data)
        }
      } catch (error) {
        console.error('Error submitting form:', error)
      }
    } else {
      console.log('error submit!', fields)
    }
  })
}

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
}

const options = Array.from({ length: 10000 }).map((_, idx) => ({
  value: `${idx + 1}`,
  label: `${idx + 1}`
}))
</script>
