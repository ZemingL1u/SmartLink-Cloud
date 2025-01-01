<template>
  <BasicLayout>
    <template #wrapper>
      <div class="app-container">

        <!-- 侧边框组件，初始时隐藏在视窗外，使用v-show控制显示 -->
        <div v-show="isSidebarVisible" class="sidebar">
          <div class="back-wrapper" @click="toggleSidebar1">
            <i class="fas fa-arrow-left" /> 返回
          </div>
          <div id="options">
            <!-- "详情" 选项卡 -->
            <button class="tab-button" :class="{ active: currentTab === 'content' }"
              @click="showContent('content')">详情</button>
            <!-- "物模型数据" 选项卡 -->
            <button class="tab-button" :class="{ active: currentTab === 'data' }"
              @click="showContent('data')">物模型数据</button>
            <!-- "任务管理" 选项卡 -->
            <button class="tab-button" :class="{ active: currentTab === 'task' }"
              @click="showContent('task')">任务管理</button>
          </div>
          <div v-show="currentTab === 'content'" class="content">
            <h4>设备图片</h4>
            <div v-if="deviceInfo && deviceInfo.type" class="device-image">
              <!-- <img :src="`/${deviceInfo.type.logo}`" :style="{ width: '100px', height: '100px' }"> -->
              <img :src="`${locationOrigin}/${deviceInfo.type.logo}`" :style="{ width: '100px', height: '100px' }">
            </div>
            <h4>设备信息</h4>
            <div>
              <table class="striped-table">
                <thead>
                  <tr>
                    <th>属性</th>
                    <th>值</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-if="deviceInfo && deviceInfo.deviceName">
                    <td>设备名称</td>
                    <td>{{ deviceInfo.deviceName }}</td>
                  </tr>
                  <tr v-if="deviceInfo && deviceInfo.deviceId">
                    <td>设备ID</td>
                    <td>{{ deviceInfo.deviceId }}</td>
                  </tr>
                  <tr v-if="deviceInfo && deviceInfo.type">
                    <td>设备类别</td>
                    <td>{{ deviceInfo.type.typeName }}</td>
                  </tr>
                  <tr v-if="deviceInfo && deviceInfo.sNCode">
                    <td>设备序列号</td>
                    <td>{{ deviceInfo.sNCode }}</td>
                  </tr>
                  <tr v-if="deviceInfo && deviceInfo.loc">
                    <td>区域</td>
                    <td>{{ deviceInfo.loc.locName }}</td>
                  </tr>
                  <tr v-if="deviceInfo && deviceInfo.loc">
                    <td>更新时间</td>
                    <td>{{ deviceInfo.loc.updatedAt }}</td>
                  </tr>
                  <tr v-if="deviceInfo && deviceInfo.QRCode">
                    <td>二维码</td>
                    <td>{{ deviceInfo.QRCode }}</td>
                  </tr>
                </tbody>
              </table>
            </div>
            <div class="button-group">
              <button @click="copyToClipboard(deviceInfo.deviceId)">复制设备ID</button>
              <button @click="copyToClipboard(deviceInfo.key)">复制设备秘钥</button>
            </div>
            <h4>设备实景图</h4>
            <!--          <div id="mapContainer" />-->
                      <div style="display:flex;">
              <div class="device-img" v-for="item in deviceInfo.image">
                <!-- xxxx.jpg -->
                <img :src="`data:image/png;base64,${item}`" :style="{ width: '200px', height: '200px' }">
              </div>
            </div>
              <el-upload
  ref="uploadImgSrc"
  class="upload-demo"
  :limit="2"
  accept="jpg,jpeg,png"
  :show-file-list="false"
  multiple
  action=""
  :http-request="(file) => httpRequest(deviceInfo.deviceId, file)"
>
  <i class="el-icon-upload"></i>
  <div class="el-upload__text"><em>点击上传</em></div>
  <div class="el-upload__tip" slot="tip">只能上传jpg/jpeg/png文件</div>
