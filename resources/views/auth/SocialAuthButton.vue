<template>
    <span>
        <span class="btn social-auth"  v-if="provider.toLowerCase() === 'facebook'" @click="login('facebook')">
            <img src="/img/facebook-logo.png" alt="facebook"/>
            <span class="social-name">Facebook</span>
        </span>
        <span class="btn social-auth" v-else-if="provider.toLowerCase() === 'google'" @click="login('google')">
            <img src="/img/google_logo.svg" alt="google"/>
            <span class="social-name">Google</span>
        </span>
        <br v-if="authError != null">
        <span v-if="authError != null" style="color: #ff0000">
            {{authError}}
        </span>
    </span>
</template>

<script>
    export default {
        name: "SocialAuthButton",
        props: ["provider"],
        data(){
            return {
                authError : null
            }
        },
        methods: {
            async login(provider){
                window.addEventListener("message", this.onMessage, false);
                //Open empty window for starters
                const authWindow = openWindow("","Login with "+provider);
                try{
                    let providerUrl = await this.$http.POST("/auth/"+provider);
                    console.log(providerUrl);
                    providerUrl = providerUrl.data.url;
                    //Push provider redirect url into opened window
                    authWindow.location.href = providerUrl;
                }catch (reason) {
                    window.removeEventListener("message", this.onMessage);
                    authWindow.close();
                    console.error(reason);
                }
            },

            async onMessage(message) {
                //We should only trigger this for right message
                if(!message) return;
                if(message.origin !== window.origin) return;
                if (message.data.access_token == null || message.data.refresh_token == null) {
                    return;
                }
                if(message.data.error !== "") {
                    this.authError = message.data.error;
                    return;
                }
                window.removeEventListener("message", this.onMessage);
                //Sets new tokens to the state
                try {
                    await this.$store.commit("setTokens", message.data);
                    const data = await this.$http.GET("/user/details");
                    await this.$store.commit("setUser", data.data);
                    await this.$router.push({path: "/user/invoices"});
                } catch (err) {
                    console.log(err.response);
                    await this.$router.push({path: "/sign-in"});
                }
            }
        }
    }
    /**
     * Just opens a new windows with given options
     * @param url
     * @param title
     * @param  {Object} options
     * @return {Window}
     */
    function openWindow (url, title, options = {}) {
        if (typeof url === 'object') {
            options = url;
            url = '';
        }
        options = { url, title, width: 600, height: 720, ...options };
        const dualScreenLeft = window.screenLeft !== undefined ? window.screenLeft : window.screen.left;
        const dualScreenTop = window.screenTop !== undefined ? window.screenTop : window.screen.top;
        const width = window.innerWidth || document.documentElement.clientWidth || window.screen.width;
        const height = window.innerHeight || document.documentElement.clientHeight || window.screen.height;
        options.left = ((width / 2) - (options.width / 2)) + dualScreenLeft;
        options.top = ((height / 2) - (options.height / 2)) + dualScreenTop;
        const optionsStr = Object.keys(options).reduce((acc, key) => {
            acc.push(`${key}=${options[key]}`);
            return acc;
        }, []).join(',');
        const newWindow = window.open(url, title, optionsStr);
        if (window.focus)
            newWindow.focus();

        return newWindow;
    }
</script>

<style scoped>
    .social-auth > img {
        width: 22px;
        height: 22px;
        margin-left: 9px;
        margin-top: 3px
    }
    .social-name {
        display: inline-block;
        margin-left: 17px;
        position: relative;
        top: -5px;
    }

</style>