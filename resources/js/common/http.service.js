import Vue from "vue";
import Axios from "axios";
import {getToken} from "@js/common/jwt.service";
import {API_URL} from "@js/common/config";

class HttpService {
    /**
     * Given base Url new http service instance
     * This class is wrapper of Axios
     * @ApiUrl base url of api
     */
    constructor(ApiUrl) {
        //TODO: CSRF
        this._axios = Axios.create({
            baseURL: ApiUrl,
            headers: {
                "Accept" : "application/json",
            }
        });
        this._axios.defaults.headers.common["Access-Control-Allow-Origin"] = "*";
        this._axios.defaults.headers.common["accept"] = "application/json";
        this.setExistingJwtHeader();
    }

    /**
     * Given a token, sets it in Axios default headers
     * as a Bearer authorization token
     *
     * @param token jwt access token
     */
    setJwtHeader(token) {
        this._axios.defaults.headers.common['Authorization'] = 'Bearer ' + token;
    }

    /**
     * Sets already existed token in local storage as Bearer access token
     */
    setExistingJwtHeader() {
        let token = getToken();
        this._axios.defaults.headers.common['Authorization'] = 'Bearer ' + token;
    }

    /**
     * Removes Bearer jwt token from Axios headers
     * Replaces it with a junk
     */
    removeJwtHeader() {
        this._axios.defaults.headers.common['Authorization'] = 'Bearer Dummy';
    }

    /**
     * Wrapper for Axios GET request
     * @uri relative path to api route
     * @urlParams Parameters that should be passed with get request
     *
     * @return Promise with either data or error
     */
    async GET(uri, urlParams = null) {
        return new Promise((resolve, reject) => {
            this._axios.get(uri, {params: urlParams})
                .then(value => {
                    resolve(value);
                })
                .catch(reason => {
                    reject(reason);
                })
        });
    }

    /**
     * Wrapper for Axios POST request
     * @uri relative path to api route
     * @baggage Data that should be sent with this request
     *
     * @return Promise with either data or error
     */
    async POST(uri, baggage = null, headers=null) {
        let promise = (headers == null) ? this._axios.post(uri, baggage) : this._axios.post(uri, baggage, {headers: headers});
        return new Promise((resolve, reject) => {
            promise
                .then(value => {
                    resolve(value);
                })
                .catch(reason => {
                    reject(reason);
                })
        });
    }

    /**
     * Wrapper for Axios, DELETE request
     * @uri relative path to api route
     * @baggage Data that should be sent with this request
     * packs _method : delete as a hidden input to let some apis know our intent
     *
     * @return Promise with either data or error
     */
    async DELETE(uri, baggage = null) {
        if (!!baggage)
            baggage["_method"] = "DELETE";
        else
            baggage = {"_method": "DELETE"};

        return new Promise((resolve, reject) => {
            this._axios.delete(uri, baggage)
                .then(value => {
                    resolve(value);
                })
                .catch(reason => {
                    reject(reason);
                })
        });
    }

    /**
     * Wrapper for Axios, PUT request
     * @uri relative path to api route
     * @baggage Data that should be sent with this request
     * packs _method : put as a hidden input to let some apis know our intent
     *
     * @return Promise with either data or error
     */
    async PUT(uri, baggage) {
        if (!!baggage)
            baggage["_method"] = "PUT";
        else
            baggage = {"_method": "PUT"};

        return new Promise((resolve, reject) => {
            this._axios.post(uri, baggage)
                .then(value => {
                    resolve(value);
                })
                .catch(reason => {
                    reject(reason);
                })
        });
    }

    /**
     * Wrapper for Axios, PATCH request
     * @uri relative path to api route
     * @baggage Data that should be sent with this request
     * packs _method : put as a hidden input to let some apis know our intent
     *
     * @return Promise with either data or error
     */
    async PATCH(uri, baggage) {
        if (!!baggage)
            baggage["_method"] = "PATCH";
        else
            baggage = {"_method": "PATCH"};

        return new Promise((resolve, reject) => {
            this._axios.post(uri, baggage)
                .then(value => {
                    resolve(value);
                })
                .catch(reason => {
                    reject(reason);
                })
        });
    }

    getHeaders(){
        return this._axios.defaults.headers.common;
    }

    getBase(){
        return this._axios;
    }

    initializeInterceptors(store, router, progressBar){
        /**
         * Is called before every call to api
         * If user doesn't have right to access resource or
         * bearer token has been expired, we log the user out
         */
        //Request intercept for starting progress bar
        this._axios.interceptors.request.use(value => {
            progressBar.start();
            return Promise.resolve(value);
        }, error => {
            progressBar.fail();
            return Promise.reject(error);
        });
        //Response intercept
        this._axios.interceptors.response.use((value) => {
            progressBar.finish();
            return Promise.resolve(value);
        }, async (err) => {
            // console.error(err);
            if(err.response.status === 401){
                //If frontend states that user is authenticated, indicates that access token might have expired
                let originalRequest = err.config;
                //Also check if we have already retried this request
                if(store.getters.isAuthenticated && !originalRequest._rtry) {
                    originalRequest._rtry = true; //We will retry here anyway
                    //Try to renew tokens
                    let tokensRefreshed = await refreshTokens(store);
                    if (tokensRefreshed) {
                        console.log("Tokens refreshed");
                        //Return original request after tokens are refreshed
                        //If this fails we will know second try also failed and logout user
                        originalRequest.response = undefined;
                        originalRequest.headers["Authorization"] = "Bearer " + store.getters.accessToken;
                        return this._axios.request(originalRequest);
                    }
                    console.error("Refresh token expired");
                }
                store.commit("logout");
                router.push({name: "SignIn"});
            }
            progressBar.fail();
            return Promise.reject(err);
        });
    }
}

//Try to refresh token
async function refreshTokens(store) {
    let success = false;
    //Making request from default axios, to not to go through same interceptors
    Axios.defaults.headers.common["accept"] = "application/json";
    try {
        let tokens = await Axios.post(API_URL+"/token/refresh", {"refresh_token" : store.getters.refreshToken} );
        store.commit("setTokens", tokens.data);
        success = true;
    } catch (ignored) {
        success = false;
    }
    return success;
}

const httpInstance = new HttpService(API_URL);
Vue.prototype.$http = httpInstance;
export default httpInstance;