</el-upload>
          </div>
          <div v-show="currentTab === 'data'" id="data" style="display: none;">
            <!-- 使用 v-for 遍历 devicedata数组 -->
            <div v-for="( dataSegment, index ) in deviceDataSegments.slice(1)" :key="index" class="card">
              <!-- 使用 v-html 渲染带有换行的 HTML 内容 -->
              <div class="mask" v-html="convertCommasToBreaks(dataSegment)" />
            </div>
          </div>
          <div v-show="currentTab === 'task'" class="task">
            <hr>
            <div class="search-box">
              <!-- 搜索框 -->
              <label class="search-label">查询任务</label>
              <input type="text" v-model="searchKeyword" placeholder="请输入任务描述" class="form-control search-input" />
              <el-col :span="1.5">
                <el-button v-permisaction="['admin:sysDevice:add']" type="primary" icon="el-icon-plus" size="mini"
                           @click="showAddModal = true">新增任务</el-button>
              </el-col>
            </div>
            <hr>
            <table class="striped-table">
              <thead>
                <tr>
                  <th>任务描述</th>
                  <!-- <th>调度类型</th> -->
                  <th>状态</th>
                  <!-- <th>设备操作</th> -->
                  <th>操作</th>
                </tr>
              </thead>
              <tbody>
                <!-- 使用计算属性filteredTasks来渲染任务 -->
                <tr v-for="task in filteredTasks" :key="task.id">
                  <td>{{ task.jobDesc }}</td>
                  <!-- <td>{{ task.scheduleType }}</td> -->
                  <td>{{ task.triggerStatus }}</td>
                  <!-- <td>{{ task.executorHandler }}</td> -->
                  <td>
                    <!-- <el-button v-permisaction="['admin:sysDevice:edit']" type="success" icon="el-icon-edit"
                      size="mini" @click="showEditModal1(task)"
                      style="background-color: #13ce66;border-color: #13ce66;">修改</el-button>
                    <el-button size="mini" type="text" icon="el-icon-delete"
                      style="background-color:#ff4949;color: white;width: 70px; padding-left: 5px;padding-right: 5px;"
                      @click="deleteTask(task)">删除</el-button> -->

                      <div>
                        <span class="is-link" v-for="item in operations" :key="item.command" @click="handleDropdownCommand(item.command, task)">{{  item.text  }}</span>
                      </div>
                    <!-- <el-dropdown @command="handleDropdownCommand($event, task)">
                      <span class="el-dropdown-link">
                        请选择<i class="el-icon-arrow-down el-icon--right"></i>
                      </span>
                      <el-dropdown-menu slot="dropdown">
                        <el-dropdown-item v-for="item in operations" :key="item.command" :command="item.command"
                          :task="task">
                          {{ item.text }}
                        </el-dropdown-item>
                      </el-dropdown-menu>
                    </el-dropdown> -->
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
      <el-card class="box-card">
        <el-row :gutter="20">
          <!--设备型号-->
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
          <!--设备数据-->
          <el-col :span="20" :xs="24">
            <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
              <el-form-item label="设备名称" prop="deviceName">
                <el-input v-model="queryParams.deviceName" placeholder="请输入设备名称" clearable size="small"
                  style="width: 160px" @keyup.enter.native="handleQuery" />
              </el-form-item>
              <el-form-item label="设备类型" prop="typeName">
                <el-input v-model="queryParams.type.typeName" placeholder="请输设备类型" clearable size="small"
                  style="width: 160px" @keyup.enter.native="handleQuery" />
              </el-form-item>
              <el-form-item label="设备区域" prop="locId">
                <el-input v-model="queryParams.loc.locName" placeholder="请输入设备区域" clearable size="small"
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
                  :disabled="single" @click="handleUpdate">修改</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button v-permisaction="['admin:sysDevice:remove']" type="danger" icon="el-icon-delete" size="mini"
                  :disabled="multiple" @click="handleDelete">删除</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button v-permisaction="['admin:sysDevice:add']" type="primary" icon="el-icon-plus" size="mini"
                           :disabled="multiple" @click="showAddModal = true">新增任务</el-button>
              </el-col>
               <!-- <el-col :span="1.5">
                <el-button v-permisaction="['admin:sysDevice:edit']" type="success" icon="el-icon-edit" size="mini"
                @click="handleAddDevice" 
                style="background-color: blue">新增设备
              </el-button>
              </el-col> -->
            </el-row>


            <el-table v-loading="loading" :data="DeviceList" border @selection-change="handleSelectionChange"
              @sort-change="handleSortChang">
              <el-table-column type="selection" width="55" align="center" />
              <el-table-column label="编号" width="75" prop="编号" sortable="custom" align="center"/>
              <!-- <el-table-column label="编号" width="75" prop="userId" sortable="custom" /> -->
              <el-table-column label="设备名称" prop="deviceName" :show-overflow-tooltip="true" align="center"/>
              <el-table-column label="区域" prop="loc.locName" :show-overflow-tooltip="true" align="center"/>
              <el-table-column label="设备类型" prop="type.typeName" width="170" align="center"/>
              <el-table-column label="序列号" prop="sNCode" :show-overflow-tooltip="true" align="center"/>
              <el-table-column label="状态" width="80" prop="status" align="center"/>
              <el-table-column label="操作" width="250" fix="right" class-name="small-padding fixed-width" align="center">
                <template slot-scope="scope">
                  <!--v-permisaction="['admin:sysUser:edit']"-->
                  <el-button size="mini" type="text" icon="el-icon-edit"
                    @click="toggleSidebar(scope.row)">详情</el-button>
                  <el-button size="mini" type="text" icon="el-icon-plus"
                    @click="showAddModal1(scope.row)">新增任务</el-button>
                  <el-button v-if="scope.row.userId !== 1" v-permisaction="['admin:sysUser:remove']" size="mini"
                    type="text" icon="el-icon-edit" @click="handleUpdate(scope.row)">修改</el-button>
                  <el-button v-permisaction="['admin:sysUser:resetPassword']" size="mini" type="text"
                    icon="el-icon-delete" @click="handleDelete(scope.row)">删除</el-button>
                </template>
              </el-table-column>
            </el-table>

            <pagination v-show="total > 0" :total="total" :page.sync="queryParams.pageIndex"
              :limit.sync="queryParams.pageSize" @pagination="getList" />
          </el-col>
        </el-row>
      </el-card>
      <!-- 修改参数配置对话框 -->
      <el-dialog :title="title" :visible.sync="open" width="600px" :close-on-click-modal="false">
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
                <!-- <el-input v-model="addDeviceInfo.type.typeName" maxlength="11" /> -->
                <el-select v-model="addInfoForm.typeId">
                  <el-option v-for="item in categoryOptions" :key="typeId" :label="item.typeName" :value="item.typeId"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="序列号" prop="sNCode">
                <el-input v-model="deviceInfo.sNCode" placeholder="请输入序列号" maxlength="50" />
              </el-form-item>
            </el-col>
            <!-- <el-col :span="24">
              <el-form-item label="物理信息" prop="deviceData">
                <el-input v-model="deviceInfo.deviceData" type="textarea" />
              </el-form-item>
            </el-col> -->
            <!-- <el-col :span="24">
              <el-form-item label="更新时间" prop="updateAt">
                <el-input v-model="deviceInfo.updatedAt" maxlength="50" />
              </el-form-item>
            </el-col> -->
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

      <!--添加设备对话框-->
       <el-dialog :title="title" :visible.sync="addDevice" width="600px" :close-on-click-modal="false">
        <el-form ref="form" :model="addInfoForm" :rules="rules" label-width="80px" @close="handleClose">
          <el-row>
            <el-col :span="12">
              <el-form-item label="设备名称" prop="deviceName">
                <el-input v-model="addInfoForm.deviceName" placeholder="请输入设备名称" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="归属区域" prop="locId">
                <treeselect v-model="addInfoForm.locId" :options="locOptions" placeholder="请选择归属区域" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="设备类型" prop="deviceType">
                <!-- <el-input v-model="addDeviceInfo.type.typeName" maxlength="11" /> -->
                <el-select v-model="addInfoForm.typeId">
                  <el-option v-for="item in categoryOptions" :key="typeId" :label="item.typeName" :value="item.typeId"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="序列号" prop="sNCode">
                <el-input v-model="addInfoForm.sNCode" placeholder="请输入序列号" maxlength="50" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="备注">
                <el-input v-model="addInfoForm.key" placeholder="请输入备注" maxlength="50" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="设备状态">
                <el-select v-model="addInfoForm.status">
                  <el-option label="在线" value="在线"></el-option>
                  <el-option label="待机" value="待机"></el-option>
                  <el-option label="故障" value="故障"></el-option>
                  <el-option label="待激活" value="待激活"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
          </el-row><!--名称 类别 区域 密钥 备注-->
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button type="primary" @click.prevent="submitForm">确 定</el-button>
          <el-button @click="canceladd">取 消</el-button>
        </div>
      </el-dialog>
      <!-- 用户导入对话框 -->
      <el-dialog :title="upload.title" :visible.sync="upload.open" width="400px" :close-on-click-modal="false">
        <el-upload ref="upload" :limit="1" accept=".xlsx, .xls" :headers="upload.headers"
          :action="upload.url + '?updateSupport=' + upload.updateSupport" :disabled="upload.isUploading"
          :on-progress="handleFileUploadProgress" :on-success="handleFileSuccess" :auto-upload="false" drag>
          <i class="el-icon-upload" />
          <div class="el-upload__text">
            将文件拖到此处，或
            <em>点击上传</em>
          </div>
          <div slot="tip" class="el-upload__tip">
            <el-checkbox v-model="upload.updateSupport" />是否更新已经存在的用户数据
            <el-link type="info" style="font-size:12px" @click="importTemplate">下载模板</el-link>
          </div>
          <div slot="tip" class="el-upload__tip" style="color:red">提示：仅允许导入“xls”或“xlsx”格式文件！</div>
        </el-upload>
        <div slot="footer" class="dialog-footer">
          <el-button type="primary" @click="submitFileForm">确 定</el-button>
          <el-button @click="upload.open = false">取 消</el-button>
        </div>
      </el-dialog>

 <!-- job新增.模态框 -->
 <div class="modal-backdrop" id="addModal" tabindex="-1" role="dialog" aria-hidden="true" v-if="showAddModal">
        <div class="modal">
          <div class="modal-content">
            <div class="modal-header">
              <h4 class="modal-title">新增任务</h4>
            </div>
            <!--任务模板-->
            <div class="col-sm-4">
              <p style="margin: 0 0 10px;text-align: left;border-bottom: 1px solid #e5e5e5;color: gray;">任务模板</p>
              <label class="col-sm-2 control-label" style="margin-left: 93px;margin-right: 15px;">任务模板</label>
              <select class="form-control scheduleType" v-model="selectedOption"
                @change="selectTemplate(selectedOption)" style="width: 205px;">
                <option value="">请选择任务模板</option>
                <option v-for="option in templateData" :key="option.templateId" :value="option.templateId">
                  {{ option.templateName }}
                </option>
                <option value="aaa">自定义</option>
                <option value="bbb">新增模板</option>
              </select>
            </div>

            <div class="modal-body">
              <form class="form-horizontal form" role="form">

                <p style="margin: 0 0 10px;text-align: left;border-bottom: 1px solid #e5e5e5;color: gray;">基础配置</p>
                <!-- <div class="form-group">
                  <label for="firstname" class="col-sm-2 control-label">执行器<font color="red">*</font></label>
                  <div class="col-sm-4">
                    <select class="form-control" name="jobGroup">
                      <option value="1">示例执行器</option>
                    </select>
                  </div>

                  <label for="lastname" class="col-sm-2 control-label">任务描述<font color="red">*</font></label>
                  <div class="col-sm-4"><input type="text" class="form-control" name="jobDesc" placeholder="请输入任务描述"
                      maxlength="50"></div>
                </div> -->
                <div class="form-group">
                  <label for="lastname" class="col-sm-2 control-label">负责人<font color="red">*</font></label>
                  <div class="col-sm-4"><input type="text" class="form-control" v-model="author" placeholder="请输入负责人"
                      maxlength="50" disabled></div>
                  <label for="lastname" class="col-sm-2 control-label">报警邮件<font color="black">*</font></label>
                  <div class="col-sm-4"><input type="text" class="form-control" v-model="alarmEmail"
                      placeholder="输入多个则用逗号分隔" maxlength="100" disabled></div>
                </div>


                <p style="margin: 0 0 10px;text-align: left;border-bottom: 1px solid #e5e5e5;color: gray;">调度配置</p>
                <div class="form-group">
                  <label for="firstname" class="col-sm-2 control-label">调度类型<font color="red">*</font></label>
                  <div class="col-sm-4" style="margin-left: -15px;">
                    <select class="form-control scheduleType" name="scheduleType" v-model="selectedScheduleType"
                      style="width: 202px;">
                      <option value="NONE">无</option>
                      <option value="CRON" selected>CRON</option>
                      <option value="FIX_RATE">固定速度</option>
                    </select>
                  </div>

                  <!-- 使用v-if来判断是否显示CRON输入框 -->
                  <div v-if="selectedScheduleType === 'CRON'" class="col-sm-4" style="margin-left: 20px;">
                    <label for="lastname" class="control-label">Cron<font color="red">*</font></label>
                    <input type="text" class="form-control" v-model="schedule_conf_CRON" placeholder="请输入Cron"
                      maxlength="128" style="margin-left: 5px;">
                  </div>

                  <!-- 使用v-else来判断是否显示固定速度输入框 -->
                  <div v-else-if="selectedScheduleType === 'FIX_RATE'" class="col-sm-4" style="margin-left: 20px;">
                    <label for="lastname" class="control-label" style="margin-left: -20px;">固定速度<font color="red">*
                      </font>
                    </label>
                    <input type="text" class="form-control" v-model="schedule_conf_FIX_RATE" placeholder="请输入 （ 秒 ）"
                      maxlength="10" onkeyup="this.value=this.value.replace(/\D/g,'')"
                      onafterpaste="this.value=this.value.replace(/\D/g,'')" style="margin-left: -5px;">
                  </div>
                </div>

                <p style="margin: 0 0 10px;text-align: left;border-bottom: 1px solid #e5e5e5;color: gray;">设备配置</p>
                <div class="form-group">
                  <label for="firstname" class="col-sm-2 control-label">区域<font color="red">*</font></label>
                  <div class="col-sm-4" style="margin-left: 15px;"><input type="text" class="form-control"
                      v-model="locName" maxlength="100" disabled>
                  </div>
                  <label for="firstname" class="col-sm-2 control-label">设备类型<font color="red">*</font></label>
                  <div class="col-sm-4" style="margin-left: -3px">
                    <input type="text" class="form-control" v-model="typeName" maxlength="100" disabled>
                  </div>
                </div>


                <p style="margin: 0 0 10px;text-align: left;border-bottom: 1px solid #e5e5e5;color: gray;">任务配置</p>

                <div class="form-group">
                  <label for="firstname" class="col-sm-2 control-label">设备操作<font color="red">*</font></label>
                  <div class="col-sm-4" style="margin-left: -15px;">
                    <select class="form-control deviceOperate" v-model="executorHandler" style="width:205px;">
                      <option value="get" selected>获取传感器数据</option>
                      <option value="post">上传传感器数据</option>
                    </select>
                  </div>
                  <label for="firstname" class="col-sm-2 control-label">任务描述<font color="red">*</font></label>
                  <div class="col-sm-4" style="margin-left: -5px;">
                    <input type="text" class="form-control" v-model="jobDesc" placeholder="请输入任务描述" maxlength="100">
                  </div>
                </div>

                <div class="form-group">
                  <label for="firstname" class="col-sm-2 control-label">任务参数<font color="black">*</font></label>
                  <div class="col-sm-10" style="margin-left: -15px;">
                    <textarea class="textarea form-control" v-model="executorParam" placeholder="请输入任务参数"
                      maxlength="512" style="height: 63px; line-height: 1.2;width: 572px;"></textarea>
                  </div>
                </div>

                <div class="form-group">
                  <label for="lastname" class="col-sm-2 control-label">任务超时时间</label>
                  <div class="col-sm-4" style="margin-left: -20px;">
                    <input type="text" class="form-control" v-model="executorTimeout" placeholder="任务超时时间(秒)，大于零时生效"
                      maxlength="6" onkeyup="this.value=this.value.replace(/\D/g,'')"
                      onafterpaste="this.value=this.value.replace(/\D/g,'')" style="width: 245px;">
                  </div>
                  <label for="lastname" class="col-sm-2 control-label" style="margin-left: -60px;">失败重试次数</label>
                  <div class="col-sm-4">
                    <input type="text" class="form-control" v-model="executorFailRetryCount" placeholder="失败重试次数，大于零时生效"
                      maxlength="4" onkeyup="this.value=this.value.replace(/\D/g,'')"
                      onafterpaste="this.value=this.value.replace(/\D/g,'')" style="width: 220px;margin-left: -20px;">
                  </div>
                </div>

                <div class="form-group">
                  <div class="col-sm-offset-3 col-sm-6">
                    <button type="submit" class="btn btn-primary" @click.prevent="saveFormData">保存</button>
                    <button type="button" class="btn btn-default" data-dismiss="modal"
                      @click="showAddModal = false">取消</button>
                  </div>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
      <!-- job批量新增.模态框 -->
      <div class="modal-backdrop" id="addModal" tabindex="-1" role="dialog" aria-hidden="true" v-if="showAddModal2">
        <div class="modal">
          <div class="modal-content">
            <div class="modal-header">
              <h4 class="modal-title">批量新增任务</h4>
            </div>
            <!--任务模板-->
            <div class="col-sm-4">
              <label class="col-sm-2 control-label" style="margin-left: 93px;margin-right: 15px;">任务模板</label>
              <select class="form-control scheduleType" v-model="selectedOption"
                @change="selectTemplate(selectedOption)" style="width: 205px;">
                <option value="">请选择任务模板</option>
                <option v-for="option in templateData" :key="option.templateId" :value="option.templateId">
                  {{ option.templateName }}
                </option>
                <option value="aaa">自定义</option>
              </select>
            </div>
            <div class="modal-body">
              <form class="form-horizontal form" role="form">

                <p style="margin: 0 0 10px;text-align: left;border-bottom: 1px solid #e5e5e5;color: gray;">基础配置</p>
                <!-- <div class="form-group">
                  <label for="firstname" class="col-sm-2 control-label">执行器<font color="red">*</font></label>
                  <div class="col-sm-4">
                    <select class="form-control" name="jobGroup">
                      <option value="1">示例执行器</option>
                    </select>
                  </div>

                  <label for="lastname" class="col-sm-2 control-label">任务描述<font color="red">*</font></label>
                  <div class="col-sm-4"><input type="text" class="form-control" name="jobDesc" placeholder="请输入任务描述"
                      maxlength="50"></div>
                </div> -->
                <div class="form-group">
                  <label for="lastname" class="col-sm-2 control-label">负责人<font color="red">*</font></label>
                  <div class="col-sm-4"><input type="text" class="form-control" v-model="author" placeholder="请输入负责人"
                      maxlength="50" disabled></div>
                  <label for="lastname" class="col-sm-2 control-label">报警邮件<font color="black">*</font></label>
                  <div class="col-sm-4"><input type="text" class="form-control" v-model="alarmEmail"
                      placeholder="输入多个则用逗号分隔" maxlength="100" disabled></div>
                </div>

                <br>
                <p style="margin: 0 0 10px;text-align: left;border-bottom: 1px solid #e5e5e5;color: gray;">调度配置</p>
                <div class="form-group">
                  <label for="firstname" class="col-sm-2 control-label">调度类型<font color="red">*</font></label>
                  <div class="col-sm-4" style="margin-left: -15px;">
                    <select class="form-control scheduleType" name="scheduleType" v-model="selectedScheduleType"
                      style="width: 202px;">
                      <option value="NONE">无</option>
                      <option value="CRON" selected>CRON</option>
                      <option value="FIX_RATE">固定速度</option>
                    </select>
                  </div>

                  <!-- 使用v-if来判断是否显示CRON输入框 -->
                  <div v-if="selectedScheduleType === 'CRON'" class="col-sm-4" style="margin-left: 20px;">
                    <label for="lastname" class="control-label">Cron<font color="red">*</font></label>
                    <input type="text" class="form-control" v-model="schedule_conf_CRON" placeholder="请输入Cron"
                      maxlength="128" style="margin-left: 5px;">
                  </div>

                  <!-- 使用v-else来判断是否显示固定速度输入框 -->
                  <div v-else-if="selectedScheduleType === 'FIX_RATE'" class="col-sm-4" style="margin-left: 20px;">
                    <label for="lastname" class="control-label" style="margin-left: -20px;">固定速度<font color="red">*
                      </font>
                    </label>
                    <input type="text" class="form-control" v-model="schedule_conf_FIX_RATE" placeholder="请输入 （ 秒 ）"
                      maxlength="10" onkeyup="this.value=this.value.replace(/\D/g,'')"
                      onafterpaste="this.value=this.value.replace(/\D/g,'')" style="margin-left: -5px;">
                  </div>
                </div>

                <br>
                <p style="margin: 0 0 10px;text-align: left;border-bottom: 1px solid #e5e5e5;color: gray;">设备配置</p>
                <div class="form-group">
                  <label for="firstname" class="col-sm-2 control-label">区域<font color="red">*</font></label>
                  <div class="col-sm-4" style="margin-left: 15px;"><input type="text" class="form-control"
                      v-model="locName" maxlength="100" disabled>
                  </div>
                  <label for="firstname" class="col-sm-2 control-label">设备类型<font color="red">*</font></label>
                  <div class="col-sm-4" style="margin-left: -3px">
                    <input type="text" class="form-control" placeholder="0" maxlength="100" disabled>
                  </div>
                </div>

                <br>
                <p style="margin: 0 0 10px;text-align: left;border-bottom: 1px solid #e5e5e5;color: gray;">任务配置</p>

                <div class="form-group">
                  <label for="firstname" class="col-sm-2 control-label">设备操作<font color="red">*</font></label>
                  <div class="col-sm-4" style="margin-left: -15px;">
                    <select class="form-control deviceOperate" v-model="executorHandler" style="width:205px;">
                      <option value="get" selected>获取传感器数据</option>
                      <option value="post">上传传感器数据</option>
                    </select>
                  </div>
                  <label for="firstname" class="col-sm-2 control-label">任务描述<font color="red">*</font></label>
                  <div class="col-sm-4" style="margin-left: -5px;">
                    <input type="text" class="form-control" v-model="jobDesc" placeholder="请输入任务描述" maxlength="100">
                  </div>
                </div>

                <div class="form-group">
                  <label for="firstname" class="col-sm-2 control-label">任务参数<font color="black">*</font></label>
                  <div class="col-sm-10" style="margin-left: -15px;">
                    <textarea class="textarea form-control" v-model="executorParam" placeholder="请输入任务参数"
                      maxlength="512" style="height: 63px; line-height: 1.2;width: 572px;"></textarea>
                  </div>
                </div>

                <div class="form-group">
                  <label for="lastname" class="col-sm-2 control-label">任务超时时间</label>
                  <div class="col-sm-4" style="margin-left: -20px;">
                    <input type="text" class="form-control" v-model="executorTimeout" placeholder="任务超时时间(秒)，大于零时生效"
                      maxlength="6" onkeyup="this.value=this.value.replace(/\D/g,'')"
                      onafterpaste="this.value=this.value.replace(/\D/g,'')" style="width: 245px;">
                  </div>
                  <label for="lastname" class="col-sm-2 control-label" style="margin-left: -60px;">失败重试次数</label>
                  <div class="col-sm-4">
                    <input type="text" class="form-control" v-model="executorFailRetryCount" placeholder="失败重试次数，大于零时生效"
                      maxlength="4" onkeyup="this.value=this.value.replace(/\D/g,'')"
                      onafterpaste="this.value=this.value.replace(/\D/g,'')" style="width: 220px;margin-left: -20px;">
                  </div>
                </div>

                <div class="form-group">
                  <div class="col-sm-offset-3 col-sm-6">
                    <button type="submit" class="btn btn-primary" @click.prevent="batchSave">保存</button>
                    <button type="button" class="btn btn-default" data-dismiss="modal"
                      @click="showAddModal2 = false">取消</button>
                  </div>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
      <!-- job修改.模态框 -->
      <div class="modal-backdrop" id="editModal" tabindex="-1" role="dialog" aria-hidden="true" v-if="showEditModal">
        <div class="modal">
          <div class="modal-content">
            <div class="modal-header">
              <h4 class="modal-title">修改任务</h4>
            </div>
            <div class="modal-body">
              <form class="form-horizontal form" role="form">

                <p style="margin: 0 0 10px;text-align: left;border-bottom: 1px solid #e5e5e5;color: gray;">基础配置</p>
                <!-- <div class="form-group">
                  <label for="firstname" class="col-sm-2 control-label">执行器<font color="red">*</font></label>
                  <div class="col-sm-4">
                    <select class="form-control" name="jobGroup">
                      <option value="1">示例执行器</option>
                    </select>
                  </div>

                  <label for="lastname" class="col-sm-2 control-label">任务描述<font color="red">*</font></label>
                  <div class="col-sm-4"><input type="text" class="form-control" name="jobDesc" placeholder="请输入任务描述"
                      maxlength="50"></div>
                </div> -->
                <div class="form-group">
                  <label for="lastname" class="col-sm-2 control-label">负责人<font color="red">*</font></label>
                  <div class="col-sm-4"><input type="text" class="form-control" v-model="author" placeholder="请输入负责人"
                      maxlength="50"></div>
                  <label for="lastname" class="col-sm-2 control-label">报警邮件<font color="black">*</font></label>
                  <div class="col-sm-4"><input type="text" class="form-control" v-model="alarmEmail"
                      placeholder="输入多个则用逗号分隔" maxlength="100"></div>
                </div>

                <br>
                <p style="margin: 0 0 10px;text-align: left;border-bottom: 1px solid #e5e5e5;color: gray;">调度配置</p>
                <div class="form-group">
                  <label for="firstname" class="col-sm-2 control-label">调度类型<font color="red">*</font></label>
                  <div class="col-sm-4" style="margin-left: -15px;">
                    <select class="form-control scheduleType" name="scheduleType" v-model="selectedScheduleType"
                      style="width: 202px;">
                      <option value="NONE">无</option>
                      <option value="CRON" selected>CRON</option>
                      <option value="FIX_RATE">固定速度</option>
                    </select>
                  </div>

                  <!-- 使用v-if来判断是否显示CRON输入框 -->
                  <div v-if="selectedScheduleType === 'CRON'" class="col-sm-4" style="margin-left: 20px;">
                    <label for="lastname" class="control-label">Cron<font color="red">*</font></label>
                    <input type="text" class="form-control" v-model="schedule_conf_CRON" placeholder="请输入Cron"
                      maxlength="128" style="margin-left: 5px;">
                  </div>


                  <!-- 使用v-else来判断是否显示固定速度输入框 -->
                  <div v-else-if="selectedScheduleType === 'FIX_RATE'" class="col-sm-4" style="margin-left: 20px;">
                    <label for="lastname" class="control-label" style="margin-left: -20px;">固定速度<font color="red">*
                      </font>
                    </label>
                    <input type="text" class="form-control" v-model="schedule_conf_FIX_RATE" placeholder="请输入 （ 秒 ）"
                      maxlength="10" onkeyup="this.value=this.value.replace(/\D/g,'')"
                      onafterpaste="this.value=this.value.replace(/\D/g,'')" style="margin-left: -5px;">
                  </div>
                </div>

                <br>
                <p style="margin: 0 0 10px;text-align: left;border-bottom: 1px solid #e5e5e5;color: gray;">设备配置</p>
                <div class="form-group">
                  <label for="firstname" class="col-sm-2 control-label">区域<font color="red">*</font></label>
                  <div class="col-sm-4" style="margin-left: 15px;">
                    <select class="form-control area" name="area" style="width: 205px" disabled>
                      <option value="L1-203">L1-203</option>
                      <option value="L1-204" selected>L1-204</option>
                      <option value="L1-205">L1-205</option>
                    </select>
                  </div>
                  <label for="firstname" class="col-sm-2 control-label">设备类型<font color="red">*</font></label>
                  <div class="col-sm-4" style="margin-left: -3px">
                    <select class="form-control deviceType" name="deviceType" style="width: 200px" disabled>
                      <option value="数传终端">数传终端</option>
                      <option value="传感器" selected>传感器</option>
                      <option value="温度传感器">温度传感器</option>
                    </select>
                  </div>
                </div>

                <br>
                <p style="margin: 0 0 10px;text-align: left;border-bottom: 1px solid #e5e5e5;color: gray;">任务配置</p>

                <div class="form-group">
                  <label for="firstname" class="col-sm-2 control-label">设备操作<font color="red">*</font></label>
                  <div class="col-sm-4" style="margin-left: -15px;">
                    <select class="form-control deviceOperate" v-model="executorHandler" style="width:205px;">
                      <option value="get" selected>获取传感器数据</option>
                      <option value="post">上传传感器数据</option>
                    </select>
                  </div>
                  <label for="firstname" class="col-sm-2 control-label">任务描述<font color="red">*</font></label>
                  <div class="col-sm-4" style="margin-left: -5px;">
                    <input type="text" class="form-control" v-model="jobDesc" placeholder="请输入任务描述" maxlength="100">
                  </div>
                </div>

                <div class="form-group">
                  <label for="firstname" class="col-sm-2 control-label">任务参数<font color="black">*</font></label>
                  <div class="col-sm-10" style="margin-left: -15px;">
                    <textarea class="textarea form-control" v-model="executorParam" placeholder="请输入任务参数"
                      maxlength="512" style="height: 63px; line-height: 1.2;width: 572px;"></textarea>
                  </div>
                </div>

                <div class="form-group">
                  <label for="lastname" class="col-sm-2 control-label">任务超时时间</label>
                  <div class="col-sm-4" style="margin-left: -20px;">
                    <input type="text" class="form-control" v-model="executorTimeout" placeholder="任务超时时间(秒)，大于零时生效"
                      maxlength="6" onkeyup="this.value=this.value.replace(/\D/g,'')"
                      onafterpaste="this.value=this.value.replace(/\D/g,'')" style="width: 245px;">
                  </div>
                  <label for="lastname" class="col-sm-2 control-label" style="margin-left: -60px;">失败重试次数</label>
                  <div class="col-sm-4">
                    <input type="text" class="form-control" v-model="executorFailRetryCount" placeholder="失败重试次数，大于零时生效"
                      maxlength="4" onkeyup="this.value=this.value.replace(/\D/g,'')"
                      onafterpaste="this.value=this.value.replace(/\D/g,'')" style="width: 220px;margin-left: -20px;">
                  </div>
                </div>

                <div class="form-group">
                  <div class="col-sm-offset-3 col-sm-6">
                    <button type="submit" class="btn btn-primary" @click.prevent="saveFormData2">保存</button>
                    <button type="button" class="btn btn-default" data-dismiss="modal"
                      @click="showEditModal = false">取消</button>
                  </div>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
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
import { listDevice, addDevice, getDeviceIdByLocId, getlistByDeviceIds } from '@/api/sys-device'
import { listCategory } from '@/api/sys-Category'
import {JobStart, logadd, Single} from "@/api/sys-task";

