import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Layout from './components/layout/Layout.jsx';
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

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="/About" element={<About />} />
          <Route path="/Adopt" element={<Adopt />} />
          <Route path="/Support" element={<Support />} />
          <Route path="/Volunteer" element={<Volunteer />} />
          <Route path="/Surrendering" element={<Surrendering />} />
          <Route path="/Sponsors" element={<Sponsors />} />
          <Route path="/Contact" element={<Contact />} />
          <Route path="LoginPage" element = {<LoginPage/>}/>
          <Route path="Registration" element={<Registration />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}