<template>
  <div class="dashboard-editor-container">
    <el-row :gutter="12">
      <el-col :sm="24" :xs="24" :md="6" :xl="6" :lg="6" :style="{ marginBottom: '12px' }">
        <chart-card>
          <div class="card-wrap">
            <img src="data:image/png;base64,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" alt="logo">
            <div class="info-wrap">
              <p class="value">{{ deviceCount }}</p>
              <p class="title">设备数</p>
            </div>
          </div>
          <template slot="footer"><p>今日新增 0</p></template>
        </chart-card>
      </el-col>
      <el-col :sm="24" :xs="24" :md="6" :xl="6" :lg="6" :style="{ marginBottom: '12px' }">
        <chart-card>
          <div class="card-wrap">
            <img src="data:image/png;base64,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" alt="logo">
            <div class="info-wrap">
              <p class="value">{{ typeCount }}</p>
              <p class="title">类型数</p>
            </div>
          </div>
          <template slot="footer"><p>今日新增 0</p></template>
        </chart-card>
      </el-col>
      <el-col :sm="24" :xs="24" :md="6" :xl="6" :lg="6" :style="{ marginBottom: '12px' }">
        <chart-card>
          <div class="card-wrap">
            <img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAALQAAAC0CAMAAAAKE/YAAAAAAXNSR0IArs4c6QAAAjFQTFRFAAAA/////1VV/01N6kBA70BA8Tk58jY28jMz8z098zo69Dc39T099Ts79jk59jc39j4+9js77zo69zo670BA8D4+8Do68T4+8jw88js78z4+8z097Ts77kBA9D4+7j0970FB7z4+8EFB8EBA8T8/8T098UFB7UBA8kJC7kBA7j8/70JC70FB7EBA7T8/8UFB8UBA8UJC7kFB7kBA7j8/7z4+7z097z097D8/7D4+7T098D098EBA8D8/8D4+8D4+7kBA7j4+7kBA7D8/7D4+7D4+7EBA7z4+8D8/7j8/7EBA7D8/70BA7z8/70FB70BA7UBA7EFB7EBA7D8/7EFB7EBA7kBA70FB70FB70BA60FB7EFB7EBA7kBA7j8/70BA7z8/7z8/7T4+7EBA7D8/7D8/7D4+7j8/7kBA7kBA7j8/7j8/7UBA7UBA7T8/7D8/7EBA7D8/7D8/7kFB7kBA7kBA7j8/7kFB7kBA7UBA7T8/7UFB7EBA7D8/7EBA7EBA7EBA7kFB7kBA7UBA7UBA7UBA7EBA7EBA7EBA7EFB7UBA7UBA7T8/7T8/7T8/7EBA7D8/7D8/7T8/7kBA7D8/7UBA7T8/7T8/7UBA7UBA7T8/7EBA7UBA7EBA7EBA7UBA7T8/7UBA7EBA7UBA7UFB7EBA7EBA7EBA7EBA7UBA7UBA7UBA7UBA7UFB7EBA7EFB7EFB7EBA7EBA7EFB7UFB7UBA7UBA7UFB7UFB7UBA7EBAOvnA1gAAALp0Uk5TAAEGCgwQEhMUFRYXGRobHB0eHx8gISMlJicpKissLS4vMTM0NTY3ODo8PT4/REVHSElLTE1OT1BRUlNTVFVWV1haXF1eX2BiZWZrbm9xcnN0dnh5ent7fX5/goaHiImMjY6PkJGSk5WXmJmam5ydnqChoqKjpKWmp6ipqqytrq+wsbO0tre6u7y9wsPFxsfIycrKzM3Q0dLT1NXX2Nrc3+Hi5OTl5ufo6uzu7/Dx8vT19vf4+fr7/P3+8csvowAAA9ZJREFUeNrt3fdTE0EUB/AEyIViFESwYUEUsVdUsGIHxYoNG1awIEIMQUVRsWCj20ARG2pAaZGyf517l6BJbsNMIjO763zfT7z3LpNPls1u7n5Zg0ENY0iYSRE+TGEhRsOfCJFA7HaHuMnGUEWiCHUNtlRmqtbmhiJZ0BliNMmGNhnlG2h1qMPkQ4cZTPKhTQZFwgAaaKCBBhpooIEGGmiggQYaaKCBBhpooIEGGmiggQYaaKCBBhpooIEWGz36fkeRWTb0PkJIumzocxSdJT7aLCH6dF9z0jDoxU0dGcKhEymyIdwvOq6NEIdw6IQBqjziF20nIqKVK5TVM8sPeh3NyF7x0NEfqKvWzETHfqLZMxFXjzR1NA8w0SU06Z4h5Dqt0roS3ckumixhfhyx0GPVSXDGnUQWvz021Kii9TqzmGhlLcXlMuo2r68od3RMtFe6++U1C+OquEfVK70Loziidzh/5UUF/Kpx9sEvC/ih6TZHmuYzW7FTkibHsCfRZ/qqCn7o1+qy0Hc80rtqybC+6FQ7pL3+Qmq4z7fVqnUu8UOntGiCV3P+liIz73URz3DY13su5h+1Yu0Ejl9ES8GganBuGCpsaSH6qB9asJVsLe85ZOa5eijK0vcqo9qVLKoj7Kh0r3itatKQzH2dHlNIB/uW9ufhAeIvurZqVzyl/5XccBE2lxUPbOoUjbKTYWLwqHrppPKKFJF2xPgaMnzcihqptxoxdORzFtTpmdiEQ19lmYsavdL9gqH3sMzF5p9eef8qodBz+xhmqzLVp+KIFwldxTDTKbzGt5YvEDqVYbbTXe+gb7F3mjjoRr35urpTW3XlMmHQq/Xmm9qvi3pdfWC6KOhina3ctVX/0H+aHFHQX31lt13mBMZUfyIIeqEv7E6Ex6MDn+iLEwOd5+O66zYrOawdJ1MM9GNv1bedWTRm08Z2W2lpaVnPyC/VI4BuZQ3oRs/nTp7xUAx0Pws9098PqSYh0BaWudd9e1Kk67QKgR7PQr9x9Qr1ne9CoCcyb1O01kVGxyEuWrsjLCByoTfTRj6RDE0fbZwnkqGdEcpZIhu6WTlFpENXnCTyoduIhGgCNNBAA/1/oi3ZAUWWEGgOAbQ06E35NwKLvGW80VFVJPC4zBl9ggQT2/ii3wWFruSL7gwK3cgXXRMUuoQvOj0Yc3cS5yUvsz1gc+ty7ut0xLz0gCItGTsi0EADDTTQQAMNNNBAAw000EADDTTQQAMNNNBAAw000EADDTTQQAMN9L+gpTw+U8qDSqU8ElbKw3elPOZYzgOl5Ty6W8pD0uU7jv43LZPJrU5hi/8AAAAASUVORK5CYII=" alt="logo">
            <div class="info-wrap">
              <p class="value">{{ areaCount }}</p>
              <p class="title">区域数</p>
            </div>
          </div>
          <template slot="footer"><p>今日新增 0</p></template>
        </chart-card>
      </el-col>
      <el-col :sm="24" :xs="24" :md="6" :xl="6" :lg="6" :style="{ marginBottom: '12px' }">
        <chart-card>
          <div class="card-wrap">
            <img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAALQAAAC0CAMAAAAKE/YAAAAAAXNSR0IArs4c6QAAARpQTFRFAAAA//8A/9UA/8wA6tUA778A788A8cYA8skA884A88UA9MgA9cwA9sgA9soA/84A+MYA+MgA+MkA+MsA+cYA+ccA+cgA/8sA+c0A+s0A+skA+8kA98oA98oA98oA+8kA/MsA/MoA+coA+coA+skA98gA+soA+MkA+MkA+MoA+MgA+8kA+8oA+cgA+cgA+ckA+ckA+coA+soA+ssA+sgA+cgA+ckA+soA+ssA+8kA+ckA+ckA+coA+coA+MoA+soA+skA+ckA+coA+soA+soA+skA+skA+skA+ckA+coA+coA+coA+coA+ckA+ckA+ckA+ckA+coA+coA+skA+soA+soA+skA+soA+soA+soA+skA+skA+soA+ckAQIu+owAAAF10Uk5TAAEGCgwQEBITFRYXGRwdHyQlJicoKSosLjM0OT4/Q0dUV1tcXmFlZ2hpa3JzeHmIiouMmJqfoKSmqq+wsba7v8PIycrO09TV1tfY2dre3+Hj5ebp7O3y9PX5+/z9LWG3TwAAAmVJREFUeNrt3VlP20AUhmE7JDkQljqBdANCN0oXlhZKWdIChdYQQqEbtITm//+NxvHYJEqvOsb1kd5zN59G8qPR2JatkY7jBOXm8gXJfBXyOdeJK1cUJVXMGbI7JIpqKFxsVeaOurs3RFl1dohb1IYuuvoWOljqvD503inoQxccUVigQYMGDRo0aNCgQYMGDRo0aNCgQYMGDRo0aNCgQYMGDRo0aNCgQYNOGD298u79P1d9Zfp/oBcv21Z1uZg+erbVtqzWbOrotbZ1raWO3rNH76WOPrBHH4AGDRo0aNCgQYMGDbpbZ1tvP/UFJxsbJ33BxzdbZ9lCvyp1oqc/r7//loI5S9ffkb8ed8al11lCb4bZyzhYDYPVOHgeBpsZQt8z4Q8zvhoPx+NXJvhmJtzPDvoiCndN0IyCpgl2ouAiM+jfEyaM7sXzaNa5CT6Y8UgrO9vjSZjdiUm1MKjFN2Y1DB5laE9/LQdRaTcO/O7aT/hxsDMSBJXPWXrkfX9x13vY+6A+na/enj/tFT64VX72hTciaNCgQYMGDRo0aNCg00TX7dH11NHL9ujl1NGjR7bmo9HU0eJtW534aG17kj5aZGymNlBz/qDPnxucNzNmceHED155A2rfS/oayZ8W8w77zYeJm2/iiFu50WtulEUDWio96kZFdKBl8jgyH0+KFrRMmf/SzSnRg5bqfmDer4omtAwvrK8vDIsu9I0WaNCgQYMGDRo0aNCgQYMGDRo0aNCgQYMGDRo0aNCgQYMGDRo0aNCg/4pW2T5TZaNSlS1hVTbfVdnmWGdDaZ2tu1U2SdfXjv4PSA7YHqfsVcEAAAAASUVORK5CYII=" alt="logo">
            <div class="info-wrap">
              <p class="value">{{ userCount }}</p>
              <p class="title">用户数</p>
            </div>
          </div>
          <template slot="footer"><p>今日新增 0</p></template>
        </chart-card>
      </el-col>
    </el-row>

    <el-row :gutter="12">
      <el-col :sm="24" :xs="24" :md="12" :xl="12" :lg="12" :style="{ marginBottom: '12px' }">
        <el-card>
          <h4 class="card-title">设备数量统计（类型）</h4>
          <pie-chart :typeDeviceMap="typeDeviceMap"/>
        </el-card>
      </el-col>
      <el-col :sm="24" :xs="24" :md="12" :xl="12" :lg="12" :style="{ marginBottom: '12px' }">
        <el-card>
          <h4 class="card-title">设备状态统计</h4>
          <round-chart :statusDeviceMap="statusDeviceMap"/>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="12">
      <el-col :sm="24" :xs="24" :md="24" :xl="24" :lg="24" :style="{ marginBottom: '12px' }">
        <el-card>
          <line-chart/>
        </el-card>
      </el-col>
    </el-row>

  </div>
