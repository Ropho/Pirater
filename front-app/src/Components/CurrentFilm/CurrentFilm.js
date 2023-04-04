import React, { useEffect, useState } from 'react'
import { BACKEND_URL } from '../config/Constants'
import { useParams } from 'react-router-dom'
import Container from '@mui/material/Container';
import Box from '@mui/material/Box';
import  './CurrentFilm.css'
import { Typography } from '@mui/material';

function ImageContainer()
{
  return (
    <div className = "afisha--container">
      <img className = "afisha--img--content" src="/john.jpeg" />
      <div className = "afisha--text--content">
      <div className='film--name--block'> 
        <div className='film--name'>Джон Кеков</div>
        <div className='film--category'> Шутер, повседневность, этти </div>
      </div>
      <div className='film--description'> John Wick is an American action thriller media franchise created by 
      Derek Kolstad and centered around John Wick, a former hitman who is forced back into the criminal underworld he had previously abandoned. 
      The franchise began with the release of John Wick in 2014, followed by three sequels: Chapter 2 in 2017, Chapter 3 – Parabellum in 2019, and Chapter 4 in 2023.</div>
      </div>
    </div>
  );
}


export default function CurrentFilm()
{
    const [filmData, setFilmData] = useState({})

    let params = useParams()

    useEffect(() => {

        fetch(BACKEND_URL + `/currentFilm/${params.hash}`)
            .then(response => {
                if (response.ok) {
                    return response.json();
                }
                throw new Error('response is not OK');                 
            })
            .then(data => {
                console.log(data);
                setFilmData(data);
            })
            .catch(err => {
                console.log('CurrentFilm: ' + err.message)
            })
    }, []);
    

    return(
        <div className= "current--container">
          <ImageContainer/>
        </div>
    )
}