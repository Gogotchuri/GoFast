<template>
    <nav id="navigation-bar">
        <div class="nav-container">
            <router-link :to="{name: 'Home'}">
                Home
            </router-link>
            <router-link v-if="!isLoggedIn" :to="{name: 'SignIn'}">
                Sign in
            </router-link>
            <router-link v-if="!isLoggedIn" :to="{name: 'SignUp'}">
                Sign up
            </router-link>
            <router-link v-if="isLoggedIn" :to="{name: 'Logout'}">
                Logout
            </router-link>
        </div>
    </nav>
</template>

<script>
    export default {
        name: "MainNav",
        computed: {
            isLoggedIn() {
                return this.$store.getters.isAuthenticated;
            }
        }
    };
</script>