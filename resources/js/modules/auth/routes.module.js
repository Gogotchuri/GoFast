import SignIn from "@views/auth/SignIn";
import SignUp from "@views/auth/SignUp";
import Logout from "@views/auth/LogOut";


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
    },
    {
        path: "/logout",
        name: "Logout",
        component: Logout,
        meta: {
            authRequired: true
        }
    },
];