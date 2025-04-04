import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Layout from './components/Layout.jsx';
import Home from './pages/Home.jsx';
import About from './pages/About.jsx';
import Adopt from './pages/Adopt.jsx';
import Volunteer from './pages/Volunteer.jsx';
import Surrendering from './pages/Surrendering.jsx';
import Support from './pages/Support.jsx';
import Sponsors from './pages/Sponsors.jsx';
import Contact from './pages/Contact.jsx';

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="../pages/About" element={<About />} />
          <Route path="../pages/Adopt" element={<Adopt />} />
          <Route path="../pages/Support" element={<Support />} />
          <Route path="../pages/Volunteer" element={<Volunteer />} />
          <Route path="../pages/Surrendering" element={<Surrendering />} />
          <Route path="../pages/Sponsors" element={<Sponsors />} />
          <Route path="../pages/Contact" element={<Contact />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}