import Vue from 'vue';
import VueRouter from 'vue-router';
import Vuex from 'vuex';
import VueProgressBar from 'vue-progressbar';
import VueI18n from "vue-i18n";
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'

import HTTP from '@js/common/http.service';
import store from '@js/store';
import App from '@views/App';
import router from '@js/router';
import i18n from "@js/locale";

Vue.use(VueRouter);
Vue.use(Vuex);
Vue.use(VueI18n);
Vue.use(BootstrapVue);
Vue.use(IconsPlugin);
Vue.use(VueProgressBar,
{
	thickness: "3px",
	transition: {
		speed: '0.2s',
		opacity: '1s',
		termination: 300
	}
});

export const app = new Vue({
	el: '#app',
	store,
	components: { App },
	router,
	i18n
});
HTTP.initializeInterceptors(store, router, app.$Progress);