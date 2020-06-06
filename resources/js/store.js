import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

/*Store data.*/
const storeData = {
    modules : {}
};

const store = new Vuex.Store(storeData);
export default store;