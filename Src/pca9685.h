#ifndef __PCF9685_H
#define __PCF9685_H

#include "main.h"

#define PCA9685_adrr 0x80//  1+A5+A4+A3+A2+A1+A0+w/r 
                                           //片选地址，将焊接点置1可改变地址，
                                            //        当IIC总线上有多片PCA9685或相同地址时才需焊接
#define PCA9685_SUBADR1 0x2
#define PCA9685_SUBADR2 0x3
#define PCA9685_SUBADR3 0x4


#define PCA9685_MODE1 0x0
#define PCA9685_PRESCALE 0xFE


#define LED0_ON_L 0x6
#define LED0_ON_H 0x7
#define LED0_OFF_L 0x8
#define LED0_OFF_H 0x9


#define ALLLED_ON_L 0xFA
#define ALLLED_ON_H 0xFB
#define ALLLED_OFF_L 0xFC
#define ALLLED_OFF_H 0xFD

#define SERVOMIN  115 // this is the 'minimum' pulse length count (out of 4096)
#define SERVOMAX  590 // this is the 'maximum' pulse length count (out of 4096)
#define SERVO000  170 //0度对应4096的脉宽计数值
#define SERVO180  55  //180度对应4096的脉宽计算值，四个值可根据不同舵机修改

void PCA9685_Reset();
void PCA9685_Go();

void PCA9685_write(uint8_t reg,uint8_t  data);
uint8_t PCA9685_read(uint8_t reg);
void setPWMFreq(float freq);
void setPWM(uint8_t num, uint16_t on, uint16_t off);

void pca9685_init(I2C_HandleTypeDef *hi2c, uint8_t address);
void pca9685_pwm(I2C_HandleTypeDef *hi2c, uint8_t address, uint8_t num, uint16_t on, uint16_t off);
#endif
