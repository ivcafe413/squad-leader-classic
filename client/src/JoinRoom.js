import React, {
    useState
} from "react";
import { useNavigate, useLocation } from "react-router-dom";

function JoinRoom() {
    const { state } = useLocation();
    const [room, setRoom] = useState(state.room);
    const [username, setUsername] = useState(state.username);

    const navigate = useNavigate();

    let handleSubmit = async(event) => {
        event.preventDefault();

        try {
            let response = await fetch("/api/JoinRoom", {
                method: "POST",
                body: JSON.stringify({
                    roomID: room,
                    username: username
                }),
                headers: { "Content-Type": "application/json" }
            });

            let responseText = await response.text();

            if(response.status === 200) {
                alert("Joining room " + room + ".....");
                navigate("/Lobby", {
                    state: {
                        room: room,
                        user: username
                    }
                });
            } else {
                console.log(responseText);
                alert("Could not Join Room");
            }
        } catch (err) {
            console.log(err);
            alert("Error in Joining Room");
        }
    };

    return (
        <div>
            <h2>Join Room</h2>
            <form onSubmit={handleSubmit}>
            <label>
                Room #:
                <input type="text" value={room}
                    onChange={(ev) => setRoom(ev.target.value)}></input>
            </label>
            <label>
                User:
                <input type="text" value={username}
                    onChange={(ev) => setUsername(ev.target.value)}></input>
            </label>
            <input type="submit" value="Join Room"></input>
            </form>
        </div>
    );
}

export default JoinRoom;
