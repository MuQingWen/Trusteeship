<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
      <el-form-item label="班级序号:">
    <el-input v-model="formData.cno" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="中文名称:">
    <el-input v-model="formData.cname" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="班级学费:">
    <el-input v-model.number="formData.caccount" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="学生缴费:">
    <el-input v-model.number="formData.saccount" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="学生电话:">
    <el-input v-model="formData.sphone" clearable placeholder="请输入" />
    </el-form-item>
    <el-form-item>
        <el-button type="primary" @click="save">保存</el-button>
        <el-button type="primary" @click="back">返回</el-button>
      </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
  createAccount,
  updateAccount,
  findAccount
} from '@/api/account' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Account',
  mixins: [infoList],
  data() {
    return {
      type: '',
      formData: {
            cno: '',
            cname: '',
            caccount: 0,
            saccount: 0,
            sphone: '',
            
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findAccount({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.reaccount
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createAccount(this.formData)
          break
        case 'update':
          res = await updateAccount(this.formData)
          break
        default:
          res = await createAccount(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
      }
    },
    back() {
      this.$router.go(-1)
    }
  }
}
</script>

<style>
</style>