package com.example.slc_actuator.service;

import com.example.slc_actuator.mapper.DeviceMapper;
import com.example.slc_actuator.mapper.JobInfoMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class JobInfoService {
    @Autowired
    private JobInfoMapper jobInfoMapper;

    public int getDeviceId(int jobid) {
        // 获取设备id
        return jobInfoMapper.getDeviceId(jobid);
    }
}
