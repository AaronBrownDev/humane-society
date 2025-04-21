import AdminSideBar from './AdminSideBar.jsx';
import {Outlet, Navigate} from "react-router-dom";
import useAuth from "../../hooks/useAuth";
import AdminNavBar from "./AdminNavBar.jsx";
import '../../styles/Admin.css'

export default function AdminLayout() {
    /**
    const { user, isAuthenticated } = useAuth();

    if (!isAuthenticated) {
        return <Navigate to="/login" replace />;
    }

    if (user.role !== 'admin' && user.role !== 'staff') {
        return <Navigate to="/Unauthorized" replace />;
    }
    */

    return (
        <div className="dashboard">
            <AdminSideBar />
            <div className="layout-container">
              <AdminNavBar/>
                <Outlet/>
            </div>



        </div>
    )
}