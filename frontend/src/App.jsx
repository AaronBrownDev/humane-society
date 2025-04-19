import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Layout from './components/layout/Layout.jsx';
import Home from './pages/Home.jsx';
import About from './pages/About.jsx';
import Adopt from './pages/Adopt.jsx';
import Volunteer from './pages/Volunteer.jsx';
import Surrendering from './pages/Surrendering.jsx';
import Support from './pages/Support.jsx';
import Sponsors from './pages/Sponsors.jsx';
import Contact from './pages/Contact.jsx';
import LoginPage from "./pages/LoginPage.jsx";
import Registration from './pages/Registration.jsx';

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