import VueRouter from "vue-router";

import PublicRoutes from "@js/modules/public/routes.module";
import AuthRoutes from "@js/modules/auth/routes.module";

import store from "@js/store";
/**
 * Concatenating routes from different modules
 */
let routes = []
    .concat(PublicRoutes)
    .concat(AuthRoutes);

//Creating and exporting Vue router instance
export const router = new VueRouter({
    mode: 'history',
    routes
});

/**
 * Getting auth required property and redirecting user properly
 */
router.beforeEach(async (to, from, next) => {
    const authRequired = to.matched.some(record => record.meta.authRequired);
    const user = store.getters.currentUser;
    if(authRequired && !user && to.path !== "/logout") {
        //Using Get request query param to redirect after
        // Redirection to login cause of unauthorized request
        next({name: "SignIn", query: {redirect: to.path}});
    }else if ((to.path === "/sign-in" || to.path === "/sign-up") && !!user) {
        //Can't sign in if you already are
        next({path:"/"});
    } else {
        next();
    }

});

//TODO:SEO Optimization
export default router;