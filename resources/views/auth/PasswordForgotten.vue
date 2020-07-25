<template>
    <section>
        <form @submit.prevent="requestReset" class="form">
            <div class="form-group">
                <label for="email" class="small">Email</label>
                <br>
                <input id="email" type="email" v-model="email" required autofocus>
            </div>

            <div class="full-width">
                <input type="submit" value="Request Reset" class="btn btn-auth">
            </div>
        </form>
    </section>
</template>

<script>
    export default {
        name: "PasswordForgotten",
        data() {
            return {
                email: ""
            }
        },
        methods: {
            requestReset() {
                this.$http.POST("/forgot-password", {email: this.email})
                    .then(() => window.alert("Check email"))
                    .catch( err => console.error(err.response));
            }
        }
    }
</script>

<style scoped>

</style>