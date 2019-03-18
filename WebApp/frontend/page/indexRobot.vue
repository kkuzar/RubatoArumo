<template>
    <v-app id="inspire">
        <v-content>
            <v-container fluid fill-height>
                <v-layout align-center justify-center>
                    <v-flex xs12 sm8 md8>
                        <v-card class="elevation-12">
                            <v-card-title primary-title>
                                <div>
                                    <div class="headline">{{ controllerName }}</div>
                                    <span class="grey--text">{{ isConnText }}</span>
                                </div>
                            </v-card-title>

                            <v-card-text>
                                <v-form>
                                    <div v-gamepad:left-analog-up.repeat="lup"
                                         v-gamepad:left-analog-down.repeat="ldown"
                                         v-gamepad:right-analog-up.repeat="MotorUp"
                                         v-gamepad:right-analog-down.repeat="MotorDown"

                                         v-gamepad:right-analog-left.repeat="StepUp"
                                         v-gamepad:right-analog-right.repeat="StepDown"

                                    >
                                        <p>Value : {{upVal}}</p>
                                        <p>Motor : {{Motor}}</p>
                                        <p>Step : {{Step}}</p>
                                    </div>
                                </v-form>
                            </v-card-text>
                            <v-card-actions>
                                <v-spacer></v-spacer>
                                <v-btn  v-gamepad:button-a.repeat="pressed" color="primary">{{ toggleLed }}</v-btn>
                            </v-card-actions>
                        </v-card>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-content>
    </v-app>
</template>

<script>
import axios from 'axios'


export default {
    name: "indexRobot",
    data: function () {
        return {
            count: 0,
            isConnText: 'Robot ARM!',
            gamPadIndex: 0,
            toggleLed: 0,
            portList: [],
            controllerName: 'Xbox GamePad Connected',
            gamepad: {},
            axisChanged: false,
            newval: 0,
            ori: "L",
            upVal: 0,
            Motor: 0,
            Step: 5,
        }
    },
    watch: {
        upVal: function (oldva, newval) {
            this.cmdServoMotor( this.Step , this.Motor  , this.ori);
        }
    },
    methods: {
        fetchSerialList: function () {
            var that = this
            axios.get('/portlist').then(function (res) {
                that.portList = res.data
            }).catch(function (err) {
                console.log(err)
            })

            console.log(this.portList);
        },
        cmdServoMotor: function (num , pos , ori) {

            var formData = new FormData();
            formData.append('num', num);
            formData.append('pos', pos);
            formData.append('ori', ori);

           axios({
               url: '/motor',
               method: "POST",
               data: formData,
           }).then(function (res) {

           }).catch(function (err) {
               console.log(err)
           })
        },
        LightLed: function (Led) {

            axios.get('/led?state=' + Led).then(function (res) {

            }).catch(function (err) {
                console.log(err)
            })
        },
        pressed() {
            if (parseInt(this.toggleLed) == 1) {
                this.toggleLed = 0
                this.LightLed(0)
            }

            if (parseInt(this.toggleLed) == 0) {
                this.toggleLed = 1
                this.LightLed(1)
            }
        },
        lup() {
            this.ori = "L";

            if (parseInt(this.upVal) == 500) {
                this.upVal = 500;
            } else {
                this.upVal ++
            }
        },
        ldown(){
            this.ori = "R";
            if (parseInt(this.upVal) == -500) {
                this.upVal = -500;
            } else {
                this.upVal --;
            }

        },
        MotorUp(){
            if (this.Motor == 5) {

            } else {
                this.Motor += 1;
            }
        },
        MotorDown(){
            if (this.Motor == 0) {

            } else {
                this.Motor -= 1;
            }
        },
        StepUp(){
            if (this.Step == 100) {

            } else {
                this.Step += 1;
            }
        },
        StepDown(){
            if (this.Step == 1) {

            } else {
                this.Step -= 1;
            }
        },
        released () {
            console.log("released!");
        }
    },
    beforeMount(){

    },
    mounted() {
        var that = this
        that.fetchSerialList();
    },
    updated() {

    },

}
</script>

<style scoped>
#inspire {
    background: url("/public/img/indexrbt.jpg") no-repeat center center fixed;
    -webkit-background-size: cover;
    -moz-background-size: cover;
    -o-background-size: cover;
    background-size: cover;
}
</style>