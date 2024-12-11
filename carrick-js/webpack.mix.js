let mix = require('laravel-mix');
// mix.js('src/app.js', 'dist/tracking.js')
// mix.js('src/skimlinks/app.js', 'dist/skimlinks/init.js')
// mix.js('src/amazon/app.js', 'dist/skimlinks/dog-gear-tracking.js')
mix.js('src/amazon/app.js', 'dist/reviewed/tracking.js')
    .setPublicPath('dist')
    .webpackConfig({
        devServer: {
            port: 5020,
            static: './examples/'
        }
    });
