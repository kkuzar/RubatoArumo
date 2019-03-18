let mix = require('laravel-mix');

/*
 |--------------------------------------------------------------------------
 | Mix Asset Management
 |--------------------------------------------------------------------------
 |
 | Mix provides a clean, fluent API for defining some Webpack build steps
 | for your Laravel application. By default, we are compiling the Sass
 | file for the application as well as bundling up all the JS files.
 |
 */
mix.setPublicPath('public');


mix.js('frontend/index.js', 'public/js');
mix.js('frontend/admin_login.js', 'public/js');
mix.js('frontend/index_robot.js', 'public/js');
// mix.js('resources/assets/js/ai.js', 'public/js')
//     .sass('resources/assets/sass/ai.scss', 'public/css');