</template>

<script>
import ChartCard from '@/components/ChartCard'
import PieChart from '@/views/dashboard/admin/components/pie-chart.vue'
import RoundChart from '@/views/dashboard/admin/components/round-chart.vue'
import LineChart from '@/views/dashboard/admin/components/line-chart.vue'
import { home } from '@/api/admin/dashboard/index'
export default {
  name: 'DashboardAdmin',
  components: {
    ChartCard,
    PieChart,
    RoundChart,
    LineChart
  },
  data() {
    return {
      deviceCount: 0,
      typeCount: 0,
      areaCount: 0,
      userCount: 0,
      typeDeviceMap: {},
      statusDeviceMap: {}
    }
  },
  mounted() {
    this.getData()
  },
  methods: {
    getData() {
      home().then(res => {
        if(res.data) {
          const {deviceCount, typeCount, areaCount, userCount, typeDeviceMap, statusDeviceMap} = res.data;
          this.deviceCount = deviceCount;
          this.typeCount = typeCount;
          this.areaCount = areaCount;
          this.userCount = userCount;
          this.typeDeviceMap = typeDeviceMap;
          this.statusDeviceMap = statusDeviceMap;
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard-editor-container {
  padding: 12px;
  background-color: rgb(240, 242, 245);
  position: relative;

  .github-corner {
    position: absolute;
    top: 0;
    border: 0;
    right: 0;
  }

  .chart-wrapper {
    background: #fff;
    padding: 16px 16px 0;
    margin-bottom: 32px;
  }
}

::v-deep .el-tabs__item{
   padding-left: 16px!important;
   height: 50px;
   line-height: 50px;
}

@media (max-width:1024px) {
  .chart-wrapper {
    padding: 8px;
  }
}

.card-title {
  padding-bottom: 20px;
  border-bottom: 1px solid #e8e8e8;
  line-height: 0;
}

.card-wrap {
  display: flex;
  align-items: center;
  img {
    width: 80px;
    height: 80px;
  }
  .info-wrap {
    flex: 1;
    margin-left: 10px;
    line-height: 0;
    .value {
      font-size: 24px;
      font-weight: bold
    }
    .title {
      font-size: 16px;
    }
  }
}
</style>
