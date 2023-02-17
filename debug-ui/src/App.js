import './App.css';

import React, {
  useState,
  createContext
} from "react";

import Room from './Room';
//import UsersContext from './UsersContext';

export const RoomContext = createContext();

//const usersContext = createContext();
const users = {
  user1: "TestUser1",
  user2: "TestUser2"
};

export default function App() {
  const [room, setRoom] = useState();

  let createRoom = async() => {
    try {
      let response = await fetch("/api/CreateRoom", {
        method: "POST",
        body: JSON.stringify({
          username: users.user1
        }),
        headers: { "Content-Type": "application/json"}
      });

      let responseJSON = await response.json();

      if(response.status === 200) {
        setRoom(responseJSON);
      } else {
        alert("Server cannot create room");
      }
    } catch (err) {
      console.log(err);
      alert("Error in Create Room");
    }
  };
  
  return (
    <div className="App">
      <h1>SL Debug - version {process.env.REACT_APP_VERSION}</h1>
      <button type="button" onClick={createRoom}>Create New Room</button>

      <RoomContext.Provider value={{ room, users }}>
        {room ? <Room /> : <></>}
      </RoomContext.Provider>
    </div>
  );
};