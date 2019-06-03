import React from 'react';
import logo from './logo.svg';
import './App.css';

function App() {
  return (
    <div className="welcome-screen">
      <header className="header">
        <h1>
          Welcome to Agenda!
        </h1>
        <h3>
          Turn-based game of political intrigue played via e-mail.
        </h3>
        <div className="login-buttons">
          <button className="button">
            Register
          </button>
          <button
            className="button">
            Log in
          </button>
        </div>
      </header>
    </div>
  );
}

export default App;
