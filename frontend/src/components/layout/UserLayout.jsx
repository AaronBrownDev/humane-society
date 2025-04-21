import UserNav from "./UserNav.jsx";
import {Outlet, Navigate} from "react-router-dom";
import Section8 from "../public/sections/Section8.jsx";
import useAuth from "../../hooks/useAuth.js";

export default function UserLayout() {
    const isAuthenticated = useAuth();
    if (!isAuthenticated) {
        return <Navigate to="/login" replace />;
    }
    return (
        <>
            <UserNav />
            <Outlet/>
            <Section8/>
        </>
    )
}