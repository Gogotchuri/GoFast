<template>
    <div>
        <form @submit.prevent="register">
            <div>
                <label for="first_name">First Name</label>
                <div>
                    <input id="first_name" type="text" v-model="form.first_name" required autofocus>
                </div>
            </div>
            <div>
                <label for="last_name">Last Name</label>
                <div>
                    <input id="last_name" type="text" v-model="form.last_name" required autofocus>
                </div>
            </div>
            <div>
                <label for="email">Email</label>
                <div>
                    <input id="email" type="email" v-model="form.email" required>
                </div>
            </div>
            <div>
                <label for="password">Password</label>
                <div class="col-md-12">
                    <input id="password" type="password" v-model="form.password" required>
                </div>
            </div>
            <div>
                <label for="password-confirm">Confirm Password</label>
                <div>
                    <input id="password-confirm" type="password" v-model="password_confirmation" required>
                </div>
            </div>
            <div>
                <div>
                    <button type="submit">
                        Sign Up
                    </button>
                </div>
                <social-auth-section/>
            </div>
        </form>
    </div>
</template>

<script>
    import SocialAuthSection from "@views/auth/SocialAuthSection";
    export default {
        name: "SignUp",
        components: {SocialAuthSection},
        data(){
            return {
                form :{
                    first_name : "",
                    last_name: "",
                    email : "",
                    password : "",

                },
                password_confirmation : "",
                authErrors: null

            }
        },
        methods : {
            register(){
                if(this.noMatch){
                    window.alert("Please enter a matching password first!");
                    return;
                }
                this.$store.dispatch("signUp", this.form)
                    .then( () => {
                        window.alert("Registered successfully! Please Sign in")
                        this.$router.push({name: "SignIn"})
                    }).catch(reason => {
                        let errors = reason.response.data.errors;
                        console.error(reason.response.data.errors);
                        this.authErrors = errors;
                        console.error(errors);
                    });
            },
        },
        computed: {
            noMatch(){
                return this.form.password !== this.password_confirmation;
            }
        }
    }
</script>
