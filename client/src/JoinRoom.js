import React, {
    useState
} from "react";
import useWebSocket, { ReadyState } from "react-use-websocket";

function JoinRoom() {
    const socketUrl = "ws://127.0.0.1:3001";
    const { sendMessage, lastMessage, readyState } = useWebSocket(socketUrl)
}
    render() {
        return (
            <div>
                <h2>Join Room</h2>
                <form>
                <label>
                    Room #:
                    <input type="text" name="roomID"></input>
                </label>
                <label>
                    User:
                    <input type="text" name="username"></input>
                </label>
                <input type="submit" value="Join Room"></input>
                </form>
            </div>
        )
    }
}

export default JoinRoom;
