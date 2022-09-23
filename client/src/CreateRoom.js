import React, {
    useState
} from "react";
import { useNavigate } from 'react-router-dom';

//class CreateRoom extends Component {
function CreateRoom() {
    const [username, setUsername] = useState("");
    const navigate = useNavigate();

    let handleSubmit = async(event) => {
        event.preventDefault();

        try {
            let response = await fetch("/api/CreateRoom", {
                method: "POST",
                body: JSON.stringify({
                    username: username,
                }),
                headers: { "Content-Type": "application/json" },
            });

            let responseJSON = await response.json();

            if(response.status === 200) {
                alert("Room ID: " + responseJSON.roomID + " Created!");
                navigate("/JoinRoom", {
                    state: {
                        username: username,
                        roomID: responseJSON.roomID
                    }
                });
            } else {
                alert("Cannot create the room");
            }
        } catch (err) {
            console.log(err);
            alert("Error in Create Room request");
        }
    };

    return (
        <div>
            <h2>Create Room</h2>
            <form onSubmit={handleSubmit}>
            <label>
                User:
                <input type="text" value={username} onChange={(ev) => setUsername(ev.target.value)}></input>
            </label>
            <input type="submit" value="Create Room"></input>
            </form>
        </div>
    );
}


export default CreateRoom;
