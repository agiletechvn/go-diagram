/**
 * The IO between the server and client.
 */
let conn;
class Connection {
  static setUp() {
    console.log("Initializing websockets...");
    conn = new WebSocket("ws://localhost:8080/ws?lastMod=143918dd9ce16851");
    conn.onclose = e => {
      console.log("Connection closed");
      window.location.reload();
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
