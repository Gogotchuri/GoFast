/**
 * This is a Vuex store module for auth
 * Managing registration and login currently
 */

import {destroyUser, getUser, storeUser, storeTokens} from "@js/common/jwt.service";
import Http from "@js/common/http.service";
import {changeLocale} from "@js/common/persistent_locale";

/**
 * State for auth module
 */
const localUser = getUser();

const state = {
    currentUser : localUser,
    authErrors : null
};

/**
 * Getters for auth module
 */
const getters = {
    currentUser(state){
        return state.currentUser;
    },

    userAccount(state){
        return state.currentUser.user.account;
    },

    isAuthenticated(state){
        return !!state.currentUser;
    },

    refreshToken(state){
        if(state.currentUser == null) {
            return "Dummy";
        }
        return state.currentUser.refresh_token ? state.currentUser.refresh_token : "Dummy";
    },

    accessToken(state){
        if(state.currentUser == null) {
            return "Dummy";
        }
        return state.currentUser.access_token ? state.currentUser.access_token : "Dummy";
    },

    authErrors(state){
        return state.authErrors;
    }
};

/**
 * Mutations for auth module are changing frontend app state
 * Mutations are being "commit"ed
 */
const mutations = {
    changeLocale(state, loc) {
        changeLocale(loc);
    },
    login(state, user, errors){
        if(errors){
            state.authErrors = errors;
            state.currentUser = null;
        }else if(user){
            state.authErrors = null;
            state.currentUser = user;
            console.log(user)
            console.log(user.access_token)
            Http.setJwtHeader(user.access_token);
            storeUser(user);
        }else{
            console.error("Something went Wrong during authentication");
        }
    },
    //Set new set of tokens to the existing user
    setTokens(state, tokens) {
        if(state.currentUser == null) {
            state.currentUser = {};
        }
        state.currentUser.access_token = tokens.access_token;
        state.currentUser.refresh_token = tokens.refresh_token;
        Http.setJwtHeader(tokens.access_token);
        //TODO might need to add user data fetching logic here
        storeTokens(tokens.access_token, tokens.refresh_token);
    },

    setUser(state, user) {
        state.currentUser.user = user;
        storeUser(state.currentUser);
    },

    logout(state){
        Http.removeJwtHeader();
        state.currentUser = null;
        destroyUser();
    }
};

/**
 * Actions for auth module are external functions
 * Actions are being "dispatch"ed
 */
const actions = {
    signIn(context, credentials) {
        return new Promise((resolve, reject) => {
            Http.POST("/sign-in", credentials)
                .then(value => {
                    let user = value.data;
                    context.commit("login", user, null);
                    resolve();
                })
                .catch(reason => {
                    console.error(reason);
                    context.commit("login", null, reason);
                    reject();
                })
        });
    },

    logout(context){
        return new Promise((resolve, reject) => {
            Http.POST("/user/logout")
                .then(() => {
                    context.commit("logout");
                    resolve("Logged out!");
                })
                .catch(reason => {
                    reject(reason);
                }).finally(() => context.commit("logout")  )
        });
    },
    /**
     *
     * @param context
     * @param data object {email, first_name, last_name, password}
     */
    signUp(context, data) {
        return new Promise((resolve, reject) => {
            Http.POST("/sign-up", data)
                .then(() => resolve())
                .catch(reason => reject(reason))
        });
    }
};


/**
 * Exporting auth module store
 */
export default {
    state,
    getters,
    mutations,
    actions
};
