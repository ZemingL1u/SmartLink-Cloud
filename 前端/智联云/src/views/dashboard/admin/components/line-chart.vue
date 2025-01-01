<template>
  <div id="line-chart"/>
</template>

<script>
import echarts from 'echarts'
import resize from './mixins/resize'
import { line } from '@/api/admin/dashboard/index'
export default {
  mixins: [resize],
  data() {
    return {
      chart: null
    }
  },
  mounted() {
    this.$nextTick(() => {
      this.getLine()
    })
  },
  beforeDestroy() {
    if (!this.chart) {
      return
    }
    this.chart.dispose()
    this.chart = null
  },
  methods: {
    initChart(data) {
      const chartDom = document.getElementById('line-chart');
      this.chart = echarts.init(chartDom)

      const option = {
        tooltip: {
          trigger: 'axis'
        },
        xAxis: {
          type: 'category',
          data: data.xAxis
        },
        yAxis: {
          type: 'value'
        },
        legend: {
          data: ['粤海校区', '丽湖校区']
        },
        series: [
          {
            name: '粤海校区',
            data: data.yAxis1,
            type: 'line'
          },
          {
            name: '丽湖校区',
            data: data.yAxis2,
            type: 'line'
          }
        ]
      };

      this.chart.setOption(option)
    },
    // 这里是获取折线图的接口
    getLine() {
      /* line().then(res => {
        this.initChart(res.data)
      }) */

      // 这里是模拟的
      this.initChart({
        xAxis: ['1月', '2月', '3月', '4月', '5月', '6月', '7月'],
        yAxis1: [150, 230, 224, 218, 135, 147, 260],
        yAxis2: [110, 210, 234, 248, 155, 157, 60]
      })
    }
  }
}
</script>
<style lang="scss" scoped>
#line-chart {
  height: 300px;
}
</style>