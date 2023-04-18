import React from 'react'
import {Splide, SplideSlide } from '@splidejs/react-splide';
import {Link} from 'react-router-dom'
import '@splidejs/react-splide/css';
import './Carousel.css'

export default function Carousel(props)
{
    const img = props.data.map((curImg) => {
        return(
            <SplideSlide key={curImg.hash}>
                <Link to = {`/film/${curImg.hash}`} >
                    <img className = 'carousel--img' src = {curImg.header_url} alt = {curImg.name}></img>
                </Link>
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
        interval     : 3000,
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