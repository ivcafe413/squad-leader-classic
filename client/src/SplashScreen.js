import React, { Component } from "react";
import {
    NavLink
} from "react-router-dom"

class SplashScreen extends Component {
    render() {
        return (
            <div>
                <NavLink to="/CreateRoom">Create Room</NavLink>
                <NavLink to="/JoinRoom">Join Room</NavLink>
            </div>
        )
    }
}

export default SplashScreen;
