<template>
    <section>
        <form @submit.prevent="performReset" class="form">
            <div class="form-group">
                <label for="password">Password</label>
                <br>
                <input id="password" type="password" v-model="password" required>
            </div>
            <div class="form-group">
                <label for="password-confirm">Confirm password</label>
                <br>
                <input id="password-confirm" type="password" v-model="password_confirmation" required>
            </div>
            <div class="full-width">
                <input type="submit" value="Reset" class="btn btn-auth">
            </div>
        </form>
    </section>
</template>

<script>
    export default {
        name: "PasswordReset",
        data() {
            return {
                password: "",
                password_confirmation: ""
            }
        },
        computed: {
            noMatch(){
                return this.password !== this.password_confirmation;
            }
        },
        methods: {
            performReset() {
                if(this.noMatch) return;
                this.$http.POST("/password-reset", {password: this.password, token: this.$route.query.token})
                    .then(() => this.$router.push({name: "SignIn"}))
                    .catch(() => window.alert(":("))
            }
        }
    }
</script>

<style scoped>

</style>