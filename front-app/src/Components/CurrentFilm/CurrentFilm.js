import React, { useEffect, useState } from 'react'
import { PATH_CUR_FILM } from '../config/Constants'
import { useParams } from 'react-router-dom'
import VideoPlayer from '../VideoPlayer/VideoPlayerMain'
import  './CurrentFilm.css'


function AfishaContainer(props)
{
  const category = () => {
    const categories = props.data.categories.map(element => (<div>{element}</div>))
    return categories;
  }

  return (
    <div className = "afisha--container">
      <img className = "afisha--img--content" src={props.data.header_url} />
      <div className = "afisha--text--content">
      <div className='film--name--block'> 
        <div className='film--name'>{props.data.name}</div>
        <div className='film--category'> {category()} </div>
      </div>
      <div className='film--description'> {props.data.description}</div>
      </div>
    </div>
  );
}


export default function CurrentFilm()
{
    window.scrollTo(0, 0)
    
    const [filmData, setFilmData] = useState({categories:[], video_url:"",})

    const params = useParams()
    console.log("Ya yebal rot")

    useEffect(() => {
        fetch(PATH_CUR_FILM + `/${params.hash}`)
            .then(response => {
                if (response.ok) {
                    return response.json();
                }
                throw new Error('response is not OK');                 
            })
            .then(data => {
                if (data === undefined)
                {
                  throw 'undefined data';
                }
                setFilmData(data);
            })
            .catch(err => {
                console.log('CurrentFilm error: ' + err.message)
            })
    }, []);
    

    return(
        <div className= "current--container">
          <AfishaContainer data = {filmData} />
          <div className='videoPlayer--container'>
            <VideoPlayer video_url = {filmData.video_url}/>
          </div>
        </div>
    )
}