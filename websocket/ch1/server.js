var WebSocketServer = require('ws').Server;

var ws = new WebSocketServer({ port: 8181 });

ws.on('connection', () => {
  console.log('connection establish');
  ws.on('message', message => {
    console.log('receive message ', message);
  })
})

