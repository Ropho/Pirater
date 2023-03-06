import './App.css';
import NavigationBar from "./Components/Navbar"
import MainPage      from './Components/MainPage/MainPage';
import React from 'react'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';



function App() {
  return (
    <Router>
      <div>
        <NavigationBar/>
        <Routes>
            <Route path = "/" element={<MainPage/>}/>
            <Route path = "/support" element={<h1>Help support</h1>}/>
        </Routes>
      </div>
    </Router>
  );
}

export default App;
