import React from "react";
import "./sidebar.css"

export default function Sidebar(){
    return (
    <div className="mainSidebarContainer">
        <div className="mainSidebar">

            <ul className="ulContainer">
                <h4 className="menu">Menu</h4>
                <li className="liContainer">
                    <p className="itemNames">Home</p>
                </li>
                <li className="liContainer">
                    <p className="itemNames">Forms</p>
                </li>
                <li className="liContainer">
                    <p className="itemNames">Dogs</p>
                </li>
                <li className="liContainer">
                    <p className="itemNames">Medications</p>
                </li>
                <li className="liContainer">
                    <p className="itemNames">Supplies</p>
                </li>
                <li className="liContainer">
                    <p className="itemNames">Sign Out</p>
                </li>
            </ul>
            <div>
                <h4></h4>
            </div>
        </div>
    </div>
    )
}