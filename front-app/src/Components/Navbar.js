import React from 'react'
import "./Navbar.css"

export default function NavigationBar()
{
    return(
        <nav>
                <div className='navbar-name'>
                    <h1> KENTUKY FRIED CINEMA </h1>
                </div>
                <ul className='navbar-pannel-container'>
                    <li>Главная</li>
                    <li>Фильмы</li>
                    <li>Сериалы</li>
                    <li>Поддержка</li>
                </ul>
                <button type='button' className='navbar-button'> Log in</button>
        </nav>
    )
}