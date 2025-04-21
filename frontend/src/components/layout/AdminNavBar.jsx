import '../../styles/AdminNavBar.css'
import SearchIcon from '@mui/icons-material/Search';

export default function AdminNavBar() {

    return (
        <div className="admin-nav-bar">
            <div className="wrapper">
                <div className="search">
                    <input type='text' placeholder='Search...' />
                    <SearchIcon/>
                </div>
            </div>
        </div>
    )
}





