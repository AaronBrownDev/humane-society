import React from 'react'
import '../../styles/manageUsers.css'


export default function ManageUsers() {
    const [user,setUser] = React.useState([

    ])
    //Api call
    async function handleSubmit(formData) {
        try {
            const response = await fetch('https://api.example.com/data', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    FirstName: '',
                    LastName: '',
                    Email: '',
                    Phone: '',
                    Password: '',
                    Confirmation: '',

                }),
            });

            const result = await response.json();
            console.log('Success:', result);
        } catch (error) {
            console.error('Error:', error);
        }

    }


    function handleChange(event) {
        const newUser = event.target.newUser;
        const inputValue = event.target.value;
        setUser(prevState => ({...prevState, [newUser]: inputValue}));
    }
    return (
        <div className={"users-container"}>

            <div className="top">
                <h1>Add Users</h1>
            </div>
            <div className="bottom">
                <div className="left">
                    <img src="../../assets/placeholder-1024x1024.png"/>
                </div>
                <div className="right">
                    <form action={handleSubmit}>
                        <label> First Name</label>
                        <input
                            type="text"
                            name="FirstName"
                            value ={user.FirstName}
                            onChange={handleChange}
                            placeholder="First Name"
                        />
                        <label> Last Name</label>
                        <input
                            type="text"
                            name="Last Name"
                            value ={user.LastName}
                            onChange={handleChange}
                            placeholder="Last Name"
                        />

                        <label> Email</label>
                        <input
                            type="email"
                            name="Email"
                            value ={user.Email}
                            onChange={handleChange}
                            placeholder="email@gmail.com"
                        />

                        <label> Phone Number</label>
                        <input
                            type="tel"
                            name="Phone"
                            value ={user.Phone}
                            onChange={handleChange}
                            placeholder="123-123-1234"
                        />

                        <label>Password</label>
                        <input
                            type="password"
                            name="password"
                            value ={user.password}
                            onChange={handleChange}

                        />

                        <label> Confirm Password</label>
                        <input
                            type="password"
                            name="confirmpassword"

                        />
                        <button type="submit">Add User</button>
                    </form>

                </div>
            </div>



        </div>
    )

}