import React, {
    useState,
    useContext
} from "react";

import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

//import UsersContext from "./UsersContext";
import { RoomContext } from "./App";
import User from "./User";

//class Room extends Component {
export default function Room() {
    const { room, users } = useContext(RoomContext);

    return(
        <div>
            <h2>Room {room.RoomID}</h2>
            <Container>
                <Row>
                    <Col>
                        {/* <h3>{users.user1}</h3> */}
                        <User username={users.user1} />
                    </Col>
                    <Col>
                        {/* <h3>{users.user2}</h3> */}
                        <User username={users.user2} />
                    </Col>
                </Row>
            </Container>
        </div>
    );
}
