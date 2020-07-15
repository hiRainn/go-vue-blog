module.exports = {
  publicPath: '/',
  //编译打包存放的目录默认dist
  outputDir: 'dist',

  // 如果你不需要使用eslint，把lintOnSave设为false即可
  lintOnSave: false,

  // 设为false打包时不生成.map文件
  productionSourceMap: false,

  assetsDir: 'static',
  
  indexPath:'index.html',
  
  filenameHashing:false,
  pages:undefined,
  runtimeCompiler:false,
  chainWebpack: config =>  {
    config.module
      .rule('vue')
      .use('vue-loader')
      .loader('vue-loader')
      .tap(options => {
        options.compilerOptions.preserveWhitespace = true
        return options
      })
      .end()
   },
  
  /*css: {
    loaderOptions: {
      css: {},
      postcss: {
        plugins: [
          // 补全css前缀(解决兼容性)
          require("autoprefixer")(),
          // 把px单位换算成rem单位
          require("postcss-pxtorem")({
            rootValue: 32, // 换算的基数(设计图750的根字体为32)
            selectorBlackList: [".van", ".my-van"],// 要忽略的选择器并保留为px。
            propList: ["*"], //可以从px更改为rem的属性。
            minPixelValue: 2 // 设置要替换的最小像素值。
          })
        ]
      }
    }
  },*/

  devServer: {
    port: 9000,
    // proxy: {/**处理跨域，本地代理转发*/
    //     '/internal': {
    //         target: 'http://192.168.2.75:9501/',  // 接口域名
    //         secure: false,  // 如果是https接口，需要配置这个参数
    //         changeOrigin: true,  //是否跨域}
    //     },
    // },
  }
}