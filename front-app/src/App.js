import './App.css';
import NavigationBar from "./Components/Navigation/Navbar"
import MainPage      from './Components/MainPage/MainPage';
import CurrentFilm   from './Components/CurrentFilm/CurrentFilm';
import React from 'react'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';



function App() {
  return (
    <Router>
      <div>
        <NavigationBar/>
        <Routes>
            <Route path = "/" exact element={<MainPage/>} />
            <Route path = "/support" exact element={<h1>Help support</h1>} />
            <Route path = "/currentFilm/:hash" 
            exact element={<CurrentFilm/>} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
