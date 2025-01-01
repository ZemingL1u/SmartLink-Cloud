<template>
  <div  id="special-page">
    <!-- 头部 -->
    <header>
      <h1>{{this.Data.locName}}</h1>
      <div class="show-time">{{ currentTime }}</div>
    </header>

    <!-- 页面主体 -->
    <section class="mainbox">
      <!-- 左侧盒子 -->
      <div class="column">
        <div class="panel bar">
          <h2>设备状态统计</h2>
          <!-- 图表放置盒子 -->
          <div data-v-69b8304a="" data-v-a361ec26="" id="Round-chart">
            <div data-v-69b8304a="" class="Round-wrap">
              <div data-v-69b8304a="" class="Round" style="border-color: rgb(190, 188, 137);">
                <p data-v-69b8304a="">16</p>
              </div>
              <p data-v-69b8304a="" class="status">在线</p>
            </div>

            <div data-v-69b8304a="" class="Round-wrap">
              <div data-v-69b8304a="" class="Round" style="border-color: rgb(92, 44, 101);">
                <p data-v-69b8304a="">1</p>
              </div><p data-v-69b8304a="" class="status">待机</p>
            </div>

            <div data-v-69b8304a="" class="Round-wrap">
              <div data-v-69b8304a="" class="Round" style="border-color: rgb(132, 167, 98);">
                <p data-v-69b8304a="">1</p>
              </div>
              <p data-v-69b8304a="" class="status">待激活</p>
            </div>
            <div data-v-69b8304a="" class="Round-wrap">
              <div data-v-69b8304a="" class="Round" style="border-color: rgb(128, 90, 24);">
                <p data-v-69b8304a="">2</p>
              </div>
            <p data-v-69b8304a="" class="status">故障</p>
            </div>
          </div>
          <!-- 伪元素绘制盒子下边角 -->
          <div class="panel-footer"></div>
        </div>
        <div class="panel line">
          <h2>数量统计</h2>
          <div class="chart">
            <div class="info-item">
              <img src="../../assets/work1.png" alt="设备数" >
              <p>设备数：{{ total }}</p>
            </div>
            <div class="info-item">
              <img src="../../assets/work2.png" alt="类型数" >
              <p>类型数：4</p>
            </div>
          </div>
          <div class="panel-footer"></div>
        </div>
        <div class="panel pie">
          <h2>设备类型统计</h2>
          <div class="chart"></div>
          <div class="panel-footer"></div>
        </div>
      </div>



      <!-- 中间盒子 -->
      <div class="column">

        <!-- 头部 no模块 -->
        <div class="no">
          <div class="no-hd">
            <ul>
              <li>{{this.Data.sceneName}}</li>
            </ul>
          </div>
          <div class="no-bd">
            <ul>
              <li>场景名称</li>
            </ul>
          </div>
        </div>
        <!-- map模块 -->
        <div class="map">
          <div class="map1"></div>
          <div class="map2"></div>
          <div class="map3"></div>
          <div class="chart"></div>
        </div>
      </div>


      <!-- 右侧盒子 -->
      <div class="column">
        <div class="panel">
          <h2>场景缩略图</h2>
          <div class="center-mysrc">
            <img style="width: 300px;height: 150px;margin: 0 auto" :src="`data:image/png;base64,${Data.image}`" alt="">
          </div>
        </div>
        <div class="panel">
          <h2>设备实际数据</h2>
          <div class="chart">
            <div class="infotwo-item">
              <img src="../../assets/work4.png" alt="智能电表" >
                <div class="text-container">
                  <p>A相电流：0.055安培/A</p>
                  <p>A相电压：234.5伏特 / V</p>
                  <p>有功功率：8.4瓦特 / W</p>
                  <p>有功总电能：0.12千瓦时 / kw*h</p>
                </div>
            </div>
            <div class="infotwo-item">
              <img src="../../assets/work3.png" alt="温湿度传感器" >
                <div class="text-container">
                  <p>温度：30℃</p>
                  <p>湿度：74%</p>
                </div>
            </div>
          </div>
          <div class="panel-footer"></div>
        </div>
        <div class="panel linepart2">
          <h2>温度与湿度的历史与预测
            <a class="a-active" href="javascript:">日</a>
            <a href="javascript:">时</a>
            <a href="javascript:">月</a></h2>
          <div class="chart"></div>
          <div class="panel-footer"></div>
        </div>
      </div>
    </section>
  </div>
</template>




