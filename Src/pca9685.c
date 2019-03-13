
//#include "main.h"
#include "pca9685.h"
#include <math.h>

extern I2C_HandleTypeDef hi2c1;

void PCA9685_write(uint8_t reg,uint8_t data){
    //Send address to start reading from.
    uint8_t tx[2];
    tx[0]=reg;
    tx[1]=data;
    HAL_I2C_Master_Transmit(&hi2c1,PCA9685_adrr, tx,2,10000);
}

uint8_t PCA9685_read(uint8_t reg){

//Send address to start reading from.
    uint8_t tx[1];
    uint8_t buffer[1];
    tx[0]=reg;

    HAL_I2C_Master_Transmit(&hi2c1,PCA9685_adrr, tx,1,10000);
    HAL_I2C_Master_Receive(&hi2c1,PCA9685_adrr,buffer,1,10000);
    return buffer[0];
}

void setPWMFreq(float freq) {
   uint32_t prescale,oldmode,newmode;
    float prescaleval;
    freq *= 0.92;  // Correct for overshoot in the frequency setting 
    prescaleval = 25000000;
    prescaleval /= 4096;
    prescaleval /= freq;
    prescaleval -= 1;
    prescale = floor(prescaleval + 0.5);

    oldmode = PCA9685_read(PCA9685_MODE1);
    newmode = (oldmode&0x7F) | 0x10; // sleep
    PCA9685_write(PCA9685_MODE1, newmode); // go to sleep
    PCA9685_write(PCA9685_PRESCALE, prescale); // set the prescaler
    PCA9685_write(PCA9685_MODE1, oldmode);
    PCA9685_write(PCA9685_MODE1, oldmode | 0xa1);
}

void setPWM(uint8_t num, uint16_t on, uint16_t off){
    PCA9685_write(LED0_ON_L+4*num,on);
    PCA9685_write(LED0_ON_H+4*num,on>>8);
    PCA9685_write(LED0_OFF_L+4*num,off);
    PCA9685_write(LED0_OFF_H+4*num,off>>8);
}

void PCA9685_Reset()
{
    PCA9685_write(PCA9685_MODE1,0x00);
}

void PCA9685_Go()
{
    PCA9685_Reset();
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
