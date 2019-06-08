import React from 'react';
import './App.css';
import CreateGame from './CreateGame';
import Home from './Home';
import { BrowserRouter as Router, Route } from "react-router-dom";

class AppRouter extends React.Component {
  render() {
    return (
      <Router>
        <Route path="/" exact component={Home} />
        <Route path="/new-game/" component={CreateGame} />
      </Router>
    );
  }
}

export default AppRouter;
