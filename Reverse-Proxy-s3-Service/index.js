const express = require("express");
const httpProxy = require("http-proxy");

const app = express();
const PORT = 8000;

const BASE_PATH = "https://go-vercel.s3.ap-south-1.amazonaws.com/build";

const proxy = httpProxy.createProxy();

app.use((req, res) => {
  const hostname = req.hostname;
  console.log(hostname);
  const subdomain = hostname.split(".")[0];
  console.log(subdomain);

  // Custom Domain - DB Query
  const resolvesTo = `${BASE_PATH}/${subdomain}`;
  console.log(resolvesTo);

  return proxy.web(req, res, { target: resolvesTo, changeOrigin: true });
});

proxy.on("proxyReq", (proxyReq, req, res) => {
  const url = req.url;
  if (url === "/") proxyReq.path += "index.html";
  console.log(url);
});

app.listen(PORT, () => console.log(`Reverse Proxy Running..${PORT}`));
