import Vue from "vue";
import Vuex from "vuex";
import Auth from "@js/modules/auth/store.module";

Vue.use(Vuex);

/*Store data.*/
const storeData = {
    modules : {
        Auth
    }
};

const store = new Vuex.Store(storeData);
export default store;