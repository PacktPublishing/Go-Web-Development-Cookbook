var path = require('path');

module.exports = {
    resolve: {
      extensions: ['.js', '.jsx']
    },
    mode: 'development',
    entry: './app/main.js',
    cache: true,
    output: {
        path: __dirname,
        filename: './assets/script.js'
    },
    module: {
        rules: [
            {
                test: path.join(__dirname, '.'),
                exclude: /(node_modules)/,
                loader: 'babel-loader',
                query: {
                    cacheDirectory: true,
                    presets: ['es2015', 'react']
                }
            }
        ]
    }
};
