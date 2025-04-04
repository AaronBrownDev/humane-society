import Nav from "./Nav.jsx"
import { Outlet } from "react-router-dom"

export default function Layout() {
    return (
      <>
        <Nav />
        <div className="layout-content">
          <Outlet />
        </div>
      </>
    );
  }