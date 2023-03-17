import React, { useState, useEffect } from 'react'
import {ReactComponent as Logo} from '../logo.svg'
import {NavLink, useMatch, useNavigate, useResolvedPath} from "react-router-dom"
import "./Navbar.css"
import Modal from 'react-modal';
import AuthorizationForm from "./AuthorizationForm"
import "./Authorization.css"

Modal.setAppElement('#root');

export default function NavigationBar(props)
{

  const [user, setUser] = React.useState({
    login: "",
    password: "",
    isLogin:false
  });



  const [modalIsOpen, setModalIsOpen] = useState(false);

  let activeStyle = {
    textDecoration: "underline",
    textDecorationColor: "#E13737",
  }; 

  function logInHandler()
  {
    setModalIsOpen(true);
  }

return (
  <div>
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
      <button onClick ={logInHandler}>{user.isLogin ? user.login : "Login"}</button>
    </div>
  </nav>
  <Modal 
    isOpen={modalIsOpen} 
    onRequestClose={() => setModalIsOpen(false)} 
    className="Modal--container"
    overlayClassName="Modal--overlay">
  <AuthorizationForm 
    modalHandler ={setModalIsOpen} 
    userHandler={setUser} 
    user = {user}/>
  </Modal>
  </div>
);
}
