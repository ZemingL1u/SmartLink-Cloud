<template>
  <div id="pie-chart"/>
</template>

<script>
import echarts from 'echarts'
import resize from './mixins/resize'

export default {
  mixins: [resize],
  props: {
    typeDeviceMap: {
      type: Object,
      default: () => ({})
    }
  },
  data() {
    return {
      chart: null,
    }
  },
  beforeDestroy() {
    if (!this.chart) {
      return
    }
    this.chart.dispose()
    this.chart = null
  },
  methods: {
    initChart() {
      const chartDom = document.getElementById('pie-chart');
      this.chart = echarts.init(chartDom)

      const option = {
        tooltip: {
          trigger: 'item'
        },
        legend: {
          orient: 'vertical',
          top: 'center',
          right: '0',
          type: 'scroll',
        },
        series: [
          {
            name: '设备数量统计（类型）',
            type: 'pie',
            radius: ['40%', '70%'],
            avoidLabelOverlap: false,
            itemStyle: {
              borderRadius: 10,
              borderColor: '#fff',
              borderWidth: 2
            },
            emphasis: {
              label: {
                show: true,
                fontSize: 18,
                fontWeight: 'bold'
              }
            },
            labelLine: {
              show: false
            },
            data: Object.entries(this.typeDeviceMap).map(([key, val]) => ({ value: val, name: key }))
            // [{ value: 1048, name: 'Search Engine' }]
          }
        ]
      };

      this.chart.setOption(option)
    }
  },
  watch: {
    typeDeviceMap(val) {
      this.$nextTick(() => {
        this.initChart()
      })
    }
  }
}
</script>
<style lang="scss" scoped>
#pie-chart {
  height: 200px;
}
</style>