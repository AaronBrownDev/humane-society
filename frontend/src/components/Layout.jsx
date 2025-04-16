import Nav from "./Nav.jsx"
import { Outlet } from "react-router-dom"
import Section8 from "./Section8.jsx";

export default function Layout() {
    return (
      <>
        <Nav />
        <div className="layout-content">
            <Outlet />
        </div>
        <Section8/>
      </>
    );
  }