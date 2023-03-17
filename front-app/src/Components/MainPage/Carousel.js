import React from 'react'
import {Splide, SplideSlide } from '@splidejs/react-splide';
import '@splidejs/react-splide/css';
import './Carousel.css'

export default function Carousel(props)
{
    
    const img = props.data.map((curImg) => {
        console.log(curImg.url)
        return(
            <SplideSlide key={curImg.id}>
                <img className = 'carousel--img' src = {curImg.url} alt = {curImg.name}></img>
            </SplideSlide>
        )
    })

    const options = {    
        type         : 'loop',
        gap          : '10pt',
        autoplay     : true,
        pauseOnHover : false,
        resetProgress: false,
        autowidth    : true,
        arrows       : false,
        interval     : 4000,
        pagination   : false,
        speed        : 1000,
        perPage: 3,
        perMove: 1,
        start: 1,
        classes: {
            active:  'is-active',
            visible: 'is-visible',
            prev:    'is-prev',
            next:    'is-next',
        },
    };

    return(
        <div className='splide--container'>
            <Splide aria-label="Afisha" tag ="section" options = {options} className = 'splide--body'>
                {img}
            </Splide>
        </div>
    );
}