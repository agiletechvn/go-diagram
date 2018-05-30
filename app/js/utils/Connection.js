/**
 * The IO between the server and client.
 */
let conn;
class Connection {
  static setUp() {
    console.log("Initializing websockets...");
    conn = new WebSocket(
      "ws://localhost:8080/ws?lastMod=" + new Date().getTime()
    );
    conn.onclose = e => {
      console.log("Connection closed");
      // setTimeout(() => window.location.reload(), 2000);
    };
  }

  static onMessage(callback) {
    conn.onmessage = callback;
  }

  static sendMessage(newPackageData) {
    conn.send(JSON.stringify(newPackageData));
  }
}

export default Connection;
