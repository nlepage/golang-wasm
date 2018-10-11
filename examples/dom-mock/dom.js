const { JSDOM } = require("jsdom")

const dom = new JSDOM(
  `
  <!DOCTYPE html>
  <html>
    <head>
      <title>Go wasm</title>
    </head>
    <body></body>
  </html>
  `,
  {
    url: "http://localhost:8080/wasm_exec.html"
  },
)

global.window = dom.window
global.document = window.document
