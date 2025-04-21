import PublicNavbar from "./PublicNavbar.jsx"
import { Outlet } from "react-router-dom"
import Section8 from "../public/sections/Section8.jsx";

export default function PublicLayout() {
    return (
      <>
        <PublicNavbar />
        <div className="layout-content">
            <Outlet />
        </div>
        <Section8/>
      </>
    );
  }