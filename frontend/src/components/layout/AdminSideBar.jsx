import '../../styles/AdminSideBar.css'
import {NavLink} from 'react-router-dom'




export default function AdminSideBar() {

    return (
        <div className="navigation">
            <div className="admin-side-bar-container">
                <div className='admin-side-bar'>
                    <div className='header'>
                        <img src='../../assets/humane-society-2022_orig.png' alt='Humane Society Logo' />
                        <h4> Management System</h4>
                    </div>
                    <ul className="admin-ulContainer">
                        <h4 className="menu"> Menu </h4>
                        <li className="admin-liContainer"><NavLink to='/Dashboard'>Dashboard</NavLink></li>
                        <li className="admin-liContainer"><NavLink to='ManageUsers'>Manage Users</NavLink> </li>
                        <li className="admin-liContainer"><NavLink to='ManageDogs'>Manage Dogs</NavLink> </li>

                    </ul>
                </div>
            </div>
         </div>
    )
}

