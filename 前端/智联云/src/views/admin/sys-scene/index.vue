<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-row :gutter="20">
          <!--区域树搜索-->
          <el-col :span="4" :xs="24">
            <div class="head-container">
              <el-input v-model="treeInputvalue" placeholder="请输入区域名称" clearable size="small"
                prefix-icon="el-icon-search" style="margin-bottom: 20px" />
            </div>
            <div class="head-container">
              <el-tree ref="tree" :data="locOptions" :props="defaultProps" :expand-on-click-node="false"
                :filter-node-method="filterNode" default-expand-all @node-click="handleNodeClick" />
            </div>
          </el-col>
          <!--两个输入框-->
          <el-col :span="20" :xs="24">
            <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
              <el-form-item label="场景名称" prop="deviceName">
                <el-input v-model="queryParams.deviceName" placeholder="请输入场景名称" clearable size="small"
                  style="width: 160px" @keyup.enter.native="handleQuery" />
              </el-form-item>
              <el-form-item label="区域名称" prop="typeName">
                <el-input v-model="queryParams.type.typeName" placeholder="请输区域名称" clearable size="small"
                  style="width: 160px" @keyup.enter.native="handleQuery" />
              </el-form-item>
              <el-form-item>
                <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
                <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
              </el-form-item>
            </el-form>

            <el-row :gutter="10" class="mb8">
              <el-col :span="1.5">
                <el-button v-permisaction="['admin:sysDevice:edit']" type="success" icon="el-icon-edit" size="mini"
                  @click="handleAddScene">新增</el-button>
              </el-col>

              <el-col :span="1.5">
                <el-button v-permisaction="['admin:sysDevice:edit']" type="success" icon="el-icon-edit" size="mini"
                  :disabled="single" @click="handleUpdate">修改</el-button>
              </el-col>

              <el-col :span="1.5">
                <el-button v-permisaction="['admin:sysDevice:remove']" type="danger" icon="el-icon-delete" size="mini"
                  :disabled="multiple" @click="handleDelete">删除</el-button>
              </el-col>
            </el-row>


            <el-table v-loading="loading" :data="tempViceList" border @selection-change="handleSelectionChange"
              @sort-change="handleSortChang">
              <el-table-column type="selection" width="55" align="center" />
              <el-table-column label="编号" width="75" align="center" prop="sceneId" sortable="custom" />
              <!-- <el-table-column label="编号" width="75" prop="userId" sortable="custom" /> -->
              <el-table-column label="场景名称" align="center" prop="sceneName" :show-overflow-tooltip="true" />
              <el-table-column label="区域名称" align="center" prop="locName" :show-overflow-tooltip="true" />
              <el-table-column label="设备数量" align="center" prop="deviceNum" width="120" />

              <!--这里要传图片-->
              <el-table-column label="场景缩略图" align="center" prop="sNCode" width="330" high="400"
                :show-overflow-tooltip="true">
                <template slot-scope="scope">
                  <img style="width: 300px;height: 150px;" :src="`data:image/png;base64,${scope.row.image}`" alt="">

                </template>

              </el-table-column>


              <el-table-column label="操作" width="250" align="center" fix="right" class-name="small-padding fixed-width">
                <template slot-scope="scope">
                  <!--v-permisaction="['admin:sysUser:edit']"-->
                  <el-button size="mini" type="text" icon="el-icon-edit"
                    @click="goToDetail(scope.row, 1)">编辑</el-button>
                  <el-button size="mini" type="text" icon="el-icon-edit"
                    @click="goToDetail(scope.row, 2)">查看</el-button>


                  <el-button size="mini" type="text" icon="el-icon-s-opportunity"
                    @click="navigateToLargeScreen(scope.row)">统计</el-button>
                  <el-button v-if="scope.row.userId !== 1" v-permisaction="['admin:sysUser:remove']" size="mini"
                    type="text" icon="el-icon-edit" @click="handleUpdate(scope.row)">修改</el-button>

                  <el-button v-permisaction="['admin:sysUser:resetPassword']" size="mini" type="text"
                    icon="el-icon-delete" @click="handleDelete(scope.row)">删除</el-button>
                </template>

              </el-table-column>
            </el-table>

            <pagination v-show="total > 0" :total="total" :current-page="queryParams.pageIndex"
              :page-size="queryParams.pageSize" @pagination="handlePagination" />
          </el-col>
        </el-row>
      </el-card>

      <!-- 添加或修改参数配置对话框 -->
      <el-dialog :title="title" :visible.sync="open" width="600px" :close-on-click-modal="false" class="default-dialog">
        <el-form ref="form" :model="deviceInfo" :rules="rules" label-width="80px" @close="handleClose">
          <el-row>
            <el-col :span="12">
              <el-form-item label="设备名称" prop="deviceName">
                <el-input v-model="deviceInfo.deviceName" placeholder="请输入设备名称" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="归属区域" prop="locId">
                <treeselect v-model="deviceInfo.locId" :options="locOptions" placeholder="请选择归属区域" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="设备类型" prop="deviceType">
                <el-input v-model="deviceInfo.type.typeName" maxlength="11" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="序列号" prop="sNCode">
                <el-input v-model="deviceInfo.sNCode" placeholder="请输入序列号" maxlength="50" />
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="物理信息" prop="deviceData">
                <el-input v-model="deviceInfo.deviceData" type="textarea" />
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="更新时间" prop="updateAt">
                <el-input v-model="deviceInfo.updatedAt" maxlength="50" />
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="备注">
                <el-input v-model="deviceInfo.remark" type="textarea" placeholder="请输入内容" />
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button type="primary" @click.prevent="submitForm">确 定</el-button>
          <el-button @click="cancel">取 消</el-button>
        </div>
      </el-dialog>

      <!--新增弹出-->
      <el-dialog :title="title" :visible.sync="addSceneDialogVisible" width="600px" :close-on-click-modal="false"
        @close="restAddFrom" class="default-dialog">
        <!-- 表单内容 -->
        <!-- 表单项 -->
        <el-form ref="addqueryForm" :rules="{}" :model="addFrom" :inline="true" label-width="68px">
          <el-form-item label="场景名称" prop="sceneName">
            <el-input v-model="addFrom.sceneName" placeholder="请输入场景名称" clearable size="small" style="width: 160px" />
          </el-form-item>
          <el-form-item label="区域名称" prop="locId">
            <!-- <treeselect
            v-model="form.loclObj"
            :options="locOptions"
            placeholder="请选择归属区域"
            value-consists-of="'ALL'"
          /> -->
            <el-cascader v-model="addFrom.locId" :props="{ value: 'id', checkStrictly: true, emitPath: false }"
              @change="handleChangeCascaser" ref="cascader" :options="locOptions"></el-cascader>

            <!-- <el-input v-model="queryParams.type.typeName" placeholder="请输区域名称" clearable size="small"
              style="width: 160px" @keyup.enter.native="handleQuery" /> -->
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button type="primary" @click.prevent="submitForm">确 定</el-button>
          <el-button @click="addSceneDialogVisible = false">取 消</el-button>
        </div>
      </el-dialog>

    </template>

  </BasicLayout>
