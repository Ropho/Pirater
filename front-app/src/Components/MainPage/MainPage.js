import React, { useEffect, useState } from 'react'
import Carousel     from './Carousel'
import FilmContainer from './FilmContainer'
import {BACKEND_URL} from '../config/Constants'
import './MainPage.css'

const numberOfElementInCarousel = 3;
const numberOfElementInNews     = 8;

export default function MainPage()
{
    const [dataCarousel, setDataCarousel] = useState([{header_url:'', name:'', hash:''}]);
    const [dataNews, setDataNews]   = useState([{name:'', hash:'', afisha_url:''}]);

    useEffect(() => {
        fetch(BACKEND_URL + `/api/carousel?count=${numberOfElementInCarousel}`)
            .then(response => {
                if (response.ok) {
                    return response.json()
                }
                throw new Error('response no OK');                 
            })
            .then(data => {
                setDataCarousel(data);
            })
            .catch(err => {
                console.log('carousel: ' + err.message)
            })
    }, []);

    useEffect(() => {
        fetch(BACKEND_URL + `/api/newFilms?count=${numberOfElementInNews}`)
            .then(response => {
                if (response.ok) {
                    return response.json()
                }
                throw new Error('response no OK'); 
            })
            .then(data => {
                setDataNews(data);
            })
            .catch(error => {
                console.log('news: ' + error.message)
            })
    }, []);


    return(
        <main>
            <Carousel      data = {dataCarousel}/>
            <FilmContainer data = {dataNews} />
        </main>
    ); 
}