import SignIn from "@views/auth/SignIn";
import SignUp from "@views/auth/SignUp";
import PasswordForgotten from "@views/auth/PasswordForgotten";
import PasswordReset from "@views/auth/PasswordReset";
import Logout from "@views/auth/Logout";
import EmailVerification from "@views/auth/EmailVerification";

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
        path: "/password-forgotten",
        name: "PasswordForgotten",
        component: PasswordForgotten
    },
    {
        path: "/password-reset",
        name: "PasswordReset",
        component: PasswordReset
    },
    {
        path: "/logout",
        name: "Logout",
        component: Logout,
        meta: {
            authRequired: true
        }
    },
    {
        path: "/verify-email",
        name: "EmailVerification",
        component: EmailVerification
    },
];