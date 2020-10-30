When we had only Comet and Ajax, to create a "long" connection between server
and client:
	* Relying on hacks in client and server to keep connections open longer
	and fake a long-running connection => **resource allocation** on server
	* Long polling makes unnecessary requests and keeps a constant stream of
	keeps a constant stream of opening and closing connections for your servers
	to deal with

WebSocket:
	* the ability to use an upgraded HTTP request
	* send data in a message-based, similar to UDP and with all reliability of TCP
	  with neligible penalty in resource utilization
	* can also layer another protocol on top of Websocket, and provide it in a
	  secure way over TLS

