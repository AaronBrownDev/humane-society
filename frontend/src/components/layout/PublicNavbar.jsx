
import { NavLink } from "react-router-dom";


export default function PublicNavbar(){
	return(
	
        <div>
            <nav className = "nav-bar">
                <img src = "/src/assets/Humane-Society-logo2.png" alt= "Humane-Society-logo"/>
                <ul className="nav-list" >
                    <li><NavLink to= "/"> Home</NavLink></li>
                    <li><NavLink to= "/About"> About</NavLink></li>
                    {/**Needs to be available dogs and not adopt dogs or can*/}
                    {/**<li><NavLink to="/Adopt"> Adopt</NavLink></li>*/}
                    <li><a href="https://www.paypal.com/donate/?cmd=_s-xclick&hosted_button_id=THJZGG93QJ3JS&ssrt=1742335653866"> Donate</a></li>
                    <li><NavLink to= "/Support"> Other Ways to Support</NavLink></li>
                    <li><NavLink to= "/Sponsors"> Sponsors</NavLink></li>
                    <li><NavLink to= "/Contact"> Contact</NavLink></li>
                    <li><NavLink to="/LoginPage">Login</NavLink></li>

                </ul>
            </nav>
            
        </div>
		
	)
	
}