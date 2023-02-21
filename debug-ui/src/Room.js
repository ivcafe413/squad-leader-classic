import React, {
    useContext
} from "react";

import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

import { RoomContext } from "./App";
import User from "./User";
import MessageHistory from "./MessageHistory";

//class Room extends Component {
export default function Room() {
    const { room, users } = useContext(RoomContext);

    return(
        <div>
            <h2>Room {room.RoomID}</h2>
            <Container>
                <Row>
                    <Col>
                        <User username={users.user1} />
                    </Col>
                    <Col>
                        <MessageHistory />
                    </Col>
                    <Col>
                        <User username={users.user2} />
                    </Col>
                </Row>
            </Container>
        </div>
    );
}
