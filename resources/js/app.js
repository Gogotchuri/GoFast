import Vue from 'vue';
import VueRouter from 'vue-router';
import Vuex from 'vuex';

import store from '@js/store';
import App from '@views/App';
import router from '@js/router';

Vue.use(VueRouter);
Vue.use(Vuex);

export const app = new Vue({
	el: '#app',
	store,
	components: { App },
	router
});a