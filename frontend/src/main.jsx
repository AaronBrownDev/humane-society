import {createRoot} from 'react-dom/client';
import App from './App.jsx';
import './index.css';
import AuthProvider from '/src/context/AuthProvider';
import {BrowserRouter, Route, Routes} from "react-router-dom";


const root = createRoot(document.getElementById("root"));

root.render(
    <AuthProvider>
        <BrowserRouter>
            <Routes>
                <Route path="/*" element={<App/>}/>
            </Routes>
        </BrowserRouter>

    </AuthProvider>

   )

