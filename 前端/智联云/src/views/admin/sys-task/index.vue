<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-row :gutter="20">
          <!--部门数据-->
          <el-col :span="4" :xs="24">
            <div class="head-container">
              <el-input
                v-model="locName"
                placeholder="请输入区域名称"
                clearable
                size="small"
                prefix-icon="el-icon-search"
                style="margin-bottom: 20px"
              />
            </div>
            <div class="head-container">
              <el-tree
                ref="tree"
                :data="locOptions"
                :props="defaultProps"
                :expand-on-click-node="false"
                :filter-node-method="filterNode"
                default-expand-all
                @node-click="handleNodeClick"
              />
            </div>
          </el-col>
          <!--用户数据-->
          <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
            <el-form-item label="任务描述" prop="jobDesc">
              <el-input
                v-model="queryParams.jobDesc"
                placeholder="请输入任务描述"
                clearable
                size="small"
                style=" width: 160px"
                @keyup.enter.native="handleQuery"
              />
            </el-form-item>
            <el-form-item label="负责人" prop="author">
              <el-input
                v-model="queryParams.author"
                placeholder="请输入负责人"
                clearable
                size="small"
                style="width: 160px"
                @keyup.enter.native="handleQuery"
              />
            </el-form-item>
            <el-form-item label="状态" prop="triggerStatus">
              <el-select
                v-model="queryParams.triggerStatus"
                placeholder="状态"
                clearable
                size="small"
                style="width: 160px"
              >
                <el-option label="全部" value="-1" />
                <el-option label="STOP" value="0" />
                <el-option label="RUNNING" value="1" />
              </el-select>
            </el-form-item>

            <el-form-item>
              <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
              <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
            </el-form-item>
          </el-form>

          <el-col :span="20" :xs="24">
            <el-row :gutter="10" class="mb8">
              <el-col :span="1.5">
                <el-button
                  v-permisaction="['admin:sysTask:edit']"
                  type="success"
                  icon="el-icon-edit"
                  size="mini"
                  :disabled="single"
                  @click=""
                >修改</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  v-permisaction="['admin:sysTask:remove']"
                  type="danger"
                  icon="el-icon-delete"
                  size="mini"
                  :disabled="multiple"
                  @click="DeleteMulJob"
                >删除</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  v-permisaction="['admin:sysTask:start']"
                  class="Custom-Button"
                  type="primary"
                  icon="el-icon-circle-check"
                  size="mini"
                  :disabled="multiple"
                  @click="TaskMulstart"
                >批量启动</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  v-permisaction="['admin:sysTask:start']"
                  class="Custom-Button-Stop"
                  type="primary"
                  icon="el-icon-circle-check"
                  size="mini"
                  :disabled="multiple"
                  @click="TaskMulstop"
                >批量停止</el-button>
              </el-col>
            </el-row>

            <el-table
              v-loading="loading"
              :data="indexedUserList"
              border
              @selection-change="handleSelectionChange"
            >
              <el-table-column type="selection" width="55" align="center" />
              <el-table-column label="编号" width="75" prop="ID" align="center" />
              <!-- <el-table-column label="编号" width="75" prop="userId" sortable="custom" /> -->
              <el-table-column label="负责人" width="150" prop="author" align="center" :show-overflow-tooltip="true" />
              <el-table-column label="任务描述" prop="jobDesc" align="center" :show-overflow-tooltip="true" />
              <el-table-column label="设备名称" prop="deviceName" align="center" :show-overflow-tooltip="true" width="250" />
              <el-table-column label="状态" width="130" align="center">
                <template slot-scope="scope">
                  <el-tag :type="scope.row.triggerStatus === 0 ? 'lightcoral' : 'success'" effect="dark" style="width: 80px;">
                    {{ scope.row.triggerStatus === 0 ? 'STOP' : 'RUNNING' }}
                  </el-tag>
                </template>
              </el-table-column>
              <el-table-column label="设备地点" prop="locName" align="center" width="200" />
              <el-table-column
                label="操作"
                width="200"
                align="center"
                fix="right"
                class-name="small-padding fixed-width"
              >
                <template slot-scope="scope">
                  <el-dropdown trigger="click" size="large" @command="handleCommand">
                    <span class="el-dropdown-link">
                      操作<i class="el-icon-arrow-down el-icon--right" />
                    </span>
                    <el-dropdown-menu slot="default">
                      <el-dropdown-item command="start" :command-data="scope.row">启动任务</el-dropdown-item>
                      <el-dropdown-item command="SingleStart" :command-data="scope.row">单次执行</el-dropdown-item>
                      <el-dropdown-item command="Stop" :command-data="scope.row">停止任务</el-dropdown-item>
                      <el-dropdown-item command="remove" :command-data="scope.row">删除任务</el-dropdown-item>
                      <el-dropdown-item command="viewLog" :command-data="scope.row">查询日志</el-dropdown-item>
                      <el-dropdown-item command="nextExecutionTime" :command-data="scope.row">下次执行时间</el-dropdown-item>
                    </el-dropdown-menu>
                  </el-dropdown>
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

      <!-- 单次执行对话框：-->
      <el-dialog title="执行一次" :visible.sync="SingledialogVisible" width="40%" :close-on-click-modal="false">
        <el-form ref="form" label-width="100px" :model="form" :rules="rules">
          <el-form-item label="任务参数" prop="executorParam">
            <!-- 使用 v-model 将输入框的值绑定到任务参数的属性上 -->
            <el-input v-model="form.executorParam" type="textarea" :rows="4" placeholder="请输入任务参数" />
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
          <el-button @click="SingledialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="TaskConfirm">确 定</el-button>
        </span>
      </el-dialog>

      <!--操作日志对话框：-->
      <el-dialog
        title="操作日志"
        :visible.sync="LogsdialogVisible"
        width="60%"
      >
        <el-table
          v-loading="loading"
          :data="LogsListSort"
          style="width: 100%"
          stripe
          border
          :fit="true"
        >
          <el-table-column prop="ID" label="ID" width="50" align="center" />
          <el-table-column prop="name" label="姓名" width="150" align="center" />
          <el-table-column prop="tel" label="电话" width="150" align="center" />
          <el-table-column prop="operator" label="任务操作" width="300" align="center" />
          <el-table-column prop="createdAt" label="时间" width="255" align="center" />
          <!-- <el-table-column fix="right" label="操作" align="center">
            <template slot-scope="scope">
              <el-button
                v-if="scope.row.userId !== 1"
                v-permisaction="['admin:sysUser:remove']"
                size="mini"
                type="text"
                icon="el-icon-delete"
                @click="LogsDelete(scope.row)"
              >删除</el-button>
            </template></el-table-column> -->
        </el-table>
        <pagination
            v-show="logsqueryParams.count>0"
            :total="logsqueryParams.count"
            :page.sync="logsqueryParams.pageIndex"
            :limit.sync="logsqueryParams.pageSize"
            @pagination="GetLogsList"
        />
        <span slot="footer" class="dialog-footer">
          <el-button @click="LogsdialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="LogsConfirm">确 定</el-button>
        </span>
      </el-dialog>
      <!--下次执行时间对话框-->
      <el-dialog
        title="下次执行时间"
        :visible.sync="TimeModal"
        width="30%"

      >
        <div v-for="(time, index) in nextExecutionTimes" :key="index" align="center" class="large-text">
          {{ time }}
        </div>
      </el-dialog>
    </template>
  </BasicLayout>
