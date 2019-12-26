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
  }
};
