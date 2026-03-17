import { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import { loginUser } from "../api/api";

function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleLogin = async () => {
    const data = await loginUser(email, password);

    if (data.token) {
      localStorage.setItem("token", data.token);
      navigate("/dashboard");
    } else {
      setError(data.error || "Login failed");
    }
  };

  return (
    <div style={{ padding: "40px" }}>
      <h2>Login</h2>

      <input placeholder="Email" value={email} onChange={e => setEmail(e.target.value)} />
      <br /><br />

      <input type="password" placeholder="Password" value={password} onChange={e => setPassword(e.target.value)} />
      <br /><br />

      <button onClick={handleLogin}>Login</button>

      <p style={{ color: "red" }}>{error}</p>

      <Link to="/register">Create Account</Link>
    </div>
  );
}

export default Login;
