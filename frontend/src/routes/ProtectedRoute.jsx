import {useLocation, Navigate, Outlet} from "react-router-dom";
import useAuth from "../hooks/useAuth.js"

export default function ProtectedRoute ({allowedRoles}) {
    const {auth} = useAuth();
    const location = useLocation();


    return (
        auth?.roles?.find(roles => allowedRoles?.includes(roles))
            ? <Outlet/>
            : auth?.user
                ? <Navigate to="/" state={{from: location}} replace/>
                : <Navigate to="/login" state={{from: location}} replace/>
    )
}