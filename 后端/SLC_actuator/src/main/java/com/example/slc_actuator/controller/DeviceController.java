package com.example.slc_actuator.controller;

import com.example.slc_actuator.pojo.Device;
import com.example.slc_actuator.pojo.Result;
import com.example.slc_actuator.service.DeviceService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class DeviceController {
    @Autowired
    private DeviceService deviceService;

    // 修改设备状态
    public void updateDeviceStatus(int deviceid,String status) {
            deviceService.updateDevice(deviceid,status);
    }


}
