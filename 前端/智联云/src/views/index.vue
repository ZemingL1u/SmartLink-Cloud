<template>
  <div class="app-container">
    <!-- 点击“详情”文字后弹出侧边框 -->
    <div>
      <button @click="toggleSidebar">详情</button>
    </div>
    <!-- 侧边框组件，初始时隐藏在视窗外，使用v-show控制显示 -->
    <div class="sidebar" v-show="isSidebarVisible">
      <div class="back-wrapper" @click="toggleSidebar">
        <i class="fas fa-arrow-left"></i> 返回
      </div>
      <div id="options">
        <!-- "详情" 选项卡 -->
        <button class="tab-button" :class="{ active: currentTab === 'content' }"
          @click="showContent('content')">详情</button>
        <!-- "物模型数据" 选项卡 -->
        <button class="tab-button" :class="{ active: currentTab === 'data' }"
          @click="showContent('data')">物模型数据</button>
      </div>
      <div class="content" v-show="currentTab === 'content'">
        <h4>设备图片</h4>
        <div class="device-image" v-if="deviceInfo && deviceInfo.type">
          <img :src="deviceInfo.type.logo">
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
              <tr>
                <td>二维码</td>
                <td></td>
              </tr>
            </tbody>
          </table>
        </div>
        <div class="button-group">
          <button @click="copyToClipboard(deviceInfo.deviceId)">复制设备ID</button>
          <button @click="copyToClipboard(deviceInfo.key)">复制设备秘钥</button>
        </div>
        <h4>设备位置</h4>
        <div id="mapContainer"></div>
      </div>
      <div id="data" v-show="currentTab === 'data'" style="display: none;">
        <!-- 使用 v-for 遍历 devicedata数组 -->
        <div class="card" v-for="(   dataSegment, index   ) in    deviceDataSegments   " :key="index">
          <!-- 使用 v-html 渲染带有换行的 HTML 内容 -->
          <div class="mask" v-html="convertCommasToBreaks(dataSegment)"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
//import AMapLoader from '@amap/amap-jsapi-loader';

// 用于存放地图实例
let map = null;
// 地图容器的 ID
const containerId = 'mapContainer';
// 地图中心点坐标
const center = [113.940052, 22.528612];
// 地图初始缩放级别
const zoom = 11;


// 控制侧边框显示状态的响应式变量
const isSidebarVisible = ref(false);
// 切换侧边框显示的方法
function toggleSidebar() {
  isSidebarVisible.value = !isSidebarVisible.value;
}

// 控制当前选中的选项卡
const currentTab = ref('content');
// 根据当前选中的选项卡显示对应的内容
const showContent = (tab) => {
  currentTab.value = tab;
};
//将逗号替换为换行
const convertCommasToBreaks = (text) => {
  return text.replace(/,/g, '<br><br>');
};

// 定义响应式数据来存储设备信息
const deviceInfo = ref({});
const devicedata = ref({});
const deviceDataSegments = ref([]);
const deviceID = '1'; // 你要查询的设备ID
const queryParams = { token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhc2NvcGUiOiIiLCJleHAiOjQ4NjY5NjU0OTUsImlkZW50aXR5IjoxLCJuaWNlIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTcxMzMyOTQ5NSwicm9sZWlkIjoxLCJyb2xla2V5IjoiYWRtaW4iLCJyb2xlbmFtZSI6Iuezu-e7n-euoeeQhuWRmCJ9.6KE1GFMEFuLllv6xxws0GP4XsJriRn-29yJKJpAzWLU'/* 参数对象，例如 { id: 1, name: 'John' } */ };
// 在组件挂载后发送请求获取数据
onMounted(async () => {
  try {
    const queryString = new URLSearchParams(queryParams).toString();
    const response = await fetch(` http://120.79.59.158:8000/api/v1/device/` + deviceID + `?${queryString}`);
    // const response = await fetch(` http://120.79.59.158:8000/api/v1/device/` + deviceID + `?${queryString}`);
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json();
    // console.log(data)
    deviceInfo.value = data.data;
    devicedata.value = data.data.deviceData;
    console.log(deviceInfo.value)
    // console.log(devicedata.value)
    // 使用 split 来分割字符串，然后使用 filter 移除空字符串
    deviceDataSegments.value = devicedata.value.split(';').map(segment => segment.trim()).filter(segment => segment);
  } catch (error) {
    console.error("Error fetching device data:", error);
  }

  AMapLoader.load({
    key: '34a27c980fd0fbda62d731e769c25e01', // 请替换成APIKey
    version: '2.0', // 高德地图 API 版本号
    plugins: ['AMap.Geocoder'], // 此处可指定需要使用的插件，如比例尺'AMap.Scale'
  })
    .then((AMap) => {
      // 初始化地图
      map = new AMap.Map(containerId, {
        center: center,
        zoom: zoom,
        // viewMode: '3D', // 使用3D视图
      });

      // 地图加载完成后，执行定位
      console.log(deviceInfo.value.loc.locName)
      // locateDevice(AMap, deviceInfo.value.loc.locName);
      // 创建一个标记并添加到地图上
      const marker = new AMap.Marker({
        position: center, // 标记的位置
        title: "设备位置", // 鼠标悬停时显示的标题
      });

      // // 将标记添加到地图上
      map.add(marker); // 地图实例的 add 方法接受 Marker 实例作为参数
    })

});
//不要删！！！！！！（后面可能会用）
// async function locateDevice(AMap, locationName) {
//   // 创建Geocoder地理编码类实例，进行地理编码操作
//   const geocoder = new AMap.Geocoder({
//     city: "深圳", // 指定城市，不指定表示全国
//     extensions: "all", // 返回地址组成要素详情
//   });

//   // 使用geocoder进行地理编码
//   try {
//     const status = await geocoder.getLocation(locationName);
//     if (status === 'complete' && geocoder.getCityList().length) {
//       // 地理编码成功，设置中心点并添加标记
//       const center = geocoder.getCityList()[0].center;
//       map.setCenter(center);
//       const marker = new AMap.Marker({
//         position: center, // 标记的位置
//         title: "设备位置", // 鼠标悬停时显示的标题
//       });
//       map.add(marker); // 地图实例的 add 方法接受 Marker 实例作为参数
//     } else {
//       console.log('地理编码失败');
//     }
//   } catch (error) {
//     console.error('地理编码出错:', error);
//   }
// }

// 定义复制到剪贴板的函数
const copyToClipboard = (text) => {
  if (!navigator.clipboard) {
    console.error('Clipboard API not available');
    return;
  }

  navigator.clipboard.writeText(text).then(() => {
    console.log('Text copied to clipboard');
    // 可以在这里显示一个提示，告知用户文本已复制
    alert("已复制到粘贴板")
  }).catch((error) => {
    console.error('Could not copy text: ', error);
  });
};


</script>

<style>
/* //@import '@fortawesome/fontawesome-free/css/all.css'; */

body {
  background-color: rgb(0, 115, 255);
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
  width: 51px;
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

#mapContainer {
  width: 100%;
  height: 400px;
}
</style>