<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-row :gutter="20">
          <!--设备数据-->
          <el-col :span="20" :xs="24">
            <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
              <el-form-item label="设备类型" prop="typeID">
                <el-input
                  v-model="queryParams.typeName"
                  placeholder="请输设备类型"
                  clearable
                  size="small"
                  style="width: 160px"
                  @keyup.enter.native="handleQuery"
                />
              </el-form-item>
              <el-form-item>
                <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
                <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
              </el-form-item>

              <el-row :gutter="10" class="mb8">
                <el-col :span="1.5">
                  <el-button
                    v-permisaction="['admin:sysDevice:add']"
                    type="primary"
                    icon="el-icon-plus"
                    size="mini"
                    @click="handleAddDevice"
                  >新增</el-button>
                </el-col>
              </el-row>
            </el-form>


            <el-table
              v-loading="loading"
              :data="CategoryList"
              row-key="typeId"
              border
              :tree-props="{children: 'children', hasChildren: 'hasChildren'}"
            >
              <el-table-column label="类别名称"  prop="typeName" :show-overflow-tooltip="true" />
              <el-table-column label="配置类型" align="center" prop="config" :show-overflow-tooltip="true" />
              <el-table-column label="接入协议" align="center" prop="protocol" :show-overflow-tooltip="true" />
              <el-table-column label="默认配置" align="center"  width="250">
              <template slot-scope="scope">
                <el-tag  :type="scope.row.triggerStatus === 0 ? 'lightcoral' : 'success'" effect="dark" style="width: 80px;">
                  {{ scope.row.triggerStatus === 0 ? '否' : '是' }}
                </el-tag>
              </template>
              </el-table-column>
              <el-table-column label="LOGO" prop="typePath" :show-overflow-tooltip="true">
               <template slot-scope="scope">
                <img v-if="scope.row.logo" :src="getImagePath(scope.row.logo)" style="width: 50px; height: 50px;"  alt=""/>
               </template>
              </el-table-column>
              <el-table-column
                label="操作"
                width="200"
                fix="right"
                class-name="small-padding fixed-width"
                align="center"
              >
                <template slot-scope="scope">
                  <el-button
                    v-if="scope.row.userId !== 1"
                    v-permisaction="['admin:sysUser:remove']"
                    size="mini"
                    type="text"
                    icon="el-icon-edit"
                    @click="handleUpdate(scope.row)"
                  >修改</el-button>
                  <el-button
                    v-permisaction="['admin:sysUser:resetPassword']"
                    size="mini"
                    type="text"
                    icon="el-icon-delete"
                    @click="handleDelete(scope.row)"
                  >删除</el-button>
                </template>
              </el-table-column>
            </el-table>

            <pagination
              v-show="total>0"
              :total="total"
              :page.sync="queryParams.pageIndex"
              :limit.sync="queryParams.pageSize"
              @pagination="getList"
            />
          </el-col>
        </el-row>
      </el-card>
      <!-- 添加或修改参数配置对话框 -->
      <el-dialog :title="title" :visible.sync="open" width="600px" :close-on-click-modal="false">
        <el-form ref="form" :model="form"  label-width="80px">
          <el-row>
            <el-col :span="12">
              <el-form-item label="类别名称" prop="typeName">
                <el-input v-model="form.typeName" placeholder="请输入设备名称" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="归属类别" prop="parentId">
                <treeselect
                  v-model="form.parentId"
                  :options="locOptions"
                  :normalizer="normalizer"
                  :show-count="true"
                  placeholder="请选择归属区域"
                />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="配置类型" prop="config">
                <el-select v-model="form.config">
                  <el-option value="default" label="default"></el-option>
                  <el-option value="" label="none"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="接入协议" prop="protocol">
                <el-select v-model="form.protocol" :key="form.protocol">
                  <el-option label="TCP" value="TCP"></el-option>
                  <el-option label="无" value="none"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item  prop="roleId">
                <el-checkbox v-model="form.roleId" label="默认配置" />
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button type="primary" @click="submitForm">确 定</el-button>
          <el-button @click="cancel">取 消</el-button>
        </div>
      </el-dialog>
    </template>
  </BasicLayout>
