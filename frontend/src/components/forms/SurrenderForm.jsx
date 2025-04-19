import React from 'react';


export default function SurrenderForm(){
	
    const [inputs,setInputs] = React.useState({});


    const handleChange = (event) => {
        const name = event.target.name;
        const inputValue = event.target.value;
        setInputs(prevState => ({...prevState, [name]: inputValue}));
    }

    //Change to do api call;
    function handleSubmit(formData){
		const data = Object.fromEntries(formData)
		console.log(data)
	}
	
    

    return (
        
        <form  action={handleSubmit}>
            <label>FirstName *:
                <input
                    type="text"
                    name="first_name"
                    value = {inputs.first_name}
                    onChange={handleChange}
                    required/>
                </label>
            <label>
                Last Name *:
                <input
                    type="text"
                    name="last_name"
                    value = {inputs.last_name}
                    onChange={handleChange}
                    required/>
            </label>
            <label>Address *: </label>
                <input
                    name="surr_address"
                    value = {inputs.surr_address}
                    onChange={handleChange}
                    required/>

            <label>Email *: </label>
                <input
                    type="email"
                    name="email"
                    value = {inputs.email}
                    onChange={handleChange}
                    required/>

            <label>Phone Number *: </label>
                <input
                    type="tel"
                    name="phone"
                    value = {inputs.phone}
                    onChange={handleChange}
                    required/>

            <label>Dog's Name *: </label>
                <input
                    type="text"
                    name="dog_name"
                    value = {inputs.dog_name}
                    onChange={handleChange}
                    required/>

            <label>Where Did You Get the Dog? *: </label>
                <input
                    type="text"
                    name="dog_source"
                    value = {inputs.dog_source}
                    onChange={handleChange}
                    required/>

            <label>Breed *: </label>
                <input
                    type="text"
                    name="breed"
                    value = {inputs.breed}
                    onChange={handleChange}
                    required/>


            <label>Date of Last Vet Visit *: </label>
                <input type="date"
                       name="last_vet_visit"
                       value = {inputs.last_vet_visit}
                       onChange={handleChange}
                       required/>
            <label>Gender *: </label>
                <select value = {inputs.gender}
                        onChange={handleChange}
                        name="gender" required>
                    <option  value="Male">Male</option>
                    <option  value="Female">Female</option>
                </select>

            <label>Current on Vaccinations? *: </label>
                <select
                    name="vaccinations"
                    value = {inputs.vaccinations}
                    onChange={handleChange}
                    required>
                <option  defaultValue="--Choose Yes or no">--Choose an Option</option>
                <option value="Yes">Yes</option>
                <option value="No">No</option>
                </select>

            <label>Name of Current Veterinarian *: </label>
                <input
                    type="text"
                    name="vet_name"
                    value = {inputs.vet_name}
                    onChange={handleChange}
                    required/>

            <label>Number of Current Veterinarian *: </label>
                <input
                    type="tel"
                    name="vet_number"
                    value = {inputs.vet_number}
                    onChange={handleChange}
                    required/>
            
            <label>Any Known Medical Problems *: </label>
                <input
                    type="text"
                    name="medical_problems"
                    value = {inputs.medical_problems}
                    onChange={handleChange}
                    required/>

            <label>Current on Heartworm Prevention? *: </label>
                <select
                    value = {inputs.heartworm_prevention}
                    onChange={handleChange}
                    name="heartworm_prevention"
                    required>
                <option value="Yes">Yes</option>
                <option value="No">No</option>
            </select>

            <label>Microchip Number (if applicable):</label>
                <input
                    type="text"
                    name="microchip"
                    value = {inputs.microchip}
                    onChange={handleChange}
                />

            <label>Housetrained? *:</label>
                <select
                    name="housetrained"
                    value = {inputs.housetrained}
                    onChange={handleChange}
                    required>
                <option value="Yes">Yes</option>
                <option value="No">No</option>
            </select>

            <label>Any Bite History?:</label>
                <input
                    type="text"
                    name="bite_history"
                    value = {inputs.bite_history}
                    onChange={handleChange}
                    required
                />

            <label>How Long Have You Owned the Dog? *:</label>
                <input
                    type="text"
                    name="ownership_duration"
                    value = {inputs.ownership_duration}
                    onChange={handleChange}
                    required/>
            <label>Age *: </label>
                <input
                    type="text"
                    name="age"
                    value = {inputs.age}
                    onChange={handleChange}
                    required/>
            <label>Weight *: </label>
                <input
                    type="text"
                    name="dog_weight"
                    value = {inputs.dog_weight}
                    onChange={handleChange}
                    required/>
            <label>Spayed/Neutered? *: </label>
                <select
                    name="spayed_neutered"
                    value = {inputs.spayed_neutered}
                    onChange={handleChange}
                    required>
                    <option value="Yes">Yes</option>
                    <option value="No">No</option>
            </select>
            <label>Good with Strangers? *: </label>
                <select
                    name="good_with_strangers"
                    value = {inputs.good_with_strangers}
                    onChange={handleChange}
                    required>
                <option value="Yes">Yes</option>
                <option value="No">No</option>
            </select>
            <label>Good with Dogs? *: </label>
                <select
                    name="good_with_dogs"
                    value = {inputs.good_with_dogs}
                    onChange={handleChange}
                    required>
                <option value="Yes">Yes</option>
                <option value="No">No</option>
            </select>
            <label>Good with Cats? *: </label>
                <select
                    name="good_with_cats"
                    value = {inputs.good_with_cats}
                    onChange={handleChange}
                    required>
                <option value="Yes">Yes</option>
                <option value="No">No</option>
            </select>
            <label>Good with Kids? *: </label>
                <select
                    name="good_with_kids"
                    value = {inputs.good_with_kids}
                    onChange={handleChange}
                    required>
                <option value="Yes">Yes</option>
                <option value="No">No</option>
            </select>

            <label>Inside or Outside Dog? *: </label>
                <select
                    name="inside_outside"
                    value = {inputs.inside_outside}
                    onChange={handleChange}
                    required>
                <option value="Inside">Inside</option>
                <option value="Outside">Outside</option>
            </select>
            <label>Why Are You Re-homing? Any Other Comments?: </label>
                <input
                    name="rehoming_reason"
                    value={inputs.rehoming_reason||''}
                    onChange={handleChange}
                ></input>
            <label>Picture of the Dog: </label>
                <input
                    type="file"
                    name="dog_picture"
                    value = {inputs.dog_picture}
                    onChange={handleChange}
                    accept="image/*"/>
            <button type= "submit"> Submit </button>
    </form>
     
        
    
    );

}


