process.env.VUE_APP_NAME = "Balanar";
process.env.VUE_APP_VERSION = require("./package.json").version;

module.exports = {
  devServer: {
    host: "localhost",
    proxy: {
      "^(/api|/ws)": {
        target: "http://localhost:3000",
        ws: true,
        changeOrigin: false
      }
    }
  },
  transpileDependencies: ["vuetify"]
};
