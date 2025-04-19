import Navbar from "./Navbar.jsx"
import { Outlet } from "react-router-dom"
import Section8 from "../sections/Section8.jsx";

export default function Layout() {
    return (
      <>
        <Navbar />
        <div className="layout-content">
            <Outlet />
        </div>
        <Section8/>
      </>
    );
  }