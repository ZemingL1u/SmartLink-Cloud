package mqtt

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"strings"
	"time"
)

// MessageRecord 用于保存消息和接收时间
type MessageRecord struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

// MessageCollector 用于收集每小时的消息
type MessageCollector struct {
	Channel   chan MessageRecord
	StartTime time.Time
}

// NewMessageCollector 创建一个新的消息收集器
func NewMessageCollector(startTime time.Time) *MessageCollector {
	return &MessageCollector{
		Channel:   make(chan MessageRecord, 1000),
		StartTime: startTime,
	}
}

// CollectMessage 收集消息
func (c *MessageCollector) CollectMessage(msg string) {
	record := MessageRecord{
		Timestamp: time.Now(),
		Message:   msg,
	}
	c.Channel <- record
}

// FlushData 每小时调用一次，上传数据到OSS
func (c *MessageCollector) FlushData(client oss.Client, bucketName string) {
	hourPassed := time.Since(c.StartTime).Hours()
	if hourPassed >= 1 {
		// 聚合数据
		records := make([]MessageRecord, 0)
		for {
			select {
			case record := <-c.Channel:
				records = append(records, record)
			default:
				break
			}
		}

		// 转换为JSON
		jsonData, err := json.Marshal(records)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}

		// 上传到OSS
		objectKey := fmt.Sprintf("mqtt_data_%s.json", c.StartTime.Format("20060102_150405"))
		bucket, err := client.Bucket(bucketName)
		if err != nil {
			fmt.Println("Error getting bucket:", err)
			return
		}
		err = bucket.PutObject(objectKey, strings.NewReader(string(jsonData)))
		if err != nil {
			fmt.Println("Error uploading object:", err)
			return
		}
		fmt.Printf("Data uploaded successfully to %s\n", objectKey)

		// 重置开始时间
		c.StartTime = time.Now().Add(-1 * time.Duration(hourPassed) * time.Hour)
	}
}
