import React, { useEffect, useState } from 'react'
import Carousel     from './Carousel'
import FilmContainer from './FilmContainer'
import dataNews     from './data_news'
import {BACKEND_URL} from '../config/Constants'
import './MainPage.css'



export default function MainPage()
{
    const [dataCarousel, setDataCarousel] = useState([]);

    useEffect(() => {
        fetch('http://192.168.31.100:8080/api/carousel')
            .then(response => {
                if (response.ok) {
                    return response.json()
                }
            })
            .then(data => {
                console.log(data)
                setDataCarousel(data);
            })
    }, []);

    return(
        <main>
            <Carousel      data = {dataCarousel}/>
            <FilmContainer data = {dataNews} />
        </main>
    ); 
}