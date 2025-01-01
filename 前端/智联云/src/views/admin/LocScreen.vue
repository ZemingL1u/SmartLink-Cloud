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
              <p>类型数：3</p>
            </div>
          </div>
          <div class="panel-footer"></div>
        </div>
        <div class="panel pie">
          <h2>用能结构</h2>
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
              <li>{{this.Data.leader}}</li>
              <li>{{this.Data.children.length}}</li>
            </ul>
          </div>
          <div class="no-bd">
            <ul>
              <li>负责人</li>
              <li>子区域</li>
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
          <h2>地图</h2>
          <div class="center-mysrc">
            <img style="width: 200px;height: 150px;margin: 0 auto" src="../../assets/work5.png" alt="">
          </div>
        </div>
        <div class="panel bar2">
          <h2>设备数量对总数量的占比</h2>
          <div class="chart">
          </div>
          <div class="panel-footer"></div>
        </div>
        <div class="panel linepart2">
          <h2>地区总耗能实时监控</h2>
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
        color: ["#1089E7", "#00fd00", "#56D0E3", "#F8B448"],
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
          name: '用能',
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
            value: 110,
            name: "用水"
          },
            {
              value: 224,
              name: "电量"
            },
            {
              value: 50,
              name: "汽油"
            },
            {
              value: 30,
              name: "燃气"
            },
          ]
        }]
      };

      myChart.setOption(option);
      window.addEventListener('resize', function () {
        myChart.resize();
      })


    },

    initChartsbar2() {
      console.log('柱状图')
      // 示例：初始化柱形图
      const myChart = echarts.init(document.querySelector('.bar2 .chart'));
      var myColor = ["#1089E7", "#F57474", "#56D0E3", "#F8B448", "#8B78F6"];
      // 2.指定配置项和数据
      var option = {
        grid: {
          top: "10%",
          left: '22%',
          bottom: '10%',
          // containLabel: true
        },
        xAxis: {
          // 不显示x轴相关信息
          show: false
        },
        yAxis: [{
          type: 'category',
          // y轴数据反转，与数组的顺序一致
          inverse: true,
          // 不显示y轴线和刻度
          axisLine: {
            show: false
          },
          axisTick: {
            show: false
          },
          // 将刻度标签文字设置为白色
          axisLabel: {
            color: "#fff"
          },
          data: ["信号覆盖", "环境", "能耗与联动", "安全", "人流统计"]
        }, {
          // y轴数据反转，与数组的顺序一致
          inverse: true,
          show: true,
          // 不显示y轴线和刻度
          axisLine: {
            show: false
          },
          axisTick: {
            show: false
          },
          // 将刻度标签文字设置为白色
          axisLabel: {
            color: "#fff"
          },
          data: [15, 10, 8, 11, 5]
        }],
        series: [{
          // 第一组柱子（条状）
          name: '条',
          type: 'bar',
          // 柱子之间的距离
          barCategoryGap: 50,
          // 柱子的宽度
          barWidth: 10,
          // 层级 相当于z-index
          yAxisIndex: 0,
          // 柱子更改样式
          itemStyle: {
            barBorderRadius: 20,
            // 此时的color可以修改柱子的颜色
            color: function (params) {
              // params 传进来的是柱子的对象
              // dataIndex 是当前柱子的索引号
              // console.log(params);
              return myColor[params.dataIndex];
            }
          },
          data: [60, 40, 37.5, 19.1, 40],
          // 显示柱子内的百分比文字
          label: {
            show: true,
            position: "inside",
            // {c} 会自动解析为数据（data内的数据）
            formatter: "{c}%"
          }
        },
          {
            // 第二组柱子（框状 border）
            name: '框',
            type: 'bar',
            // 柱子之间的距离
            barCategoryGap: 50,
            // 柱子的宽度
            barWidth: 14,
            // 层级 相当于z-index
            yAxisIndex: 1,
            // 柱子修改样式
            itemStyle: {
              color: "none",
              borderColor: "#00c1de",
              borderWidth: 2,
              barBorderRadius: 15,
            },
            data: [100, 100, 100, 100, 100]
          }
        ]
      };
      // 3.把配置项给实例对象
      myChart.setOption(option);

      // 4.让图表随屏幕自适应
      window.addEventListener('resize', function () {
        myChart.resize();
      })


    },


    initChartLisewet(){
      console.log('线图2')
      console.log(this.TempData)
      console.log(this.HumiData)

      const myChart = echarts.init(document.querySelector(".linepart2 .chart"));

      const option = {
        color: ['#2f89cf'],
        // 提示框组件
        tooltip: {
          trigger: 'axis',
          axisPointer: { // 坐标轴指示器，坐标轴触发有效
            type: 'shadow' // 默认为直线，可选为：'line' | 'shadow'
          }
        },
        // 修改图表位置大小
        grid: {
          left: '0%',
          top: '10px',
          right: '0%',
          bottom: '4%',
          containLabel: true
        },
        // x轴相关配置
        xAxis: [{
          type: 'category',
          data: ["8：00", "8：30", "9：00", "9：30", "10：00", "10：30", "11：00"],
          axisTick: {
            alignWithLabel: true
          },
          // 修改刻度标签，相关样式
          axisLabel: {
            color: "rgba(255,255,255,0.8)",
            fontSize: 10
          },
          // x轴样式不显示
          axisLine: {
            show: false
          }
        }],
        // y轴相关配置
        yAxis: [{
          type: 'value',
          // 修改刻度标签，相关样式
          axisLabel: {
            color: "rgba(255,255,255,0.6)",
            fontSize: 12
          },
          // y轴样式修改
          axisLine: {
            lineStyle: {
              color: "rgba(255,255,255,0.6)",
              width: 2
            }
          },
          // y轴分割线的颜色
          splitLine: {
            lineStyle: {
              color: "rgba(255,255,255,0.1)"
            }
          }
        }],
        // 系列列表配置
        series: [{
          name: '耗能',
          type: 'bar',
          barWidth: '35%',
          // ajax传动态数据
          data: [200, 300, 300, 400, 700, 350, 450],
          itemStyle: {
            // 修改柱子圆角
            barBorderRadius: 5
          }
        }]
      };
      // 3.把配置项给实例对象
      myChart.setOption(option);

      // 4.让图表随屏幕自适应
      window.addEventListener('resize', function () {
        myChart.resize();
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

      this.initChartsbar2()
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


