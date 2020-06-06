<template>
    <div>
        <div>
            <form @submit.prevent="authenticate">
                <div>
                    <label for="email">Email</label>

                    <div class="col-md-12">
                        <input id="email" type="email" v-model="credentials.email" required autofocus>
                    </div>
                </div>

                <div>
                    <label for="password">Password</label>

                    <div>
                        <input id="password" type="password" v-model="credentials.password" required>
                    </div>
                </div>

                <div v-if="error">
                    <div class="error">
                        Invalid credentials
                    </div>
                </div>
                <div>
                    <div>
                        <button type="submit">
                           Sign In
                        </button>
<!--                        TODO:Social auth, password forgotten-->
                    </div>
                    <social-auth-section/>
                </div>
            </form>
        </div>
    </div>
</template>

<script>
    import SocialAuthSection from "@views/auth/SocialAuthSection";
    export default {
        name: "SignIn",
        components: {SocialAuthSection},
        data(){
            return{
                credentials: {
                    email : "",
                    password : ""
                },
                error: false
            }
        },
        methods: {
            authenticate(){
                this.$store.dispatch("signIn", this.credentials)
                    .then(() => {
                        let redirectionUrl = this.$route.query.redirect;
                        this.$router.push({path : redirectionUrl || "/", query: this.$route.query});
                    })
                    .catch(() => this.error = true);
            }

        },
    }
</script>

<style scoped>
    .error{
        color:red;
        text-align: center;
    }
</style>