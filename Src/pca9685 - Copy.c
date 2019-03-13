
//#include "main.h"
#include "pca9685.h"
#include "myi2c.h"
#include <math.h>

void PCA9685_write(uint8_t reg,uint8_t data){
    IIC_Start();
    IIC_Send_Byte(PCA9685_adrr);//发送器件地址+写命令
    IIC_Wait_Ack();         //等待应答
    IIC_Send_Byte(reg);         //写寄存器地址
    IIC_Wait_Ack();             //等待应答
    IIC_Send_Byte(data);        //发送数据
    IIC_Wait_Ack();             //等待应答
    IIC_Stop();
}

uint8_t PCA9685_read(uint8_t reg){

	  uint8_t res;
    IIC_Start();
    IIC_Send_Byte(PCA9685_adrr); //发送器件地址+写命令
    IIC_Wait_Ack();             //等待应答
    IIC_Send_Byte(reg);         //写寄存器地址
    IIC_Wait_Ack();             //等待应答
	  IIC_Start();                
    IIC_Send_Byte(PCA9685_adrr|0X01); //发送器件地址+读命令
    IIC_Wait_Ack();             //等待应答
    res = IIC_Read_Byte(0);		//读数据,发送nACK  
    IIC_Stop();                 //产生一个停止条件
    return res;  
}

void setPWMFreq(uint8_t freq) {
   uint8_t prescale,oldmode,newmode,prescaleval;
   freq *= 0.9;  // Correct for overshoot in the frequency setting 
   prescaleval = (uint8_t)25000000/4096;
   prescaleval /= freq;
   prescaleval -= 1;
   prescale = (uint8_t)floor(prescaleval + 0.5);
                
   oldmode = PCA9685_read(PCA9685_MODE1);
   newmode = (oldmode&0x7F) | 0x10; // sleep
   PCA9685_write(PCA9685_MODE1, newmode); // go to sleep
   PCA9685_write(PCA9685_PRESCALE, prescale); // set the prescaler
//   oldmode &= 0xef;	//清除sleep位
   PCA9685_write(PCA9685_MODE1, oldmode);
   HAL_Delay(5);
   PCA9685_write(PCA9685_MODE1, oldmode | 0xa1); 
}

void setPWM(uint8_t num, uint16_t on, uint16_t off){
  PCA9685_write(LED0_ON_L+4*num,on);
  PCA9685_write(LED0_ON_H+4*num,on>>8);
  PCA9685_write(LED0_OFF_L+4*num,off);
  PCA9685_write(LED0_OFF_H+4*num,off>>8);
}

void pca9685_init(I2C_HandleTypeDef *hi2c, uint8_t address)
{
 uint8_t initStruct[2];
 uint8_t prescale = 3; // hardcoded
 HAL_I2C_Master_Transmit(hi2c, address, PCA9685_MODE1, 1, 1);
 uint8_t oldmode = 0; // hardcoded
 // HAL_I2C_Master_Receive(hi2c, address, &oldmode, 1, 1);
 uint8_t newmode = ((oldmode & 0x7F) | 0x10);
 initStruct[0] = PCA9685_MODE1;
 initStruct[1] = newmode;
 HAL_I2C_Master_Transmit(hi2c, address, initStruct, 2, 1);
 initStruct[1] = prescale;
 HAL_I2C_Master_Transmit(hi2c, address, initStruct, 2, 1);
 initStruct[1] = oldmode;
 HAL_I2C_Master_Transmit(hi2c, address, initStruct, 2, 1);
 HAL_Delay(5);
 initStruct[1] = (oldmode | 0xA1);
 HAL_I2C_Master_Transmit(hi2c, address, initStruct, 2, 1);
} 


void pca9685_pwm(I2C_HandleTypeDef *hi2c, uint8_t address, uint8_t num, uint16_t on, uint16_t off)
{
 uint8_t outputBuffer[5] = {0x06 + 4*num, on, (on >> 8), off, (off >> 8)};
 HAL_I2C_Master_Transmit(hi2c, address, outputBuffer, 5, 1);
} 
