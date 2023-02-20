import React from 'react'
import "./Navbar.css"

export default function NavigationBar()
{
    return(
        <div>
            <nav>
                <div className='navbar-name'>
                    <h1> KENTUKY FRIED CINEMA </h1>
                </div>
                <div className='navbar-pannel-container'>
                    <div>Главная</div>
                    <div> Фильмы</div>
                    <div> Сериалы</div>
                    <div> ТВ</div>
                    <div> Поддержка</div>
                </div>
                <button className='navbar-button'> Log in</button>
            </nav>
        </div>
    )
}