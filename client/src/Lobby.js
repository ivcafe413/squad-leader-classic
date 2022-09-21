import React, {
    useState
} from "react";
import { useNavigate, useLocation } from "react-router-dom";
import useWebSocket, { ReadyState } from "react-use-websocket";

function Lobby() {
    const { state } = useLocation();
    
    const socketUrl = "ws://127.0.0.1:3001";
    const { sendMessage, lastMessage, readyState } = useWebSocket(socketUrl)
}