</template>

<script>
import { getToken } from '@/utils/auth'
import { listRole } from '@/api/admin/sys-role'
import { treeselect } from '@/api/admin/sys-dept'

import Treeselect from '@riophae/vue-treeselect'
import '@riophae/vue-treeselect/dist/vue-treeselect.css'
import {
  Findtime,
  GetLogs, JobfirstRemove,
  JobsecondRemove,
  JobStart, JobStop,
  listfirstTask,
  listsecondTask, logadd,
  logremove,
  Single,
  TreeChild
} from '@/api/sys-task'
import row from 'element-ui/packages/row'

export default {
  name: 'SysUserManage',
  components: { Treeselect },
  data() {
    return {
      编号: [],
      showPassword: false, // 控制密码可见性的变量，默认为 false
      // 任务启动成功返回信息：
      Massage: undefined,
      // 遮罩层
      loading: true,
      // 单次执行弹框
      SingledialogVisible: false,
      // 操作日志弹框
      LogsdialogVisible: false,
      // 操作日志列表
      LogsList: [],
      // 单次操作执行参数
      executorParam: undefined,
      // 控制模态框显示的变量
      showAddModal: false,
      // 控制模态框显示的变量
      showEditModal: false,
      // 下次执行时间弹窗
      TimeModal: false,
      // 存储下次执行时间数据
      nextExecutionTimes: [],
      // 封装后的id数组
      Ids: [],
      // 接口第一次getlist查询返回数据存储处
      tasklist: [],
      // 接口第二次Getlist查询返回数据存储处
      TaskList: [],
      // 日志ID
      oplId: undefined,
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 分页时使用的总条数
      total: 0,
      // 用户表格数据
      userList: null,
      // 弹出层标题
      title: '',
      // 部门树选项
      locOptions: undefined,
      // 是否显示弹出层
      open: false,
      // 部门名称
      locName: undefined,
      // 日期范围
      dateRange: [],
      // 状态数据字典
      statusOptions: [],
      // 性别状态字典
      sexOptions: [],
      // 表单参数
      form: {},
      defaultProps: {
        children: 'children',
        label: 'label'
      },

      // 封装查询参数
      queryParams: {
        page: 1,
        pageIndex: 1,
        pageSize: 10,
        start: 0,
        length: 10,
        deviceId: undefined,
        triggerStatus: '全部',
        JobId: undefined,
        typeId: undefined,
        locId: undefined,
        jobDesc: undefined,
        author: undefined,
        ids: undefined
      },

      // 封装logs查询参数
      logsqueryParams: {
        jobId: undefined,
        pageIndex: 1,
        pageSize: 10,
        count: 0
      },

      // 表单校验
      rules: {
        executorParam: [{ required: true, message: '任务参数不能为空', trigger: 'blur' }]
      }
    }
  },
  computed: {
    // 计算编号的计算属性
    indexedUserList() {
      // 根据列表数据和当前分页的索引计算编号
      const { TaskList, queryParams: { pageIndex, pageSize }} = this
      console.log(TaskList)
      return TaskList.map((item, index) => {
        return {
          ...item, // 保留原有数据
          ID: (pageIndex - 1) * pageSize + index + 1 // 计算编号
        }
      })
    },
    LogsListSort() {

      const { LogsList, logsqueryParams: { pageIndex, pageSize }} = this
      console.log(LogsList)
      if (!Array.isArray(LogsList)) {
        console.error("LogsList is not an array", LogsList);
        return [];
      }
      return LogsList.map((item, index) => {
        return {
          ...item, // 保留原有数据
          ID: (pageIndex - 1) * pageSize + index + 1 // 计算编号
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
    console.log(this.queryParams.triggerStatus)
    this.getList()
    this.GetLogsList()
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
  methods: {
    GetLogsList() {
      GetLogs(this.logsqueryParams).then(response => {
        this.LogsList = this.TimeTranslators(response.data.list)
        console.log(response.data)
      })
    },
    /** 查询用户列表 */
    getList() {
      this.loading = true
      listfirstTask(this.queryParams).then(response => {
        this.tasklist = response.data
        this.total = response.data.count
        console.log(this.tasklist)
        this.loading = false
        const JobIdArray = this.tasklist.list.map(item => item.jobId)

        // 将 JobId 数组转换为以逗号分隔的字符串
        this.queryParams.ids = JobIdArray.join(',')
        console.log(this.queryParams.ids)
        if (this.queryParams.ids) {
          this.queryParams.length = this.queryParams.pageSize
          if (this.queryParams.triggerStatus === '全部') {
            this.queryParams.triggerStatus = -1
          } else if (this.queryParams.triggerStatus === '') {
            this.queryParams.triggerStatus = -1
          } else if (this.queryParams.triggerStatus === 'RUNNING') {
            this.queryParams.triggerStatus = '1'
          } else if (this.queryParams.triggerStatus === 'STOP') {
            this.queryParams.triggerStatus = '0'
          }
          this.GetList()
        } else {
          this.TaskList = []
        }
      })
    },

    GetList() {
      this.loading = true
      console.log(this.queryParams)
      listsecondTask(this.queryParams).then(response => {
        this.TaskList = response.content.data.map(item =>{
          const task = this.tasklist.list.find(task => task.jobId === item.id)
          const locName = task ? task.loc.locName : ''
          const deviceName = task && task.device !== null ? task.device.deviceName : ''

          return {
            ...item,
            locName: locName,
            deviceName: deviceName
          }
        })
        this.loading = false
        console.log(this.TaskList)
      })
      if (this.queryParams.triggerStatus === -1) {
        this.queryParams.triggerStatus = '全部'
        // eslint-disable-next-line no-empty
      } else if (this.queryParams.triggerStatus === 1) {
        this.queryParams.triggerStatus = 'RUNNING'
      } else if (this.queryParams.triggerStatus === 0) {
        this.queryParams.triggerStatus = 'STOP'
      }
    },

    /** 查询部门下拉树结构 */
    getTreeselect() {
      treeselect().then(response => {
        this.locOptions = response.data
      })
    },
    // 筛选节点
    filterNode(value, data) {
      if (!value) return true
      return data.label.indexOf(value) !== -1
    },
    // 节点单击事件
    handleNodeClick(data) {
      TreeChild({ locId: data.id }).then(response => {
        const dataArray = response.data
        console.log(dataArray)
        if(dataArray != null){
          const idsString = dataArray.join(',')
          this.queryParams.ids = idsString
          console.log(this.queryParams)

          this.queryParams.jobDesc = ''
          this.queryParams.author = ''
          this.queryParams.locId = ''
          if (this.queryParams.triggerStatus === '全部') {
            this.queryParams.triggerStatus = -1
          } else if (this.queryParams.triggerStatus === '') {
            this.queryParams.triggerStatus = -1
          } else if (this.queryParams.triggerStatus === 'RUNNING') {
            this.queryParams.triggerStatus = '1'
          } else if (this.queryParams.triggerStatus === 'STOP') {
            this.queryParams.triggerStatus = '0'
          }
          this.GetList()
        }else {
          this.queryParams.ids = ''
          this.TaskList = []
        }
      })
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
    // 取消按钮
    cancel() {
      this.open = false
      this.reset()
    },
    // 表单重置
    reset() {
      this.form = {
        executorParam: undefined,
        ids: undefined
      }
      this.resetForm('form')
    },
    // queryParams重置
    ResetqueryParams() {
      this.queryParams = {
        page: 1,
        pageIndex: 1,
        pageSize: 10,
        start: 0,
        length: 10,
        deviceId: undefined,
        triggerStatus: '全部',
        JobId: undefined,
        typeId: undefined,
        locId: undefined,
        jobDesc: undefined,
        author: undefined,
        ids: undefined
      }
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.page = 1
      console.log(this.queryParams.author)
      console.log(this.queryParams)
      console.log('handleQuery函数启动正常')
      this.getList()
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.queryParams.jobDesc = ''
      this.queryParams.author = ''
      this.queryParams.locId = ''
      this.queryParams.triggerStatus = '全部'
      this.handleQuery()
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      console.log('handleSelectionChange函数执行成功')
      console.log(selection)
      this.Ids = selection.map(item => item.id)
      this.queryParams.ids = this.Ids.join(',')
      this.single = selection.length !== 1
      console.log(this.single)
      this.multiple = !selection.length
      console.log(this.single)
    },
    /** 执行一次提交按钮 */
    TaskConfirm: function() {
      console.log(this.form)
      this.$refs['form'].validate(valid => {
        if (valid) {
          Single(this.form).then(response => {
            this.getList()
            this.SingledialogVisible = false
          })
        }
      })
    },
    /** 操作日志提交按钮 */
    LogsConfirm: function() {
      this.LogsdialogVisible = false
    },
    /** 删除按钮操作 */
    TaskDelete(row) {
      const Ids = (row.userId && [row.userId]) || this.Ids
      this.$confirm('是否确认删除用户编号为"' + Ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return delUser({ 'Ids': Ids })
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
    handleCommand(command, row) {
      console.log('handleCommand函数执行正常')
      console.log(row)
      console.log(command)
      switch (command) {
        case 'start':
          this.taskstart(row)
          break
        case 'SingleStart':
          this.SingleTask(row)
          break
        case 'Stop':
          this.taskstop(row)
          break
        case 'remove':
          this.DeleteJob(row)
          break
        case 'viewLog':
          this.ViewLog(row)
          break
        case 'nextExecutionTime':
          this.nextExecutionTime(row)
          break
        default:
          break
      }
    },
    /* 单个任务启动 */
    taskstart(row) {
      console.log(row.$attrs['command-data'])
      console.log(row.$attrs['command-data'].id)
      this.queryParams.ids = row.$attrs['command-data'].id
      JobStart(this.queryParams).then(response => {
        this.LogsAdd(this.queryParams)
        this.getList()
      })
    },
    /* 多个任务启动 */
    TaskMulstart() {
      JobStart(this.queryParams).then(response => {
        this.LogsAdd(this.queryParams)
        this.getList()
      })
    },
    /* 执行一次任务 */
    SingleTask(row) {
      this.reset()
      this.form.ids = String(row.$attrs['command-data'].id).trim()
      this.SingledialogVisible = true
    },
    // 单个任务停止
    taskstop(row) {
      console.log(row.$attrs['command-data'])
      console.log(row.$attrs['command-data'].id)
      this.queryParams.ids = row.$attrs['command-data'].id
      JobStop(this.queryParams).then(response => {
        this.LogsStop(this.queryParams)
        this.getList()
      })
    },
    /* 多个任务终止 */
    TaskMulstop() {
      JobStop(this.queryParams).then(response => {
        this.LogsStop(this.queryParams)
        this.getList()
      })
    },
    /* 操作日志 */
    ViewLog(row) {
      this.LogsdialogVisible = true
      this.logsqueryParams.jobId = row.$attrs['command-data'].id
      GetLogs(this.logsqueryParams).then(response => {
        this.LogsList = this.TimeTranslators(response.data.list)
        this.logsqueryParams.count = response.data.count
        console.log(this.LogsList)
        console.log(this.logsqueryParams.count)
      })
    },
    // 删除单行任务
    DeleteJob(row) {
      this.queryParams.ids = row.$attrs['command-data'].id
      console.log(row.$attrs['command-data'])
      const self = this // 存储当前组件实例的引用
      console.log(self.queryParams.ids)
      this.$confirm('是否删除编号为' + row.$attrs['command-data'].ID + '的任务', '警告', {
        cancelButtonText: '取消',
        confirmButtonText: '确定',
        type: 'warning'
      }).then(async function() {
        console.log(self.queryParams) // 使用存储的this指向当前组件实例
        await JobsecondRemove(self.queryParams)
        // 执行第二个异步操作
        return JobfirstRemove({ ids: [self.queryParams.ids] })
      }).then(response => {
        console.log(self.queryParams.ids)
        self.ResetqueryParams()
        self.getList()
      })
    },
    // 删除多行任务
    async DeleteMulJob() {
      console.log(this.queryParams)
      await JobsecondRemove(this.queryParams)
      const arr = this.queryParams.ids.split(',')
      console.log(arr)
      console.log(this.queryParams)
      for (let i = 0; i < arr.length; i++) {
        const id = parseInt(arr[i], 10)
        await JobfirstRemove({ ids: [id] }).then(response => {
          // 在异步操作的回调函数内执行这两个函数
          this.ResetqueryParams()
          this.getList()
        });
      }
      // 在这里执行第二个异步操作
    },

    // 删除日志
    LogsDelete(row) {
      console.log(row)
      const temp = { ids: [row.oplId] }
      logremove(temp).then(response => {
        this.GetLogsList()
      })
    },
    // 增加日志RUN
    async LogsAdd(row) {
      console.log(row.ids)
      if(typeof row.ids === 'string'){
        const jobIdArray = row.ids.split(',')
        const jobIdCount = jobIdArray.length
        try {
          for (let i = 0; i < jobIdCount; i++) {
            const jobId = parseInt(jobIdArray[i])
            const temp = { jobId: jobId, name: 'admin', operator: '设备' + jobIdArray[i] + '启动', tel: '123456' }
            await logadd(temp); // 等待日志添加操作完成
          }
          // 所有日志添加操作完成后，刷新日志列表
          await this.GetLogsList();
        } catch (error) {
          console.error("Error adding logs:", error);
        }
      }else {
        const temp = { jobId: row.ids, name: 'admin', operator: '设备' + row.ids + '启动', tel: '123456' }
        await logadd(temp); // 等待日志添加操作完成
      }

    },

    // 增加日志STOP
    async LogsStop(row) {
      console.log(row.ids)
      if(typeof row.ids === 'string'){
        const jobIdArray = row.ids.split(',')
        const jobIdCount = jobIdArray.length
        try {
          for (let i = 0; i < jobIdCount; i++) {
            const jobId = parseInt(jobIdArray[i])
            const temp = { jobId: jobId, name: 'admin', operator: '设备' + jobIdArray[i] + '停止', tel: '123456' }
            await logadd(temp); // 等待日志添加操作完成
          }
          // 所有日志添加操作完成后，刷新日志列表
          await this.GetLogsList();
        } catch (error) {
          console.error("Error adding logs:", error);
        }
      }else {
        const temp = { jobId: row.ids, name: 'admin', operator: '设备' + row.ids + '停止', tel: '123456' }
        await logadd(temp); // 等待日志添加操作完成
      }
    },
    // 下次执行时间
    nextExecutionTime(row) {
      console.log(row)
      Findtime({
        scheduleType: row.$attrs['command-data'].scheduleType,
        scheduleConf: row.$attrs['command-data'].scheduleConf
      }).then(response => {
        this.nextExecutionTimes = response.content
        this.TimeModal = true
      })
    },

    //时间字符串转化
    TimeTranslators(logList) {
      console.log(logList)
  return logList.map(log => {
    console.log(log)
    const originalCreatedAt = new Date(log.createdAt); // 将ISO 8601时间字符串转换为Date对象
    const year = originalCreatedAt.getFullYear();
    const month = String(originalCreatedAt.getMonth() + 1).padStart(2, '0');
    const day = String(originalCreatedAt.getDate()).padStart(2, '0');
    const hours = String(originalCreatedAt.getHours()).padStart(2, '0');
    const minutes = String(originalCreatedAt.getMinutes()).padStart(2, '0');
    const seconds = String(originalCreatedAt.getSeconds()).padStart(2, '0');
    const formattedCreatedAt = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`; // 格式化日期时间
    console.log(formattedCreatedAt)
    return {
      ...log,
      createdAt: formattedCreatedAt
    }
  })
}


}
}

</script>
@/api/admin/sys-loc@/api/admin/sys-dept

<style>
.Custom-Button {
  background-color: #e6f7ff; /* 浅蓝色背景 */
  border-color: #91d5ff; /* 浅蓝色边框 */
  color: #409eff; /* 文本颜色 */
}
.Custom-Button-Stop {
  background-color: #ffe6f9; /* 浅蓝色背景 */
  border-color: #ff91e7; /* 浅蓝色边框 */
  color: #ff4073; /* 文本颜色 */
}

.large-text {
  font-size: 25px;
}

</style>
