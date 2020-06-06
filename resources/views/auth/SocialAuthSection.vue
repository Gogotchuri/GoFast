<template>
    <div>
        <social-auth-button provider="facebook" class="social-login fb-login"></social-auth-button>
        <social-auth-button provider="google" class="social-login g-login"></social-auth-button>
    </div>
</template>

<script>
    import SocialAuthButton from "@views/auth/SocialAuthButton";
    export default {
        name: "SocialAuthSection",
        components : {
            SocialAuthButton
        }
    }
</script>

<style scoped>

</style>