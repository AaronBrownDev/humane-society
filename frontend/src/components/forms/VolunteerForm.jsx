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
                <label>First Name *: <input type="text" name="first_name" required/> Last Name*<input type="text" name="last_name" required/></label>
                <label>Age *: <input type="text" name="age" required/></label>
                <label>Email *: <input type="email" name="email" required/></label>
                <label>Full Home Address *: <input type="textarea " name="address" required/></label>
                <label>Phone Number *: <input type="tel" name="phone" required/></label>
                <label>Do You Have Current Animals? *:
                    <select name="current_animals" required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>
                </label>
                <label>Animal Name(s): <input type="text" name="animal_names"/></label>
                <label>If So, How Many?: <input type="number" name="animal_count"/></label>
                <label>What Kind of Animals Do You Have?: <input type="text" name="breed"/></label>
                <label>Vet Name: <input type="text" name="veterinarian_name"/></label>
                <label>Vet Number: <input type="text" name="veterinarian_number"/></label>
                
                <label>List the Client's Name the Veterinarian Records Are Under: <input type="text" name="vet_client_name"/></label>
                <label>Consent to Contact Veterinarian for Records: <input type="text" name="vet_consent"/></label>
                <label>Authorization to Release Records: <input type="text" name="vet_authorization"/></label>
                <label>Are You Willing to Promote Spay/Neuter and Pet Education? *:
                    <select name="promote_spay_neuter" required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>
                </label>
                <label>Do You Believe in Breeding of Dogs and Cats? *:
                    <select name="believe_breeding" required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>
                </label>
                <label>Are You Willing to Clean Kennels and Scoop Yards? *:
                    <select name="cleaning_kennels" required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>
                </label>
                <label>Are You Willing to Brush and Play with Dogs? *:
                    <select name="brush_play_dogs" required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>
                </label>
                <label>Are You Allergic to Dogs? *:
                    <select name="allergic_dogs" required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>
                </label>
                <label>Do You Have Any Limitations for Certain Duties? *:
                    <select name="limitations" required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>
                </label>
                <label>Are You Wanting to Volunteer for Community Service Hours? *:
                    <select name="community_service" required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>
                </label>
                <label>If Yes, Please Explain Why Needed: <input type="text" name="community_service_reason"/></label>
                <label>If Yes, How Many Hours Are Needed and By When?: <input type="text" name="community_service_hours"/></label>
                <label>How Did You Hear About the HSNWLA? *: <input type="text" name="referral_source" required/></label>
                <label>Any Other Questions or Comments?: <input name="additional_comments"></input></label>
                <label>Available Volunteer Shifts: <input type="text" name="volunteer_shifts"/></label>
                
                <button type="submit">Submit</button>
            </form>
        </div>
)
}