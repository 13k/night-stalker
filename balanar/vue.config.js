const pkg = require("./package.json");

process.env.VUE_APP_NAME = pkg.displayName;
process.env.VUE_APP_VERSION = pkg.version;

module.exports = {
  devServer: {
    proxy: {
      "^(/api|/ws)": {
        target: "http://localhost:3000",
        ws: true,
        changeOrigin: false,
      },
    },
  },
  transpileDependencies: ["vuetify"],
};
