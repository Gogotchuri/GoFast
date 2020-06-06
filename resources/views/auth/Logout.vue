<template>
    <div>
        <p v-if="errors === null">Successfully logged out!</p>
        <p v-else>Couldn't logout</p>
    </div>
</template>

<script>
    import store from "@js/store";
    export default {
        name: "Logout",
        data() {
            return {
                errors: null,
            };
        },

        beforeRouteEnter(to, from, next){
            const Auth = store.getters.isAuthenticated;
            if(!Auth){
                next({name: "Home"});
            }else{
                next();
            }
        },

        beforeMount() {
            this.logout();
        },

        mounted(){
            const redirect = () => this.$router.push({name: "Home"});
            setTimeout(redirect, 5);
        },
        methods: {
            logout(){
                this.$store.dispatch("logout")
                    .catch(err => {
                        this.errors = err;
                        console.error(err.response);
                    });
            }
        }

    }
</script>