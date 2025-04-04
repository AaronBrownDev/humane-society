

/*
export default function Nav(){
    return(
        
            <header className= "header">
            <nav className = "nav-bar">
                <img src = "./src/assets/Humane-Society-logo2.png" />
                <ul className ="nav-list">
                    <li><a href=" ">Home</a></li>
                    <li><a href="">About</a></li>
                    <li><a href="ADOPT.jsx">Adopt</a></li>
                    <li><a href="https://www.paypal.com/donate/?cmd=_s-xclick&hosted_button_id=THJZGG93QJ3JS&ssrt=1742335653866">Donate</a></li>
                    <li><a href="">Other Ways to Support</a></li>
                    <li><a href="">Volunteer</a></li>
                    <li><a href="">Surrendering</a></li>
                    <li><a href="">Sponsors</a></li>
                    <li><a href="">Contact</a></li>
                </ul>
            </nav>

        </header>

        
        
        
    )
}
*/

import { NavLink } from "react-router-dom";


export default function Nav(){
	return(
	
        <div>
            <nav className = "nav-bar">
                <img src = "/src/assets/Humane-Society-logo2.png" alt= "Humane-Society-logo"/>
                <ul className="nav-list" >
                    <li><NavLink to= "../pages/Home"> Home</NavLink></li>
                    <li><NavLink to= "/About"> About</NavLink></li>
                    <li><NavLink to= "/Adopt"> Adopt</NavLink></li>
                    <li><a href="https://www.paypal.com/donate/?cmd=_s-xclick&hosted_button_id=THJZGG93QJ3JS&ssrt=1742335653866"> Donate</a></li>
                    <li><NavLink to= "/Support"> Other Ways to Support</NavLink></li>
                    <li><NavLink to= "/Volunteer"> Volunteering</NavLink></li>
                    <li><NavLink to= "/Surrendering"> Surrendering</NavLink></li>
                    <li><NavLink to= "/Sponsors"> Sponsors</NavLink></li>
                    <li><NavLink to= "/Contact"> Contact</NavLink></li>
                </ul>
            </nav>
            
        </div>
		
	)
	
}