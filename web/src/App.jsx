import { useState } from "react";
import axios from "axios";

function App() {
  const [itemId, setItemId] = useState("");
  const [email, setEmail] = useState("");
  const [msg, setMsg] = useState("");

  const subscribe = async () => {
    setMsg("Sending...");

    try {
      const res = await axios.post("http://localhost:8080/api/subscribe", {
        itemId,
        email,
      });

      setMsg("✔ Subscription successful!");
    } catch (err) {
      console.error(err);
      setMsg("❌ Error: Could not subscribe");
    }
  };

  return (
    <div style={styles.container}>
      <h2>Woolworths Price Tracker</h2>

      <input
        style={styles.input}
        type="text"
        placeholder="Item ID"
        value={itemId}
        onChange={(e) => setItemId(e.target.value)}
      />

      <input
        style={styles.input}
        type="email"
        placeholder="Your email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />

      <button style={styles.button} onClick={subscribe}>
        Subscribe
      </button>

      <p>{msg}</p>
    </div>
  );
}

const styles = {
  container: {
    maxWidth: "400px",
    margin: "60px auto",
    fontFamily: "Arial",
    textAlign: "center",
  },
  input: {
    width: "100%",
    padding: "12px",
    margin: "10px 0",
    fontSize: "16px",
  },
  button: {
    width: "100%",
    padding: "12px",
    marginTop: "10px",
    backgroundColor: "#4CAF50",
    border: "none",
    color: "white",
    fontSize: "16px",
    cursor: "pointer",
  },
};

export default App;
