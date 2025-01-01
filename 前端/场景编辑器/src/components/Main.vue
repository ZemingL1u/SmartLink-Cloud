<script setup>
import Icons from "./Icons.vue";
import Meta2d from "./Meta2d.vue";
import Setting from "./Setting.vue";
// import Treeselect from 'vue3-treeselect'
// import 'vue3-treeselect/dist/vue3-treeselect.css'
import { ref } from 'vue';
// 定义一个方法来处理从 setting 组件传来的数据
const handleActiveData = (data) => {
  console.log('Received data from setting component:', data);
  activeData.value = data;
};
// 定义一个 ref 来存储从 setting 组件传来的数据
const activeData = ref(null);


// 获取当前页面的 URL 查询字符串部分
const queryString = window.location.search;

// 创建 URLSearchParams 对象
const urlParams = new URLSearchParams(queryString);

// 从查询字符串中获取 image 和 token 参数
const sceneId = urlParams.get('sceneId');
const token = urlParams.get('token');
const state = urlParams.get('state');
// 设备类型选项列表
const deviceTypeOptions = ref([
  { value: '6', label: '室内网关' },
  { value: '7', label: '室内环境传感器' },
  { value: '8', label: '噪声传感器' },
  { value: '9', label: '数传终端' },
  { value: '10', label: '电流互感器' },
  { value: '11', label: '智能插座面板' },
  { value: '12', label: '智能开关面板' },
  { value: '13', label: '光照/PIR传感器' },
  { value: '14', label: '空间人数传感器' },
  { value: '15', label: '风机盘管温控器' },
  { value: '16', label: '小型数传终端' },
  { value: '17', label: '情景面板' },
  { value: '18', label: '无线门磁' },
  { value: '19', label: '独立式光电感烟火灾探测报警器' },
  { value: '20', label: 'AI ToF人数统计传感器' }
]);

// 响应式对象
const deviceInfo = ref({
  deviceName: '',
  locId: 12,
  locName: "致理楼L1",
  status: "在线",
  typeId: '',
  sNCode: '',
  key: '',
  updatedAt: '',
  remark: ''
});

// 控制弹窗显示的变量
const open = ref(false);
// 弹窗标题
const title = ref('添加设备');
// 提交表单的方法
const submitForm = async () => {
  console.log("触发提交函数");
  try {
    // 构建请求体
    const requestBody = {
      ...deviceInfo.value,
      typeId: parseInt(deviceInfo.value.typeId, 10) // 转换 typeId 为整数
    };
    console.log(requestBody);
    console.log(token)

    // 发起 API 请求
    const response = await fetch(`http://120.79.59.158:8000/api/v1/device`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`, // 使用从 sharedState 获取的 token
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(requestBody) // 将请求体转换为 JSON 字符串
    });

    if (!response.ok) {
      // 如果响应状态码不是 2xx，抛出错误
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    // 解析响应体为 JSON
    const data = await response.json();
    console.log('接口返回的数据：', data);

    // 提交成功后关闭弹窗
    open.value = false;

    // 根据返回的数据给用户反馈
    if (data) {
      const secondRequestData = {
        deviceId: parseInt(data.data, 10),
        sceneId: parseInt(sceneId, 10),
        elementId: activeData.value,
      };
      console.log(secondRequestData);
      const secondResponse = await fetch(`http://120.79.59.158:8000/api/v1/sceneDevice`, {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(secondRequestData)
      });

      if (!secondResponse.ok) {
        throw new Error(`HTTP error! status: ${secondResponse.status}`);
      }

      // 解析第二个请求的响应体为 JSON
      const secondData = await secondResponse.json();
      console.log('第二个接口返回的数据：', secondData);
    } else {
      console.error('操作失败');
    }
  } catch (error) {
    console.error('表单提交失败：', error);
  }
};
// 取消操作的方法
const cancel = () => {
  open.value = false;
};
function updateOpen() {
  // 设置一个两秒的定时器
  setTimeout(() => {
    // 两秒后执行，修改 open 的值为 true
    open.value = true;
  }, 1000); // 2000 表示两秒，单位是毫秒
}
</script>

<template>
  <div class="main">
    <Icons v-if="state == 1" @openChanged="updateOpen"></Icons>
    <Meta2d></Meta2d>
    <el-dialog :title="title" v-model="open" width="600px" :close-on-click-modal="false">
      <el-form ref="form" :model="deviceInfo" label-width="80px" @close="open = false">
        <el-row>
          <el-col :span="12">
            <el-form-item label="设备名称" prop="deviceName">
              <el-input v-model="deviceInfo.deviceName" placeholder="请输入设备名称" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="归属区域" prop="locId">
              <!-- <treeselect v-model="deviceInfo.locId" :options="locOptions" placeholder="请选择归属区域" />-->
              <el-input placeholder="致理楼L1" disabled />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="设备类型" prop="deviceType">
              <el-select v-model="deviceInfo.typeId" placeholder="请选择设备类型">
                <el-option v-for="item in deviceTypeOptions" :key="item.value" :label="item.label"
                  :value="item.value"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="序列号" prop="sNCode">
              <el-input v-model="deviceInfo.sNCode" placeholder="请输入序列号" maxlength="50" />
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="设备秘钥" prop="key">
              <el-input v-model="deviceInfo.key" type="textarea" />
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="备注">
              <el-input v-model="deviceInfo.remark" type="textarea" placeholder="请输入内容" />
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <div class="dialog-footer">
        <el-button type="primary" @click="submitForm">确 定</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>
    <Setting @update:active-data="handleActiveData"></Setting>
  </div>
</template>

<style scoped>
.main {
  flex: 1;
  height: calc(100% - 59px);
  display: flex;
  justify-content: flex-start;
}
</style>