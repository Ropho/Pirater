import React from 'react'
import ReactHlsPlayer from 'react-hls-player';
import { BACKEND_URL } from '../config/Constants';


export default function VideoPlayer(props)
{
    return(
    <ReactHlsPlayer
      src={BACKEND_URL + "/video/outputlist.m3u8"}
      autoPlay={false}
      controls={true}
      width="100%"
      height="auto"
    />
    );
}