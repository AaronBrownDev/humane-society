// frontend/src/routes/ProtectedRoute.jsx - Update your existing file
import { Navigate, Outlet, useLocation } from "react-router-dom";
import useAuth from "../hooks/useAuth";

export default function ProtectedRoute() {
    const { auth, loading } = useAuth();
    const location = useLocation();

    console.log("ProtectedRoute - auth state:", auth);
    console.log("ProtectedRoute - loading:", loading);

    // If still loading auth state, show nothing or a loading spinner
    if (loading) {
        return <div>Loading...</div>;
    }

    // If not authenticated, redirect to login
    if (!auth.isAuthenticated) {
        console.log("Not authenticated, redirecting to login");
        return <Navigate to="/LoginPage" state={{ from: location }} replace />;
    }

    // If authenticated, render the child routes
    console.log("Authenticated, rendering children");
    return <Outlet />;
}