import React, {
    useState,
    useEffect,
    useMemo
} from "react";
import { useNavigate, useLocation } from "react-router-dom";
import useWebSocket, { ReadyState } from "react-use-websocket";

function Lobby() {
    const { state } = useLocation();
    const socketUrl = "ws://127.0.0.1:3001/ws/" + state.roomID + "/" + state.username;
    const { sendMessage, lastMessage, readyState } = useWebSocket(socketUrl);

    const [lobby, setLobby] = useState({})
    const [userReady, setUserReady] = useState({})

    const allReady = useMemo(() => {
        for (const key of Object.keys(lobby)) {
            if (!lobby[key]) {
                return false;
            }
        }
        return true;
    });

    useEffect(() => {
        setUserReady(lobby[state.username]);
    }, [lobby, state.username]);

    useEffect(() => {
        if (lastMessage !== null) {
            var payload = JSON.parse(lastMessage.data);
            //console.log(payload);
            setLobby(payload);
        }
    }, [lastMessage]);

    let handleChangeReadyClick = (ev) => {
        ev.preventDefault();
        let updatedValue = lobby;
        updatedValue[state.username] = !updatedValue[state.username];
        //console.log("trying to ready " + state.username)
        //lobby[state.username] = !lobby[state.username];
        setLobby(ps => ({
            ...ps,
            ...updatedValue
        }));
    };

    let startGame = () => {
        console.log("START!");
    };

    return (
        <div>
            <h2>Lobby {state.roomID}</h2>
            <ul>
                {
                    Object.keys(lobby).map((key, i) => {
                        return (
                            <li key={i}>
                                <span>{key}: {lobby[key] ? "Ready" : "Not Ready"}</span>
                            </li>
                        );
                    })
                }
            </ul>
            <input type="button" onClick={handleChangeReadyClick} value={userReady ? "Ready Down" : "Ready Up"}></input>
            <input type="button" onClick={startGame} value="Start Game!" disabled={!allReady}></input>
        </div>
    );
}

export default Lobby;
