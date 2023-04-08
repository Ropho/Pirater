import React, { useEffect, useState } from 'react'
import { PATH_CUR_FILM } from '../config/Constants'
import { useParams } from 'react-router-dom'
import VideoPlayer from '../VideoPlayer/VideoPlayerMain'
import  './CurrentFilm.css'


function AfishaContainer(props)
{
  return (
    <div className = "afisha--container">
      <img className = "afisha--img--content" src={props.data.pic_url} />
      <div className = "afisha--text--content">
      <div className='film--name--block'> 
        <div className='film--name'>{props.data.name}</div>
        <div className='film--category'> {() => props.data.categories.map((element) => element)} </div>
      </div>
      <div className='film--description'> {props.data.description}</div>
      </div>
    </div>
  );
}


export default function CurrentFilm()
{
    const [filmData, setFilmData] = useState({})

    let params = useParams()

    useEffect(() => {
        fetch(PATH_CUR_FILM + `/${params.hash}`)
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
          <AfishaContainer data = {filmData} />
          <div className='videoPlayer--container'>
            <VideoPlayer />
          </div>
        </div>
    )
}