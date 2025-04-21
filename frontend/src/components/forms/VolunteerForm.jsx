import React from "react";

export default function VolunteerForm(){
    const [inputs, setInputs]= React.useState([]);

    const handleChange = (event)=>{
        const name = event.target.name;
        const inputValue = event.target.value;
        setInputs(prevState => ({...prevState, [name]: inputValue}));
    }

    //Enter API call
	function handleSubmit(formData){
		const data = Object.fromEntries(formData)
		console.log(data)
	}
	
    return(
        <div className="volunteer-form">
        <form action={handleSubmit} >
                <label>First Name *: </label>
                    <input
                        type="text"
                        name="first_name"
                        value={inputs.first_name}
                        onChange={handleChange}
                        required/>
                <label>Last Name*</label>
                    <input
                        type="text"
                        name="last_name"
                        value={inputs.last_name}
                        onChange={handleChange}
                        required/>
                <label>Age *:</label>
                    <input
                        type="number"
                        name="age"
                        value={inputs.age}
                        onChange={handleChange}
                        required/>
                <label>Email *:</label>
                    <input
                        type="email"
                        name="email"
                        value={inputs.email}
                        onChange={handleChange}
                        required/>
                <label>Full Home Address *:</label>
                    <input
                        type="textarea "
                        name="address"
                        value={inputs.address}
                        onChange={handleChange}
                        required/>
                <label>Phone Number *:</label>
                    <input
                        type="tel"
                        name="phone"
                        value={inputs.phone}
                        onChange={handleChange}
                        placeholder='123-123-1234'
                        required/>
                <label>Do You Have Current Animals? *:</label>
                    <select
                        name="current_animals"
                        value={inputs.current_animals}
                        onChange={handleChange}
                        required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>

                <label>Animal Name(s):</label>
                    <input
                        type="text"
                        name="animal_names"
                        value={inputs.animal_names}
                        onChange={handleChange}
                    />
                <label>If So, How Many?:</label>
                    <input
                        type="number"
                        name="animal_count"
                        value={inputs.animal_count}
                        onChange={handleChange}
                    />
                <label>What Kind of Animals Do You Have?:</label>
                    <input
                        type="text"
                        name="breed"
                        value={inputs.breed}
                        onChange={handleChange}
                    />
            {/** Vet's Info*/}
                <label>Vet Name:</label>
                    <input
                        type="text"
                        name="veterinarian_name"
                        value={inputs.veterinarian_name}
                        onChange={handleChange}
                    />
                <label>Vet Number:</label>
                    <input
                        type="text"
                        name="veterinarian_number"
                        value={inputs.veterinarian_number}
                        onChange={handleChange}
                    />
                <label>List the Client's Name the Veterinarian Records Are Under: </label>
                    <input
                        type="text"
                        name="vet_client_name"
                        value={inputs.veterinarian_number}
                        onChange={handleChange}
                    />
                <label>Consent to Contact Veterinarian for Records:</label>
                    <input
                        type="text"
                        name="vet_consent"
                        value={inputs.vet_consent}
                        onChange={handleChange}
                    />
                <label>Authorization to Release Records:</label>
                    <input
                        type="text"
                        name="vet_authorization"
                        value={inputs.vet_authorization}
                        onChange={handleChange}
                    />
            {/**Volunteer Knowledge*/}
                <label>Are You Willing to Promote Spay/Neuter and Pet Education? *:</label>
                    <select
                        name="promote_spay_neuter"
                        value={inputs.promote_spay_neuter}
                        onChange={handleChange}
                        required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>

                <label>Do You Believe in Breeding of Dogs and Cats? *:</label>
                    <select
                        name="believe_breeding"
                        value={inputs.believe_breeding}
                        onChange={handleChange}
                        required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>

            {/**Volunteer Preferences*/}
                <label>Are You Willing to Clean Kennels and Scoop Yards? *:</label>
                    <select
                        name="cleaning_kennels"
                        value={inputs.cleaning_kennels}
                        onChange={handleChange}
                        required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>

                <label>Are You Willing to Brush and Play with Dogs? *:</label>
                    <select
                        name="brush_play_dogs"
                        value={inputs.brush_play_dogs}
                        onChange={handleChange}
                        required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>

                <label>Are You Allergic to Dogs? *: </label>
                    <select
                        name="allergic_dogs"
                        value={inputs.brush_play_dogs}
                        onChange={handleChange}
                        required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>

                <label>Do You Have Any Limitations for Certain Duties? *:</label>
                    <select
                        name="limitations"
                        value={inputs.limitations}
                        onChange={handleChange}
                        required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>

                <label>Are You Wanting to Volunteer for Community Service Hours? *:</label>
                    <select
                        name="community_service"
                        value={inputs.community_service}
                        onChange={handleChange}
                        required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>

                <label>If Yes, Please Explain Why Needed:</label>
                    <input
                        type="text"
                        name="community_service_reason"
                        value={inputs.community_service_reason}
                        onChange={handleChange}
                    />
                <label>If Yes, How Many Hours Are Needed and By When?:</label>
                    <input
                        type="number"
                        name="community_service_hours"
                        value={inputs.community_service_hours}
                        onChange={handleChange}
                    />
                <label>How Did You Hear About the HSNWLA? *:</label>
                    <input
                        type="text"
                        name="referral_source"
                        value={inputs.referral_source}
                        onChange={handleChange}
                        required/>
                <label>Any Other Questions or Comments?:</label>
                    <input
                        type="text"
                        name="additional_comments"
                        value={inputs.additional_comments}
                        onChange={handleChange}
                    />
                <label>Available Volunteer Shifts:</label>
                    <input
                        type="text"
                        name="volunteer_shifts"
                        value={inputs.volunteer_shifts}
                        onChange={handleChange}
                        required
                    />
                
                <button type="submit">Submit</button>
            </form>
        </div>
)
}