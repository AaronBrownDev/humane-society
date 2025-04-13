import Nav from "./Nav.jsx"
import Section8 from "./Section8.jsx"
import { Outlet } from "react-router-dom"

export default function Layout() {
    return (
      <>
        <Nav />
        <div className="layout-content">
          <Outlet />
        </div>
        <Section8 />
      </>
    );
  }