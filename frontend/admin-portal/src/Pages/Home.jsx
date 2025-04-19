import React from 'react';
import Sidebar from '../components/sidebar.jsx';
import HomeRightBar from '../components/homerighbar/homerightbar.jsx';
import "./Home.css"

export default function Home(){
    return(
        <div className="mainHomeContainer">
            <Sidebar/>
            <HomeRightBar/>
        </div>
    )
}