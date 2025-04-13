export default function VolunteerForm(formData){
	
	function handleSubmit(formData){
		const data = Object.fromEntries(formData)
		console.log(data)
	}
	
    return(
        <form action={handleSubmit}>
                <label>Name *: <input type="text" name="first_name" required/> <input type="text" name="last_name" required/></label><br></br>
                <label>Age *: <input type="number" name="age" required/></label><br></br>
                <label>Email *: <input type="email" name="email" required/></label><br></br>
                <label>Full Home Address *: <input type="textarea " name="address" required/></label><br></br>
                <label>Phone Number *: <input type="tel" name="phone" required/></label><br></br>
                <label>Do You Have Current Animals? *:
                    <select name="current_animals" required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>
                </label><br></br>
                <label>Animal Name(s): <input type="text" name="animal_names"/></label><br></br>
                <label>If So, How Many?: <input type="number" name="animal_count"/></label><br></br>
                <label>What Kind of Animals Do You Have?: <input type="text" name="breed"/></label><br></br>
                <label>Vet Name: <input type="text" name="veterinarian_name"/></label><br></br>
                <label>Vet Number: <input type="text" name="veterinarian_number"/></label><br></br>
                
                <label>List the Client's Name the Veterinarian Records Are Under: <input type="text" name="vet_client_name"/></label><br></br>
                <label>Consent to Contact Veterinarian for Records: <input type="text" name="vet_consent"/></label><br></br>
                <label>Authorization to Release Records: <input type="text" name="vet_authorization"/></label><br></br>
                <label>Are You Willing to Promote Spay/Neuter and Pet Education? *:
                    <select name="promote_spay_neuter" required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>
                </label><br></br>
                <label>Do You Believe in Breeding of Dogs and Cats? *:
                    <select name="believe_breeding" required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>
                </label><br></br>
                <label>Are You Willing to Clean Kennels and Scoop Yards? *:
                    <select name="cleaning_kennels" required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>
                </label><br></br>
                <label>Are You Willing to Brush and Play with Dogs? *:
                    <select name="brush_play_dogs" required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>
                </label><br></br>
                <label>Are You Allergic to Dogs? *:
                    <select name="allergic_dogs" required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>
                </label><br></br>
                <label>Do You Have Any Limitations for Certain Duties? *:
                    <select name="limitations" required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>
                </label><br></br>
                <label>Are You Wanting to Volunteer for Community Service Hours? *:
                    <select name="community_service" required>
                        <option value="Yes">Yes</option>
                        <option value="No">No</option>
                    </select>
                </label><br></br>
                <label>If Yes, Please Explain Why Needed: <input type="text" name="community_service_reason"/></label><br></br>
                <label>If Yes, How Many Hours Are Needed and By When?: <input type="text" name="community_service_hours"/></label><br></br>
                <label>How Did You Hear About the HSNWLA? *: <input type="text" name="referral_source" required/></label><br></br>
                <label>Any Other Questions or Comments?: <textarea name="additional_comments"></textarea></label><br></br>
                <label>Available Volunteer Shifts: <input type="text" name="volunteer_shifts"/></label><br></br>
                
                <button type="submit">Submit</button>
            </form>
)
}