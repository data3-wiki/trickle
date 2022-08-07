const path = require('path')
const webpack = require('webpack')

module.exports = [{
    mode: 'production',
    entry: './src/decoder.ts',
    module: {
        rules: [
            {
                test: /\.ts$/,
                use: 'ts-loader',
                exclude: /node_modules/,
            },
        ],
    },
    plugins: [
        // Fix "Buffer is not defined" error.
        new webpack.ProvidePlugin({
            Buffer: ['buffer', 'Buffer']
        })
    ],
    resolve: {
        extensions: ['.ts', '.js'],
        fallback: {
            assert: false,
        }
    },
    output: {
        path: path.resolve(__dirname, 'dist'),
        filename: 'decoder.js',
        library: {
            type: 'var',
            name: 'decoder'
        }
    }
}]