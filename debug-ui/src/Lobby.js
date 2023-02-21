import React, {
    useState,
    useContext,
    useEffect,
    useMemo
}
from "react";
import useWebsocket, { ReadyState } from "react-use-websocket";

import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

import { RoomContext } from "./App";

export default function Lobby({username}) {
    const { room, wsPrefix } = useContext(RoomContext);
    const { sendMessage, lastMessage, readyState } = useWebsocket(wsPrefix + "JoinRoom/" + room.roomID + "/" + username, {
        onOpen: () => {
            console.log(username + " successful websocket connection to lobby");
        },
        onError: (ev) => {
            console.log(ev);
        },
        onClose: () => {
            console.log(username + " websocket connection closed");
        }
    });

    const [lobby, setLobby] = useState({});
    const [userReady, setUserReady] = useState();

    const allReady = useMemo(() => {
        for(const key of Object.keys(lobby)) {
            if(!lobby[key]) {
                return false;
            }
        }
        return true;
    }, [lobby]);

    useEffect(() => {
        if(lastMessage !== null) {
            var msg = JSON.parse(lastMessage.data)
            console.log("Server lobby WS message received: " + msg.toString())
            setLobby(msg);
        }
    }, [lastMessage]);

    useEffect(() => {
        setUserReady(lobby[username]);
    }, [lobby, username]);

    let toggleReady = () => {
        // let newLobby = lobby;
        // newLobby[username] = !newLobby[username];
        // setUserReady(newLobby[username]);
        // setLobby(newLobby);
        var readyResult = !lobby[username];
        var msg = readyResult ? "ready" : "not ready";
        if(readyState === ReadyState.OPEN) {
            console.log("Sending ready change message")
            sendMessage(msg);
        }
        // lobby => ({
        //     ...lobby,
        //     ...msg
        // })
    };

    let startGame = () => {
        console.log("Game Start!");
    };

    return(
        <div className="square border">
            <Container>
                {
                    Object.keys(lobby).map((key) => {
                        return (
                            <Row key={key}>
                                <span>{key}: {lobby[key] ? "Ready" : "Not Ready"}</span>
                            </Row>
                        );
                    })
                }
                <Row>
                    <Col>
                        <button type="button" onClick={toggleReady}>
                            {userReady ? "Un-Ready" : "Ready"}
                        </button>
                    </Col>
                    <Col>
                        { room.owner === username ?
                            <button type="button" onClick={startGame} disabled={!allReady}>
                                Start!
                            </button> :
                            <></>
                        }
                    </Col>
                </Row>
            </Container>
        </div>
    );
};