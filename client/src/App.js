import './App.css';

import React, { Component } from "react";
import {
  Route, Routes,
  HashRouter,
} from "react-router-dom";

import SplashScreen from './SplashScreen';
import CreateRoom from './CreateRoom';
import JoinRoom from './JoinRoom';
import Lobby from './Lobby';

class App extends Component {
  render() {
    return (
      <HashRouter>
        <div className="App">
          <h1>Squad Leader Application</h1>
        </div>
        <div className="content">
          <Routes>
            <Route path="/" element={<SplashScreen/>}/>
            <Route path="/CreateRoom" element={<CreateRoom/>}/>
            <Route path="/JoinRoom" element={<JoinRoom/>}/>
            <Route path="/Lobby" element={<Lobby/>}/>
          </Routes>
        </div>
      </HashRouter>
    );
  }
}

export default App;
