<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-row :gutter="20">
          <!-- 树形菜单 -->
          <el-col :span="4" :xs="24">
            <div class="head-container">
              <el-input v-model="treeInputvalue"
              placeholder="请输入区域名称"
              clearable size="small"
                prefix-icon="el-icon-search"
                style="margin-bottom: 20px" />
            </div>
            <div class="head-container">
              <el-tree ref="tree"
              :data="locOptions"
              :props="defaultProps"
              :expand-on-click-node="false"
                :filter-node-method="filterNode"
                default-expand-all
                @node-click="handleNodeClick" />
            </div>
          </el-col>
          <el-col :span="20" :xs="24">
            <el-form :inline="true">
              <el-form-item label="任务ID">
                <el-input v-model="queryParams.joblds"
                placeholder="请输入任务ID"
                clearable
                size="small"
                @keyup.enter.native="handleQuery" />
              </el-form-item>
              <el-form-item label="状态">
                <el-select v-model="queryParams.logStatus" placeholder="菜单状态" clearable size="small">
                  <el-option label="全部" :value="-1" />
                  <el-option label="成功" :value="1" />
                  <el-option label="失败" :value="2" />
                  <el-option label="进行中" :value="3" />
                </el-select>
              </el-form-item>
              <el-form-item label="任务时间">
                <el-date-picker size="small" type="datetimerange" value-format="yyyy-MM-dd HH:mm:ss"
                  @change="handleChangeTime" />
              </el-form-item>
              <el-form-item>
                <el-button type="primary" icon="el-icon-search" size="mini" @click="getList">搜索</el-button>
                <el-button type="primary" icon="el-icon-refresh" size="mini" @click="getList">重置</el-button>
              </el-form-item>
            </el-form>

            <el-table
              v-loading="loading"
              :data="CounttaskList"
              border row-key="id"
            >
              <el-table-column align="center"  prop="ID" label="编号" width="70" />
              <el-table-column v-align="middle" prop="id" label="任务ID" :show-overflow-tooltip="true" align="center"
                header-align="center" />
              <el-table-column prop="triggerTime" label="任务时间" align="center" />
              <el-table-column prop="handleMsg" label="任务备注" header-align="center" >
                <template slot-scope="{row}">
                  <el-button type="primary" plain style="display: block; margin: 0 auto;" >查看</el-button>

                </template>
              </el-table-column>
              <el-table-column prop="handleCode" label="任务结果" width="120" align="center">
                <template slot-scope="{row}">
                  <span :style="{ color: row.handleCode === 200 ? 'green' : 'red' }">
                    {{ row.handleCode === 200 ? '成功' : '失败' }}
                  </span>
                </template>
              </el-table-column>
              <el-table-column prop="executorHandler" label="执行日志" header-align="center" >
                <template slot-scope="{row}">
                  <el-button type="primary" plain style="display: block; margin: 0 auto;">查看</el-button>
                </template>
              </el-table-column>
            </el-table>

            <pagination
              v-show="total>0"
              :total="total"
              :page.sync="queryParams.pageindex"
              :limit.sync="queryParams.length"
              @pagination="getList"
            />
            <!-- 日志清理模态框 -->
            <div v-if="isModalVisible" class="modal-dialog">
              <div class="modal-content">
                <div class="modal-header">
                  <h4 class="modal-title">日志清理</h4>
                  <button type="button" class="close" @click="closeModal">
                    <span>&times;</span>
                  </button>
                </div>
                <div class="modal-body">
                  <form class="form-horizontal form" role="form">
                    <!-- 模态框内容 -->
                    <!-- ... 表单输入 ... -->
                    <div class="form-group">
                      <label class="col-sm-3 control-label">执行器：</label>
                      <div class="col-sm-9">
                        <input type="text" class="custom-input" readonly="">
							<input type="hidden" name="jobGroup" value="0">
                      </div>
                    </div>
                    <div class="form-group">
                      <label class="col-sm-3 control-label">任务：</label>
                      <div class="col-sm-9">
                        <input type="text" class="custom-input" readonly="">
                        <input type="hidden" name="jobId" value="0">
                      </div>
                    </div>
                    <div class="form-group">
                      <label class="col-sm-3 control-label">清理方式：</label>
                      <div class="col-sm-9">
                        <select v-model="logchose" class="custom-select" name="请选择">
                          <option value="1">清理一个月之前日志数据</option>
                          <option value="2">清理三个月之前日志数据</option>
                          <option value="3">清理六个月之前日志数据</option>
                          <option value="4">清理一年之前日志数据</option>
                          <option value="5">清理一千条以前日志数据</option>
                          <option value="6">清理一万条以前日志数据</option>
                          <option value="7">清理三万条以前日志数据</option>
                          <option value="8">清理十万条以前日志数据</option>
                          <option value="9">清理所有日志数据</option>
                        </select>
                      </div>
                    </div>
                    <hr>
                    <div class="form-group">
                      <div class="col-sm-offset-3 col-sm-6">
                        <button type="button" class="btn btn-primary ok" @click="confirmDelete">确定</button>
                        <button type="button" class="btn btn-default" @click="closeModal">取消</button>
                      </div>
                    </div>
                  </form>
                </div>
              </div>
            </div>
          </el-col>
        </el-row>

      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import { pageList } from '@/api/admin/sys-menu'
