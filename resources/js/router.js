import VueRouter from "vue-router";

import PublicRoutes from "@js/modules/public/routes.module";
import AuthRoutes from "@js/modules/auth/routes.module";

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

//TODO:SEO Optimization
export default router;