</template>

<script>
// eslint-disable-next-line no-unused-vars
import { getToken } from '@/utils/auth'

import { listRole } from '@/api/admin/sys-role'
import Treeselect from '@riophae/vue-treeselect'
import '@riophae/vue-treeselect/dist/vue-treeselect.css'
import { listDevices } from '@/api/sys-device'
import {addCate, delCate, getCateList, listCategory, updateCate} from '@/api/sys-Category'
import {getDeptList} from "@/api/admin/sys-dept";

export default {
  name: 'SysDeviceManage',
  components: { Treeselect },
  data() {
    return {
      default:"default",
      编号: [],
      showPaslocOptionssword: false, // 控制密码可见性的变量，默认为 false
      locOptions: [],
      // 遮罩层
      loading: true,
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 设备类型表格数据
      CategoryList: [],
      // 弹出层标题
      title: '',
      // 类型树选项
      cateOptions: undefined,
      // 是否显示弹出层
      open: false,
      // 设备名称
      deviceName: undefined,
      // 设备类型
      deviceType: undefined,
      // 默认密码
      initPassword: undefined,
      // 日期范围
      dateRange: [],
      // 状态数据字典
      statusOptions: [],
      // 性别状态字典
      sexOptions: [],
      // 岗位选项
      postOptions: [],
      // 角色选项
      roleOptions: [],
      // 表单参数
      form: {
        typeID : undefined,
        typeName : undefined,
        parentId : undefined,
        config : undefined,
        protocol : undefined,
        roleId :undefined
      },
      defaultProps: {
        children: 'children',
        label: 'label'
      },
      // 用户导入参数
      upload: {
        // 是否显示弹出层（用户导入）
        open: false,
        // 弹出层标题（用户导入）
        title: '',
        // 是否禁用上传
        isUploading: false,
        // 是否更新已经存在的用户数据
        updateSupport: 0,
        // 设置上传的请求头部
        headers: { Authorization: 'Bearer ' + getToken() },
        // 上传的地址
        url: process.env.VUE_APP_BASE_API + '/system/user/importData'
      },
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        typeID: undefined,
        locID: undefined,
        deviceID: undefined,
        token: undefined,
        typeName: undefined
      }
    }
  },
  computed: {
    // 计算编号的计算属性
    indexedDeviceList() {
      // 根据列表数据和当前分页的索引计算编号
      const { DeviceList, queryParams: { pageIndex, pageSize }} = this
      return DeviceList.map((item, index) => {
        return {
          ...item, // 保留原有数据
          编号: (pageIndex - 1) * pageSize + index + 1 // 计算编号
        }
      })
    }
  },
  watch: {
    // 根据名称筛选部门树
    locName(val) {
      this.$refs.tree.filter(val)
    }
  },
  created() {
    this.getList()
    this.getDicts('sys_normal_disable').then(response => {
      this.statusOptions = response.data
    })
    this.getDicts('sys_user_sex').then(response => {
      this.sexOptions = response.daundefinedta
    })
    this.getConfigKey('sys_user_initPassword').then(response => {
      this.initPassword = response.data.configValue
    })
  },
  methods: {

    getImagePath(imageName) {
      return process.env.BASE_URL + imageName;
    },
    /** 密码点击可见 */
    togglePasswordVisibility() {
      this.showPassword = !this.showPassword // 切换 showPassword 的值
    },
    /** 查询用户列表 */
    getList() {
      this.loading = true
      listCategory(this.queryParams).then(response => {
        this.CategoryList = response.data
        this.loading = false
      }
      )
      console.log(this.CategoryList)
    },
    /** 转换菜单数据结构 */
    normalizer(node) {
      if (node.children && !node.children.length) {
        delete node.children
      }
      return {
        id: node.id,
        label: node.typeName,
        children: node.children
      }
    },
    /** 排序回调函数 */
    handleSortChang(column, prop, order) {
      prop = column.prop
      order = column.order
      if (this.order !== '' && this.order !== prop + 'Order') {
        this.queryParams[this.order] = undefined
      }
      if (order === 'descending') {
        this.queryParams[prop + 'Order'] = 'desc'
        this.order = prop + 'Order'
      } else if (order === 'ascending') {
        this.queryParams[prop + 'Order'] = 'asc'
        this.order = prop + 'Order'
      } else {
        this.queryParams[prop + 'Order'] = undefined
      }
      this.getList()
    },
    // 用户状态修改
    handleStatusChange(row) {
      const text = row.status === '2' ? '启用' : '停用'
      this.$confirm('确认要"' + text + '""' + row.username + '"用户吗?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return changeUserStatus(row)
      }).then(() => {
        this.msgSuccess(text + '成功')
      }).catch(function() {
        row.status = row.status === '2' ? '1' : '2'
      })
    },
    // 取消按钮
    cancel() {
      this.open = false
      this.resetDeviceForm()
    },
    // 表单重置
    resetDeviceForm() {
      this.form = {
        userId: undefined,
        locId: undefined,
        username: undefined,
        roleId: undefined,
        password: undefined,
        phone: undefined,
        email: undefined,
        sex: undefined,
        status: '2',
        remark: undefined,
        postIds: undefined,
        roleIds: undefined
      }
      this.resetForm('form')
    },
    /** 搜索按钮操作 */
    handleQuery() {
      console.log('search')
      this.queryParams.page = 1
      this.getList()
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.$refs.queryForm.resetFields()
      this.queryParams = {
        pageIndex: 1,
        pageSize: 10,
        typeID: '',
        locID: '',
        deviceID: ''
      }
      this.handleQuery()
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.userId)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    /** 新增按钮操作 */
    handleAddDevice() {
      this.Formreset()
      this.getTreeselect('Add', () => {
        this.open = true
        this.title = '添加设备类型'
      })
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.getTreeselect('update', () => {
        console.log(row)
        this.form.typeID = row.typeId || this.typeId
        this.form.typeName = row.typeName || ''
        this.form.parentId = row.parentId || this.parentId
        this.form.config = row.config || 'default'
        this.form.protocol = row.protocol || 'TCP'
        this.form.roleId = row.triggerStatus === 1
        this.open = true
        this.title = '修改类型'
      })
    },
    /** 提交按钮 */
    submitForm: function() {
        this.$refs['form'].validate(valid => {
        console.log(this.form)
        if (valid) {
          if (this.form.typeID === '') {
            console.log('调用addCate')
            addCate({parentId:  this.form.parentId, typeName: this.form.typeName , config: this.form.config , protocol: this.form.protocol}).then(response => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            console.log('调用update')
            updateCate({typeID :this.form.typeID, typeName: this.form.typeName , parentId:  this.form.parentId,  config: this.form.config , protocol: this.form.protocol},this.form.typeID).then(response => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          }
        }
      })
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      const Ids = (row.typeId && [row.typeId]) || this.typeId
      this.$confirm('是否确认删除用户编号为"' + Ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return delCate({ 'ids': Ids })
      }).then((response) => {
        if (response.code === 200) {
          this.msgSuccess(response.msg)
          this.open = false
          this.getList()
        } else {
          this.msgError(response.msg)
        }
      }).catch(function() {})
    },
    getTreeselect(type, callback) {
      getCateList().then(response => {
        this.locOptions = []
        if (type === 'Add') {
          const loc = { id: 0, typeName: '主类目', children: [] }
          loc.children = this.processData(response.data)
          this.locOptions.push(loc)
        } else {
          const loc = { id: 0, typeName: '主类目', children: [], isDisabled: true }
          loc.children = this.processData(response.data)
          this.locOptions.push(loc)
        }
        console.log(this.locOptions)
        if (callback && typeof callback === 'function') {
          callback()
        }
      })
    },
    processData(data) {
      return data.map(item => {
        return {
          ...item,
            id: item.typeId, // 确保每个节点有唯一的 id
          children: this.processData(item.children || [])
        }
      })
    },
    Formreset(){
      this.form.typeID = ''
      this.form.typeName = ''
      this.form.parentId = 0
      this.form.config = ''
      this.form.protocol = ''
      this.form.roleId =''
    }
  }
}
</script>

<style>

.el-table .el-table__header th .cell {
  text-align: center; /* 标题居中 */
}
</style>