import { typetreeselect, treeselect } from '@/api/admin/sys-dept'
import { clearLog1 } from '@/api/admin/sys-menu'
import { locChose } from '@/api/admin/sys-menu'
export default {
  name: 'SysMenuManage',
  data() {
    return {
      //分页数据
      total: 0,
      // 设备名称
      treeInputvalue: undefined,
      // 类型树选项
      cateOptions: undefined,
      locOptions: undefined,
      defaultProps: {
        children: 'children',
        label: 'label'
      },
      // 遮罩层
      loading: true,
      taskList: [],
      // 查询参数
      queryParams: {
        pageindex: 1,
        joblds: null,
        logStatus: -1,
        filterTime: null,
        start: null,
        length: 10
      },

      locquery: {
        locId: null
      },
      logquery: {
        jobId: 0, // 假设 row 是包含 jobId 的对象
        type: null // 假设这是根据某种逻辑获取的 type 值
      },
      logchose: null,
      isModalVisible: false, // 控制模态框可见性的属性
      deleteModalData: null // 用于存储被选择删除的行数据
    }
  },
  computed: {
    CounttaskList() {
      const { taskList, queryParams: { pageindex, length }} = this
      return taskList.map((item, index) => {
        return {
          ...item, // 保留原有数据
          ID: (pageindex - 1) * length + index + 1 // 计算编号
        }
      })
    }
  },
  watch: {
    // 根据名称筛选部门树
    treeInputvalue(val) {
      this.$refs.tree.filter(val)
    }
  },
  created() {
    this.getList()
    this.getTreeselect()
  },
  methods: {
    getTreeselect() {
      treeselect().then(response => {
        this.locOptions = response.data
        console.log('treeOptions', this.locOptions)
      })
    },
    // 筛选节点
    filterNode(value, data) {

      if (!value) return true
      return data.label.indexOf(value) !== -1
    },
    // 树节点单击事件
    handleNodeClick(row) {
      // --todo 获取到节点数据 id ,添加到 queryParams请求参数中,需要确定查询接口对应的区域字段是什么,类型是什么
      this.locquery.locId = row.id
      locChose({locId: this.locquery.locId}).then(response => {
        const dataArray = response.data
        console.log()
        console.log(this.joblds)
        if(dataArray != null){
          this.joblds = dataArray.join(',')
          console.log(this.queryParams)
          this.getListtree(this.joblds) // 调用查询接口
        }else {
          this.queryParams.joblds = ''
          this.taskList = []
        }

      })
    },

    /** 确认删除 */
    confirmDelete(row) {
      this.$confirm('此操作将永久删除该日志, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        // 临时构建一个对象来保存所需的参数
        // 调用 API 进行删除操作
        console.log('Job ID to be deleted:', this.logchose)
        clearLog1(this.logquery).then(response => {
          if (response.code === 200) {
            // 删除成功的逻辑
            this.closeModal()
            this.getList() // 重新加载列表数据
          } else {
            // 删除失败的逻辑
            this.$message.error('操作失败')
          }
        }).catch(error => {
          // 处理错误情况
          console.error('发生错误:', error)
          this.$message.error('发生未知错误')
        })
      }).catch(() => {
        // 用户取消了删除操作
        this.$message.info('已取消删除')
      })
    },
    /** 关闭模态框 */
    closeModal() {
      this.isModalVisible = false
    },


    handleChangeTime(value) {
      if (value) {
        this.queryParams.filterTime = value.join(' - ')
      } else {
        this.queryParams.filterTime = null
      }
    },
    getListtree(joblds) {
      this.queryParams.jobIds = joblds
      console.log(this.queryParams.jobIds)
      pageList(this.queryParams).then(response => {
        this.taskList = this.TimeTranslators(response.content.data)
        console.log(this.taskList)
        this.loading = false
      })
    },
    getList() {
      this.loading = true
      console.log(this.queryParams)
      this.queryParams.start = 1 + ( ( this.queryParams.pageindex-1 ) * this.queryParams.length )
      console.log(this.queryParams.start)
      pageList(this.queryParams).then(response => {
        this.taskList = this.TimeTranslators(response.content.data)
        this.total=response.content.recordsTotal
        this.loading = false
        console.log('respon', this)
        console.log(this.taskList)
      })
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      this.$confirm('是否确认删除?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
      })
    },
    //时间字符串转化
    TimeTranslators(logList) {
      console.log(logList)
      return logList.map(log => {
        const originaltriggerTime = new Date(log.triggerTime); // 将ISO 8601时间字符串转换为Date对象
        const year = originaltriggerTime.getFullYear();
        const month = String(originaltriggerTime.getMonth() + 1).padStart(2, '0');
        const day = String(originaltriggerTime.getDate()).padStart(2, '0');
        const hours = String(originaltriggerTime.getHours()).padStart(2, '0');
        const minutes = String(originaltriggerTime.getMinutes()).padStart(2, '0');
        const seconds = String(originaltriggerTime.getSeconds()).padStart(2, '0');
        const formattedtriggerTime = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`; // 格式化日期时间
        return {
          ...log,
          triggerTime: formattedtriggerTime
        }
      })
    }
  }
}

</script>
<style lang="css">

.panel .el-transfer__buttons {
  width: 150px;
}

.panel .el-transfer__buttons .el-button+.el-button {
  margin-left: 0;
}

.panel .el-transfer-panel {
  width: 300px;
}

.el-col {
  padding: 0 5px;
}

.el-drawer__header {
  margin-bottom: 0;
}

.modal-dialog {
  position: fixed;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
}

.modal-content {
  background: #fff;
  border-radius: 5px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  width: 400px;
  padding: 20px;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-title {
  margin: 0;
  font-size: 18px;
}

.close {
  background: none;
  border: none;
  font-size: 30px;
  color: #ccc;
}

.close:focus {
  outline: none;
}

.btn {
  border: none;
  border-radius: 5px;
  padding: 10px 15px;
  cursor: pointer;
}

.btn-primary {
  background-color: #409eff;
  color: white;
}

.btn-default {
  background-color: #fff;
  color: #409eff;
}
</style>
<style scoped>
/* 自定义输入框，模仿 Element UI 的样式 */
.custom-input {
  display: inline-block;
  width: 100%;
  padding: 0 11px;
  height: 40px; /* Adjust this value based on Element UI's input size */
  line-height: 40px; /* Match the height */
  font-size: 14px;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  background-color: #fff;
  background-image: none;
  appearance: none;
  box-shadow: inset 0 1px 3px rgba(0,0,0,0.1);
  transition: border-color .2s cubic-bezier(.645,.045,.355,1), box-shadow .2s cubic-bezier(.645,.045,.355,1);
}

.custom-input:focus {
  border-color: #409eff;
  outline: none;
  box-shadow: 0 0 0 2px rgba(64,158,255,.2);
}

.custom-input[disabled] {
  background-color: #f5f7fa;
  border-color: #e4e7ed;
  cursor: not-allowed;
}

.custom-input::placeholder {
  color: #a8abb2;
}
/* 自定义下拉选择框，模仿 Element UI 的样式 */
.custom-select {
  display: inline-block;
  width: 100%;
  padding: 0 11px;
  height: 40px; /* Adjust this value based on Element UI's select size */
  line-height: 40px; /* Match the height */
  font-size: 14px;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  background-color: #fff;
  background-image: none;
  appearance: none;
  box-shadow: inset 0 1px 3px rgba(0,0,0,0.1);
  transition: border-color .2s cubic-bezier(.645,.045,.355,1), box-shadow .2s cubic-bezier(.645,.045,.355,1);
}

.custom-select:focus {
  border-color: #409eff;
  outline: none;
  box-shadow: 0 0 0 2px rgba(64,158,255,.2);
}

.custom-select[disabled] {
  background-color: #f5f7fa;
  border-color: #e4e7ed;
  cursor: not-allowed;
}
</style>
