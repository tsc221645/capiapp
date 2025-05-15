const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true
})
// capiapp/vue.config.js
module.exports = {
  devServer: {
    host: '0.0.0.0',
    port: 5173
    
  }
}
