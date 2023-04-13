import React, { useState} from 'react'
import {ReactComponent as Logo} from '../../logo.svg'
import {NavLink} from "react-router-dom"
import "./Navbar.css"
import AuthorizationForm from "./AuthorizationForm"


let activeStyle = {
  textDecoration: "underline",
  textDecorationColor: "#E13737",
}; 



export default function NavigationBar(props)
{
  const [modalIsOpen, setModalIsOpen] = useState(false);

  console.log(props.userData.isLogin)
  console.log('hehe')

  return (  
    <div>

      <NavBarElements setModalIsOpen = {setModalIsOpen} isLogin = {props.userData.isLogin} name = {props.userData.Email} />

      <AuthorizationForm 
        setModalIsOpen ={setModalIsOpen} 
        modalIsOpen  ={modalIsOpen}
        userData     = {props.userData}
        setUserData  = {props.setUserData}
      />

    </div>
  );
}


function NavBarElements(props)
{
  return(
  <nav className="Navbar">
    <NavLink to="/">
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

      {props.isLogin ? <p onClick ={() => {props.setModalIsOpen(true)}}> props.name </p> :
                       <button onClick ={() => {props.setModalIsOpen(true)}}> Login </button>
      }
    </div>
  </nav>
  )
}