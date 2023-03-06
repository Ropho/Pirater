import React, {useState} from 'react'
import "./FilmContainer.css"


function ImageContainer(props)
{
  return (
    <div 
      className="film-card"
      id ={props.id}
    >
      <img src={props.URL} alt={props.name} className="film-image"/>
      <h3 className="film-title">{props.name}</h3>
    </div>
  );
}


export default function FilmsGrid(props) {
  
  let images = props.data.map((curFilm) => (
    <ImageContainer 
    key  = {curFilm.id} 
    id   = {curFilm.id} 
    URL  = {curFilm.URL} 
    name = {curFilm.name} 
    description = {curFilm.description} 
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