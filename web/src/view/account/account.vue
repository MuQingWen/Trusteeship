<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">  
        <el-form-item label="中文名称">
          <el-input placeholder="搜索条件" v-model="searchInfo.cname" />
        </el-form-item>    
        <el-form-item label="班级学费">
          <el-input placeholder="搜索条件" v-model="searchInfo.caccount" />
        </el-form-item>    
        <el-form-item label="学生缴费">
          <el-input placeholder="搜索条件" v-model="searchInfo.saccount" />
        </el-form-item>    
        <el-form-item label="学生电话">
          <el-input placeholder="搜索条件" v-model="searchInfo.sphone" />
        </el-form-item>    
        <el-form-item>
          <el-button type="primary" @click="onSubmit">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="openDialog">新增账目概览</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover v-model="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
              <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
            </div>
            <el-button slot="reference" icon="el-icon-delete" size="mini" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      ref="multipleTable"
      border
      stripe
      style="width: 100%"
      tooltip-effect="dark"
      :data="tableData"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column label="日期" width="180">
        <template slot-scope="scope">{{ scope.row.CreatedAt|formatDate }}</template>
      </el-table-column>
      <el-table-column label="班级序号" prop="cno" width="120" /> 
      <el-table-column label="中文名称" prop="cname" width="120" /> 
      <el-table-column label="班级学费" prop="caccount" width="120" /> 
      <el-table-column label="学生缴费" prop="saccount" width="120" /> 
      <el-table-column label="学生电话" prop="sphone" width="120" /> <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="updateAccount(scope.row)">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      layout="total, sizes, prev, pager, next, jumper"
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />
    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="班级序号:">
      
          <el-input v-model="formData.cno" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="中文名称:">
      
          <el-input v-model="formData.cname" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="班级学费:">
      
          <el-input v-model.number="formData.caccount" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="学生缴费:">
      
          <el-input v-model.number="formData.saccount" clearable placeholder="请输入" />
      </el-form-item>
        <el-form-item label="学生电话:">
      
          <el-input v-model="formData.sphone" clearable placeholder="请输入" />
      </el-form-item>
     </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="enterDialog">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  createAccount,
  deleteAccount,
  deleteAccountByIds,
  updateAccount,
  findAccount,
  getAccountList
} from '@/api/account' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList'
export default {
  name: 'Account',
  mixins: [infoList],
  data() {
    return {
      listApi: getAccountList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      
      formData: {
        cno: '',
          cname: '',
          caccount: 0,
          saccount: 0,
          sphone: '',
          
      }
    }
  },
  filters: {
    formatDate: function(time) {
      if (time !== null && time !== '') {
        var date = new Date(time);
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss');
      } else {
        return ''
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? '是' : '否'
      } else {
        return ''
      }
    }
  },
  async created() {
    await this.getTableData()
    
  },
  methods: {
  // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10     
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteAccount(row)
      })
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteAccountByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateAccount(row) {
      const res = await findAccount({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.reaccount
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        cno: '',
          cname: '',
          caccount: 0,
          saccount: 0,
          sphone: '',
          
      }
    },
    async deleteAccount(row) {
      const res = await deleteAccount({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === 1 && this.page > 1 ) {
          this.page--
        }
        this.getTableData()
      }
    },
    async enterDialog() {
      let res
      switch (this.type) {
        case "create":
          res = await createAccount(this.formData)
          break
        case "update":
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
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    }
  },
}
</script>

<style>
</style>
