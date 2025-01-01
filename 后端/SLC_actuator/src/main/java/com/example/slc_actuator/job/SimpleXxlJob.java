package com.example.slc_actuator.job;

import com.xxl.job.core.handler.annotation.XxlJob;
import org.springframework.stereotype.Component;

import java.util.Date;

@Component
public class SimpleXxlJob {
    @XxlJob("simpleJobHandler")
    public void simpleJobHandler() {
        System.out.println("执行定时任务，执行时间："+new Date());
    }
}