</template>

<script>
// eslint-disable-next-line no-unused-vars
import { listUser, getUser, getDevice, delUser, addUser, updateUser, updateDeviceProfile, delDevice, exportUser, resetUserPwd, changeUserStatus, importTemplate, uploadFile } from '@/api/admin/sys-user'
import { getToken } from '@/utils/auth'

import { listRole } from '@/api/admin/sys-role'
import { typetreeselect, treeselect } from '@/api/admin/sys-dept'

import Treeselect from '@riophae/vue-treeselect'
import '@riophae/vue-treeselect/dist/vue-treeselect.css'
import { listDevice, listDevices } from '@/api/sys-device'
import { lisSceneApi, addSceneApi, updateSceneApi, delSence } from '@/api/admin/sys-scene'
import { JobStart, logadd, Single } from "@/api/sys-task";



import LargeScreen from '@/views/admin/LargeScreen.vue'

export default {
  name: 'SysDeviceManage',
  components: { Treeselect },
  data() {
    return {

      operations: [
        {
          text: '修改',
          command: 'edit',
          click: (task) => this.handleDropdownCommand('edit', task)
        },
        {
          text: '删除',
          command: 'delete',
          click: (task) => this.handleDropdownCommand('delete', task)
        },

      ],
      addSceneDialogVisible: false, // 控制新增场景对话框显示
      // 是否显示弹出层
      open: false,
      deviceId: '',
      typeId: '',
      locId: '',
      jobId: '',
      currentTime: "",

      // 控制大屏幕显示
      Screen: false,

      // 定义响应式数据来存储设备信息
      deviceDataSegments: [],
      // 编号: [],
      showPassword: false, // 控制密码可见性的变量，默认为 false
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
      // 设备表格数据
      DeviceList: [],
      tempViceList: [],
      // 弹出层标题
      title: '',
      // 类型树选项
      cateOptions: undefined,
      locOptions: undefined,

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
      form: {},
      defaultProps: {
        children: 'children',
        label: 'label'
      },
      treeInputvalue: '', // tree 输入框字段
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
        token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhc2NvcGUiOiIiLCJleHAiOjQ4Njk5MjY0NDEsImlkZW50aXR5IjoxLCJuaWNlIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTcxNjI5MDQ0MSwicm9sZWlkIjoxLCJyb2xla2V5IjoiYWRtaW4iLCJyb2xlbmFtZSI6Iuezu-e7n-euoeeQhuWRmCJ9.vm6F3qdUKdq1r4hVg2AB1tNkNEgEBCYO6yaT-6IPrNw', /* 请替换成实际的token */
        pageIndex: 1,
        pageSize: 20,
        deviceName: undefined,
        type: {
          typeName: ''
        },
        loc: {
          locName: ''
        },
        sceneName: '',
        locId: '',
        locName: '',
        locIdArr: []
      },
      addFrom: {
        token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhc2NvcGUiOiIiLCJleHAiOjQ4Njk5MjY0NDEsImlkZW50aXR5IjoxLCJuaWNlIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTcxNjI5MDQ0MSwicm9sZWlkIjoxLCJyb2xla2V5IjoiYWRtaW4iLCJyb2xlbmFtZSI6Iuezu-e7n-euoeeQhuWRmCJ9.vm6F3qdUKdq1r4hVg2AB1tNkNEgEBCYO6yaT-6IPrNw', /* 请替换成实际的token */
        sceneName: '',
        locId: '',
        locName: '',
        locIdArr: [],
        sceneId: null
      },
      // 表单校验
      rules: {
        username: [{ required: true, message: '用户名称不能为空', trigger: 'blur' }],
        nickName: [{ required: true, message: '用户昵称不能为空', trigger: 'blur' }],
        locId: [{ required: true, message: '归属部门不能为空', trigger: 'blur' }],
        password: [{ required: true, message: '用户密码不能为空', trigger: 'blur' }],
        email: [
          { required: true, message: '邮箱地址不能为空', trigger: 'blur' },
          { type: 'email', message: "'请输入正确的邮箱地址", trigger: ['blur', 'change'] }
        ],
        roleId: [{ required: true, message: '用户角色不能为空', trigger: 'blur' }],
        phone: [
          { required: true, message: '手机号码不能为空', trigger: 'blur' },
          { pattern: /^1[3|4|5|6|7|8|9][0-9]\d{8}$/, message: '请输入正确的手机号码', trigger: 'blur' }
        ]
      },
      deviceInfo: {
        deviceName: '',
        locId: null,
        deviceId: '',
        type: {
          typeName: '',
          typeId: 7
        },
        sNCode: '',
        loc: {
          locName: ''
        },
        updatedAt: '',
        remark: ''
      }
    }
  },
  computed: {

  },
  watch: {
    // 根据名称筛选部门树
    selectedLocId(newVal, oldVal) {
      // 当selectedLocId变化时，重新筛选表格
      this.filterTableBySelectedLoc();
    }
  },
  created() {
    this.getList()
    this.getTreeselect()
    this.getDicts('sys_normal_disable').then(response => {
      this.statusOptions = response.data
    })
    this.getDicts('sys_user_sex').then(response => {
      this.sexOptions = response.data
    })
    this.getConfigKey('sys_user_initPassword').then(response => {
      this.initPassword = response.data.configValue
    })
  },
  mounted() {
    this.updateTime();
    this.timer = setInterval(this.updateTime, 1000);

    this.initChart()
  },
  methods: {
    goToDetail(row, state) {
      const token = getToken();
      const queryParams = `?sceneId=${encodeURIComponent(row.sceneId)}&token=${encodeURIComponent(token)}&state=${encodeURIComponent(state)}`;
      window.location.href = `http://localhost:8080/${queryParams}`;
    },

    //分页函数
    initChart() {
      // 初始化 ECharts 实例
      const barChart = echarts.init(this.$refs.barChart);
      // 设置图表配置项
      const option = {
        title: {
          text: '某设备状态统计',
          subtext: '示例数据'
        },
        tooltip: {
          trigger: 'axis'
        },
        legend: {
          data: ['在线设备', '离线设备']
        },
        xAxis: {
          type: 'category',
          data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
        },
        yAxis: {
          type: 'value'
        },
        series: [
          {
            name: '在线设备',
            type: 'bar',
            data: [120, 200, 150, 80, 70, 110, 130] // 示例数据，可以替换为你的真实数据
          },
          {
            name: '离线设备',
            type: 'bar',
            data: [100, 180, 140, 90, 60, 100, 120] // 示例数据，可以替换为你的真实数据
          }
        ]
      };

      // 使用配置项设置图表
      barChart.setOption(option)
    },
    handlePagination(pageIndex, pageSize) {
      this.queryParams.pageIndex = pageIndex;
      this.queryParams.pageSize = pageSize;
      console.log(pageIndex, pageIndex)
      this.getList();
    },

    handleAddScene() {
      // 重置表单数据为初始状态
      this.deviceInfo = {
        deviceName: '',
        locId: null,
        deviceId: '',
        type: { typeName: '', typeId: 7 },
        sNCode: '',
        loc: { locName: '' },
        updatedAt: '',
        remark: ''
      };
      // 打开弹出层
      this.addSceneDialogVisible = true;
      // 设置弹出层标题
      this.title = '添加场景';
    },


    getSecondData(url) {
      return fetch(url, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${getToken()}` // 使用Bearer Token进行认证
        }
      })
        .then(response => response.json());
    },

    // 显示第二个接口返回的数据的方法
    displaySecondData(data) {
      if (data && Array.isArray(data)) {
        // 假设您要在表格中显示这些数据
        this.tasks = data
      } else {
        console.error('Invalid data format')
      }
    },





    /** 密码点击可见 */
    togglePasswordVisibility() {
      this.showPassword = !this.showPassword // 切换 showPassword 的值
    },

    /** 时间设置 */
    formatDate(dateString) {
      // 将日期字符串分割为数组，分别代表年、月、日
      const [year, month, day] = dateString.split('-')
      // 返回格式化后的日期字符串
      return `${year}-${month}-${day}`
    },






    /** 查询列表 */
    async getList() {
      this.loading = true

      lisSceneApi(this.queryParams).then(
        response => {
          const data = response.data.list; // 获取原始设备列表数据
          const pageIndex = this.queryParams.pageIndex; // 当前页码
          const pageSize = this.queryParams.pageSize; // 每页显示数量
          const updatedList = data.map((item, index) => {
            return {
              ...item, // 保留原有数据
            };
          });
          this.DeviceList = updatedList
          this.tempViceList = [...this.DeviceList]
          this.total = response.data.count
          this.loading = false
        }
      )
    },
    /** 查询部门下拉树结构 */
    getDeviceTreeselect() {
      typetreeselect().then(response => {
        this.typeOptions = response.data
      })
    },
    getTreeselect() {
      treeselect().then(response => {
        this.locOptions = response.data
      })
      console.log(this.locOptions)
    },

    //区域树筛选表格
    filterTableBySelectedLoc() {
      // 根LocId筛选DeviceList
      this.tempViceList = this.DeviceList.filter(item => {
        return item.locId === this.locId;
      });
      // 如果需要，更新总条数total
      this.total = this.tempViceList.length;
      console.log(this.locId)
    },
    // 筛选节点
    filterNode(value, data) {
      if (!value) return true
      return data.label.indexOf(value) !== -1
    },
    // 节点单击事件
    handleNodeClick(data) {
      this.locId = data.id
      console.log(this.locId)
      this.filterTableBySelectedLoc();

    },
    /** 转换菜单数据结构 */
    normalizer(node) {
      if (node.children && !node.children.length) {
        delete node.children
      }
      return {
        id: node.id,
        label: node.label,
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
      }).then(function () {
        return changeUserStatus(row)
      }).then(() => {
        this.msgSuccess(text + '成功')
      }).catch(function () {
        row.status = row.status === '2' ? '1' : '2'
      })
    },
    // 取消按钮
    cancel() {
      this.open = false
      this.resetDeviceForm()
    },
    /** 搜索按钮操作 */
    handleQuery() {
      console.log('search')
      this.queryParams.pageIndex = 1
      this.getList()
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.$refs.queryForm.resetFields()
      this.queryParams = {
        pageIndex: 1,
        pageSize: 10,
        deviceName: undefined,
        type: {
          typeName: ''
        },
        loc: {
          locName: ''
        }
      }
      this.handleQuery()
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.deviceId)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      console.log('Update triggered', row)
      const { sceneId, locId, locName, sceneName } = row

      this.addFrom.sceneId = sceneId
      this.addFrom.locId = locId
      this.addFrom.locName = locName
      this.addFrom.sceneName = sceneName

      this.title = '修改场景'
      this.addSceneDialogVisible = true
    },
    handleChangeCascaser(value) {
      const that = this
      const list = that.$refs['cascader'] ? Array.from(that.$refs['cascader'].getCheckedNodes()) : []
      if (list) {
        // that.$refs['cascader'] 获取cascader 组件实例， getCheckedNodes 方法获取当前选中的node 节点
        console.log('list', list)
        const obj = list.length > 0 ? list[0] : { value: '', label: '' };
        const { value, label } = obj
        that.addFrom.locId = value
        that.addFrom.locName = label
        console.log('addFrom11', this.addFrom)
      } else {
        that.addFrom.locName = ''
      }
    },
    restAddFrom() {
      this.addFrom.sceneId = ''
      this.addFrom.locId = ''
      this.addFrom.locName = ''
      this.addFrom.sceneName = ''
      this.addFrom.locIdArr = []
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs['addqueryForm'].validate(valid => {
        if (valid) {
          const { locId, locName, sceneName, sceneId, token } = this.addFrom;
          const params = { locId, locName, sceneName, sceneId };
          let api = this.title.indexOf('修改') > -1 ? updateSceneApi : addSceneApi
          api(params, token).then(response => {
            if (response.code === 200) {
              this.msgSuccess(response.msg);
              // 关闭弹窗
              this.addSceneDialogVisible = false; // 确保使用了正确的属性名
              this.open = false; // 如果 this.open 也用于控制弹窗显示，则也需要设置为 false
              this.getList(); // 刷新列表
              // this.restAddFrom(); // 重置表单数据
            } else {
              this.msgError(response.msg);
            }
          });
        }
      });
    },

    /** 删除按钮操作 */
    handleDelete(row) {
      const Ids = (row.sceneId && [row.sceneId]) || this.ids
      console.log(row.sceneId)
      this.$confirm('是否确认删除设备编号为"' + Ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        // 调用删除操作的API
        delSence({ 'ids': Ids }).then(response => {
          if (response.code === 200) { // 假设响应码200表示操作成功
            // 显示成功消息
            this.$message.success('删除成功');
            // 重新获取数据列表
            this.getList();
          } else {
            // 显示失败消息
            this.$message.error('删除失败');
          }
        }).catch(error => {
          // 处理错误情况
          console.error('删除操作发生错误：', error);
          this.$message.error('删除操作发生错误');
        });
      }).catch(() => {
        // 用户取消删除操作
        this.$message.info('已取消删除');
      });
    },

    /** 导出按钮操作 */
    handleExport() {
      const queryParams = this.queryParams
      this.$confirm('是否确认导出所有用户数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function () {
        return exportUser(queryParams)
      }).then(response => {
        this.download(response.msg)
      }).catch(function () { })
    },
    /** 导入按钮操作 */
    handleImport() {
      this.upload.title = '用户导入'
      this.upload.open = true
    },
    /** 下载模板操作 */
    importTemplate() {
      importTemplate().then(response => {
        this.download(response.msg)
      })
    },
    // 文件上传中处理
    handleFileUploadProgress(event, file, fileList) {
      this.upload.isUploading = true
    },
    // 文件上传成功处理
    handleFileSuccess(response, file, fileList) {
      this.upload.open = false
      this.upload.isUploading = false
      this.$refs.upload.clearFiles()
      this.$alert(response.msg, '导入结果', { dangerouslyUseHTMLString: true })
      this.getList()
    },
    // 提交上传文件
    submitFileForm() {
      this.$refs.upload.submit()
    },
    handleClose() {
      this.deviceInfo = this.$options.data().deviceInfo
    },
    reset() {
      this.form = {
        executorParam: undefined,
        ids: undefined
      }
      this.resetForm('form')
    },
    navigateToLargeScreen(data) {
      // 使用编程式导航跳转到 LargeScreen 组件
      console.log(data)
      this.$router.push({
        name: 'LargeScreen',
        query: {
          chartData: JSON.stringify(data),
        }
      })
    }
  },


}
</script>

<style>
@import '~@fortawesome/fontawesome-free/css/all.css';

.el-dropdown-link {
  cursor: pointer;
  /* 鼠标悬停时显示手形图标 */
}

.el-dropdown-link:hover {
  background-color: #f5f5f5;
  /* 鼠标悬停时的背景色 */
  color: #409eff;
  /* 鼠标悬停时的文字颜色 */
}


.search-box {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
  margin-top: 10px;
  margin-left: 23px;
}

.search-label {
  margin-right: 8px;
  /* 为标签和输入框之间添加一些间距 */
  white-space: nowrap;
  /* 防止文本换行 */
  margin-top: -5px
}

.search-input {
  height: 35px;
  width: 180px;
  border-radius: 4px;
  border-color: rgba(128, 128, 128, 0.114);
  font-size: smaller;
  font-weight: normal;
  padding-left: 10px;
  margin-right: 5px;
}


.form-group {
  display: flex;
  padding-left: 20px;
}

.form-group label {
  padding-left: 75px;
  margin-right: 25px;
}

.modal-backdrop {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 1000;
  /* 确保遮罩层在其他元素之上 */
  display: flex;
  align-items: center;
  justify-content: center;
}

@keyframes slide-down {
  from {
    transform: translateY(-100%);
  }

  to {
    transform: translateY(0);
  }
}

.modal {
  width: 850px;
  height: 630px;
  border-radius: 4px;
  /* 模态框的宽度，根据需要调整 */
  background: #fff;
  z-index: 1001;
  /* 确保模态框在遮罩层之上 */
  animation: slide-down 0.3s;
  /* 下滑动画 */
}


/* 模态框的样式 */
.modal-content {
  border-radius: 4px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.5);
  /* 添加阴影效果 */
  width: 850px;
  height: 100%;
}

/* 模态框头部的样式 */
.modal-header {
  border-radius: 4px;
  background-color: #fff;
  /* 设置背景颜色 */
  border-bottom: none;
  /* 移除边框 */
  padding: 5px 20px;
  /* 设置内边距 */
}

.modal p {
  padding-left: 20px;
}

/* 模态框标题的样式 */
.modal-title {
  font-size: 1.25rem;
  /* 设置字体大小 */
  color: #333;
  /* 设置字体颜色 */
}

/* 表单分组的样式 */
.form-group {
  margin-bottom: 15px;
  /* 设置外边距 */
}

/* 表单标签的样式 */
.control-label {
  font-weight: bold;
  /* 设置字体为粗体 */
  color: #666;
  /* 设置字体颜色 */
  margin-top: 2px;
}

/* 输入框和文本域的样式 */
.form-control {
  height: 35px;
  width: 202px;
  border-radius: 4px;
  /* box-shadow: none; */
  border-color: rgba(128, 128, 128, 0.114);
  margin-top: -5px;
  font-size: smaller;
  font-weight: normal;
  padding-left: 10px;
}

/* 按钮的样式 */
.btn {
  border-radius: 4px;
  /* 移除边框圆角 */
  padding: 10px 20px;
  /* 设置内边距 */
  font-size: 0.875rem;
  /* 设置字体大小 */
  margin-left: 220px;
  margin-top: 5px;
  cursor: pointer;
}

.btn-primary {
  background-color: #0099FF;
  /* 设置主要按钮的背景颜色 */
  border-color: #0099FF;
  /* 设置主要按钮的边框颜色 */
  color: #fff;
}

.btn-default {
  background-color: #fff;
  /* 设置默认按钮的背景颜色 */
  color: #333;
  /* 设置默认按钮的字体颜色 */
  border-color: #ddd;
  /* 设置默认按钮的边框颜色 */
}

/* 分割线的样式 */
hr {
  margin-top: 20px;
  /* 设置上边距 */
  margin-bottom: 20px;
  /* 设置下边距 */
}

/* 用于表示必填字段的红色星号 */
.required-field::before {
  content: '*';
  color: red;
  font-size: 1.2em;
  margin-right: 4px;
}

/* 自定义滚动条样式 */
::-webkit-scrollbar {
  width: 8px;
  background-color: #f5f5f5;
}

::-webkit-scrollbar-thumb {
  background-color: #c1c1c1;
  border-radius: 4px;
}

::-webkit-scrollbar-thumb:hover {
  background-color: #a8a8a8;
}

::-webkit-scrollbar-track {
  background-color: transparent;
}

::-webkit-scrollbar-track:hover {
  background-color: #e0e0e0;
}



.back-wrapper {
  display: flex;
  align-items: center;
  cursor: pointer;
  font-size: 16px;
  width: 85px;
  border-radius: 5px;
  padding: 10px 15px;
  margin-top: 10px;
  margin-left: 5px;
  margin-bottom: -10px;
  transition: background-color 0.3s;
}


.back-wrapper i {
  margin-right: 5px;
  /* 为图标和文本之间添加一些间距 */
  color: black;
}

/* 选项卡未选中时的样式 */
.tab-button {
  color: black;
  /* 默认文字颜色为黑色 */
  border: none;
  /* 移除边框 */
  background-color: transparent;
  /* 透明背景 */
  padding: 5px 15px;
  cursor: pointer;
  border-bottom: 2px solid transparent;
  /* 默认不显示下划线 */
  transition: border-bottom-color 0.3s;
  /* 添加过渡效果 */
  margin-left: 24px;
  margin-top: 10px;
  font-size: medium;
}

/* 选项卡选中时的样式 */
.tab-button.active {
  color: #0099FF;
  /* 选中时的文字颜色 */
  border-bottom-color: #0099FF;
  /* 选中时显示下划线 */
}

.content {
  padding: 20px;
  padding-top: 0px;
  /* border: 1px solid #ddd; */
  border-top: none;
  border-bottom-left-radius: 5px;
  border-bottom-right-radius: 5px;
  overflow-y: auto;
  /* 如果内容过长，提供滚动条 */
  max-height: 100vh;
  /* 限制内容的最大高度，防止溢出 */
  position: absolute;
  top: 75px;
  right: 0;
  bottom: 0;
  left: 0;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  /* 轻微的阴影效果 */
  transition: all 0.3s ease;
  /* 过渡动画效果 */
  margin-top: 8px;
}

#data {
  height: 100%;
}


.device-image img {
  width: 30%;
  height: auto;
}

.striped-table {
  width: 100%;
  /* 表格宽度 */
  border-collapse: collapse;
  /* 合并单元格边框 */
  margin-top: 10px;
  /* 表格上方的外边距 */
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  /* 字体 */
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  /* 表格阴影 */
}

.striped-table thead {
  background-color: rgba(128, 128, 128, 0.103);
  color: black;
  /* 表头文字颜色 */
  text-transform: uppercase;
  /* 表头文字小写转大写 */
}

.striped-table th,
.striped-table td {
  padding: 10px;
  /* 单元格内边距 */
  text-align: center;
  /* 文字左对齐 */
  border-bottom: 1px solid #ddd;
  /* 单元格底部边框 */
  vertical-align: middle;
  /* 垂直居中对齐 */
}

.striped-table tbody tr:nth-child(odd) {
  background-color: #ffffffc0;
  /* 奇数行背景色 */
}

.striped-table tbody tr:nth-child(even) {
  background-color: rgba(128, 128, 128, 0.103);
  /* 偶数行背景色 */
}

.button-group {
  margin-top: 20px;
  margin-left: -5px;
}

.button-group button {
  padding: 5px 15px;
  /* 按钮的内边距 */
  margin: 0 5px;
  /* 按钮之间的外边距 */
  border: none;
  /* 移除按钮默认边框 */
  border-radius: 5px;
  /* 按钮圆角 */
  background-color: #0099FF;
  /* 按钮背景颜色 */
  color: white;
  /* 按钮文字颜色 */
  cursor: pointer;
  /* 鼠标悬停时变为手指形状 */
  transition: background-color 0.3s;
  /* 背景颜色变化的过渡效果 */
}

.button-group button:hover {
  background-color: #0077CC;
  /* 鼠标悬停时的背景颜色 */
}

.button-group button:active {
  background-color: #005FA3;
  /* 按钮被点击时的背景颜色 */
}


.card {
  border: 1px solid #ddd;
  border-radius: 12px;
  padding: 20px;
  /* 增加内边距 */
  margin-bottom: 10px;
  /* 增加下外边距 */
  width: 285px;
  height: 150px;
  display: inline-block;
  margin-left: 30px;
  margin-top: 40px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  /* 添加阴影效果 */
  transition: box-shadow 0.3s ease;
  /* 添加阴影过渡效果 */
  background-color: rgba(255, 255, 255, 0.7);
  /* 设置背景颜色 */
}

.card:hover {
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
  /* 鼠标悬停时增加阴影深度 */
}

.card div {
  margin-top: 15px;
}

.device-img {
  width: 300px;

  img {
    width: 100%;
  }
}

.upload-demo {
  text-align: center
}
</style>
