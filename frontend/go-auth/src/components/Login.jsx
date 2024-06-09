import axios from "axios";
import React, { useState } from "react";

function Login() {
  const [user, setUser] = useState({
    email: "",
    password: "",
  });
  const [message, setMessage] = useState("");
  const [isLoggedIn, setIsLoggedIn] = useState(false);
const [login,setLogin]=useState("Login")
  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post("http://localhost:3000/api/login", user);
      console.log(response.data);
      console.log(response.status);

      if (response.status === 200) {
        setIsLoggedIn(true);
        setMessage("Login successful!");
        setLogin("")
      }
    } catch (error) {
      console.error(error.message);
      setMessage("Login failed. Please try again.");
    }
  };

  const handleLogout = async () => {
    try {
      const response = await axios.get("http://localhost:3000/api/logout");
      if (response.status === 200) {
        setIsLoggedIn(false);
        setUser({ email: "", password: "" });
        setMessage("Logged out successfully.");
      }
    } catch (error) {
      console.error(error.message);
      setMessage("Logout failed. Please try again.");
    }
  };

  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <form onSubmit={handleLogin} className="w-full max-w-sm">
        <div className="md:flex md:items-center mb-6">
          <div className="md:w-1/3">
            <label
              className="block text-gray-500 font-bold md:text-right mb-1 md:mb-0 pr-4"
              htmlFor="inline-email"
            >
              Email
            </label>
          </div>
          <div className="md:w-2/3">
            <input
              className="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
              id="inline-email"
              type="text"
              placeholder="Email"
              value={user.email}
              onChange={(e) => setUser({ ...user, email: e.target.value })}
            />
          </div>
        </div>
        <div className="md:flex md:items-center mb-6">
          <div className="md:w-1/3">
            <label
              className="block text-gray-500 font-bold md:text-right mb-1 md:mb-0 pr-4"
              htmlFor="inline-password"
            >
              Password
            </label>
          </div>
          <div className="md:w-2/3">
            <input
              className="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
              id="inline-password"
              type="password"
              placeholder="******************"
              value={user.password}
              onChange={(e) => setUser({ ...user, password: e.target.value })}
            />
          </div>
        </div>
        <div className="flex justify-center">
          <button className="font-3xl bold" type="submit">
           {login}
          </button>
        </div>
      </form>

      {isLoggedIn && (
        <button onClick={handleLogout} className="font-3xl bold mt-4">
          Logout
        </button>
      )}

      {message && <p className="mt-4">{message}</p>}
    </div>
  );
}

export default Login;
