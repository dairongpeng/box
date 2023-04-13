可以使用浏览器来测试这个 WebSocket 服务。在浏览器的控制台中，使用以下代码连接到服务器。

```js
var socket = new WebSocket("ws://localhost:8080/websocket");

socket.onopen = function() {
    console.log("Connected");
    socket.send("Hello, server!");
};

socket.onmessage = function(event) {
    console.log("Received message: " + event.data);
};

socket.onclose = function(event) {
    console.log("Disconnected");
};
```

这将连接到 ws://localhost:8080/websocket 地址，并在连接成功后发送一条消息 "Hello, server!"，
然后在收到服务器的响应消息时输出它们。当连接关闭时，将输出 "Disconnected"。