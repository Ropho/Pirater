import React, { useState} from 'react'
import {ReactComponent as Logo} from '../../logo.svg'
import {NavLink} from "react-router-dom"
import "./Navbar.css"
import AuthorizationForm from "./AuthorizationForm"
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
import Avatar from '@mui/material/Avatar';
import { OUT_URL } from '../config/Constants'


let activeStyle = {
  textDecoration: "underline",
  textDecorationColor: "#E13737",
}; 


export default function NavigationBar(props)
{
  const [modalIsOpen, setModalIsOpen] = useState(false);

  return (  
    <div>
      <NavBarElements setModalIsOpen = {setModalIsOpen} userData = {props.userData} setUserData = {props.setUserData}/>

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
      {props.userData.isLogin ? <UserMenu userData = {props.userData} setUserData = {props.setUserData}/> :
                       <button onClick ={() => {props.setModalIsOpen(true)}}> Login </button>
      }
    </div>
  </nav>
  )
}

function UserMenu(props)
{
  const [anchorEl, setAnchorEl] = React.useState(null);
  const open = Boolean(anchorEl);

  const handleClick = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const handleLogout = () => {

    fetch(OUT_URL)
    .then(response => {
      if (response.ok)
      {
        props.setUserData(() => {
          return{
            email:  "",
            rights: "",    
            modified:"",
            registered: "",
            isLogin: false,
          }
        })
      }
      throw 'Log out error'
    })
    .catch((err) => {
      alert(err)
    })


  };

  return (
    <div>
      <Avatar
        id="user-icon"
        sx={{bgcolor: stringToColor(props.userData.email), width:50, height: 48}}
        onClick={handleClick}
      >
        {stringAvatar(props.userData.email)}
      </Avatar>
      <Menu
        id="basic-menu"
        anchorEl={anchorEl}
        open={open}
        onClose={handleClose}
        MenuListProps={{
          'aria-labelledby': 'basic-button',
        }}
      >
        <MenuItem onClick={handleClose}>{props.userData.email}</MenuItem>
        <MenuItem onClick={handleLogout}>Logout</MenuItem>
      </Menu>
    </div>
  );
}


function stringToColor(string) {
  let hash = 0;
  let i;

  for (i = 0; i < string.length; i += 1) {
    hash = string.charCodeAt(i) + ((hash << 5) - hash);
  }

  let color = '#';

  for (i = 0; i < 3; i += 1) {
    const value = (hash >> (i * 8)) & 0xff;
    color += `00${value.toString(16)}`.slice(-2);
  }

  return color;
}

function stringAvatar(name) {

  return (
    `${name[0].toUpperCase()+name[1].toUpperCase()}`
  );
}
