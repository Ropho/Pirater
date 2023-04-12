import './App.css';
import NavigationBar from "./Components/Navigation/Navbar"
import MainPage      from './Components/MainPage/MainPage';
import CurrentFilm   from './Components/CurrentFilm/CurrentFilm';
import React, { useEffect, useState } from 'react'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { WHOAMI_URL } from './Components/config/Constants';



function App() {

  const [userData, setUserData] = useState({

    Email: "",
    Right: "",
    isLogin: false,
  })

  useEffect(() => {
    fetch(WHOAMI_URL)
    .then(response => {
      if (response.ok)
      {
        return response.json()
      }
      throw 'cancel';
    })
    .then(data => {
      data = {...data, isLogin: true}
      setUserData(data)

      console.log(userData)
    })
    .catch(() => {
      console.log(userData)
    }
    )
  }, [])

  return (
    <Router>
      <div>
        <NavigationBar userData = {userData} setUserData = {setUserData}/>
        <Routes>
            <Route path = "/" exact element={<MainPage/>} />
            <Route path = "/support" exact element={<h1>Help support</h1>} />
            <Route path = "/film/:hash" 
            exact element={<CurrentFilm/>} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
