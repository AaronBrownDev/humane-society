export default function surrenderform(){
	
    function handleSubmit(formData){
		const data = Object.fromEntries(formData)
		console.log(data)
	}
	
    

    return (
        
        <form  action={handleSubmit}>
            <label>Date *: <input type="date" name="date" required/></label>
            <label>FirstName *: <input type="text" name="first_name" required/> Last Name *:<input type="text" name="last_name" required/></label>
            <label>Address *: <input type="textarea" name="address" required/></label>
            <label>Email *: <input type="email" name="email" required/></label>
            <label>Phone Number *: <input type="tel" name="phone" required/></label>
            <label>Dog's Name *: <input type="text" name="dog_name" required/></label>
            <label>Where Did You Get the Dog? *: <input type="text" name="dog_source" required/></label>
            <label>Breed *: <input type="text" name="breed" required/></label>
            <label>Date of Last Vet Visit *: <input type="date" name="last_vet_visit" required/></label>
            <label>Gender *: <select name="gender" required>
                <option value="Male">Male</option>
                <option value="Female">Female</option>
            </select></label>
            <label>Current on Vaccinations? *: <select name="vaccinations" required>
                <option value="Yes">Yes</option>
                <option value="No">No</option>
            </select></label>
            <label>Name of Current Veterinarian *: <input type="text" name="vet_name" required/></label>
            <label>Number of Current Veterinarian *: <input type="text" name="vet_number" required/></label>
            
            <label>Any Known Medical Problems *: <input type="text" name="medical_problems" required/></label>
            <label>Current on Heartworm Prevention? *: <select name="heartworm_prevention" required>
                <option value="Yes">Yes</option>
                <option value="No">No</option>
            </select></label>
            <label>Microchip Number (if applicable): <input type="text" name="microchip"/></label>
            <label>Housetrained? *: <select name="housetrained" required>
                <option value="Yes">Yes</option>
                <option value="No">No</option>
            </select></label>
            <label>Any Bite History?: <textarea name="bite_history"></textarea></label>
            <label>How Long Have You Owned the Dog? *: <input type="text" name="ownership_duration" required/></label>
            <label>Age *: <input type="text" name="age" required/></label>
            <label>Weight *: <input type="text" name="weight" required/></label>
            <label>Spayed/Neutered? *: <select name="spayed_neutered" required>
                <option value="Yes">Yes</option>
                <option value="No">No</option>
            </select></label>
            <label>Good with Strangers? *: <select name="good_with_strangers" required>
                <option value="Yes">Yes</option>
                <option value="No">No</option>
            </select></label>
            <label>Good with Dogs? *: <select name="good_with_dogs" required>
                <option value="Yes">Yes</option>
                <option value="No">No</option>
            </select></label>
            <label>Good with Cats? *: <select name="good_with_cats" required>
                <option value="Yes">Yes</option>
                <option value="No">No</option>
            </select></label>
            <label>Good with Kids? *: <select name="good_with_kids" required>
                <option value="Yes">Yes</option>
                <option value="No">No</option>
            </select></label>
            <label>Inside or Outside Dog? *: <select name="inside_outside" required>
                <option value="Inside">Inside</option>
                <option value="Outside">Outside</option>
            </select></label>
            <label>Why Are You Re-homing? Any Other Comments?: <textarea name="rehoming_reason"></textarea></label>
            <label>Picture of the Dog: <input type="file" name="dog_picture" accept="image/*"/></label>
            <button type= "submit"> Submit </button>
    </form>
     
        
    
    );

}


