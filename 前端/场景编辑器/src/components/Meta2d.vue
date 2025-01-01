<script setup>
import { Meta2d } from "@meta2d/core"
import { flowPens } from "@meta2d/flow-diagram";
import {
  activityDiagram,
  activityDiagramByCtx,
} from "@meta2d/activity-diagram";
import { sequencePens, sequencePensbyCtx } from "@meta2d/sequence-diagram";
import { classPens } from "@meta2d/class-diagram";
import { collapseChildPlugin } from "@meta2d/plugin-mind-collapse"
import { mindBoxPlugin } from "@meta2d/plugin-mind-core"
// import { myTriangle, myTriangleAnchors} from "../../public/path2D/mypath2d/myTriangle.js";
import { register as registerEcharts, registerHighcharts, registerLightningChart } from "@meta2d/chart-diagram"; // 引入echarts注册函数，原函数名为register 为了与其他注册函数区分这里重命名为registerEcharts
import { formPens } from '@meta2d/form-diagram';
import { onMounted } from "vue";
// import { canvasTriangle, canvasTriangleAnchors } from "../../public/canvasDraw/myCanvasDraw/canvasTriangle.js";
import { useEventbus } from "../hooks/useEventbus.js";
import Contextmenu from "./Contextmenu.vue";

import mapModule from '../data/defaultsConfig';
const event = useEventbus()


import sharedState from './Main.vue';
import deviceInfo from './Main.vue';

onMounted(async () => {

  // 获取当前页面的 URL 查询字符串部分
  const queryString = window.location.search;

  // 创建 URLSearchParams 对象
  const urlParams = new URLSearchParams(queryString);

  // 从查询字符串中获取 image 和 token 参数
  const sceneId = urlParams.get('sceneId');
  const token = urlParams.get('token');
  // 更新共享状态
  sharedState.value = { sceneId, token };

  // 现在你可以使用 image 和 token 变量了
  console.log(sceneId, token);
  // 一旦获取了 sceneId 和 token，立即执行这个函数
  fetchDataAndProcess(sceneId, token);





  // 创建meta2d对象
  let meta2d = new Meta2d("meta2d")
  meta2d.register(flowPens())

  // 注册注册活动图元
  meta2d.register(activityDiagram())
  // 原生canvas绘画的图库，支持逻辑复杂的需求
  meta2d.registerCanvasDraw(activityDiagramByCtx())

  // 注册时序图
  meta2d.register(sequencePens())
  meta2d.registerCanvasDraw(sequencePensbyCtx())
  meta2d.installPenPlugins({ name: 'mindNode2' }, [{
    plugin: mindBoxPlugin
  }, {
    plugin: collapseChildPlugin
  }])

  // 注册类图
  meta2d.register(classPens())
  // uninstallPlugin('toolBox')
  // 注册表单图元
  meta2d.registerCanvasDraw(formPens())
  // 直接调用Echarts的注册函数
  registerEcharts()

  // 直接调用HighCharts的注册函数
  registerHighcharts()
  // 直接调用LightningChart的注册函数
  registerLightningChart()

  // installPlugin(CollapseChildPlugin)

  // a.setFuncList({
  //   'root':[
  //     {
  //       name: "测试",
  //       event: 'click',
  //       func(self,pen){
  //         console.log(self,pen)
  //       },
  //
  //       setDom(self,dom){
  //         let html = `<div>
  //               测试按钮
  //           </div>`
  //         let css = `<style>
  //               div {
  //               color: pink;
  //               }
  //           </style>`
  //         return html + css
  //       }
  //     }
  //   ],
  //   'leaf':[{
  //     name: "测试2"
  //   }
  //   ]
  // })
  //注册自定义path2d图元
  // meta2d.register({ myTriangle })
  // 注册自定义图元的m锚点信息
  // meta2d.registerAnchors({ myTriangle: myTriangleAnchors });

  // 注册自定义canvasDraw函数
  // meta2d.registerCanvasDraw({ canvasTriangle })
  //注册锚点
  // meta2d.registerAnchors({ canvasTriangle: canvasTriangleAnchors })
  event.customEmit('opened')
  event.customEmit('load')


  // // 获取插件
  // let a = await getPlugin('toolBox')
  // // a.setFuncList({'A':[{
  // //     name:"颜色"
  // //   },],'B':[{
  // //   name:'666'
  // //   }]})
  // // a.getFuncList = (pen)=>{
  // //   return a.funcList['A']
  // // };
  // a.childrenGap = 100
})


// 定义一个函数，根据 sceneId 和 token 获取数据并处理
async function fetchDataAndProcess(sceneId, token) {
  const response = await fetchData(sceneId, token);
  try {
    // 检查响应码和消息
    if (response.code === 200 && response.msg === '查询成功' && response.data) {
      // 现在 data 是实际的绘图数据
      const data = response.data;
      deviceInfo.value.locId = data.locId;
      deviceInfo.value.locName = data.locName;
      const json = JSON.parse(data.data);
      // console.log('Processing data:', data);
      if (json && json.pens) {
        mapModule.initializeMap(new Map(json.pens.map(device => [device.id, device])));
        window.meta2d.open(json); // 确保 window.meta2d.open 能够处理传入的 data 对象
        useEventbus().customEmit('opened');
      } else {
        console.error('Invalid data structure within response.data:', data);
      }
    } else {
      console.error('Unexpected response format or error in response:', response);
    }
  } catch (error) {
    console.error('Error processing fetched data:', error);
  }
}



// 修改 fetchData 函数以调用接口
async function fetchData(sceneId, token) {
  // 构建请求选项
  const options = {
    method: 'GET',
    headers: {
      // 添加请求头，传递 token
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json'
    }
  };

  // 构造请求 URL，包含 sceneId 和 token
  const url = `http://120.79.59.158:8000/api/v1/scene/${sceneId}?token=${token}`;

  try {
    // 发起 fetch 请求
    const response = await fetch(url, options);
    if (!response.ok) {
      // 如果响应状态码不是 2xx，抛出错误
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    // 解析响应体为 JSON
    const data = await response.json();
    return data; // 返回获取的数据
  } catch (error) {
    console.error('Fetching data failed:', error);
    // 处理错误情况
    throw error;
  }
}
</script>

<template>
  <!--  meta2d图纸所站位置-->
  <div id="meta2d">
  </div>
  <contextmenu></contextmenu>

</template>

<style scoped>
#meta2d {
  position: relative;
  height: 100%;
  width: 100%;
  overflow: hidden;
  background-color: #eee;
}
</style>