<script>
import "../../layout/ECharts-CSS/flexible.js"
import "../../layout/ECharts-CSS/echarts.min.js"
import "../../layout/ECharts-CSS/jquery.js"

import echarts from 'echarts'
import {listDevice} from "@/api/sys-device";

import $ from 'jquery'
import {getHumiData, getTempData} from "@/api/LargeScreen";

export default {
  name: 'LargeScreen',
  data(){
    return {
      currentTime:'',

      Data:'',

      //设备数
      total: '',

      //温度数据：
      TempData:'',
      //湿度数据：
      HumiData:'',

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
    }
  },
  mounted() {
    const DataString = this.$route.query.chartData
    if (DataString) {
      this.Data = JSON.parse(DataString);
      // 使用this.chartData初始化图表
      console.log(this.Data)
    }
    //
    this.getList()
    this.FindData()
    this.updateTime();
    this.timer = setInterval(this.updateTime, 1000);

  },
  created() {

  },
  methods: {
    //时间显示
    updateTime() {
      const dt = new Date();
      const y = dt.getFullYear();
      const mt = dt.getMonth() + 1;
      const day = dt.getDate();
      const h = dt.getHours();
      const m = dt.getMinutes();
      const s = dt.getSeconds();
      this.currentTime = `当前时间：${y}年${mt}月${day}日-${h}时${m}分${s}秒`;
    },
      initChartsbar() {
        console.log('柱状图')
      // 示例：初始化柱形图
      const myChart = echarts.init(document.querySelector('.pie .chart'));
        var option = {
          color: ["#1089E7", "#F57474", "#56D0E3"],
          tooltip: {
            trigger: 'item',
            formatter: '{a} <br/>{b}: {c} ({d}%)'
          },
          legend: {
            // 垂直居中,默认水平居中
            // orient: 'vertical',

            bottom: 0,
            left: 10,
            // 小图标的宽度和高度
            itemWidth: 10,
            itemHeight: 10,
            // 修改图例组件的文字为 12px
            textStyle: {
              color: "rgba(255,255,255,.5)",
              fontSize: "10"
            }
          },
          series: [{
            name: '类型统计',
            type: 'pie',
            // 设置饼形图在容器中的位置
            center: ["50%", "42%"],
            // 修改饼形图大小，第一个为内圆半径，第二个为外圆半径
            radius: ['40%', '60%'],
            avoidLabelOverlap: false,
            // 图形上的文字
            label: {
              show: false,
              position: 'center'
            },
            // 链接文字和图形的线
            labelLine: {
              show: false
            },
            data: [{
              value: 1,
              name: "智能电表"
            },
              {
                value: 1,
                name: "温湿度传感器"
              },
              {
                value: 2,
                name: "智能开关面板"
              }
            ]
          }]
      };

      myChart.setOption(option);
      window.addEventListener('resize', function () {
        myChart.resize();
      })


    },




    initChartLisewet(){
      console.log('线图2')
      console.log(this.TempData)
      console.log(this.HumiData)
      const yearData = [];
      for (const timePeriod in this.TempData) {
        if (this.TempData.hasOwnProperty(timePeriod)) {
          const tempData = this.TempData[timePeriod];
          const humiData = this.HumiData[timePeriod];

          // 处理当前时间段的数据
          const timeData = {
            time: timePeriod,
            data: [[], []] // 初始化两个空数组，用于存放温度和湿度的数据
          };

          // 遍历当前时间段的日期或小时数据
          for (const key in tempData) {
            if (tempData.hasOwnProperty(key) && humiData.hasOwnProperty(key)) {
              const tempValues = tempData[key][0]; // 假设每个 key 对应的值是一个数组，取第一个元素
              const humiValues = humiData[key][0]; // 假设每个 key 对应的值是一个数组，取第一个元素

              // 将温度和湿度分别放入对应的数组中
              timeData.data[0].push(tempValues); // 温度数据
              timeData.data[1].push(humiValues); // 湿度数据
            }
          }

          // 将处理好的时间段数据推入 yearData 数组
          yearData.push(timeData);
        }
      }

      console.log(yearData);


      let xAxisDataDay = [];
      let xAxisDataHour = [];
      let xAxisDataMonth = [];

// 遍历 originalData 中的 day 对象的属性，将日期作为 x 轴的数据
      for (const timeKey in  this.TempData.day) {
        if ( this.TempData.day.hasOwnProperty(timeKey)) {
          // 这里可以根据实际情况进行格式化
          xAxisDataDay.push(timeKey); // 将日期添加到 day 的 x 轴数据数组中
        }
      }

// 遍历 originalData 中的 hour 对象的属性，将小时作为 x 轴的数据
      for (const timeKey in this.TempData.hour) {
        if (this.TempData.hour.hasOwnProperty(timeKey)) {
          // 这里可以根据实际情况进行格式化
          xAxisDataHour.push(timeKey); // 将小时添加到 hour 的 x 轴数据数组中
        }
      }

// 遍历 originalData 中的 month 对象的属性，将月份作为 x 轴的数据
      for (const timeKey in this.TempData.month) {
        if (this.TempData.month.hasOwnProperty(timeKey)) {
          // 这里可以根据实际情况进行格式化
          xAxisDataMonth.push(timeKey); // 将月份添加到 month 的 x 轴数据数组中
        }
      }

      let currentXAxisData = xAxisDataDay

      // const yearData = [{
      //   time: "hour", // 年份
      //   data: [
      //     // 两个数组是因为有两条线
      //     [24, 40, 101, 134, 90, 230, 210, 230, 120, 230, 210, 120],
      //     [40, 64, 191, 324, 290, 330, 310, 213, 180, 200, 180, 79]
      //   ]
      // },
      //   {
      //     time: "day", // 年份
      //     data: [
      //       // 两个数组是因为有两条线
      //       [123, 175, 112, 197, 121, 67, 98, 21, 43, 64, 76, 38],
      //       [143, 131, 165, 123, 178, 21, 82, 64, 43, 60, 19, 34]
      //     ]
      //   },
      //   {
      //     time: "month", // 年份
      //     data: [
      //       // 两个数组是因为有两条线
      //       [ 67, 98, 21, 43, 64, 76,123, 175, 112, 197, 121,  38],
      //       [ 82, 64, 43, 60, 19, 143, 131, 165, 123, 178, 21, 34]
      //     ]
      //   }
      // ];

      const myChart = echarts.init(document.querySelector(".linepart2 .chart"));

      const option = {
        // 修改两条线的颜色
        color: ['#00f2f1', '#ed3f35'],
        tooltip: {
          trigger: 'axis'
        },
        // 图例组件
        legend: {
          // 当serise 有name值时， legend 不需要写data
          // 修改图例组件文字颜色
          textStyle: {
            color: '#4c9bfd'
          },
          right: '10%',
        },
        grid: {
          top: "20%",
          left: '3%',
          right: '4%',
          bottom: '3%',
          containLabel: true,
          show: true, // 显示边框
          borderColor: '#012f4a' // 边框颜色
        },
        xAxis: {
          type: 'category',
          boundaryGap: false, // 去除轴间距
          data: currentXAxisData,
          // 去除刻度线
          axisTick: {
            show: false
          },
          axisLabel: {
            color: "#4c9bfb" // x轴文本颜色
          },
          axisLine: {
            show: false // 去除轴线
          }
        },
        yAxis: {
          type: 'value',
          // 去除刻度线
          axisTick: {
            show: false
          },
          axisLabel: {
            color: "#4c9bfb" // x轴文本颜色
          },
          axisLine: {
            show: false // 去除轴线
          },
          splitLine: {
            lineStyle: {
              color: "#012f4a"
            }
          }
        },
        series: [{
          type: 'line',
          smooth: true, // 圆滑的线
          name: '温度',
          data: yearData[0].data[0]
        },
          {
            type: 'line',
            smooth: true, // 圆滑的线
            name: '湿度',
            data: yearData[0].data[1]
          }
        ]
      };

      myChart.setOption(option);

      // 4.让图表随屏幕自适应
      window.addEventListener('resize', function () {
        myChart.resize();
      })

      // 5.点击切换hour 和 day 和 month 的数据
      $('.linepart2 h2 a').on('click', function () {
        // console.log($(this).index());
        // 点击a 之后 根据当前a的索引号 找到对应的 yearData 相关对象
        // console.log(yearData[$(this).index()]);
        const obj = yearData[$(this).index()];
        option.series[0].data = obj.data[0];
        option.series[1].data = obj.data[1];
        console.log(obj)
        if(obj.time ==='hour') {
          currentXAxisData = xAxisDataHour
        }
        else if(obj.time ==='day'){
          currentXAxisData = xAxisDataDay
        }
        else if(obj.time ==='month'){
          currentXAxisData = xAxisDataMonth
        }
        else {
          currentXAxisData = xAxisDataDay
        }
        option.xAxis.data = currentXAxisData;

        // 选中年份高亮
        $('.linepart2 h2 a').removeClass('a-active');
        $(this).addClass('a-active');

        // 需要重新渲染
        myChart.setOption(option);
      })

    },

    //查询
    getList() {
      this.loading = true
      listDevice(this.queryParams).then(response => {
          this.total = response.data.count
        console.log(this.total)
        }
      )
    },
    async FindData() {
      await getTempData().then(response =>{
        this.TempData = response
      })
      console.log( this.TempData)
      await getHumiData().then(response =>{
        this.HumiData = response
      })
      console.log( this.HumiData)

      // 在 mounted 钩子中初始化图表
      this.initChartsbar()
      //湿度预测
      this.initChartLisewet()
    },
  }
}