export default {
  name: 'SysDeviceManage',
  components: { Treeselect },
  data() {
    return {
      locationOrigin: location.origin,
      //上传图片
      uploadType: 1, // 默认上传一张图片
      files: [],

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
        {
          text: '启动',
           command: 'start',
           click: (task) => this.handleDropdownCommand('start', task)
        },
        {
          text: '单次执行',
           command: 'SingleStart',
           click: (task) => this.handleDropdownCommand('SingleStart', task)
        },
      ],
      SingledialogVisible: undefined,
      ifHaveAddTemplate: 0,//判断是否已经添加过模板
      isCustomTemplate: false, //判断是否选择了自定义模板
      searchKeyword: '',//搜索关键词
      templateData: [],//任务输入模板
      currentTemplate: null, // 当前选中的模板数据
      selectedOption: '',// 当前选中的选项ID
      tasks: [], // 用于存储任务数据
      showAddModal: false, // 控制模态框显示的变量
      showAddModal2: false, // 控制模态框显示的变量
      showEditModal: false, // 控制模态框显示的变量
      author: '',            // 负责人
      alarmEmail: '',        // 报警邮件
      selectedScheduleType: 'CRON', // 调度类型
      schedule_conf_CRON: '', // CRON表达式
      schedule_conf_FIX_RATE: '', // 固定速度（秒）
      executorHandler: '',   // 设备操作
      jobDesc: '',           // 任务描述
      executorParam: '',     // 任务参数
      executorTimeout: '',   // 任务超时时间
      executorFailRetryCount: '', // 失败重试次数
      deviceId: '',
      typeId: '',
      locId: '',
      jobId: '',
      // 控制侧边框显示状态的响应式变量
      isSidebarVisible: false,
      currentTab: 'content',
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
      DeviceList: null,
      // 弹出层标题
      title: '',
      // 类型树选项
      cateOptions: undefined,
      locOptions: undefined,
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
        token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhc2NvcGUiOiIiLCJleHAiOjQ4NjY5NjU0OTUsImlkZW50aXR5IjoxLCJuaWNlIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTcxMzMyOTQ5NSwicm9sZWlkIjoxLCJyb2xla2V5IjoiYWRtaW4iLCJyb2xlbmFtZSI6Iuezu-e7n-euoeeQhuWRmCJ9.6KE1GFMEFuLllv6xxws0GP4XsJriRn-29yJKJpAzWLU', /* 请替换成实际的token */
        pageIndex: 1,
        pageSize: 10,
        deviceName: undefined,
        type: {
          typeName: ''
        },
        loc: {
          locName: ''
        }
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
        type: {
          typeName: '',
          typeId: ''
        },
        sNCode: '',
        loc: {
          locName: ''
        },
      },
      //添加设备参数
      addDeviceInfo: {
        deviceName: '',
        locId: null,
        deviceId: '',
        type: {
          typeName: '',
          typeId: ''
        },
        sNCode: '',
        loc: {
          locName: ''
        },
        updatedAt: '',
        remark: ''
      },
      categoryOptions: [],
      addInfoForm: {
        typeId: null,
        locId: null,
        deviceName: '',
        sNCode: '',
        key: '',
        status: ''
      }
    }
  },
  computed: {

    //上传图片
    uploadLimit() {
    return this.uploadType === 2 ? 2 : 1;
  },

    // 计算编号的计算属性
    // indexedDeviceList() {
    //   // 根据列表数据和当前分页的索引计算编号
    //   const { DeviceList, queryParams: { pageIndex, pageSize } } = this
    //   return DeviceList.map((item, index) => {
    //     return {
    //       ...item, // 保留原有数据
    //       编号: (pageIndex - 1) * pageSize + index + 1 // 计算编号
    //     }
    //   })
    // }
     // 根据搜索关键词过滤任务列表
     filteredTasks() {
      return this.tasks.filter(task => {
        // 检查任务描述是否包含搜索关键词（不区分大小写）
        return task.jobDesc.toLowerCase().includes(this.searchKeyword.toLowerCase());
      });
    },
    // 计算属性，根据当前选中的模板数据来确定要在表格中显示的数据
    displayData() {
      // 检查是否已经有选中的模板数据
      if (!this.currentTemplate) {
        return []; // 如果没有选中的模板数据，则显示空数组
      }

      // 返回当前模板对应的数据数组
      return this.currentTemplate.data;
    }
  },
  watch: {
    // 根据名称筛选部门树
    treeInputvalue(val) {
      this.$refs.tree.filter(val)
    }
  },
  created() {
    this.getListCategory()
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
  methods: {
     // 处理下拉菜单点击事件
     handleDropdownCommand(command, task) {
      switch (command) {
        case 'edit':
          this.showEditModal1(task);
          break;
        case 'delete':
          this.deleteTask(task);
          break;
         case 'start':
           this.Taskstart(task);
           break;
         case 'SingleStart':
           this.Singletask(task);
           break;
      }
    },
     // 当用户从下拉菜单中选择一个选项时调用此方法
     selectTemplate(optionId) {
      // 更新当前选中的模板数据
      if (optionId === '') {
        // 用户尚未选择有效模板，可以显示提示信息或执行其他逻辑
        this.currentTemplate = null;
      }
      else if (optionId === 'aaa') {
        //自定义无操作
      }
      else if (optionId === 'bbb') {
        this.isCustomTemplate = true;
      }
      else {
        // 用户选择了一个模板，执行更新操作
        this.currentTemplate = this.templateData.find(
          (template) => template.templateId === optionId
        );
        this.executorFailRetryCount = this.currentTemplate.executorFailRetryCount
        this.executorHandler = this.currentTemplate.executorHandler
        this.executorParam = this.currentTemplate.executorParam
        this.executorTimeout = this.currentTemplate.executorTimeout
        this.glueType = this.currentTemplate.glueType
        this.jobDesc = this.currentTemplate.jobDesc
        this.scheduleConf = this.currentTemplate.scheduleConf
        this.schedule_conf_CRON = this.currentTemplate.scheduleConfCron
        this.scheduleConfFixDelay = this.currentTemplate.scheduleConfFixDelay
        this.schedule_conf_FIX_RATE = this.currentTemplate.scheduleConfFixRate
        this.scheduleType = this.currentTemplate.scheduleType
      }

    },
    getListCategory() {
      listCategory().then(res => {
        if(res.data) {
          this.categoryOptions = res.data
        }
      })
    },
      //获取任务模板
      getTemplateData() {
      const url = `http://120.79.59.158:8000/api/v1/jobTemplate/pageList?typeId=${encodeURIComponent(this.typeId)}`;
      fetch(url, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${getToken()}` // 使用Bearer Token进行认证
        }
      })
        .then(response => response.json())
        .then(data => {
          this.templateData = data.data.list;
          console.log(this.templateData)
        })
        .catch(error => {
          console.error('Error fetching schedule types:', error);
          // 错误处理
        });
    },
    //新增任务框
    showAddModal1(row) {
      console.log(row)
      this.showAddModal = true
      this.deviceId = row.deviceId
      this.locId = row.locId
      this.typeId = row.typeId
      this.author = row.loc.leader
      this.alarmEmail = row.loc.email
      this.locName = row.loc.locName
      this.typeName = row.type.typeName
      this.ifHaveAddTemplate = 0
      this.getTemplateData()
    },
    //批量新增任务框
    showAddModalp(row) {
      console.log(row)
      this.showAddModal2 = true
      this.deviceId = row.deviceId
      this.locId = row.locId
      this.typeId = row.typeId
      this.author = row.loc.leader
      this.alarmEmail = row.loc.email
      this.locName = row.loc.locName
      this.typeName = row.type.typeName
      this.ifHaveAddTemplate = 0
    },
    showEditModal1(row) {
      this.showEditModal = true
      this.author = row.author
      this.jobDesc = row.jobDesc
      this.alarmEmail = row.alarmEmail
      this.selectedScheduleType = row.scheduleType
      this.schedule_conf_CRON = row.schedule_conf_CRON
      this.schedule_conf_FIX_RATE = row.schedule_conf_FIX_RATE
      this.executorHandler = row.executorHandler
      this.executorParam = row.executorParam
      this.executorTimeout = row.executorTimeout
      this.executorFailRetryCount = row.executorFailRetryCount
      this.jobId = row.id
      console.log(row)
      console.log(this.jobId)
    },
 //添加新模板
    addNewTemplate() {
      // 弹框要求输入新模板名称
      const newTemplateName = prompt('请输入新任务模板的名称：', '');
      // 检查用户是否输入了模板名称
      if (newTemplateName === null) {
        // 用户点击了取消或关闭了对话框
        alert('模板名称不能为空，操作已取消。');
        return;
      }
      // 假设新增模板需要以下参数
      const newTemplateParams = {
        typeId: this.typeId,
        templateName: newTemplateName,
        jobDesc: this.jobDesc,
        scheduleType: this.selectedScheduleType,
        scheduleConf: this.schedule_conf_CRON,
        conGenDisplay: this.schedule_conf_CRON,
        scheduleConfCron: this.schedule_conf_CRON,
        scheduleConfFixRate: this.schedule_conf_FIX_RATE,
        scheduleConfFixDelay: "",
        glueType: "BEAN",
        executorHandler: this.executorHandler,
        executorParam: this.executorParam,
        executorTimeout: parseInt(this.executorTimeout, 10),
        executorFailRetryCount: parseInt(this.executorFailRetryCount, 10)
      };
      console.log(newTemplateParams)

      return fetch('http://120.79.59.158:8000/api/v1/jobTemplate', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${getToken()}`
        },
        body: JSON.stringify(newTemplateParams),
      })
        .then(response => response.json())
        .then(data => {
          if (data.code === 200) {
            // 成功逻辑
            return data;
          } else {
            throw new Error(data.message);
          }
        })
        .catch(error => {
          throw error;
        });
    },
    //新增任务
    saveFormData() {
      // 判断是否选择了自定义
      if (this.isCustomTemplate && this.ifHaveAddTemplate === 0) {
        this.ifHaveAddTemplate = 1
        this.addNewTemplate().then(() => {

        }).catch(error => {
          console.error('Error adding new template:', error);
          // 错误处理，如弹出错误消息等
        });
      }
      this.showAddModal = false
      // 使用 Vue 的 this 访问组件内的数据
      const params = {
        jobDesc: this.jobDesc,
        author: this.author,
        alarmEmail: this.alarmEmail,
        scheduleType: this.selectedScheduleType,
        scheduleConf: this.schedule_conf_CRON,
        cronGen_display: this.schedule_conf_CRON,
        schedule_conf_CRON: this.schedule_conf_CRON,
        schedule_conf_FIX_RATE: this.schedule_conf_FIX_RATE,
        schedule_conf_FIX_DELAY: "",
        glueType: "BEAN",
        executorHandler: this.executorHandler,
        executorParam: this.executorParam,
        executorTimeout: this.executorTimeout,
        executorFailRetryCount: this.executorFailRetryCount
      };
      console.log(params)

      // 使用 URLSearchParams 构建表单数据
      const formData = new URLSearchParams();
      for (let key in params) {
        if (params.hasOwnProperty(key)) {
          formData.append(key, params[key]);
        }
      }

      // 使用 fetch 发送异步请求
      fetch('http://120.79.59.158:8080/xxl-job-admin/jobinfo/add', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded', // 修改头部为表单类型
        },
        body: formData, // 使用构建好的表单数据
      })
        .then(response => {
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }
          return response.json();
        })
        .then(data => {
          console.log('Success:', data);
          const jobId = parseInt(data.content, 10); // 确保 jobId 是一个整数
          const token = getToken()
          // 处理返回的数据
          const additionalParams = {
            deviceId: this.deviceId,
            locId: this.locId,
            typeId: 1,
            jobId: jobId, // 将返回的 content 添加到 additionalParams
            token: token
          };
          // 将 additionalParams 转换为 JSON 字符串
          const jsonBody = JSON.stringify(additionalParams);
          // 使用 fetch 发送第二个异步请求，传递 JSON 格式的参数
          return fetch('http://120.79.59.158:8000/api/v1/jobinfo', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json', // 指定内容类型为 JSON
              'Authorization': `Bearer ${token}` // 在请求头中添加令牌
            },
            body: jsonBody, // 使用 JSON 字符串作为请求体
          }).then(response => {
            if (!response.ok) {
              throw new Error('Second network response was not ok');
            }
            return response.json();
          })
            .then(secondData => {
              console.log('Second Success:', secondData);
              // 处理第二个接口调用成功后的数据
            });
        })
        .catch(error => {
          console.error('Error:', error);
          // 处理错误情况，如弹出错误消息等
        })
        .catch(error => {
          console.error('Error:', error);
          // 处理错误情况，如弹出错误消息等
        });
    },
    batchSave() {
      this.ids.forEach(deviceId => {
        this.deviceId = deviceId; // 设置当前的 deviceId
        this.saveFormData(); // 调用保存数据的方法
      });
    },
    //修改任务
    saveFormData2() {
      // 使用 Vue 的 this 访问组件内的数据
      const params = {
        jobDesc: this.jobDesc,
        author: this.author,
        alarmEmail: this.alarmEmail,
        scheduleType: this.selectedScheduleType,
        scheduleConf: this.schedule_conf_CRON,
        cronGen_display: this.schedule_conf_CRON,
        schedule_conf_CRON: this.schedule_conf_CRON,
        schedule_conf_FIX_RATE: this.schedule_conf_FIX_RATE,
        schedule_conf_FIX_DELAY: "",
        glueType: "BEAN",
        executorHandler: this.executorHandler,
        executorParam: this.executorParam,
        executorTimeout: this.executorTimeout,
        executorFailRetryCount: this.executorFailRetryCount,
        id: this.jobId
      };
      console.log(params);

      // 调用接口，传递 params 参数

      // 使用 URLSearchParams 构建表单数据
      const formData = new URLSearchParams();
      for (let key in params) {
        if (params.hasOwnProperty(key)) {
          formData.append(key, params[key]);
        }
      }

      // 使用 fetch 发送异步请求
      fetch('http://120.79.59.158:8080/xxl-job-admin/jobinfo/update', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded', // 修改头部为表单类型
        },
        body: formData, // 使用构建好的表单数据
      })
        .then(response => {
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }
          return response.json();
        })
        .then(data => {
          console.log('Success:', data);
          this.getTask(this.deviceId)
          this.showEditModal = false
          // 处理返回的数据，例如 jobId
        })
        .catch(error => {
          console.error('Error:', error);
          // 处理错误情况，如弹出错误消息等
        });
    },
    //任务管理
    getTask(deviceId) {
      // console.log(deviceId);
      if (!deviceId) return; // 如果没有设备ID，不执行请求

      const url = `http://120.79.59.158:8000/api/v1/jobinfo/pageList?deviceId=${encodeURIComponent(deviceId)}`;

      fetch(url, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${getToken()}` // 使用Bearer Token进行认证
        }
      })
        .then(response => {
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }
          return response.json();
        })
        .then(data => {
          if (data && data.code === 200) { // 确保数据存在并且请求成功
            console.log('First Success:', data.data.list);
            const jobIds = data.data.list.map(item => item.jobId); // 收集所有的jobId
            console.log(jobIds);
            const secondApiUrl = `http://120.79.59.158:8080/xxl-job-admin/jobinfo/pageLists?ids=${encodeURIComponent(jobIds.join(','))}&triggerStatus=0`;
            return this.getSecondData(secondApiUrl); // 调用第二个接口的方法
          } else {
            console.error('No tasks data received or request failed');
          }
        })
        .then(secondData => {
          if (secondData && secondData.code === 200) { // 处理第二个接口返回的数据
            this.displaySecondData(secondData.content.data); // 显示数据的方法
          } else {
            console.error('Second API request failed or no data received');
          }
        })
        .catch(error => {
          console.error('Error:', error);
        });
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
    deleteTask(row) {
      console.log(row)
      // 删除第一个任务
      return fetch(`http://120.79.59.158:8080/xxl-job-admin/jobinfo/remove?id=${row.id}`, {
        method: 'POST'
      }).then(response => {
        if (!response.ok) {
          throw new Error('Failed to delete the first task')
        }

        const jobInfo = {
          ids: [row.id]
        }
        console.log(jobInfo)
        // 将jobInfo转换为JSON字符串
        const body = JSON.stringify(jobInfo)

        // 删除第二个任务
        return fetch(`http://120.79.59.158:8000/api/v1/jobinfo`, {
          method: 'DELETE',
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
            'Authorization': `Bearer ${getToken()}` // 使用Bearer Token进行认证
          },
          body: body
        })
      }).then(response => {
        if (!response.ok) {
          throw new Error('Failed to delete the second task')
        }
        console.log('Second task deleted successfully')
        // 处理删除后的操作，例如更新UI或状态
        this.getTask(this.deviceId)
      }).catch(error => {
        console.error('Error:', error)
        // 处理错误情况
      })
    },


    /**控制详情页弹出 */
    toggleSidebar(row) {
      const deviceId = row.deviceId || this.ids
      this.deviceId = deviceId
      this.currentTab = 'content'

      getDevice(deviceId).then(response => {
        this.deviceInfo = response.data
        console.log(this.deviceInfo.type.logo)
        // this.open = true
        const devicedata = this.deviceInfo.deviceData || '';
        // this.deviceDataSegments = devicedata.split(';').map(segment => segment.trim()).filter(segment => segment);
        this.deviceDataSegments = devicedata.split(',').map(segment => segment.trim()).filter(segment => segment);
        this.isSidebarVisible = true
        this.title = '设备详情'
      })
    },

    toggleSidebar1() {

      this.isSidebarVisible = !this.isSidebarVisible
      console.log(this.isSidebarVisible)

    },

    /**详情页物模型切换 */
    showContent(tab) {
      this.currentTab = tab
      if (tab === 'task' && this.deviceInfo.deviceId) {
        const deviceId = this.deviceInfo.deviceId
        this.getTask(deviceId)
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

    /**详情页相关 */
    convertCommasToBreaks(text) {
      if(text) {
       /*  const obj = JSON.parse(text);
        const s1 = '设备ID：' + (obj['DeviceID'] || 0) + '<br><br>'
        const s2 = '湿度：' + (obj['Humi'] || 0) + '<br><br>'
        const s3 = '温度：' + (obj['Temp'] || 0) + '<br><br>'
        const s4 = '其他：' + (obj['Led'] || 0)
        return s1 + s2 + s3 + s4; */
        return text.replace(/{/g, '').replace(/}/g, '').replace(/,/g, '<br><br>').replace(/"Humi"/g, '湿度：').replace(/"Temp"/g, '温度').replace(/"Led"/g, '其他')
      }
      return '-';
    },
    copyToClipboard(text) {
      if (!navigator.clipboard) {
        console.error('Clipboard API not available')
        return
      }
      navigator.clipboard.writeText(text).then(() => {
        console.log('Text copied to clipboard')
        alert("已复制到粘贴板")
      }).catch((error) => {
        console.error('Could not copy text: ', error)
      })
    },



  /** 查询用户列表 */
  getList() {
      this.loading = true
      listDevice(this.queryParams).then(response => {
        const data = response.data.list; // 获取原始设备列表数据
        const pageIndex = this.queryParams.pageIndex || 1; // 当前页码
        const pageSize = this.queryParams.pageSize || 10; // 每页显示数量
        const updatedList = data.map((item, index) => {
          return {
            ...item, // 保留原有数据
            编号: (pageIndex - 1) * pageSize + index + 1, // 计算编号
          };
        });
        this.DeviceList = updatedList
        // console.log(this.DeviceList)
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
    // 筛选节点
    filterNode(value, data) {
      if (!value) return true
      return data.label.indexOf(value) !== -1
    },
    // 节点单击事件
    handleNodeClick(data) {
      this.queryParams.locId = data.id
      this.loading = true
      getDeviceIdByLocId(this.queryParams).then(res => {
        const data = res.data;
        if(!data) {
          this.$message.warning('无数据')
          this.loading = false
          return false;
        }
        getlistByDeviceIds(JSON.stringify({deviceIds: data})).then(response => {
          const data = response.data.list;
          const pageIndex = this.queryParams.pageIndex || 1;
          const pageSize = this.queryParams.pageSize || 10;
          const updatedList = data.map((item, index) => {
            return {
              ...item,
              编号: (pageIndex - 1) * pageSize + index + 1,
            };
          });
          this.DeviceList = updatedList
          this.total = response.data.count
          this.loading = false
        })
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
    //添加设备取消按钮
    canceladd() {
      this.addDevice = false
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
    /** 新增按钮操作 */
    handleAddDevice() {
      this.deviceInfo = this.$options.data().deviceInfo
      // this.addDevice = true
      this.open = true
      this.title = '添加设备'
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      console.log('Update triggered', row)

      const deviceId = row.deviceId || this.ids
      getDevice(deviceId).then(response => {
        this.deviceInfo = response.data
        this.open = true
        this.title = '修改设备'
      })

      listRole({ pageSize: 1000 }).then(response => {
        this.roleOptions = response.data.list
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.deviceInfo.deviceId !== undefined) {
            updateDeviceProfile(this.deviceInfo).then(response => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addDevice(this.addInfoForm).then(response => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
                this.open = false
                this.addDevice = false
                this.getList()
                this.addInfoForm = this.$options.data().addInfoForm
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
      const Ids = (row.deviceId && [row.deviceId]) || this.ids
      this.$confirm('是否确认删除设备编号为"' + Ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.getList()
        return delDevice({ 'ids': Ids })
        this.$message.warning('删除设备接口调用')
      })

    },
    //上传图片相关
    handleUpload(file) {
    this.files.push(file);
    if (this.files.length === this.uploadLimit) {
      const formData = new FormData();
      this.files.forEach((fileItem, index) => {
        formData.append(`file${index + 1}`, fileItem.raw);
      });
      uploadFile(formData, this.uploadType)
        .then(response => {
          console.log('上传成功:', response);
        })
        .catch(error => {
          console.error('上传失败:', error);
        });
      this.files = [];
    }
  },
  setUploadType(type) {
    this.uploadType = type;
    this.files = [];
  },
  /** 导出按钮操作 */
  handleExport() {
    const queryParams = this.queryParams
    this.$confirm('是否确认导出所有用户数据项?', '警告', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(function() {
      return exportUser(queryParams)
    }).then(response => {
      this.download(response.msg)
    }).catch(function() { })
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
  Taskstart(row) {
    console.log(row)
    this.queryParams.ids = row.id
    JobStart(this.queryParams).then(response => {
      this.LogsAdd(this.queryParams)
      this.getList()
    })
  },
  // 增加日志RUN
  LogsAdd(row) {
    console.log(row.ids)
    const temp = { jobId: row.ids, name: 'admin', operator: '设备' + row.ids + '启动', tel: '123456' }
    logadd(temp).then(response => {
      this.GetLogsList()
    })
  },
  /* 执行一次任务 */
  Singletask(row) {
    this.reset()
    this.form.ids = String(row.id).trim()
    this.SingledialogVisible = true
  },
  reset() {
    this.form = {
      executorParam: undefined,
      ids: undefined
    }
    this.resetForm('form')
  },
  //上传图片相关
  httpRequest(deviceId, file) {
    console.log('filefile', file)
    const formData = new FormData()
    formData.append('file', file.file)
    formData.append('deviceId', deviceId)

    uploadFile(formData).then(res => {
        if (res.code === 200) {
          this.$refs.uploadImgSrc.clearFiles()
          getDevice(deviceId).then(response => {
            this.deviceInfo = response.data
            // imgSrc = ["xxx.jpg", "xxxx.jpg"]
          })
        }
      })
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

/* 侧边框的样式 */
.sidebar {
  position: fixed;
  right: 0;
  /* 初始时隐藏在视窗右侧 */
  top: 0;
  width: 50%;
  /* 侧边框的宽度 */
  height: 100%;
  /* background-color: white; */
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.3);
  transition: right 3s ease;
  z-index: 100;
  /* 确保侧边框在内容之上 */
  background-color: rgba(255, 255, 255, 0.8);
  /* 半透明的白色背景 */
  backdrop-filter: blur(10px);
  /* 添加模糊效果 */
}

/* 侧边框可见时的样式 */
.sidebar-visible {
  right: 0;
  /* 显示在视窗内 */
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
  /* margin-top: 15px; */
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
.is-link {
  margin-right: 5px;
  color: #0099FF;
  font-size: 12px;
  cursor: pointer;
}
</style>
