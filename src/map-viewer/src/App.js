import React from 'react';
import './App.css';
import Map from './Map';
import { BrowserRouter as Router, Route } from "react-router-dom";

class AppRouter extends React.Component {
  render() {
    return (
      <Router>
        <Route path="/" exact component={Map} />
      </Router>
    );
  }
}

export default AppRouter;
