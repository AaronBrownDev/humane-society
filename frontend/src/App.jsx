import { Routes, Route } from 'react-router-dom';
import PublicLayout from './components/layout/PublicLayout.jsx';
import Home from './pages/public/Home.jsx';
import About from './pages/public/About.jsx';
import Adopt from './pages/user/Adopt.jsx';
import Volunteer from './pages/user/Volunteer.jsx';
import Surrendering from './pages/user/Surrendering.jsx';
import Support from './pages/public/Support.jsx';
import Sponsors from './pages/public/Sponsors.jsx';
import Contact from './pages/public/Contact.jsx';
import LoginPage from "./pages/public/LoginPage.jsx";
import Registration from './pages/public/Registration.jsx';
import ProtectedRoute from './routes/ProtectedRoute.jsx'
import Dashboard from './pages/admin/Dashboard.jsx';
import ManageDogs from './pages/admin/ManageDogs.jsx';
import ManageUsers from './pages/admin/ManageUsers.jsx';
import AuthProvider from "./context/AuthProvider.jsx";
import AdminLayout from "./components/layout/AdminLayout.jsx";
import UserLayout from "./components/layout/UserLayout.jsx";
import Unauthorized from "./pages/public/Unauthorized.jsx";


export default function App() {
  return (

      <Routes>
        <Route path="/" element={<PublicLayout />}>
          {/** Public Routes*/}
          <Route index element={<Home />} />
          <Route path="/About" element={<About />} />
          <Route path="/Sponsors" element={<Sponsors />} />
          <Route path="/Contact" element={<Contact />} />
          <Route path="/Support" element={<Support />} />
          <Route path="/LoginPage" element = {<LoginPage/>}/>
          <Route path="/Registration" element={<Registration />} />
          <Route path="/Unauthorized" element={<Unauthorized />} />
            <Route path="/Adopt" element={<Adopt />} />
        </Route>

          {/**User Routes */}
          <Route element={ <UserLayout/> } >

            <Route path="/Volunteer" element={<Volunteer />} />
            <Route path="/Surrendering" element={<Surrendering />} />
          </Route>


          {/** Admin Routes (Protected*/}
        <Route path ="/" element={ <AdminLayout/> }>
            <Route path="/Dashboard" element={<Dashboard />} />
            <Route path="/ManageDogs" element={<ManageDogs />}/>
            <Route path="/ManageUsers" element={<ManageUsers />}/>
          </Route>



      </Routes>

  );
}