import React, {useRef} from 'react'
import ReactHlsPlayer from 'react-hls-player';
import { BACKEND_URL } from '../config/Constants';


export default function VideoPlayer(props)
{

  return(
    <ReactHlsPlayer
      src={"http://192.168.31.100/video/3.m3u8"}
      autoPlay={false}
      controls={true}
      width="100%"
      height="auto"
    />
  );
}