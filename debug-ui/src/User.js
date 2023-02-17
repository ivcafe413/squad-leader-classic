import React, {
    useState
} from "react";
import useWebsocket, { ReadyState } from "react-use-websocket";

// import Container from 'react-bootstrap/Container';
// import Row from 'react-bootstrap/Row';
// import Col from 'react-bootstrap/Col';

//import UsersContext from "./UsersContext";

export default function User({username}) {
    return(
        <div>
            <h3>{username}</h3>

            {/* <Lobby/> */}
        </div>
    );
};