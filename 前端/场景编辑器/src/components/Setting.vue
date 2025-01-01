<script setup>
import { onMounted, ref } from "vue";
import MapProps from "./MapProps/Map.vue";
import Global from "./MapProps/Global.vue";
import Appearance from "./PenProps/Appearance.vue";
import Event from "./PenProps/Event.vue";
import Animate from "./PenProps/Animate.vue";
const props = defineProps(['activeType'])
let activeName1 = ref('map')
let activeName2 = ref('appearance')
let activeName3 = ref('appearance')
let activePen = ref(false)
let multiPen = ref(false)
import { defineEmits } from 'vue';
const emit = defineEmits(['update:active-data']);

const queryString = window.location.search;
const urlParams = new URLSearchParams(queryString);
const sceneId = urlParams.get('sceneId');
const token = urlParams.get('token');
const requestData = ref({
  sceneId: parseInt(sceneId, 10),
  elementId: ''
});
onMounted(() => {
  // 监听鼠标点击画布上的设备图元
  meta2d.on('active', (args) => {
    console.log('active', args[0].id)//打印所选择的设备信息
    requestData.value.elementId = args[0].id;
    // console.log('requestData', requestData.value)
    fetch(`http://120.79.59.158:8000/api/v1/sceneDevice/get`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(requestData.value)
    })
      .then(response => response.json())
      .then(data => {
        // console.log('data', data)
        const deviceId = parseInt(data.data.deviceId, 10)
        const secondApiUrl = `http://120.79.59.158:8000/api/v1/device/${encodeURIComponent(deviceId)}`;

        // 发起第二个请求
        fetch(secondApiUrl, {
          method: 'GET',
          headers: {
            'Authorization': `Bearer ${token}`
          }
        })
          .then(response => response.json()) // 解析第二个请求的响应为JSON
          .then(secondApiResponse => {
            // 处理第二个请求的响应数据
            // console.log('secondApiResponse', secondApiResponse);
            parseSecondApiResponse(secondApiResponse.data.deviceData, dataList);
          })
          .catch(secondApiError => {
            console.error('Error calling second API:', secondApiError);
          });
      })
      .catch(error => {
        console.error('Error:', error);
      });


    if (args.length >= 1) activePen.value = true
    if (args.length > 1) {
      multiPen.value = true
    } else {
      multiPen.value = false
    }
    if (args.length >= 1 && args[0] && typeof args[0].id !== 'undefined') {
      // 安全地发出自定义事件
      emit('update:active-data', args[0].id);
    }
  })
  meta2d.on('inactive', () => {
    activePen.value = false
    multiPen.value = false
  })
})
function parseSecondApiResponse(deviceDataString, dataListRef) {
  // 去除字符串前后的大括号
  const trimmedDataString = deviceDataString.replace(/^\{|\}$/g, '');
  // 将字符串按逗号分割，得到键值对数组
  const keyValuePairs = trimmedDataString.split(',');
  // console.log('keyValuePairs', keyValuePairs)

  // 创建一个空对象来存储解析后的键值对
  const deviceData = {};

  // 遍历键值对数组，并将它们添加到对象中
  keyValuePairs.forEach(pair => {
    // 去除键值对中的等号两边的空格和双引号
    const [key, value] = pair.split(':').map(item => item.trim().replace(/"/g, ''));
    deviceData[key] = value; // 假设所有的值都是数字
  });

  // 使用 deviceData 对象的属性创建 dataList 条目
  const dataListItems = Object.entries(deviceData).map(([key, value]) => {
    return {
      name: key, // 使用属性名作为条目的名称
      description: value.toString() // 使用属性值的字符串表示作为描述
    };
  });

  // 直接更新 dataList 为新的数据列表
  dataListRef.value = dataListItems;
}
// 操作列表数据
const operateList = ref([
  { name: '报警器' },
  { name: 'LED' },
]);
const dataList = ref([
  { name: '111', description: '222' },
  { name: '333', description: '444' },
])

// 操作方法
function switchChanged(rowData) {
  // console.log('开关状态改变:', rowData);
  // 检查开关状态
  if (rowData.isSwitchOn) {
    if (rowData.name === '报警器') {
      console.log('执行报警器开启操作');
      beepController(1);
    }
    else {
      console.log('执行LED开启操作');
      ledController(1);
    }
  } else {
    if (rowData.name === '报警器') {
      console.log('执行报警器关闭操作');
      beepController(0);
    }
    else {
      console.log('执行LED关闭操作');
      ledController(0);
    }
  }
}
function ledController(state) {
  const url = `http://120.79.59.158:8000/api/v1/device/controlLed?state=${encodeURIComponent(state)}`;
  fetch(url, {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${token}`
    }
  })
    .then(response => response.json())
    .then(data => console.log(data))
    .catch(error => console.error('Error:', error));
}

function beepController(state) {
  const url = `http://120.79.59.158:8000/api/v1/device/controlBeep?state=${encodeURIComponent(state)}`;
  fetch(url, {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${token}`
    }
  })
    .then(response => response.json())
    .then(data => console.log(data))
    .catch(error => console.error('Error:', error));
}
</script>

<template>

  <div class="setting">

    <div class="tz_props" v-show="!activePen">
      <el-tabs v-model="activeName1" class="demo-tabs" active-name="first">
        <el-tab-pane label="图纸" name="map" class="tab">
          <MapProps></MapProps>
        </el-tab-pane>
        <el-tab-pane label="全局配置" name="global" class="tab">
          <Global></Global>
        </el-tab-pane>
      </el-tabs>
    </div>

    <div class="ty_props" v-show="activePen && !multiPen">
      <el-tabs v-model="activeName2" class="demo-tabs">
        <el-tab-pane label="外观" name="appearance">
          <Animate></Animate>
          <Appearance></Appearance>
        </el-tab-pane>

        <el-tab-pane label="数据" name="data">
          <el-table :data="dataList">
            <el-table-column prop="name" label="数据名称"></el-table-column>
            <el-table-column prop="description" label="描述"></el-table-column>
          </el-table>
        </el-tab-pane>

        <el-tab-pane label="操作" name="interface">
          <!-- 操作列表 -->
          <el-table :data="operateList">
            <el-table-column prop="name" label="名称"></el-table-column>
            <el-table-column label="操作">
              <template #default="scope">
                <el-switch v-model="scope.row.isSwitchOn" @change="switchChanged(scope.row)">
                </el-switch>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>

        <el-tab-pane label="事件" name="event">
          <Event></Event>
        </el-tab-pane>
      </el-tabs>
    </div>

    <div class="ty_props" v-show="activePen && multiPen">
      <el-tabs v-model="activeName3" class="demo-tabs">
        <el-tab-pane label="外观" name="appearance">
          <Appearance></Appearance>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<style scoped>
.setting {
  position: relative;
  display: flex;
  width: 350px;
  padding: 16px 0 0 16px;
  overflow: auto;
  box-shadow: 0 2px 4px 0 #dad7d7;
}

:deep(.el-tabs__header) {
  margin: 0;
}

:deep(.el-tabs__content::-webkit-scrollbar) {
  /*滚动条整体样式*/
  width: 10px;
  /*高宽分别对应横竖滚动条的尺寸*/
  height: 1px;
}

:deep(.el-tabs__content::-webkit-scrollbar-thumb) {
  /*滚动条里面小方块*/
  border-radius: 10px;
  height: 20px;
  -webkit-box-shadow: inset 0 0 5px rgba(0, 0, 0, 0.2);
  background: #4e4e4e;
}

:deep(.el-tabs__content::-webkit-scrollbar-track) {
  /*滚动条里面轨道*/
  -webkit-box-shadow: inset 0 0 5px rgba(0, 0, 0, 0.2);
  border-radius: 10px;
  background: #ffffff;
}

.tz_props {
  display: flex;
  width: 100%;
}

.ty_props {
  display: flex;
  width: 100%;
}

:deep(.el-tabs__header) {
  height: 40px;
}

:deep(.el-tabs__content) {
  flex: 1;
  overflow: auto;
}

:deep(.el-tabs) {
  display: flex;
  flex-direction: column;
  width: 100%;
}
</style>