</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
li {
  list-style: none;
}
@font-face {
  font-family: electronicFont;
  src: url(../../font/DS-DIGIT.TTF);
}
body {
  background: url(../../assets/images/bg.jpg) no-repeat;
  line-height: 1.15;
}
header {
  position: relative;
  height: 1.25rem;
  background: url(../../assets/images/head_bg.png) no-repeat;
  background-size: 100% 100%;
}
header h1 {
  font-size: 0.475rem;
  color: rgba(255, 255, 255, 0.87);
  text-align: center;
  line-height: 1rem;
}
header .show-time {
  position: absolute;
  top: 0;
  right: 0.375rem;
  line-height: 0.9375rem;
  color: rgba(255, 255, 255, 0.7);
  font-size: 0.25rem;
}
.mainbox {
  display: flex;
  min-width: 1024px;
  max-height: 1920px;
  margin: 0 auto;
  padding: 0.125rem 0.125rem 0;
}
.mainbox .column {
  flex: 3;
//background: #52c41a;
}
.mainbox .column:nth-child(2) {
  flex: 5;
  margin: 0 0.125rem 0.1875rem;
  overflow: hidden;
//background: #ff91e7;
}
.mainbox .panel {
  position: relative;
  height: 3.875rem;
  padding: 0 0.1875rem 0.5rem;
  margin-bottom: 0.1875rem;
  border: 1px solid rgba(25, 186, 139, 0.17);
  background: url(../../assets/images/line.png) rgba(255, 255, 255, 0.03);
}
.mainbox .panel::before {
  position: absolute;
  top: 0;
  left: 0;
  width: 10px;
  height: 10px;
  border-left: 2px solid #02a6b5;
  border-top: 2px solid #02a6b5;
  content: '';
}
.mainbox .panel::after {
  position: absolute;
  top: 0;
  right: 0;
  width: 10px;
  height: 10px;
  border-right: 2px solid #02a6b5;
  border-top: 2px solid #02a6b5;
  content: '';
}
.mainbox .panel .panel-footer {
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
}
.mainbox .panel .panel-footer::before {
  position: absolute;
  bottom: 0;
  left: 0;
  width: 10px;
  height: 10px;
  border-left: 2px solid #02a6b5;
  border-bottom: 2px solid #02a6b5;
  content: '';
}
.mainbox .panel .panel-footer::after {
  position: absolute;
  bottom: 0;
  right: 0;
  width: 10px;
  height: 10px;
  border-right: 2px solid #02a6b5;
  border-bottom: 2px solid #02a6b5;
  content: '';
}
.mainbox .panel h2 {
  height: 0.6rem;
  color: #fff;
  line-height: 0.6rem;
  text-align: center;
  font-size: 0.25rem;
  font-weight: 400;
}
.mainbox .panel h2 a {
  color: #a7a7a7;
  text-decoration: none;
  margin: 0 0.125rem;
}
.mainbox .panel h2 .a-active {
  color: #fff;
}
.mainbox .panel .chart {
  height: 3rem;
}
.no {
  background-color: rgba(101, 132, 226, 0.1);
  padding: 0.1875rem;
}
.no .no-hd {
  position: relative;
  border: 1px solid rgba(25, 186, 139, 0.17);
}
.no .no-hd::before {
  position: absolute;
  top: 0;
  left: 0;
  width: 0.375rem;
  height: 0.125rem;
  border-top: 2px solid #02a6b5;
  border-left: 2px solid #02a6b5;
  content: '';
}
.no .no-hd::after {
  position: absolute;
  bottom: 0;
  right: 0;
  width: 0.375rem;
  height: 0.125rem;
  border-bottom: 2px solid #02a6b5;
  border-right: 2px solid #02a6b5;
  content: '';
}
.no .no-hd ul {
  display: flex;
}
.no .no-hd ul li {
  position: relative;
  flex: 1;
  line-height: 1rem;
  font-size: 0.875rem;
  color: #ffeb7b;
  text-align: center;
  font-family: electronicFont;
}
.no .no-hd ul li:first-child::after {
  content: '';
  position: absolute;
  top: 25%;
  right: 0;
  height: 50%;
  width: 1px;
  background-color: rgba(255, 255, 255, 0.3);
}
.no .no-bd ul {
  display: flex;
}
.no .no-bd ul li {
  flex: 1;
  text-align: center;
  color: rgba(255, 255, 255, 0.7);
  font-size: 0.225rem;
  height: 0.5rem;
  line-height: 0.5rem;
  padding-top: 0.125rem;
}
.map {
  position: relative;
  height: 10.125rem;
}
.map .map1 {
  width: 6.475rem;
  height: 6.475rem;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background: url(../../assets/images/map.png);
  background-size: 100% 100%;
  opacity: 0.3;
}
.map .map2 {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 8.0375rem;
  height: 8.0375rem;
  background: url(../../assets/images/lbx.png);
  background-size: 100% 100%;
  animation: rotate1 15s linear infinite;
  opacity: 0.6;
}
.map .map3 {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 7.075rem;
  height: 7.075rem;
  background: url(../../assets/images/jt.png);
  background-size: 100% 100%;
  animation: rotate2 15s linear infinite;
}
.map .chart {
  position: absolute;
  top: 0;
  left: 0;
  height: 10.125rem;
  width: 100%;
}
@keyframes rotate1 {
  from {
    transform: translate(-50%, -50%) rotate(0deg);
  }
  to {
    transform: translate(-50%, -50%) rotate(360deg);
  }
}
@keyframes rotate2 {
  from {
    transform: translate(-50%, -50%) rotate(360deg);
  }
  to {
    transform: translate(-50%, -50%) rotate(0deg);
  }
}
/* 约束屏幕尺寸 */
@media screen and (max-width: 1024px) {
  html {
    font-size: 42px !important;
  }
}
@media screen and (min-width: 1920px) {
  html {
    font-size: 80px !important;
  }
}

