//const { defineConfig } = require('@vue/cli-service')
// module.exports = defineConfig({
//   transpileDependencies: true,
//
//   pluginOptions: {
//     vuetify: {
// 			// https://github.com/vuetifyjs/vuetify-loader/tree/next/packages/vuetify-loader
// 		}
//   },
//
//     devServer: {
//         port: 8080,
//         proxy: {
//             "^/task": {
//                 target: "http://api:8081",
//                 changeOrigin: true,
//                 secure:false,
//                 logLevel: 'debug',
//                 pathRewrite: { "^/task": "/task" }
//             }
//         }
//     }
// })

module.exports = {
    publicPath: './',
    devServer: {
        port: 8080,
        proxy: {
            "^/task": {
                target: "http://api:8081",
                changeOrigin: true,
                secure: false,
                logLevel: 'debug',
                pathRewrite: {"^/task": "/task"}
            }
        }
    },
    pluginOptions: {
    vuetify: {
			// https://github.com/vuetifyjs/vuetify-loader/tree/next/packages/vuetify-loader
		}
    },
}
