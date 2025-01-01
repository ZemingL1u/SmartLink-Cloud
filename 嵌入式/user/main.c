
//单片机头文件
#include "stm32f10x.h"

//网络协议层
#include "onenet.h"

//网络设备
#include "esp8266.h"

//硬件驱动
#include "delay.h"
#include "usart.h"
#include "led.h"
#include "key.h"
#include "dht11.h"
#include "lcd_init.h"
#include "lcd.h"
#include "pic.h"
//C库
#include <string.h>

//#define ESP8266_ONENET_INFO		"AT+CIPSTART=\"TCP\",\"mqtts.heclouds.com\",1883\r\n"
#define ESP8266_ONENET_INFO		"AT+CIPSTART=\"TCP\",\"mqtts.heclouds.com\",1883\r\n"
void Hardware_Init(void);
void Display_Init(void);
void Refresh_Data(void);

u8 temp,humi;


/*
************************************************************
*	函数名称：	main
*
*	函数功能：	
*
*	入口参数：	无
*
*	返回参数：	0
*
*	说明：		
************************************************************
*/
int main(void)
{
	
	unsigned short timeCount = 0;	//发送间隔变量
	unsigned char *dataPtr = NULL;
	Hardware_Init();				//初始化外围硬件

	ESP8266_Init();					//初始化ESP8266
	LCD_Fill(0,0,160,160,WHITE);
  UsartPrintf(USART_DEBUG, "Connect MQTTs Server...\r\n");
	
	while(ESP8266_SendCmd(ESP8266_ONENET_INFO, "CONNECT"))
		DelayXms(500);
	LCD_ShowString(10,50,"Device Login...", DARKBLUE ,WHITE,16,0);
	
	while(OneNet_DevLink())			//接入OneNET
	{
		ESP8266_SendCmd(ESP8266_ONENET_INFO, "CONNECT");
		DelayXms(500);
	}
	LCD_ShowString(10,80,"Login Success!", DARKBLUE ,WHITE,16,0);
	
	OneNET_Subscribe();
	Display_Init();
	while(1)
	{

		if(++timeCount >= 20)									//发送间隔5s
		{
			DHT11_Read_Data(&temp,&humi);
			OneNet_SendData();									//发送数据
			
			timeCount = 0;
			ESP8266_Clear();
		}
		
		dataPtr = ESP8266_GetIPD(0);
		if(dataPtr != NULL)
			OneNet_RevPro(dataPtr);
		
		Refresh_Data();
		
		DelayMs(10);
	
	}

}
/*
************************************************************
*	函数名称：	Hardware_Init
*
*	函数功能：	硬件初始化
*
*	入口参数：	无
*
*	返回参数：	无
*
*	说明：		初始化单片机功能以及外接设备
************************************************************
*/
void Hardware_Init(void)
{
	
	NVIC_PriorityGroupConfig(NVIC_PriorityGroup_2);	//中断控制器分组设置

	Delay_Init();									//systick初始化
	
	Usart1_Init(115200);							//串口1，打印信息用
	
	Usart2_Init(115200);							//串口2，驱动ESP8266用
	
  Key_Init();
	Led_Init();
  LCD_Init();//LCD初始化
	LCD_Fill(0,0,160,160,WHITE);
	//LCD_ShowString(15,15,"LCD Init End.", DARKBLUE ,WHITE,16,0);
	UsartPrintf(USART_DEBUG, "LCD Init OK.\r\n");
	
	while(DHT11_Init())
	{
		UsartPrintf(USART_DEBUG, "DHT11 Error \r\n");
		DelayMs(1000);
	}
	UsartPrintf(USART_DEBUG, " Hardware init OK\r\n");
}

void Display_Init(void)
{
		LCD_Fill(0,0,160,160,WHITE);
		LCD_ShowString(20,20,"Temp",DARKBLUE,WHITE,16,0);
		LCD_ShowString(20,45,"Humi",DARKBLUE,WHITE,16,0);
		LCD_ShowString(20,70,"LED ",DARKBLUE,WHITE,16,0);
//		LCD_ShowString(45,15,":", DARKBLUE ,WHITE,16,0);
//		LCD_ShowString(45,35,":", DARKBLUE ,WHITE,16,0);

		LCD_ShowString(97,20,"C",DARKBLUE,WHITE,16,0);
		LCD_ShowString(100,45,"%",DARKBLUE,WHITE,16,0);
}
void Refresh_Data(void)
{
	LCD_ShowIntNum(55,20,temp,4, DARKBLUE ,WHITE,16);
	
	LCD_ShowIntNum(55,45,humi,4,DARKBLUE,WHITE,16);
	
	LCD_ShowString(70,70,led_info.Led_Status ? "ON  " : "OFF",DARKBLUE,WHITE,16,0);
//	if(led_info.Led_Status) OLED_ShowCHinese(54,6,8);//亮
//	else OLED_ShowCHinese(54,6,9);//灭
	
}
