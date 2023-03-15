import React, { useEffect, useState } from 'react'
import Carousel     from './Carousel'
import FilmContainer from './FilmContainer'
import dataCarousel from './data_carousel'
import dataNews     from './data_news'
import './MainPage.css'



export default function MainPage()
{

    return(
        <main>
            <Carousel      data = {dataCarousel}/>
            <FilmContainer data = {dataNews} />
        </main>
    ); 
}