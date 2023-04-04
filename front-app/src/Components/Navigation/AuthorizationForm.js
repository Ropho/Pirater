import React from 'react'
import "./Authorization.css"

export default function LoginForm(props)
{

    function handleSubmit(event)
    {
        event.preventDefault()
        props.modalHandler(false);
        props.userHandler((prevFormData) => {
            return{
                ...prevFormData,
                isLogin:true,
            }
        })
        console.log(props.user)
    }

    function handleChange(event)
    {   
        props.userHandler(prevFormData => {
        return{
            ...prevFormData,
            [event.target.name]: event.target.value
        }
        })
    }


    return(
    <div className='form--container'>
    <form onSubmit={handleSubmit}>
        <h3>Run a Rig</h3>
        <input type="text" 
               name="login" 
               placeholder='Login' 
               value={props.user.login} 
               onChange={handleChange}/>

        <input type="password" 
               name="password" 
               placeholder='Password' 
               value={props.user.password}
               onChange={handleChange}/>

        <button>Login</button>
    </form>
    </div>
    );
}