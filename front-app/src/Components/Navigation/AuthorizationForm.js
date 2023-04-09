import React from 'react'
import Modal from 'react-modal';
import "./Authorization.css"

Modal.setAppElement('#root');


export default function LoginForm(props)
{
    const [userFormData, setUserFormData] = React.useState({
        email: "",
        pass: "",
    });

    const [isReg, setIsReg] = React.useState(false)


    function modalClose()
    {
        props.setModalIsOpen(false)
        setUserFormData({
            email: "",
            pass: "",
        })
        setIsReg(false);
    }

    function handleSubmit(event)
    {
        modalClose()
        event.preventDefault()
        console.log(userFormData)
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
        <div className='signup--text' onClick = {() => {setIsReg(true)}}>
            {isReg ? "" : "Sign up"}
        </div>
    </Modal>
    );
}