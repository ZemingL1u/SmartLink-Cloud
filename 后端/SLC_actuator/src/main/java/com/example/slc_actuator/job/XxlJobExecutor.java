package com.example.slc_actuator.job;
import com.example.slc_actuator.controller.DeviceController;
import com.example.slc_actuator.controller.JobInfoController;
import com.example.slc_actuator.pojo.Result;
import com.xxl.job.core.context.XxlJobHelper;
import com.xxl.job.core.handler.annotation.XxlJob;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;



@Component
public class XxlJobExecutor {
    @Autowired
    private DeviceController deviceController;
    @Autowired
    private JobInfoController jobInfoController;
    @XxlJob("turn_on")
    public void turnOn() {
        long jobId= XxlJobHelper.getJobId();
        if (jobId >= Integer.MIN_VALUE && jobId <= Integer.MAX_VALUE) {
            int jobIdAsInt = (int) jobId;
            // 现在可以使用 jobIdAsInt 作为 int 类型的变量
            int deviceId = jobInfoController.getDeviceId(jobIdAsInt);
            deviceController.updateDeviceStatus(deviceId,"在线");
        } else {
            // jobId 的值超出了 int 类型的范围，需要进行错误处理
            throw new ArithmeticException("jobId value is out of int range");
        }
    }

    @XxlJob("turn_off")
    public void turnOff() {
        long jobId= XxlJobHelper.getJobId();
        if (jobId >= Integer.MIN_VALUE && jobId <= Integer.MAX_VALUE) {
            int jobIdAsInt = (int) jobId;
            // 现在可以使用 jobIdAsInt 作为 int 类型的变量
            int deviceId = jobInfoController.getDeviceId(jobIdAsInt);
            deviceController.updateDeviceStatus(deviceId,"待机");
        } else {
            // jobId 的值超出了 int 类型的范围，需要进行错误处理
            throw new ArithmeticException("jobId value is out of int range");
        }
    }

    @XxlJob("get_sensor_info")
    public void getInfo() {
       System.out.println("读取数据");
    }

}
