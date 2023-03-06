import React, { useState, useEffect } from 'react'
import {ReactComponent as Logo} from '../logo.svg'
import {NavLink, useMatch, useResolvedPath} from "react-router-dom"
import "./Navbar.css"


export default function NavigationBar(props)
{
  const [data, setData] = useState([]);
  
  const apiGet = () => {
    fetch("http://192.168.1.122:8080/static/kek.json")
      .then((response) => response.json())
      .then((json) => {
      console.log(json);
      setData(json);
      });
    };

  let activeStyle = {
    textDecoration: "underline",
    textDecorationColor: "#E13737",
  }; 

return (
  <nav className="Navbar">
    <NavLink exact to="/">
      <Logo className = "Navbar--logo"/>
    </NavLink>
    <div className="Navbar--nav">
      <ul>
        <li>
          <NavLink to="/" style = {({isActive}) => isActive ? activeStyle : undefined}>
            Главная
          </NavLink>
        </li>
        <li>
          <NavLink to="/films" style={({ isActive }) => isActive ? activeStyle : undefined}>
            Фильмы
          </NavLink>
        </li>
        <li>
          <NavLink to="/serials" style={({ isActive }) => isActive ? activeStyle : undefined}>
            Сериалы
          </NavLink>
        </li>
        <li>
          <NavLink to="/support" style={({ isActive }) => isActive ? activeStyle : undefined}>
            Поддержка
          </NavLink>
        </li>
      </ul>
    </div>
    <div className="Navbar--auth">
      <button onClick ={apiGet}>Login</button>
    </div>
  </nav>
);
}
