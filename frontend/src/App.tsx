import { useEffect, useState } from "react";

function App() {
  const [message, setMessage] = useState("Loading...");

  useEffect(() => {
    fetch("http://localhost:8080/health")
      .then((response) => response.text())
      .then((data) => setMessage(data))
      .catch(() => setMessage("Failed to connect to backend"));
  }, []);

  return (
    <div>
      <h1>Economic Data Dashboard</h1>
      <p>Backend status: {message}</p>
    </div>
  );
}

export default App;