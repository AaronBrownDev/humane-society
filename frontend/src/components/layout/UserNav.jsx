import {NavLink} from 'react-router-dom';

export default function UserNav() {

    return (
        <>
            <nav className='user-nav'>
                <img src = "/src/assets/Humane-Society-logo2.png" alt= "Humane-Society-logo"/>
                <ul className="nav-list" >
                    <li><NavLink to= "/"> Home</NavLink></li>
                    <li><NavLink to= "/About"> About</NavLink></li>

                    <li><NavLink to="/Adopt">Available Dogs</NavLink></li>
                    <li><a href="https://www.paypal.com/donate/?cmd=_s-xclick&hosted_button_id=THJZGG93QJ3JS&ssrt=1742335653866"> Donate</a></li>
                    <li><NavLink to= "/Support"> Other Ways to Support</NavLink></li>
                    <li><NavLink to= "/Volunteer"> Volunteering</NavLink></li>
                    <li><NavLink to= "/Surrendering"> Surrendering</NavLink></li>
                    <li><NavLink to= "/Sponsors"> Sponsors</NavLink></li>
                    <li><NavLink to= "/Contact"> Contact</NavLink></li>
                    <li><NavLink to="/LoginPage">Login Page</NavLink></li>

                </ul>

            </nav>
        </>

    )
}