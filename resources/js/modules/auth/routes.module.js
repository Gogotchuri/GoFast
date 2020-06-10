import SignIn from "@views/auth/SignIn.vue";
import SignUp from "@views/auth/SignUp.vue";

export default [
    {
        path: "/sign-in",
        name: "SignIn",
        component: SignIn
    },
    {
        path: "/sign-up",
        name: "SignUp",
        component: SignUp
    }
];