import React, { useState } from 'react'
import Modal from 'react-modal';
import "./Authorization.css"
import {REGISTRATION_URL, SESSIONS_URL} from "../config/Constants"

Modal.setAppElement('#root');



export default function LoginForm(props)
{
    const [userFormData, setUserFormData] = useState({
        email: "",
        pass: "",
    });

    const [isReg, setIsReg] = useState(false)
    const [errorMessage, setErrorMessage]   = useState("") 


    function modalClose()
    {
        props.setModalIsOpen(false)

        setErrorMessage("")

        setUserFormData({
            email: "",
            pass: "",
        })

        setIsReg(false);
    }

    function succeedReq()
    {
        props.setUserData(() => {
            return{
                Email  : userFormData.email,
                Right  : "user",
                isLogin: true,
            }
        })

        modalClose()
    }

    function handleSubmit(event)
    {
        if (isReg)
        {
            fetch(REGISTRATION_URL, {
                method: "POST",
                mode: "cors",
                body: JSON.stringify(userFormData),
            })
            .then(response => {
                if (response.ok) 
                {
                    succeedReq();
                }
                else
                {
                    return response.text();
                }
            })
            .then(text => {
                setErrorMessage(text)
            })
        }
        else
        {
            fetch(SESSIONS_URL, {
                method: "POST",
                mode: "cors",
                body: JSON.stringify(userFormData),
            })
            .then(response => {
                if (response.ok){
                    console.log("HERE")
                    succeedReq();
                }
                else
                {
                    return response.text();
                }
            })
            .then(text => {
                setErrorMessage(text)
            })
        }


        event.preventDefault();
    }

    function handleChange(event)
    {   
        setUserFormData(prevFormData => {
        return{
            ...prevFormData,
            [event.target.name]: event.target.value
        }
        })
    }


    return(
    <Modal
    isOpen={props.modalIsOpen} 
    onRequestClose={() => modalClose()} 
    className="Modal--container"
    overlayClassName="Modal--overlay">

        <form onSubmit={handleSubmit}>
            <h3>{isReg ? "Welcome aboard" : "Run a Rig"}</h3>
            <input type="email" 
                   name="email" 
                   placeholder='Email' 
                   value={userFormData.email} 
                   onChange={handleChange}/>

            <input type="password" 
                   name="pass" 
                   placeholder='Password' 
                   value={userFormData.pass}
                   onChange={handleChange}/>

            <button>{ isReg ? "Sign Up" : "Sign In"}</button>
        </form>
        
        <div className='error--text'>
           {errorMessage}
        </div>
        
        <div className='signup--text' onClick = {() => {
            setIsReg(!isReg)
            setErrorMessage("")
            }}
        >
            {isReg ? "Sign In" : "Sign up"}
        </div>
    </Modal>
    );
}