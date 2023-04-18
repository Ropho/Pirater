import React from 'react'
import ReactHlsPlayer from 'react-hls-player';


export default function VideoPlayer(props)
{

  return(
    <ReactHlsPlayer
      src={props.video_url}
      autoPlay={false}
      controls={true}
      width="100%"
      height="auto"
    />
  );
}