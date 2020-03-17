const path = require('path');
const LiveReloadPlugin = require('webpack-livereload-plugin');
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const UglifyJsPlugin = require("uglifyjs-webpack-plugin");
const OptimizeCSSAssetsPlugin = require("optimize-css-assets-webpack-plugin");
const devMode = process.env.NODE_ENV !== 'production';
var CopyWebpackPlugin = require('copy-webpack-plugin');
const MomentLocalesPlugin = require('moment-locales-webpack-plugin');

module.exports = {
    entry: './client/boot.js',
    output: {
        filename: 'main.js',
        path: path.resolve(__dirname, 'public/assets')
    },
    devtool: "source-map",
    watchOptions: {
        aggregateTimeout: 300,
        ignored: /node_modules/,
    },
    optimization: {
        usedExports: true,
        minimizer: [
            new UglifyJsPlugin({
                cache: true,
                parallel: true,
                sourceMap: true // set to true if you want JS source maps
            }),
            new OptimizeCSSAssetsPlugin({})
        ]
    },
    resolve: {
        extensions: [
            ".js"
        ],
        modules: [
            path.resolve(__dirname, "scripts"),
            "node_modules"
        ]
    },
    module: {
        rules: [{
            test: /\.js$/,
            exclude: /node_modules/,
            use: {
                loader: 'babel-loader',
                options: {
                    presets: ['@babel/preset-env']
                }
            }
        },
        {
            test: /\.woff(2)?(\?v=[0-9]\.[0-9]\.[0-9])?$/,
            loader: "url-loader?limit=10000&mimetype=application/font-woff"
        },
        {
            test: /\.(ttf|eot|svg)(\?v=[0-9]\.[0-9]\.[0-9])?$/,
            loader: "file-loader"
        },
        {
            test: /\.css$/,
            use: [MiniCssExtractPlugin.loader, "css-loader"],
        },
        {
            test: /\.scss$/,
            use: [
                //devMode ? 'style-loader' : MiniCssExtractPlugin.loader,
                MiniCssExtractPlugin.loader,
                'css-loader',
                'sass-loader',
            ],
        },
        {
            test: /\.(png|jpg|jpeg|gif|svg)$/,
            use: 'url-loader?limit=25000'
        },

        ]
    },

    plugins: [
        // To strip all locales except “en”. See https://momentjs.com/docs/ how for other locales
        new MomentLocalesPlugin(),

        new LiveReloadPlugin({
            appendScriptTag: devMode
        }),
        new MiniCssExtractPlugin({
            filename: "[name].css",
            chunkFilename: "[id].css"
        }),
        new CopyWebpackPlugin([{
            from: 'client/images',
            to: '../images'
        },
        {
            from: 'client/vendor',
            to: '../vendor'
        },
        ]),
    ]
};
