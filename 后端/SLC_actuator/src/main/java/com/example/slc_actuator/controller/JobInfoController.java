package com.example.slc_actuator.controller;

import com.example.slc_actuator.pojo.Result;
import com.example.slc_actuator.service.JobInfoService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class JobInfoController {
    @Autowired
    private JobInfoService jobInfoService;

    // 根据任务ID获取设备ID
    public int getDeviceId(int jobid) {
            // 调用服务层方法获取设备ID
             return jobInfoService.getDeviceId(jobid);
    }
}
