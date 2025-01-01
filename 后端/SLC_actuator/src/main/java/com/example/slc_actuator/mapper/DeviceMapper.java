package com.example.slc_actuator.mapper;

import com.example.slc_actuator.pojo.Device;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

@Mapper
public interface DeviceMapper {
    // 根据设备id更改设备状态
    @Update("UPDATE device SET status = #{status} WHERE device_id = #{deviceid}")
    void updateDevice(int deviceid,String status);



}
