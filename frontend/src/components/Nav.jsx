import { NavLink } from "react-router-dom";


export default function Nav(){
	return(
	
        <div className="header">
            <nav className = "nav-bar">
                <img src = "/src/assets/Humane-Society-logo2.png" alt= "Humane-Society-logo"/>
                <ul className="nav-list" >
                    <li><NavLink to= "/"> Home</NavLink></li>
                    <li><NavLink to= "../pages/About"> About</NavLink></li>
                    <li><NavLink to= "../pages/Adopt"> Adopt</NavLink></li>
                    <li><a href="https://www.paypal.com/donate/?cmd=_s-xclick&hosted_button_id=THJZGG93QJ3JS&ssrt=1742335653866"> Donate</a></li>
                    <li><NavLink to= "../pages/Support"> Other Ways to Support</NavLink></li>
                    <li><NavLink to= "../pages/Volunteer"> Volunteering</NavLink></li>
                    <li><NavLink to= "../pages/Surrendering"> Surrendering</NavLink></li>
                    <li><NavLink to= "../pages/Sponsors"> Sponsors</NavLink></li>
                    <li><NavLink to= "../pages/Contact"> Contact</NavLink></li>
                </ul>
            </nav>
            
        </div>
	)
	
}