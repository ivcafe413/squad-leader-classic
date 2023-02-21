import React, {
    useState,
} from "react";

import Lobby from "./Lobby";

export default function User({username}) {
    const [ userJoined, setUserJoined ] = useState(false);

    let joinRoom = () => {
        console.log(username + " is joining the lobby")
        setUserJoined(true);
    };

    return(
        <div className="square border">
            <h3>{username}</h3>
            {userJoined ?
                <Lobby username={username}/> :
                <button type="button" onClick={joinRoom}>
                    Join Lobby
                </button>
            }
        </div>
    );
};