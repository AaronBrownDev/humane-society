import {useContext} from "react";
import AuthProvider from "../context/AuthProvider.jsx"

export default function useAuth() {
    return useContext(AuthProvider);
}

