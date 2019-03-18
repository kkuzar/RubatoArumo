import vuetify from 'vuetify';
import axios  from 'axios'
import VueClip from 'vue-clip'
import VueGamepad from 'vue-gamepad'


window.Vue = require('vue');


Vue.component('indexRobot', require('./page/indexRobot'));
Vue.use(vuetify);
Vue.use(VueClip);
Vue.use(axios);
Vue.use(VueGamepad, {
    analogThreshold: 0.5,
    buttonRepeatTimeout: 7500,
});

const app = new Vue({
    el: '#app',
});
