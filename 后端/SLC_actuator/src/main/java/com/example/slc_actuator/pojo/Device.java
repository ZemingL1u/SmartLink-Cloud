package com.example.slc_actuator.pojo;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Device {
    private int deviceid; //设备id
    private int typeid; // 设备类型id
    private int locid; // 位置id
    private String devicename; // 设备名称
    private String devicedata; // 设备数据
    private String status; // 设备状态
}
