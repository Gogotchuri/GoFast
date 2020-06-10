<template>
    <nav id="navigation-bar">
        <div class="nav-container">
            <router-link :to="{name: 'Home'}">
                Home
            </router-link>
            <router-link :to="{name: 'SignIn'}">
                Sign in
            </router-link>
            <router-link :to="{name: 'SignUp'}">
                Sign up
            </router-link>
        </div>
    </nav>
</template>

<script>
    export default {
        name: "MainNav"
    };
</script>