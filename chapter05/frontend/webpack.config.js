module.exports = {
    mode: "production",
    entry: "./src/index.tsx",
    output: {
        filename: "bundle.js",
        path: __dirname + "/dist"
    },

    devtool: "source-map",

    resolve: {
        extensions: [".ts", ".tsx", ".css", ".js"]
    },

    module: {
        /*loaders: [
            {
                test: /\.tsc?$/,
                loader: "awesome-typescript-loader"
            }
        ],*/

        rules: [
            {
                test: /\.tsx?$/,
                //loader: "awesome-typescript-loader"
                use: ["ts-loader"]
            },
            {
                test: /.jsx?$/,
                use: ['babel-loader']
            },
            {
                test: /\.css$/,
                use: ["style-loader", "css-loader"]
            },
            {
                test: /\.(woff|woff2)(\?v=\d+\.\d+\.\d+)?$/,
                use: ['url-loader?limit=10000&mimetype=application/font-woff']
            },
            {
                test: /\.ttf(\?v=\d+\.\d+\.\d+)?$/,
                use: ['url-loader?limit=10000&mimetype=application/octet-stream']
            },
            {
                test: /\.eot(\?v=\d+\.\d+\.\d+)?$/,
                use: ['file-loader']
            },
            {
                test: /\.svg(\?v=\d+\.\d+\.\d+)?$/,
                use: ['url-loader?limit=10000&mimetype=image/svg+xml']
            }
            /*,
            {
                test: /\.js$/,
                enforce: "pre",
                loader: "source-map-loader"
            }*/
        ]
    },

    externals: {
        "react": "React",
        "react-dom": "ReactDOM"
    }
};