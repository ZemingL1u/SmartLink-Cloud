<template>
  <div id="round-chart">
    <div class="round-wrap" v-for="(key, idx) in Object.keys(statusDeviceMap)" :key="idx">
      <div class="round" :style="{'border-color': getRandomColor()}">
        <p>{{ statusDeviceMap[key] }}</p>
      </div>
      <p class="status">{{ key }}</p>
    </div>
  </div>
</template>

<script>
import echarts from 'echarts'
import resize from './mixins/resize'

export default {
  mixins: [resize],
  props: {
    statusDeviceMap: {
      type: Object,
      default: () => ({})
    }
  },
  data() {
    return {
      chart: null
    }
  },
  mounted() {
    this.$nextTick(() => {
      this.initChart()
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
    initChart() {
      const chartDom = document.getElementById('pie-chart');
      this.chart = echarts.init(chartDom)

      const option = {
      tooltip: {
        trigger: 'item'
      },
      legend: {
        top: '5%',
        left: 'center'
      },
      series: [
        {
          name: 'Access From',
          type: 'pie',
          radius: ['40%', '70%'],
          avoidLabelOverlap: false,
          itemStyle: {
            borderRadius: 10,
            borderColor: '#fff',
            borderWidth: 2
          },
          label: {
            show: false,
            position: 'center'
          },
          emphasis: {
            label: {
              show: true,
              fontSize: 40,
              fontWeight: 'bold'
            }
          },
          labelLine: {
            show: false
          },
          data: [
            { value: 1048, name: 'Search Engine' },
            { value: 735, name: 'Direct' },
            { value: 580, name: 'Email' },
            { value: 484, name: 'Union Ads' },
            { value: 300, name: 'Video Ads' }
          ]
        }
      ]
    };

      this.chart.setOption(option)
    },
    
    getRandomColor() {
      const r = Math.floor(Math.random() * 256);
      const g = Math.floor(Math.random() * 256);
      const b = Math.floor(Math.random() * 256);
      return 'rgb(' + r + ',' + g + ',' + b + ')';
    }
  }
}
</script>
<style lang="scss" scoped>
#round-chart {
  height: 200px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  .round-wrap {
    text-align: center;
    font-weight: bold;
    .round {
      width: 100px;
      height: 100px;
      border: 10px solid greenyellow;
      border-bottom: 10px dashed transparent !important;
      border-radius: 50%;
      display: flex;
      justify-content: center;
      align-items: center;
      p {
        font-size: 20px;
      }
    }
  }
}
.status {
  position: relative;
  &::before {
    content: "";
    width: 3px;
    height: 40px;
    background-color: #e8e8e8;
    position: absolute;
    top: -24px;
    left: 50%;
    transform: translateY(-50%);
  }
}
</style>

