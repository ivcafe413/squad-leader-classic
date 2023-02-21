import './App.css';

import React, {
  useState,
  createContext
} from "react";

import Room from './Room';

export const RoomContext = createContext();

const users = {
  user1: "TestUser1",
  user2: "TestUser2"
};
const wsPrefix = "ws://127.0.0.1:3001/ws/";

export default function App() {
  const [room, setRoom] = useState();

  let createRoom = async() => {
    try {
      console.log("Creating a new room...")
      let response = await fetch("/api/CreateRoom", {
        method: "POST",
        body: JSON.stringify({
          username: users.user1
        }),
        headers: { "Content-Type": "application/json"}
      });

      let responseJSON = await response.json();

      if(response.status === 200) {
        console.log("Room created successfully!")
        setRoom(responseJSON);
      } else {
        console.log("Server cannot create room");
      }
    } catch (err) {
      console.log(err);
    }
  };
  
  return (
    <div className="App">
      <h1>SL Debug - version {process.env.REACT_APP_VERSION}</h1>
      <button type="button" onClick={createRoom}>Create New Room</button>

      <RoomContext.Provider value={{
        room,
        users,
        wsPrefix
      }}>
        {room ? <Room /> : <></>}
      </RoomContext.Provider>
    </div>
  );
};