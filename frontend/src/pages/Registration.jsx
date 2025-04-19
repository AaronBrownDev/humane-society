import React, {useState} from 'react'
import "../styles/Registration.css"
import {NavLink} from "react-router-dom";

export default function Registration() {
    let isValid =true;
    const [formData, setFormData] = useState({
        first_name:'',
        last_name:'',
        email:'',
        password:'',
        confirmpassword:'',
    })

    const [error,setErrors] = useState({})
    const [valid,setValid] = useState(true)



    function handleSubmit() {

        let validationErrors = {}
        if (formData.first_name === "" || formData.first_name === null) {
            isValid = false;
            validationErrors.first_name = "First name is required";
        }
        if (formData.last_name === "" || formData.last_name === null) {
            isValid = false;
            validationErrors.last_name = "Last name is required";
        }

        if (formData.password === "" || formData.password === null) {
            isValid = false;
            validationErrors.password = "Password is required";
        } else if (formData.password.length < 6) {
            isValid = false;
            validationErrors.password = "Password need 6 or more characters";
        }

        if (formData.confirmpassword !== formData.password) {
            isValid = false;
            validationErrors.password = "Passwords do not match";
        }

        setErrors(validationErrors);
        setValid(isValid);

        if (Object.keys(!validationErrors).length === 0) {
            alert("Registration unsuccessful!");
        }else if (Object.keys(validationErrors).length === 0) {
            //Call api?
            alert("Registration Successful!");
            console.log(formData)
        }

    }
    return (
        <div className="registrationContainer">
            <div className="formconcontainer">
                {
                    valid ? <></>:
                        <span className = "text-danger">
                            {error.first_name}; {error.last_name};{error.password};{error.confirmpassword}

                        </span>
                }
                <form action={handleSubmit}>
                    <label htmlFor="first_name">First Name</label>
                    <input
                        type="text"
                        id="first_name"
                        name="first_name"
                        placeholder="First Name"
                        onChange={(event) => setFormData({...formData, first_name: event.target.value})}
                    />
                    <label htmlFor="last_name">Last Name</label>
                    <input
                        type="text"
                        id="last_name"
                        name="last_name"
                        placeholder="Last Name"
                        onChange={(event) => setFormData({...formData, last_name: event.target.value})}
                    />
                    <label htmlFor="email">Email</label>
                    <input
                        type = "email"
                        id='email'
                        name = "email"
                        placeholder="Email"
                        onChange={(event) => setFormData({...formData, email: event.target.value})}
                    />
                    <label htmlFor="password">Password</label>
                    <input
                        type="password"
                        id="password"
                        name="password"
                        placeholder="Password"
                        onChange={(event) => setFormData({...formData, password: event.target.value})}
                        />
                    <label htmlFor="confirmPassword">Confirm Password</label>
                    <input
                        type="password"
                        id="confirmPassword"
                        name="confirmPassword"
                        placeholder="Confirm Password"
                        onChange={(event) => setFormData({...formData, confirmpassword: event.target.value})}
                    />

                    <button type="submit">Submit</button>

                    <p>
                        Already have an Account?<br/>
                        <span className='line'>
                        <NavLink to ="/LoginPage"> Login </NavLink>
                    </span>

                    </p>

                </form>
            </div>

        </div>

    )
}