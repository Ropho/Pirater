import React, {useState} from 'react'
import {Link} from 'react-router-dom'
import "./FilmContainer.css"


function ImageContainer(props)
{
  return (
    <div
      className="film-card"
      id ={props.id}
    >
      <Link to = {`/currentFilm/${props.hash}`} >
      <img src={props.imgUrl} alt={props.name} className="film-image"/>
      <h3 className="film-title">{props.name}</h3>
      </Link>
    </div>
  );
}


export default function FilmsGrid(props) {
  
  let images = props.data.map((curFilm) => (
    <ImageContainer 
    key     = {curFilm.id} 
    id      = {curFilm.id} 
    imgUrl  = {curFilm.url} 
    name    = {curFilm.name} 
    hash    = {curFilm.hash}
    />
  ))

  return (
    <div className="films-container">
      <h2 className="header">New!</h2>
      <div className="films-grid">
        {images}   
      </div>
    </div>
  );
}