<template>
    <header>
        <div class="logo">
            GoFast
        </div>
        <nav>
            <ul class="nav__links">
                <li>
                    <router-link :to="{name: 'Home'}">
                        Home
                    </router-link>
                </li>
                <li>
                    <router-link v-if="!isLoggedIn" :to="{name: 'SignIn'}">
                        Sign in
                    </router-link>
                </li>
                <li>
                    <router-link v-if="!isLoggedIn" :to="{name: 'SignUp'}">
                        Sign up
                    </router-link>
                </li>
                <li>
                    <router-link v-if="isLoggedIn" :to="{name: 'Logout'}">
                        Logout
                    </router-link>
                </li>
            </ul>
        </nav>
    </header>
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

<style scoped>
* {
    box-sizing: border-box;
    margin: 0;
    padding: 0;
    background-color: #24252A;
}

li, a, button {
    font-family: sans-serif;
    font-weight: 450;
    font-size: 15px;
    color: white;
    text-decoration: none;
}

header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 20px 10%;
}

.nav__links {
    list-style: none;
}

.nav__links li {
    display: inline-block;
    padding: 0px 20px;
}

.nav__links li a {
    transition: all 0.3s ease 0s;
}

.nav__links li a:hover {
    color: #0088a9
}

button {
    padding: 9px 25px;
    background-color: rgba(0,136,169,1);
    border: none;
    border-radius: 50px;
    cursor: pointer;
    transition: all 0.3s ease 0s;
}

button:hover {
    background-color: rgba(0,136,169,0.8);
}

.logo {
    color: white;
    font-family: sans-serif;
    font-weight: 700;
    font-size: 35px;
}
</style>
