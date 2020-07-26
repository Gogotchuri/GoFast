<template>
    <div>
        <form @submit.prevent="verify">
            <div>
                <label for="code">Verification code</label>

                <div class="col-md-12">
                    <input id="code" type="code" v-model="code" required autofocus>
                </div>
            </div>
            <div>
                <div>
                    <button type="submit">
                        Verify
                    </button>
                </div>
            </div>
        </form>
        <div slot="footer">
            <input type="button" class="btn btn-green px-2" value="Send verification mail" @click="resendVerification">
        </div>
    </div>
</template>

<script>
    import HTTP from "@js/common/http.service";

    export default {
        name: "EmailVerification",
        data() {
            return {
                success: false,
                verification_sent: false,
                code: "",
            }
        },
        methods: {
            resendVerification() {
                this.$http.POST("/user/resend-code")
                    .then(() => this.verification_sent = true)
                    .catch(() => this.verification_sent = false)
            },
            async verify() {
                try {
                    this.$http.POST("/user/verify-email", {"otac": this.code})
                    .then(() => {
                        this.success = true;
                        let curUser = this.$store.getters.currentUser.user;
                        curUser.email_verified = true;
                        this.$store.commit("setUser", curUser);
                        window.alert("Email verified!");
                        this.$router.push({name: "Home"});
                    })
                    .catch(()=>
                        window.alert("Verification failed!")
                    );
                } catch (e) {
                    console.error(e.response);
                    this.success = false;
                    window.alert("Verification failed!")
                }
            }
        }
    }
</script>