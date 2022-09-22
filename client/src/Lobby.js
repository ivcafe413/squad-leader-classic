import React, {
    useState,
    useEffect
} from "react";
import { useNavigate, useLocation } from "react-router-dom";
import useWebSocket, { ReadyState } from "react-use-websocket";

function Lobby() {
    const { state } = useLocation();
    const socketUrl = "ws://127.0.0.1:3001/ws/" + state.room + "/" + state.user;
    const { sendMessage, lastMessage, readyState } = useWebSocket(socketUrl);

    const [lobby, setLobby] = useState({})

    useEffect(() => {
        if (lastMessage !== null) {
            var payload = JSON.parse(lastMessage.data);
            console.log(payload);
            setLobby(payload);
        }
    }, [lastMessage]);

    return (
        <div>
            <h2>Lobby {state.room}</h2>
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
        </div>
    );
}

export default Lobby;
