<template>
    <div class="login-form">
        <form @submit.prevent="register">
            <h1>Sign up</h1>
            <div class="form-group">
                <input id="first_name" type="text" placeholder="First Name" v-model="form.first_name" required autofocus>
            </div>
            <div class="form-group">
                <input id="last_name" type="text" placeholder="Last Name" v-model="form.last_name" required>
            </div>
            <div class="form-group">
                <input id="email" type="email" placeholder="Email" v-model="form.email" required>
            </div>
            <div class="form-group">
                <input id="password" type="password" placeholder="Password" v-model="form.password" required>
            </div>
            <div class="form-group">
                <input id="password-confirm" type="password" placeholder="Password" v-model="password_confirmation" required>
            </div>
            <button type="submit" class="login-btn">
                Sign Up
            </button>
            <div class="seperator"><b>or</b></div>
            <p>Sign in with your social media account</p>

            <div class="soc-auth">
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
                authErrors: null,
                registrationSuccess: false
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


<style scoped>
body {
    background: #607D8B;
    font-family: arial;
}
.login-form h1 {
    font-size: 36px;
    text-align: center;
    color: #45aba6;
    margin-bottom: 30px;
    font-weight: normal;
}
.login-form .social-icon {
    width: 100%;
    font-size: 20px;
    padding-top: 20px;
    color: #fff;
    text-align: center;
    float: left;
}
.login-form {
    background: #fff;
    width: 450px;
    border-radius: 6px;
    margin: 0 auto;
    display: table;
    padding: 15px 30px 30px;
    box-sizing: border-box;
}
.form-group {
  float: left;
  width: 100%;
  margin: 0 0 15px;
  position: relative;
}
.login-form input {
    width: 100%;
    padding: 5px;
    height: 56px;
    border-radius: 74px;
    border: 1px solid #ccc;
    box-sizing: border-box;
    font-size: 15px;
    padding-left: 75px;
}
.login-form .form-group .input-icon {
    font-size: 15px;
    display: -webkit-box;
    display: -webkit-flex;
    display: -moz-box;
    display: -ms-flexbox;
    display: flex;
    align-items: center;
    position: absolute;
    border-radius: 25px;
    bottom: 0;
    height: 100%;
    padding-left: 35px;
    color: #666;
}
.login-form .login-btn {
    background: #45aba6;
    padding: 11px 50px;
    border-color: #45aba6;
    color: #fff;
    text-align: center;
    margin: 0 auto;
    font-size: 20px;
    border: 1px solid #45aba6;
    border-radius: 44px;
    width: 100%;
    height: 56px;
    cursor: pointer;
}
.login-form .reset-psw {
    float: left;
    width: 100%;
    text-decoration: none;
    color: #45aba6;
    font-size: 14px;
    text-align: center;
    margin-top: 11px;
}
.login-form button:hover{
    opacity: 0.9;
}
.login-form .seperator {
    float: left;
    width: 100%;
    border-top: 1px solid #ccc;
    text-align: center;
    margin: 50px 0 0;
}
.login-form .seperator b {
    width: 40px;
    height: 40px;
    font-size: 16px;
    text-align: center;
    line-height: 40px;
    background: #fff;
    display: inline-block;
    border: 1px solid #e0e0e0;
    border-radius: 50%;
    position: relative;
    top: -22px;
    z-index: 1;
}
.login-form p {
    float: left;
    width: 100%;
    text-align: center;
    font-size: 16px;
    margin: 0 0 10px;
}
@media screen and (max-width:767px) {
    .login-form {
        width: 90%;
        padding: 15px 15px 30px;
    }
}
.soc-auth {
    text-align: center;
}
</style>
