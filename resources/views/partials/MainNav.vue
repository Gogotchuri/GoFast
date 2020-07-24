<template>
    <nav id="navigation-bar">
        <div class="nav-container">
            <router-link :to="{name: 'Home'}">
                Home
            </router-link>
            <router-link :to="{name: 'SignIn'}">
<<<<<<< HEAD
                Sign In
=======
                Sign in
            </router-link>
            <router-link :to="{name: 'SignUp'}">
                Sign up
>>>>>>> 16aab656726afa8de02689a35db1b483b018bc0f
            </router-link>
        </div>
    </nav>
</template>

<script>
    export default {
        name: "MainNav"
    };
</script>