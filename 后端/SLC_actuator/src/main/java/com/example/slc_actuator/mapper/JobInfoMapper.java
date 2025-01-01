package com.example.slc_actuator.mapper;

import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;

@Mapper
public interface JobInfoMapper {

    // 根据任务id查找设备id
    @Select("SELECT device_id FROM jobinfo WHERE job_id = #{jobid}")
    int getDeviceId(int jobid);


}
