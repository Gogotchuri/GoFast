class WebPackConf {
    webpackConfig(webpackConfig){
        webpackConfig.resolve.alias = {
            "vue$": "vue/dist/vue.esm.js",
            "@": __dirname + "/resources",
            "@js": __dirname + "/resources/js",
            "@views": __dirname + "/resources/views",
            "@lang": __dirname + "/resources/lang"
        };
    }
}

const mix = require("laravel-mix");
mix.extend("customConfig", new WebPackConf);
mix.customConfig();

mix.js("resources/js/app.js", "public/js")
   .sass("resources/sass/app.scss", "public/css");