#special-page {
  background: url(../../assets/images/bg.jpg) no-repeat top center;
  line-height: 1.15;
}

.panel .chart {
  display: flex !important;
  flex-direction: column !important;
  align-items: flex-start !important;
}

.panel .chart .info-item {
  display: flex !important;
  align-items: center !important;
  margin-bottom: 10px !important;
}

.panel .chart .info-item img {
  width: 70px !important;
  height: 70px !important;
  margin-left: 30px !important; /* 调整图片与容器左侧之间的间距 */
  margin-right: 20px !important;
}

.panel .chart .info-item p {
  margin: 0 !important;
  font-size: 20px !important;
  color: #fff !important;
}

.panel .chart .infotwo-item {
  display: flex !important;
  align-items: center !important;
  margin-bottom: 10px !important;
}

.panel .chart .infotwo-item img {
  width: 70px !important;
  height: 70px !important;
  margin-left: 30px !important; /* 调整图片与容器左侧之间的间距 */
  margin-right: 20px !important;
}

.panel .chart .infotwo-item .text-container {
  display: flex;
  flex-direction: column; /* 使文字内部上下排列 */
}

.panel .chart .infotwo-item p {
  margin: 2px !important;
  font-size: 12px !important;
  color: #fff !important;
}


.center-mysrc {
  text-align: center!important;
}

.center-mysrc img {
  display: inline-block !important;
  //margin-left: 30px !important; /* 调整图片与容器左侧之间的间距 */
  //margin-right: 85px !important;
}


.Round-wrap .status::before {
  content: "";
  display: block;
  width: 3px; /* 竖线的宽度 */
  height: 22px; /* 竖线的高度，可以根据需要调整 */
  background-color: rgb(255, 255, 255); /* 竖线的颜色 */
  position: absolute;
  left: 40px; /* 根据需要调整竖线的位置 */
  top: -15px;
}
.Round-wrap .status{
  color: #ffffff;
}
</style>

<style>
#Round-chart {
  height: 150px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  .Round-wrap {
    text-align: center;
    font-weight: bold;
    .Round {
      width: 80px;
      height: 80px;
      border: 10px solid greenyellow;
      border-bottom: 10px dashed transparent !important;
      border-radius: 50%;
      display: flex;
      justify-content: center;
      align-items: center;
      p {
        font-size: 20px;
        color: #ffffff;
      }
    }
  }
}
</style>


