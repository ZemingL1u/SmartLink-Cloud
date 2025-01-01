package com.example.slc_actuator.service;
import com.example.slc_actuator.mapper.DeviceMapper;
import com.example.slc_actuator.pojo.Device;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class DeviceService {

    @Autowired
    private DeviceMapper deviceMapper;

    public void updateDevice(int deviceid,String status) {
        // 更新设备信息
        deviceMapper.updateDevice(deviceid,status);
    }

}

